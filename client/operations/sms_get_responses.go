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

// SmsGetReader is a Reader for the SmsGet structure.
type SmsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SmsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSmsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSmsGetOK creates a SmsGetOK with default headers values
func NewSmsGetOK() *SmsGetOK {
	return &SmsGetOK{}
}

/*SmsGetOK handles this case with default header values.

SmsGetOK sms get o k
*/
type SmsGetOK struct {
	Payload []*models.SM
}

func (o *SmsGetOK) Error() string {
	return fmt.Sprintf("[GET /sms][%d] smsGetOK  %+v", 200, o.Payload)
}

func (o *SmsGetOK) GetPayload() []*models.SM {
	return o.Payload
}

func (o *SmsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
