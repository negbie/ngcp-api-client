// Code generated by go-swagger; DO NOT EDIT.

package invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// CreateanewInvoiceReader is a Reader for the CreateanewInvoice structure.
type CreateanewInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateanewInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateanewInvoiceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateanewInvoiceUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateanewInvoiceCreated creates a CreateanewInvoiceCreated with default headers values
func NewCreateanewInvoiceCreated() *CreateanewInvoiceCreated {
	return &CreateanewInvoiceCreated{
		Location: "nostrud velit non",
	}
}

/*CreateanewInvoiceCreated handles this case with default header values.

CreateanewInvoiceCreated createanew invoice created
*/
type CreateanewInvoiceCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty45
}

func (o *CreateanewInvoiceCreated) Error() string {
	return fmt.Sprintf("[POST /invoices][%d] createanewInvoiceCreated  %+v", 201, o.Payload)
}

func (o *CreateanewInvoiceCreated) GetPayload() []*models.Thenewlycreateditemorempty45 {
	return o.Payload
}

func (o *CreateanewInvoiceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateanewInvoiceUnprocessableEntity creates a CreateanewInvoiceUnprocessableEntity with default headers values
func NewCreateanewInvoiceUnprocessableEntity() *CreateanewInvoiceUnprocessableEntity {
	return &CreateanewInvoiceUnprocessableEntity{}
}

/*CreateanewInvoiceUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CreateanewInvoiceUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CreateanewInvoiceUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /invoices][%d] createanewInvoiceUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateanewInvoiceUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CreateanewInvoiceUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
