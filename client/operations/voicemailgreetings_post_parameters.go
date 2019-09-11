// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	models "github.com/negbie/ngcp-api/models"
)

// NewVoicemailgreetingsPostParams creates a new VoicemailgreetingsPostParams object
// with the default values initialized.
func NewVoicemailgreetingsPostParams() *VoicemailgreetingsPostParams {
	var ()
	return &VoicemailgreetingsPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVoicemailgreetingsPostParamsWithTimeout creates a new VoicemailgreetingsPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVoicemailgreetingsPostParamsWithTimeout(timeout time.Duration) *VoicemailgreetingsPostParams {
	var ()
	return &VoicemailgreetingsPostParams{

		timeout: timeout,
	}
}

// NewVoicemailgreetingsPostParamsWithContext creates a new VoicemailgreetingsPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewVoicemailgreetingsPostParamsWithContext(ctx context.Context) *VoicemailgreetingsPostParams {
	var ()
	return &VoicemailgreetingsPostParams{

		Context: ctx,
	}
}

// NewVoicemailgreetingsPostParamsWithHTTPClient creates a new VoicemailgreetingsPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVoicemailgreetingsPostParamsWithHTTPClient(client *http.Client) *VoicemailgreetingsPostParams {
	var ()
	return &VoicemailgreetingsPostParams{
		HTTPClient: client,
	}
}

/*VoicemailgreetingsPostParams contains all the parameters to send to the API endpoint
for the voicemailgreetings post operation typically these are written to a http.Request
*/
type VoicemailgreetingsPostParams struct {

	/*Body*/
	Body *models.CreateanewVoicemailGreetingRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the voicemailgreetings post params
func (o *VoicemailgreetingsPostParams) WithTimeout(timeout time.Duration) *VoicemailgreetingsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the voicemailgreetings post params
func (o *VoicemailgreetingsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the voicemailgreetings post params
func (o *VoicemailgreetingsPostParams) WithContext(ctx context.Context) *VoicemailgreetingsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the voicemailgreetings post params
func (o *VoicemailgreetingsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the voicemailgreetings post params
func (o *VoicemailgreetingsPostParams) WithHTTPClient(client *http.Client) *VoicemailgreetingsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the voicemailgreetings post params
func (o *VoicemailgreetingsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the voicemailgreetings post params
func (o *VoicemailgreetingsPostParams) WithBody(body *models.CreateanewVoicemailGreetingRequest) *VoicemailgreetingsPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the voicemailgreetings post params
func (o *VoicemailgreetingsPostParams) SetBody(body *models.CreateanewVoicemailGreetingRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the voicemailgreetings post params
func (o *VoicemailgreetingsPostParams) WithContentType(contentType string) *VoicemailgreetingsPostParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the voicemailgreetings post params
func (o *VoicemailgreetingsPostParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *VoicemailgreetingsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// header param Content-Type
	if err := r.SetHeaderParam("Content-Type", o.ContentType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
