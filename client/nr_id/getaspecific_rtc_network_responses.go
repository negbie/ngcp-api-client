// Code generated by go-swagger; DO NOT EDIT.

package nr_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// GetaspecificRtcNetworkReader is a Reader for the GetaspecificRtcNetwork structure.
type GetaspecificRtcNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificRtcNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificRtcNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificRtcNetworkOK creates a GetaspecificRtcNetworkOK with default headers values
func NewGetaspecificRtcNetworkOK() *GetaspecificRtcNetworkOK {
	return &GetaspecificRtcNetworkOK{}
}

/*GetaspecificRtcNetworkOK handles this case with default header values.

GetaspecificRtcNetworkOK getaspecific rtc network o k
*/
type GetaspecificRtcNetworkOK struct {
	Payload *models.RtcNetworks
}

func (o *GetaspecificRtcNetworkOK) Error() string {
	return fmt.Sprintf("[GET /rtcnetworks/{id}][%d] getaspecificRtcNetworkOK  %+v", 200, o.Payload)
}

func (o *GetaspecificRtcNetworkOK) GetPayload() *models.RtcNetworks {
	return o.Payload
}

func (o *GetaspecificRtcNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RtcNetworks)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
