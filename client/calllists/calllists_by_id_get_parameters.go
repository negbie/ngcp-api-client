// Code generated by go-swagger; DO NOT EDIT.

package calllists

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

// NewCalllistsByIDGetParams creates a new CalllistsByIDGetParams object
// with the default values initialized.
func NewCalllistsByIDGetParams() *CalllistsByIDGetParams {
	var ()
	return &CalllistsByIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCalllistsByIDGetParamsWithTimeout creates a new CalllistsByIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCalllistsByIDGetParamsWithTimeout(timeout time.Duration) *CalllistsByIDGetParams {
	var ()
	return &CalllistsByIDGetParams{

		timeout: timeout,
	}
}

// NewCalllistsByIDGetParamsWithContext creates a new CalllistsByIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCalllistsByIDGetParamsWithContext(ctx context.Context) *CalllistsByIDGetParams {
	var ()
	return &CalllistsByIDGetParams{

		Context: ctx,
	}
}

// NewCalllistsByIDGetParamsWithHTTPClient creates a new CalllistsByIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCalllistsByIDGetParamsWithHTTPClient(client *http.Client) *CalllistsByIDGetParams {
	var ()
	return &CalllistsByIDGetParams{
		HTTPClient: client,
	}
}

/*CalllistsByIDGetParams contains all the parameters to send to the API endpoint
for the calllists by Id get operation typically these are written to a http.Request
*/
type CalllistsByIDGetParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the calllists by Id get params
func (o *CalllistsByIDGetParams) WithTimeout(timeout time.Duration) *CalllistsByIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the calllists by Id get params
func (o *CalllistsByIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the calllists by Id get params
func (o *CalllistsByIDGetParams) WithContext(ctx context.Context) *CalllistsByIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the calllists by Id get params
func (o *CalllistsByIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the calllists by Id get params
func (o *CalllistsByIDGetParams) WithHTTPClient(client *http.Client) *CalllistsByIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the calllists by Id get params
func (o *CalllistsByIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the calllists by Id get params
func (o *CalllistsByIDGetParams) WithID(id string) *CalllistsByIDGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the calllists by Id get params
func (o *CalllistsByIDGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CalllistsByIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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