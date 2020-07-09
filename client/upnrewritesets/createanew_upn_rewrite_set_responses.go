// Code generated by go-swagger; DO NOT EDIT.

package upnrewritesets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// CreateanewUpnRewriteSetReader is a Reader for the CreateanewUpnRewriteSet structure.
type CreateanewUpnRewriteSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateanewUpnRewriteSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateanewUpnRewriteSetCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateanewUpnRewriteSetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateanewUpnRewriteSetCreated creates a CreateanewUpnRewriteSetCreated with default headers values
func NewCreateanewUpnRewriteSetCreated() *CreateanewUpnRewriteSetCreated {
	return &CreateanewUpnRewriteSetCreated{
		Location: "nostrud velit non",
	}
}

/*CreateanewUpnRewriteSetCreated handles this case with default header values.

CreateanewUpnRewriteSetCreated createanew upn rewrite set created
*/
type CreateanewUpnRewriteSetCreated struct {
	Location string

	Payload string
}

func (o *CreateanewUpnRewriteSetCreated) Error() string {
	return fmt.Sprintf("[POST /upnrewritesets][%d] createanewUpnRewriteSetCreated  %+v", 201, o.Payload)
}

func (o *CreateanewUpnRewriteSetCreated) GetPayload() string {
	return o.Payload
}

func (o *CreateanewUpnRewriteSetCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateanewUpnRewriteSetUnprocessableEntity creates a CreateanewUpnRewriteSetUnprocessableEntity with default headers values
func NewCreateanewUpnRewriteSetUnprocessableEntity() *CreateanewUpnRewriteSetUnprocessableEntity {
	return &CreateanewUpnRewriteSetUnprocessableEntity{}
}

/*CreateanewUpnRewriteSetUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CreateanewUpnRewriteSetUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CreateanewUpnRewriteSetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /upnrewritesets][%d] createanewUpnRewriteSetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateanewUpnRewriteSetUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CreateanewUpnRewriteSetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
