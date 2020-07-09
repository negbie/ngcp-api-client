// Code generated by go-swagger; DO NOT EDIT.

package callrecordings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// GetCallRecordingitemsReader is a Reader for the GetCallRecordingitems structure.
type GetCallRecordingitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCallRecordingitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCallRecordingitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCallRecordingitemsOK creates a GetCallRecordingitemsOK with default headers values
func NewGetCallRecordingitemsOK() *GetCallRecordingitemsOK {
	return &GetCallRecordingitemsOK{}
}

/*GetCallRecordingitemsOK handles this case with default header values.

GetCallRecordingitemsOK get call recordingitems o k
*/
type GetCallRecordingitemsOK struct {
	Payload []*models.CallRecordings
}

func (o *GetCallRecordingitemsOK) Error() string {
	return fmt.Sprintf("[GET /callrecordings][%d] getCallRecordingitemsOK  %+v", 200, o.Payload)
}

func (o *GetCallRecordingitemsOK) GetPayload() []*models.CallRecordings {
	return o.Payload
}

func (o *GetCallRecordingitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
