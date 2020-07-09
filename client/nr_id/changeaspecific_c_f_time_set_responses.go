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

// ChangeaspecificCFTimeSetReader is a Reader for the ChangeaspecificCFTimeSet structure.
type ChangeaspecificCFTimeSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificCFTimeSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeaspecificCFTimeSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificCFTimeSetOK creates a ChangeaspecificCFTimeSetOK with default headers values
func NewChangeaspecificCFTimeSetOK() *ChangeaspecificCFTimeSetOK {
	return &ChangeaspecificCFTimeSetOK{}
}

/*ChangeaspecificCFTimeSetOK handles this case with default header values.

ChangeaspecificCFTimeSetOK changeaspecific c f time set o k
*/
type ChangeaspecificCFTimeSetOK struct {
	Payload *models.CFTimeSets
}

func (o *ChangeaspecificCFTimeSetOK) Error() string {
	return fmt.Sprintf("[PATCH /cftimesets/{id}][%d] changeaspecificCFTimeSetOK  %+v", 200, o.Payload)
}

func (o *ChangeaspecificCFTimeSetOK) GetPayload() *models.CFTimeSets {
	return o.Payload
}

func (o *ChangeaspecificCFTimeSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CFTimeSets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
