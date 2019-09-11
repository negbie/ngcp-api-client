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

// PeeringinboundrulesByIDGetReader is a Reader for the PeeringinboundrulesByIDGet structure.
type PeeringinboundrulesByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PeeringinboundrulesByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPeeringinboundrulesByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPeeringinboundrulesByIDGetOK creates a PeeringinboundrulesByIDGetOK with default headers values
func NewPeeringinboundrulesByIDGetOK() *PeeringinboundrulesByIDGetOK {
	return &PeeringinboundrulesByIDGetOK{}
}

/*PeeringinboundrulesByIDGetOK handles this case with default header values.

PeeringinboundrulesByIDGetOK peeringinboundrules by Id get o k
*/
type PeeringinboundrulesByIDGetOK struct {
	Payload string
}

func (o *PeeringinboundrulesByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /peeringinboundrules/{id}][%d] peeringinboundrulesByIdGetOK  %+v", 200, o.Payload)
}

func (o *PeeringinboundrulesByIDGetOK) GetPayload() string {
	return o.Payload
}

func (o *PeeringinboundrulesByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
