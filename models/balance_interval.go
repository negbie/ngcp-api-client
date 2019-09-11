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

// BalanceInterval BalanceInterval
// swagger:model BalanceInterval
type BalanceInterval struct {

	// billing profile id
	// Required: true
	BillingProfileID *string `json:"billing_profile_id"`

	// cash balance
	// Required: true
	CashBalance *string `json:"cash_balance"`

	// cash debit
	// Required: true
	CashDebit *string `json:"cash_debit"`

	// free time balance
	// Required: true
	FreeTimeBalance *string `json:"free_time_balance"`

	// free time spent
	// Required: true
	FreeTimeSpent *string `json:"free_time_spent"`

	// initial cash balance
	// Required: true
	InitialCashBalance *string `json:"initial_cash_balance"`

	// initial free time balance
	// Required: true
	InitialFreeTimeBalance *string `json:"initial_free_time_balance"`

	// is actual
	// Required: true
	IsActual *string `json:"is_actual"`

	// notopup discard expiry
	// Required: true
	NotopupDiscardExpiry *string `json:"notopup_discard_expiry"`

	// start
	// Required: true
	Start *string `json:"start"`

	// stop
	// Required: true
	Stop *string `json:"stop"`

	// timely topup count
	// Required: true
	TimelyTopupCount *string `json:"timely_topup_count"`

	// timely topup start
	// Required: true
	TimelyTopupStart *string `json:"timely_topup_start"`

	// timely topup stop
	// Required: true
	TimelyTopupStop *string `json:"timely_topup_stop"`

	// topup count
	// Required: true
	TopupCount *string `json:"topup_count"`

	// underrun lock
	// Required: true
	UnderrunLock *string `json:"underrun_lock"`

	// underrun profiles
	// Required: true
	UnderrunProfiles *string `json:"underrun_profiles"`
}

// Validate validates this balance interval
func (m *BalanceInterval) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingProfileID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCashBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCashDebit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFreeTimeBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFreeTimeSpent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitialCashBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitialFreeTimeBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsActual(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotopupDiscardExpiry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStop(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimelyTopupCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimelyTopupStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimelyTopupStop(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopupCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnderrunLock(formats); err != nil {
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

func (m *BalanceInterval) validateBillingProfileID(formats strfmt.Registry) error {

	if err := validate.Required("billing_profile_id", "body", m.BillingProfileID); err != nil {
		return err
	}

	return nil
}

func (m *BalanceInterval) validateCashBalance(formats strfmt.Registry) error {

	if err := validate.Required("cash_balance", "body", m.CashBalance); err != nil {
		return err
	}

	return nil
}

func (m *BalanceInterval) validateCashDebit(formats strfmt.Registry) error {

	if err := validate.Required("cash_debit", "body", m.CashDebit); err != nil {
		return err
	}

	return nil
}

func (m *BalanceInterval) validateFreeTimeBalance(formats strfmt.Registry) error {

	if err := validate.Required("free_time_balance", "body", m.FreeTimeBalance); err != nil {
		return err
	}

	return nil
}

func (m *BalanceInterval) validateFreeTimeSpent(formats strfmt.Registry) error {

	if err := validate.Required("free_time_spent", "body", m.FreeTimeSpent); err != nil {
		return err
	}

	return nil
}

func (m *BalanceInterval) validateInitialCashBalance(formats strfmt.Registry) error {

	if err := validate.Required("initial_cash_balance", "body", m.InitialCashBalance); err != nil {
		return err
	}

	return nil
}

func (m *BalanceInterval) validateInitialFreeTimeBalance(formats strfmt.Registry) error {

	if err := validate.Required("initial_free_time_balance", "body", m.InitialFreeTimeBalance); err != nil {
		return err
	}

	return nil
}

func (m *BalanceInterval) validateIsActual(formats strfmt.Registry) error {

	if err := validate.Required("is_actual", "body", m.IsActual); err != nil {
		return err
	}

	return nil
}

func (m *BalanceInterval) validateNotopupDiscardExpiry(formats strfmt.Registry) error {

	if err := validate.Required("notopup_discard_expiry", "body", m.NotopupDiscardExpiry); err != nil {
		return err
	}

	return nil
}

func (m *BalanceInterval) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("start", "body", m.Start); err != nil {
		return err
	}

	return nil
}

func (m *BalanceInterval) validateStop(formats strfmt.Registry) error {

	if err := validate.Required("stop", "body", m.Stop); err != nil {
		return err
	}

	return nil
}

func (m *BalanceInterval) validateTimelyTopupCount(formats strfmt.Registry) error {

	if err := validate.Required("timely_topup_count", "body", m.TimelyTopupCount); err != nil {
		return err
	}

	return nil
}

func (m *BalanceInterval) validateTimelyTopupStart(formats strfmt.Registry) error {

	if err := validate.Required("timely_topup_start", "body", m.TimelyTopupStart); err != nil {
		return err
	}

	return nil
}

func (m *BalanceInterval) validateTimelyTopupStop(formats strfmt.Registry) error {

	if err := validate.Required("timely_topup_stop", "body", m.TimelyTopupStop); err != nil {
		return err
	}

	return nil
}

func (m *BalanceInterval) validateTopupCount(formats strfmt.Registry) error {

	if err := validate.Required("topup_count", "body", m.TopupCount); err != nil {
		return err
	}

	return nil
}

func (m *BalanceInterval) validateUnderrunLock(formats strfmt.Registry) error {

	if err := validate.Required("underrun_lock", "body", m.UnderrunLock); err != nil {
		return err
	}

	return nil
}

func (m *BalanceInterval) validateUnderrunProfiles(formats strfmt.Registry) error {

	if err := validate.Required("underrun_profiles", "body", m.UnderrunProfiles); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BalanceInterval) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BalanceInterval) UnmarshalBinary(b []byte) error {
	var res BalanceInterval
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
