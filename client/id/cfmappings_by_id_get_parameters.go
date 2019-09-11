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

// NewCfmappingsByIDGetParams creates a new CfmappingsByIDGetParams object
// with the default values initialized.
func NewCfmappingsByIDGetParams() *CfmappingsByIDGetParams {
	var ()
	return &CfmappingsByIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCfmappingsByIDGetParamsWithTimeout creates a new CfmappingsByIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCfmappingsByIDGetParamsWithTimeout(timeout time.Duration) *CfmappingsByIDGetParams {
	var ()
	return &CfmappingsByIDGetParams{

		timeout: timeout,
	}
}

// NewCfmappingsByIDGetParamsWithContext creates a new CfmappingsByIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCfmappingsByIDGetParamsWithContext(ctx context.Context) *CfmappingsByIDGetParams {
	var ()
	return &CfmappingsByIDGetParams{

		Context: ctx,
	}
}

// NewCfmappingsByIDGetParamsWithHTTPClient creates a new CfmappingsByIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCfmappingsByIDGetParamsWithHTTPClient(client *http.Client) *CfmappingsByIDGetParams {
	var ()
	return &CfmappingsByIDGetParams{
		HTTPClient: client,
	}
}

/*CfmappingsByIDGetParams contains all the parameters to send to the API endpoint
for the cfmappings by Id get operation typically these are written to a http.Request
*/
type CfmappingsByIDGetParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cfmappings by Id get params
func (o *CfmappingsByIDGetParams) WithTimeout(timeout time.Duration) *CfmappingsByIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cfmappings by Id get params
func (o *CfmappingsByIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cfmappings by Id get params
func (o *CfmappingsByIDGetParams) WithContext(ctx context.Context) *CfmappingsByIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cfmappings by Id get params
func (o *CfmappingsByIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cfmappings by Id get params
func (o *CfmappingsByIDGetParams) WithHTTPClient(client *http.Client) *CfmappingsByIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cfmappings by Id get params
func (o *CfmappingsByIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the cfmappings by Id get params
func (o *CfmappingsByIDGetParams) WithID(id string) *CfmappingsByIDGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the cfmappings by Id get params
func (o *CfmappingsByIDGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CfmappingsByIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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