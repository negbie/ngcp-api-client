// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/negbie/ngcp-api/models"
)

// VoicemailgreetingsPostReader is a Reader for the VoicemailgreetingsPost structure.
type VoicemailgreetingsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VoicemailgreetingsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVoicemailgreetingsPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewVoicemailgreetingsPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVoicemailgreetingsPostCreated creates a VoicemailgreetingsPostCreated with default headers values
func NewVoicemailgreetingsPostCreated() *VoicemailgreetingsPostCreated {
	return &VoicemailgreetingsPostCreated{
		Location: "\"<string>\"",
	}
}

/*VoicemailgreetingsPostCreated handles this case with default header values.

VoicemailgreetingsPostCreated voicemailgreetings post created
*/
type VoicemailgreetingsPostCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty46
}

func (o *VoicemailgreetingsPostCreated) Error() string {
	return fmt.Sprintf("[POST /voicemailgreetings/][%d] voicemailgreetingsPostCreated  %+v", 201, o.Payload)
}

func (o *VoicemailgreetingsPostCreated) GetPayload() []*models.Thenewlycreateditemorempty46 {
	return o.Payload
}

func (o *VoicemailgreetingsPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoicemailgreetingsPostUnprocessableEntity creates a VoicemailgreetingsPostUnprocessableEntity with default headers values
func NewVoicemailgreetingsPostUnprocessableEntity() *VoicemailgreetingsPostUnprocessableEntity {
	return &VoicemailgreetingsPostUnprocessableEntity{}
}

/*VoicemailgreetingsPostUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type VoicemailgreetingsPostUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *VoicemailgreetingsPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /voicemailgreetings/][%d] voicemailgreetingsPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *VoicemailgreetingsPostUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *VoicemailgreetingsPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
