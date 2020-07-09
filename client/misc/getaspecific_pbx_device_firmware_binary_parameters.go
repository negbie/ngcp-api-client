// Code generated by go-swagger; DO NOT EDIT.

package misc

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

// NewGetaspecificPbxDeviceFirmwareBinaryParams creates a new GetaspecificPbxDeviceFirmwareBinaryParams object
// with the default values initialized.
func NewGetaspecificPbxDeviceFirmwareBinaryParams() *GetaspecificPbxDeviceFirmwareBinaryParams {
	var ()
	return &GetaspecificPbxDeviceFirmwareBinaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificPbxDeviceFirmwareBinaryParamsWithTimeout creates a new GetaspecificPbxDeviceFirmwareBinaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificPbxDeviceFirmwareBinaryParamsWithTimeout(timeout time.Duration) *GetaspecificPbxDeviceFirmwareBinaryParams {
	var ()
	return &GetaspecificPbxDeviceFirmwareBinaryParams{

		timeout: timeout,
	}
}

// NewGetaspecificPbxDeviceFirmwareBinaryParamsWithContext creates a new GetaspecificPbxDeviceFirmwareBinaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificPbxDeviceFirmwareBinaryParamsWithContext(ctx context.Context) *GetaspecificPbxDeviceFirmwareBinaryParams {
	var ()
	return &GetaspecificPbxDeviceFirmwareBinaryParams{

		Context: ctx,
	}
}

// NewGetaspecificPbxDeviceFirmwareBinaryParamsWithHTTPClient creates a new GetaspecificPbxDeviceFirmwareBinaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificPbxDeviceFirmwareBinaryParamsWithHTTPClient(client *http.Client) *GetaspecificPbxDeviceFirmwareBinaryParams {
	var ()
	return &GetaspecificPbxDeviceFirmwareBinaryParams{
		HTTPClient: client,
	}
}

/*GetaspecificPbxDeviceFirmwareBinaryParams contains all the parameters to send to the API endpoint
for the getaspecific pbx device firmware binary operation typically these are written to a http.Request
*/
type GetaspecificPbxDeviceFirmwareBinaryParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific pbx device firmware binary params
func (o *GetaspecificPbxDeviceFirmwareBinaryParams) WithTimeout(timeout time.Duration) *GetaspecificPbxDeviceFirmwareBinaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific pbx device firmware binary params
func (o *GetaspecificPbxDeviceFirmwareBinaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific pbx device firmware binary params
func (o *GetaspecificPbxDeviceFirmwareBinaryParams) WithContext(ctx context.Context) *GetaspecificPbxDeviceFirmwareBinaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific pbx device firmware binary params
func (o *GetaspecificPbxDeviceFirmwareBinaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific pbx device firmware binary params
func (o *GetaspecificPbxDeviceFirmwareBinaryParams) WithHTTPClient(client *http.Client) *GetaspecificPbxDeviceFirmwareBinaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific pbx device firmware binary params
func (o *GetaspecificPbxDeviceFirmwareBinaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific pbx device firmware binary params
func (o *GetaspecificPbxDeviceFirmwareBinaryParams) WithID(id string) *GetaspecificPbxDeviceFirmwareBinaryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific pbx device firmware binary params
func (o *GetaspecificPbxDeviceFirmwareBinaryParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificPbxDeviceFirmwareBinaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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