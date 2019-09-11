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

// NewLnpcarriersByIDPutParams creates a new LnpcarriersByIDPutParams object
// with the default values initialized.
func NewLnpcarriersByIDPutParams() *LnpcarriersByIDPutParams {
	var ()
	return &LnpcarriersByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLnpcarriersByIDPutParamsWithTimeout creates a new LnpcarriersByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLnpcarriersByIDPutParamsWithTimeout(timeout time.Duration) *LnpcarriersByIDPutParams {
	var ()
	return &LnpcarriersByIDPutParams{

		timeout: timeout,
	}
}

// NewLnpcarriersByIDPutParamsWithContext creates a new LnpcarriersByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewLnpcarriersByIDPutParamsWithContext(ctx context.Context) *LnpcarriersByIDPutParams {
	var ()
	return &LnpcarriersByIDPutParams{

		Context: ctx,
	}
}

// NewLnpcarriersByIDPutParamsWithHTTPClient creates a new LnpcarriersByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLnpcarriersByIDPutParamsWithHTTPClient(client *http.Client) *LnpcarriersByIDPutParams {
	var ()
	return &LnpcarriersByIDPutParams{
		HTTPClient: client,
	}
}

/*LnpcarriersByIDPutParams contains all the parameters to send to the API endpoint
for the lnpcarriers by Id put operation typically these are written to a http.Request
*/
type LnpcarriersByIDPutParams struct {

	/*Body*/
	Body *Replace1changeaspecificLnpCarrierRequest
	/*ContentType*/
	ContentType string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the lnpcarriers by Id put params
func (o *LnpcarriersByIDPutParams) WithTimeout(timeout time.Duration) *LnpcarriersByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lnpcarriers by Id put params
func (o *LnpcarriersByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lnpcarriers by Id put params
func (o *LnpcarriersByIDPutParams) WithContext(ctx context.Context) *LnpcarriersByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lnpcarriers by Id put params
func (o *LnpcarriersByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lnpcarriers by Id put params
func (o *LnpcarriersByIDPutParams) WithHTTPClient(client *http.Client) *LnpcarriersByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lnpcarriers by Id put params
func (o *LnpcarriersByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the lnpcarriers by Id put params
func (o *LnpcarriersByIDPutParams) WithBody(body *Replace1changeaspecificLnpCarrierRequest) *LnpcarriersByIDPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the lnpcarriers by Id put params
func (o *LnpcarriersByIDPutParams) SetBody(body *Replace1changeaspecificLnpCarrierRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the lnpcarriers by Id put params
func (o *LnpcarriersByIDPutParams) WithContentType(contentType string) *LnpcarriersByIDPutParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the lnpcarriers by Id put params
func (o *LnpcarriersByIDPutParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the lnpcarriers by Id put params
func (o *LnpcarriersByIDPutParams) WithID(id string) *LnpcarriersByIDPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the lnpcarriers by Id put params
func (o *LnpcarriersByIDPutParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LnpcarriersByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// header param Content-Type
	if err := r.SetHeaderParam("Content-Type", o.ContentType); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
