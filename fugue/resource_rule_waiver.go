package fugue

import (
	"context"
	"fmt"

	"github.com/fugue/fugue-client/client/rule_waivers"
	"github.com/fugue/fugue-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRuleWaiver() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRuleWaiverCreate,
		ReadContext:   resourceRuleWaiverRead,
		UpdateContext: resourceRuleWaiverUpdate,
		DeleteContext: resourceRuleWaiverDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Required: true,
			},
			"environment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_provider": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rule_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRuleWaiverCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	client := m.(*Client)

	name := d.Get("name").(string)
	comment := d.Get("comment").(string)
	envID := d.Get("environment_id").(string)
	ruleID := d.Get("rule_id").(string)
	resourceProvider := d.Get("resource_provider").(string)
	resourceType := d.Get("resource_type").(string)
	resourceID := d.Get("resource_id").(string)

	validRule, err := isValidRuleID(client, ruleID)
	if err != nil {
		return diag.FromErr(err)
	}
	if !validRule {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Invalid rule ID specified",
			Detail:   fmt.Sprintf("'%s' is not a valid rule ID.", ruleID),
		})
		return diags
	}

	params := rule_waivers.NewCreateRuleWaiverParams()
	params.Input = &models.CreateRuleWaiverInput{
		Name:             &name,
		Comment:          comment,
		EnvironmentID:    &envID,
		RuleID:           &ruleID,
		ResourceProvider: &resourceProvider,
		ResourceType:     &resourceType,
		ResourceID:       &resourceID,
	}

	var waiverID string

	err = resource.Retry(EnvironmentRetryTimeout, func() *resource.RetryError {
		resp, err := client.RuleWaivers.CreateRuleWaiver(params, client.Auth)
		if err != nil {
			switch err.(type) {
			case *rule_waivers.CreateRuleWaiverForbidden:
				return resource.NonRetryableError(err)
			case *rule_waivers.CreateRuleWaiverUnauthorized:
				return resource.NonRetryableError(err)
			default:
				return resource.RetryableError(err)
			}
		}
		waiverID = *resp.Payload.ID
		return nil
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(waiverID)
	resourceRuleWaiverRead(ctx, d, m)
	return diags
}

func resourceRuleWaiverRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	client := m.(*Client)

	params := rule_waivers.NewGetRuleWaiverParams()
	params.RuleWaiverID = d.Id()

	var waiver *models.RuleWaiver

	err := resource.Retry(EnvironmentRetryTimeout, func() *resource.RetryError {
		resp, err := client.RuleWaivers.GetRuleWaiver(params, client.Auth)
		if err != nil {
			switch err.(type) {
			case *rule_waivers.GetRuleWaiverForbidden:
				return resource.NonRetryableError(err)
			case *rule_waivers.GetRuleWaiverUnauthorized:
				return resource.NonRetryableError(err)
			default:
				return resource.RetryableError(err)
			}
		}
		waiver = resp.Payload
		return nil
	})
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", waiver.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("comment", waiver.Comment); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("environment_id", waiver.EnvironmentID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("resource_provider", waiver.ResourceProvider); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("resource_type", waiver.ResourceType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("resource_id", waiver.ResourceID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("rule_id", waiver.RuleID); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceRuleWaiverUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*Client)

	if !d.HasChange("name") && !d.HasChange("comment") {
		return resourceRuleWaiverRead(ctx, d, m)
	}

	params := rule_waivers.NewUpdateRuleWaiverParams()
	params.RuleWaiverID = d.Id()
	params.Input = &models.UpdateRuleWaiverInput{}
	params.Input.Name = d.Get("name").(string)
	params.Input.Comment = d.Get("comment").(string)

	err := resource.Retry(EnvironmentRetryTimeout, func() *resource.RetryError {
		_, err := client.RuleWaivers.UpdateRuleWaiver(params, client.Auth)
		if err != nil {
			switch err.(type) {
			case *rule_waivers.UpdateRuleWaiverForbidden:
				return resource.NonRetryableError(err)
			case *rule_waivers.UpdateRuleWaiverUnauthorized:
				return resource.NonRetryableError(err)
			case *rule_waivers.UpdateRuleWaiverNotFound:
				return resource.NonRetryableError(err)
			default:
				return resource.RetryableError(err)
			}
		}
		return nil
	})
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceRuleWaiverRead(ctx, d, m)
}

func resourceRuleWaiverDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*Client)

	params := rule_waivers.NewDeleteRuleWaiverParams()
	params.RuleWaiverID = d.Id()

	err := resource.Retry(EnvironmentRetryTimeout, func() *resource.RetryError {
		_, err := client.RuleWaivers.DeleteRuleWaiver(params, client.Auth)
		if err != nil {
			switch err.(type) {
			case *rule_waivers.DeleteRuleWaiverForbidden:
				return resource.NonRetryableError(err)
			case *rule_waivers.DeleteRuleWaiverUnauthorized:
				return resource.NonRetryableError(err)
			case *rule_waivers.DeleteRuleWaiverNotFound:
				return resource.NonRetryableError(err)
			default:
				return resource.RetryableError(err)
			}
		}
		return nil
	})
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
