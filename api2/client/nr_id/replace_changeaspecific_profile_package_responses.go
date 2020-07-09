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

// ReplaceChangeaspecificProfilePackageReader is a Reader for the ReplaceChangeaspecificProfilePackage structure.
type ReplaceChangeaspecificProfilePackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceChangeaspecificProfilePackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceChangeaspecificProfilePackageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceChangeaspecificProfilePackageOK creates a ReplaceChangeaspecificProfilePackageOK with default headers values
func NewReplaceChangeaspecificProfilePackageOK() *ReplaceChangeaspecificProfilePackageOK {
	return &ReplaceChangeaspecificProfilePackageOK{}
}

/*ReplaceChangeaspecificProfilePackageOK handles this case with default header values.

ReplaceChangeaspecificProfilePackageOK replace changeaspecific profile package o k
*/
type ReplaceChangeaspecificProfilePackageOK struct {
	Payload *models.ProfilePackages
}

func (o *ReplaceChangeaspecificProfilePackageOK) Error() string {
	return fmt.Sprintf("[PUT /profilepackages/{id}][%d] replaceChangeaspecificProfilePackageOK  %+v", 200, o.Payload)
}

func (o *ReplaceChangeaspecificProfilePackageOK) GetPayload() *models.ProfilePackages {
	return o.Payload
}

func (o *ReplaceChangeaspecificProfilePackageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProfilePackages)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}