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

// Thenewlycreateditemorempty33 Thenewlycreateditemorempty33
// swagger:model Thenewlycreateditemorempty33
type Thenewlycreateditemorempty33 struct {

	// balance interval start mode
	// Required: true
	BalanceIntervalStartMode *string `json:"balance_interval_start_mode"`

	// balance interval unit
	// Required: true
	BalanceIntervalUnit *string `json:"balance_interval_unit"`

	// balance interval value
	// Required: true
	BalanceIntervalValue *string `json:"balance_interval_value"`

	// carry over mode
	// Required: true
	CarryOverMode *string `json:"carry_over_mode"`

	// description
	// Required: true
	Description *string `json:"description"`

	// initial balance
	// Required: true
	InitialBalance *string `json:"initial_balance"`

	// initial profiles
	// Required: true
	InitialProfiles []*InitialProfile `json:"initial_profiles"`

	// name
	// Required: true
	Name *string `json:"name"`

	// notopup discard intervals
	// Required: true
	NotopupDiscardIntervals *string `json:"notopup_discard_intervals"`

	// reseller id
	// Required: true
	ResellerID *string `json:"reseller_id"`

	// service charge
	// Required: true
	ServiceCharge *string `json:"service_charge"`

	// timely duration unit
	// Required: true
	TimelyDurationUnit *string `json:"timely_duration_unit"`

	// timely duration value
	// Required: true
	TimelyDurationValue *string `json:"timely_duration_value"`

	// topup lock level
	// Required: true
	TopupLockLevel *string `json:"topup_lock_level"`

	// topup profiles
	// Required: true
	TopupProfiles []*TopupProfile `json:"topup_profiles"`

	// underrun lock level
	// Required: true
	UnderrunLockLevel *string `json:"underrun_lock_level"`

	// underrun lock threshold
	// Required: true
	UnderrunLockThreshold *string `json:"underrun_lock_threshold"`

	// underrun profile threshold
	// Required: true
	UnderrunProfileThreshold *string `json:"underrun_profile_threshold"`

	// underrun profiles
	// Required: true
	UnderrunProfiles []*UnderrunProfile `json:"underrun_profiles"`
}

// Validate validates this thenewlycreateditemorempty33
func (m *Thenewlycreateditemorempty33) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBalanceIntervalStartMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBalanceIntervalUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBalanceIntervalValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCarryOverMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitialBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitialProfiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotopupDiscardIntervals(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResellerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceCharge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimelyDurationUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimelyDurationValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopupLockLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopupProfiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnderrunLockLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnderrunLockThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnderrunProfileThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnderrunProfiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Thenewlycreateditemorempty33) validateBalanceIntervalStartMode(formats strfmt.Registry) error {

	if err := validate.Required("balance_interval_start_mode", "body", m.BalanceIntervalStartMode); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty33) validateBalanceIntervalUnit(formats strfmt.Registry) error {

	if err := validate.Required("balance_interval_unit", "body", m.BalanceIntervalUnit); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty33) validateBalanceIntervalValue(formats strfmt.Registry) error {

	if err := validate.Required("balance_interval_value", "body", m.BalanceIntervalValue); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty33) validateCarryOverMode(formats strfmt.Registry) error {

	if err := validate.Required("carry_over_mode", "body", m.CarryOverMode); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty33) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty33) validateInitialBalance(formats strfmt.Registry) error {

	if err := validate.Required("initial_balance", "body", m.InitialBalance); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty33) validateInitialProfiles(formats strfmt.Registry) error {

	if err := validate.Required("initial_profiles", "body", m.InitialProfiles); err != nil {
		return err
	}

	for i := 0; i < len(m.InitialProfiles); i++ {
		if swag.IsZero(m.InitialProfiles[i]) { // not required
			continue
		}

		if m.InitialProfiles[i] != nil {
			if err := m.InitialProfiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("initial_profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Thenewlycreateditemorempty33) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty33) validateNotopupDiscardIntervals(formats strfmt.Registry) error {

	if err := validate.Required("notopup_discard_intervals", "body", m.NotopupDiscardIntervals); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty33) validateResellerID(formats strfmt.Registry) error {

	if err := validate.Required("reseller_id", "body", m.ResellerID); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty33) validateServiceCharge(formats strfmt.Registry) error {

	if err := validate.Required("service_charge", "body", m.ServiceCharge); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty33) validateTimelyDurationUnit(formats strfmt.Registry) error {

	if err := validate.Required("timely_duration_unit", "body", m.TimelyDurationUnit); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty33) validateTimelyDurationValue(formats strfmt.Registry) error {

	if err := validate.Required("timely_duration_value", "body", m.TimelyDurationValue); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty33) validateTopupLockLevel(formats strfmt.Registry) error {

	if err := validate.Required("topup_lock_level", "body", m.TopupLockLevel); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty33) validateTopupProfiles(formats strfmt.Registry) error {

	if err := validate.Required("topup_profiles", "body", m.TopupProfiles); err != nil {
		return err
	}

	for i := 0; i < len(m.TopupProfiles); i++ {
		if swag.IsZero(m.TopupProfiles[i]) { // not required
			continue
		}

		if m.TopupProfiles[i] != nil {
			if err := m.TopupProfiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("topup_profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Thenewlycreateditemorempty33) validateUnderrunLockLevel(formats strfmt.Registry) error {

	if err := validate.Required("underrun_lock_level", "body", m.UnderrunLockLevel); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty33) validateUnderrunLockThreshold(formats strfmt.Registry) error {

	if err := validate.Required("underrun_lock_threshold", "body", m.UnderrunLockThreshold); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty33) validateUnderrunProfileThreshold(formats strfmt.Registry) error {

	if err := validate.Required("underrun_profile_threshold", "body", m.UnderrunProfileThreshold); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty33) validateUnderrunProfiles(formats strfmt.Registry) error {

	if err := validate.Required("underrun_profiles", "body", m.UnderrunProfiles); err != nil {
		return err
	}

	for i := 0; i < len(m.UnderrunProfiles); i++ {
		if swag.IsZero(m.UnderrunProfiles[i]) { // not required
			continue
		}

		if m.UnderrunProfiles[i] != nil {
			if err := m.UnderrunProfiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("underrun_profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Thenewlycreateditemorempty33) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Thenewlycreateditemorempty33) UnmarshalBinary(b []byte) error {
	var res Thenewlycreateditemorempty33
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
