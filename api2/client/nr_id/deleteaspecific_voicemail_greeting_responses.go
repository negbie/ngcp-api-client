// Code generated by go-swagger; DO NOT EDIT.

package nr_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteaspecificVoicemailGreetingReader is a Reader for the DeleteaspecificVoicemailGreeting structure.
type DeleteaspecificVoicemailGreetingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteaspecificVoicemailGreetingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteaspecificVoicemailGreetingNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteaspecificVoicemailGreetingNoContent creates a DeleteaspecificVoicemailGreetingNoContent with default headers values
func NewDeleteaspecificVoicemailGreetingNoContent() *DeleteaspecificVoicemailGreetingNoContent {
	return &DeleteaspecificVoicemailGreetingNoContent{}
}

/*DeleteaspecificVoicemailGreetingNoContent handles this case with default header values.

DeleteaspecificVoicemailGreetingNoContent deleteaspecific voicemail greeting no content
*/
type DeleteaspecificVoicemailGreetingNoContent struct {
	Payload interface{}
}

func (o *DeleteaspecificVoicemailGreetingNoContent) Error() string {
	return fmt.Sprintf("[DELETE /voicemailgreetings/{id}][%d] deleteaspecificVoicemailGreetingNoContent  %+v", 204, o.Payload)
}

func (o *DeleteaspecificVoicemailGreetingNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteaspecificVoicemailGreetingNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}