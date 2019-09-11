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

// CftimesetsGetReader is a Reader for the CftimesetsGet structure.
type CftimesetsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CftimesetsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCftimesetsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCftimesetsGetOK creates a CftimesetsGetOK with default headers values
func NewCftimesetsGetOK() *CftimesetsGetOK {
	return &CftimesetsGetOK{}
}

/*CftimesetsGetOK handles this case with default header values.

CftimesetsGetOK cftimesets get o k
*/
type CftimesetsGetOK struct {
	Payload []*models.CFTimeSets
}

func (o *CftimesetsGetOK) Error() string {
	return fmt.Sprintf("[GET /cftimesets][%d] cftimesetsGetOK  %+v", 200, o.Payload)
}

func (o *CftimesetsGetOK) GetPayload() []*models.CFTimeSets {
	return o.Payload
}

func (o *CftimesetsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}