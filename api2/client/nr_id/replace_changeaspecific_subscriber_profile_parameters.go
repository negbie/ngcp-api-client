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

// NewReplaceChangeaspecificSubscriberProfileParams creates a new ReplaceChangeaspecificSubscriberProfileParams object
// with the default values initialized.
func NewReplaceChangeaspecificSubscriberProfileParams() *ReplaceChangeaspecificSubscriberProfileParams {
	var ()
	return &ReplaceChangeaspecificSubscriberProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceChangeaspecificSubscriberProfileParamsWithTimeout creates a new ReplaceChangeaspecificSubscriberProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceChangeaspecificSubscriberProfileParamsWithTimeout(timeout time.Duration) *ReplaceChangeaspecificSubscriberProfileParams {
	var ()
	return &ReplaceChangeaspecificSubscriberProfileParams{

		timeout: timeout,
	}
}

// NewReplaceChangeaspecificSubscriberProfileParamsWithContext creates a new ReplaceChangeaspecificSubscriberProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceChangeaspecificSubscriberProfileParamsWithContext(ctx context.Context) *ReplaceChangeaspecificSubscriberProfileParams {
	var ()
	return &ReplaceChangeaspecificSubscriberProfileParams{

		Context: ctx,
	}
}

// NewReplaceChangeaspecificSubscriberProfileParamsWithHTTPClient creates a new ReplaceChangeaspecificSubscriberProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceChangeaspecificSubscriberProfileParamsWithHTTPClient(client *http.Client) *ReplaceChangeaspecificSubscriberProfileParams {
	var ()
	return &ReplaceChangeaspecificSubscriberProfileParams{
		HTTPClient: client,
	}
}

/*ReplaceChangeaspecificSubscriberProfileParams contains all the parameters to send to the API endpoint
for the replace changeaspecific subscriber profile operation typically these are written to a http.Request
*/
type ReplaceChangeaspecificSubscriberProfileParams struct {

	/*Body*/
	Body *Replace1changeaspecificSubscriberProfileRequest
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

// WithTimeout adds the timeout to the replace changeaspecific subscriber profile params
func (o *ReplaceChangeaspecificSubscriberProfileParams) WithTimeout(timeout time.Duration) *ReplaceChangeaspecificSubscriberProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace changeaspecific subscriber profile params
func (o *ReplaceChangeaspecificSubscriberProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace changeaspecific subscriber profile params
func (o *ReplaceChangeaspecificSubscriberProfileParams) WithContext(ctx context.Context) *ReplaceChangeaspecificSubscriberProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace changeaspecific subscriber profile params
func (o *ReplaceChangeaspecificSubscriberProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace changeaspecific subscriber profile params
func (o *ReplaceChangeaspecificSubscriberProfileParams) WithHTTPClient(client *http.Client) *ReplaceChangeaspecificSubscriberProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace changeaspecific subscriber profile params
func (o *ReplaceChangeaspecificSubscriberProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace changeaspecific subscriber profile params
func (o *ReplaceChangeaspecificSubscriberProfileParams) WithBody(body *Replace1changeaspecificSubscriberProfileRequest) *ReplaceChangeaspecificSubscriberProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace changeaspecific subscriber profile params
func (o *ReplaceChangeaspecificSubscriberProfileParams) SetBody(body *Replace1changeaspecificSubscriberProfileRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the replace changeaspecific subscriber profile params
func (o *ReplaceChangeaspecificSubscriberProfileParams) WithContentType(contentType string) *ReplaceChangeaspecificSubscriberProfileParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the replace changeaspecific subscriber profile params
func (o *ReplaceChangeaspecificSubscriberProfileParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the replace changeaspecific subscriber profile params
func (o *ReplaceChangeaspecificSubscriberProfileParams) WithID(id string) *ReplaceChangeaspecificSubscriberProfileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace changeaspecific subscriber profile params
func (o *ReplaceChangeaspecificSubscriberProfileParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceChangeaspecificSubscriberProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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