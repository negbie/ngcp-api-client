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

// CfmappingsByIDGetReader is a Reader for the CfmappingsByIDGet structure.
type CfmappingsByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CfmappingsByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCfmappingsByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCfmappingsByIDGetOK creates a CfmappingsByIDGetOK with default headers values
func NewCfmappingsByIDGetOK() *CfmappingsByIDGetOK {
	return &CfmappingsByIDGetOK{}
}

/*CfmappingsByIDGetOK handles this case with default header values.

CfmappingsByIDGetOK cfmappings by Id get o k
*/
type CfmappingsByIDGetOK struct {
	Payload *models.CFMappings
}

func (o *CfmappingsByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /cfmappings/{id}][%d] cfmappingsByIdGetOK  %+v", 200, o.Payload)
}

func (o *CfmappingsByIDGetOK) GetPayload() *models.CFMappings {
	return o.Payload
}

func (o *CfmappingsByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CFMappings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}