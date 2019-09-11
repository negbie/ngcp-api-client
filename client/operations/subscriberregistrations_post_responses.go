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

// SubscriberregistrationsPostReader is a Reader for the SubscriberregistrationsPost structure.
type SubscriberregistrationsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscriberregistrationsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSubscriberregistrationsPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewSubscriberregistrationsPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSubscriberregistrationsPostCreated creates a SubscriberregistrationsPostCreated with default headers values
func NewSubscriberregistrationsPostCreated() *SubscriberregistrationsPostCreated {
	return &SubscriberregistrationsPostCreated{
		Location: "\"<string>\"",
	}
}

/*SubscriberregistrationsPostCreated handles this case with default header values.

SubscriberregistrationsPostCreated subscriberregistrations post created
*/
type SubscriberregistrationsPostCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty13
}

func (o *SubscriberregistrationsPostCreated) Error() string {
	return fmt.Sprintf("[POST /subscriberregistrations/][%d] subscriberregistrationsPostCreated  %+v", 201, o.Payload)
}

func (o *SubscriberregistrationsPostCreated) GetPayload() []*models.Thenewlycreateditemorempty13 {
	return o.Payload
}

func (o *SubscriberregistrationsPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriberregistrationsPostUnprocessableEntity creates a SubscriberregistrationsPostUnprocessableEntity with default headers values
func NewSubscriberregistrationsPostUnprocessableEntity() *SubscriberregistrationsPostUnprocessableEntity {
	return &SubscriberregistrationsPostUnprocessableEntity{}
}

/*SubscriberregistrationsPostUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type SubscriberregistrationsPostUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *SubscriberregistrationsPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /subscriberregistrations/][%d] subscriberregistrationsPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *SubscriberregistrationsPostUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *SubscriberregistrationsPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
