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

// ChangeaspecificResellerReader is a Reader for the ChangeaspecificReseller structure.
type ChangeaspecificResellerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificResellerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeaspecificResellerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificResellerOK creates a ChangeaspecificResellerOK with default headers values
func NewChangeaspecificResellerOK() *ChangeaspecificResellerOK {
	return &ChangeaspecificResellerOK{}
}

/*ChangeaspecificResellerOK handles this case with default header values.

ChangeaspecificResellerOK changeaspecific reseller o k
*/
type ChangeaspecificResellerOK struct {
	Payload *models.Resellers
}

func (o *ChangeaspecificResellerOK) Error() string {
	return fmt.Sprintf("[PATCH /resellers/{id}][%d] changeaspecificResellerOK  %+v", 200, o.Payload)
}

func (o *ChangeaspecificResellerOK) GetPayload() *models.Resellers {
	return o.Payload
}

func (o *ChangeaspecificResellerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Resellers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
