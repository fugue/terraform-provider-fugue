package fugue

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FUGUE_API_ID", nil),
			},
			"client_secret": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("FUGUE_API_SECRET", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"fugue_aws_environment": resourceAwsEnvironment(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"fugue_aws_types": dataSourceAwsTypes(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	var diags diag.Diagnostics

	clientID := d.Get("client_id").(string)
	clientSecret := d.Get("client_secret").(string)

	client, err := getFugueClient(clientID, clientSecret)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Fugue client",
			Detail:   err.Error(),
		})
		return nil, diags
	}

	return client, diags
}
