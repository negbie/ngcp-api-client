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

// NewPreferencesmetaentriesByIDGetParams creates a new PreferencesmetaentriesByIDGetParams object
// with the default values initialized.
func NewPreferencesmetaentriesByIDGetParams() *PreferencesmetaentriesByIDGetParams {
	var ()
	return &PreferencesmetaentriesByIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPreferencesmetaentriesByIDGetParamsWithTimeout creates a new PreferencesmetaentriesByIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPreferencesmetaentriesByIDGetParamsWithTimeout(timeout time.Duration) *PreferencesmetaentriesByIDGetParams {
	var ()
	return &PreferencesmetaentriesByIDGetParams{

		timeout: timeout,
	}
}

// NewPreferencesmetaentriesByIDGetParamsWithContext creates a new PreferencesmetaentriesByIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPreferencesmetaentriesByIDGetParamsWithContext(ctx context.Context) *PreferencesmetaentriesByIDGetParams {
	var ()
	return &PreferencesmetaentriesByIDGetParams{

		Context: ctx,
	}
}

// NewPreferencesmetaentriesByIDGetParamsWithHTTPClient creates a new PreferencesmetaentriesByIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPreferencesmetaentriesByIDGetParamsWithHTTPClient(client *http.Client) *PreferencesmetaentriesByIDGetParams {
	var ()
	return &PreferencesmetaentriesByIDGetParams{
		HTTPClient: client,
	}
}

/*PreferencesmetaentriesByIDGetParams contains all the parameters to send to the API endpoint
for the preferencesmetaentries by Id get operation typically these are written to a http.Request
*/
type PreferencesmetaentriesByIDGetParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the preferencesmetaentries by Id get params
func (o *PreferencesmetaentriesByIDGetParams) WithTimeout(timeout time.Duration) *PreferencesmetaentriesByIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the preferencesmetaentries by Id get params
func (o *PreferencesmetaentriesByIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the preferencesmetaentries by Id get params
func (o *PreferencesmetaentriesByIDGetParams) WithContext(ctx context.Context) *PreferencesmetaentriesByIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the preferencesmetaentries by Id get params
func (o *PreferencesmetaentriesByIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the preferencesmetaentries by Id get params
func (o *PreferencesmetaentriesByIDGetParams) WithHTTPClient(client *http.Client) *PreferencesmetaentriesByIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the preferencesmetaentries by Id get params
func (o *PreferencesmetaentriesByIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the preferencesmetaentries by Id get params
func (o *PreferencesmetaentriesByIDGetParams) WithID(id string) *PreferencesmetaentriesByIDGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the preferencesmetaentries by Id get params
func (o *PreferencesmetaentriesByIDGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PreferencesmetaentriesByIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
