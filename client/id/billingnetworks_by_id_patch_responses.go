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

// BillingnetworksByIDPatchReader is a Reader for the BillingnetworksByIDPatch structure.
type BillingnetworksByIDPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingnetworksByIDPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBillingnetworksByIDPatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBillingnetworksByIDPatchOK creates a BillingnetworksByIDPatchOK with default headers values
func NewBillingnetworksByIDPatchOK() *BillingnetworksByIDPatchOK {
	return &BillingnetworksByIDPatchOK{}
}

/*BillingnetworksByIDPatchOK handles this case with default header values.

BillingnetworksByIDPatchOK billingnetworks by Id patch o k
*/
type BillingnetworksByIDPatchOK struct {
	Payload *models.BillingNetworks
}

func (o *BillingnetworksByIDPatchOK) Error() string {
	return fmt.Sprintf("[PATCH /billingnetworks/{id}][%d] billingnetworksByIdPatchOK  %+v", 200, o.Payload)
}

func (o *BillingnetworksByIDPatchOK) GetPayload() *models.BillingNetworks {
	return o.Payload
}

func (o *BillingnetworksByIDPatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingNetworks)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
