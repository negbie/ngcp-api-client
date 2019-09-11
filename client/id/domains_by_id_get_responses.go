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

// DomainsByIDGetReader is a Reader for the DomainsByIDGet structure.
type DomainsByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainsByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainsByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDomainsByIDGetOK creates a DomainsByIDGetOK with default headers values
func NewDomainsByIDGetOK() *DomainsByIDGetOK {
	return &DomainsByIDGetOK{}
}

/*DomainsByIDGetOK handles this case with default header values.

DomainsByIDGetOK domains by Id get o k
*/
type DomainsByIDGetOK struct {
	Payload *models.Domains
}

func (o *DomainsByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /domains/{id}][%d] domainsByIdGetOK  %+v", 200, o.Payload)
}

func (o *DomainsByIDGetOK) GetPayload() *models.Domains {
	return o.Payload
}

func (o *DomainsByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Domains)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
