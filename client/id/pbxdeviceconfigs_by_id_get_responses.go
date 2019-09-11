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

// PbxdeviceconfigsByIDGetReader is a Reader for the PbxdeviceconfigsByIDGet structure.
type PbxdeviceconfigsByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PbxdeviceconfigsByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPbxdeviceconfigsByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPbxdeviceconfigsByIDGetOK creates a PbxdeviceconfigsByIDGetOK with default headers values
func NewPbxdeviceconfigsByIDGetOK() *PbxdeviceconfigsByIDGetOK {
	return &PbxdeviceconfigsByIDGetOK{}
}

/*PbxdeviceconfigsByIDGetOK handles this case with default header values.

PbxdeviceconfigsByIDGetOK pbxdeviceconfigs by Id get o k
*/
type PbxdeviceconfigsByIDGetOK struct {
	Payload *models.PbxDeviceConfigs
}

func (o *PbxdeviceconfigsByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /pbxdeviceconfigs/{id}][%d] pbxdeviceconfigsByIdGetOK  %+v", 200, o.Payload)
}

func (o *PbxdeviceconfigsByIDGetOK) GetPayload() *models.PbxDeviceConfigs {
	return o.Payload
}

func (o *PbxdeviceconfigsByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PbxDeviceConfigs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
