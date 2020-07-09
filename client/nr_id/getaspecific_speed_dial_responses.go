// Code generated by go-swagger; DO NOT EDIT.

package nr_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// GetaspecificSpeedDialReader is a Reader for the GetaspecificSpeedDial structure.
type GetaspecificSpeedDialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificSpeedDialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificSpeedDialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificSpeedDialOK creates a GetaspecificSpeedDialOK with default headers values
func NewGetaspecificSpeedDialOK() *GetaspecificSpeedDialOK {
	return &GetaspecificSpeedDialOK{}
}

/*GetaspecificSpeedDialOK handles this case with default header values.

GetaspecificSpeedDialOK getaspecific speed dial o k
*/
type GetaspecificSpeedDialOK struct {
	Payload *models.SpeedDials
}

func (o *GetaspecificSpeedDialOK) Error() string {
	return fmt.Sprintf("[GET /speeddials/{id}][%d] getaspecificSpeedDialOK  %+v", 200, o.Payload)
}

func (o *GetaspecificSpeedDialOK) GetPayload() *models.SpeedDials {
	return o.Payload
}

func (o *GetaspecificSpeedDialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SpeedDials)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}