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

// NewDomainsByIDDeleteParams creates a new DomainsByIDDeleteParams object
// with the default values initialized.
func NewDomainsByIDDeleteParams() *DomainsByIDDeleteParams {
	var ()
	return &DomainsByIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDomainsByIDDeleteParamsWithTimeout creates a new DomainsByIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDomainsByIDDeleteParamsWithTimeout(timeout time.Duration) *DomainsByIDDeleteParams {
	var ()
	return &DomainsByIDDeleteParams{

		timeout: timeout,
	}
}

// NewDomainsByIDDeleteParamsWithContext creates a new DomainsByIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDomainsByIDDeleteParamsWithContext(ctx context.Context) *DomainsByIDDeleteParams {
	var ()
	return &DomainsByIDDeleteParams{

		Context: ctx,
	}
}

// NewDomainsByIDDeleteParamsWithHTTPClient creates a new DomainsByIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDomainsByIDDeleteParamsWithHTTPClient(client *http.Client) *DomainsByIDDeleteParams {
	var ()
	return &DomainsByIDDeleteParams{
		HTTPClient: client,
	}
}

/*DomainsByIDDeleteParams contains all the parameters to send to the API endpoint
for the domains by Id delete operation typically these are written to a http.Request
*/
type DomainsByIDDeleteParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the domains by Id delete params
func (o *DomainsByIDDeleteParams) WithTimeout(timeout time.Duration) *DomainsByIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domains by Id delete params
func (o *DomainsByIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domains by Id delete params
func (o *DomainsByIDDeleteParams) WithContext(ctx context.Context) *DomainsByIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domains by Id delete params
func (o *DomainsByIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domains by Id delete params
func (o *DomainsByIDDeleteParams) WithHTTPClient(client *http.Client) *DomainsByIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domains by Id delete params
func (o *DomainsByIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the domains by Id delete params
func (o *DomainsByIDDeleteParams) WithID(id string) *DomainsByIDDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the domains by Id delete params
func (o *DomainsByIDDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DomainsByIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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