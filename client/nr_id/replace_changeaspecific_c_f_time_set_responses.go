// Code generated by go-swagger; DO NOT EDIT.

package nr_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// ReplaceChangeaspecificCFTimeSetReader is a Reader for the ReplaceChangeaspecificCFTimeSet structure.
type ReplaceChangeaspecificCFTimeSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceChangeaspecificCFTimeSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceChangeaspecificCFTimeSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceChangeaspecificCFTimeSetOK creates a ReplaceChangeaspecificCFTimeSetOK with default headers values
func NewReplaceChangeaspecificCFTimeSetOK() *ReplaceChangeaspecificCFTimeSetOK {
	return &ReplaceChangeaspecificCFTimeSetOK{}
}

/*ReplaceChangeaspecificCFTimeSetOK handles this case with default header values.

ReplaceChangeaspecificCFTimeSetOK replace changeaspecific c f time set o k
*/
type ReplaceChangeaspecificCFTimeSetOK struct {
	Payload *models.CFTimeSets
}

func (o *ReplaceChangeaspecificCFTimeSetOK) Error() string {
	return fmt.Sprintf("[PUT /cftimesets/{id}][%d] replaceChangeaspecificCFTimeSetOK  %+v", 200, o.Payload)
}

func (o *ReplaceChangeaspecificCFTimeSetOK) GetPayload() *models.CFTimeSets {
	return o.Payload
}

func (o *ReplaceChangeaspecificCFTimeSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CFTimeSets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}