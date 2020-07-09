// Code generated by go-swagger; DO NOT EDIT.

package pbxdevices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// CreateanewPbxDeviceReader is a Reader for the CreateanewPbxDevice structure.
type CreateanewPbxDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateanewPbxDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateanewPbxDeviceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateanewPbxDeviceUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateanewPbxDeviceCreated creates a CreateanewPbxDeviceCreated with default headers values
func NewCreateanewPbxDeviceCreated() *CreateanewPbxDeviceCreated {
	return &CreateanewPbxDeviceCreated{
		Location: "nostrud velit non",
	}
}

/*CreateanewPbxDeviceCreated handles this case with default header values.

CreateanewPbxDeviceCreated createanew pbx device created
*/
type CreateanewPbxDeviceCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty23
}

func (o *CreateanewPbxDeviceCreated) Error() string {
	return fmt.Sprintf("[POST /pbxdevices][%d] createanewPbxDeviceCreated  %+v", 201, o.Payload)
}

func (o *CreateanewPbxDeviceCreated) GetPayload() []*models.Thenewlycreateditemorempty23 {
	return o.Payload
}

func (o *CreateanewPbxDeviceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateanewPbxDeviceUnprocessableEntity creates a CreateanewPbxDeviceUnprocessableEntity with default headers values
func NewCreateanewPbxDeviceUnprocessableEntity() *CreateanewPbxDeviceUnprocessableEntity {
	return &CreateanewPbxDeviceUnprocessableEntity{}
}

/*CreateanewPbxDeviceUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CreateanewPbxDeviceUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CreateanewPbxDeviceUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pbxdevices][%d] createanewPbxDeviceUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateanewPbxDeviceUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CreateanewPbxDeviceUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
