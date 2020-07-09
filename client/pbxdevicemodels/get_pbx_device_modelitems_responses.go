// Code generated by go-swagger; DO NOT EDIT.

package pbxdevicemodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// GetPbxDeviceModelitemsReader is a Reader for the GetPbxDeviceModelitems structure.
type GetPbxDeviceModelitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPbxDeviceModelitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPbxDeviceModelitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPbxDeviceModelitemsOK creates a GetPbxDeviceModelitemsOK with default headers values
func NewGetPbxDeviceModelitemsOK() *GetPbxDeviceModelitemsOK {
	return &GetPbxDeviceModelitemsOK{}
}

/*GetPbxDeviceModelitemsOK handles this case with default header values.

GetPbxDeviceModelitemsOK get pbx device modelitems o k
*/
type GetPbxDeviceModelitemsOK struct {
	Payload []*models.PbxDeviceModels1
}

func (o *GetPbxDeviceModelitemsOK) Error() string {
	return fmt.Sprintf("[GET /pbxdevicemodels][%d] getPbxDeviceModelitemsOK  %+v", 200, o.Payload)
}

func (o *GetPbxDeviceModelitemsOK) GetPayload() []*models.PbxDeviceModels1 {
	return o.Payload
}

func (o *GetPbxDeviceModelitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
