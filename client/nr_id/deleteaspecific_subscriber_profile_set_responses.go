// Code generated by go-swagger; DO NOT EDIT.

package nr_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteaspecificSubscriberProfileSetReader is a Reader for the DeleteaspecificSubscriberProfileSet structure.
type DeleteaspecificSubscriberProfileSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteaspecificSubscriberProfileSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteaspecificSubscriberProfileSetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteaspecificSubscriberProfileSetNoContent creates a DeleteaspecificSubscriberProfileSetNoContent with default headers values
func NewDeleteaspecificSubscriberProfileSetNoContent() *DeleteaspecificSubscriberProfileSetNoContent {
	return &DeleteaspecificSubscriberProfileSetNoContent{}
}

/*DeleteaspecificSubscriberProfileSetNoContent handles this case with default header values.

DeleteaspecificSubscriberProfileSetNoContent deleteaspecific subscriber profile set no content
*/
type DeleteaspecificSubscriberProfileSetNoContent struct {
	Payload interface{}
}

func (o *DeleteaspecificSubscriberProfileSetNoContent) Error() string {
	return fmt.Sprintf("[DELETE /subscriberprofilesets/{id}][%d] deleteaspecificSubscriberProfileSetNoContent  %+v", 204, o.Payload)
}

func (o *DeleteaspecificSubscriberProfileSetNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteaspecificSubscriberProfileSetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}