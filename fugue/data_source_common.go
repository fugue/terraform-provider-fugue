package fugue

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ryanuber/go-glob"
)

func dataSourceFiltersSchema() *schema.Schema {
	/* Copied from https://github.com/hashicorp/terraform-provider-aws/blob/ba0539575904c75afb86c9d6bd450d2c21a62556/aws/data_source_aws_common_schema.go#L25-L44
	* (also MPL-licensed) */
	return &schema.Schema{
		Type:        schema.TypeSet,
		Optional:    true,
		Description: "Filters. The result is restricted to the intersection of the result of all filters.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Description: "Name of the field you want to apply this filter to.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"values": {
					Description: "Value to match",
					Type:        schema.TypeList,
					Required:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
				},
			},
		},
	}
}

type dataSourcePluralInfo struct {
	SingularEquivalent string
	Description        string
}

func dataSourcePluralSchema(info dataSourcePluralInfo, readFunc schema.ReadContextFunc) *schema.Resource {
	var description string
	if info.Description != "" {
		description = info.Description
	} else if info.SingularEquivalent != "" {
		description = fmt.Sprintf("Multi-result version of `%s` where the results are placed in the `ids` element.", info.SingularEquivalent)
	}
	return &schema.Resource{
		Description: description,
		ReadContext: readFunc,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"ids": {
				Description: "List of the ID of each result.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceVerifySingleResult(filteredLength int) diag.Diagnostics {
	switch filteredLength {
	case 0:
		return diag.Errorf("Query returned zero results.")
	case 1:
		return nil
	default:
		return diag.Errorf("Query returned multiple results.")
	}
}

func dataSourceHashFilter(resourceType string, filters interface{}) string {
	var filterData []interface{} = nil
	if filters != nil {
		filterData = filters.(*schema.Set).List()
	}
	input := map[string]interface{}{resourceType: filterData}
	hv := sha256.Sum256([]byte(fmt.Sprint(input)))
	dst := make([]byte, hex.EncodedLen(len(hv)))
	hex.Encode(dst, hv[:])
	return string(dst)
}

// Transforms filters in the given resource into a simple map of attribute name
// to a slice of attribute values to filter on. For example, a filter with name
// "name" and values "staging" and "prod" would return a map in this form:
// { "name": ["staging", "prod"] }
func getDataSourceFilters(d *schema.ResourceData) map[string][]string {
	result := map[string][]string{}
	if filters := d.Get("filter"); filters != nil {
		for _, filter := range filters.(*schema.Set).List() {
			filterMap, ok := filter.(map[string]interface{})
			if !ok {
				continue
			}
			filterName, ok := filterMap["name"].(string)
			if !ok {
				continue
			}
			var values []string
			for _, value := range filterMap["values"].([]interface{}) {
				if valueStr, ok := value.(string); ok {
					values = append(values, valueStr)
				}
			}
			result[filterName] = values
		}
	}
	return result
}

// Given a set of name:values filters, build a JSON-formatted query string that
// is compatible with the Fugue API "query" parameter. This looks something like
// ["name:staging,prod", "id:1234"] if searching on both name and ID. If
// wildcards ("*") are present in the filter strings, the longest non-wildcard
// segment is passed to the server via the query parameter. The rest of the
// matching is then done client-side. The server doesn't support wildcards.
func getQueryJSON(filters map[string][]string, maxValues int) (string, error) {
	var filterStrings []string
	for filterName, filterValue := range filters {
		// Not all API endpoints support querying for multiple values.
		if maxValues > 0 && len(filterValue) > maxValues {
			if maxValues == 1 {
				return "", fmt.Errorf("filter is limited to 1 value; got %d",
					len(filterValue))
			} else {
				return "", fmt.Errorf("filter is limited to %d values; got %d",
					maxValues, len(filterValue))
			}
		}
		// Extract the largest section from each glob so we can use glob locally
		// but stiff filter out most of the entries on the server.
		serverFilters := make([]string, 0, len(filterValue))
		for filter := range filterValue {
			components := strings.Split(filterValue[filter], "*")
			var longestComponent string
			longestComponentLength := -1
			for component := range components {
				if len(components[component]) > longestComponentLength {
					longestComponent = components[component]
					longestComponentLength = len(longestComponent)
				}
			}
			serverFilters = append(serverFilters, longestComponent)
		}
		filterStrings = append(
			filterStrings,
			fmt.Sprintf("%s:%s", filterName, strings.Join(serverFilters, ",")),
		)
	}
	if len(filterStrings) == 0 {
		return "", nil
	}
	query, err := json.Marshal(filterStrings)
	if err != nil {
		return "", err
	}
	return string(query), nil
}

// Checks if the given attribute name and value are matched by the configured
// filters. This is a case insensitive match using glob patterns in the filters.
func dataSourceCheckFilter(filters map[string][]string, name, value string) bool {
	patterns, ok := filters[name]
	if !ok || len(patterns) == 0 {
		// No filter is set for this attribute: match!
		return true
	}
	valueLower := strings.ToLower(value)
	for _, pattern := range patterns {
		if glob.Glob(strings.ToLower(pattern), valueLower) {
			// Case-insensitive glob match!
			return true
		}
	}
	// None of the patterns matched
	return false
}

func dataSourceCheckFilterP(filters map[string][]string, filterName string, value *string) bool {
	if value == nil {
		def := ""
		value = &def
	}
	return dataSourceCheckFilter(filters, filterName, *value)
}

func copyFilterMap(filters map[string][]string) map[string][]string {
	result := make(map[string][]string, len(filters))
	for k, v := range filters {
		result[k] = v
	}
	return result
}

// Adjusts filters for server-side application. Currently this only applies one
// transformation, which is to rename "cloud_provider" to "provider". The latter
// had to be avoided as an attribute name here in the Terraform provider since
// that is a reserved attribute name.
func serverSideFilters(filters map[string][]string) map[string][]string {
	modified := copyFilterMap(filters)
	if values, ok := modified["cloud_provider"]; ok {
		delete(modified, "cloud_provider")
		modified["provider"] = values
	}
	return modified
}
