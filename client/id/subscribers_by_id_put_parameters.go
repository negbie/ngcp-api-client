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

// NewSubscribersByIDPutParams creates a new SubscribersByIDPutParams object
// with the default values initialized.
func NewSubscribersByIDPutParams() *SubscribersByIDPutParams {
	var ()
	return &SubscribersByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubscribersByIDPutParamsWithTimeout creates a new SubscribersByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubscribersByIDPutParamsWithTimeout(timeout time.Duration) *SubscribersByIDPutParams {
	var ()
	return &SubscribersByIDPutParams{

		timeout: timeout,
	}
}

// NewSubscribersByIDPutParamsWithContext creates a new SubscribersByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubscribersByIDPutParamsWithContext(ctx context.Context) *SubscribersByIDPutParams {
	var ()
	return &SubscribersByIDPutParams{

		Context: ctx,
	}
}

// NewSubscribersByIDPutParamsWithHTTPClient creates a new SubscribersByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSubscribersByIDPutParamsWithHTTPClient(client *http.Client) *SubscribersByIDPutParams {
	var ()
	return &SubscribersByIDPutParams{
		HTTPClient: client,
	}
}

/*SubscribersByIDPutParams contains all the parameters to send to the API endpoint
for the subscribers by Id put operation typically these are written to a http.Request
*/
type SubscribersByIDPutParams struct {

	/*Body*/
	Body *Replace1changeaspecificSubscriberRequest
	/*ContentType*/
	ContentType string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the subscribers by Id put params
func (o *SubscribersByIDPutParams) WithTimeout(timeout time.Duration) *SubscribersByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscribers by Id put params
func (o *SubscribersByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscribers by Id put params
func (o *SubscribersByIDPutParams) WithContext(ctx context.Context) *SubscribersByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscribers by Id put params
func (o *SubscribersByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscribers by Id put params
func (o *SubscribersByIDPutParams) WithHTTPClient(client *http.Client) *SubscribersByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscribers by Id put params
func (o *SubscribersByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the subscribers by Id put params
func (o *SubscribersByIDPutParams) WithBody(body *Replace1changeaspecificSubscriberRequest) *SubscribersByIDPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the subscribers by Id put params
func (o *SubscribersByIDPutParams) SetBody(body *Replace1changeaspecificSubscriberRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the subscribers by Id put params
func (o *SubscribersByIDPutParams) WithContentType(contentType string) *SubscribersByIDPutParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the subscribers by Id put params
func (o *SubscribersByIDPutParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the subscribers by Id put params
func (o *SubscribersByIDPutParams) WithID(id string) *SubscribersByIDPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the subscribers by Id put params
func (o *SubscribersByIDPutParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SubscribersByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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