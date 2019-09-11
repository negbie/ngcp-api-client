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

// NewSystemcontactsByIDPatchParams creates a new SystemcontactsByIDPatchParams object
// with the default values initialized.
func NewSystemcontactsByIDPatchParams() *SystemcontactsByIDPatchParams {
	var ()
	return &SystemcontactsByIDPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSystemcontactsByIDPatchParamsWithTimeout creates a new SystemcontactsByIDPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSystemcontactsByIDPatchParamsWithTimeout(timeout time.Duration) *SystemcontactsByIDPatchParams {
	var ()
	return &SystemcontactsByIDPatchParams{

		timeout: timeout,
	}
}

// NewSystemcontactsByIDPatchParamsWithContext creates a new SystemcontactsByIDPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewSystemcontactsByIDPatchParamsWithContext(ctx context.Context) *SystemcontactsByIDPatchParams {
	var ()
	return &SystemcontactsByIDPatchParams{

		Context: ctx,
	}
}

// NewSystemcontactsByIDPatchParamsWithHTTPClient creates a new SystemcontactsByIDPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSystemcontactsByIDPatchParamsWithHTTPClient(client *http.Client) *SystemcontactsByIDPatchParams {
	var ()
	return &SystemcontactsByIDPatchParams{
		HTTPClient: client,
	}
}

/*SystemcontactsByIDPatchParams contains all the parameters to send to the API endpoint
for the systemcontacts by Id patch operation typically these are written to a http.Request
*/
type SystemcontactsByIDPatchParams struct {

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

// WithTimeout adds the timeout to the systemcontacts by Id patch params
func (o *SystemcontactsByIDPatchParams) WithTimeout(timeout time.Duration) *SystemcontactsByIDPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the systemcontacts by Id patch params
func (o *SystemcontactsByIDPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the systemcontacts by Id patch params
func (o *SystemcontactsByIDPatchParams) WithContext(ctx context.Context) *SystemcontactsByIDPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the systemcontacts by Id patch params
func (o *SystemcontactsByIDPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the systemcontacts by Id patch params
func (o *SystemcontactsByIDPatchParams) WithHTTPClient(client *http.Client) *SystemcontactsByIDPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the systemcontacts by Id patch params
func (o *SystemcontactsByIDPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the systemcontacts by Id patch params
func (o *SystemcontactsByIDPatchParams) WithBody(body string) *SystemcontactsByIDPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the systemcontacts by Id patch params
func (o *SystemcontactsByIDPatchParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the systemcontacts by Id patch params
func (o *SystemcontactsByIDPatchParams) WithContentType(contentType string) *SystemcontactsByIDPatchParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the systemcontacts by Id patch params
func (o *SystemcontactsByIDPatchParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the systemcontacts by Id patch params
func (o *SystemcontactsByIDPatchParams) WithID(id string) *SystemcontactsByIDPatchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the systemcontacts by Id patch params
func (o *SystemcontactsByIDPatchParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SystemcontactsByIDPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
