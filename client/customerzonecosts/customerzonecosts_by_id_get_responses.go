// Code generated by go-swagger; DO NOT EDIT.

package customerzonecosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CustomerzonecostsByIDGetReader is a Reader for the CustomerzonecostsByIDGet structure.
type CustomerzonecostsByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerzonecostsByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerzonecostsByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomerzonecostsByIDGetOK creates a CustomerzonecostsByIDGetOK with default headers values
func NewCustomerzonecostsByIDGetOK() *CustomerzonecostsByIDGetOK {
	return &CustomerzonecostsByIDGetOK{}
}

/*CustomerzonecostsByIDGetOK handles this case with default header values.

CustomerzonecostsByIDGetOK customerzonecosts by Id get o k
*/
type CustomerzonecostsByIDGetOK struct {
	Payload string
}

func (o *CustomerzonecostsByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /customerzonecosts/{id}][%d] customerzonecostsByIdGetOK  %+v", 200, o.Payload)
}

func (o *CustomerzonecostsByIDGetOK) GetPayload() string {
	return o.Payload
}

func (o *CustomerzonecostsByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}