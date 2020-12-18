package fugue

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/fugue/fugue-client/client/environments"
	"github.com/fugue/fugue-client/client/metadata"
	"github.com/fugue/fugue-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAwsEnvironment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAwsEnvironmentCreate,
		ReadContext:   resourceAwsEnvironmentRead,
		UpdateContext: resourceAwsEnvironmentUpdate,
		DeleteContext: resourceAwsEnvironmentDelete,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"govcloud": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"regions": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 100,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"role_arn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_types": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1000,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"compliance_families": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 100,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"scan_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  86400,
			},
			"scan_schedule_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"scan_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAwsEnvironmentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	client := m.(*Client)

	regions := []string{"*"}
	if regionsSetting, ok := d.GetOk("regions"); ok {
		regions = expandStringSet(regionsSetting.(*schema.Set))
		if len(regions) == 0 {
			return diag.FromErr(errors.New("Must specify a region"))
		}
	}

	provider := "aws"
	if d.Get("govcloud").(bool) {
		provider = "aws_govcloud"
	}

	singleRegion := getSingleAwsRegion(provider, regions)

	var surveyTypes []string
	if resourceTypesSetting, ok := d.GetOk("resource_types"); ok {
		surveyTypes = expandStringSet(resourceTypesSetting.(*schema.Set))
	}
	if len(surveyTypes) == 0 {
		// Default to all available types
		getTypesParams := metadata.NewGetResourceTypesParams()
		getTypesParams.Provider = provider
		getTypesParams.Region = &singleRegion
		resp, err := client.Metadata.GetResourceTypes(getTypesParams, client.Auth)
		if err != nil {
			return diag.FromErr(err)
		}
		surveyTypes = resp.Payload.ResourceTypes
	}

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
		Provider:            provider,
		ScanInterval:        scanIntervalPtr,
		SurveyResourceTypes: surveyTypes,
		ScanScheduleEnabled: &scanScheduleEnabled,
	}

	providerOpts := &models.ProviderOptionsAws{
		Regions: regions,
		RoleArn: d.Get("role_arn").(string),
	}
	if provider == "aws_govcloud" {
		params.Environment.ProviderOptions = &models.ProviderOptions{AwsGovcloud: providerOpts}
	} else {
		params.Environment.ProviderOptions = &models.ProviderOptions{Aws: providerOpts}
	}

	resp, err := client.Environments.CreateEnvironment(params, client.Auth)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.Payload.ID)

	resourceAwsEnvironmentRead(ctx, d, m)

	return diags
}

func resourceAwsEnvironmentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	client := m.(*Client)

	params := environments.NewGetEnvironmentParams()
	params.EnvironmentID = d.Id()

	resp, err := client.Environments.GetEnvironment(params, client.Auth)
	if err != nil {
		return diag.FromErr(err)
	}
	env := resp.Payload

	if err := d.Set("name", env.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("scan_interval", int(env.ScanInterval)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("scan_schedule_enabled", env.ScanScheduleEnabled); err != nil {
		return diag.FromErr(err)
	}

	resourceTypes := env.SurveyResourceTypes
	sort.Strings(resourceTypes)
	if err := d.Set("resource_types", resourceTypes); err != nil {
		return diag.FromErr(err)
	}

	complianceFamilies := env.ComplianceFamilies
	sort.Strings(complianceFamilies)
	if err := d.Set("compliance_families", complianceFamilies); err != nil {
		return diag.FromErr(err)
	}

	var providerOpts *models.ProviderOptionsAws

	if strings.ToLower(env.Provider) == "aws_govcloud" {
		if err := d.Set("govcloud", true); err != nil {
			return diag.FromErr(err)
		}
		providerOpts = env.ProviderOptions.AwsGovcloud
	} else {
		if err := d.Set("govcloud", false); err != nil {
			return diag.FromErr(err)
		}
		providerOpts = env.ProviderOptions.Aws
	}

	roleArn := providerOpts.RoleArn
	regions := providerOpts.Regions
	sort.Strings(regions)

	if err := d.Set("regions", regions); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("role_arn", roleArn); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("scan_status", env.ScanStatus); err != nil {
		return diag.FromErr(err)
	}

	// diags = append(diags, diag.Diagnostic{
	// 	Severity: diag.Warning,
	// 	Summary:  "Fugue Debugging",
	// 	Detail:   fmt.Sprintf("Regions: %+v", regions),
	// })

	return diags
}

func resourceAwsEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	client := m.(*Client)

	params := environments.NewUpdateEnvironmentParams()
	params.EnvironmentID = d.Id()
	params.Environment = &models.UpdateEnvironmentInput{}

	providerOptsInput := &models.ProviderOptionsUpdateInput{}
	providerOptsInput.Aws = &models.ProviderOptionsAwsUpdateInput{}

	if d.HasChange("name") {
		params.Environment.Name = d.Get("name").(string)
	}
	if d.HasChange("regions") {
		regions := []string{"*"}
		if regionsSetting, ok := d.GetOk("regions"); ok {
			regions = expandStringSet(regionsSetting.(*schema.Set))
			if len(regions) == 0 {
				return diag.FromErr(errors.New("Must specify a region"))
			}
		}
		providerOptsInput.Aws.Regions = regions
		params.Environment.ProviderOptions = providerOptsInput
	}
	if d.HasChange("role_arn") {
		providerOptsInput.Aws.RoleArn = d.Get("role_arn").(string)
		params.Environment.ProviderOptions = providerOptsInput
	}
	if d.HasChange("resource_types") {
		if resourceTypesSetting, ok := d.GetOk("resource_types"); ok {
			params.Environment.SurveyResourceTypes = expandStringSet(resourceTypesSetting.(*schema.Set))
		}
	}

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Fugue Debugging",
		Detail:   fmt.Sprintf("Has Change? %+v", d.HasChange("compliance_families")),
	})

	if d.HasChange("compliance_families") {
		complianceFamilies := []string{}
		if complianceFamiliesSetting, ok := d.GetOk("compliance_families"); ok {
			complianceFamilies = expandStringSet(complianceFamiliesSetting.(*schema.Set))
		}
		params.Environment.ComplianceFamilies = complianceFamilies

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Fugue Debugging",
			Detail:   fmt.Sprintf("Up families: %+v", complianceFamilies),
		})
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

	_, err := client.Environments.UpdateEnvironment(params, client.Auth)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceAwsEnvironmentRead(ctx, d, m)
}

func resourceAwsEnvironmentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	client := m.(*Client)

	params := environments.NewDeleteEnvironmentParams()
	params.EnvironmentID = d.Id()

	_, err := client.Environments.DeleteEnvironment(params, client.Auth)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
