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

// PeeringinboundrulesByIDPutReader is a Reader for the PeeringinboundrulesByIDPut structure.
type PeeringinboundrulesByIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PeeringinboundrulesByIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPeeringinboundrulesByIDPutNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPeeringinboundrulesByIDPutNoContent creates a PeeringinboundrulesByIDPutNoContent with default headers values
func NewPeeringinboundrulesByIDPutNoContent() *PeeringinboundrulesByIDPutNoContent {
	return &PeeringinboundrulesByIDPutNoContent{}
}

/*PeeringinboundrulesByIDPutNoContent handles this case with default header values.

PeeringinboundrulesByIDPutNoContent peeringinboundrules by Id put no content
*/
type PeeringinboundrulesByIDPutNoContent struct {
	Payload interface{}
}

func (o *PeeringinboundrulesByIDPutNoContent) Error() string {
	return fmt.Sprintf("[PUT /peeringinboundrules/{id}][%d] peeringinboundrulesByIdPutNoContent  %+v", 204, o.Payload)
}

func (o *PeeringinboundrulesByIDPutNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PeeringinboundrulesByIDPutNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
