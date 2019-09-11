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

// PbxdeviceprofilepreferencesByIDPutReader is a Reader for the PbxdeviceprofilepreferencesByIDPut structure.
type PbxdeviceprofilepreferencesByIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PbxdeviceprofilepreferencesByIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPbxdeviceprofilepreferencesByIDPutNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPbxdeviceprofilepreferencesByIDPutNoContent creates a PbxdeviceprofilepreferencesByIDPutNoContent with default headers values
func NewPbxdeviceprofilepreferencesByIDPutNoContent() *PbxdeviceprofilepreferencesByIDPutNoContent {
	return &PbxdeviceprofilepreferencesByIDPutNoContent{}
}

/*PbxdeviceprofilepreferencesByIDPutNoContent handles this case with default header values.

PbxdeviceprofilepreferencesByIDPutNoContent pbxdeviceprofilepreferences by Id put no content
*/
type PbxdeviceprofilepreferencesByIDPutNoContent struct {
	Payload interface{}
}

func (o *PbxdeviceprofilepreferencesByIDPutNoContent) Error() string {
	return fmt.Sprintf("[PUT /pbxdeviceprofilepreferences/{id}][%d] pbxdeviceprofilepreferencesByIdPutNoContent  %+v", 204, o.Payload)
}

func (o *PbxdeviceprofilepreferencesByIDPutNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PbxdeviceprofilepreferencesByIDPutNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
