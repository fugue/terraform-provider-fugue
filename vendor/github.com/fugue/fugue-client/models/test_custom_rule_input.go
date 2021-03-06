// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TestCustomRuleInput Input request for testing a custom rule.
//
// swagger:model TestCustomRuleInput
type TestCustomRuleInput struct {

	// Resource type to which the custom rule applies
	ResourceType string `json:"resource_type,omitempty"`

	// The rego source code for the rule
	// Required: true
	RuleText *string `json:"rule_text"`

	// Scan to test the custom rule with
	// Required: true
	ScanID *string `json:"scan_id"`
}

// Validate validates this test custom rule input
func (m *TestCustomRuleInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRuleText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestCustomRuleInput) validateRuleText(formats strfmt.Registry) error {

	if err := validate.Required("rule_text", "body", m.RuleText); err != nil {
		return err
	}

	return nil
}

func (m *TestCustomRuleInput) validateScanID(formats strfmt.Registry) error {

	if err := validate.Required("scan_id", "body", m.ScanID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestCustomRuleInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestCustomRuleInput) UnmarshalBinary(b []byte) error {
	var res TestCustomRuleInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
