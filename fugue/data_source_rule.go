package fugue

import (
	"context"
	"log"

	"github.com/fugue/fugue-client/client/custom_rules"
	"github.com/fugue/fugue-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRule() *schema.Resource {
	return &schema.Resource{
		Description: "`fugue_rule` data source can be used to retrieve information about a Fugue rule.",

		ReadContext: dataSourceRuleRead,

		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),

			"id": {
				Description: "The unique ID for this rule.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The name of this rule.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"cloud_provider": {
				Description: "The provider of this rule.",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"description": {
				Description: "The description of this rule.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"resource_type": {
				Description: "The resource type of this rule.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"rule_text": {
				Description: "The text of this rule.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"severity": {
				Description: "The severity of this rule.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"source": {
				Description: "The source of this rule.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"status": {
				Description: "The status of this rule.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceRuleCommonRead(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*models.CustomRule, diag.Diagnostics) {
	client := m.(*Client)

	var filtered []*models.CustomRule

	err := resource.RetryContext(context.Background(), EnvironmentRetryTimeout, func() *resource.RetryError {
		params := custom_rules.NewListCustomRulesParams()
		offset := int64(0)
		maxItems := int64(100)
		isTruncated := true

		params.Offset = &offset
		params.MaxItems = &maxItems

		for isTruncated {
			resp, err := client.CustomRules.ListCustomRules(params, client.Auth)
			if err != nil {
				log.Printf("[WARN] Get rule error: %s", err.Error())
				switch err.(type) {
				case *custom_rules.GetCustomRuleInternalServerError:
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

func dataSourceRuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	filtered, diags := dataSourceRuleCommonRead(ctx, d, m)
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
	if err := d.Set("description", result.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("resource_type", result.ResourceType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("rule_text", result.RuleText); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("severity", result.Severity); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("source", result.Source); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("status", result.Status); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func dataSourceRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	filtered, diags := dataSourceRuleCommonRead(ctx, d, m)
	if diags != nil {
		return diags
	}

	var ids []string
	for _, env := range filtered {
		ids = append(ids, env.ID)
	}

	d.SetId(dataSourceHashFilter("fugue_rules", d.Get("filter")))
	if err := d.Set("ids", ids); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
