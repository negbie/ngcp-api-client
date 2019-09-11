// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/negbie/ngcp-api/models"
)

// ConversationsGetReader is a Reader for the ConversationsGet structure.
type ConversationsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConversationsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConversationsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewConversationsGetOK creates a ConversationsGetOK with default headers values
func NewConversationsGetOK() *ConversationsGetOK {
	return &ConversationsGetOK{}
}

/*ConversationsGetOK handles this case with default header values.

ConversationsGetOK conversations get o k
*/
type ConversationsGetOK struct {
	Payload []*models.Conversation
}

func (o *ConversationsGetOK) Error() string {
	return fmt.Sprintf("[GET /conversations][%d] conversationsGetOK  %+v", 200, o.Payload)
}

func (o *ConversationsGetOK) GetPayload() []*models.Conversation {
	return o.Payload
}

func (o *ConversationsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
