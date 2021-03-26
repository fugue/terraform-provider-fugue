package fugue

import (
	"context"
	"log"
	"sort"

	"github.com/fugue/fugue-client/client/environments"
	"github.com/fugue/fugue-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGoogleEnvironment() *schema.Resource {
	return &schema.Resource{
		Description:   "`fugue_google_environment` manages an Environment in Fugue corresponding to one Google account.",
		CreateContext: resourceGoogleEnvironmentCreate,
		ReadContext:   resourceGoogleEnvironmentRead,
		UpdateContext: resourceGoogleEnvironmentUpdate,
		DeleteContext: resourceGoogleEnvironmentDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The unique ID for this environment as generated by Fugue.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The name for the environment.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"service_account_email": {
				Description: "The Google Service Account email used to provide Fugue secure access to the Google account.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "The google project ID (not required).",
				Type:        schema.TypeString,
				Required:    false,
				Optional:    true,
			},
			"compliance_families": {
				Description: `The set of compliance families to enable in this environment, e.g. ["CIS-Google_v1.1.0"].`,
				Type:        schema.TypeSet,
				Optional:    true,
				MaxItems:    100,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"scan_interval": {
				Description: "Controls the time in seconds between scheduled scans of this environment.",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     86400,
			},
			"scan_schedule_enabled": {
				Description: "Controls whether this environment is scanned on a schedule.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"scan_status": {
				Description: "Indicates whether a scan on this environment is currently running.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceGoogleEnvironmentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	client := m.(*Client)

	scanInterval := int64(0)
	if scanIntervalSetting, ok := d.GetOk("scan_interval"); ok {
		scanInterval = int64(scanIntervalSetting.(int))
	}
	scanScheduleEnabled := d.Get("scan_schedule_enabled").(bool)
	var scanIntervalPtr *int64
	if scanScheduleEnabled {
		scanIntervalPtr = &scanInterval
	}
	complianceFamilies := []string{}
	if complianceFamiliesSetting, ok := d.GetOk("compliance_families"); ok {
		complianceFamilies = expandStringSet(complianceFamiliesSetting.(*schema.Set))
	}

	params := environments.NewCreateEnvironmentParams()
	params.Environment = &models.CreateEnvironmentInput{
		ComplianceFamilies:  complianceFamilies,
		Name:                d.Get("name").(string),
		Provider:            "google",
		ScanInterval:        scanIntervalPtr,
		ScanScheduleEnabled: &scanScheduleEnabled,
	}

	projectId := ""
	if projectIdSetting, ok := d.GetOk("project_id"); ok {
		projectId = projectIdSetting.(string)
	}

	providerOpts := &models.ProviderOptionsGoogle{
		ProjectID:           projectId,
		ServiceAccountEmail: d.Get("service_account_email").(string),
	}

	params.Environment.ProviderOptions = &models.ProviderOptions{Google: providerOpts}

	var environmentID string

	err := resource.RetryContext(ctx, EnvironmentRetryTimeout, func() *resource.RetryError {
		resp, err := client.Environments.CreateEnvironment(params, client.Auth)
		if err != nil {
			log.Printf("[WARN] Create environment error: %s", err.Error())
			switch err.(type) {
			case *environments.CreateEnvironmentInternalServerError:
				return resource.RetryableError(err)
			default:
				return resource.NonRetryableError(err)
			}
		}
		environmentID = resp.Payload.ID
		return nil
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(environmentID)
	resourceGoogleEnvironmentRead(ctx, d, m)
	return diags
}

func resourceGoogleEnvironmentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	client := m.(*Client)

	params := environments.NewGetEnvironmentParams()
	params.EnvironmentID = d.Id()
	var env *models.EnvironmentWithSummary

	err := resource.RetryContext(ctx, EnvironmentRetryTimeout, func() *resource.RetryError {
		resp, err := client.Environments.GetEnvironment(params, client.Auth)
		if err != nil {
			log.Printf("[WARN] Get environment error: %s", err.Error())
			switch err.(type) {
			case *environments.GetEnvironmentInternalServerError:
				return resource.RetryableError(err)
			default:
				return resource.NonRetryableError(err)
			}
		}
		env = resp.Payload
		return nil
	})
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", env.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("scan_interval", int(env.ScanInterval)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("scan_schedule_enabled", env.ScanScheduleEnabled); err != nil {
		return diag.FromErr(err)
	}

	complianceFamilies := env.ComplianceFamilies
	sort.Strings(complianceFamilies)
	if err := d.Set("compliance_families", complianceFamilies); err != nil {
		return diag.FromErr(err)
	}

	providerOpts := env.ProviderOptions.Google
	projectId := providerOpts.ProjectID
	if err := d.Set("project_id", projectId); err != nil {
		return diag.FromErr(err)
	}
	serviceAccountEmail := providerOpts.ServiceAccountEmail
	if err := d.Set("service_account_email", serviceAccountEmail); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("scan_status", env.ScanStatus); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceGoogleEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*Client)
	params := environments.NewUpdateEnvironmentParams()
	params.EnvironmentID = d.Id()
	params.Environment = &models.UpdateEnvironmentInput{}

	providerOptsInput := &models.ProviderOptionsUpdateInput{}
	providerOptsInput.Google = &models.ProviderOptionsGoogleUpdateInput{}

	if d.HasChange("name") {
		params.Environment.Name = d.Get("name").(string)
	}

	if d.HasChange("service_account_email") {
		providerOptsInput.Google.ServiceAccountEmail = d.Get("service_account_email").(string)
		params.Environment.ProviderOptions = providerOptsInput
	}

	if d.HasChange("compliance_families") {
		complianceFamilies := []string{}
		if complianceFamiliesSetting, ok := d.GetOk("compliance_families"); ok {
			complianceFamilies = expandStringSet(complianceFamiliesSetting.(*schema.Set))
		}
		params.Environment.ComplianceFamilies = complianceFamilies
	}

	if d.HasChange("scan_interval") {
		scanInterval := int64(0)
		if scanIntervalSetting, ok := d.GetOk("scan_interval"); ok {
			scanInterval = int64(scanIntervalSetting.(int))
		}
		params.Environment.ScanInterval = scanInterval
	}

	if d.HasChange("scan_schedule_enabled") {
		scanScheduleEnabled := d.Get("scan_schedule_enabled").(bool)
		params.Environment.ScanScheduleEnabled = &scanScheduleEnabled
	}

	err := resource.RetryContext(ctx, EnvironmentRetryTimeout, func() *resource.RetryError {
		_, err := client.Environments.UpdateEnvironment(params, client.Auth)
		if err != nil {
			log.Printf("[WARN] Update environment error: %s", err.Error())
			switch err.(type) {
			case *environments.UpdateEnvironmentInternalServerError:
				return resource.RetryableError(err)
			default:
				return resource.NonRetryableError(err)
			}
		}
		return nil
	})
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceGoogleEnvironmentRead(ctx, d, m)
}

func resourceGoogleEnvironmentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*Client)
	params := environments.NewDeleteEnvironmentParams()
	params.EnvironmentID = d.Id()

	err := resource.RetryContext(ctx, EnvironmentRetryTimeout, func() *resource.RetryError {
		_, err := client.Environments.DeleteEnvironment(params, client.Auth)
		if err != nil {
			log.Printf("[WARN] Delete environment error: %s", err.Error())
			switch err.(type) {
			case *environments.DeleteEnvironmentInternalServerError:
				return resource.RetryableError(err)
			default:
				return resource.NonRetryableError(err)
			}
		}
		return nil
	})
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
