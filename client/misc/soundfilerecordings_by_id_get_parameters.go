// Code generated by go-swagger; DO NOT EDIT.

package misc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSoundfilerecordingsByIDGetParams creates a new SoundfilerecordingsByIDGetParams object
// with the default values initialized.
func NewSoundfilerecordingsByIDGetParams() *SoundfilerecordingsByIDGetParams {
	var ()
	return &SoundfilerecordingsByIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSoundfilerecordingsByIDGetParamsWithTimeout creates a new SoundfilerecordingsByIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSoundfilerecordingsByIDGetParamsWithTimeout(timeout time.Duration) *SoundfilerecordingsByIDGetParams {
	var ()
	return &SoundfilerecordingsByIDGetParams{

		timeout: timeout,
	}
}

// NewSoundfilerecordingsByIDGetParamsWithContext creates a new SoundfilerecordingsByIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSoundfilerecordingsByIDGetParamsWithContext(ctx context.Context) *SoundfilerecordingsByIDGetParams {
	var ()
	return &SoundfilerecordingsByIDGetParams{

		Context: ctx,
	}
}

// NewSoundfilerecordingsByIDGetParamsWithHTTPClient creates a new SoundfilerecordingsByIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSoundfilerecordingsByIDGetParamsWithHTTPClient(client *http.Client) *SoundfilerecordingsByIDGetParams {
	var ()
	return &SoundfilerecordingsByIDGetParams{
		HTTPClient: client,
	}
}

/*SoundfilerecordingsByIDGetParams contains all the parameters to send to the API endpoint
for the soundfilerecordings by Id get operation typically these are written to a http.Request
*/
type SoundfilerecordingsByIDGetParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the soundfilerecordings by Id get params
func (o *SoundfilerecordingsByIDGetParams) WithTimeout(timeout time.Duration) *SoundfilerecordingsByIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the soundfilerecordings by Id get params
func (o *SoundfilerecordingsByIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the soundfilerecordings by Id get params
func (o *SoundfilerecordingsByIDGetParams) WithContext(ctx context.Context) *SoundfilerecordingsByIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the soundfilerecordings by Id get params
func (o *SoundfilerecordingsByIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the soundfilerecordings by Id get params
func (o *SoundfilerecordingsByIDGetParams) WithHTTPClient(client *http.Client) *SoundfilerecordingsByIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the soundfilerecordings by Id get params
func (o *SoundfilerecordingsByIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the soundfilerecordings by Id get params
func (o *SoundfilerecordingsByIDGetParams) WithID(id string) *SoundfilerecordingsByIDGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the soundfilerecordings by Id get params
func (o *SoundfilerecordingsByIDGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SoundfilerecordingsByIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}