// Code generated by go-swagger; DO NOT EDIT.

package id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/negbie/ngcp-api/models"
)

// RemindersByIDGetReader is a Reader for the RemindersByIDGet structure.
type RemindersByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemindersByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemindersByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemindersByIDGetOK creates a RemindersByIDGetOK with default headers values
func NewRemindersByIDGetOK() *RemindersByIDGetOK {
	return &RemindersByIDGetOK{}
}

/*RemindersByIDGetOK handles this case with default header values.

RemindersByIDGetOK reminders by Id get o k
*/
type RemindersByIDGetOK struct {
	Payload *models.Reminders
}

func (o *RemindersByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /reminders/{id}][%d] remindersByIdGetOK  %+v", 200, o.Payload)
}

func (o *RemindersByIDGetOK) GetPayload() *models.Reminders {
	return o.Payload
}

func (o *RemindersByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Reminders)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}