// Code generated by go-swagger; DO NOT EDIT.

package balanceintervals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/negbie/ngcp-api/models"
)

// BalanceintervalsGetReader is a Reader for the BalanceintervalsGet structure.
type BalanceintervalsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BalanceintervalsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBalanceintervalsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBalanceintervalsGetOK creates a BalanceintervalsGetOK with default headers values
func NewBalanceintervalsGetOK() *BalanceintervalsGetOK {
	return &BalanceintervalsGetOK{}
}

/*BalanceintervalsGetOK handles this case with default header values.

BalanceintervalsGetOK balanceintervals get o k
*/
type BalanceintervalsGetOK struct {
	Payload []*models.BalanceInterval
}

func (o *BalanceintervalsGetOK) Error() string {
	return fmt.Sprintf("[GET /balanceintervals][%d] balanceintervalsGetOK  %+v", 200, o.Payload)
}

func (o *BalanceintervalsGetOK) GetPayload() []*models.BalanceInterval {
	return o.Payload
}

func (o *BalanceintervalsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}