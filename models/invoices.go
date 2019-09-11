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

// Invoices Invoices
// swagger:model Invoices
type Invoices struct {

	// amount net
	// Required: true
	AmountNet *string `json:"amount_net"`

	// amount total
	// Required: true
	AmountTotal *string `json:"amount_total"`

	// amount vat
	// Required: true
	AmountVat *string `json:"amount_vat"`

	// customer id
	// Required: true
	CustomerID *string `json:"customer_id"`

	// period
	// Required: true
	Period *string `json:"period"`

	// period end
	// Required: true
	PeriodEnd *string `json:"period_end"`

	// period start
	// Required: true
	PeriodStart *string `json:"period_start"`

	// sent date
	// Required: true
	SentDate *string `json:"sent_date"`

	// serial
	// Required: true
	Serial *string `json:"serial"`

	// template id
	// Required: true
	TemplateID *string `json:"template_id"`
}

// Validate validates this invoices
func (m *Invoices) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmountNet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmountTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmountVat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeriodEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeriodStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSentDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Invoices) validateAmountNet(formats strfmt.Registry) error {

	if err := validate.Required("amount_net", "body", m.AmountNet); err != nil {
		return err
	}

	return nil
}

func (m *Invoices) validateAmountTotal(formats strfmt.Registry) error {

	if err := validate.Required("amount_total", "body", m.AmountTotal); err != nil {
		return err
	}

	return nil
}

func (m *Invoices) validateAmountVat(formats strfmt.Registry) error {

	if err := validate.Required("amount_vat", "body", m.AmountVat); err != nil {
		return err
	}

	return nil
}

func (m *Invoices) validateCustomerID(formats strfmt.Registry) error {

	if err := validate.Required("customer_id", "body", m.CustomerID); err != nil {
		return err
	}

	return nil
}

func (m *Invoices) validatePeriod(formats strfmt.Registry) error {

	if err := validate.Required("period", "body", m.Period); err != nil {
		return err
	}

	return nil
}

func (m *Invoices) validatePeriodEnd(formats strfmt.Registry) error {

	if err := validate.Required("period_end", "body", m.PeriodEnd); err != nil {
		return err
	}

	return nil
}

func (m *Invoices) validatePeriodStart(formats strfmt.Registry) error {

	if err := validate.Required("period_start", "body", m.PeriodStart); err != nil {
		return err
	}

	return nil
}

func (m *Invoices) validateSentDate(formats strfmt.Registry) error {

	if err := validate.Required("sent_date", "body", m.SentDate); err != nil {
		return err
	}

	return nil
}

func (m *Invoices) validateSerial(formats strfmt.Registry) error {

	if err := validate.Required("serial", "body", m.Serial); err != nil {
		return err
	}

	return nil
}

func (m *Invoices) validateTemplateID(formats strfmt.Registry) error {

	if err := validate.Required("template_id", "body", m.TemplateID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Invoices) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Invoices) UnmarshalBinary(b []byte) error {
	var res Invoices
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
