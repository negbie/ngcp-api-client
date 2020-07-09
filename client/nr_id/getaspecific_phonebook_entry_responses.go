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

// GetaspecificPhonebookEntryReader is a Reader for the GetaspecificPhonebookEntry structure.
type GetaspecificPhonebookEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificPhonebookEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificPhonebookEntryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificPhonebookEntryOK creates a GetaspecificPhonebookEntryOK with default headers values
func NewGetaspecificPhonebookEntryOK() *GetaspecificPhonebookEntryOK {
	return &GetaspecificPhonebookEntryOK{}
}

/*GetaspecificPhonebookEntryOK handles this case with default header values.

GetaspecificPhonebookEntryOK getaspecific phonebook entry o k
*/
type GetaspecificPhonebookEntryOK struct {
	Payload *models.PhonebookEntries
}

func (o *GetaspecificPhonebookEntryOK) Error() string {
	return fmt.Sprintf("[GET /phonebookentries/{id}][%d] getaspecificPhonebookEntryOK  %+v", 200, o.Payload)
}

func (o *GetaspecificPhonebookEntryOK) GetPayload() *models.PhonebookEntries {
	return o.Payload
}

func (o *GetaspecificPhonebookEntryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PhonebookEntries)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
