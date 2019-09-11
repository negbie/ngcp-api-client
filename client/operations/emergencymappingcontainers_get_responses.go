// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/negbie/ngcp-api/models"
)

// EmergencymappingcontainersGetReader is a Reader for the EmergencymappingcontainersGet structure.
type EmergencymappingcontainersGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmergencymappingcontainersGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmergencymappingcontainersGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEmergencymappingcontainersGetOK creates a EmergencymappingcontainersGetOK with default headers values
func NewEmergencymappingcontainersGetOK() *EmergencymappingcontainersGetOK {
	return &EmergencymappingcontainersGetOK{}
}

/*EmergencymappingcontainersGetOK handles this case with default header values.

EmergencymappingcontainersGetOK emergencymappingcontainers get o k
*/
type EmergencymappingcontainersGetOK struct {
	Payload []*models.EmergencyMappingContainer
}

func (o *EmergencymappingcontainersGetOK) Error() string {
	return fmt.Sprintf("[GET /emergencymappingcontainers][%d] emergencymappingcontainersGetOK  %+v", 200, o.Payload)
}

func (o *EmergencymappingcontainersGetOK) GetPayload() []*models.EmergencyMappingContainer {
	return o.Payload
}

func (o *EmergencymappingcontainersGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}