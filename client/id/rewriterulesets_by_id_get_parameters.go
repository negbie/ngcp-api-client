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

// NewRewriterulesetsByIDGetParams creates a new RewriterulesetsByIDGetParams object
// with the default values initialized.
func NewRewriterulesetsByIDGetParams() *RewriterulesetsByIDGetParams {
	var ()
	return &RewriterulesetsByIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRewriterulesetsByIDGetParamsWithTimeout creates a new RewriterulesetsByIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRewriterulesetsByIDGetParamsWithTimeout(timeout time.Duration) *RewriterulesetsByIDGetParams {
	var ()
	return &RewriterulesetsByIDGetParams{

		timeout: timeout,
	}
}

// NewRewriterulesetsByIDGetParamsWithContext creates a new RewriterulesetsByIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewRewriterulesetsByIDGetParamsWithContext(ctx context.Context) *RewriterulesetsByIDGetParams {
	var ()
	return &RewriterulesetsByIDGetParams{

		Context: ctx,
	}
}

// NewRewriterulesetsByIDGetParamsWithHTTPClient creates a new RewriterulesetsByIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRewriterulesetsByIDGetParamsWithHTTPClient(client *http.Client) *RewriterulesetsByIDGetParams {
	var ()
	return &RewriterulesetsByIDGetParams{
		HTTPClient: client,
	}
}

/*RewriterulesetsByIDGetParams contains all the parameters to send to the API endpoint
for the rewriterulesets by Id get operation typically these are written to a http.Request
*/
type RewriterulesetsByIDGetParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the rewriterulesets by Id get params
func (o *RewriterulesetsByIDGetParams) WithTimeout(timeout time.Duration) *RewriterulesetsByIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rewriterulesets by Id get params
func (o *RewriterulesetsByIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rewriterulesets by Id get params
func (o *RewriterulesetsByIDGetParams) WithContext(ctx context.Context) *RewriterulesetsByIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rewriterulesets by Id get params
func (o *RewriterulesetsByIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rewriterulesets by Id get params
func (o *RewriterulesetsByIDGetParams) WithHTTPClient(client *http.Client) *RewriterulesetsByIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rewriterulesets by Id get params
func (o *RewriterulesetsByIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the rewriterulesets by Id get params
func (o *RewriterulesetsByIDGetParams) WithID(id string) *RewriterulesetsByIDGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the rewriterulesets by Id get params
func (o *RewriterulesetsByIDGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RewriterulesetsByIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
