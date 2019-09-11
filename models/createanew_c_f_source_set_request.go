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

// CreateanewCFSourceSetRequest CreateanewCFSourceSetRequest
// swagger:model CreateanewCFSourceSetRequest
type CreateanewCFSourceSetRequest struct {

	// is regex
	// Required: true
	IsRegex *string `json:"is_regex"`

	// mode
	// Required: true
	Mode *string `json:"mode"`

	// name
	// Required: true
	Name *string `json:"name"`

	// sources
	// Required: true
	Sources []*Source `json:"sources"`

	// subscriber id
	// Required: true
	SubscriberID *string `json:"subscriber_id"`
}

// Validate validates this createanew c f source set request
func (m *CreateanewCFSourceSetRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsRegex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSources(formats); err != nil {
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

func (m *CreateanewCFSourceSetRequest) validateIsRegex(formats strfmt.Registry) error {

	if err := validate.Required("is_regex", "body", m.IsRegex); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewCFSourceSetRequest) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewCFSourceSetRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewCFSourceSetRequest) validateSources(formats strfmt.Registry) error {

	if err := validate.Required("sources", "body", m.Sources); err != nil {
		return err
	}

	for i := 0; i < len(m.Sources); i++ {
		if swag.IsZero(m.Sources[i]) { // not required
			continue
		}

		if m.Sources[i] != nil {
			if err := m.Sources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateanewCFSourceSetRequest) validateSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_id", "body", m.SubscriberID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateanewCFSourceSetRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateanewCFSourceSetRequest) UnmarshalBinary(b []byte) error {
	var res CreateanewCFSourceSetRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
