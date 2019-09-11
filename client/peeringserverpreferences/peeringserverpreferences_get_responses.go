// Code generated by go-swagger; DO NOT EDIT.

package peeringserverpreferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PeeringserverpreferencesGetReader is a Reader for the PeeringserverpreferencesGet structure.
type PeeringserverpreferencesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PeeringserverpreferencesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPeeringserverpreferencesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPeeringserverpreferencesGetOK creates a PeeringserverpreferencesGetOK with default headers values
func NewPeeringserverpreferencesGetOK() *PeeringserverpreferencesGetOK {
	return &PeeringserverpreferencesGetOK{}
}

/*PeeringserverpreferencesGetOK handles this case with default header values.

PeeringserverpreferencesGetOK peeringserverpreferences get o k
*/
type PeeringserverpreferencesGetOK struct {
	Payload []string
}

func (o *PeeringserverpreferencesGetOK) Error() string {
	return fmt.Sprintf("[GET /peeringserverpreferences][%d] peeringserverpreferencesGetOK  %+v", 200, o.Payload)
}

func (o *PeeringserverpreferencesGetOK) GetPayload() []string {
	return o.Payload
}

func (o *PeeringserverpreferencesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}