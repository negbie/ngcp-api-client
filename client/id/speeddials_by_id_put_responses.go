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

// SpeeddialsByIDPutReader is a Reader for the SpeeddialsByIDPut structure.
type SpeeddialsByIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SpeeddialsByIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSpeeddialsByIDPutNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSpeeddialsByIDPutNoContent creates a SpeeddialsByIDPutNoContent with default headers values
func NewSpeeddialsByIDPutNoContent() *SpeeddialsByIDPutNoContent {
	return &SpeeddialsByIDPutNoContent{}
}

/*SpeeddialsByIDPutNoContent handles this case with default header values.

SpeeddialsByIDPutNoContent speeddials by Id put no content
*/
type SpeeddialsByIDPutNoContent struct {
	Payload interface{}
}

func (o *SpeeddialsByIDPutNoContent) Error() string {
	return fmt.Sprintf("[PUT /speeddials/{id}][%d] speeddialsByIdPutNoContent  %+v", 204, o.Payload)
}

func (o *SpeeddialsByIDPutNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *SpeeddialsByIDPutNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
