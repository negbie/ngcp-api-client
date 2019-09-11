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

// PeeringrulesPostReader is a Reader for the PeeringrulesPost structure.
type PeeringrulesPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PeeringrulesPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPeeringrulesPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewPeeringrulesPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPeeringrulesPostCreated creates a PeeringrulesPostCreated with default headers values
func NewPeeringrulesPostCreated() *PeeringrulesPostCreated {
	return &PeeringrulesPostCreated{
		Location: "\"<string>\"",
	}
}

/*PeeringrulesPostCreated handles this case with default header values.

PeeringrulesPostCreated peeringrules post created
*/
type PeeringrulesPostCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty39
}

func (o *PeeringrulesPostCreated) Error() string {
	return fmt.Sprintf("[POST /peeringrules/][%d] peeringrulesPostCreated  %+v", 201, o.Payload)
}

func (o *PeeringrulesPostCreated) GetPayload() []*models.Thenewlycreateditemorempty39 {
	return o.Payload
}

func (o *PeeringrulesPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPeeringrulesPostUnprocessableEntity creates a PeeringrulesPostUnprocessableEntity with default headers values
func NewPeeringrulesPostUnprocessableEntity() *PeeringrulesPostUnprocessableEntity {
	return &PeeringrulesPostUnprocessableEntity{}
}

/*PeeringrulesPostUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type PeeringrulesPostUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *PeeringrulesPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /peeringrules/][%d] peeringrulesPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PeeringrulesPostUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *PeeringrulesPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
