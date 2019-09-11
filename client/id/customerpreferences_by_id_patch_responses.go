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

// CustomerpreferencesByIDPatchReader is a Reader for the CustomerpreferencesByIDPatch structure.
type CustomerpreferencesByIDPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerpreferencesByIDPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCustomerpreferencesByIDPatchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomerpreferencesByIDPatchNoContent creates a CustomerpreferencesByIDPatchNoContent with default headers values
func NewCustomerpreferencesByIDPatchNoContent() *CustomerpreferencesByIDPatchNoContent {
	return &CustomerpreferencesByIDPatchNoContent{}
}

/*CustomerpreferencesByIDPatchNoContent handles this case with default header values.

CustomerpreferencesByIDPatchNoContent customerpreferences by Id patch no content
*/
type CustomerpreferencesByIDPatchNoContent struct {
	Payload interface{}
}

func (o *CustomerpreferencesByIDPatchNoContent) Error() string {
	return fmt.Sprintf("[PATCH /customerpreferences/{id}][%d] customerpreferencesByIdPatchNoContent  %+v", 204, o.Payload)
}

func (o *CustomerpreferencesByIDPatchNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerpreferencesByIDPatchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
