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

// CfdestinationsetsPostReader is a Reader for the CfdestinationsetsPost structure.
type CfdestinationsetsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CfdestinationsetsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCfdestinationsetsPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCfdestinationsetsPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCfdestinationsetsPostCreated creates a CfdestinationsetsPostCreated with default headers values
func NewCfdestinationsetsPostCreated() *CfdestinationsetsPostCreated {
	return &CfdestinationsetsPostCreated{
		Location: "\"<string>\"",
	}
}

/*CfdestinationsetsPostCreated handles this case with default header values.

CfdestinationsetsPostCreated cfdestinationsets post created
*/
type CfdestinationsetsPostCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty38
}

func (o *CfdestinationsetsPostCreated) Error() string {
	return fmt.Sprintf("[POST /cfdestinationsets/][%d] cfdestinationsetsPostCreated  %+v", 201, o.Payload)
}

func (o *CfdestinationsetsPostCreated) GetPayload() []*models.Thenewlycreateditemorempty38 {
	return o.Payload
}

func (o *CfdestinationsetsPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCfdestinationsetsPostUnprocessableEntity creates a CfdestinationsetsPostUnprocessableEntity with default headers values
func NewCfdestinationsetsPostUnprocessableEntity() *CfdestinationsetsPostUnprocessableEntity {
	return &CfdestinationsetsPostUnprocessableEntity{}
}

/*CfdestinationsetsPostUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CfdestinationsetsPostUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CfdestinationsetsPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /cfdestinationsets/][%d] cfdestinationsetsPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CfdestinationsetsPostUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CfdestinationsetsPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}