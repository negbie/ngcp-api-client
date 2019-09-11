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

// LnpcarriersByIDPutReader is a Reader for the LnpcarriersByIDPut structure.
type LnpcarriersByIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LnpcarriersByIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLnpcarriersByIDPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLnpcarriersByIDPutOK creates a LnpcarriersByIDPutOK with default headers values
func NewLnpcarriersByIDPutOK() *LnpcarriersByIDPutOK {
	return &LnpcarriersByIDPutOK{}
}

/*LnpcarriersByIDPutOK handles this case with default header values.

LnpcarriersByIDPutOK lnpcarriers by Id put o k
*/
type LnpcarriersByIDPutOK struct {
	Payload *models.LnpCarriers
}

func (o *LnpcarriersByIDPutOK) Error() string {
	return fmt.Sprintf("[PUT /lnpcarriers/{id}][%d] lnpcarriersByIdPutOK  %+v", 200, o.Payload)
}

func (o *LnpcarriersByIDPutOK) GetPayload() *models.LnpCarriers {
	return o.Payload
}

func (o *LnpcarriersByIDPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LnpCarriers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}