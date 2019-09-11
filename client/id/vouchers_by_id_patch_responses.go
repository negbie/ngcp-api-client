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

// VouchersByIDPatchReader is a Reader for the VouchersByIDPatch structure.
type VouchersByIDPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VouchersByIDPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVouchersByIDPatchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVouchersByIDPatchNoContent creates a VouchersByIDPatchNoContent with default headers values
func NewVouchersByIDPatchNoContent() *VouchersByIDPatchNoContent {
	return &VouchersByIDPatchNoContent{}
}

/*VouchersByIDPatchNoContent handles this case with default header values.

VouchersByIDPatchNoContent vouchers by Id patch no content
*/
type VouchersByIDPatchNoContent struct {
	Payload interface{}
}

func (o *VouchersByIDPatchNoContent) Error() string {
	return fmt.Sprintf("[PATCH /vouchers/{id}][%d] vouchersByIdPatchNoContent  %+v", 204, o.Payload)
}

func (o *VouchersByIDPatchNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *VouchersByIDPatchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
