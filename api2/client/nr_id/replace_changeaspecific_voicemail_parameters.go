// Code generated by go-swagger; DO NOT EDIT.

package nr_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewReplaceChangeaspecificVoicemailParams creates a new ReplaceChangeaspecificVoicemailParams object
// with the default values initialized.
func NewReplaceChangeaspecificVoicemailParams() *ReplaceChangeaspecificVoicemailParams {
	var ()
	return &ReplaceChangeaspecificVoicemailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceChangeaspecificVoicemailParamsWithTimeout creates a new ReplaceChangeaspecificVoicemailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceChangeaspecificVoicemailParamsWithTimeout(timeout time.Duration) *ReplaceChangeaspecificVoicemailParams {
	var ()
	return &ReplaceChangeaspecificVoicemailParams{

		timeout: timeout,
	}
}

// NewReplaceChangeaspecificVoicemailParamsWithContext creates a new ReplaceChangeaspecificVoicemailParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceChangeaspecificVoicemailParamsWithContext(ctx context.Context) *ReplaceChangeaspecificVoicemailParams {
	var ()
	return &ReplaceChangeaspecificVoicemailParams{

		Context: ctx,
	}
}

// NewReplaceChangeaspecificVoicemailParamsWithHTTPClient creates a new ReplaceChangeaspecificVoicemailParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceChangeaspecificVoicemailParamsWithHTTPClient(client *http.Client) *ReplaceChangeaspecificVoicemailParams {
	var ()
	return &ReplaceChangeaspecificVoicemailParams{
		HTTPClient: client,
	}
}

/*ReplaceChangeaspecificVoicemailParams contains all the parameters to send to the API endpoint
for the replace changeaspecific voicemail operation typically these are written to a http.Request
*/
type ReplaceChangeaspecificVoicemailParams struct {

	/*Body*/
	Body *Replace1changeaspecificVoicemailRequest
	/*ContentType*/
	ContentType string
	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace changeaspecific voicemail params
func (o *ReplaceChangeaspecificVoicemailParams) WithTimeout(timeout time.Duration) *ReplaceChangeaspecificVoicemailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace changeaspecific voicemail params
func (o *ReplaceChangeaspecificVoicemailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace changeaspecific voicemail params
func (o *ReplaceChangeaspecificVoicemailParams) WithContext(ctx context.Context) *ReplaceChangeaspecificVoicemailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace changeaspecific voicemail params
func (o *ReplaceChangeaspecificVoicemailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace changeaspecific voicemail params
func (o *ReplaceChangeaspecificVoicemailParams) WithHTTPClient(client *http.Client) *ReplaceChangeaspecificVoicemailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace changeaspecific voicemail params
func (o *ReplaceChangeaspecificVoicemailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace changeaspecific voicemail params
func (o *ReplaceChangeaspecificVoicemailParams) WithBody(body *Replace1changeaspecificVoicemailRequest) *ReplaceChangeaspecificVoicemailParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace changeaspecific voicemail params
func (o *ReplaceChangeaspecificVoicemailParams) SetBody(body *Replace1changeaspecificVoicemailRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the replace changeaspecific voicemail params
func (o *ReplaceChangeaspecificVoicemailParams) WithContentType(contentType string) *ReplaceChangeaspecificVoicemailParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the replace changeaspecific voicemail params
func (o *ReplaceChangeaspecificVoicemailParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the replace changeaspecific voicemail params
func (o *ReplaceChangeaspecificVoicemailParams) WithID(id string) *ReplaceChangeaspecificVoicemailParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace changeaspecific voicemail params
func (o *ReplaceChangeaspecificVoicemailParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceChangeaspecificVoicemailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}