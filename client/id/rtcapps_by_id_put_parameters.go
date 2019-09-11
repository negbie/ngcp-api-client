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

// NewRtcappsByIDPutParams creates a new RtcappsByIDPutParams object
// with the default values initialized.
func NewRtcappsByIDPutParams() *RtcappsByIDPutParams {
	var ()
	return &RtcappsByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRtcappsByIDPutParamsWithTimeout creates a new RtcappsByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRtcappsByIDPutParamsWithTimeout(timeout time.Duration) *RtcappsByIDPutParams {
	var ()
	return &RtcappsByIDPutParams{

		timeout: timeout,
	}
}

// NewRtcappsByIDPutParamsWithContext creates a new RtcappsByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewRtcappsByIDPutParamsWithContext(ctx context.Context) *RtcappsByIDPutParams {
	var ()
	return &RtcappsByIDPutParams{

		Context: ctx,
	}
}

// NewRtcappsByIDPutParamsWithHTTPClient creates a new RtcappsByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRtcappsByIDPutParamsWithHTTPClient(client *http.Client) *RtcappsByIDPutParams {
	var ()
	return &RtcappsByIDPutParams{
		HTTPClient: client,
	}
}

/*RtcappsByIDPutParams contains all the parameters to send to the API endpoint
for the rtcapps by Id put operation typically these are written to a http.Request
*/
type RtcappsByIDPutParams struct {

	/*Body*/
	Body *Replace1changeaspecificRtcAppRequest
	/*ContentType*/
	ContentType string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the rtcapps by Id put params
func (o *RtcappsByIDPutParams) WithTimeout(timeout time.Duration) *RtcappsByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rtcapps by Id put params
func (o *RtcappsByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rtcapps by Id put params
func (o *RtcappsByIDPutParams) WithContext(ctx context.Context) *RtcappsByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rtcapps by Id put params
func (o *RtcappsByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rtcapps by Id put params
func (o *RtcappsByIDPutParams) WithHTTPClient(client *http.Client) *RtcappsByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rtcapps by Id put params
func (o *RtcappsByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the rtcapps by Id put params
func (o *RtcappsByIDPutParams) WithBody(body *Replace1changeaspecificRtcAppRequest) *RtcappsByIDPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the rtcapps by Id put params
func (o *RtcappsByIDPutParams) SetBody(body *Replace1changeaspecificRtcAppRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the rtcapps by Id put params
func (o *RtcappsByIDPutParams) WithContentType(contentType string) *RtcappsByIDPutParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the rtcapps by Id put params
func (o *RtcappsByIDPutParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the rtcapps by Id put params
func (o *RtcappsByIDPutParams) WithID(id string) *RtcappsByIDPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the rtcapps by Id put params
func (o *RtcappsByIDPutParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RtcappsByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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