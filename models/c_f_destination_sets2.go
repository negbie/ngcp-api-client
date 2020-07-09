// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CFDestinationSets2 CFDestinationSets2
//
// swagger:model CFDestinationSets2
type CFDestinationSets2 struct {

	// destinations
	// Required: true
	Destinations []*Destination24 `json:"destinations"`

	// name
	// Required: true
	Name *string `json:"name"`

	// subscriber id
	// Required: true
	SubscriberID *float64 `json:"subscriber_id"`
}

// Validate validates this c f destination sets2
func (m *CFDestinationSets2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CFDestinationSets2) validateDestinations(formats strfmt.Registry) error {

	if err := validate.Required("destinations", "body", m.Destinations); err != nil {
		return err
	}

	for i := 0; i < len(m.Destinations); i++ {
		if swag.IsZero(m.Destinations[i]) { // not required
			continue
		}

		if m.Destinations[i] != nil {
			if err := m.Destinations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("destinations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CFDestinationSets2) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CFDestinationSets2) validateSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_id", "body", m.SubscriberID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CFDestinationSets2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CFDestinationSets2) UnmarshalBinary(b []byte) error {
	var res CFDestinationSets2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}