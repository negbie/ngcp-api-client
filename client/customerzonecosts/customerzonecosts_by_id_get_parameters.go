// Code generated by go-swagger; DO NOT EDIT.

package customerzonecosts

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

// NewCustomerzonecostsByIDGetParams creates a new CustomerzonecostsByIDGetParams object
// with the default values initialized.
func NewCustomerzonecostsByIDGetParams() *CustomerzonecostsByIDGetParams {
	var ()
	return &CustomerzonecostsByIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerzonecostsByIDGetParamsWithTimeout creates a new CustomerzonecostsByIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerzonecostsByIDGetParamsWithTimeout(timeout time.Duration) *CustomerzonecostsByIDGetParams {
	var ()
	return &CustomerzonecostsByIDGetParams{

		timeout: timeout,
	}
}

// NewCustomerzonecostsByIDGetParamsWithContext creates a new CustomerzonecostsByIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerzonecostsByIDGetParamsWithContext(ctx context.Context) *CustomerzonecostsByIDGetParams {
	var ()
	return &CustomerzonecostsByIDGetParams{

		Context: ctx,
	}
}

// NewCustomerzonecostsByIDGetParamsWithHTTPClient creates a new CustomerzonecostsByIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerzonecostsByIDGetParamsWithHTTPClient(client *http.Client) *CustomerzonecostsByIDGetParams {
	var ()
	return &CustomerzonecostsByIDGetParams{
		HTTPClient: client,
	}
}

/*CustomerzonecostsByIDGetParams contains all the parameters to send to the API endpoint
for the customerzonecosts by Id get operation typically these are written to a http.Request
*/
type CustomerzonecostsByIDGetParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customerzonecosts by Id get params
func (o *CustomerzonecostsByIDGetParams) WithTimeout(timeout time.Duration) *CustomerzonecostsByIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customerzonecosts by Id get params
func (o *CustomerzonecostsByIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customerzonecosts by Id get params
func (o *CustomerzonecostsByIDGetParams) WithContext(ctx context.Context) *CustomerzonecostsByIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customerzonecosts by Id get params
func (o *CustomerzonecostsByIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customerzonecosts by Id get params
func (o *CustomerzonecostsByIDGetParams) WithHTTPClient(client *http.Client) *CustomerzonecostsByIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customerzonecosts by Id get params
func (o *CustomerzonecostsByIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customerzonecosts by Id get params
func (o *CustomerzonecostsByIDGetParams) WithID(id string) *CustomerzonecostsByIDGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customerzonecosts by Id get params
func (o *CustomerzonecostsByIDGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerzonecostsByIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
