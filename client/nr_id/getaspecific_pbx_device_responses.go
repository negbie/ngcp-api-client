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

// GetaspecificPbxDeviceReader is a Reader for the GetaspecificPbxDevice structure.
type GetaspecificPbxDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificPbxDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificPbxDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificPbxDeviceOK creates a GetaspecificPbxDeviceOK with default headers values
func NewGetaspecificPbxDeviceOK() *GetaspecificPbxDeviceOK {
	return &GetaspecificPbxDeviceOK{}
}

/*GetaspecificPbxDeviceOK handles this case with default header values.

GetaspecificPbxDeviceOK getaspecific pbx device o k
*/
type GetaspecificPbxDeviceOK struct {
	Payload *models.PbxDevices
}

func (o *GetaspecificPbxDeviceOK) Error() string {
	return fmt.Sprintf("[GET /pbxdevices/{id}][%d] getaspecificPbxDeviceOK  %+v", 200, o.Payload)
}

func (o *GetaspecificPbxDeviceOK) GetPayload() *models.PbxDevices {
	return o.Payload
}

func (o *GetaspecificPbxDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PbxDevices)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
