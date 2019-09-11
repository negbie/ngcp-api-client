// Code generated by go-swagger; DO NOT EDIT.

package id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PbxfielddevicepreferencesByIDGetReader is a Reader for the PbxfielddevicepreferencesByIDGet structure.
type PbxfielddevicepreferencesByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PbxfielddevicepreferencesByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPbxfielddevicepreferencesByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPbxfielddevicepreferencesByIDGetOK creates a PbxfielddevicepreferencesByIDGetOK with default headers values
func NewPbxfielddevicepreferencesByIDGetOK() *PbxfielddevicepreferencesByIDGetOK {
	return &PbxfielddevicepreferencesByIDGetOK{}
}

/*PbxfielddevicepreferencesByIDGetOK handles this case with default header values.

PbxfielddevicepreferencesByIDGetOK pbxfielddevicepreferences by Id get o k
*/
type PbxfielddevicepreferencesByIDGetOK struct {
	Payload string
}

func (o *PbxfielddevicepreferencesByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /pbxfielddevicepreferences/{id}][%d] pbxfielddevicepreferencesByIdGetOK  %+v", 200, o.Payload)
}

func (o *PbxfielddevicepreferencesByIDGetOK) GetPayload() string {
	return o.Payload
}

func (o *PbxfielddevicepreferencesByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}