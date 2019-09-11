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

// RemindersPostReader is a Reader for the RemindersPost structure.
type RemindersPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemindersPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRemindersPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewRemindersPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemindersPostCreated creates a RemindersPostCreated with default headers values
func NewRemindersPostCreated() *RemindersPostCreated {
	return &RemindersPostCreated{
		Location: "\"<string>\"",
	}
}

/*RemindersPostCreated handles this case with default header values.

RemindersPostCreated reminders post created
*/
type RemindersPostCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty24
}

func (o *RemindersPostCreated) Error() string {
	return fmt.Sprintf("[POST /reminders/][%d] remindersPostCreated  %+v", 201, o.Payload)
}

func (o *RemindersPostCreated) GetPayload() []*models.Thenewlycreateditemorempty24 {
	return o.Payload
}

func (o *RemindersPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemindersPostUnprocessableEntity creates a RemindersPostUnprocessableEntity with default headers values
func NewRemindersPostUnprocessableEntity() *RemindersPostUnprocessableEntity {
	return &RemindersPostUnprocessableEntity{}
}

/*RemindersPostUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type RemindersPostUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *RemindersPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /reminders/][%d] remindersPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *RemindersPostUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *RemindersPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
