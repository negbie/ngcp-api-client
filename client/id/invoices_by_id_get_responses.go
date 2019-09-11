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

// InvoicesByIDGetReader is a Reader for the InvoicesByIDGet structure.
type InvoicesByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InvoicesByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInvoicesByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInvoicesByIDGetOK creates a InvoicesByIDGetOK with default headers values
func NewInvoicesByIDGetOK() *InvoicesByIDGetOK {
	return &InvoicesByIDGetOK{}
}

/*InvoicesByIDGetOK handles this case with default header values.

InvoicesByIDGetOK invoices by Id get o k
*/
type InvoicesByIDGetOK struct {
	Payload *models.Invoices
}

func (o *InvoicesByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /invoices/{id}][%d] invoicesByIdGetOK  %+v", 200, o.Payload)
}

func (o *InvoicesByIDGetOK) GetPayload() *models.Invoices {
	return o.Payload
}

func (o *InvoicesByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Invoices)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
