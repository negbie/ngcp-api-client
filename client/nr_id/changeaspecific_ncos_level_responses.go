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

// ChangeaspecificNcosLevelReader is a Reader for the ChangeaspecificNcosLevel structure.
type ChangeaspecificNcosLevelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificNcosLevelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeaspecificNcosLevelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificNcosLevelOK creates a ChangeaspecificNcosLevelOK with default headers values
func NewChangeaspecificNcosLevelOK() *ChangeaspecificNcosLevelOK {
	return &ChangeaspecificNcosLevelOK{}
}

/*ChangeaspecificNcosLevelOK handles this case with default header values.

ChangeaspecificNcosLevelOK changeaspecific ncos level o k
*/
type ChangeaspecificNcosLevelOK struct {
	Payload *models.NcosLevels
}

func (o *ChangeaspecificNcosLevelOK) Error() string {
	return fmt.Sprintf("[PATCH /ncoslevels/{id}][%d] changeaspecificNcosLevelOK  %+v", 200, o.Payload)
}

func (o *ChangeaspecificNcosLevelOK) GetPayload() *models.NcosLevels {
	return o.Payload
}

func (o *ChangeaspecificNcosLevelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NcosLevels)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
