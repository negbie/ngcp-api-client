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

// CreateanewAdminRequest CreateanewAdminRequest
// swagger:model CreateanewAdminRequest
type CreateanewAdminRequest struct {

	// billing data
	// Required: true
	BillingData *string `json:"billing_data"`

	// call data
	// Required: true
	CallData *string `json:"call_data"`

	// is active
	// Required: true
	IsActive *string `json:"is_active"`

	// is master
	// Required: true
	IsMaster *string `json:"is_master"`

	// is superuser
	// Required: true
	IsSuperuser *string `json:"is_superuser"`

	// lawful intercept
	// Required: true
	LawfulIntercept *string `json:"lawful_intercept"`

	// login
	// Required: true
	Login *string `json:"login"`

	// password
	// Required: true
	Password *string `json:"password"`

	// read only
	// Required: true
	ReadOnly *string `json:"read_only"`

	// reseller id
	// Required: true
	ResellerID *string `json:"reseller_id"`

	// show passwords
	// Required: true
	ShowPasswords *string `json:"show_passwords"`
}

// Validate validates this createanew admin request
func (m *CreateanewAdminRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsMaster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsSuperuser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLawfulIntercept(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadOnly(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResellerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShowPasswords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateanewAdminRequest) validateBillingData(formats strfmt.Registry) error {

	if err := validate.Required("billing_data", "body", m.BillingData); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewAdminRequest) validateCallData(formats strfmt.Registry) error {

	if err := validate.Required("call_data", "body", m.CallData); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewAdminRequest) validateIsActive(formats strfmt.Registry) error {

	if err := validate.Required("is_active", "body", m.IsActive); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewAdminRequest) validateIsMaster(formats strfmt.Registry) error {

	if err := validate.Required("is_master", "body", m.IsMaster); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewAdminRequest) validateIsSuperuser(formats strfmt.Registry) error {

	if err := validate.Required("is_superuser", "body", m.IsSuperuser); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewAdminRequest) validateLawfulIntercept(formats strfmt.Registry) error {

	if err := validate.Required("lawful_intercept", "body", m.LawfulIntercept); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewAdminRequest) validateLogin(formats strfmt.Registry) error {

	if err := validate.Required("login", "body", m.Login); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewAdminRequest) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewAdminRequest) validateReadOnly(formats strfmt.Registry) error {

	if err := validate.Required("read_only", "body", m.ReadOnly); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewAdminRequest) validateResellerID(formats strfmt.Registry) error {

	if err := validate.Required("reseller_id", "body", m.ResellerID); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewAdminRequest) validateShowPasswords(formats strfmt.Registry) error {

	if err := validate.Required("show_passwords", "body", m.ShowPasswords); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateanewAdminRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateanewAdminRequest) UnmarshalBinary(b []byte) error {
	var res CreateanewAdminRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
