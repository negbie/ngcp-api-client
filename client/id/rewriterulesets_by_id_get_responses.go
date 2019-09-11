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

// RewriterulesetsByIDGetReader is a Reader for the RewriterulesetsByIDGet structure.
type RewriterulesetsByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RewriterulesetsByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRewriterulesetsByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRewriterulesetsByIDGetOK creates a RewriterulesetsByIDGetOK with default headers values
func NewRewriterulesetsByIDGetOK() *RewriterulesetsByIDGetOK {
	return &RewriterulesetsByIDGetOK{}
}

/*RewriterulesetsByIDGetOK handles this case with default header values.

RewriterulesetsByIDGetOK rewriterulesets by Id get o k
*/
type RewriterulesetsByIDGetOK struct {
	Payload *models.RewriteRuleSets
}

func (o *RewriterulesetsByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /rewriterulesets/{id}][%d] rewriterulesetsByIdGetOK  %+v", 200, o.Payload)
}

func (o *RewriterulesetsByIDGetOK) GetPayload() *models.RewriteRuleSets {
	return o.Payload
}

func (o *RewriterulesetsByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RewriteRuleSets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}