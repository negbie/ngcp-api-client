// Code generated by go-swagger; DO NOT EDIT.

package id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RemindersByIDDeleteReader is a Reader for the RemindersByIDDelete structure.
type RemindersByIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemindersByIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRemindersByIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemindersByIDDeleteNoContent creates a RemindersByIDDeleteNoContent with default headers values
func NewRemindersByIDDeleteNoContent() *RemindersByIDDeleteNoContent {
	return &RemindersByIDDeleteNoContent{}
}

/*RemindersByIDDeleteNoContent handles this case with default header values.

RemindersByIDDeleteNoContent reminders by Id delete no content
*/
type RemindersByIDDeleteNoContent struct {
	Payload interface{}
}

func (o *RemindersByIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /reminders/{id}][%d] remindersByIdDeleteNoContent  %+v", 204, o.Payload)
}

func (o *RemindersByIDDeleteNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *RemindersByIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
