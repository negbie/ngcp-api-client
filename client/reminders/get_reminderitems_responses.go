// Code generated by go-swagger; DO NOT EDIT.

package reminders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// GetReminderitemsReader is a Reader for the GetReminderitems structure.
type GetReminderitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReminderitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReminderitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReminderitemsOK creates a GetReminderitemsOK with default headers values
func NewGetReminderitemsOK() *GetReminderitemsOK {
	return &GetReminderitemsOK{}
}

/*GetReminderitemsOK handles this case with default header values.

GetReminderitemsOK get reminderitems o k
*/
type GetReminderitemsOK struct {
	Payload []*models.Reminders1
}

func (o *GetReminderitemsOK) Error() string {
	return fmt.Sprintf("[GET /reminders][%d] getReminderitemsOK  %+v", 200, o.Payload)
}

func (o *GetReminderitemsOK) GetPayload() []*models.Reminders1 {
	return o.Payload
}

func (o *GetReminderitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
