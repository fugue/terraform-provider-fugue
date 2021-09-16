package fugue

import (
	"context"
	"log"

	"github.com/fugue/fugue-client/client/rule_waivers"
	"github.com/fugue/fugue-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRuleWaiver() *schema.Resource {
	return &schema.Resource{
		Description: "`fugue_rule_waiver` data source can be used to retrieve information about a Fugue rule waiver.",

		ReadContext: dataSourceRuleWaiverRead,

		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),

			"id": {
				Description: "The unique ID for this waiver.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The name of this waiver.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"environment_id": {
				Description: "The ID of the environment this waiver is applied to.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"comment": {
				Description: "The comment associated with this waiver.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"rule_id": {
				Description: "The ID of the rule being waived.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"resource_id": {
				Description: "The ID of the resource the rule is being waived for.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"resource_provider": {
				Description: "The provider of the resource the rule is being waived for.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"resource_type": {
				Description: "The type of the resource the rule is being waived for.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"resource_tag": {
				Description: "The type of the resource the rule is being waived for.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceRuleWaiverCommonRead(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*models.RuleWaiver, diag.Diagnostics) {
	client := m.(*Client)

	var filtered []*models.RuleWaiver

	err := resource.RetryContext(context.Background(), EnvironmentRetryTimeout, func() *resource.RetryError {
		params := rule_waivers.NewListRuleWaiversParams()
		offset := int64(0)
		maxItems := int64(100)
		isTruncated := true

		params.Offset = &offset
		params.MaxItems = &maxItems

		for isTruncated {
			resp, err := client.RuleWaivers.ListRuleWaivers(params, client.Auth)
			if err != nil {
				log.Printf("[WARN] Get waiver error: %s", err.Error())
				switch err.(type) {
				case *rule_waivers.GetRuleWaiverInternalServerError:
					return resource.RetryableError(err)
				default:
					return resource.NonRetryableError(err)
				}
			}

			for _, waiver := range resp.Payload.Items {
				if !dataSourceCheckFilterP(d, "name", waiver.Name) {
					continue
				}
				if !dataSourceCheckFilterP(d, "id", waiver.ID) {
					continue
				}
				if !dataSourceCheckFilterP(d, "environment_id", waiver.EnvironmentID) {
					continue
				}
				if !dataSourceCheckFilter(d, "comment", waiver.Comment) {
					continue
				}
				if !dataSourceCheckFilterP(d, "rule_id", waiver.RuleID) {
					continue
				}
				if !dataSourceCheckFilterP(d, "resource_id", waiver.ResourceID) {
					continue
				}
				if !dataSourceCheckFilterP(d, "resource_provider", waiver.ResourceProvider) {
					continue
				}
				if !dataSourceCheckFilterP(d, "resource_type", waiver.ResourceType) {
					continue
				}
				if !dataSourceCheckFilter(d, "resource_tag", waiver.ResourceTag) {
					continue
				}

				filtered = append(filtered, waiver)
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

func dataSourceRuleWaiverRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	filtered, diags := dataSourceRuleWaiverCommonRead(ctx, d, m)
	if diags != nil {
		return diags
	}

	if diags := dataSourceVerifySingleResult(len(filtered)); diags != nil {
		return diags
	}

	result := filtered[0]

	d.SetId(*result.ID)
	if err := d.Set("name", result.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("environment_id", result.EnvironmentID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("comment", result.Comment); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("rule_id", result.RuleID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("resource_id", result.ResourceID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("resource_provider", result.ResourceProvider); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("resource_type", result.ResourceType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("resource_tag", result.ResourceTag); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func dataSourceRuleWaiversRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	filtered, diags := dataSourceRuleWaiverCommonRead(ctx, d, m)
	if diags != nil {
		return diags
	}

	var ids []string
	for _, waiver := range filtered {
		ids = append(ids, *waiver.ID)
	}

	d.SetId(dataSourceHashFilter("fugue_waivers", d.Get("filter")))
	if err := d.Set("ids", ids); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
