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

// ProfilepackagesByIDDeleteReader is a Reader for the ProfilepackagesByIDDelete structure.
type ProfilepackagesByIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProfilepackagesByIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewProfilepackagesByIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProfilepackagesByIDDeleteNoContent creates a ProfilepackagesByIDDeleteNoContent with default headers values
func NewProfilepackagesByIDDeleteNoContent() *ProfilepackagesByIDDeleteNoContent {
	return &ProfilepackagesByIDDeleteNoContent{}
}

/*ProfilepackagesByIDDeleteNoContent handles this case with default header values.

ProfilepackagesByIDDeleteNoContent profilepackages by Id delete no content
*/
type ProfilepackagesByIDDeleteNoContent struct {
	Payload interface{}
}

func (o *ProfilepackagesByIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /profilepackages/{id}][%d] profilepackagesByIdDeleteNoContent  %+v", 204, o.Payload)
}

func (o *ProfilepackagesByIDDeleteNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ProfilepackagesByIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
