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

// SystemcontactsByIDGetReader is a Reader for the SystemcontactsByIDGet structure.
type SystemcontactsByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemcontactsByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemcontactsByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSystemcontactsByIDGetOK creates a SystemcontactsByIDGetOK with default headers values
func NewSystemcontactsByIDGetOK() *SystemcontactsByIDGetOK {
	return &SystemcontactsByIDGetOK{}
}

/*SystemcontactsByIDGetOK handles this case with default header values.

SystemcontactsByIDGetOK systemcontacts by Id get o k
*/
type SystemcontactsByIDGetOK struct {
	Payload *models.SystemContacts
}

func (o *SystemcontactsByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /systemcontacts/{id}][%d] systemcontactsByIdGetOK  %+v", 200, o.Payload)
}

func (o *SystemcontactsByIDGetOK) GetPayload() *models.SystemContacts {
	return o.Payload
}

func (o *SystemcontactsByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemContacts)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
