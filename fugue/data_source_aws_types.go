package fugue

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/fugue/fugue-client/client/metadata"
	"github.com/fugue/fugue-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAwsTypes() *schema.Resource {
	return &schema.Resource{
		Description: "`fugue_aws_environment` data source can be used to retrieve information about a Fugue AWS environment.",
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
			"types": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceAwsTypesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	client := m.(*Client)

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

	var typeMetadata *models.ResourceTypeMetadata

	err := resource.Retry(EnvironmentRetryTimeout, func() *resource.RetryError {
		resp, err := client.Metadata.GetResourceTypes(params, client.Auth)
		if err != nil {
			log.Printf("[WARN] Get resource types error: %s, retrying", err.Error())
			switch err.(type) {
			case *metadata.GetResourceTypesInternalServerError:
				return resource.RetryableError(err)
			default:
				return resource.NonRetryableError(err)
			}
		}
		typeMetadata = resp.Payload
		return nil
	})
	if err != nil {
		return diag.FromErr(err)
	}

	var types []string
	for _, resourceTypeName := range typeMetadata.ResourceTypes {
		types = append(types, resourceTypeName)
	}
	if err := d.Set("types", types); err != nil {
		return diag.FromErr(err)
	}

	// Always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
