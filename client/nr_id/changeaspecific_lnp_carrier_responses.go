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

// ChangeaspecificLnpCarrierReader is a Reader for the ChangeaspecificLnpCarrier structure.
type ChangeaspecificLnpCarrierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificLnpCarrierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeaspecificLnpCarrierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificLnpCarrierOK creates a ChangeaspecificLnpCarrierOK with default headers values
func NewChangeaspecificLnpCarrierOK() *ChangeaspecificLnpCarrierOK {
	return &ChangeaspecificLnpCarrierOK{}
}

/*ChangeaspecificLnpCarrierOK handles this case with default header values.

ChangeaspecificLnpCarrierOK changeaspecific lnp carrier o k
*/
type ChangeaspecificLnpCarrierOK struct {
	Payload *models.LnpCarriers
}

func (o *ChangeaspecificLnpCarrierOK) Error() string {
	return fmt.Sprintf("[PATCH /lnpcarriers/{id}][%d] changeaspecificLnpCarrierOK  %+v", 200, o.Payload)
}

func (o *ChangeaspecificLnpCarrierOK) GetPayload() *models.LnpCarriers {
	return o.Payload
}

func (o *ChangeaspecificLnpCarrierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LnpCarriers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
