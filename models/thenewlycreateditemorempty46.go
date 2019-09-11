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

// Thenewlycreateditemorempty46 Thenewlycreateditemorempty46
// swagger:model Thenewlycreateditemorempty46
type Thenewlycreateditemorempty46 struct {

	// dir
	// Required: true
	Dir *string `json:"dir"`

	// greetingfile
	// Required: true
	Greetingfile *string `json:"greetingfile"`

	// subscriber id
	// Required: true
	SubscriberID *string `json:"subscriber_id"`
}

// Validate validates this thenewlycreateditemorempty46
func (m *Thenewlycreateditemorempty46) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDir(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGreetingfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Thenewlycreateditemorempty46) validateDir(formats strfmt.Registry) error {

	if err := validate.Required("dir", "body", m.Dir); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty46) validateGreetingfile(formats strfmt.Registry) error {

	if err := validate.Required("greetingfile", "body", m.Greetingfile); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty46) validateSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_id", "body", m.SubscriberID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Thenewlycreateditemorempty46) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Thenewlycreateditemorempty46) UnmarshalBinary(b []byte) error {
	var res Thenewlycreateditemorempty46
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
