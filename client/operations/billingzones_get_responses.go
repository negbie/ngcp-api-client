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

// BillingzonesGetReader is a Reader for the BillingzonesGet structure.
type BillingzonesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingzonesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBillingzonesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBillingzonesGetOK creates a BillingzonesGetOK with default headers values
func NewBillingzonesGetOK() *BillingzonesGetOK {
	return &BillingzonesGetOK{}
}

/*BillingzonesGetOK handles this case with default header values.

BillingzonesGetOK billingzones get o k
*/
type BillingzonesGetOK struct {
	Payload []*models.BillingZones
}

func (o *BillingzonesGetOK) Error() string {
	return fmt.Sprintf("[GET /billingzones][%d] billingzonesGetOK  %+v", 200, o.Payload)
}

func (o *BillingzonesGetOK) GetPayload() []*models.BillingZones {
	return o.Payload
}

func (o *BillingzonesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
