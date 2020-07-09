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

// ChangeaspecificBillingFeeReader is a Reader for the ChangeaspecificBillingFee structure.
type ChangeaspecificBillingFeeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificBillingFeeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeaspecificBillingFeeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificBillingFeeOK creates a ChangeaspecificBillingFeeOK with default headers values
func NewChangeaspecificBillingFeeOK() *ChangeaspecificBillingFeeOK {
	return &ChangeaspecificBillingFeeOK{}
}

/*ChangeaspecificBillingFeeOK handles this case with default header values.

ChangeaspecificBillingFeeOK changeaspecific billing fee o k
*/
type ChangeaspecificBillingFeeOK struct {
	Payload *models.BillingFees
}

func (o *ChangeaspecificBillingFeeOK) Error() string {
	return fmt.Sprintf("[PATCH /billingfees/{id}][%d] changeaspecificBillingFeeOK  %+v", 200, o.Payload)
}

func (o *ChangeaspecificBillingFeeOK) GetPayload() *models.BillingFees {
	return o.Payload
}

func (o *ChangeaspecificBillingFeeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingFees)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
