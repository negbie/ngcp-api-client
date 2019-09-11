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

// PbxdeviceprofilesByIDPatchReader is a Reader for the PbxdeviceprofilesByIDPatch structure.
type PbxdeviceprofilesByIDPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PbxdeviceprofilesByIDPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPbxdeviceprofilesByIDPatchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPbxdeviceprofilesByIDPatchNoContent creates a PbxdeviceprofilesByIDPatchNoContent with default headers values
func NewPbxdeviceprofilesByIDPatchNoContent() *PbxdeviceprofilesByIDPatchNoContent {
	return &PbxdeviceprofilesByIDPatchNoContent{}
}

/*PbxdeviceprofilesByIDPatchNoContent handles this case with default header values.

PbxdeviceprofilesByIDPatchNoContent pbxdeviceprofiles by Id patch no content
*/
type PbxdeviceprofilesByIDPatchNoContent struct {
	Payload interface{}
}

func (o *PbxdeviceprofilesByIDPatchNoContent) Error() string {
	return fmt.Sprintf("[PATCH /pbxdeviceprofiles/{id}][%d] pbxdeviceprofilesByIdPatchNoContent  %+v", 204, o.Payload)
}

func (o *PbxdeviceprofilesByIDPatchNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PbxdeviceprofilesByIDPatchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
