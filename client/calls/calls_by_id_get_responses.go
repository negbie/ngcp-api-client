// Code generated by go-swagger; DO NOT EDIT.

package calls

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/negbie/ngcp-api/models"
)

// CallsByIDGetReader is a Reader for the CallsByIDGet structure.
type CallsByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CallsByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCallsByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCallsByIDGetOK creates a CallsByIDGetOK with default headers values
func NewCallsByIDGetOK() *CallsByIDGetOK {
	return &CallsByIDGetOK{}
}

/*CallsByIDGetOK handles this case with default header values.

CallsByIDGetOK calls by Id get o k
*/
type CallsByIDGetOK struct {
	Payload *models.Calls
}

func (o *CallsByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /calls/{id}][%d] callsByIdGetOK  %+v", 200, o.Payload)
}

func (o *CallsByIDGetOK) GetPayload() *models.Calls {
	return o.Payload
}

func (o *CallsByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Calls)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
