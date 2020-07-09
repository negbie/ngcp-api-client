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

// ReplaceChangeaspecificPeeringRuleReader is a Reader for the ReplaceChangeaspecificPeeringRule structure.
type ReplaceChangeaspecificPeeringRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceChangeaspecificPeeringRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceChangeaspecificPeeringRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceChangeaspecificPeeringRuleOK creates a ReplaceChangeaspecificPeeringRuleOK with default headers values
func NewReplaceChangeaspecificPeeringRuleOK() *ReplaceChangeaspecificPeeringRuleOK {
	return &ReplaceChangeaspecificPeeringRuleOK{}
}

/*ReplaceChangeaspecificPeeringRuleOK handles this case with default header values.

ReplaceChangeaspecificPeeringRuleOK replace changeaspecific peering rule o k
*/
type ReplaceChangeaspecificPeeringRuleOK struct {
	Payload *models.PeeringRules
}

func (o *ReplaceChangeaspecificPeeringRuleOK) Error() string {
	return fmt.Sprintf("[PUT /peeringrules/{id}][%d] replaceChangeaspecificPeeringRuleOK  %+v", 200, o.Payload)
}

func (o *ReplaceChangeaspecificPeeringRuleOK) GetPayload() *models.PeeringRules {
	return o.Payload
}

func (o *ReplaceChangeaspecificPeeringRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PeeringRules)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
