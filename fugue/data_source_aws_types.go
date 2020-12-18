package fugue

import (
	"context"
	"strconv"
	"time"

	"github.com/fugue/fugue-client/client/metadata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAwsTypes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAwsTypesRead,
		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "us-east-1",
			},
			"govcloud": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"types": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceAwsTypesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client, auth := getClient()
	params := metadata.NewGetResourceTypesParams()

	params.Provider = "aws"
	if d.Get("govcloud").(bool) {
		params.Provider = "aws_govcloud"
	}

	regionForTypes := "us-east-1"
	if region, ok := d.GetOk("region"); ok {
		regionForTypes = region.(string)
	}
	params.Region = &regionForTypes

	resp, err := client.Metadata.GetResourceTypes(params, auth)
	if err != nil {
		return diag.FromErr(err)
	}
	var types []string
	for _, resourceTypeName := range resp.Payload.ResourceTypes {
		types = append(types, resourceTypeName)
	}
	if err := d.Set("types", types); err != nil {
		return diag.FromErr(err)
	}

	// Always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
