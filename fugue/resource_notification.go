package fugue

import (
	"context"
	"log"

	"github.com/fugue/fugue-client/client/notifications"
	"github.com/fugue/fugue-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var notificationAllowedEvents = map[string]bool{
	"compliance":  true,
	"drift":       true,
	"remediation": true,
}

func resourceNotification() *schema.Resource {
	return &schema.Resource{
		Description:   "`fugue_notification` manages a notification in Fugue.",
		CreateContext: resourceNotificationCreate,
		ReadContext:   resourceNotificationRead,
		UpdateContext: resourceNotificationUpdate,
		DeleteContext: resourceNotificationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID for this notification as generated by Fugue.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description:  "The name of the notification. Must be unique.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 250),
			},
			"emails": {
				Description: "The email addresses to be notified.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				MaxItems:    100,
			},
			"environments": {
				Description: "The environments to be monitored.",
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"events": {
				Description: "Event types to be monitored.",
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"topic_arn": {
				Description: "SNS topic ARN to be notified.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceNotificationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Client)

	name := d.Get("name").(string)
	emails := getStringSlice(d.Get("emails").([]interface{}))
	environments := getStringSlice(d.Get("environments").([]interface{}))
	events := getStringSlice(d.Get("events").([]interface{}))
	topicArn := d.Get("topic_arn").(string)

	for _, event := range events {
		if !notificationAllowedEvents[event] {
			return diag.Errorf("invalid event type: %s; expected: compliance, drift, remediation", event)
		}
	}

	params := notifications.NewCreateNotificationParams()
	params.Notification = &models.CreateNotificationInput{
		Name:         name,
		Emails:       emails,
		Environments: environments,
		Events:       events,
		TopicArn:     topicArn,
	}

	var notificationID string
	err := resource.RetryContext(ctx, EnvironmentRetryTimeout, func() *resource.RetryError {
		resp, err := client.Notifications.CreateNotification(params, client.Auth)
		if err != nil {
			switch err.(type) {
			case *notifications.CreateNotificationBadRequest,
				*notifications.CreateNotificationForbidden,
				*notifications.CreateNotificationUnauthorized:
				return resource.NonRetryableError(err)
			default:
				return resource.RetryableError(err)
			}
		}
		notificationID = resp.Payload.NotificationID
		return nil
	})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(notificationID)
	resourceNotificationRead(ctx, d, m)
	return diags
}

func resourceNotificationGetById(ctx context.Context, d *schema.ResourceData, m interface{}, id string) (*models.Notification, diag.Diagnostics) {

	client := m.(*Client)
	var result *models.Notification

	err := resource.RetryContext(ctx, EnvironmentRetryTimeout, func() *resource.RetryError {
		params := notifications.NewListNotificationsParams()
		offset := int64(0)
		maxItems := int64(100)
		isTruncated := true
		params.Offset = &offset
		params.MaxItems = &maxItems
		for isTruncated {
			resp, err := client.Notifications.ListNotifications(params, client.Auth)
			if err != nil {
				log.Printf("[WARN] List notifications error: %s", err.Error())
				switch err.(type) {
				case *notifications.ListNotificationsBadRequest,
					*notifications.ListNotificationsUnauthorized,
					*notifications.ListNotificationsForbidden:
					return resource.NonRetryableError(err)
				default:
					return resource.RetryableError(err)
				}
			}
			for _, notification := range resp.Payload.Items {
				if notification.NotificationID == id {
					result = notification
					return nil
				}
			}
			isTruncated = resp.Payload.IsTruncated
			offset = resp.Payload.NextOffset
		}
		return nil
	})
	if err != nil {
		return nil, diag.FromErr(err)
	}
	return result, nil
}

func resourceNotificationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	notification, diags := resourceNotificationGetById(ctx, d, m, d.Id())
	if diags != nil {
		return diags
	}
	if notification == nil {
		return diag.Errorf("Unable to locate notification with id '%s'.", d.Id())
	}
	if err := d.Set("name", notification.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("emails", notification.Emails); err != nil {
		return diag.FromErr(err)
	}
	var environmentIDs []string
	for envID := range notification.Environments {
		environmentIDs = append(environmentIDs, envID)
	}
	if err := d.Set("environments", environmentIDs); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("events", notification.Events); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("topic_arn", notification.TopicArn); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceNotificationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*Client)
	params := notifications.NewUpdateNotificationParams()
	params.NotificationID = d.Id()
	params.Notification = &models.UpdateNotificationInput{}

	params.Notification.Name = d.Get("name").(string)
	params.Notification.TopicArn = d.Get("topic_arn").(string)

	emails := []string{}
	if emailsList, ok := d.GetOk("emails"); ok {
		emails = expandStringSet(emailsList.(*schema.Set))
	}
	params.Notification.Emails = emails

	environments := []string{}
	if envList, ok := d.GetOk("environments"); ok {
		environments = expandStringSet(envList.(*schema.Set))
	}
	params.Notification.Environments = environments

	events := []string{}
	if eventsList, ok := d.GetOk("events"); ok {
		events = expandStringSet(eventsList.(*schema.Set))
	}
	params.Notification.Events = events

	for _, event := range events {
		if !notificationAllowedEvents[event] {
			return diag.Errorf("invalid event type: %s; expected: compliance, drift, remediation", event)
		}
	}

	err := resource.RetryContext(ctx, EnvironmentRetryTimeout, func() *resource.RetryError {
		_, err := client.Notifications.UpdateNotification(params, client.Auth)
		if err != nil {
			switch err.(type) {
			case *notifications.UpdateNotificationBadRequest,
				*notifications.UpdateNotificationForbidden,
				*notifications.UpdateNotificationUnauthorized,
				*notifications.UpdateNotificationNotFound:
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
	return resourceNotificationRead(ctx, d, m)
}

func resourceNotificationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*Client)
	params := notifications.NewDeleteNotificationParams()
	params.NotificationID = d.Id()

	err := resource.RetryContext(ctx, EnvironmentRetryTimeout, func() *resource.RetryError {
		_, err := client.Notifications.DeleteNotification(params, client.Auth)
		if err != nil {
			switch err.(type) {
			case *notifications.DeleteNotificationForbidden,
				*notifications.DeleteNotificationUnauthorized,
				*notifications.DeleteNotificationNotFound:
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
