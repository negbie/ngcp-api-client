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

// VoicemailsByIDDeleteReader is a Reader for the VoicemailsByIDDelete structure.
type VoicemailsByIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VoicemailsByIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVoicemailsByIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVoicemailsByIDDeleteNoContent creates a VoicemailsByIDDeleteNoContent with default headers values
func NewVoicemailsByIDDeleteNoContent() *VoicemailsByIDDeleteNoContent {
	return &VoicemailsByIDDeleteNoContent{}
}

/*VoicemailsByIDDeleteNoContent handles this case with default header values.

VoicemailsByIDDeleteNoContent voicemails by Id delete no content
*/
type VoicemailsByIDDeleteNoContent struct {
	Payload interface{}
}

func (o *VoicemailsByIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /voicemails/{id}][%d] voicemailsByIdDeleteNoContent  %+v", 204, o.Payload)
}

func (o *VoicemailsByIDDeleteNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *VoicemailsByIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
