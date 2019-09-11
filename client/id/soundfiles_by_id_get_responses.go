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

// SoundfilesByIDGetReader is a Reader for the SoundfilesByIDGet structure.
type SoundfilesByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SoundfilesByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSoundfilesByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSoundfilesByIDGetOK creates a SoundfilesByIDGetOK with default headers values
func NewSoundfilesByIDGetOK() *SoundfilesByIDGetOK {
	return &SoundfilesByIDGetOK{}
}

/*SoundfilesByIDGetOK handles this case with default header values.

SoundfilesByIDGetOK soundfiles by Id get o k
*/
type SoundfilesByIDGetOK struct {
	Payload *models.SoundFiles
}

func (o *SoundfilesByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /soundfiles/{id}][%d] soundfilesByIdGetOK  %+v", 200, o.Payload)
}

func (o *SoundfilesByIDGetOK) GetPayload() *models.SoundFiles {
	return o.Payload
}

func (o *SoundfilesByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SoundFiles)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}