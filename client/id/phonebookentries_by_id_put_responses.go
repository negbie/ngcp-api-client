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

// PhonebookentriesByIDPutReader is a Reader for the PhonebookentriesByIDPut structure.
type PhonebookentriesByIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PhonebookentriesByIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPhonebookentriesByIDPutNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPhonebookentriesByIDPutNoContent creates a PhonebookentriesByIDPutNoContent with default headers values
func NewPhonebookentriesByIDPutNoContent() *PhonebookentriesByIDPutNoContent {
	return &PhonebookentriesByIDPutNoContent{}
}

/*PhonebookentriesByIDPutNoContent handles this case with default header values.

PhonebookentriesByIDPutNoContent phonebookentries by Id put no content
*/
type PhonebookentriesByIDPutNoContent struct {
	Payload interface{}
}

func (o *PhonebookentriesByIDPutNoContent) Error() string {
	return fmt.Sprintf("[PUT /phonebookentries/{id}][%d] phonebookentriesByIdPutNoContent  %+v", 204, o.Payload)
}

func (o *PhonebookentriesByIDPutNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PhonebookentriesByIDPutNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}