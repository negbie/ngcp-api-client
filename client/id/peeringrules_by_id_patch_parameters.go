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

// NewPeeringrulesByIDPatchParams creates a new PeeringrulesByIDPatchParams object
// with the default values initialized.
func NewPeeringrulesByIDPatchParams() *PeeringrulesByIDPatchParams {
	var ()
	return &PeeringrulesByIDPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPeeringrulesByIDPatchParamsWithTimeout creates a new PeeringrulesByIDPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPeeringrulesByIDPatchParamsWithTimeout(timeout time.Duration) *PeeringrulesByIDPatchParams {
	var ()
	return &PeeringrulesByIDPatchParams{

		timeout: timeout,
	}
}

// NewPeeringrulesByIDPatchParamsWithContext creates a new PeeringrulesByIDPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPeeringrulesByIDPatchParamsWithContext(ctx context.Context) *PeeringrulesByIDPatchParams {
	var ()
	return &PeeringrulesByIDPatchParams{

		Context: ctx,
	}
}

// NewPeeringrulesByIDPatchParamsWithHTTPClient creates a new PeeringrulesByIDPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPeeringrulesByIDPatchParamsWithHTTPClient(client *http.Client) *PeeringrulesByIDPatchParams {
	var ()
	return &PeeringrulesByIDPatchParams{
		HTTPClient: client,
	}
}

/*PeeringrulesByIDPatchParams contains all the parameters to send to the API endpoint
for the peeringrules by Id patch operation typically these are written to a http.Request
*/
type PeeringrulesByIDPatchParams struct {

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

// WithTimeout adds the timeout to the peeringrules by Id patch params
func (o *PeeringrulesByIDPatchParams) WithTimeout(timeout time.Duration) *PeeringrulesByIDPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the peeringrules by Id patch params
func (o *PeeringrulesByIDPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the peeringrules by Id patch params
func (o *PeeringrulesByIDPatchParams) WithContext(ctx context.Context) *PeeringrulesByIDPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the peeringrules by Id patch params
func (o *PeeringrulesByIDPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the peeringrules by Id patch params
func (o *PeeringrulesByIDPatchParams) WithHTTPClient(client *http.Client) *PeeringrulesByIDPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the peeringrules by Id patch params
func (o *PeeringrulesByIDPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the peeringrules by Id patch params
func (o *PeeringrulesByIDPatchParams) WithBody(body string) *PeeringrulesByIDPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the peeringrules by Id patch params
func (o *PeeringrulesByIDPatchParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the peeringrules by Id patch params
func (o *PeeringrulesByIDPatchParams) WithContentType(contentType string) *PeeringrulesByIDPatchParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the peeringrules by Id patch params
func (o *PeeringrulesByIDPatchParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the peeringrules by Id patch params
func (o *PeeringrulesByIDPatchParams) WithID(id string) *PeeringrulesByIDPatchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the peeringrules by Id patch params
func (o *PeeringrulesByIDPatchParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PeeringrulesByIDPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
