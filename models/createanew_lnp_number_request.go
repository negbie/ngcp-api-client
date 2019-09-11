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

// CreateanewLnpNumberRequest CreateanewLnpNumberRequest
// swagger:model CreateanewLnpNumberRequest
type CreateanewLnpNumberRequest struct {

	// carrier id
	// Required: true
	CarrierID *string `json:"carrier_id"`

	// end
	// Required: true
	End *string `json:"end"`

	// number
	// Required: true
	Number *string `json:"number"`

	// routing number
	// Required: true
	RoutingNumber *string `json:"routing_number"`

	// start
	// Required: true
	Start *string `json:"start"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this createanew lnp number request
func (m *CreateanewLnpNumberRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCarrierID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutingNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateanewLnpNumberRequest) validateCarrierID(formats strfmt.Registry) error {

	if err := validate.Required("carrier_id", "body", m.CarrierID); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewLnpNumberRequest) validateEnd(formats strfmt.Registry) error {

	if err := validate.Required("end", "body", m.End); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewLnpNumberRequest) validateNumber(formats strfmt.Registry) error {

	if err := validate.Required("number", "body", m.Number); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewLnpNumberRequest) validateRoutingNumber(formats strfmt.Registry) error {

	if err := validate.Required("routing_number", "body", m.RoutingNumber); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewLnpNumberRequest) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("start", "body", m.Start); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewLnpNumberRequest) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateanewLnpNumberRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateanewLnpNumberRequest) UnmarshalBinary(b []byte) error {
	var res CreateanewLnpNumberRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
