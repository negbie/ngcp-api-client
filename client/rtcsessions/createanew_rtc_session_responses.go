// Code generated by go-swagger; DO NOT EDIT.

package rtcsessions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// CreateanewRtcSessionReader is a Reader for the CreateanewRtcSession structure.
type CreateanewRtcSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateanewRtcSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateanewRtcSessionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateanewRtcSessionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateanewRtcSessionCreated creates a CreateanewRtcSessionCreated with default headers values
func NewCreateanewRtcSessionCreated() *CreateanewRtcSessionCreated {
	return &CreateanewRtcSessionCreated{
		Location: "nostrud velit non",
	}
}

/*CreateanewRtcSessionCreated handles this case with default header values.

CreateanewRtcSessionCreated createanew rtc session created
*/
type CreateanewRtcSessionCreated struct {
	Location string

	Payload []string
}

func (o *CreateanewRtcSessionCreated) Error() string {
	return fmt.Sprintf("[POST /rtcsessions][%d] createanewRtcSessionCreated  %+v", 201, o.Payload)
}

func (o *CreateanewRtcSessionCreated) GetPayload() []string {
	return o.Payload
}

func (o *CreateanewRtcSessionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateanewRtcSessionUnprocessableEntity creates a CreateanewRtcSessionUnprocessableEntity with default headers values
func NewCreateanewRtcSessionUnprocessableEntity() *CreateanewRtcSessionUnprocessableEntity {
	return &CreateanewRtcSessionUnprocessableEntity{}
}

/*CreateanewRtcSessionUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CreateanewRtcSessionUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CreateanewRtcSessionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /rtcsessions][%d] createanewRtcSessionUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateanewRtcSessionUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CreateanewRtcSessionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
