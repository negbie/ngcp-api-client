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

// DeleteaspecificCustomerLocationReader is a Reader for the DeleteaspecificCustomerLocation structure.
type DeleteaspecificCustomerLocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteaspecificCustomerLocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteaspecificCustomerLocationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteaspecificCustomerLocationNoContent creates a DeleteaspecificCustomerLocationNoContent with default headers values
func NewDeleteaspecificCustomerLocationNoContent() *DeleteaspecificCustomerLocationNoContent {
	return &DeleteaspecificCustomerLocationNoContent{}
}

/*DeleteaspecificCustomerLocationNoContent handles this case with default header values.

DeleteaspecificCustomerLocationNoContent deleteaspecific customer location no content
*/
type DeleteaspecificCustomerLocationNoContent struct {
	Payload interface{}
}

func (o *DeleteaspecificCustomerLocationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /customerlocations/{id}][%d] deleteaspecificCustomerLocationNoContent  %+v", 204, o.Payload)
}

func (o *DeleteaspecificCustomerLocationNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteaspecificCustomerLocationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}