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

// InterceptionsPostReader is a Reader for the InterceptionsPost structure.
type InterceptionsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterceptionsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewInterceptionsPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewInterceptionsPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInterceptionsPostCreated creates a InterceptionsPostCreated with default headers values
func NewInterceptionsPostCreated() *InterceptionsPostCreated {
	return &InterceptionsPostCreated{
		Location: "\"<string>\"",
	}
}

/*InterceptionsPostCreated handles this case with default header values.

InterceptionsPostCreated interceptions post created
*/
type InterceptionsPostCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty22
}

func (o *InterceptionsPostCreated) Error() string {
	return fmt.Sprintf("[POST /interceptions/][%d] interceptionsPostCreated  %+v", 201, o.Payload)
}

func (o *InterceptionsPostCreated) GetPayload() []*models.Thenewlycreateditemorempty22 {
	return o.Payload
}

func (o *InterceptionsPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterceptionsPostUnprocessableEntity creates a InterceptionsPostUnprocessableEntity with default headers values
func NewInterceptionsPostUnprocessableEntity() *InterceptionsPostUnprocessableEntity {
	return &InterceptionsPostUnprocessableEntity{}
}

/*InterceptionsPostUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type InterceptionsPostUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *InterceptionsPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /interceptions/][%d] interceptionsPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *InterceptionsPostUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *InterceptionsPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}