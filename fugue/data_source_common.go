package fugue

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

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
		Description: "Filters.  The result is restricted to the intersection of the result of all filter.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Description: "Name of the field you want to apply this filter to.",
					Type:        schema.TypeString,
					Required:    true,
				},

				"values": {
					Description: "Filter to apply (using glob syntax).",
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

func dataSourceCheckFilterV(d *schema.ResourceData, filterName string, values []string) bool {
	if filters := d.Get("filter"); filters != nil {
		for _, filt := range filters.(*schema.Set).List() {
			if cfName, ok := filt.(map[string]interface{})["name"]; ok && cfName.(string) == filterName {
				if cfVals, ok := filt.(map[string]interface{})["values"]; ok {
					for _, val := range cfVals.([]interface{}) {
						for _, cv := range values {
							if glob.Glob(val.(string), cv) {
								return true
							}
						}
					}
					return false
				} else {
					continue
				}
			} else {
				continue
			}
		}
		return true
	} else {
		return false
	}
}

func dataSourceCheckFilter(d *schema.ResourceData, filterName string, value string) bool {
	return dataSourceCheckFilterV(d, filterName, []string{value})
}
