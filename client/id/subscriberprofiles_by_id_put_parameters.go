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

// NewSubscriberprofilesByIDPutParams creates a new SubscriberprofilesByIDPutParams object
// with the default values initialized.
func NewSubscriberprofilesByIDPutParams() *SubscriberprofilesByIDPutParams {
	var ()
	return &SubscriberprofilesByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubscriberprofilesByIDPutParamsWithTimeout creates a new SubscriberprofilesByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubscriberprofilesByIDPutParamsWithTimeout(timeout time.Duration) *SubscriberprofilesByIDPutParams {
	var ()
	return &SubscriberprofilesByIDPutParams{

		timeout: timeout,
	}
}

// NewSubscriberprofilesByIDPutParamsWithContext creates a new SubscriberprofilesByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubscriberprofilesByIDPutParamsWithContext(ctx context.Context) *SubscriberprofilesByIDPutParams {
	var ()
	return &SubscriberprofilesByIDPutParams{

		Context: ctx,
	}
}

// NewSubscriberprofilesByIDPutParamsWithHTTPClient creates a new SubscriberprofilesByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSubscriberprofilesByIDPutParamsWithHTTPClient(client *http.Client) *SubscriberprofilesByIDPutParams {
	var ()
	return &SubscriberprofilesByIDPutParams{
		HTTPClient: client,
	}
}

/*SubscriberprofilesByIDPutParams contains all the parameters to send to the API endpoint
for the subscriberprofiles by Id put operation typically these are written to a http.Request
*/
type SubscriberprofilesByIDPutParams struct {

	/*Body*/
	Body *Replace1changeaspecificSubscriberProfileRequest
	/*ContentType*/
	ContentType string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the subscriberprofiles by Id put params
func (o *SubscriberprofilesByIDPutParams) WithTimeout(timeout time.Duration) *SubscriberprofilesByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscriberprofiles by Id put params
func (o *SubscriberprofilesByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscriberprofiles by Id put params
func (o *SubscriberprofilesByIDPutParams) WithContext(ctx context.Context) *SubscriberprofilesByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscriberprofiles by Id put params
func (o *SubscriberprofilesByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscriberprofiles by Id put params
func (o *SubscriberprofilesByIDPutParams) WithHTTPClient(client *http.Client) *SubscriberprofilesByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscriberprofiles by Id put params
func (o *SubscriberprofilesByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the subscriberprofiles by Id put params
func (o *SubscriberprofilesByIDPutParams) WithBody(body *Replace1changeaspecificSubscriberProfileRequest) *SubscriberprofilesByIDPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the subscriberprofiles by Id put params
func (o *SubscriberprofilesByIDPutParams) SetBody(body *Replace1changeaspecificSubscriberProfileRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the subscriberprofiles by Id put params
func (o *SubscriberprofilesByIDPutParams) WithContentType(contentType string) *SubscriberprofilesByIDPutParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the subscriberprofiles by Id put params
func (o *SubscriberprofilesByIDPutParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the subscriberprofiles by Id put params
func (o *SubscriberprofilesByIDPutParams) WithID(id string) *SubscriberprofilesByIDPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the subscriberprofiles by Id put params
func (o *SubscriberprofilesByIDPutParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SubscriberprofilesByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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