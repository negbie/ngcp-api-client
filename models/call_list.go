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

// CallList CallList
// swagger:model CallList
type CallList struct {

	// call id
	// Required: true
	CallID *string `json:"call_id"`

	// clir
	// Required: true
	Clir *string `json:"clir"`

	// customer cost
	// Required: true
	CustomerCost *string `json:"customer_cost"`

	// customer free time
	// Required: true
	CustomerFreeTime *string `json:"customer_free_time"`

	// direction
	// Required: true
	Direction *string `json:"direction"`

	// duration
	// Required: true
	Duration *string `json:"duration"`

	// init time
	// Required: true
	InitTime *string `json:"init_time"`

	// intra customer
	// Required: true
	IntraCustomer *string `json:"intra_customer"`

	// mos average
	// Required: true
	MosAverage *string `json:"mos_average"`

	// mos average jitter
	// Required: true
	MosAverageJitter *string `json:"mos_average_jitter"`

	// mos average packetloss
	// Required: true
	MosAveragePacketloss *string `json:"mos_average_packetloss"`

	// mos average roundtrip
	// Required: true
	MosAverageRoundtrip *string `json:"mos_average_roundtrip"`

	// other cli
	// Required: true
	OtherCli *string `json:"other_cli"`

	// own cli
	// Required: true
	OwnCli *string `json:"own_cli"`

	// rating status
	// Required: true
	RatingStatus *string `json:"rating_status"`

	// start time
	// Required: true
	StartTime *string `json:"start_time"`

	// status
	// Required: true
	Status *string `json:"status"`

	// total customer cost
	// Required: true
	TotalCustomerCost *string `json:"total_customer_cost"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this call list
func (m *CallList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClir(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerFreeTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntraCustomer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMosAverage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMosAverageJitter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMosAveragePacketloss(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMosAverageRoundtrip(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOtherCli(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnCli(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRatingStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCustomerCost(formats); err != nil {
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

func (m *CallList) validateCallID(formats strfmt.Registry) error {

	if err := validate.Required("call_id", "body", m.CallID); err != nil {
		return err
	}

	return nil
}

func (m *CallList) validateClir(formats strfmt.Registry) error {

	if err := validate.Required("clir", "body", m.Clir); err != nil {
		return err
	}

	return nil
}

func (m *CallList) validateCustomerCost(formats strfmt.Registry) error {

	if err := validate.Required("customer_cost", "body", m.CustomerCost); err != nil {
		return err
	}

	return nil
}

func (m *CallList) validateCustomerFreeTime(formats strfmt.Registry) error {

	if err := validate.Required("customer_free_time", "body", m.CustomerFreeTime); err != nil {
		return err
	}

	return nil
}

func (m *CallList) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *CallList) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *CallList) validateInitTime(formats strfmt.Registry) error {

	if err := validate.Required("init_time", "body", m.InitTime); err != nil {
		return err
	}

	return nil
}

func (m *CallList) validateIntraCustomer(formats strfmt.Registry) error {

	if err := validate.Required("intra_customer", "body", m.IntraCustomer); err != nil {
		return err
	}

	return nil
}

func (m *CallList) validateMosAverage(formats strfmt.Registry) error {

	if err := validate.Required("mos_average", "body", m.MosAverage); err != nil {
		return err
	}

	return nil
}

func (m *CallList) validateMosAverageJitter(formats strfmt.Registry) error {

	if err := validate.Required("mos_average_jitter", "body", m.MosAverageJitter); err != nil {
		return err
	}

	return nil
}

func (m *CallList) validateMosAveragePacketloss(formats strfmt.Registry) error {

	if err := validate.Required("mos_average_packetloss", "body", m.MosAveragePacketloss); err != nil {
		return err
	}

	return nil
}

func (m *CallList) validateMosAverageRoundtrip(formats strfmt.Registry) error {

	if err := validate.Required("mos_average_roundtrip", "body", m.MosAverageRoundtrip); err != nil {
		return err
	}

	return nil
}

func (m *CallList) validateOtherCli(formats strfmt.Registry) error {

	if err := validate.Required("other_cli", "body", m.OtherCli); err != nil {
		return err
	}

	return nil
}

func (m *CallList) validateOwnCli(formats strfmt.Registry) error {

	if err := validate.Required("own_cli", "body", m.OwnCli); err != nil {
		return err
	}

	return nil
}

func (m *CallList) validateRatingStatus(formats strfmt.Registry) error {

	if err := validate.Required("rating_status", "body", m.RatingStatus); err != nil {
		return err
	}

	return nil
}

func (m *CallList) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("start_time", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

func (m *CallList) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *CallList) validateTotalCustomerCost(formats strfmt.Registry) error {

	if err := validate.Required("total_customer_cost", "body", m.TotalCustomerCost); err != nil {
		return err
	}

	return nil
}

func (m *CallList) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CallList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CallList) UnmarshalBinary(b []byte) error {
	var res CallList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
