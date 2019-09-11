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

// NewEmergencymappingcontainersByIDDeleteParams creates a new EmergencymappingcontainersByIDDeleteParams object
// with the default values initialized.
func NewEmergencymappingcontainersByIDDeleteParams() *EmergencymappingcontainersByIDDeleteParams {
	var ()
	return &EmergencymappingcontainersByIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmergencymappingcontainersByIDDeleteParamsWithTimeout creates a new EmergencymappingcontainersByIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmergencymappingcontainersByIDDeleteParamsWithTimeout(timeout time.Duration) *EmergencymappingcontainersByIDDeleteParams {
	var ()
	return &EmergencymappingcontainersByIDDeleteParams{

		timeout: timeout,
	}
}

// NewEmergencymappingcontainersByIDDeleteParamsWithContext creates a new EmergencymappingcontainersByIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmergencymappingcontainersByIDDeleteParamsWithContext(ctx context.Context) *EmergencymappingcontainersByIDDeleteParams {
	var ()
	return &EmergencymappingcontainersByIDDeleteParams{

		Context: ctx,
	}
}

// NewEmergencymappingcontainersByIDDeleteParamsWithHTTPClient creates a new EmergencymappingcontainersByIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmergencymappingcontainersByIDDeleteParamsWithHTTPClient(client *http.Client) *EmergencymappingcontainersByIDDeleteParams {
	var ()
	return &EmergencymappingcontainersByIDDeleteParams{
		HTTPClient: client,
	}
}

/*EmergencymappingcontainersByIDDeleteParams contains all the parameters to send to the API endpoint
for the emergencymappingcontainers by Id delete operation typically these are written to a http.Request
*/
type EmergencymappingcontainersByIDDeleteParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the emergencymappingcontainers by Id delete params
func (o *EmergencymappingcontainersByIDDeleteParams) WithTimeout(timeout time.Duration) *EmergencymappingcontainersByIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the emergencymappingcontainers by Id delete params
func (o *EmergencymappingcontainersByIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the emergencymappingcontainers by Id delete params
func (o *EmergencymappingcontainersByIDDeleteParams) WithContext(ctx context.Context) *EmergencymappingcontainersByIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the emergencymappingcontainers by Id delete params
func (o *EmergencymappingcontainersByIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the emergencymappingcontainers by Id delete params
func (o *EmergencymappingcontainersByIDDeleteParams) WithHTTPClient(client *http.Client) *EmergencymappingcontainersByIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the emergencymappingcontainers by Id delete params
func (o *EmergencymappingcontainersByIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the emergencymappingcontainers by Id delete params
func (o *EmergencymappingcontainersByIDDeleteParams) WithID(id string) *EmergencymappingcontainersByIDDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the emergencymappingcontainers by Id delete params
func (o *EmergencymappingcontainersByIDDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EmergencymappingcontainersByIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
