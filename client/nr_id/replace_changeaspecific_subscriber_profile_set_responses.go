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

// ReplaceChangeaspecificSubscriberProfileSetReader is a Reader for the ReplaceChangeaspecificSubscriberProfileSet structure.
type ReplaceChangeaspecificSubscriberProfileSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceChangeaspecificSubscriberProfileSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceChangeaspecificSubscriberProfileSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceChangeaspecificSubscriberProfileSetOK creates a ReplaceChangeaspecificSubscriberProfileSetOK with default headers values
func NewReplaceChangeaspecificSubscriberProfileSetOK() *ReplaceChangeaspecificSubscriberProfileSetOK {
	return &ReplaceChangeaspecificSubscriberProfileSetOK{}
}

/*ReplaceChangeaspecificSubscriberProfileSetOK handles this case with default header values.

ReplaceChangeaspecificSubscriberProfileSetOK replace changeaspecific subscriber profile set o k
*/
type ReplaceChangeaspecificSubscriberProfileSetOK struct {
	Payload *models.SubscriberProfileSets
}

func (o *ReplaceChangeaspecificSubscriberProfileSetOK) Error() string {
	return fmt.Sprintf("[PUT /subscriberprofilesets/{id}][%d] replaceChangeaspecificSubscriberProfileSetOK  %+v", 200, o.Payload)
}

func (o *ReplaceChangeaspecificSubscriberProfileSetOK) GetPayload() *models.SubscriberProfileSets {
	return o.Payload
}

func (o *ReplaceChangeaspecificSubscriberProfileSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriberProfileSets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
