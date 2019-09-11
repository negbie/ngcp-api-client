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

// Slot Slot
// swagger:model Slot
type Slot struct {

	// destination
	// Required: true
	Destination *string `json:"destination"`

	// slot
	// Required: true
	Slot *string `json:"slot"`
}

// Validate validates this slot
func (m *Slot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Slot) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("destination", "body", m.Destination); err != nil {
		return err
	}

	return nil
}

func (m *Slot) validateSlot(formats strfmt.Registry) error {

	if err := validate.Required("slot", "body", m.Slot); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Slot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Slot) UnmarshalBinary(b []byte) error {
	var res Slot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
