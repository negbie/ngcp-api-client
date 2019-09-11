// Code generated by go-swagger; DO NOT EDIT.

package id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpnrewritesetsByIDGetReader is a Reader for the UpnrewritesetsByIDGet structure.
type UpnrewritesetsByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpnrewritesetsByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpnrewritesetsByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpnrewritesetsByIDGetOK creates a UpnrewritesetsByIDGetOK with default headers values
func NewUpnrewritesetsByIDGetOK() *UpnrewritesetsByIDGetOK {
	return &UpnrewritesetsByIDGetOK{}
}

/*UpnrewritesetsByIDGetOK handles this case with default header values.

UpnrewritesetsByIDGetOK upnrewritesets by Id get o k
*/
type UpnrewritesetsByIDGetOK struct {
	Payload string
}

func (o *UpnrewritesetsByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /upnrewritesets/{id}][%d] upnrewritesetsByIdGetOK  %+v", 200, o.Payload)
}

func (o *UpnrewritesetsByIDGetOK) GetPayload() string {
	return o.Payload
}

func (o *UpnrewritesetsByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}