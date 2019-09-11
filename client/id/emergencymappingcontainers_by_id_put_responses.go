// Code generated by go-swagger; DO NOT EDIT.

package id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// EmergencymappingcontainersByIDPutReader is a Reader for the EmergencymappingcontainersByIDPut structure.
type EmergencymappingcontainersByIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmergencymappingcontainersByIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewEmergencymappingcontainersByIDPutNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEmergencymappingcontainersByIDPutNoContent creates a EmergencymappingcontainersByIDPutNoContent with default headers values
func NewEmergencymappingcontainersByIDPutNoContent() *EmergencymappingcontainersByIDPutNoContent {
	return &EmergencymappingcontainersByIDPutNoContent{}
}

/*EmergencymappingcontainersByIDPutNoContent handles this case with default header values.

EmergencymappingcontainersByIDPutNoContent emergencymappingcontainers by Id put no content
*/
type EmergencymappingcontainersByIDPutNoContent struct {
	Payload interface{}
}

func (o *EmergencymappingcontainersByIDPutNoContent) Error() string {
	return fmt.Sprintf("[PUT /emergencymappingcontainers/{id}][%d] emergencymappingcontainersByIdPutNoContent  %+v", 204, o.Payload)
}

func (o *EmergencymappingcontainersByIDPutNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *EmergencymappingcontainersByIDPutNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
