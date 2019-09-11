// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/negbie/ngcp-api/models"
)

// CfdestinationsetsGetReader is a Reader for the CfdestinationsetsGet structure.
type CfdestinationsetsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CfdestinationsetsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCfdestinationsetsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCfdestinationsetsGetOK creates a CfdestinationsetsGetOK with default headers values
func NewCfdestinationsetsGetOK() *CfdestinationsetsGetOK {
	return &CfdestinationsetsGetOK{}
}

/*CfdestinationsetsGetOK handles this case with default header values.

CfdestinationsetsGetOK cfdestinationsets get o k
*/
type CfdestinationsetsGetOK struct {
	Payload []*models.CFDestinationSet
}

func (o *CfdestinationsetsGetOK) Error() string {
	return fmt.Sprintf("[GET /cfdestinationsets][%d] cfdestinationsetsGetOK  %+v", 200, o.Payload)
}

func (o *CfdestinationsetsGetOK) GetPayload() []*models.CFDestinationSet {
	return o.Payload
}

func (o *CfdestinationsetsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
