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

// NewReplaceChangeaspecificCustomerPreferenceParams creates a new ReplaceChangeaspecificCustomerPreferenceParams object
// with the default values initialized.
func NewReplaceChangeaspecificCustomerPreferenceParams() *ReplaceChangeaspecificCustomerPreferenceParams {
	var ()
	return &ReplaceChangeaspecificCustomerPreferenceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceChangeaspecificCustomerPreferenceParamsWithTimeout creates a new ReplaceChangeaspecificCustomerPreferenceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceChangeaspecificCustomerPreferenceParamsWithTimeout(timeout time.Duration) *ReplaceChangeaspecificCustomerPreferenceParams {
	var ()
	return &ReplaceChangeaspecificCustomerPreferenceParams{

		timeout: timeout,
	}
}

// NewReplaceChangeaspecificCustomerPreferenceParamsWithContext creates a new ReplaceChangeaspecificCustomerPreferenceParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceChangeaspecificCustomerPreferenceParamsWithContext(ctx context.Context) *ReplaceChangeaspecificCustomerPreferenceParams {
	var ()
	return &ReplaceChangeaspecificCustomerPreferenceParams{

		Context: ctx,
	}
}

// NewReplaceChangeaspecificCustomerPreferenceParamsWithHTTPClient creates a new ReplaceChangeaspecificCustomerPreferenceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceChangeaspecificCustomerPreferenceParamsWithHTTPClient(client *http.Client) *ReplaceChangeaspecificCustomerPreferenceParams {
	var ()
	return &ReplaceChangeaspecificCustomerPreferenceParams{
		HTTPClient: client,
	}
}

/*ReplaceChangeaspecificCustomerPreferenceParams contains all the parameters to send to the API endpoint
for the replace changeaspecific customer preference operation typically these are written to a http.Request
*/
type ReplaceChangeaspecificCustomerPreferenceParams struct {

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

// WithTimeout adds the timeout to the replace changeaspecific customer preference params
func (o *ReplaceChangeaspecificCustomerPreferenceParams) WithTimeout(timeout time.Duration) *ReplaceChangeaspecificCustomerPreferenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace changeaspecific customer preference params
func (o *ReplaceChangeaspecificCustomerPreferenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace changeaspecific customer preference params
func (o *ReplaceChangeaspecificCustomerPreferenceParams) WithContext(ctx context.Context) *ReplaceChangeaspecificCustomerPreferenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace changeaspecific customer preference params
func (o *ReplaceChangeaspecificCustomerPreferenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace changeaspecific customer preference params
func (o *ReplaceChangeaspecificCustomerPreferenceParams) WithHTTPClient(client *http.Client) *ReplaceChangeaspecificCustomerPreferenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace changeaspecific customer preference params
func (o *ReplaceChangeaspecificCustomerPreferenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace changeaspecific customer preference params
func (o *ReplaceChangeaspecificCustomerPreferenceParams) WithBody(body string) *ReplaceChangeaspecificCustomerPreferenceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace changeaspecific customer preference params
func (o *ReplaceChangeaspecificCustomerPreferenceParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the replace changeaspecific customer preference params
func (o *ReplaceChangeaspecificCustomerPreferenceParams) WithContentType(contentType string) *ReplaceChangeaspecificCustomerPreferenceParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the replace changeaspecific customer preference params
func (o *ReplaceChangeaspecificCustomerPreferenceParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the replace changeaspecific customer preference params
func (o *ReplaceChangeaspecificCustomerPreferenceParams) WithID(id string) *ReplaceChangeaspecificCustomerPreferenceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace changeaspecific customer preference params
func (o *ReplaceChangeaspecificCustomerPreferenceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceChangeaspecificCustomerPreferenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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