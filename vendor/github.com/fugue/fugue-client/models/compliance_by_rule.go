// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ComplianceByRule Compliance rule and result.
//
// swagger:model ComplianceByRule
type ComplianceByRule struct {

	// List of resource types that failed to satisfy the rule due to a required resource being omitted and associated error messages.
	FailedResourceTypes []*ComplianceByRuleFailedResourceTypesItems0 `json:"failed_resource_types"`

	// List of resources that failed to satisfy the rule due to a misconfiguration in the resource and associated error messages.
	FailedResources []*ComplianceByRuleFailedResourcesItems0 `json:"failed_resources"`

	// Name of the compliance family.
	Family string `json:"family,omitempty"`

	// Result of the rule.
	// Enum: [PASS FAIL UNKNOWN]
	Result string `json:"result,omitempty"`

	// Name of the compliance rule.
	Rule string `json:"rule,omitempty"`

	// List of resource types that were not surveyed and caused the result to be unknown.
	UnsurveyedResourceTypes []string `json:"unsurveyed_resource_types"`
}

// Validate validates this compliance by rule
func (m *ComplianceByRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFailedResourceTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailedResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComplianceByRule) validateFailedResourceTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.FailedResourceTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.FailedResourceTypes); i++ {
		if swag.IsZero(m.FailedResourceTypes[i]) { // not required
			continue
		}

		if m.FailedResourceTypes[i] != nil {
			if err := m.FailedResourceTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("failed_resource_types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComplianceByRule) validateFailedResources(formats strfmt.Registry) error {

	if swag.IsZero(m.FailedResources) { // not required
		return nil
	}

	for i := 0; i < len(m.FailedResources); i++ {
		if swag.IsZero(m.FailedResources[i]) { // not required
			continue
		}

		if m.FailedResources[i] != nil {
			if err := m.FailedResources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("failed_resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var complianceByRuleTypeResultPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PASS","FAIL","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		complianceByRuleTypeResultPropEnum = append(complianceByRuleTypeResultPropEnum, v)
	}
}

const (

	// ComplianceByRuleResultPASS captures enum value "PASS"
	ComplianceByRuleResultPASS string = "PASS"

	// ComplianceByRuleResultFAIL captures enum value "FAIL"
	ComplianceByRuleResultFAIL string = "FAIL"

	// ComplianceByRuleResultUNKNOWN captures enum value "UNKNOWN"
	ComplianceByRuleResultUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *ComplianceByRule) validateResultEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, complianceByRuleTypeResultPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ComplianceByRule) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	// value enum
	if err := m.validateResultEnum("result", "body", m.Result); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComplianceByRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComplianceByRule) UnmarshalBinary(b []byte) error {
	var res ComplianceByRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ComplianceByRuleFailedResourceTypesItems0 Resource type that failed to satisfy the rule due to a required resource being omitted and associated error messages.
//
// swagger:model ComplianceByRuleFailedResourceTypesItems0
type ComplianceByRuleFailedResourceTypesItems0 struct {

	// Messages why the rule failed.
	Messages []string `json:"messages"`

	// Resource type that failed to satisfy the rule.
	ResourceType string `json:"resource_type,omitempty"`
}

// Validate validates this compliance by rule failed resource types items0
func (m *ComplianceByRuleFailedResourceTypesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComplianceByRuleFailedResourceTypesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComplianceByRuleFailedResourceTypesItems0) UnmarshalBinary(b []byte) error {
	var res ComplianceByRuleFailedResourceTypesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ComplianceByRuleFailedResourcesItems0 Resource that failed to satisfy the rule due to a misconfiguration in the resource and associated error messages.
//
// swagger:model ComplianceByRuleFailedResourcesItems0
type ComplianceByRuleFailedResourcesItems0 struct {

	// Messages why the rule failed.
	Messages []string `json:"messages"`

	// resource
	Resource *Resource `json:"resource,omitempty"`
}

// Validate validates this compliance by rule failed resources items0
func (m *ComplianceByRuleFailedResourcesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComplianceByRuleFailedResourcesItems0) validateResource(formats strfmt.Registry) error {

	if swag.IsZero(m.Resource) { // not required
		return nil
	}

	if m.Resource != nil {
		if err := m.Resource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComplianceByRuleFailedResourcesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComplianceByRuleFailedResourcesItems0) UnmarshalBinary(b []byte) error {
	var res ComplianceByRuleFailedResourcesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
