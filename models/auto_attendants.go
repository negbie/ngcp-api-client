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

// AutoAttendants AutoAttendants
// swagger:model AutoAttendants
type AutoAttendants struct {

	// slots
	// Required: true
	Slots []*Slot `json:"slots"`
}

// Validate validates this auto attendants
func (m *AutoAttendants) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSlots(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoAttendants) validateSlots(formats strfmt.Registry) error {

	if err := validate.Required("slots", "body", m.Slots); err != nil {
		return err
	}

	for i := 0; i < len(m.Slots); i++ {
		if swag.IsZero(m.Slots[i]) { // not required
			continue
		}

		if m.Slots[i] != nil {
			if err := m.Slots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("slots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutoAttendants) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoAttendants) UnmarshalBinary(b []byte) error {
	var res AutoAttendants
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
