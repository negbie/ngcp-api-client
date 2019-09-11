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

// CfbnumbersetsByIDGetReader is a Reader for the CfbnumbersetsByIDGet structure.
type CfbnumbersetsByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CfbnumbersetsByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCfbnumbersetsByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCfbnumbersetsByIDGetOK creates a CfbnumbersetsByIDGetOK with default headers values
func NewCfbnumbersetsByIDGetOK() *CfbnumbersetsByIDGetOK {
	return &CfbnumbersetsByIDGetOK{}
}

/*CfbnumbersetsByIDGetOK handles this case with default header values.

CfbnumbersetsByIDGetOK cfbnumbersets by Id get o k
*/
type CfbnumbersetsByIDGetOK struct {
	Payload *models.CFBNumberSets
}

func (o *CfbnumbersetsByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /cfbnumbersets/{id}][%d] cfbnumbersetsByIdGetOK  %+v", 200, o.Payload)
}

func (o *CfbnumbersetsByIDGetOK) GetPayload() *models.CFBNumberSets {
	return o.Payload
}

func (o *CfbnumbersetsByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CFBNumberSets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}