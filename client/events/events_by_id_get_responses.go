// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/negbie/ngcp-api/models"
)

// EventsByIDGetReader is a Reader for the EventsByIDGet structure.
type EventsByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EventsByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEventsByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEventsByIDGetOK creates a EventsByIDGetOK with default headers values
func NewEventsByIDGetOK() *EventsByIDGetOK {
	return &EventsByIDGetOK{}
}

/*EventsByIDGetOK handles this case with default header values.

EventsByIDGetOK events by Id get o k
*/
type EventsByIDGetOK struct {
	Payload *models.Events
}

func (o *EventsByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /events/{id}][%d] eventsByIdGetOK  %+v", 200, o.Payload)
}

func (o *EventsByIDGetOK) GetPayload() *models.Events {
	return o.Payload
}

func (o *EventsByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Events)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
