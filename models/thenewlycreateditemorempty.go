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

// Thenewlycreateditemorempty Thenewlycreateditemorempty
// swagger:model Thenewlycreateditemorempty
type Thenewlycreateditemorempty struct {

	// authoritative
	// Required: true
	Authoritative *string `json:"authoritative"`

	// name
	// Required: true
	Name *string `json:"name"`

	// prefix
	// Required: true
	Prefix *string `json:"prefix"`

	// skip rewrite
	// Required: true
	SkipRewrite *string `json:"skip_rewrite"`
}

// Validate validates this thenewlycreateditemorempty
func (m *Thenewlycreateditemorempty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthoritative(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkipRewrite(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Thenewlycreateditemorempty) validateAuthoritative(formats strfmt.Registry) error {

	if err := validate.Required("authoritative", "body", m.Authoritative); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty) validatePrefix(formats strfmt.Registry) error {

	if err := validate.Required("prefix", "body", m.Prefix); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty) validateSkipRewrite(formats strfmt.Registry) error {

	if err := validate.Required("skip_rewrite", "body", m.SkipRewrite); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Thenewlycreateditemorempty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Thenewlycreateditemorempty) UnmarshalBinary(b []byte) error {
	var res Thenewlycreateditemorempty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
