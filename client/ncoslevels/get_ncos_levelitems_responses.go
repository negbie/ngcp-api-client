// Code generated by go-swagger; DO NOT EDIT.

package ncoslevels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// GetNcosLevelitemsReader is a Reader for the GetNcosLevelitems structure.
type GetNcosLevelitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNcosLevelitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNcosLevelitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNcosLevelitemsOK creates a GetNcosLevelitemsOK with default headers values
func NewGetNcosLevelitemsOK() *GetNcosLevelitemsOK {
	return &GetNcosLevelitemsOK{}
}

/*GetNcosLevelitemsOK handles this case with default header values.

GetNcosLevelitemsOK get ncos levelitems o k
*/
type GetNcosLevelitemsOK struct {
	Payload []*models.NcosLevels3
}

func (o *GetNcosLevelitemsOK) Error() string {
	return fmt.Sprintf("[GET /ncoslevels][%d] getNcosLevelitemsOK  %+v", 200, o.Payload)
}

func (o *GetNcosLevelitemsOK) GetPayload() []*models.NcosLevels3 {
	return o.Payload
}

func (o *GetNcosLevelitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
