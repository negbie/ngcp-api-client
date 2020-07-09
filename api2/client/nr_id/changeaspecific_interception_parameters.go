// Code generated by go-swagger; DO NOT EDIT.

package nr_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewChangeaspecificInterceptionParams creates a new ChangeaspecificInterceptionParams object
// with the default values initialized.
func NewChangeaspecificInterceptionParams() *ChangeaspecificInterceptionParams {
	var ()
	return &ChangeaspecificInterceptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeaspecificInterceptionParamsWithTimeout creates a new ChangeaspecificInterceptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeaspecificInterceptionParamsWithTimeout(timeout time.Duration) *ChangeaspecificInterceptionParams {
	var ()
	return &ChangeaspecificInterceptionParams{

		timeout: timeout,
	}
}

// NewChangeaspecificInterceptionParamsWithContext creates a new ChangeaspecificInterceptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeaspecificInterceptionParamsWithContext(ctx context.Context) *ChangeaspecificInterceptionParams {
	var ()
	return &ChangeaspecificInterceptionParams{

		Context: ctx,
	}
}

// NewChangeaspecificInterceptionParamsWithHTTPClient creates a new ChangeaspecificInterceptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeaspecificInterceptionParamsWithHTTPClient(client *http.Client) *ChangeaspecificInterceptionParams {
	var ()
	return &ChangeaspecificInterceptionParams{
		HTTPClient: client,
	}
}

/*ChangeaspecificInterceptionParams contains all the parameters to send to the API endpoint
for the changeaspecific interception operation typically these are written to a http.Request
*/
type ChangeaspecificInterceptionParams struct {

	/*Body*/
	Body string
	/*ContentType*/
	ContentType string
	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the changeaspecific interception params
func (o *ChangeaspecificInterceptionParams) WithTimeout(timeout time.Duration) *ChangeaspecificInterceptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the changeaspecific interception params
func (o *ChangeaspecificInterceptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the changeaspecific interception params
func (o *ChangeaspecificInterceptionParams) WithContext(ctx context.Context) *ChangeaspecificInterceptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the changeaspecific interception params
func (o *ChangeaspecificInterceptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the changeaspecific interception params
func (o *ChangeaspecificInterceptionParams) WithHTTPClient(client *http.Client) *ChangeaspecificInterceptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the changeaspecific interception params
func (o *ChangeaspecificInterceptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the changeaspecific interception params
func (o *ChangeaspecificInterceptionParams) WithBody(body string) *ChangeaspecificInterceptionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the changeaspecific interception params
func (o *ChangeaspecificInterceptionParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the changeaspecific interception params
func (o *ChangeaspecificInterceptionParams) WithContentType(contentType string) *ChangeaspecificInterceptionParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the changeaspecific interception params
func (o *ChangeaspecificInterceptionParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the changeaspecific interception params
func (o *ChangeaspecificInterceptionParams) WithID(id string) *ChangeaspecificInterceptionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the changeaspecific interception params
func (o *ChangeaspecificInterceptionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeaspecificInterceptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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