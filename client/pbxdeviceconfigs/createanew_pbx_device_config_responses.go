// Code generated by go-swagger; DO NOT EDIT.

package pbxdeviceconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// CreateanewPbxDeviceConfigReader is a Reader for the CreateanewPbxDeviceConfig structure.
type CreateanewPbxDeviceConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateanewPbxDeviceConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateanewPbxDeviceConfigCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateanewPbxDeviceConfigUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateanewPbxDeviceConfigCreated creates a CreateanewPbxDeviceConfigCreated with default headers values
func NewCreateanewPbxDeviceConfigCreated() *CreateanewPbxDeviceConfigCreated {
	return &CreateanewPbxDeviceConfigCreated{
		Location: "nostrud velit non",
	}
}

/*CreateanewPbxDeviceConfigCreated handles this case with default header values.

CreateanewPbxDeviceConfigCreated createanew pbx device config created
*/
type CreateanewPbxDeviceConfigCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty24
}

func (o *CreateanewPbxDeviceConfigCreated) Error() string {
	return fmt.Sprintf("[POST /pbxdeviceconfigs][%d] createanewPbxDeviceConfigCreated  %+v", 201, o.Payload)
}

func (o *CreateanewPbxDeviceConfigCreated) GetPayload() []*models.Thenewlycreateditemorempty24 {
	return o.Payload
}

func (o *CreateanewPbxDeviceConfigCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateanewPbxDeviceConfigUnprocessableEntity creates a CreateanewPbxDeviceConfigUnprocessableEntity with default headers values
func NewCreateanewPbxDeviceConfigUnprocessableEntity() *CreateanewPbxDeviceConfigUnprocessableEntity {
	return &CreateanewPbxDeviceConfigUnprocessableEntity{}
}

/*CreateanewPbxDeviceConfigUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CreateanewPbxDeviceConfigUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CreateanewPbxDeviceConfigUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pbxdeviceconfigs][%d] createanewPbxDeviceConfigUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateanewPbxDeviceConfigUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CreateanewPbxDeviceConfigUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
