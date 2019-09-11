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

// CreateanewPeeringServerRequest CreateanewPeeringServerRequest
// swagger:model CreateanewPeeringServerRequest
type CreateanewPeeringServerRequest struct {

	// enabled
	// Required: true
	Enabled *string `json:"enabled"`

	// group id
	// Required: true
	GroupID *string `json:"group_id"`

	// host
	// Required: true
	Host *string `json:"host"`

	// ip
	// Required: true
	IP *string `json:"ip"`

	// name
	// Required: true
	Name *string `json:"name"`

	// port
	// Required: true
	Port *string `json:"port"`

	// probe
	// Required: true
	Probe *string `json:"probe"`

	// transport
	// Required: true
	Transport *string `json:"transport"`

	// via route
	// Required: true
	ViaRoute *string `json:"via_route"`

	// weight
	// Required: true
	Weight *string `json:"weight"`
}

// Validate validates this createanew peering server request
func (m *CreateanewPeeringServerRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProbe(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViaRoute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateanewPeeringServerRequest) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewPeeringServerRequest) validateGroupID(formats strfmt.Registry) error {

	if err := validate.Required("group_id", "body", m.GroupID); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewPeeringServerRequest) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewPeeringServerRequest) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewPeeringServerRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewPeeringServerRequest) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewPeeringServerRequest) validateProbe(formats strfmt.Registry) error {

	if err := validate.Required("probe", "body", m.Probe); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewPeeringServerRequest) validateTransport(formats strfmt.Registry) error {

	if err := validate.Required("transport", "body", m.Transport); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewPeeringServerRequest) validateViaRoute(formats strfmt.Registry) error {

	if err := validate.Required("via_route", "body", m.ViaRoute); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewPeeringServerRequest) validateWeight(formats strfmt.Registry) error {

	if err := validate.Required("weight", "body", m.Weight); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateanewPeeringServerRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateanewPeeringServerRequest) UnmarshalBinary(b []byte) error {
	var res CreateanewPeeringServerRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
