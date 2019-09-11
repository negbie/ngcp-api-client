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

// CustomerlocationsByIDPutReader is a Reader for the CustomerlocationsByIDPut structure.
type CustomerlocationsByIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerlocationsByIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCustomerlocationsByIDPutNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomerlocationsByIDPutNoContent creates a CustomerlocationsByIDPutNoContent with default headers values
func NewCustomerlocationsByIDPutNoContent() *CustomerlocationsByIDPutNoContent {
	return &CustomerlocationsByIDPutNoContent{}
}

/*CustomerlocationsByIDPutNoContent handles this case with default header values.

CustomerlocationsByIDPutNoContent customerlocations by Id put no content
*/
type CustomerlocationsByIDPutNoContent struct {
	Payload interface{}
}

func (o *CustomerlocationsByIDPutNoContent) Error() string {
	return fmt.Sprintf("[PUT /customerlocations/{id}][%d] customerlocationsByIdPutNoContent  %+v", 204, o.Payload)
}

func (o *CustomerlocationsByIDPutNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerlocationsByIDPutNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
