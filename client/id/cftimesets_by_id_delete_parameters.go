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

// NewCftimesetsByIDDeleteParams creates a new CftimesetsByIDDeleteParams object
// with the default values initialized.
func NewCftimesetsByIDDeleteParams() *CftimesetsByIDDeleteParams {
	var ()
	return &CftimesetsByIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCftimesetsByIDDeleteParamsWithTimeout creates a new CftimesetsByIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCftimesetsByIDDeleteParamsWithTimeout(timeout time.Duration) *CftimesetsByIDDeleteParams {
	var ()
	return &CftimesetsByIDDeleteParams{

		timeout: timeout,
	}
}

// NewCftimesetsByIDDeleteParamsWithContext creates a new CftimesetsByIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewCftimesetsByIDDeleteParamsWithContext(ctx context.Context) *CftimesetsByIDDeleteParams {
	var ()
	return &CftimesetsByIDDeleteParams{

		Context: ctx,
	}
}

// NewCftimesetsByIDDeleteParamsWithHTTPClient creates a new CftimesetsByIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCftimesetsByIDDeleteParamsWithHTTPClient(client *http.Client) *CftimesetsByIDDeleteParams {
	var ()
	return &CftimesetsByIDDeleteParams{
		HTTPClient: client,
	}
}

/*CftimesetsByIDDeleteParams contains all the parameters to send to the API endpoint
for the cftimesets by Id delete operation typically these are written to a http.Request
*/
type CftimesetsByIDDeleteParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cftimesets by Id delete params
func (o *CftimesetsByIDDeleteParams) WithTimeout(timeout time.Duration) *CftimesetsByIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cftimesets by Id delete params
func (o *CftimesetsByIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cftimesets by Id delete params
func (o *CftimesetsByIDDeleteParams) WithContext(ctx context.Context) *CftimesetsByIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cftimesets by Id delete params
func (o *CftimesetsByIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cftimesets by Id delete params
func (o *CftimesetsByIDDeleteParams) WithHTTPClient(client *http.Client) *CftimesetsByIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cftimesets by Id delete params
func (o *CftimesetsByIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the cftimesets by Id delete params
func (o *CftimesetsByIDDeleteParams) WithID(id string) *CftimesetsByIDDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the cftimesets by Id delete params
func (o *CftimesetsByIDDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CftimesetsByIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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