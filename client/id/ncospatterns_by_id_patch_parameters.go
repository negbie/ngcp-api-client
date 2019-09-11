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

// NewNcospatternsByIDPatchParams creates a new NcospatternsByIDPatchParams object
// with the default values initialized.
func NewNcospatternsByIDPatchParams() *NcospatternsByIDPatchParams {
	var ()
	return &NcospatternsByIDPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNcospatternsByIDPatchParamsWithTimeout creates a new NcospatternsByIDPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNcospatternsByIDPatchParamsWithTimeout(timeout time.Duration) *NcospatternsByIDPatchParams {
	var ()
	return &NcospatternsByIDPatchParams{

		timeout: timeout,
	}
}

// NewNcospatternsByIDPatchParamsWithContext creates a new NcospatternsByIDPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewNcospatternsByIDPatchParamsWithContext(ctx context.Context) *NcospatternsByIDPatchParams {
	var ()
	return &NcospatternsByIDPatchParams{

		Context: ctx,
	}
}

// NewNcospatternsByIDPatchParamsWithHTTPClient creates a new NcospatternsByIDPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNcospatternsByIDPatchParamsWithHTTPClient(client *http.Client) *NcospatternsByIDPatchParams {
	var ()
	return &NcospatternsByIDPatchParams{
		HTTPClient: client,
	}
}

/*NcospatternsByIDPatchParams contains all the parameters to send to the API endpoint
for the ncospatterns by Id patch operation typically these are written to a http.Request
*/
type NcospatternsByIDPatchParams struct {

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

// WithTimeout adds the timeout to the ncospatterns by Id patch params
func (o *NcospatternsByIDPatchParams) WithTimeout(timeout time.Duration) *NcospatternsByIDPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ncospatterns by Id patch params
func (o *NcospatternsByIDPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ncospatterns by Id patch params
func (o *NcospatternsByIDPatchParams) WithContext(ctx context.Context) *NcospatternsByIDPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ncospatterns by Id patch params
func (o *NcospatternsByIDPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ncospatterns by Id patch params
func (o *NcospatternsByIDPatchParams) WithHTTPClient(client *http.Client) *NcospatternsByIDPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ncospatterns by Id patch params
func (o *NcospatternsByIDPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the ncospatterns by Id patch params
func (o *NcospatternsByIDPatchParams) WithBody(body string) *NcospatternsByIDPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the ncospatterns by Id patch params
func (o *NcospatternsByIDPatchParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the ncospatterns by Id patch params
func (o *NcospatternsByIDPatchParams) WithContentType(contentType string) *NcospatternsByIDPatchParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the ncospatterns by Id patch params
func (o *NcospatternsByIDPatchParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the ncospatterns by Id patch params
func (o *NcospatternsByIDPatchParams) WithID(id string) *NcospatternsByIDPatchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ncospatterns by Id patch params
func (o *NcospatternsByIDPatchParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *NcospatternsByIDPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
