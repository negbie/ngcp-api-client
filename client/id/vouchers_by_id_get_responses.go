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

// VouchersByIDGetReader is a Reader for the VouchersByIDGet structure.
type VouchersByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VouchersByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVouchersByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVouchersByIDGetOK creates a VouchersByIDGetOK with default headers values
func NewVouchersByIDGetOK() *VouchersByIDGetOK {
	return &VouchersByIDGetOK{}
}

/*VouchersByIDGetOK handles this case with default header values.

VouchersByIDGetOK vouchers by Id get o k
*/
type VouchersByIDGetOK struct {
	Payload *models.Vouchers
}

func (o *VouchersByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /vouchers/{id}][%d] vouchersByIdGetOK  %+v", 200, o.Payload)
}

func (o *VouchersByIDGetOK) GetPayload() *models.Vouchers {
	return o.Payload
}

func (o *VouchersByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Vouchers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}