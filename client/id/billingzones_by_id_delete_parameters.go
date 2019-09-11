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

// NewBillingzonesByIDDeleteParams creates a new BillingzonesByIDDeleteParams object
// with the default values initialized.
func NewBillingzonesByIDDeleteParams() *BillingzonesByIDDeleteParams {
	var ()
	return &BillingzonesByIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBillingzonesByIDDeleteParamsWithTimeout creates a new BillingzonesByIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBillingzonesByIDDeleteParamsWithTimeout(timeout time.Duration) *BillingzonesByIDDeleteParams {
	var ()
	return &BillingzonesByIDDeleteParams{

		timeout: timeout,
	}
}

// NewBillingzonesByIDDeleteParamsWithContext creates a new BillingzonesByIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewBillingzonesByIDDeleteParamsWithContext(ctx context.Context) *BillingzonesByIDDeleteParams {
	var ()
	return &BillingzonesByIDDeleteParams{

		Context: ctx,
	}
}

// NewBillingzonesByIDDeleteParamsWithHTTPClient creates a new BillingzonesByIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBillingzonesByIDDeleteParamsWithHTTPClient(client *http.Client) *BillingzonesByIDDeleteParams {
	var ()
	return &BillingzonesByIDDeleteParams{
		HTTPClient: client,
	}
}

/*BillingzonesByIDDeleteParams contains all the parameters to send to the API endpoint
for the billingzones by Id delete operation typically these are written to a http.Request
*/
type BillingzonesByIDDeleteParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the billingzones by Id delete params
func (o *BillingzonesByIDDeleteParams) WithTimeout(timeout time.Duration) *BillingzonesByIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billingzones by Id delete params
func (o *BillingzonesByIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billingzones by Id delete params
func (o *BillingzonesByIDDeleteParams) WithContext(ctx context.Context) *BillingzonesByIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billingzones by Id delete params
func (o *BillingzonesByIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billingzones by Id delete params
func (o *BillingzonesByIDDeleteParams) WithHTTPClient(client *http.Client) *BillingzonesByIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billingzones by Id delete params
func (o *BillingzonesByIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the billingzones by Id delete params
func (o *BillingzonesByIDDeleteParams) WithID(id string) *BillingzonesByIDDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the billingzones by Id delete params
func (o *BillingzonesByIDDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *BillingzonesByIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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