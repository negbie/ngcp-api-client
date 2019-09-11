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

// NewDomainsByIDGetParams creates a new DomainsByIDGetParams object
// with the default values initialized.
func NewDomainsByIDGetParams() *DomainsByIDGetParams {
	var ()
	return &DomainsByIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDomainsByIDGetParamsWithTimeout creates a new DomainsByIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDomainsByIDGetParamsWithTimeout(timeout time.Duration) *DomainsByIDGetParams {
	var ()
	return &DomainsByIDGetParams{

		timeout: timeout,
	}
}

// NewDomainsByIDGetParamsWithContext creates a new DomainsByIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewDomainsByIDGetParamsWithContext(ctx context.Context) *DomainsByIDGetParams {
	var ()
	return &DomainsByIDGetParams{

		Context: ctx,
	}
}

// NewDomainsByIDGetParamsWithHTTPClient creates a new DomainsByIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDomainsByIDGetParamsWithHTTPClient(client *http.Client) *DomainsByIDGetParams {
	var ()
	return &DomainsByIDGetParams{
		HTTPClient: client,
	}
}

/*DomainsByIDGetParams contains all the parameters to send to the API endpoint
for the domains by Id get operation typically these are written to a http.Request
*/
type DomainsByIDGetParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the domains by Id get params
func (o *DomainsByIDGetParams) WithTimeout(timeout time.Duration) *DomainsByIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domains by Id get params
func (o *DomainsByIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domains by Id get params
func (o *DomainsByIDGetParams) WithContext(ctx context.Context) *DomainsByIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domains by Id get params
func (o *DomainsByIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domains by Id get params
func (o *DomainsByIDGetParams) WithHTTPClient(client *http.Client) *DomainsByIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domains by Id get params
func (o *DomainsByIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the domains by Id get params
func (o *DomainsByIDGetParams) WithID(id string) *DomainsByIDGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the domains by Id get params
func (o *DomainsByIDGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DomainsByIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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