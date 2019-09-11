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

// NewRtcnetworksByIDPatchParams creates a new RtcnetworksByIDPatchParams object
// with the default values initialized.
func NewRtcnetworksByIDPatchParams() *RtcnetworksByIDPatchParams {
	var ()
	return &RtcnetworksByIDPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRtcnetworksByIDPatchParamsWithTimeout creates a new RtcnetworksByIDPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRtcnetworksByIDPatchParamsWithTimeout(timeout time.Duration) *RtcnetworksByIDPatchParams {
	var ()
	return &RtcnetworksByIDPatchParams{

		timeout: timeout,
	}
}

// NewRtcnetworksByIDPatchParamsWithContext creates a new RtcnetworksByIDPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewRtcnetworksByIDPatchParamsWithContext(ctx context.Context) *RtcnetworksByIDPatchParams {
	var ()
	return &RtcnetworksByIDPatchParams{

		Context: ctx,
	}
}

// NewRtcnetworksByIDPatchParamsWithHTTPClient creates a new RtcnetworksByIDPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRtcnetworksByIDPatchParamsWithHTTPClient(client *http.Client) *RtcnetworksByIDPatchParams {
	var ()
	return &RtcnetworksByIDPatchParams{
		HTTPClient: client,
	}
}

/*RtcnetworksByIDPatchParams contains all the parameters to send to the API endpoint
for the rtcnetworks by Id patch operation typically these are written to a http.Request
*/
type RtcnetworksByIDPatchParams struct {

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

// WithTimeout adds the timeout to the rtcnetworks by Id patch params
func (o *RtcnetworksByIDPatchParams) WithTimeout(timeout time.Duration) *RtcnetworksByIDPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rtcnetworks by Id patch params
func (o *RtcnetworksByIDPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rtcnetworks by Id patch params
func (o *RtcnetworksByIDPatchParams) WithContext(ctx context.Context) *RtcnetworksByIDPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rtcnetworks by Id patch params
func (o *RtcnetworksByIDPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rtcnetworks by Id patch params
func (o *RtcnetworksByIDPatchParams) WithHTTPClient(client *http.Client) *RtcnetworksByIDPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rtcnetworks by Id patch params
func (o *RtcnetworksByIDPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the rtcnetworks by Id patch params
func (o *RtcnetworksByIDPatchParams) WithBody(body string) *RtcnetworksByIDPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the rtcnetworks by Id patch params
func (o *RtcnetworksByIDPatchParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the rtcnetworks by Id patch params
func (o *RtcnetworksByIDPatchParams) WithContentType(contentType string) *RtcnetworksByIDPatchParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the rtcnetworks by Id patch params
func (o *RtcnetworksByIDPatchParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the rtcnetworks by Id patch params
func (o *RtcnetworksByIDPatchParams) WithID(id string) *RtcnetworksByIDPatchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the rtcnetworks by Id patch params
func (o *RtcnetworksByIDPatchParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RtcnetworksByIDPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
