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

// PeeringgroupsGetReader is a Reader for the PeeringgroupsGet structure.
type PeeringgroupsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PeeringgroupsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPeeringgroupsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPeeringgroupsGetOK creates a PeeringgroupsGetOK with default headers values
func NewPeeringgroupsGetOK() *PeeringgroupsGetOK {
	return &PeeringgroupsGetOK{}
}

/*PeeringgroupsGetOK handles this case with default header values.

PeeringgroupsGetOK peeringgroups get o k
*/
type PeeringgroupsGetOK struct {
	Payload []*models.PeeringGroups
}

func (o *PeeringgroupsGetOK) Error() string {
	return fmt.Sprintf("[GET /peeringgroups][%d] peeringgroupsGetOK  %+v", 200, o.Payload)
}

func (o *PeeringgroupsGetOK) GetPayload() []*models.PeeringGroups {
	return o.Payload
}

func (o *PeeringgroupsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
