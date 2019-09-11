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

// Event Event
// swagger:model Event
type Event struct {

	// export status
	// Required: true
	ExportStatus *string `json:"export_status"`

	// exported at
	// Required: true
	ExportedAt *string `json:"exported_at"`

	// first non primary alias username after
	// Required: true
	FirstNonPrimaryAliasUsernameAfter *string `json:"first_non_primary_alias_username_after"`

	// first non primary alias username before
	// Required: true
	FirstNonPrimaryAliasUsernameBefore *string `json:"first_non_primary_alias_username_before"`

	// new status
	// Required: true
	NewStatus *string `json:"new_status"`

	// non primary alias username
	// Required: true
	NonPrimaryAliasUsername *string `json:"non_primary_alias_username"`

	// old status
	// Required: true
	OldStatus *string `json:"old_status"`

	// pilot first non primary alias username after
	// Required: true
	PilotFirstNonPrimaryAliasUsernameAfter *string `json:"pilot_first_non_primary_alias_username_after"`

	// pilot first non primary alias username before
	// Required: true
	PilotFirstNonPrimaryAliasUsernameBefore *string `json:"pilot_first_non_primary_alias_username_before"`

	// pilot primary alias username after
	// Required: true
	PilotPrimaryAliasUsernameAfter *string `json:"pilot_primary_alias_username_after"`

	// pilot primary alias username before
	// Required: true
	PilotPrimaryAliasUsernameBefore *string `json:"pilot_primary_alias_username_before"`

	// pilot primary number ac
	// Required: true
	PilotPrimaryNumberAc *string `json:"pilot_primary_number_ac"`

	// pilot primary number cc
	// Required: true
	PilotPrimaryNumberCc *string `json:"pilot_primary_number_cc"`

	// pilot primary number id
	// Required: true
	PilotPrimaryNumberID *string `json:"pilot_primary_number_id"`

	// pilot primary number sn
	// Required: true
	PilotPrimaryNumberSn *string `json:"pilot_primary_number_sn"`

	// pilot subscriber id
	// Required: true
	PilotSubscriberID *string `json:"pilot_subscriber_id"`

	// pilot subscriber profile id
	// Required: true
	PilotSubscriberProfileID *string `json:"pilot_subscriber_profile_id"`

	// pilot subscriber profile name
	// Required: true
	PilotSubscriberProfileName *string `json:"pilot_subscriber_profile_name"`

	// pilot subscriber profile set id
	// Required: true
	PilotSubscriberProfileSetID *string `json:"pilot_subscriber_profile_set_id"`

	// pilot subscriber profile set name
	// Required: true
	PilotSubscriberProfileSetName *string `json:"pilot_subscriber_profile_set_name"`

	// primary alias username after
	// Required: true
	PrimaryAliasUsernameAfter *string `json:"primary_alias_username_after"`

	// primary alias username before
	// Required: true
	PrimaryAliasUsernameBefore *string `json:"primary_alias_username_before"`

	// primary number ac
	// Required: true
	PrimaryNumberAc *string `json:"primary_number_ac"`

	// primary number cc
	// Required: true
	PrimaryNumberCc *string `json:"primary_number_cc"`

	// primary number id
	// Required: true
	PrimaryNumberID *string `json:"primary_number_id"`

	// primary number sn
	// Required: true
	PrimaryNumberSn *string `json:"primary_number_sn"`

	// reseller id
	// Required: true
	ResellerID *string `json:"reseller_id"`

	// subscriber id
	// Required: true
	SubscriberID *string `json:"subscriber_id"`

	// subscriber profile id
	// Required: true
	SubscriberProfileID *string `json:"subscriber_profile_id"`

	// subscriber profile name
	// Required: true
	SubscriberProfileName *string `json:"subscriber_profile_name"`

	// subscriber profile set id
	// Required: true
	SubscriberProfileSetID *string `json:"subscriber_profile_set_id"`

	// subscriber profile set name
	// Required: true
	SubscriberProfileSetName *string `json:"subscriber_profile_set_name"`

	// timestamp
	// Required: true
	Timestamp *string `json:"timestamp"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExportStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstNonPrimaryAliasUsernameAfter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstNonPrimaryAliasUsernameBefore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNonPrimaryAliasUsername(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOldStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotFirstNonPrimaryAliasUsernameAfter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotFirstNonPrimaryAliasUsernameBefore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotPrimaryAliasUsernameAfter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotPrimaryAliasUsernameBefore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotPrimaryNumberAc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotPrimaryNumberCc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotPrimaryNumberID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotPrimaryNumberSn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotSubscriberID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotSubscriberProfileID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotSubscriberProfileName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotSubscriberProfileSetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotSubscriberProfileSetName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryAliasUsernameAfter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryAliasUsernameBefore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryNumberAc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryNumberCc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryNumberID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryNumberSn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResellerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberProfileID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberProfileName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberProfileSetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberProfileSetName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Event) validateExportStatus(formats strfmt.Registry) error {

	if err := validate.Required("export_status", "body", m.ExportStatus); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateExportedAt(formats strfmt.Registry) error {

	if err := validate.Required("exported_at", "body", m.ExportedAt); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateFirstNonPrimaryAliasUsernameAfter(formats strfmt.Registry) error {

	if err := validate.Required("first_non_primary_alias_username_after", "body", m.FirstNonPrimaryAliasUsernameAfter); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateFirstNonPrimaryAliasUsernameBefore(formats strfmt.Registry) error {

	if err := validate.Required("first_non_primary_alias_username_before", "body", m.FirstNonPrimaryAliasUsernameBefore); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateNewStatus(formats strfmt.Registry) error {

	if err := validate.Required("new_status", "body", m.NewStatus); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateNonPrimaryAliasUsername(formats strfmt.Registry) error {

	if err := validate.Required("non_primary_alias_username", "body", m.NonPrimaryAliasUsername); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateOldStatus(formats strfmt.Registry) error {

	if err := validate.Required("old_status", "body", m.OldStatus); err != nil {
		return err
	}

	return nil
}

func (m *Event) validatePilotFirstNonPrimaryAliasUsernameAfter(formats strfmt.Registry) error {

	if err := validate.Required("pilot_first_non_primary_alias_username_after", "body", m.PilotFirstNonPrimaryAliasUsernameAfter); err != nil {
		return err
	}

	return nil
}

func (m *Event) validatePilotFirstNonPrimaryAliasUsernameBefore(formats strfmt.Registry) error {

	if err := validate.Required("pilot_first_non_primary_alias_username_before", "body", m.PilotFirstNonPrimaryAliasUsernameBefore); err != nil {
		return err
	}

	return nil
}

func (m *Event) validatePilotPrimaryAliasUsernameAfter(formats strfmt.Registry) error {

	if err := validate.Required("pilot_primary_alias_username_after", "body", m.PilotPrimaryAliasUsernameAfter); err != nil {
		return err
	}

	return nil
}

func (m *Event) validatePilotPrimaryAliasUsernameBefore(formats strfmt.Registry) error {

	if err := validate.Required("pilot_primary_alias_username_before", "body", m.PilotPrimaryAliasUsernameBefore); err != nil {
		return err
	}

	return nil
}

func (m *Event) validatePilotPrimaryNumberAc(formats strfmt.Registry) error {

	if err := validate.Required("pilot_primary_number_ac", "body", m.PilotPrimaryNumberAc); err != nil {
		return err
	}

	return nil
}

func (m *Event) validatePilotPrimaryNumberCc(formats strfmt.Registry) error {

	if err := validate.Required("pilot_primary_number_cc", "body", m.PilotPrimaryNumberCc); err != nil {
		return err
	}

	return nil
}

func (m *Event) validatePilotPrimaryNumberID(formats strfmt.Registry) error {

	if err := validate.Required("pilot_primary_number_id", "body", m.PilotPrimaryNumberID); err != nil {
		return err
	}

	return nil
}

func (m *Event) validatePilotPrimaryNumberSn(formats strfmt.Registry) error {

	if err := validate.Required("pilot_primary_number_sn", "body", m.PilotPrimaryNumberSn); err != nil {
		return err
	}

	return nil
}

func (m *Event) validatePilotSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("pilot_subscriber_id", "body", m.PilotSubscriberID); err != nil {
		return err
	}

	return nil
}

func (m *Event) validatePilotSubscriberProfileID(formats strfmt.Registry) error {

	if err := validate.Required("pilot_subscriber_profile_id", "body", m.PilotSubscriberProfileID); err != nil {
		return err
	}

	return nil
}

func (m *Event) validatePilotSubscriberProfileName(formats strfmt.Registry) error {

	if err := validate.Required("pilot_subscriber_profile_name", "body", m.PilotSubscriberProfileName); err != nil {
		return err
	}

	return nil
}

func (m *Event) validatePilotSubscriberProfileSetID(formats strfmt.Registry) error {

	if err := validate.Required("pilot_subscriber_profile_set_id", "body", m.PilotSubscriberProfileSetID); err != nil {
		return err
	}

	return nil
}

func (m *Event) validatePilotSubscriberProfileSetName(formats strfmt.Registry) error {

	if err := validate.Required("pilot_subscriber_profile_set_name", "body", m.PilotSubscriberProfileSetName); err != nil {
		return err
	}

	return nil
}

func (m *Event) validatePrimaryAliasUsernameAfter(formats strfmt.Registry) error {

	if err := validate.Required("primary_alias_username_after", "body", m.PrimaryAliasUsernameAfter); err != nil {
		return err
	}

	return nil
}

func (m *Event) validatePrimaryAliasUsernameBefore(formats strfmt.Registry) error {

	if err := validate.Required("primary_alias_username_before", "body", m.PrimaryAliasUsernameBefore); err != nil {
		return err
	}

	return nil
}

func (m *Event) validatePrimaryNumberAc(formats strfmt.Registry) error {

	if err := validate.Required("primary_number_ac", "body", m.PrimaryNumberAc); err != nil {
		return err
	}

	return nil
}

func (m *Event) validatePrimaryNumberCc(formats strfmt.Registry) error {

	if err := validate.Required("primary_number_cc", "body", m.PrimaryNumberCc); err != nil {
		return err
	}

	return nil
}

func (m *Event) validatePrimaryNumberID(formats strfmt.Registry) error {

	if err := validate.Required("primary_number_id", "body", m.PrimaryNumberID); err != nil {
		return err
	}

	return nil
}

func (m *Event) validatePrimaryNumberSn(formats strfmt.Registry) error {

	if err := validate.Required("primary_number_sn", "body", m.PrimaryNumberSn); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateResellerID(formats strfmt.Registry) error {

	if err := validate.Required("reseller_id", "body", m.ResellerID); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_id", "body", m.SubscriberID); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateSubscriberProfileID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_profile_id", "body", m.SubscriberProfileID); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateSubscriberProfileName(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_profile_name", "body", m.SubscriberProfileName); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateSubscriberProfileSetID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_profile_set_id", "body", m.SubscriberProfileSetID); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateSubscriberProfileSetName(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_profile_set_name", "body", m.SubscriberProfileSetName); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Event) UnmarshalBinary(b []byte) error {
	var res Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
