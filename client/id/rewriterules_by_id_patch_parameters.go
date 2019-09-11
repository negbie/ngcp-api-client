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

// NewRewriterulesByIDPatchParams creates a new RewriterulesByIDPatchParams object
// with the default values initialized.
func NewRewriterulesByIDPatchParams() *RewriterulesByIDPatchParams {
	var ()
	return &RewriterulesByIDPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRewriterulesByIDPatchParamsWithTimeout creates a new RewriterulesByIDPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRewriterulesByIDPatchParamsWithTimeout(timeout time.Duration) *RewriterulesByIDPatchParams {
	var ()
	return &RewriterulesByIDPatchParams{

		timeout: timeout,
	}
}

// NewRewriterulesByIDPatchParamsWithContext creates a new RewriterulesByIDPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewRewriterulesByIDPatchParamsWithContext(ctx context.Context) *RewriterulesByIDPatchParams {
	var ()
	return &RewriterulesByIDPatchParams{

		Context: ctx,
	}
}

// NewRewriterulesByIDPatchParamsWithHTTPClient creates a new RewriterulesByIDPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRewriterulesByIDPatchParamsWithHTTPClient(client *http.Client) *RewriterulesByIDPatchParams {
	var ()
	return &RewriterulesByIDPatchParams{
		HTTPClient: client,
	}
}

/*RewriterulesByIDPatchParams contains all the parameters to send to the API endpoint
for the rewriterules by Id patch operation typically these are written to a http.Request
*/
type RewriterulesByIDPatchParams struct {

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

// WithTimeout adds the timeout to the rewriterules by Id patch params
func (o *RewriterulesByIDPatchParams) WithTimeout(timeout time.Duration) *RewriterulesByIDPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rewriterules by Id patch params
func (o *RewriterulesByIDPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rewriterules by Id patch params
func (o *RewriterulesByIDPatchParams) WithContext(ctx context.Context) *RewriterulesByIDPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rewriterules by Id patch params
func (o *RewriterulesByIDPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rewriterules by Id patch params
func (o *RewriterulesByIDPatchParams) WithHTTPClient(client *http.Client) *RewriterulesByIDPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rewriterules by Id patch params
func (o *RewriterulesByIDPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the rewriterules by Id patch params
func (o *RewriterulesByIDPatchParams) WithBody(body string) *RewriterulesByIDPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the rewriterules by Id patch params
func (o *RewriterulesByIDPatchParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the rewriterules by Id patch params
func (o *RewriterulesByIDPatchParams) WithContentType(contentType string) *RewriterulesByIDPatchParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the rewriterules by Id patch params
func (o *RewriterulesByIDPatchParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the rewriterules by Id patch params
func (o *RewriterulesByIDPatchParams) WithID(id string) *RewriterulesByIDPatchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the rewriterules by Id patch params
func (o *RewriterulesByIDPatchParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RewriterulesByIDPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
