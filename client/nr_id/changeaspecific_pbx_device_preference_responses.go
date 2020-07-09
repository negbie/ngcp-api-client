// Code generated by go-swagger; DO NOT EDIT.

package nr_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ChangeaspecificPbxDevicePreferenceReader is a Reader for the ChangeaspecificPbxDevicePreference structure.
type ChangeaspecificPbxDevicePreferenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificPbxDevicePreferenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeaspecificPbxDevicePreferenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificPbxDevicePreferenceOK creates a ChangeaspecificPbxDevicePreferenceOK with default headers values
func NewChangeaspecificPbxDevicePreferenceOK() *ChangeaspecificPbxDevicePreferenceOK {
	return &ChangeaspecificPbxDevicePreferenceOK{}
}

/*ChangeaspecificPbxDevicePreferenceOK handles this case with default header values.

ChangeaspecificPbxDevicePreferenceOK changeaspecific pbx device preference o k
*/
type ChangeaspecificPbxDevicePreferenceOK struct {
	Payload string
}

func (o *ChangeaspecificPbxDevicePreferenceOK) Error() string {
	return fmt.Sprintf("[PATCH /pbxdevicepreferences/{id}][%d] changeaspecificPbxDevicePreferenceOK  %+v", 200, o.Payload)
}

func (o *ChangeaspecificPbxDevicePreferenceOK) GetPayload() string {
	return o.Payload
}

func (o *ChangeaspecificPbxDevicePreferenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}