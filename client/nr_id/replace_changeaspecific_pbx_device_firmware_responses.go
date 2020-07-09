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

// ReplaceChangeaspecificPbxDeviceFirmwareReader is a Reader for the ReplaceChangeaspecificPbxDeviceFirmware structure.
type ReplaceChangeaspecificPbxDeviceFirmwareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceChangeaspecificPbxDeviceFirmwareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceChangeaspecificPbxDeviceFirmwareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceChangeaspecificPbxDeviceFirmwareOK creates a ReplaceChangeaspecificPbxDeviceFirmwareOK with default headers values
func NewReplaceChangeaspecificPbxDeviceFirmwareOK() *ReplaceChangeaspecificPbxDeviceFirmwareOK {
	return &ReplaceChangeaspecificPbxDeviceFirmwareOK{}
}

/*ReplaceChangeaspecificPbxDeviceFirmwareOK handles this case with default header values.

ReplaceChangeaspecificPbxDeviceFirmwareOK replace changeaspecific pbx device firmware o k
*/
type ReplaceChangeaspecificPbxDeviceFirmwareOK struct {
	Payload *models.PbxDeviceFirmwares
}

func (o *ReplaceChangeaspecificPbxDeviceFirmwareOK) Error() string {
	return fmt.Sprintf("[PUT /pbxdevicefirmwares/{id}][%d] replaceChangeaspecificPbxDeviceFirmwareOK  %+v", 200, o.Payload)
}

func (o *ReplaceChangeaspecificPbxDeviceFirmwareOK) GetPayload() *models.PbxDeviceFirmwares {
	return o.Payload
}

func (o *ReplaceChangeaspecificPbxDeviceFirmwareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PbxDeviceFirmwares)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
