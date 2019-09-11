// Code generated by go-swagger; DO NOT EDIT.

package id

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

// NewVoicemailsByIDPatchParams creates a new VoicemailsByIDPatchParams object
// with the default values initialized.
func NewVoicemailsByIDPatchParams() *VoicemailsByIDPatchParams {
	var ()
	return &VoicemailsByIDPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVoicemailsByIDPatchParamsWithTimeout creates a new VoicemailsByIDPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVoicemailsByIDPatchParamsWithTimeout(timeout time.Duration) *VoicemailsByIDPatchParams {
	var ()
	return &VoicemailsByIDPatchParams{

		timeout: timeout,
	}
}

// NewVoicemailsByIDPatchParamsWithContext creates a new VoicemailsByIDPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewVoicemailsByIDPatchParamsWithContext(ctx context.Context) *VoicemailsByIDPatchParams {
	var ()
	return &VoicemailsByIDPatchParams{

		Context: ctx,
	}
}

// NewVoicemailsByIDPatchParamsWithHTTPClient creates a new VoicemailsByIDPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVoicemailsByIDPatchParamsWithHTTPClient(client *http.Client) *VoicemailsByIDPatchParams {
	var ()
	return &VoicemailsByIDPatchParams{
		HTTPClient: client,
	}
}

/*VoicemailsByIDPatchParams contains all the parameters to send to the API endpoint
for the voicemails by Id patch operation typically these are written to a http.Request
*/
type VoicemailsByIDPatchParams struct {

	/*Body*/
	Body string
	/*ContentType*/
	ContentType string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the voicemails by Id patch params
func (o *VoicemailsByIDPatchParams) WithTimeout(timeout time.Duration) *VoicemailsByIDPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the voicemails by Id patch params
func (o *VoicemailsByIDPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the voicemails by Id patch params
func (o *VoicemailsByIDPatchParams) WithContext(ctx context.Context) *VoicemailsByIDPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the voicemails by Id patch params
func (o *VoicemailsByIDPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the voicemails by Id patch params
func (o *VoicemailsByIDPatchParams) WithHTTPClient(client *http.Client) *VoicemailsByIDPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the voicemails by Id patch params
func (o *VoicemailsByIDPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the voicemails by Id patch params
func (o *VoicemailsByIDPatchParams) WithBody(body string) *VoicemailsByIDPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the voicemails by Id patch params
func (o *VoicemailsByIDPatchParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the voicemails by Id patch params
func (o *VoicemailsByIDPatchParams) WithContentType(contentType string) *VoicemailsByIDPatchParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the voicemails by Id patch params
func (o *VoicemailsByIDPatchParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the voicemails by Id patch params
func (o *VoicemailsByIDPatchParams) WithID(id string) *VoicemailsByIDPatchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the voicemails by Id patch params
func (o *VoicemailsByIDPatchParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VoicemailsByIDPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// header param Content-Type
	if err := r.SetHeaderParam("Content-Type", o.ContentType); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
