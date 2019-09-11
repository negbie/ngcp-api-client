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

// VouchersPostReader is a Reader for the VouchersPost structure.
type VouchersPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VouchersPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVouchersPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewVouchersPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVouchersPostCreated creates a VouchersPostCreated with default headers values
func NewVouchersPostCreated() *VouchersPostCreated {
	return &VouchersPostCreated{
		Location: "\"<string>\"",
	}
}

/*VouchersPostCreated handles this case with default header values.

VouchersPostCreated vouchers post created
*/
type VouchersPostCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty7
}

func (o *VouchersPostCreated) Error() string {
	return fmt.Sprintf("[POST /vouchers/][%d] vouchersPostCreated  %+v", 201, o.Payload)
}

func (o *VouchersPostCreated) GetPayload() []*models.Thenewlycreateditemorempty7 {
	return o.Payload
}

func (o *VouchersPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVouchersPostUnprocessableEntity creates a VouchersPostUnprocessableEntity with default headers values
func NewVouchersPostUnprocessableEntity() *VouchersPostUnprocessableEntity {
	return &VouchersPostUnprocessableEntity{}
}

/*VouchersPostUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type VouchersPostUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *VouchersPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /vouchers/][%d] vouchersPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *VouchersPostUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *VouchersPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}