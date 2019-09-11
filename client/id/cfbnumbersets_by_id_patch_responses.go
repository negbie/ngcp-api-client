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

// CfbnumbersetsByIDPatchReader is a Reader for the CfbnumbersetsByIDPatch structure.
type CfbnumbersetsByIDPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CfbnumbersetsByIDPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCfbnumbersetsByIDPatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCfbnumbersetsByIDPatchOK creates a CfbnumbersetsByIDPatchOK with default headers values
func NewCfbnumbersetsByIDPatchOK() *CfbnumbersetsByIDPatchOK {
	return &CfbnumbersetsByIDPatchOK{}
}

/*CfbnumbersetsByIDPatchOK handles this case with default header values.

CfbnumbersetsByIDPatchOK cfbnumbersets by Id patch o k
*/
type CfbnumbersetsByIDPatchOK struct {
	Payload *models.CFBNumberSets
}

func (o *CfbnumbersetsByIDPatchOK) Error() string {
	return fmt.Sprintf("[PATCH /cfbnumbersets/{id}][%d] cfbnumbersetsByIdPatchOK  %+v", 200, o.Payload)
}

func (o *CfbnumbersetsByIDPatchOK) GetPayload() *models.CFBNumberSets {
	return o.Payload
}

func (o *CfbnumbersetsByIDPatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CFBNumberSets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
