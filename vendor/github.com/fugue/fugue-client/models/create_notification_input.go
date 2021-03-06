// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateNotificationInput Request for creating a new notification.
//
// swagger:model CreateNotificationInput
type CreateNotificationInput struct {

	// List of email address the notification is delivered to.
	Emails []string `json:"emails"`

	// List of environment ids the notification is attached to.
	Environments []string `json:"environments"`

	// List of events the notification is triggered on.
	Events []string `json:"events"`

	// Human readable name of the notification.
	Name string `json:"name,omitempty"`

	// AWS SNS topic arn the notification is delivered to.
	TopicArn string `json:"topic_arn,omitempty"`
}

// Validate validates this create notification input
func (m *CreateNotificationInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateNotificationInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateNotificationInput) UnmarshalBinary(b []byte) error {
	var res CreateNotificationInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
