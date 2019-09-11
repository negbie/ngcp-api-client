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

// NcosLnpCarriers NcosLnpCarriers
// swagger:model NcosLnpCarriers
type NcosLnpCarriers struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// ncos level id
	// Required: true
	NcosLevelID *string `json:"ncos_level_id"`
}

// Validate validates this ncos lnp carriers
func (m *NcosLnpCarriers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNcosLevelID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NcosLnpCarriers) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *NcosLnpCarriers) validateNcosLevelID(formats strfmt.Registry) error {

	if err := validate.Required("ncos_level_id", "body", m.NcosLevelID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NcosLnpCarriers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NcosLnpCarriers) UnmarshalBinary(b []byte) error {
	var res NcosLnpCarriers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
