// Code generated by go-swagger; DO NOT EDIT.

package peeringgroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// CreateanewPeeringGroupReader is a Reader for the CreateanewPeeringGroup structure.
type CreateanewPeeringGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateanewPeeringGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateanewPeeringGroupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateanewPeeringGroupUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateanewPeeringGroupCreated creates a CreateanewPeeringGroupCreated with default headers values
func NewCreateanewPeeringGroupCreated() *CreateanewPeeringGroupCreated {
	return &CreateanewPeeringGroupCreated{
		Location: "nostrud velit non",
	}
}

/*CreateanewPeeringGroupCreated handles this case with default header values.

CreateanewPeeringGroupCreated createanew peering group created
*/
type CreateanewPeeringGroupCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty10
}

func (o *CreateanewPeeringGroupCreated) Error() string {
	return fmt.Sprintf("[POST /peeringgroups][%d] createanewPeeringGroupCreated  %+v", 201, o.Payload)
}

func (o *CreateanewPeeringGroupCreated) GetPayload() []*models.Thenewlycreateditemorempty10 {
	return o.Payload
}

func (o *CreateanewPeeringGroupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateanewPeeringGroupUnprocessableEntity creates a CreateanewPeeringGroupUnprocessableEntity with default headers values
func NewCreateanewPeeringGroupUnprocessableEntity() *CreateanewPeeringGroupUnprocessableEntity {
	return &CreateanewPeeringGroupUnprocessableEntity{}
}

/*CreateanewPeeringGroupUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CreateanewPeeringGroupUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CreateanewPeeringGroupUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /peeringgroups][%d] createanewPeeringGroupUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateanewPeeringGroupUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CreateanewPeeringGroupUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
