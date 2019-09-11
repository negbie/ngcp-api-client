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

// AutoattendantsByIDPutReader is a Reader for the AutoattendantsByIDPut structure.
type AutoattendantsByIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AutoattendantsByIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAutoattendantsByIDPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAutoattendantsByIDPutOK creates a AutoattendantsByIDPutOK with default headers values
func NewAutoattendantsByIDPutOK() *AutoattendantsByIDPutOK {
	return &AutoattendantsByIDPutOK{}
}

/*AutoattendantsByIDPutOK handles this case with default header values.

AutoattendantsByIDPutOK autoattendants by Id put o k
*/
type AutoattendantsByIDPutOK struct {
	Payload *models.AutoAttendants
}

func (o *AutoattendantsByIDPutOK) Error() string {
	return fmt.Sprintf("[PUT /autoattendants/{id}][%d] autoattendantsByIdPutOK  %+v", 200, o.Payload)
}

func (o *AutoattendantsByIDPutOK) GetPayload() *models.AutoAttendants {
	return o.Payload
}

func (o *AutoattendantsByIDPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AutoAttendants)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
