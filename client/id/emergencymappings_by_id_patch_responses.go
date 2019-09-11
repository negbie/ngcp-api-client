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

// EmergencymappingsByIDPatchReader is a Reader for the EmergencymappingsByIDPatch structure.
type EmergencymappingsByIDPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmergencymappingsByIDPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmergencymappingsByIDPatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEmergencymappingsByIDPatchOK creates a EmergencymappingsByIDPatchOK with default headers values
func NewEmergencymappingsByIDPatchOK() *EmergencymappingsByIDPatchOK {
	return &EmergencymappingsByIDPatchOK{}
}

/*EmergencymappingsByIDPatchOK handles this case with default header values.

EmergencymappingsByIDPatchOK emergencymappings by Id patch o k
*/
type EmergencymappingsByIDPatchOK struct {
	Payload *models.EmergencyMappings
}

func (o *EmergencymappingsByIDPatchOK) Error() string {
	return fmt.Sprintf("[PATCH /emergencymappings/{id}][%d] emergencymappingsByIdPatchOK  %+v", 200, o.Payload)
}

func (o *EmergencymappingsByIDPatchOK) GetPayload() *models.EmergencyMappings {
	return o.Payload
}

func (o *EmergencymappingsByIDPatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmergencyMappings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
