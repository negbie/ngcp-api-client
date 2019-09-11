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

// ProfilepreferencesByIDPutReader is a Reader for the ProfilepreferencesByIDPut structure.
type ProfilepreferencesByIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProfilepreferencesByIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewProfilepreferencesByIDPutNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProfilepreferencesByIDPutNoContent creates a ProfilepreferencesByIDPutNoContent with default headers values
func NewProfilepreferencesByIDPutNoContent() *ProfilepreferencesByIDPutNoContent {
	return &ProfilepreferencesByIDPutNoContent{}
}

/*ProfilepreferencesByIDPutNoContent handles this case with default header values.

ProfilepreferencesByIDPutNoContent profilepreferences by Id put no content
*/
type ProfilepreferencesByIDPutNoContent struct {
	Payload interface{}
}

func (o *ProfilepreferencesByIDPutNoContent) Error() string {
	return fmt.Sprintf("[PUT /profilepreferences/{id}][%d] profilepreferencesByIdPutNoContent  %+v", 204, o.Payload)
}

func (o *ProfilepreferencesByIDPutNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ProfilepreferencesByIDPutNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
