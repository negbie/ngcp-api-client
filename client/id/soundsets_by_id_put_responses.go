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

// SoundsetsByIDPutReader is a Reader for the SoundsetsByIDPut structure.
type SoundsetsByIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SoundsetsByIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSoundsetsByIDPutNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSoundsetsByIDPutNoContent creates a SoundsetsByIDPutNoContent with default headers values
func NewSoundsetsByIDPutNoContent() *SoundsetsByIDPutNoContent {
	return &SoundsetsByIDPutNoContent{}
}

/*SoundsetsByIDPutNoContent handles this case with default header values.

SoundsetsByIDPutNoContent soundsets by Id put no content
*/
type SoundsetsByIDPutNoContent struct {
	Payload interface{}
}

func (o *SoundsetsByIDPutNoContent) Error() string {
	return fmt.Sprintf("[PUT /soundsets/{id}][%d] soundsetsByIdPutNoContent  %+v", 204, o.Payload)
}

func (o *SoundsetsByIDPutNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *SoundsetsByIDPutNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
