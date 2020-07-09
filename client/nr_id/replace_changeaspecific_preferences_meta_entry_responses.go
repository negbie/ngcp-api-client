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

// ReplaceChangeaspecificPreferencesMetaEntryReader is a Reader for the ReplaceChangeaspecificPreferencesMetaEntry structure.
type ReplaceChangeaspecificPreferencesMetaEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceChangeaspecificPreferencesMetaEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceChangeaspecificPreferencesMetaEntryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceChangeaspecificPreferencesMetaEntryOK creates a ReplaceChangeaspecificPreferencesMetaEntryOK with default headers values
func NewReplaceChangeaspecificPreferencesMetaEntryOK() *ReplaceChangeaspecificPreferencesMetaEntryOK {
	return &ReplaceChangeaspecificPreferencesMetaEntryOK{}
}

/*ReplaceChangeaspecificPreferencesMetaEntryOK handles this case with default header values.

ReplaceChangeaspecificPreferencesMetaEntryOK replace changeaspecific preferences meta entry o k
*/
type ReplaceChangeaspecificPreferencesMetaEntryOK struct {
	Payload *models.PreferencesMetaEntries
}

func (o *ReplaceChangeaspecificPreferencesMetaEntryOK) Error() string {
	return fmt.Sprintf("[PUT /preferencesmetaentries/{id}][%d] replaceChangeaspecificPreferencesMetaEntryOK  %+v", 200, o.Payload)
}

func (o *ReplaceChangeaspecificPreferencesMetaEntryOK) GetPayload() *models.PreferencesMetaEntries {
	return o.Payload
}

func (o *ReplaceChangeaspecificPreferencesMetaEntryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PreferencesMetaEntries)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
