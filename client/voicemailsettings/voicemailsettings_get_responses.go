// Code generated by go-swagger; DO NOT EDIT.

package voicemailsettings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/negbie/ngcp-api/models"
)

// VoicemailsettingsGetReader is a Reader for the VoicemailsettingsGet structure.
type VoicemailsettingsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VoicemailsettingsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVoicemailsettingsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVoicemailsettingsGetOK creates a VoicemailsettingsGetOK with default headers values
func NewVoicemailsettingsGetOK() *VoicemailsettingsGetOK {
	return &VoicemailsettingsGetOK{}
}

/*VoicemailsettingsGetOK handles this case with default header values.

VoicemailsettingsGetOK voicemailsettings get o k
*/
type VoicemailsettingsGetOK struct {
	Payload []*models.VoicemailSettings1
}

func (o *VoicemailsettingsGetOK) Error() string {
	return fmt.Sprintf("[GET /voicemailsettings][%d] voicemailsettingsGetOK  %+v", 200, o.Payload)
}

func (o *VoicemailsettingsGetOK) GetPayload() []*models.VoicemailSettings1 {
	return o.Payload
}

func (o *VoicemailsettingsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
