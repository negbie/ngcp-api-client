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

// PhonebookentriesByIDPatchReader is a Reader for the PhonebookentriesByIDPatch structure.
type PhonebookentriesByIDPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PhonebookentriesByIDPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPhonebookentriesByIDPatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPhonebookentriesByIDPatchOK creates a PhonebookentriesByIDPatchOK with default headers values
func NewPhonebookentriesByIDPatchOK() *PhonebookentriesByIDPatchOK {
	return &PhonebookentriesByIDPatchOK{}
}

/*PhonebookentriesByIDPatchOK handles this case with default header values.

PhonebookentriesByIDPatchOK phonebookentries by Id patch o k
*/
type PhonebookentriesByIDPatchOK struct {
	Payload *models.PhonebookEntries
}

func (o *PhonebookentriesByIDPatchOK) Error() string {
	return fmt.Sprintf("[PATCH /phonebookentries/{id}][%d] phonebookentriesByIdPatchOK  %+v", 200, o.Payload)
}

func (o *PhonebookentriesByIDPatchOK) GetPayload() *models.PhonebookEntries {
	return o.Payload
}

func (o *PhonebookentriesByIDPatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PhonebookEntries)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}