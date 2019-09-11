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

// RtcappsByIDGetReader is a Reader for the RtcappsByIDGet structure.
type RtcappsByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RtcappsByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRtcappsByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRtcappsByIDGetOK creates a RtcappsByIDGetOK with default headers values
func NewRtcappsByIDGetOK() *RtcappsByIDGetOK {
	return &RtcappsByIDGetOK{}
}

/*RtcappsByIDGetOK handles this case with default header values.

RtcappsByIDGetOK rtcapps by Id get o k
*/
type RtcappsByIDGetOK struct {
	Payload *models.RtcApps
}

func (o *RtcappsByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /rtcapps/{id}][%d] rtcappsByIdGetOK  %+v", 200, o.Payload)
}

func (o *RtcappsByIDGetOK) GetPayload() *models.RtcApps {
	return o.Payload
}

func (o *RtcappsByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RtcApps)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
