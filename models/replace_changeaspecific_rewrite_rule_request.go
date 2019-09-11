// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReplaceChangeaspecificRewriteRuleRequest Replace/changeaspecificRewriteRuleRequest
// swagger:model Replace/changeaspecificRewriteRuleRequest
type ReplaceChangeaspecificRewriteRuleRequest struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// direction
	// Required: true
	Direction *string `json:"direction"`

	// enabled
	// Required: true
	Enabled *string `json:"enabled"`

	// field
	// Required: true
	Field *string `json:"field"`

	// match pattern
	// Required: true
	MatchPattern *string `json:"match_pattern"`

	// priority
	// Required: true
	Priority *string `json:"priority"`

	// replace pattern
	// Required: true
	ReplacePattern *string `json:"replace_pattern"`

	// set id
	// Required: true
	SetID *string `json:"set_id"`
}

// Validate validates this replace changeaspecific rewrite rule request
func (m *ReplaceChangeaspecificRewriteRuleRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateField(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatchPattern(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplacePattern(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSetID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplaceChangeaspecificRewriteRuleRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificRewriteRuleRequest) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificRewriteRuleRequest) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificRewriteRuleRequest) validateField(formats strfmt.Registry) error {

	if err := validate.Required("field", "body", m.Field); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificRewriteRuleRequest) validateMatchPattern(formats strfmt.Registry) error {

	if err := validate.Required("match_pattern", "body", m.MatchPattern); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificRewriteRuleRequest) validatePriority(formats strfmt.Registry) error {

	if err := validate.Required("priority", "body", m.Priority); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificRewriteRuleRequest) validateReplacePattern(formats strfmt.Registry) error {

	if err := validate.Required("replace_pattern", "body", m.ReplacePattern); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificRewriteRuleRequest) validateSetID(formats strfmt.Registry) error {

	if err := validate.Required("set_id", "body", m.SetID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReplaceChangeaspecificRewriteRuleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplaceChangeaspecificRewriteRuleRequest) UnmarshalBinary(b []byte) error {
	var res ReplaceChangeaspecificRewriteRuleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}