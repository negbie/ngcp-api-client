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

// NewPeeringgroupsByIDGetParams creates a new PeeringgroupsByIDGetParams object
// with the default values initialized.
func NewPeeringgroupsByIDGetParams() *PeeringgroupsByIDGetParams {
	var ()
	return &PeeringgroupsByIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPeeringgroupsByIDGetParamsWithTimeout creates a new PeeringgroupsByIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPeeringgroupsByIDGetParamsWithTimeout(timeout time.Duration) *PeeringgroupsByIDGetParams {
	var ()
	return &PeeringgroupsByIDGetParams{

		timeout: timeout,
	}
}

// NewPeeringgroupsByIDGetParamsWithContext creates a new PeeringgroupsByIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPeeringgroupsByIDGetParamsWithContext(ctx context.Context) *PeeringgroupsByIDGetParams {
	var ()
	return &PeeringgroupsByIDGetParams{

		Context: ctx,
	}
}

// NewPeeringgroupsByIDGetParamsWithHTTPClient creates a new PeeringgroupsByIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPeeringgroupsByIDGetParamsWithHTTPClient(client *http.Client) *PeeringgroupsByIDGetParams {
	var ()
	return &PeeringgroupsByIDGetParams{
		HTTPClient: client,
	}
}

/*PeeringgroupsByIDGetParams contains all the parameters to send to the API endpoint
for the peeringgroups by Id get operation typically these are written to a http.Request
*/
type PeeringgroupsByIDGetParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the peeringgroups by Id get params
func (o *PeeringgroupsByIDGetParams) WithTimeout(timeout time.Duration) *PeeringgroupsByIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the peeringgroups by Id get params
func (o *PeeringgroupsByIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the peeringgroups by Id get params
func (o *PeeringgroupsByIDGetParams) WithContext(ctx context.Context) *PeeringgroupsByIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the peeringgroups by Id get params
func (o *PeeringgroupsByIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the peeringgroups by Id get params
func (o *PeeringgroupsByIDGetParams) WithHTTPClient(client *http.Client) *PeeringgroupsByIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the peeringgroups by Id get params
func (o *PeeringgroupsByIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the peeringgroups by Id get params
func (o *PeeringgroupsByIDGetParams) WithID(id string) *PeeringgroupsByIDGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the peeringgroups by Id get params
func (o *PeeringgroupsByIDGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PeeringgroupsByIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
