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

// CreateanewSMSRequest CreateanewSMSRequest
// swagger:model CreateanewSMSRequest
type CreateanewSMSRequest struct {

	// callee
	// Required: true
	Callee *string `json:"callee"`

	// caller
	// Required: true
	Caller *string `json:"caller"`

	// direction
	// Required: true
	Direction *string `json:"direction"`

	// reason
	// Required: true
	Reason *string `json:"reason"`

	// status
	// Required: true
	Status *string `json:"status"`

	// subscriber id
	// Required: true
	SubscriberID *string `json:"subscriber_id"`

	// text
	// Required: true
	Text *string `json:"text"`
}

// Validate validates this createanew s m s request
func (m *CreateanewSMSRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCaller(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateanewSMSRequest) validateCallee(formats strfmt.Registry) error {

	if err := validate.Required("callee", "body", m.Callee); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewSMSRequest) validateCaller(formats strfmt.Registry) error {

	if err := validate.Required("caller", "body", m.Caller); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewSMSRequest) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewSMSRequest) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("reason", "body", m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewSMSRequest) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewSMSRequest) validateSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_id", "body", m.SubscriberID); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewSMSRequest) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateanewSMSRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateanewSMSRequest) UnmarshalBinary(b []byte) error {
	var res CreateanewSMSRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
