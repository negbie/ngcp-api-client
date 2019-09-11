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

// NewCcmapentriesByIDPutParams creates a new CcmapentriesByIDPutParams object
// with the default values initialized.
func NewCcmapentriesByIDPutParams() *CcmapentriesByIDPutParams {
	var ()
	return &CcmapentriesByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCcmapentriesByIDPutParamsWithTimeout creates a new CcmapentriesByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCcmapentriesByIDPutParamsWithTimeout(timeout time.Duration) *CcmapentriesByIDPutParams {
	var ()
	return &CcmapentriesByIDPutParams{

		timeout: timeout,
	}
}

// NewCcmapentriesByIDPutParamsWithContext creates a new CcmapentriesByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewCcmapentriesByIDPutParamsWithContext(ctx context.Context) *CcmapentriesByIDPutParams {
	var ()
	return &CcmapentriesByIDPutParams{

		Context: ctx,
	}
}

// NewCcmapentriesByIDPutParamsWithHTTPClient creates a new CcmapentriesByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCcmapentriesByIDPutParamsWithHTTPClient(client *http.Client) *CcmapentriesByIDPutParams {
	var ()
	return &CcmapentriesByIDPutParams{
		HTTPClient: client,
	}
}

/*CcmapentriesByIDPutParams contains all the parameters to send to the API endpoint
for the ccmapentries by Id put operation typically these are written to a http.Request
*/
type CcmapentriesByIDPutParams struct {

	/*Body*/
	Body *Replace1changeaspecificCCMapEntryRequest
	/*ContentType*/
	ContentType string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ccmapentries by Id put params
func (o *CcmapentriesByIDPutParams) WithTimeout(timeout time.Duration) *CcmapentriesByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ccmapentries by Id put params
func (o *CcmapentriesByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ccmapentries by Id put params
func (o *CcmapentriesByIDPutParams) WithContext(ctx context.Context) *CcmapentriesByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ccmapentries by Id put params
func (o *CcmapentriesByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ccmapentries by Id put params
func (o *CcmapentriesByIDPutParams) WithHTTPClient(client *http.Client) *CcmapentriesByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ccmapentries by Id put params
func (o *CcmapentriesByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the ccmapentries by Id put params
func (o *CcmapentriesByIDPutParams) WithBody(body *Replace1changeaspecificCCMapEntryRequest) *CcmapentriesByIDPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the ccmapentries by Id put params
func (o *CcmapentriesByIDPutParams) SetBody(body *Replace1changeaspecificCCMapEntryRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the ccmapentries by Id put params
func (o *CcmapentriesByIDPutParams) WithContentType(contentType string) *CcmapentriesByIDPutParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the ccmapentries by Id put params
func (o *CcmapentriesByIDPutParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the ccmapentries by Id put params
func (o *CcmapentriesByIDPutParams) WithID(id string) *CcmapentriesByIDPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ccmapentries by Id put params
func (o *CcmapentriesByIDPutParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CcmapentriesByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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