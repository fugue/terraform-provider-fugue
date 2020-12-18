package fugue

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"fugue_aws_environment": resourceAwsEnvironment(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"fugue_aws_types": dataSourceAwsTypes(),
		},
	}
}
