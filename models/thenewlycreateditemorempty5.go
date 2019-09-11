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

// Thenewlycreateditemorempty5 Thenewlycreateditemorempty5
// swagger:model Thenewlycreateditemorempty5
type Thenewlycreateditemorempty5 struct {

	// activate timestamp
	// Required: true
	ActivateTimestamp *string `json:"activate_timestamp"`

	// add vat
	// Required: true
	AddVat *string `json:"add_vat"`

	// billing profile definition
	// Required: true
	BillingProfileDefinition *string `json:"billing_profile_definition"`

	// billing profile id
	// Required: true
	BillingProfileID *string `json:"billing_profile_id"`

	// billing profiles
	// Required: true
	BillingProfiles []*BillingProfile `json:"billing_profiles"`

	// contact id
	// Required: true
	ContactID *string `json:"contact_id"`

	// create timestamp
	// Required: true
	CreateTimestamp *string `json:"create_timestamp"`

	// external id
	// Required: true
	ExternalID *string `json:"external_id"`

	// invoice email template id
	// Required: true
	InvoiceEmailTemplateID *string `json:"invoice_email_template_id"`

	// invoice template id
	// Required: true
	InvoiceTemplateID *string `json:"invoice_template_id"`

	// max subscribers
	// Required: true
	MaxSubscribers *string `json:"max_subscribers"`

	// modify timestamp
	// Required: true
	ModifyTimestamp *string `json:"modify_timestamp"`

	// passreset email template id
	// Required: true
	PassresetEmailTemplateID *string `json:"passreset_email_template_id"`

	// profile package id
	// Required: true
	ProfilePackageID *string `json:"profile_package_id"`

	// status
	// Required: true
	Status *string `json:"status"`

	// subscriber email template id
	// Required: true
	SubscriberEmailTemplateID *string `json:"subscriber_email_template_id"`

	// terminate timestamp
	// Required: true
	TerminateTimestamp *string `json:"terminate_timestamp"`

	// type
	// Required: true
	Type *string `json:"type"`

	// vat rate
	// Required: true
	VatRate *string `json:"vat_rate"`
}

// Validate validates this thenewlycreateditemorempty5
func (m *Thenewlycreateditemorempty5) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddVat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingProfileDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingProfileID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingProfiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceEmailTemplateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceTemplateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxSubscribers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifyTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassresetEmailTemplateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfilePackageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberEmailTemplateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVatRate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Thenewlycreateditemorempty5) validateActivateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("activate_timestamp", "body", m.ActivateTimestamp); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty5) validateAddVat(formats strfmt.Registry) error {

	if err := validate.Required("add_vat", "body", m.AddVat); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty5) validateBillingProfileDefinition(formats strfmt.Registry) error {

	if err := validate.Required("billing_profile_definition", "body", m.BillingProfileDefinition); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty5) validateBillingProfileID(formats strfmt.Registry) error {

	if err := validate.Required("billing_profile_id", "body", m.BillingProfileID); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty5) validateBillingProfiles(formats strfmt.Registry) error {

	if err := validate.Required("billing_profiles", "body", m.BillingProfiles); err != nil {
		return err
	}

	for i := 0; i < len(m.BillingProfiles); i++ {
		if swag.IsZero(m.BillingProfiles[i]) { // not required
			continue
		}

		if m.BillingProfiles[i] != nil {
			if err := m.BillingProfiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("billing_profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Thenewlycreateditemorempty5) validateContactID(formats strfmt.Registry) error {

	if err := validate.Required("contact_id", "body", m.ContactID); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty5) validateCreateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("create_timestamp", "body", m.CreateTimestamp); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty5) validateExternalID(formats strfmt.Registry) error {

	if err := validate.Required("external_id", "body", m.ExternalID); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty5) validateInvoiceEmailTemplateID(formats strfmt.Registry) error {

	if err := validate.Required("invoice_email_template_id", "body", m.InvoiceEmailTemplateID); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty5) validateInvoiceTemplateID(formats strfmt.Registry) error {

	if err := validate.Required("invoice_template_id", "body", m.InvoiceTemplateID); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty5) validateMaxSubscribers(formats strfmt.Registry) error {

	if err := validate.Required("max_subscribers", "body", m.MaxSubscribers); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty5) validateModifyTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("modify_timestamp", "body", m.ModifyTimestamp); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty5) validatePassresetEmailTemplateID(formats strfmt.Registry) error {

	if err := validate.Required("passreset_email_template_id", "body", m.PassresetEmailTemplateID); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty5) validateProfilePackageID(formats strfmt.Registry) error {

	if err := validate.Required("profile_package_id", "body", m.ProfilePackageID); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty5) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty5) validateSubscriberEmailTemplateID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_email_template_id", "body", m.SubscriberEmailTemplateID); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty5) validateTerminateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("terminate_timestamp", "body", m.TerminateTimestamp); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty5) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty5) validateVatRate(formats strfmt.Registry) error {

	if err := validate.Required("vat_rate", "body", m.VatRate); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Thenewlycreateditemorempty5) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Thenewlycreateditemorempty5) UnmarshalBinary(b []byte) error {
	var res Thenewlycreateditemorempty5
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}