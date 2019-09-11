// Code generated by go-swagger; DO NOT EDIT.

package id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/negbie/ngcp-api/models"
)

// VoicemailsByIDPatchReader is a Reader for the VoicemailsByIDPatch structure.
type VoicemailsByIDPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VoicemailsByIDPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVoicemailsByIDPatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVoicemailsByIDPatchOK creates a VoicemailsByIDPatchOK with default headers values
func NewVoicemailsByIDPatchOK() *VoicemailsByIDPatchOK {
	return &VoicemailsByIDPatchOK{}
}

/*VoicemailsByIDPatchOK handles this case with default header values.

VoicemailsByIDPatchOK voicemails by Id patch o k
*/
type VoicemailsByIDPatchOK struct {
	Payload *models.Voicemails
}

func (o *VoicemailsByIDPatchOK) Error() string {
	return fmt.Sprintf("[PATCH /voicemails/{id}][%d] voicemailsByIdPatchOK  %+v", 200, o.Payload)
}

func (o *VoicemailsByIDPatchOK) GetPayload() *models.Voicemails {
	return o.Payload
}

func (o *VoicemailsByIDPatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Voicemails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
