// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SpeedDials SpeedDials
// swagger:model SpeedDials
type SpeedDials struct {

	// speeddials
	// Required: true
	Speeddials []*Speeddials1 `json:"speeddials"`
}

// Validate validates this speed dials
func (m *SpeedDials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSpeeddials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SpeedDials) validateSpeeddials(formats strfmt.Registry) error {

	if err := validate.Required("speeddials", "body", m.Speeddials); err != nil {
		return err
	}

	for i := 0; i < len(m.Speeddials); i++ {
		if swag.IsZero(m.Speeddials[i]) { // not required
			continue
		}

		if m.Speeddials[i] != nil {
			if err := m.Speeddials[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("speeddials" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SpeedDials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpeedDials) UnmarshalBinary(b []byte) error {
	var res SpeedDials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
