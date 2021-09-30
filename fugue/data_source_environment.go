package fugue

import (
	"context"
	"log"

	"github.com/fugue/fugue-client/client/environments"
	"github.com/fugue/fugue-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEnvironment() *schema.Resource {
	return &schema.Resource{
		Description: "`fugue_environment` data source can be used to retrieve information about a Fugue environment.",

		ReadContext: dataSourceEnvironmentRead,

		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),

			"id": {
				Description: "The unique ID for this environment.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The name of this environment.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"cloud_provider": {
				Description: "The provider of this environment.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceEnvironmentCommonRead(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*models.Environment, diag.Diagnostics) {
	client := m.(*Client)

	var filtered []*models.Environment
	filters := getDataSourceFilters(d)
	filtersJSON, err := getQueryJSON(filters)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	err = resource.RetryContext(context.Background(), EnvironmentRetryTimeout, func() *resource.RetryError {
		params := environments.NewListEnvironmentsParams()
		offset := int64(0)
		maxItems := int64(100)
		isTruncated := true

		params.Offset = &offset
		params.MaxItems = &maxItems
		if len(filters) > 0 {
			params.Query = &filtersJSON
			log.Printf("[INFO] XXX Query: %+v", params.Query)
		} else {
			log.Printf("[INFO] XXX NO QUERY")
		}

		for isTruncated {
			resp, err := client.Environments.ListEnvironments(params, client.Auth)
			if err != nil {
				log.Printf("[WARN] Get environment error: %s", err.Error())
				switch err.(type) {
				case *environments.GetEnvironmentInternalServerError:
					return resource.RetryableError(err)
				default:
					return resource.NonRetryableError(err)
				}
			}

			for _, env := range resp.Payload.Items {
				if !dataSourceCheckFilter(d, "name", env.Name) {
					continue
				}
				if !dataSourceCheckFilter(d, "id", env.ID) {
					continue
				}
				if !dataSourceCheckFilter(d, "cloud_provider", env.Provider) {
					continue
				}

				filtered = append(filtered, env)
			}

			isTruncated = resp.Payload.IsTruncated
			offset = resp.Payload.NextOffset
		}
		return nil
	})
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return filtered, nil
}

func dataSourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	filtered, diags := dataSourceEnvironmentCommonRead(ctx, d, m)
	if diags != nil {
		return diags
	}

	if diags := dataSourceVerifySingleResult(len(filtered)); diags != nil {
		return diags
	}

	result := filtered[0]

	d.SetId(result.ID)
	if err := d.Set("name", result.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cloud_provider", result.Provider); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func dataSourceEnvironmentsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	filtered, diags := dataSourceEnvironmentCommonRead(ctx, d, m)
	if diags != nil {
		return diags
	}

	var ids []string
	for _, env := range filtered {
		ids = append(ids, env.ID)
	}

	d.SetId(dataSourceHashFilter("fugue_environments", d.Get("filter")))
	if err := d.Set("ids", ids); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
