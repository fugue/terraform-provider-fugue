// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateEnvironmentInput Structure of the body for creating a new environment.
//
// swagger:model CreateEnvironmentInput
type CreateEnvironmentInput struct {

	// List of compliance families validated against the environment.
	ComplianceFamilies []string `json:"compliance_families"`

	// Name of the environment.
	Name string `json:"name,omitempty"`

	// Name of the cloud service provider for the environment.
	// Enum: [aws aws_govcloud azure google]
	Provider string `json:"provider,omitempty"`

	// A dictionary of options for the provider.
	ProviderOptions *ProviderOptions `json:"provider_options,omitempty"`

	// List of resource types to be remediated if remediation is enabled.
	RemediateResourceTypes []string `json:"remediate_resource_types"`

	// Time in seconds between the end of one scan to the start of the next. Must also set scan_schedule_enabled to true.
	// Minimum: 300
	ScanInterval *int64 `json:"scan_interval,omitempty"`

	// Indicates if the new environment should have scans run on a schedule upon creation.
	ScanScheduleEnabled *bool `json:"scan_schedule_enabled,omitempty"`

	// List of resource types to be surveyed.
	SurveyResourceTypes []string `json:"survey_resource_types"`
}

// Validate validates this create environment input
func (m *CreateEnvironmentInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanInterval(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createEnvironmentInputTypeProviderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aws","aws_govcloud","azure","google"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createEnvironmentInputTypeProviderPropEnum = append(createEnvironmentInputTypeProviderPropEnum, v)
	}
}

const (

	// CreateEnvironmentInputProviderAws captures enum value "aws"
	CreateEnvironmentInputProviderAws string = "aws"

	// CreateEnvironmentInputProviderAwsGovcloud captures enum value "aws_govcloud"
	CreateEnvironmentInputProviderAwsGovcloud string = "aws_govcloud"

	// CreateEnvironmentInputProviderAzure captures enum value "azure"
	CreateEnvironmentInputProviderAzure string = "azure"

	// CreateEnvironmentInputProviderGoogle captures enum value "google"
	CreateEnvironmentInputProviderGoogle string = "google"
)

// prop value enum
func (m *CreateEnvironmentInput) validateProviderEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, createEnvironmentInputTypeProviderPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreateEnvironmentInput) validateProvider(formats strfmt.Registry) error {

	if swag.IsZero(m.Provider) { // not required
		return nil
	}

	// value enum
	if err := m.validateProviderEnum("provider", "body", m.Provider); err != nil {
		return err
	}

	return nil
}

func (m *CreateEnvironmentInput) validateProviderOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.ProviderOptions) { // not required
		return nil
	}

	if m.ProviderOptions != nil {
		if err := m.ProviderOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provider_options")
			}
			return err
		}
	}

	return nil
}

func (m *CreateEnvironmentInput) validateScanInterval(formats strfmt.Registry) error {

	if swag.IsZero(m.ScanInterval) { // not required
		return nil
	}

	if err := validate.MinimumInt("scan_interval", "body", int64(*m.ScanInterval), 300, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateEnvironmentInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateEnvironmentInput) UnmarshalBinary(b []byte) error {
	var res CreateEnvironmentInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
