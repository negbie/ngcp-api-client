// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/negbie/ngcp-api/models"
)

// EmergencymappingsPostReader is a Reader for the EmergencymappingsPost structure.
type EmergencymappingsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmergencymappingsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewEmergencymappingsPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewEmergencymappingsPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEmergencymappingsPostCreated creates a EmergencymappingsPostCreated with default headers values
func NewEmergencymappingsPostCreated() *EmergencymappingsPostCreated {
	return &EmergencymappingsPostCreated{
		Location: "\"<string>\"",
	}
}

/*EmergencymappingsPostCreated handles this case with default header values.

EmergencymappingsPostCreated emergencymappings post created
*/
type EmergencymappingsPostCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty42
}

func (o *EmergencymappingsPostCreated) Error() string {
	return fmt.Sprintf("[POST /emergencymappings/][%d] emergencymappingsPostCreated  %+v", 201, o.Payload)
}

func (o *EmergencymappingsPostCreated) GetPayload() []*models.Thenewlycreateditemorempty42 {
	return o.Payload
}

func (o *EmergencymappingsPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmergencymappingsPostUnprocessableEntity creates a EmergencymappingsPostUnprocessableEntity with default headers values
func NewEmergencymappingsPostUnprocessableEntity() *EmergencymappingsPostUnprocessableEntity {
	return &EmergencymappingsPostUnprocessableEntity{}
}

/*EmergencymappingsPostUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type EmergencymappingsPostUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *EmergencymappingsPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /emergencymappings/][%d] emergencymappingsPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *EmergencymappingsPostUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *EmergencymappingsPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
