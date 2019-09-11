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

// NewMailtofaxsettingsByIDGetParams creates a new MailtofaxsettingsByIDGetParams object
// with the default values initialized.
func NewMailtofaxsettingsByIDGetParams() *MailtofaxsettingsByIDGetParams {
	var ()
	return &MailtofaxsettingsByIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMailtofaxsettingsByIDGetParamsWithTimeout creates a new MailtofaxsettingsByIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMailtofaxsettingsByIDGetParamsWithTimeout(timeout time.Duration) *MailtofaxsettingsByIDGetParams {
	var ()
	return &MailtofaxsettingsByIDGetParams{

		timeout: timeout,
	}
}

// NewMailtofaxsettingsByIDGetParamsWithContext creates a new MailtofaxsettingsByIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewMailtofaxsettingsByIDGetParamsWithContext(ctx context.Context) *MailtofaxsettingsByIDGetParams {
	var ()
	return &MailtofaxsettingsByIDGetParams{

		Context: ctx,
	}
}

// NewMailtofaxsettingsByIDGetParamsWithHTTPClient creates a new MailtofaxsettingsByIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMailtofaxsettingsByIDGetParamsWithHTTPClient(client *http.Client) *MailtofaxsettingsByIDGetParams {
	var ()
	return &MailtofaxsettingsByIDGetParams{
		HTTPClient: client,
	}
}

/*MailtofaxsettingsByIDGetParams contains all the parameters to send to the API endpoint
for the mailtofaxsettings by Id get operation typically these are written to a http.Request
*/
type MailtofaxsettingsByIDGetParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the mailtofaxsettings by Id get params
func (o *MailtofaxsettingsByIDGetParams) WithTimeout(timeout time.Duration) *MailtofaxsettingsByIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the mailtofaxsettings by Id get params
func (o *MailtofaxsettingsByIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the mailtofaxsettings by Id get params
func (o *MailtofaxsettingsByIDGetParams) WithContext(ctx context.Context) *MailtofaxsettingsByIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the mailtofaxsettings by Id get params
func (o *MailtofaxsettingsByIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the mailtofaxsettings by Id get params
func (o *MailtofaxsettingsByIDGetParams) WithHTTPClient(client *http.Client) *MailtofaxsettingsByIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the mailtofaxsettings by Id get params
func (o *MailtofaxsettingsByIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the mailtofaxsettings by Id get params
func (o *MailtofaxsettingsByIDGetParams) WithID(id string) *MailtofaxsettingsByIDGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the mailtofaxsettings by Id get params
func (o *MailtofaxsettingsByIDGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *MailtofaxsettingsByIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}