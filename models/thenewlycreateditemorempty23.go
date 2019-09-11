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

// Thenewlycreateditemorempty23 Thenewlycreateditemorempty23
// swagger:model Thenewlycreateditemorempty23
type Thenewlycreateditemorempty23 struct {

	// bankname
	// Required: true
	Bankname *string `json:"bankname"`

	// bic
	// Required: true
	Bic *string `json:"bic"`

	// city
	// Required: true
	City *string `json:"city"`

	// company
	// Required: true
	Company *string `json:"company"`

	// comregnum
	// Required: true
	Comregnum *string `json:"comregnum"`

	// country
	// Required: true
	Country *string `json:"country"`

	// email
	// Required: true
	Email *string `json:"email"`

	// faxnumber
	// Required: true
	Faxnumber *string `json:"faxnumber"`

	// firstname
	// Required: true
	Firstname *string `json:"firstname"`

	// gpp0
	// Required: true
	Gpp0 *string `json:"gpp0"`

	// gpp1
	// Required: true
	Gpp1 *string `json:"gpp1"`

	// gpp2
	// Required: true
	Gpp2 *string `json:"gpp2"`

	// gpp3
	// Required: true
	Gpp3 *string `json:"gpp3"`

	// gpp4
	// Required: true
	Gpp4 *string `json:"gpp4"`

	// gpp5
	// Required: true
	Gpp5 *string `json:"gpp5"`

	// gpp6
	// Required: true
	Gpp6 *string `json:"gpp6"`

	// gpp7
	// Required: true
	Gpp7 *string `json:"gpp7"`

	// gpp8
	// Required: true
	Gpp8 *string `json:"gpp8"`

	// gpp9
	// Required: true
	Gpp9 *string `json:"gpp9"`

	// iban
	// Required: true
	Iban *string `json:"iban"`

	// lastname
	// Required: true
	Lastname *string `json:"lastname"`

	// mobilenumber
	// Required: true
	Mobilenumber *string `json:"mobilenumber"`

	// phonenumber
	// Required: true
	Phonenumber *string `json:"phonenumber"`

	// postcode
	// Required: true
	Postcode *string `json:"postcode"`

	// reseller id
	// Required: true
	ResellerID *string `json:"reseller_id"`

	// street
	// Required: true
	Street *string `json:"street"`

	// timezone
	// Required: true
	Timezone *string `json:"timezone"`

	// vatnum
	// Required: true
	Vatnum *string `json:"vatnum"`
}

// Validate validates this thenewlycreateditemorempty23
func (m *Thenewlycreateditemorempty23) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBankname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompany(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComregnum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFaxnumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpp0(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpp1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpp2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpp3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpp4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpp5(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpp6(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpp7(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpp8(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpp9(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIban(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMobilenumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhonenumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostcode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResellerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStreet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimezone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVatnum(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Thenewlycreateditemorempty23) validateBankname(formats strfmt.Registry) error {

	if err := validate.Required("bankname", "body", m.Bankname); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateBic(formats strfmt.Registry) error {

	if err := validate.Required("bic", "body", m.Bic); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateCity(formats strfmt.Registry) error {

	if err := validate.Required("city", "body", m.City); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateCompany(formats strfmt.Registry) error {

	if err := validate.Required("company", "body", m.Company); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateComregnum(formats strfmt.Registry) error {

	if err := validate.Required("comregnum", "body", m.Comregnum); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateFaxnumber(formats strfmt.Registry) error {

	if err := validate.Required("faxnumber", "body", m.Faxnumber); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateFirstname(formats strfmt.Registry) error {

	if err := validate.Required("firstname", "body", m.Firstname); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateGpp0(formats strfmt.Registry) error {

	if err := validate.Required("gpp0", "body", m.Gpp0); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateGpp1(formats strfmt.Registry) error {

	if err := validate.Required("gpp1", "body", m.Gpp1); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateGpp2(formats strfmt.Registry) error {

	if err := validate.Required("gpp2", "body", m.Gpp2); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateGpp3(formats strfmt.Registry) error {

	if err := validate.Required("gpp3", "body", m.Gpp3); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateGpp4(formats strfmt.Registry) error {

	if err := validate.Required("gpp4", "body", m.Gpp4); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateGpp5(formats strfmt.Registry) error {

	if err := validate.Required("gpp5", "body", m.Gpp5); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateGpp6(formats strfmt.Registry) error {

	if err := validate.Required("gpp6", "body", m.Gpp6); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateGpp7(formats strfmt.Registry) error {

	if err := validate.Required("gpp7", "body", m.Gpp7); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateGpp8(formats strfmt.Registry) error {

	if err := validate.Required("gpp8", "body", m.Gpp8); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateGpp9(formats strfmt.Registry) error {

	if err := validate.Required("gpp9", "body", m.Gpp9); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateIban(formats strfmt.Registry) error {

	if err := validate.Required("iban", "body", m.Iban); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateLastname(formats strfmt.Registry) error {

	if err := validate.Required("lastname", "body", m.Lastname); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateMobilenumber(formats strfmt.Registry) error {

	if err := validate.Required("mobilenumber", "body", m.Mobilenumber); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validatePhonenumber(formats strfmt.Registry) error {

	if err := validate.Required("phonenumber", "body", m.Phonenumber); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validatePostcode(formats strfmt.Registry) error {

	if err := validate.Required("postcode", "body", m.Postcode); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateResellerID(formats strfmt.Registry) error {

	if err := validate.Required("reseller_id", "body", m.ResellerID); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateStreet(formats strfmt.Registry) error {

	if err := validate.Required("street", "body", m.Street); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateTimezone(formats strfmt.Registry) error {

	if err := validate.Required("timezone", "body", m.Timezone); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty23) validateVatnum(formats strfmt.Registry) error {

	if err := validate.Required("vatnum", "body", m.Vatnum); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Thenewlycreateditemorempty23) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Thenewlycreateditemorempty23) UnmarshalBinary(b []byte) error {
	var res Thenewlycreateditemorempty23
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
