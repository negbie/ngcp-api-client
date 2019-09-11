// Code generated by go-swagger; DO NOT EDIT.

package capabilities

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

// NewCapabilitiesByIDGetParams creates a new CapabilitiesByIDGetParams object
// with the default values initialized.
func NewCapabilitiesByIDGetParams() *CapabilitiesByIDGetParams {
	var ()
	return &CapabilitiesByIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCapabilitiesByIDGetParamsWithTimeout creates a new CapabilitiesByIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCapabilitiesByIDGetParamsWithTimeout(timeout time.Duration) *CapabilitiesByIDGetParams {
	var ()
	return &CapabilitiesByIDGetParams{

		timeout: timeout,
	}
}

// NewCapabilitiesByIDGetParamsWithContext creates a new CapabilitiesByIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCapabilitiesByIDGetParamsWithContext(ctx context.Context) *CapabilitiesByIDGetParams {
	var ()
	return &CapabilitiesByIDGetParams{

		Context: ctx,
	}
}

// NewCapabilitiesByIDGetParamsWithHTTPClient creates a new CapabilitiesByIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCapabilitiesByIDGetParamsWithHTTPClient(client *http.Client) *CapabilitiesByIDGetParams {
	var ()
	return &CapabilitiesByIDGetParams{
		HTTPClient: client,
	}
}

/*CapabilitiesByIDGetParams contains all the parameters to send to the API endpoint
for the capabilities by Id get operation typically these are written to a http.Request
*/
type CapabilitiesByIDGetParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the capabilities by Id get params
func (o *CapabilitiesByIDGetParams) WithTimeout(timeout time.Duration) *CapabilitiesByIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the capabilities by Id get params
func (o *CapabilitiesByIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the capabilities by Id get params
func (o *CapabilitiesByIDGetParams) WithContext(ctx context.Context) *CapabilitiesByIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the capabilities by Id get params
func (o *CapabilitiesByIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the capabilities by Id get params
func (o *CapabilitiesByIDGetParams) WithHTTPClient(client *http.Client) *CapabilitiesByIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the capabilities by Id get params
func (o *CapabilitiesByIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the capabilities by Id get params
func (o *CapabilitiesByIDGetParams) WithID(id string) *CapabilitiesByIDGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the capabilities by Id get params
func (o *CapabilitiesByIDGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CapabilitiesByIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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