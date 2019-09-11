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

// NcoslnpcarriersGetReader is a Reader for the NcoslnpcarriersGet structure.
type NcoslnpcarriersGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NcoslnpcarriersGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNcoslnpcarriersGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNcoslnpcarriersGetOK creates a NcoslnpcarriersGetOK with default headers values
func NewNcoslnpcarriersGetOK() *NcoslnpcarriersGetOK {
	return &NcoslnpcarriersGetOK{}
}

/*NcoslnpcarriersGetOK handles this case with default header values.

NcoslnpcarriersGetOK ncoslnpcarriers get o k
*/
type NcoslnpcarriersGetOK struct {
	Payload []*models.NcosLnpCarriers
}

func (o *NcoslnpcarriersGetOK) Error() string {
	return fmt.Sprintf("[GET /ncoslnpcarriers][%d] ncoslnpcarriersGetOK  %+v", 200, o.Payload)
}

func (o *NcoslnpcarriersGetOK) GetPayload() []*models.NcosLnpCarriers {
	return o.Payload
}

func (o *NcoslnpcarriersGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
