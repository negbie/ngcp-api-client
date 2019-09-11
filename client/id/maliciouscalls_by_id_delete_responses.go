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

// MaliciouscallsByIDDeleteReader is a Reader for the MaliciouscallsByIDDelete structure.
type MaliciouscallsByIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MaliciouscallsByIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewMaliciouscallsByIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMaliciouscallsByIDDeleteNoContent creates a MaliciouscallsByIDDeleteNoContent with default headers values
func NewMaliciouscallsByIDDeleteNoContent() *MaliciouscallsByIDDeleteNoContent {
	return &MaliciouscallsByIDDeleteNoContent{}
}

/*MaliciouscallsByIDDeleteNoContent handles this case with default header values.

MaliciouscallsByIDDeleteNoContent maliciouscalls by Id delete no content
*/
type MaliciouscallsByIDDeleteNoContent struct {
	Payload interface{}
}

func (o *MaliciouscallsByIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /maliciouscalls/{id}][%d] maliciouscallsByIdDeleteNoContent  %+v", 204, o.Payload)
}

func (o *MaliciouscallsByIDDeleteNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *MaliciouscallsByIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}