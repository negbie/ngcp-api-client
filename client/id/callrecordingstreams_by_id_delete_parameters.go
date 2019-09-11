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

// NewCallrecordingstreamsByIDDeleteParams creates a new CallrecordingstreamsByIDDeleteParams object
// with the default values initialized.
func NewCallrecordingstreamsByIDDeleteParams() *CallrecordingstreamsByIDDeleteParams {
	var ()
	return &CallrecordingstreamsByIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCallrecordingstreamsByIDDeleteParamsWithTimeout creates a new CallrecordingstreamsByIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCallrecordingstreamsByIDDeleteParamsWithTimeout(timeout time.Duration) *CallrecordingstreamsByIDDeleteParams {
	var ()
	return &CallrecordingstreamsByIDDeleteParams{

		timeout: timeout,
	}
}

// NewCallrecordingstreamsByIDDeleteParamsWithContext creates a new CallrecordingstreamsByIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewCallrecordingstreamsByIDDeleteParamsWithContext(ctx context.Context) *CallrecordingstreamsByIDDeleteParams {
	var ()
	return &CallrecordingstreamsByIDDeleteParams{

		Context: ctx,
	}
}

// NewCallrecordingstreamsByIDDeleteParamsWithHTTPClient creates a new CallrecordingstreamsByIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCallrecordingstreamsByIDDeleteParamsWithHTTPClient(client *http.Client) *CallrecordingstreamsByIDDeleteParams {
	var ()
	return &CallrecordingstreamsByIDDeleteParams{
		HTTPClient: client,
	}
}

/*CallrecordingstreamsByIDDeleteParams contains all the parameters to send to the API endpoint
for the callrecordingstreams by Id delete operation typically these are written to a http.Request
*/
type CallrecordingstreamsByIDDeleteParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the callrecordingstreams by Id delete params
func (o *CallrecordingstreamsByIDDeleteParams) WithTimeout(timeout time.Duration) *CallrecordingstreamsByIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the callrecordingstreams by Id delete params
func (o *CallrecordingstreamsByIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the callrecordingstreams by Id delete params
func (o *CallrecordingstreamsByIDDeleteParams) WithContext(ctx context.Context) *CallrecordingstreamsByIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the callrecordingstreams by Id delete params
func (o *CallrecordingstreamsByIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the callrecordingstreams by Id delete params
func (o *CallrecordingstreamsByIDDeleteParams) WithHTTPClient(client *http.Client) *CallrecordingstreamsByIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the callrecordingstreams by Id delete params
func (o *CallrecordingstreamsByIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the callrecordingstreams by Id delete params
func (o *CallrecordingstreamsByIDDeleteParams) WithID(id string) *CallrecordingstreamsByIDDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the callrecordingstreams by Id delete params
func (o *CallrecordingstreamsByIDDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CallrecordingstreamsByIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
