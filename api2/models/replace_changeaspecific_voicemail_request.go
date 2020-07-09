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

// ReplaceChangeaspecificVoicemailRequest Replace/changeaspecificVoicemailRequest
//
// swagger:model Replace/changeaspecificVoicemailRequest
type ReplaceChangeaspecificVoicemailRequest struct {

	// caller
	// Required: true
	Caller *string `json:"caller"`

	// duration
	// Required: true
	Duration *string `json:"duration"`

	// folder
	// Required: true
	Folder *string `json:"folder"`

	// subscriber id
	// Required: true
	SubscriberID *string `json:"subscriber_id"`

	// time
	// Required: true
	Time *string `json:"time"`
}

// Validate validates this replace changeaspecific voicemail request
func (m *ReplaceChangeaspecificVoicemailRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCaller(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFolder(formats); err != nil {
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

func (m *ReplaceChangeaspecificVoicemailRequest) validateCaller(formats strfmt.Registry) error {

	if err := validate.Required("caller", "body", m.Caller); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificVoicemailRequest) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificVoicemailRequest) validateFolder(formats strfmt.Registry) error {

	if err := validate.Required("folder", "body", m.Folder); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificVoicemailRequest) validateSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_id", "body", m.SubscriberID); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificVoicemailRequest) validateTime(formats strfmt.Registry) error {

	if err := validate.Required("time", "body", m.Time); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReplaceChangeaspecificVoicemailRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplaceChangeaspecificVoicemailRequest) UnmarshalBinary(b []byte) error {
	var res ReplaceChangeaspecificVoicemailRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}