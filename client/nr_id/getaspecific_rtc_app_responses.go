// Code generated by go-swagger; DO NOT EDIT.

package nr_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// GetaspecificRtcAppReader is a Reader for the GetaspecificRtcApp structure.
type GetaspecificRtcAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificRtcAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificRtcAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificRtcAppOK creates a GetaspecificRtcAppOK with default headers values
func NewGetaspecificRtcAppOK() *GetaspecificRtcAppOK {
	return &GetaspecificRtcAppOK{}
}

/*GetaspecificRtcAppOK handles this case with default header values.

GetaspecificRtcAppOK getaspecific rtc app o k
*/
type GetaspecificRtcAppOK struct {
	Payload *models.RtcApps
}

func (o *GetaspecificRtcAppOK) Error() string {
	return fmt.Sprintf("[GET /rtcapps/{id}][%d] getaspecificRtcAppOK  %+v", 200, o.Payload)
}

func (o *GetaspecificRtcAppOK) GetPayload() *models.RtcApps {
	return o.Payload
}

func (o *GetaspecificRtcAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RtcApps)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}