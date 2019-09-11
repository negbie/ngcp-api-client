// Code generated by go-swagger; DO NOT EDIT.

package callrecordings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/negbie/ngcp-api/models"
)

// CallrecordingsGetReader is a Reader for the CallrecordingsGet structure.
type CallrecordingsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CallrecordingsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCallrecordingsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCallrecordingsGetOK creates a CallrecordingsGetOK with default headers values
func NewCallrecordingsGetOK() *CallrecordingsGetOK {
	return &CallrecordingsGetOK{}
}

/*CallrecordingsGetOK handles this case with default header values.

CallrecordingsGetOK callrecordings get o k
*/
type CallrecordingsGetOK struct {
	Payload []*models.CallRecordings
}

func (o *CallrecordingsGetOK) Error() string {
	return fmt.Sprintf("[GET /callrecordings][%d] callrecordingsGetOK  %+v", 200, o.Payload)
}

func (o *CallrecordingsGetOK) GetPayload() []*models.CallRecordings {
	return o.Payload
}

func (o *CallrecordingsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
