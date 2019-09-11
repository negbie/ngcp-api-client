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

// PbxDeviceProfiles PbxDeviceProfiles
// swagger:model PbxDeviceProfiles
type PbxDeviceProfiles struct {

	// config id
	// Required: true
	ConfigID *string `json:"config_id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this pbx device profiles
func (m *PbxDeviceProfiles) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PbxDeviceProfiles) validateConfigID(formats strfmt.Registry) error {

	if err := validate.Required("config_id", "body", m.ConfigID); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceProfiles) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PbxDeviceProfiles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PbxDeviceProfiles) UnmarshalBinary(b []byte) error {
	var res PbxDeviceProfiles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
