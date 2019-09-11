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

// AdminsGetReader is a Reader for the AdminsGet structure.
type AdminsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAdminsGetOK creates a AdminsGetOK with default headers values
func NewAdminsGetOK() *AdminsGetOK {
	return &AdminsGetOK{}
}

/*AdminsGetOK handles this case with default header values.

AdminsGetOK admins get o k
*/
type AdminsGetOK struct {
	Payload []*models.Admin
}

func (o *AdminsGetOK) Error() string {
	return fmt.Sprintf("[GET /admins][%d] adminsGetOK  %+v", 200, o.Payload)
}

func (o *AdminsGetOK) GetPayload() []*models.Admin {
	return o.Payload
}

func (o *AdminsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
