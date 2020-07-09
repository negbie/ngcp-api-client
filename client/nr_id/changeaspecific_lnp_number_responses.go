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

// ChangeaspecificLnpNumberReader is a Reader for the ChangeaspecificLnpNumber structure.
type ChangeaspecificLnpNumberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificLnpNumberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeaspecificLnpNumberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificLnpNumberOK creates a ChangeaspecificLnpNumberOK with default headers values
func NewChangeaspecificLnpNumberOK() *ChangeaspecificLnpNumberOK {
	return &ChangeaspecificLnpNumberOK{}
}

/*ChangeaspecificLnpNumberOK handles this case with default header values.

ChangeaspecificLnpNumberOK changeaspecific lnp number o k
*/
type ChangeaspecificLnpNumberOK struct {
	Payload *models.LnpNumbers
}

func (o *ChangeaspecificLnpNumberOK) Error() string {
	return fmt.Sprintf("[PATCH /lnpnumbers/{id}][%d] changeaspecificLnpNumberOK  %+v", 200, o.Payload)
}

func (o *ChangeaspecificLnpNumberOK) GetPayload() *models.LnpNumbers {
	return o.Payload
}

func (o *ChangeaspecificLnpNumberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LnpNumbers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
