// Code generated by go-swagger; DO NOT EDIT.

package contracts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// GetContractitemsReader is a Reader for the GetContractitems structure.
type GetContractitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContractitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContractitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetContractitemsOK creates a GetContractitemsOK with default headers values
func NewGetContractitemsOK() *GetContractitemsOK {
	return &GetContractitemsOK{}
}

/*GetContractitemsOK handles this case with default header values.

GetContractitemsOK get contractitems o k
*/
type GetContractitemsOK struct {
	Payload []*models.Contracts
}

func (o *GetContractitemsOK) Error() string {
	return fmt.Sprintf("[GET /contracts][%d] getContractitemsOK  %+v", 200, o.Payload)
}

func (o *GetContractitemsOK) GetPayload() []*models.Contracts {
	return o.Payload
}

func (o *GetContractitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}