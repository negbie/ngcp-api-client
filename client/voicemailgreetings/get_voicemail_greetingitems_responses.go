// Code generated by go-swagger; DO NOT EDIT.

package voicemailgreetings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// GetVoicemailGreetingitemsReader is a Reader for the GetVoicemailGreetingitems structure.
type GetVoicemailGreetingitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVoicemailGreetingitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVoicemailGreetingitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVoicemailGreetingitemsOK creates a GetVoicemailGreetingitemsOK with default headers values
func NewGetVoicemailGreetingitemsOK() *GetVoicemailGreetingitemsOK {
	return &GetVoicemailGreetingitemsOK{}
}

/*GetVoicemailGreetingitemsOK handles this case with default header values.

GetVoicemailGreetingitemsOK get voicemail greetingitems o k
*/
type GetVoicemailGreetingitemsOK struct {
	Payload []*models.VoicemailGreetings1
}

func (o *GetVoicemailGreetingitemsOK) Error() string {
	return fmt.Sprintf("[GET /voicemailgreetings][%d] getVoicemailGreetingitemsOK  %+v", 200, o.Payload)
}

func (o *GetVoicemailGreetingitemsOK) GetPayload() []*models.VoicemailGreetings1 {
	return o.Payload
}

func (o *GetVoicemailGreetingitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
