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

// NcospatternsPostReader is a Reader for the NcospatternsPost structure.
type NcospatternsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NcospatternsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewNcospatternsPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewNcospatternsPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNcospatternsPostCreated creates a NcospatternsPostCreated with default headers values
func NewNcospatternsPostCreated() *NcospatternsPostCreated {
	return &NcospatternsPostCreated{
		Location: "\"<string>\"",
	}
}

/*NcospatternsPostCreated handles this case with default header values.

NcospatternsPostCreated ncospatterns post created
*/
type NcospatternsPostCreated struct {
	Location string

	Payload string
}

func (o *NcospatternsPostCreated) Error() string {
	return fmt.Sprintf("[POST /ncospatterns/][%d] ncospatternsPostCreated  %+v", 201, o.Payload)
}

func (o *NcospatternsPostCreated) GetPayload() string {
	return o.Payload
}

func (o *NcospatternsPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNcospatternsPostUnprocessableEntity creates a NcospatternsPostUnprocessableEntity with default headers values
func NewNcospatternsPostUnprocessableEntity() *NcospatternsPostUnprocessableEntity {
	return &NcospatternsPostUnprocessableEntity{}
}

/*NcospatternsPostUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type NcospatternsPostUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *NcospatternsPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /ncospatterns/][%d] ncospatternsPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *NcospatternsPostUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *NcospatternsPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}