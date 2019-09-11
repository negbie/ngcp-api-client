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

// NewSubscriberpreferencesByIDPutParams creates a new SubscriberpreferencesByIDPutParams object
// with the default values initialized.
func NewSubscriberpreferencesByIDPutParams() *SubscriberpreferencesByIDPutParams {
	var ()
	return &SubscriberpreferencesByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubscriberpreferencesByIDPutParamsWithTimeout creates a new SubscriberpreferencesByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubscriberpreferencesByIDPutParamsWithTimeout(timeout time.Duration) *SubscriberpreferencesByIDPutParams {
	var ()
	return &SubscriberpreferencesByIDPutParams{

		timeout: timeout,
	}
}

// NewSubscriberpreferencesByIDPutParamsWithContext creates a new SubscriberpreferencesByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubscriberpreferencesByIDPutParamsWithContext(ctx context.Context) *SubscriberpreferencesByIDPutParams {
	var ()
	return &SubscriberpreferencesByIDPutParams{

		Context: ctx,
	}
}

// NewSubscriberpreferencesByIDPutParamsWithHTTPClient creates a new SubscriberpreferencesByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSubscriberpreferencesByIDPutParamsWithHTTPClient(client *http.Client) *SubscriberpreferencesByIDPutParams {
	var ()
	return &SubscriberpreferencesByIDPutParams{
		HTTPClient: client,
	}
}

/*SubscriberpreferencesByIDPutParams contains all the parameters to send to the API endpoint
for the subscriberpreferences by Id put operation typically these are written to a http.Request
*/
type SubscriberpreferencesByIDPutParams struct {

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

// WithTimeout adds the timeout to the subscriberpreferences by Id put params
func (o *SubscriberpreferencesByIDPutParams) WithTimeout(timeout time.Duration) *SubscriberpreferencesByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscriberpreferences by Id put params
func (o *SubscriberpreferencesByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscriberpreferences by Id put params
func (o *SubscriberpreferencesByIDPutParams) WithContext(ctx context.Context) *SubscriberpreferencesByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscriberpreferences by Id put params
func (o *SubscriberpreferencesByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscriberpreferences by Id put params
func (o *SubscriberpreferencesByIDPutParams) WithHTTPClient(client *http.Client) *SubscriberpreferencesByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscriberpreferences by Id put params
func (o *SubscriberpreferencesByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the subscriberpreferences by Id put params
func (o *SubscriberpreferencesByIDPutParams) WithBody(body string) *SubscriberpreferencesByIDPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the subscriberpreferences by Id put params
func (o *SubscriberpreferencesByIDPutParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the subscriberpreferences by Id put params
func (o *SubscriberpreferencesByIDPutParams) WithContentType(contentType string) *SubscriberpreferencesByIDPutParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the subscriberpreferences by Id put params
func (o *SubscriberpreferencesByIDPutParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the subscriberpreferences by Id put params
func (o *SubscriberpreferencesByIDPutParams) WithID(id string) *SubscriberpreferencesByIDPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the subscriberpreferences by Id put params
func (o *SubscriberpreferencesByIDPutParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SubscriberpreferencesByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
