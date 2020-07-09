// Code generated by go-swagger; DO NOT EDIT.

package topuplogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// GetTopupLogitemsReader is a Reader for the GetTopupLogitems structure.
type GetTopupLogitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTopupLogitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTopupLogitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTopupLogitemsOK creates a GetTopupLogitemsOK with default headers values
func NewGetTopupLogitemsOK() *GetTopupLogitemsOK {
	return &GetTopupLogitemsOK{}
}

/*GetTopupLogitemsOK handles this case with default header values.

GetTopupLogitemsOK get topup logitems o k
*/
type GetTopupLogitemsOK struct {
	Payload []*models.TopupLog
}

func (o *GetTopupLogitemsOK) Error() string {
	return fmt.Sprintf("[GET /topuplogs][%d] getTopupLogitemsOK  %+v", 200, o.Payload)
}

func (o *GetTopupLogitemsOK) GetPayload() []*models.TopupLog {
	return o.Payload
}

func (o *GetTopupLogitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
