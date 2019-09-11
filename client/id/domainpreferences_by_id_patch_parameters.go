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

// NewDomainpreferencesByIDPatchParams creates a new DomainpreferencesByIDPatchParams object
// with the default values initialized.
func NewDomainpreferencesByIDPatchParams() *DomainpreferencesByIDPatchParams {
	var ()
	return &DomainpreferencesByIDPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDomainpreferencesByIDPatchParamsWithTimeout creates a new DomainpreferencesByIDPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDomainpreferencesByIDPatchParamsWithTimeout(timeout time.Duration) *DomainpreferencesByIDPatchParams {
	var ()
	return &DomainpreferencesByIDPatchParams{

		timeout: timeout,
	}
}

// NewDomainpreferencesByIDPatchParamsWithContext creates a new DomainpreferencesByIDPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewDomainpreferencesByIDPatchParamsWithContext(ctx context.Context) *DomainpreferencesByIDPatchParams {
	var ()
	return &DomainpreferencesByIDPatchParams{

		Context: ctx,
	}
}

// NewDomainpreferencesByIDPatchParamsWithHTTPClient creates a new DomainpreferencesByIDPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDomainpreferencesByIDPatchParamsWithHTTPClient(client *http.Client) *DomainpreferencesByIDPatchParams {
	var ()
	return &DomainpreferencesByIDPatchParams{
		HTTPClient: client,
	}
}

/*DomainpreferencesByIDPatchParams contains all the parameters to send to the API endpoint
for the domainpreferences by Id patch operation typically these are written to a http.Request
*/
type DomainpreferencesByIDPatchParams struct {

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

// WithTimeout adds the timeout to the domainpreferences by Id patch params
func (o *DomainpreferencesByIDPatchParams) WithTimeout(timeout time.Duration) *DomainpreferencesByIDPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domainpreferences by Id patch params
func (o *DomainpreferencesByIDPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domainpreferences by Id patch params
func (o *DomainpreferencesByIDPatchParams) WithContext(ctx context.Context) *DomainpreferencesByIDPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domainpreferences by Id patch params
func (o *DomainpreferencesByIDPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domainpreferences by Id patch params
func (o *DomainpreferencesByIDPatchParams) WithHTTPClient(client *http.Client) *DomainpreferencesByIDPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domainpreferences by Id patch params
func (o *DomainpreferencesByIDPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the domainpreferences by Id patch params
func (o *DomainpreferencesByIDPatchParams) WithBody(body string) *DomainpreferencesByIDPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the domainpreferences by Id patch params
func (o *DomainpreferencesByIDPatchParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the domainpreferences by Id patch params
func (o *DomainpreferencesByIDPatchParams) WithContentType(contentType string) *DomainpreferencesByIDPatchParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the domainpreferences by Id patch params
func (o *DomainpreferencesByIDPatchParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the domainpreferences by Id patch params
func (o *DomainpreferencesByIDPatchParams) WithID(id string) *DomainpreferencesByIDPatchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the domainpreferences by Id patch params
func (o *DomainpreferencesByIDPatchParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DomainpreferencesByIDPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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