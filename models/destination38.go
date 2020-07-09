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

// Destination38 Destination38
//
// swagger:model Destination38
type Destination38 struct {

	// destination
	// Required: true
	Destination *string `json:"destination"`

	// filetype
	// Required: true
	Filetype *string `json:"filetype"`

	// incoming
	// Required: true
	Incoming *bool `json:"incoming"`

	// outgoing
	// Required: true
	Outgoing *bool `json:"outgoing"`

	// status
	// Required: true
	Status *bool `json:"status"`
}

// Validate validates this destination38
func (m *Destination38) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFiletype(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncoming(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutgoing(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Destination38) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("destination", "body", m.Destination); err != nil {
		return err
	}

	return nil
}

func (m *Destination38) validateFiletype(formats strfmt.Registry) error {

	if err := validate.Required("filetype", "body", m.Filetype); err != nil {
		return err
	}

	return nil
}

func (m *Destination38) validateIncoming(formats strfmt.Registry) error {

	if err := validate.Required("incoming", "body", m.Incoming); err != nil {
		return err
	}

	return nil
}

func (m *Destination38) validateOutgoing(formats strfmt.Registry) error {

	if err := validate.Required("outgoing", "body", m.Outgoing); err != nil {
		return err
	}

	return nil
}

func (m *Destination38) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Destination38) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Destination38) UnmarshalBinary(b []byte) error {
	var res Destination38
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}