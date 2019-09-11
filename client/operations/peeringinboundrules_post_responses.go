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

// PeeringinboundrulesPostReader is a Reader for the PeeringinboundrulesPost structure.
type PeeringinboundrulesPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PeeringinboundrulesPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPeeringinboundrulesPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewPeeringinboundrulesPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPeeringinboundrulesPostCreated creates a PeeringinboundrulesPostCreated with default headers values
func NewPeeringinboundrulesPostCreated() *PeeringinboundrulesPostCreated {
	return &PeeringinboundrulesPostCreated{
		Location: "\"<string>\"",
	}
}

/*PeeringinboundrulesPostCreated handles this case with default header values.

PeeringinboundrulesPostCreated peeringinboundrules post created
*/
type PeeringinboundrulesPostCreated struct {
	Location string

	Payload string
}

func (o *PeeringinboundrulesPostCreated) Error() string {
	return fmt.Sprintf("[POST /peeringinboundrules/][%d] peeringinboundrulesPostCreated  %+v", 201, o.Payload)
}

func (o *PeeringinboundrulesPostCreated) GetPayload() string {
	return o.Payload
}

func (o *PeeringinboundrulesPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPeeringinboundrulesPostUnprocessableEntity creates a PeeringinboundrulesPostUnprocessableEntity with default headers values
func NewPeeringinboundrulesPostUnprocessableEntity() *PeeringinboundrulesPostUnprocessableEntity {
	return &PeeringinboundrulesPostUnprocessableEntity{}
}

/*PeeringinboundrulesPostUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type PeeringinboundrulesPostUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *PeeringinboundrulesPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /peeringinboundrules/][%d] peeringinboundrulesPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PeeringinboundrulesPostUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *PeeringinboundrulesPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}