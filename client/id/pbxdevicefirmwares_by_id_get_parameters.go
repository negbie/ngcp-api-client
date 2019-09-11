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

// NewPbxdevicefirmwaresByIDGetParams creates a new PbxdevicefirmwaresByIDGetParams object
// with the default values initialized.
func NewPbxdevicefirmwaresByIDGetParams() *PbxdevicefirmwaresByIDGetParams {
	var ()
	return &PbxdevicefirmwaresByIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPbxdevicefirmwaresByIDGetParamsWithTimeout creates a new PbxdevicefirmwaresByIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPbxdevicefirmwaresByIDGetParamsWithTimeout(timeout time.Duration) *PbxdevicefirmwaresByIDGetParams {
	var ()
	return &PbxdevicefirmwaresByIDGetParams{

		timeout: timeout,
	}
}

// NewPbxdevicefirmwaresByIDGetParamsWithContext creates a new PbxdevicefirmwaresByIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPbxdevicefirmwaresByIDGetParamsWithContext(ctx context.Context) *PbxdevicefirmwaresByIDGetParams {
	var ()
	return &PbxdevicefirmwaresByIDGetParams{

		Context: ctx,
	}
}

// NewPbxdevicefirmwaresByIDGetParamsWithHTTPClient creates a new PbxdevicefirmwaresByIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPbxdevicefirmwaresByIDGetParamsWithHTTPClient(client *http.Client) *PbxdevicefirmwaresByIDGetParams {
	var ()
	return &PbxdevicefirmwaresByIDGetParams{
		HTTPClient: client,
	}
}

/*PbxdevicefirmwaresByIDGetParams contains all the parameters to send to the API endpoint
for the pbxdevicefirmwares by Id get operation typically these are written to a http.Request
*/
type PbxdevicefirmwaresByIDGetParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pbxdevicefirmwares by Id get params
func (o *PbxdevicefirmwaresByIDGetParams) WithTimeout(timeout time.Duration) *PbxdevicefirmwaresByIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pbxdevicefirmwares by Id get params
func (o *PbxdevicefirmwaresByIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pbxdevicefirmwares by Id get params
func (o *PbxdevicefirmwaresByIDGetParams) WithContext(ctx context.Context) *PbxdevicefirmwaresByIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pbxdevicefirmwares by Id get params
func (o *PbxdevicefirmwaresByIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pbxdevicefirmwares by Id get params
func (o *PbxdevicefirmwaresByIDGetParams) WithHTTPClient(client *http.Client) *PbxdevicefirmwaresByIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pbxdevicefirmwares by Id get params
func (o *PbxdevicefirmwaresByIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the pbxdevicefirmwares by Id get params
func (o *PbxdevicefirmwaresByIDGetParams) WithID(id string) *PbxdevicefirmwaresByIDGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the pbxdevicefirmwares by Id get params
func (o *PbxdevicefirmwaresByIDGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PbxdevicefirmwaresByIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
