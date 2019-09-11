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

// Faxes Faxes
// swagger:model Faxes
type Faxes struct {

	// data
	// Required: true
	Data *string `json:"data"`

	// destination
	// Required: true
	Destination *string `json:"destination"`

	// faxfile
	// Required: true
	Faxfile *string `json:"faxfile"`

	// pageheader
	// Required: true
	Pageheader *string `json:"pageheader"`

	// quality
	// Required: true
	Quality *string `json:"quality"`

	// subscriber id
	// Required: true
	SubscriberID *string `json:"subscriber_id"`
}

// Validate validates this faxes
func (m *Faxes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFaxfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageheader(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuality(formats); err != nil {
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

func (m *Faxes) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	return nil
}

func (m *Faxes) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("destination", "body", m.Destination); err != nil {
		return err
	}

	return nil
}

func (m *Faxes) validateFaxfile(formats strfmt.Registry) error {

	if err := validate.Required("faxfile", "body", m.Faxfile); err != nil {
		return err
	}

	return nil
}

func (m *Faxes) validatePageheader(formats strfmt.Registry) error {

	if err := validate.Required("pageheader", "body", m.Pageheader); err != nil {
		return err
	}

	return nil
}

func (m *Faxes) validateQuality(formats strfmt.Registry) error {

	if err := validate.Required("quality", "body", m.Quality); err != nil {
		return err
	}

	return nil
}

func (m *Faxes) validateSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_id", "body", m.SubscriberID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Faxes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Faxes) UnmarshalBinary(b []byte) error {
	var res Faxes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
