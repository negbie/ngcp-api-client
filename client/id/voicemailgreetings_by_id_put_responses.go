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

// VoicemailgreetingsByIDPutReader is a Reader for the VoicemailgreetingsByIDPut structure.
type VoicemailgreetingsByIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VoicemailgreetingsByIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVoicemailgreetingsByIDPutNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVoicemailgreetingsByIDPutNoContent creates a VoicemailgreetingsByIDPutNoContent with default headers values
func NewVoicemailgreetingsByIDPutNoContent() *VoicemailgreetingsByIDPutNoContent {
	return &VoicemailgreetingsByIDPutNoContent{}
}

/*VoicemailgreetingsByIDPutNoContent handles this case with default header values.

VoicemailgreetingsByIDPutNoContent voicemailgreetings by Id put no content
*/
type VoicemailgreetingsByIDPutNoContent struct {
	Payload interface{}
}

func (o *VoicemailgreetingsByIDPutNoContent) Error() string {
	return fmt.Sprintf("[PUT /voicemailgreetings/{id}][%d] voicemailgreetingsByIdPutNoContent  %+v", 204, o.Payload)
}

func (o *VoicemailgreetingsByIDPutNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *VoicemailgreetingsByIDPutNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
