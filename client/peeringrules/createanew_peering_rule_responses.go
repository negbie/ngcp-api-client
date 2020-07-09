// Code generated by go-swagger; DO NOT EDIT.

package peeringrules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// CreateanewPeeringRuleReader is a Reader for the CreateanewPeeringRule structure.
type CreateanewPeeringRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateanewPeeringRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateanewPeeringRuleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateanewPeeringRuleUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateanewPeeringRuleCreated creates a CreateanewPeeringRuleCreated with default headers values
func NewCreateanewPeeringRuleCreated() *CreateanewPeeringRuleCreated {
	return &CreateanewPeeringRuleCreated{
		Location: "nostrud velit non",
	}
}

/*CreateanewPeeringRuleCreated handles this case with default header values.

CreateanewPeeringRuleCreated createanew peering rule created
*/
type CreateanewPeeringRuleCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty35
}

func (o *CreateanewPeeringRuleCreated) Error() string {
	return fmt.Sprintf("[POST /peeringrules][%d] createanewPeeringRuleCreated  %+v", 201, o.Payload)
}

func (o *CreateanewPeeringRuleCreated) GetPayload() []*models.Thenewlycreateditemorempty35 {
	return o.Payload
}

func (o *CreateanewPeeringRuleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateanewPeeringRuleUnprocessableEntity creates a CreateanewPeeringRuleUnprocessableEntity with default headers values
func NewCreateanewPeeringRuleUnprocessableEntity() *CreateanewPeeringRuleUnprocessableEntity {
	return &CreateanewPeeringRuleUnprocessableEntity{}
}

/*CreateanewPeeringRuleUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CreateanewPeeringRuleUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CreateanewPeeringRuleUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /peeringrules][%d] createanewPeeringRuleUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateanewPeeringRuleUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CreateanewPeeringRuleUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
