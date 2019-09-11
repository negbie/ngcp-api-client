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

// NewRemindersByIDPutParams creates a new RemindersByIDPutParams object
// with the default values initialized.
func NewRemindersByIDPutParams() *RemindersByIDPutParams {
	var ()
	return &RemindersByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemindersByIDPutParamsWithTimeout creates a new RemindersByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemindersByIDPutParamsWithTimeout(timeout time.Duration) *RemindersByIDPutParams {
	var ()
	return &RemindersByIDPutParams{

		timeout: timeout,
	}
}

// NewRemindersByIDPutParamsWithContext creates a new RemindersByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemindersByIDPutParamsWithContext(ctx context.Context) *RemindersByIDPutParams {
	var ()
	return &RemindersByIDPutParams{

		Context: ctx,
	}
}

// NewRemindersByIDPutParamsWithHTTPClient creates a new RemindersByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemindersByIDPutParamsWithHTTPClient(client *http.Client) *RemindersByIDPutParams {
	var ()
	return &RemindersByIDPutParams{
		HTTPClient: client,
	}
}

/*RemindersByIDPutParams contains all the parameters to send to the API endpoint
for the reminders by Id put operation typically these are written to a http.Request
*/
type RemindersByIDPutParams struct {

	/*Body*/
	Body *Replace1changeaspecificReminderRequest
	/*ContentType*/
	ContentType string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the reminders by Id put params
func (o *RemindersByIDPutParams) WithTimeout(timeout time.Duration) *RemindersByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reminders by Id put params
func (o *RemindersByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reminders by Id put params
func (o *RemindersByIDPutParams) WithContext(ctx context.Context) *RemindersByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reminders by Id put params
func (o *RemindersByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reminders by Id put params
func (o *RemindersByIDPutParams) WithHTTPClient(client *http.Client) *RemindersByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reminders by Id put params
func (o *RemindersByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the reminders by Id put params
func (o *RemindersByIDPutParams) WithBody(body *Replace1changeaspecificReminderRequest) *RemindersByIDPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the reminders by Id put params
func (o *RemindersByIDPutParams) SetBody(body *Replace1changeaspecificReminderRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the reminders by Id put params
func (o *RemindersByIDPutParams) WithContentType(contentType string) *RemindersByIDPutParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the reminders by Id put params
func (o *RemindersByIDPutParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the reminders by Id put params
func (o *RemindersByIDPutParams) WithID(id string) *RemindersByIDPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the reminders by Id put params
func (o *RemindersByIDPutParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RemindersByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
