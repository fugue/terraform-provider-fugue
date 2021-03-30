package fugue

import (
	"strings"
	"time"

	"github.com/fugue/fugue-client/client/custom_rules"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// EnvironmentRetryTimeout defines the maximum time to retry on
	// errors when changing an environment
	EnvironmentRetryTimeout = 30 * time.Second
)

func getStringSlice(input []interface{}) []string {
	items := make([]string, 0, len(input))
	for _, item := range input {
		items = append(items, item.(string))
	}
	return items
}

func getSingleAwsRegion(provider string, regions []string) string {
	if len(regions) == 0 || regions[0] == "*" {
		if provider == "aws_govcloud" {
			return "us-gov-west-1"
		}
		return "us-east-1"
	}
	return regions[0]
}

func expandStringSet(configured *schema.Set) []string {
	return expandStringList(configured.List())
}

// Takes the result of flatmap.Expand for an array of strings
// and returns a []*string
func expandStringList(configured []interface{}) []string {
	vs := make([]string, 0, len(configured))
	for _, v := range configured {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, val)
		}
	}
	return vs
}

func isValidRuleID(client *Client, ruleID string) (bool, error) {

	// Fugue managed rules have IDs starting with "FG_"
	if strings.HasPrefix(ruleID, "FG_") {
		return true, nil
	}

	// Otherwise, this must be referring to a custom rule.
	// Confirm a custom rule with this ID exists.
	params := custom_rules.NewGetCustomRuleParams()
	params.RuleID = ruleID

	var ruleFound bool

	err := resource.Retry(EnvironmentRetryTimeout, func() *resource.RetryError {
		_, err := client.CustomRules.GetCustomRule(params, client.Auth)
		if err != nil {
			switch err.(type) {
			case *custom_rules.GetCustomRuleNotFound:
				ruleFound = false
				return nil
			case *custom_rules.GetCustomRuleForbidden:
				return resource.NonRetryableError(err)
			case *custom_rules.GetCustomRuleUnauthorized:
				return resource.NonRetryableError(err)
			default:
				return resource.RetryableError(err)
			}
		}
		ruleFound = true
		return nil
	})
	if err != nil {
		return false, err
	}

	return ruleFound, nil
}
