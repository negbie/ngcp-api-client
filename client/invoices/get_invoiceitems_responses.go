// Code generated by go-swagger; DO NOT EDIT.

package invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// GetInvoiceitemsReader is a Reader for the GetInvoiceitems structure.
type GetInvoiceitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInvoiceitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInvoiceitemsOK creates a GetInvoiceitemsOK with default headers values
func NewGetInvoiceitemsOK() *GetInvoiceitemsOK {
	return &GetInvoiceitemsOK{}
}

/*GetInvoiceitemsOK handles this case with default header values.

GetInvoiceitemsOK get invoiceitems o k
*/
type GetInvoiceitemsOK struct {
	Payload []*models.Invoices
}

func (o *GetInvoiceitemsOK) Error() string {
	return fmt.Sprintf("[GET /invoices][%d] getInvoiceitemsOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceitemsOK) GetPayload() []*models.Invoices {
	return o.Payload
}

func (o *GetInvoiceitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
