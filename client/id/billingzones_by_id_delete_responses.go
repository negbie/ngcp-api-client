// Code generated by go-swagger; DO NOT EDIT.

package id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// BillingzonesByIDDeleteReader is a Reader for the BillingzonesByIDDelete structure.
type BillingzonesByIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingzonesByIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewBillingzonesByIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBillingzonesByIDDeleteNoContent creates a BillingzonesByIDDeleteNoContent with default headers values
func NewBillingzonesByIDDeleteNoContent() *BillingzonesByIDDeleteNoContent {
	return &BillingzonesByIDDeleteNoContent{}
}

/*BillingzonesByIDDeleteNoContent handles this case with default header values.

BillingzonesByIDDeleteNoContent billingzones by Id delete no content
*/
type BillingzonesByIDDeleteNoContent struct {
	Payload interface{}
}

func (o *BillingzonesByIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /billingzones/{id}][%d] billingzonesByIdDeleteNoContent  %+v", 204, o.Payload)
}

func (o *BillingzonesByIDDeleteNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *BillingzonesByIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
