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

// RewriterulesetsPostReader is a Reader for the RewriterulesetsPost structure.
type RewriterulesetsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RewriterulesetsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRewriterulesetsPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewRewriterulesetsPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRewriterulesetsPostCreated creates a RewriterulesetsPostCreated with default headers values
func NewRewriterulesetsPostCreated() *RewriterulesetsPostCreated {
	return &RewriterulesetsPostCreated{
		Location: "\"<string>\"",
	}
}

/*RewriterulesetsPostCreated handles this case with default header values.

RewriterulesetsPostCreated rewriterulesets post created
*/
type RewriterulesetsPostCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty35
}

func (o *RewriterulesetsPostCreated) Error() string {
	return fmt.Sprintf("[POST /rewriterulesets/][%d] rewriterulesetsPostCreated  %+v", 201, o.Payload)
}

func (o *RewriterulesetsPostCreated) GetPayload() []*models.Thenewlycreateditemorempty35 {
	return o.Payload
}

func (o *RewriterulesetsPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRewriterulesetsPostUnprocessableEntity creates a RewriterulesetsPostUnprocessableEntity with default headers values
func NewRewriterulesetsPostUnprocessableEntity() *RewriterulesetsPostUnprocessableEntity {
	return &RewriterulesetsPostUnprocessableEntity{}
}

/*RewriterulesetsPostUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type RewriterulesetsPostUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *RewriterulesetsPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /rewriterulesets/][%d] rewriterulesetsPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *RewriterulesetsPostUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *RewriterulesetsPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}