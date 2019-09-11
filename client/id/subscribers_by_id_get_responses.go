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

// SubscribersByIDGetReader is a Reader for the SubscribersByIDGet structure.
type SubscribersByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscribersByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubscribersByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSubscribersByIDGetOK creates a SubscribersByIDGetOK with default headers values
func NewSubscribersByIDGetOK() *SubscribersByIDGetOK {
	return &SubscribersByIDGetOK{}
}

/*SubscribersByIDGetOK handles this case with default header values.

SubscribersByIDGetOK subscribers by Id get o k
*/
type SubscribersByIDGetOK struct {
	Payload *models.Subscribers
}

func (o *SubscribersByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /subscribers/{id}][%d] subscribersByIdGetOK  %+v", 200, o.Payload)
}

func (o *SubscribersByIDGetOK) GetPayload() *models.Subscribers {
	return o.Payload
}

func (o *SubscribersByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Subscribers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
