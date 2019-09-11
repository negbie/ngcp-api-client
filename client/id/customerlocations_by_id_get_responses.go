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

// CustomerlocationsByIDGetReader is a Reader for the CustomerlocationsByIDGet structure.
type CustomerlocationsByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerlocationsByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerlocationsByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomerlocationsByIDGetOK creates a CustomerlocationsByIDGetOK with default headers values
func NewCustomerlocationsByIDGetOK() *CustomerlocationsByIDGetOK {
	return &CustomerlocationsByIDGetOK{}
}

/*CustomerlocationsByIDGetOK handles this case with default header values.

CustomerlocationsByIDGetOK customerlocations by Id get o k
*/
type CustomerlocationsByIDGetOK struct {
	Payload *models.CustomerLocations
}

func (o *CustomerlocationsByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /customerlocations/{id}][%d] customerlocationsByIdGetOK  %+v", 200, o.Payload)
}

func (o *CustomerlocationsByIDGetOK) GetPayload() *models.CustomerLocations {
	return o.Payload
}

func (o *CustomerlocationsByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerLocations)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
