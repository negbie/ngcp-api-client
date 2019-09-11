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

// NewCfsourcesetsByIDPutParams creates a new CfsourcesetsByIDPutParams object
// with the default values initialized.
func NewCfsourcesetsByIDPutParams() *CfsourcesetsByIDPutParams {
	var ()
	return &CfsourcesetsByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCfsourcesetsByIDPutParamsWithTimeout creates a new CfsourcesetsByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCfsourcesetsByIDPutParamsWithTimeout(timeout time.Duration) *CfsourcesetsByIDPutParams {
	var ()
	return &CfsourcesetsByIDPutParams{

		timeout: timeout,
	}
}

// NewCfsourcesetsByIDPutParamsWithContext creates a new CfsourcesetsByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewCfsourcesetsByIDPutParamsWithContext(ctx context.Context) *CfsourcesetsByIDPutParams {
	var ()
	return &CfsourcesetsByIDPutParams{

		Context: ctx,
	}
}

// NewCfsourcesetsByIDPutParamsWithHTTPClient creates a new CfsourcesetsByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCfsourcesetsByIDPutParamsWithHTTPClient(client *http.Client) *CfsourcesetsByIDPutParams {
	var ()
	return &CfsourcesetsByIDPutParams{
		HTTPClient: client,
	}
}

/*CfsourcesetsByIDPutParams contains all the parameters to send to the API endpoint
for the cfsourcesets by Id put operation typically these are written to a http.Request
*/
type CfsourcesetsByIDPutParams struct {

	/*Body*/
	Body *Replace1changeaspecificCFSourceSetRequest
	/*ContentType*/
	ContentType string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cfsourcesets by Id put params
func (o *CfsourcesetsByIDPutParams) WithTimeout(timeout time.Duration) *CfsourcesetsByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cfsourcesets by Id put params
func (o *CfsourcesetsByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cfsourcesets by Id put params
func (o *CfsourcesetsByIDPutParams) WithContext(ctx context.Context) *CfsourcesetsByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cfsourcesets by Id put params
func (o *CfsourcesetsByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cfsourcesets by Id put params
func (o *CfsourcesetsByIDPutParams) WithHTTPClient(client *http.Client) *CfsourcesetsByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cfsourcesets by Id put params
func (o *CfsourcesetsByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cfsourcesets by Id put params
func (o *CfsourcesetsByIDPutParams) WithBody(body *Replace1changeaspecificCFSourceSetRequest) *CfsourcesetsByIDPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cfsourcesets by Id put params
func (o *CfsourcesetsByIDPutParams) SetBody(body *Replace1changeaspecificCFSourceSetRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the cfsourcesets by Id put params
func (o *CfsourcesetsByIDPutParams) WithContentType(contentType string) *CfsourcesetsByIDPutParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the cfsourcesets by Id put params
func (o *CfsourcesetsByIDPutParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the cfsourcesets by Id put params
func (o *CfsourcesetsByIDPutParams) WithID(id string) *CfsourcesetsByIDPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the cfsourcesets by Id put params
func (o *CfsourcesetsByIDPutParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CfsourcesetsByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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