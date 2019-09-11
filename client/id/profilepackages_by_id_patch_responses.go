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

// ProfilepackagesByIDPatchReader is a Reader for the ProfilepackagesByIDPatch structure.
type ProfilepackagesByIDPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProfilepackagesByIDPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewProfilepackagesByIDPatchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProfilepackagesByIDPatchNoContent creates a ProfilepackagesByIDPatchNoContent with default headers values
func NewProfilepackagesByIDPatchNoContent() *ProfilepackagesByIDPatchNoContent {
	return &ProfilepackagesByIDPatchNoContent{}
}

/*ProfilepackagesByIDPatchNoContent handles this case with default header values.

ProfilepackagesByIDPatchNoContent profilepackages by Id patch no content
*/
type ProfilepackagesByIDPatchNoContent struct {
	Payload interface{}
}

func (o *ProfilepackagesByIDPatchNoContent) Error() string {
	return fmt.Sprintf("[PATCH /profilepackages/{id}][%d] profilepackagesByIdPatchNoContent  %+v", 204, o.Payload)
}

func (o *ProfilepackagesByIDPatchNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ProfilepackagesByIDPatchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
