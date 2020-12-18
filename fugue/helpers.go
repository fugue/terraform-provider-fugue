package fugue

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
