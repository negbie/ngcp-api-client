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

// NewReplaceChangeaspecificRtcAppParams creates a new ReplaceChangeaspecificRtcAppParams object
// with the default values initialized.
func NewReplaceChangeaspecificRtcAppParams() *ReplaceChangeaspecificRtcAppParams {
	var ()
	return &ReplaceChangeaspecificRtcAppParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceChangeaspecificRtcAppParamsWithTimeout creates a new ReplaceChangeaspecificRtcAppParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceChangeaspecificRtcAppParamsWithTimeout(timeout time.Duration) *ReplaceChangeaspecificRtcAppParams {
	var ()
	return &ReplaceChangeaspecificRtcAppParams{

		timeout: timeout,
	}
}

// NewReplaceChangeaspecificRtcAppParamsWithContext creates a new ReplaceChangeaspecificRtcAppParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceChangeaspecificRtcAppParamsWithContext(ctx context.Context) *ReplaceChangeaspecificRtcAppParams {
	var ()
	return &ReplaceChangeaspecificRtcAppParams{

		Context: ctx,
	}
}

// NewReplaceChangeaspecificRtcAppParamsWithHTTPClient creates a new ReplaceChangeaspecificRtcAppParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceChangeaspecificRtcAppParamsWithHTTPClient(client *http.Client) *ReplaceChangeaspecificRtcAppParams {
	var ()
	return &ReplaceChangeaspecificRtcAppParams{
		HTTPClient: client,
	}
}

/*ReplaceChangeaspecificRtcAppParams contains all the parameters to send to the API endpoint
for the replace changeaspecific rtc app operation typically these are written to a http.Request
*/
type ReplaceChangeaspecificRtcAppParams struct {

	/*Body*/
	Body *Replace1changeaspecificRtcAppRequest
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

// WithTimeout adds the timeout to the replace changeaspecific rtc app params
func (o *ReplaceChangeaspecificRtcAppParams) WithTimeout(timeout time.Duration) *ReplaceChangeaspecificRtcAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace changeaspecific rtc app params
func (o *ReplaceChangeaspecificRtcAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace changeaspecific rtc app params
func (o *ReplaceChangeaspecificRtcAppParams) WithContext(ctx context.Context) *ReplaceChangeaspecificRtcAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace changeaspecific rtc app params
func (o *ReplaceChangeaspecificRtcAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace changeaspecific rtc app params
func (o *ReplaceChangeaspecificRtcAppParams) WithHTTPClient(client *http.Client) *ReplaceChangeaspecificRtcAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace changeaspecific rtc app params
func (o *ReplaceChangeaspecificRtcAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace changeaspecific rtc app params
func (o *ReplaceChangeaspecificRtcAppParams) WithBody(body *Replace1changeaspecificRtcAppRequest) *ReplaceChangeaspecificRtcAppParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace changeaspecific rtc app params
func (o *ReplaceChangeaspecificRtcAppParams) SetBody(body *Replace1changeaspecificRtcAppRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the replace changeaspecific rtc app params
func (o *ReplaceChangeaspecificRtcAppParams) WithContentType(contentType string) *ReplaceChangeaspecificRtcAppParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the replace changeaspecific rtc app params
func (o *ReplaceChangeaspecificRtcAppParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the replace changeaspecific rtc app params
func (o *ReplaceChangeaspecificRtcAppParams) WithID(id string) *ReplaceChangeaspecificRtcAppParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace changeaspecific rtc app params
func (o *ReplaceChangeaspecificRtcAppParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceChangeaspecificRtcAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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