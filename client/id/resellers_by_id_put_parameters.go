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

// NewResellersByIDPutParams creates a new ResellersByIDPutParams object
// with the default values initialized.
func NewResellersByIDPutParams() *ResellersByIDPutParams {
	var ()
	return &ResellersByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewResellersByIDPutParamsWithTimeout creates a new ResellersByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewResellersByIDPutParamsWithTimeout(timeout time.Duration) *ResellersByIDPutParams {
	var ()
	return &ResellersByIDPutParams{

		timeout: timeout,
	}
}

// NewResellersByIDPutParamsWithContext creates a new ResellersByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewResellersByIDPutParamsWithContext(ctx context.Context) *ResellersByIDPutParams {
	var ()
	return &ResellersByIDPutParams{

		Context: ctx,
	}
}

// NewResellersByIDPutParamsWithHTTPClient creates a new ResellersByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResellersByIDPutParamsWithHTTPClient(client *http.Client) *ResellersByIDPutParams {
	var ()
	return &ResellersByIDPutParams{
		HTTPClient: client,
	}
}

/*ResellersByIDPutParams contains all the parameters to send to the API endpoint
for the resellers by Id put operation typically these are written to a http.Request
*/
type ResellersByIDPutParams struct {

	/*Body*/
	Body *Replace1changeaspecificResellerRequest
	/*ContentType*/
	ContentType string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the resellers by Id put params
func (o *ResellersByIDPutParams) WithTimeout(timeout time.Duration) *ResellersByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resellers by Id put params
func (o *ResellersByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resellers by Id put params
func (o *ResellersByIDPutParams) WithContext(ctx context.Context) *ResellersByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resellers by Id put params
func (o *ResellersByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resellers by Id put params
func (o *ResellersByIDPutParams) WithHTTPClient(client *http.Client) *ResellersByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resellers by Id put params
func (o *ResellersByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the resellers by Id put params
func (o *ResellersByIDPutParams) WithBody(body *Replace1changeaspecificResellerRequest) *ResellersByIDPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the resellers by Id put params
func (o *ResellersByIDPutParams) SetBody(body *Replace1changeaspecificResellerRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the resellers by Id put params
func (o *ResellersByIDPutParams) WithContentType(contentType string) *ResellersByIDPutParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the resellers by Id put params
func (o *ResellersByIDPutParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the resellers by Id put params
func (o *ResellersByIDPutParams) WithID(id string) *ResellersByIDPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the resellers by Id put params
func (o *ResellersByIDPutParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ResellersByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
