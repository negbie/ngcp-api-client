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

// NewCfmappingsByIDPatchParams creates a new CfmappingsByIDPatchParams object
// with the default values initialized.
func NewCfmappingsByIDPatchParams() *CfmappingsByIDPatchParams {
	var ()
	return &CfmappingsByIDPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCfmappingsByIDPatchParamsWithTimeout creates a new CfmappingsByIDPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCfmappingsByIDPatchParamsWithTimeout(timeout time.Duration) *CfmappingsByIDPatchParams {
	var ()
	return &CfmappingsByIDPatchParams{

		timeout: timeout,
	}
}

// NewCfmappingsByIDPatchParamsWithContext creates a new CfmappingsByIDPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewCfmappingsByIDPatchParamsWithContext(ctx context.Context) *CfmappingsByIDPatchParams {
	var ()
	return &CfmappingsByIDPatchParams{

		Context: ctx,
	}
}

// NewCfmappingsByIDPatchParamsWithHTTPClient creates a new CfmappingsByIDPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCfmappingsByIDPatchParamsWithHTTPClient(client *http.Client) *CfmappingsByIDPatchParams {
	var ()
	return &CfmappingsByIDPatchParams{
		HTTPClient: client,
	}
}

/*CfmappingsByIDPatchParams contains all the parameters to send to the API endpoint
for the cfmappings by Id patch operation typically these are written to a http.Request
*/
type CfmappingsByIDPatchParams struct {

	/*Body*/
	Body string
	/*ContentType*/
	ContentType string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cfmappings by Id patch params
func (o *CfmappingsByIDPatchParams) WithTimeout(timeout time.Duration) *CfmappingsByIDPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cfmappings by Id patch params
func (o *CfmappingsByIDPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cfmappings by Id patch params
func (o *CfmappingsByIDPatchParams) WithContext(ctx context.Context) *CfmappingsByIDPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cfmappings by Id patch params
func (o *CfmappingsByIDPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cfmappings by Id patch params
func (o *CfmappingsByIDPatchParams) WithHTTPClient(client *http.Client) *CfmappingsByIDPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cfmappings by Id patch params
func (o *CfmappingsByIDPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cfmappings by Id patch params
func (o *CfmappingsByIDPatchParams) WithBody(body string) *CfmappingsByIDPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cfmappings by Id patch params
func (o *CfmappingsByIDPatchParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the cfmappings by Id patch params
func (o *CfmappingsByIDPatchParams) WithContentType(contentType string) *CfmappingsByIDPatchParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the cfmappings by Id patch params
func (o *CfmappingsByIDPatchParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the cfmappings by Id patch params
func (o *CfmappingsByIDPatchParams) WithID(id string) *CfmappingsByIDPatchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the cfmappings by Id patch params
func (o *CfmappingsByIDPatchParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CfmappingsByIDPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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