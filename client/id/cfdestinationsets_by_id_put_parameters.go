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

// NewCfdestinationsetsByIDPutParams creates a new CfdestinationsetsByIDPutParams object
// with the default values initialized.
func NewCfdestinationsetsByIDPutParams() *CfdestinationsetsByIDPutParams {
	var ()
	return &CfdestinationsetsByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCfdestinationsetsByIDPutParamsWithTimeout creates a new CfdestinationsetsByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCfdestinationsetsByIDPutParamsWithTimeout(timeout time.Duration) *CfdestinationsetsByIDPutParams {
	var ()
	return &CfdestinationsetsByIDPutParams{

		timeout: timeout,
	}
}

// NewCfdestinationsetsByIDPutParamsWithContext creates a new CfdestinationsetsByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewCfdestinationsetsByIDPutParamsWithContext(ctx context.Context) *CfdestinationsetsByIDPutParams {
	var ()
	return &CfdestinationsetsByIDPutParams{

		Context: ctx,
	}
}

// NewCfdestinationsetsByIDPutParamsWithHTTPClient creates a new CfdestinationsetsByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCfdestinationsetsByIDPutParamsWithHTTPClient(client *http.Client) *CfdestinationsetsByIDPutParams {
	var ()
	return &CfdestinationsetsByIDPutParams{
		HTTPClient: client,
	}
}

/*CfdestinationsetsByIDPutParams contains all the parameters to send to the API endpoint
for the cfdestinationsets by Id put operation typically these are written to a http.Request
*/
type CfdestinationsetsByIDPutParams struct {

	/*Body*/
	Body *Replace1changeaspecificCFDestinationSetRequest
	/*ContentType*/
	ContentType string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cfdestinationsets by Id put params
func (o *CfdestinationsetsByIDPutParams) WithTimeout(timeout time.Duration) *CfdestinationsetsByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cfdestinationsets by Id put params
func (o *CfdestinationsetsByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cfdestinationsets by Id put params
func (o *CfdestinationsetsByIDPutParams) WithContext(ctx context.Context) *CfdestinationsetsByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cfdestinationsets by Id put params
func (o *CfdestinationsetsByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cfdestinationsets by Id put params
func (o *CfdestinationsetsByIDPutParams) WithHTTPClient(client *http.Client) *CfdestinationsetsByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cfdestinationsets by Id put params
func (o *CfdestinationsetsByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cfdestinationsets by Id put params
func (o *CfdestinationsetsByIDPutParams) WithBody(body *Replace1changeaspecificCFDestinationSetRequest) *CfdestinationsetsByIDPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cfdestinationsets by Id put params
func (o *CfdestinationsetsByIDPutParams) SetBody(body *Replace1changeaspecificCFDestinationSetRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the cfdestinationsets by Id put params
func (o *CfdestinationsetsByIDPutParams) WithContentType(contentType string) *CfdestinationsetsByIDPutParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the cfdestinationsets by Id put params
func (o *CfdestinationsetsByIDPutParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the cfdestinationsets by Id put params
func (o *CfdestinationsetsByIDPutParams) WithID(id string) *CfdestinationsetsByIDPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the cfdestinationsets by Id put params
func (o *CfdestinationsetsByIDPutParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CfdestinationsetsByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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