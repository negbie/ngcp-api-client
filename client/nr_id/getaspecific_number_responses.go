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

// GetaspecificNumberReader is a Reader for the GetaspecificNumber structure.
type GetaspecificNumberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificNumberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificNumberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificNumberOK creates a GetaspecificNumberOK with default headers values
func NewGetaspecificNumberOK() *GetaspecificNumberOK {
	return &GetaspecificNumberOK{}
}

/*GetaspecificNumberOK handles this case with default header values.

GetaspecificNumberOK getaspecific number o k
*/
type GetaspecificNumberOK struct {
	Payload *models.Numbers
}

func (o *GetaspecificNumberOK) Error() string {
	return fmt.Sprintf("[GET /numbers/{id}][%d] getaspecificNumberOK  %+v", 200, o.Payload)
}

func (o *GetaspecificNumberOK) GetPayload() *models.Numbers {
	return o.Payload
}

func (o *GetaspecificNumberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Numbers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
