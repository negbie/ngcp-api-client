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

// Cft4 Cft4
// swagger:model Cft4
type Cft4 struct {

	// bnumberset
	// Required: true
	Bnumberset *string `json:"bnumberset"`

	// bnumberset id
	// Required: true
	BnumbersetID *string `json:"bnumberset_id"`

	// destinationset
	// Required: true
	Destinationset *string `json:"destinationset"`

	// destinationset id
	// Required: true
	DestinationsetID *string `json:"destinationset_id"`

	// sourceset
	// Required: true
	Sourceset *string `json:"sourceset"`

	// sourceset id
	// Required: true
	SourcesetID *string `json:"sourceset_id"`

	// timeset
	// Required: true
	Timeset *string `json:"timeset"`

	// timeset id
	// Required: true
	TimesetID *string `json:"timeset_id"`
}

// Validate validates this cft4
func (m *Cft4) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBnumberset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBnumbersetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationsetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourcesetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimesetID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cft4) validateBnumberset(formats strfmt.Registry) error {

	if err := validate.Required("bnumberset", "body", m.Bnumberset); err != nil {
		return err
	}

	return nil
}

func (m *Cft4) validateBnumbersetID(formats strfmt.Registry) error {

	if err := validate.Required("bnumberset_id", "body", m.BnumbersetID); err != nil {
		return err
	}

	return nil
}

func (m *Cft4) validateDestinationset(formats strfmt.Registry) error {

	if err := validate.Required("destinationset", "body", m.Destinationset); err != nil {
		return err
	}

	return nil
}

func (m *Cft4) validateDestinationsetID(formats strfmt.Registry) error {

	if err := validate.Required("destinationset_id", "body", m.DestinationsetID); err != nil {
		return err
	}

	return nil
}

func (m *Cft4) validateSourceset(formats strfmt.Registry) error {

	if err := validate.Required("sourceset", "body", m.Sourceset); err != nil {
		return err
	}

	return nil
}

func (m *Cft4) validateSourcesetID(formats strfmt.Registry) error {

	if err := validate.Required("sourceset_id", "body", m.SourcesetID); err != nil {
		return err
	}

	return nil
}

func (m *Cft4) validateTimeset(formats strfmt.Registry) error {

	if err := validate.Required("timeset", "body", m.Timeset); err != nil {
		return err
	}

	return nil
}

func (m *Cft4) validateTimesetID(formats strfmt.Registry) error {

	if err := validate.Required("timeset_id", "body", m.TimesetID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cft4) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cft4) UnmarshalBinary(b []byte) error {
	var res Cft4
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
