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

// NewReplaceChangeaspecificCFTimeSetParams creates a new ReplaceChangeaspecificCFTimeSetParams object
// with the default values initialized.
func NewReplaceChangeaspecificCFTimeSetParams() *ReplaceChangeaspecificCFTimeSetParams {
	var ()
	return &ReplaceChangeaspecificCFTimeSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceChangeaspecificCFTimeSetParamsWithTimeout creates a new ReplaceChangeaspecificCFTimeSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceChangeaspecificCFTimeSetParamsWithTimeout(timeout time.Duration) *ReplaceChangeaspecificCFTimeSetParams {
	var ()
	return &ReplaceChangeaspecificCFTimeSetParams{

		timeout: timeout,
	}
}

// NewReplaceChangeaspecificCFTimeSetParamsWithContext creates a new ReplaceChangeaspecificCFTimeSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceChangeaspecificCFTimeSetParamsWithContext(ctx context.Context) *ReplaceChangeaspecificCFTimeSetParams {
	var ()
	return &ReplaceChangeaspecificCFTimeSetParams{

		Context: ctx,
	}
}

// NewReplaceChangeaspecificCFTimeSetParamsWithHTTPClient creates a new ReplaceChangeaspecificCFTimeSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceChangeaspecificCFTimeSetParamsWithHTTPClient(client *http.Client) *ReplaceChangeaspecificCFTimeSetParams {
	var ()
	return &ReplaceChangeaspecificCFTimeSetParams{
		HTTPClient: client,
	}
}

/*ReplaceChangeaspecificCFTimeSetParams contains all the parameters to send to the API endpoint
for the replace changeaspecific c f time set operation typically these are written to a http.Request
*/
type ReplaceChangeaspecificCFTimeSetParams struct {

	/*Body*/
	Body *Replace1changeaspecificCFTimeSetRequest
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

// WithTimeout adds the timeout to the replace changeaspecific c f time set params
func (o *ReplaceChangeaspecificCFTimeSetParams) WithTimeout(timeout time.Duration) *ReplaceChangeaspecificCFTimeSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace changeaspecific c f time set params
func (o *ReplaceChangeaspecificCFTimeSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace changeaspecific c f time set params
func (o *ReplaceChangeaspecificCFTimeSetParams) WithContext(ctx context.Context) *ReplaceChangeaspecificCFTimeSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace changeaspecific c f time set params
func (o *ReplaceChangeaspecificCFTimeSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace changeaspecific c f time set params
func (o *ReplaceChangeaspecificCFTimeSetParams) WithHTTPClient(client *http.Client) *ReplaceChangeaspecificCFTimeSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace changeaspecific c f time set params
func (o *ReplaceChangeaspecificCFTimeSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace changeaspecific c f time set params
func (o *ReplaceChangeaspecificCFTimeSetParams) WithBody(body *Replace1changeaspecificCFTimeSetRequest) *ReplaceChangeaspecificCFTimeSetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace changeaspecific c f time set params
func (o *ReplaceChangeaspecificCFTimeSetParams) SetBody(body *Replace1changeaspecificCFTimeSetRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the replace changeaspecific c f time set params
func (o *ReplaceChangeaspecificCFTimeSetParams) WithContentType(contentType string) *ReplaceChangeaspecificCFTimeSetParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the replace changeaspecific c f time set params
func (o *ReplaceChangeaspecificCFTimeSetParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the replace changeaspecific c f time set params
func (o *ReplaceChangeaspecificCFTimeSetParams) WithID(id string) *ReplaceChangeaspecificCFTimeSetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace changeaspecific c f time set params
func (o *ReplaceChangeaspecificCFTimeSetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceChangeaspecificCFTimeSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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