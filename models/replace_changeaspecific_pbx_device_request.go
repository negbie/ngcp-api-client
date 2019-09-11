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

// ReplaceChangeaspecificPbxDeviceRequest Replace/changeaspecificPbxDeviceRequest
// swagger:model Replace/changeaspecificPbxDeviceRequest
type ReplaceChangeaspecificPbxDeviceRequest struct {

	// customer id
	// Required: true
	CustomerID *string `json:"customer_id"`

	// identifier
	// Required: true
	Identifier *string `json:"identifier"`

	// lines
	// Required: true
	Lines []*Line `json:"lines"`

	// profile id
	// Required: true
	ProfileID *string `json:"profile_id"`

	// station name
	// Required: true
	StationName *string `json:"station_name"`
}

// Validate validates this replace changeaspecific pbx device request
func (m *ReplaceChangeaspecificPbxDeviceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLines(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfileID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplaceChangeaspecificPbxDeviceRequest) validateCustomerID(formats strfmt.Registry) error {

	if err := validate.Required("customer_id", "body", m.CustomerID); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificPbxDeviceRequest) validateIdentifier(formats strfmt.Registry) error {

	if err := validate.Required("identifier", "body", m.Identifier); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificPbxDeviceRequest) validateLines(formats strfmt.Registry) error {

	if err := validate.Required("lines", "body", m.Lines); err != nil {
		return err
	}

	for i := 0; i < len(m.Lines); i++ {
		if swag.IsZero(m.Lines[i]) { // not required
			continue
		}

		if m.Lines[i] != nil {
			if err := m.Lines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReplaceChangeaspecificPbxDeviceRequest) validateProfileID(formats strfmt.Registry) error {

	if err := validate.Required("profile_id", "body", m.ProfileID); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificPbxDeviceRequest) validateStationName(formats strfmt.Registry) error {

	if err := validate.Required("station_name", "body", m.StationName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReplaceChangeaspecificPbxDeviceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplaceChangeaspecificPbxDeviceRequest) UnmarshalBinary(b []byte) error {
	var res ReplaceChangeaspecificPbxDeviceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
