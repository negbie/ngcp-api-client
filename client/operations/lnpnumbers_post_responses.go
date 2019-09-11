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

// LnpnumbersPostReader is a Reader for the LnpnumbersPost structure.
type LnpnumbersPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LnpnumbersPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewLnpnumbersPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewLnpnumbersPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLnpnumbersPostCreated creates a LnpnumbersPostCreated with default headers values
func NewLnpnumbersPostCreated() *LnpnumbersPostCreated {
	return &LnpnumbersPostCreated{
		Location: "\"<string>\"",
	}
}

/*LnpnumbersPostCreated handles this case with default header values.

LnpnumbersPostCreated lnpnumbers post created
*/
type LnpnumbersPostCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty10
}

func (o *LnpnumbersPostCreated) Error() string {
	return fmt.Sprintf("[POST /lnpnumbers/][%d] lnpnumbersPostCreated  %+v", 201, o.Payload)
}

func (o *LnpnumbersPostCreated) GetPayload() []*models.Thenewlycreateditemorempty10 {
	return o.Payload
}

func (o *LnpnumbersPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLnpnumbersPostUnprocessableEntity creates a LnpnumbersPostUnprocessableEntity with default headers values
func NewLnpnumbersPostUnprocessableEntity() *LnpnumbersPostUnprocessableEntity {
	return &LnpnumbersPostUnprocessableEntity{}
}

/*LnpnumbersPostUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type LnpnumbersPostUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *LnpnumbersPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /lnpnumbers/][%d] lnpnumbersPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *LnpnumbersPostUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *LnpnumbersPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
