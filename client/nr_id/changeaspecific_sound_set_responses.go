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

// ChangeaspecificSoundSetReader is a Reader for the ChangeaspecificSoundSet structure.
type ChangeaspecificSoundSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificSoundSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeaspecificSoundSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificSoundSetOK creates a ChangeaspecificSoundSetOK with default headers values
func NewChangeaspecificSoundSetOK() *ChangeaspecificSoundSetOK {
	return &ChangeaspecificSoundSetOK{}
}

/*ChangeaspecificSoundSetOK handles this case with default header values.

ChangeaspecificSoundSetOK changeaspecific sound set o k
*/
type ChangeaspecificSoundSetOK struct {
	Payload *models.SoundSets
}

func (o *ChangeaspecificSoundSetOK) Error() string {
	return fmt.Sprintf("[PATCH /soundsets/{id}][%d] changeaspecificSoundSetOK  %+v", 200, o.Payload)
}

func (o *ChangeaspecificSoundSetOK) GetPayload() *models.SoundSets {
	return o.Payload
}

func (o *ChangeaspecificSoundSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SoundSets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
