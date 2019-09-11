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

// Thenewlycreateditemorempty1 Thenewlycreateditemorempty1
// swagger:model Thenewlycreateditemorempty1
type Thenewlycreateditemorempty1 struct {

	// billing profile id
	// Required: true
	BillingProfileID *string `json:"billing_profile_id"`

	// billing zone id
	// Required: true
	BillingZoneID *string `json:"billing_zone_id"`

	// destination
	// Required: true
	Destination *string `json:"destination"`

	// direction
	// Required: true
	Direction *string `json:"direction"`

	// match mode
	// Required: true
	MatchMode *string `json:"match_mode"`

	// offpeak follow interval
	// Required: true
	OffpeakFollowInterval *string `json:"offpeak_follow_interval"`

	// offpeak follow rate
	// Required: true
	OffpeakFollowRate *string `json:"offpeak_follow_rate"`

	// offpeak init interval
	// Required: true
	OffpeakInitInterval *string `json:"offpeak_init_interval"`

	// offpeak init rate
	// Required: true
	OffpeakInitRate *string `json:"offpeak_init_rate"`

	// onpeak follow interval
	// Required: true
	OnpeakFollowInterval *string `json:"onpeak_follow_interval"`

	// onpeak follow rate
	// Required: true
	OnpeakFollowRate *string `json:"onpeak_follow_rate"`

	// onpeak init interval
	// Required: true
	OnpeakInitInterval *string `json:"onpeak_init_interval"`

	// onpeak init rate
	// Required: true
	OnpeakInitRate *string `json:"onpeak_init_rate"`

	// purge existing
	// Required: true
	PurgeExisting *string `json:"purge_existing"`

	// source
	// Required: true
	Source *string `json:"source"`

	// use free time
	// Required: true
	UseFreeTime *string `json:"use_free_time"`
}

// Validate validates this thenewlycreateditemorempty1
func (m *Thenewlycreateditemorempty1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingProfileID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingZoneID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatchMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOffpeakFollowInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOffpeakFollowRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOffpeakInitInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOffpeakInitRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnpeakFollowInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnpeakFollowRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnpeakInitInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnpeakInitRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurgeExisting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseFreeTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Thenewlycreateditemorempty1) validateBillingProfileID(formats strfmt.Registry) error {

	if err := validate.Required("billing_profile_id", "body", m.BillingProfileID); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty1) validateBillingZoneID(formats strfmt.Registry) error {

	if err := validate.Required("billing_zone_id", "body", m.BillingZoneID); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty1) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("destination", "body", m.Destination); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty1) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty1) validateMatchMode(formats strfmt.Registry) error {

	if err := validate.Required("match_mode", "body", m.MatchMode); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty1) validateOffpeakFollowInterval(formats strfmt.Registry) error {

	if err := validate.Required("offpeak_follow_interval", "body", m.OffpeakFollowInterval); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty1) validateOffpeakFollowRate(formats strfmt.Registry) error {

	if err := validate.Required("offpeak_follow_rate", "body", m.OffpeakFollowRate); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty1) validateOffpeakInitInterval(formats strfmt.Registry) error {

	if err := validate.Required("offpeak_init_interval", "body", m.OffpeakInitInterval); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty1) validateOffpeakInitRate(formats strfmt.Registry) error {

	if err := validate.Required("offpeak_init_rate", "body", m.OffpeakInitRate); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty1) validateOnpeakFollowInterval(formats strfmt.Registry) error {

	if err := validate.Required("onpeak_follow_interval", "body", m.OnpeakFollowInterval); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty1) validateOnpeakFollowRate(formats strfmt.Registry) error {

	if err := validate.Required("onpeak_follow_rate", "body", m.OnpeakFollowRate); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty1) validateOnpeakInitInterval(formats strfmt.Registry) error {

	if err := validate.Required("onpeak_init_interval", "body", m.OnpeakInitInterval); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty1) validateOnpeakInitRate(formats strfmt.Registry) error {

	if err := validate.Required("onpeak_init_rate", "body", m.OnpeakInitRate); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty1) validatePurgeExisting(formats strfmt.Registry) error {

	if err := validate.Required("purge_existing", "body", m.PurgeExisting); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty1) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty1) validateUseFreeTime(formats strfmt.Registry) error {

	if err := validate.Required("use_free_time", "body", m.UseFreeTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Thenewlycreateditemorempty1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Thenewlycreateditemorempty1) UnmarshalBinary(b []byte) error {
	var res Thenewlycreateditemorempty1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
