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

// NewNcoslnpcarriersByIDGetParams creates a new NcoslnpcarriersByIDGetParams object
// with the default values initialized.
func NewNcoslnpcarriersByIDGetParams() *NcoslnpcarriersByIDGetParams {
	var ()
	return &NcoslnpcarriersByIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNcoslnpcarriersByIDGetParamsWithTimeout creates a new NcoslnpcarriersByIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNcoslnpcarriersByIDGetParamsWithTimeout(timeout time.Duration) *NcoslnpcarriersByIDGetParams {
	var ()
	return &NcoslnpcarriersByIDGetParams{

		timeout: timeout,
	}
}

// NewNcoslnpcarriersByIDGetParamsWithContext creates a new NcoslnpcarriersByIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewNcoslnpcarriersByIDGetParamsWithContext(ctx context.Context) *NcoslnpcarriersByIDGetParams {
	var ()
	return &NcoslnpcarriersByIDGetParams{

		Context: ctx,
	}
}

// NewNcoslnpcarriersByIDGetParamsWithHTTPClient creates a new NcoslnpcarriersByIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNcoslnpcarriersByIDGetParamsWithHTTPClient(client *http.Client) *NcoslnpcarriersByIDGetParams {
	var ()
	return &NcoslnpcarriersByIDGetParams{
		HTTPClient: client,
	}
}

/*NcoslnpcarriersByIDGetParams contains all the parameters to send to the API endpoint
for the ncoslnpcarriers by Id get operation typically these are written to a http.Request
*/
type NcoslnpcarriersByIDGetParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ncoslnpcarriers by Id get params
func (o *NcoslnpcarriersByIDGetParams) WithTimeout(timeout time.Duration) *NcoslnpcarriersByIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ncoslnpcarriers by Id get params
func (o *NcoslnpcarriersByIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ncoslnpcarriers by Id get params
func (o *NcoslnpcarriersByIDGetParams) WithContext(ctx context.Context) *NcoslnpcarriersByIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ncoslnpcarriers by Id get params
func (o *NcoslnpcarriersByIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ncoslnpcarriers by Id get params
func (o *NcoslnpcarriersByIDGetParams) WithHTTPClient(client *http.Client) *NcoslnpcarriersByIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ncoslnpcarriers by Id get params
func (o *NcoslnpcarriersByIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ncoslnpcarriers by Id get params
func (o *NcoslnpcarriersByIDGetParams) WithID(id string) *NcoslnpcarriersByIDGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ncoslnpcarriers by Id get params
func (o *NcoslnpcarriersByIDGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *NcoslnpcarriersByIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
