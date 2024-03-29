// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Notification Describes configuration of a notification.
//
// swagger:model Notification
type Notification struct {

	// The date and time the notification was created.
	// CreatedAt

	// Principal the created the notification.
	CreatedBy string `json:"created_by,omitempty"`

	// List of email address the notification is delivered to.
	Emails []string `json:"emails"`

	// List of maps from environment id to name the notification is attached to.
	Environments map[string]string `json:"environments"`

	// List of events the notification is triggered on.
	Events []string `json:"events"`

	// Last error recorded while processing notification. If the last notification processed had no error this field will be empty.
	LastError string `json:"last_error,omitempty"`

	// Human readable name of the notification.
	Name string `json:"name,omitempty"`

	// ID of the notification.
	NotificationID string `json:"notification_id,omitempty"`

	// AWS SNS topic arn the notification is delivered to.
	TopicArn string `json:"topic_arn,omitempty"`

	// AWS The date and time the notification was last updated.
	// UpdatedAt

	// Principal that last updated the notification.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// Validate validates this notification
func (m *Notification) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Notification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Notification) UnmarshalBinary(b []byte) error {
	var res Notification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
