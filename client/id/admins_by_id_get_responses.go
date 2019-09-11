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

// AdminsByIDGetReader is a Reader for the AdminsByIDGet structure.
type AdminsByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminsByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminsByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAdminsByIDGetOK creates a AdminsByIDGetOK with default headers values
func NewAdminsByIDGetOK() *AdminsByIDGetOK {
	return &AdminsByIDGetOK{}
}

/*AdminsByIDGetOK handles this case with default header values.

AdminsByIDGetOK admins by Id get o k
*/
type AdminsByIDGetOK struct {
	Payload *models.Admins
}

func (o *AdminsByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /admins/{id}][%d] adminsByIdGetOK  %+v", 200, o.Payload)
}

func (o *AdminsByIDGetOK) GetPayload() *models.Admins {
	return o.Payload
}

func (o *AdminsByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Admins)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}