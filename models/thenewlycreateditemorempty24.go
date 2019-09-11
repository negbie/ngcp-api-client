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

// Thenewlycreateditemorempty24 Thenewlycreateditemorempty24
// swagger:model Thenewlycreateditemorempty24
type Thenewlycreateditemorempty24 struct {

	// active
	// Required: true
	Active *string `json:"active"`

	// recur
	// Required: true
	Recur *string `json:"recur"`

	// subscriber id
	// Required: true
	SubscriberID *string `json:"subscriber_id"`

	// time
	// Required: true
	Time *string `json:"time"`
}

// Validate validates this thenewlycreateditemorempty24
func (m *Thenewlycreateditemorempty24) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecur(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Thenewlycreateditemorempty24) validateActive(formats strfmt.Registry) error {

	if err := validate.Required("active", "body", m.Active); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty24) validateRecur(formats strfmt.Registry) error {

	if err := validate.Required("recur", "body", m.Recur); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty24) validateSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_id", "body", m.SubscriberID); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty24) validateTime(formats strfmt.Registry) error {

	if err := validate.Required("time", "body", m.Time); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Thenewlycreateditemorempty24) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Thenewlycreateditemorempty24) UnmarshalBinary(b []byte) error {
	var res Thenewlycreateditemorempty24
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
