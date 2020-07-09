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

// GetaspecificCCMapEntryReader is a Reader for the GetaspecificCCMapEntry structure.
type GetaspecificCCMapEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificCCMapEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificCCMapEntryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificCCMapEntryOK creates a GetaspecificCCMapEntryOK with default headers values
func NewGetaspecificCCMapEntryOK() *GetaspecificCCMapEntryOK {
	return &GetaspecificCCMapEntryOK{}
}

/*GetaspecificCCMapEntryOK handles this case with default header values.

GetaspecificCCMapEntryOK getaspecific c c map entry o k
*/
type GetaspecificCCMapEntryOK struct {
	Payload *models.CCMapEntries
}

func (o *GetaspecificCCMapEntryOK) Error() string {
	return fmt.Sprintf("[GET /ccmapentries/{id}][%d] getaspecificCCMapEntryOK  %+v", 200, o.Payload)
}

func (o *GetaspecificCCMapEntryOK) GetPayload() *models.CCMapEntries {
	return o.Payload
}

func (o *GetaspecificCCMapEntryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CCMapEntries)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
