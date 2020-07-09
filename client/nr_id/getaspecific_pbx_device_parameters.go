// Code generated by go-swagger; DO NOT EDIT.

package nr_id

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

// NewGetaspecificPbxDeviceParams creates a new GetaspecificPbxDeviceParams object
// with the default values initialized.
func NewGetaspecificPbxDeviceParams() *GetaspecificPbxDeviceParams {
	var ()
	return &GetaspecificPbxDeviceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificPbxDeviceParamsWithTimeout creates a new GetaspecificPbxDeviceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificPbxDeviceParamsWithTimeout(timeout time.Duration) *GetaspecificPbxDeviceParams {
	var ()
	return &GetaspecificPbxDeviceParams{

		timeout: timeout,
	}
}

// NewGetaspecificPbxDeviceParamsWithContext creates a new GetaspecificPbxDeviceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificPbxDeviceParamsWithContext(ctx context.Context) *GetaspecificPbxDeviceParams {
	var ()
	return &GetaspecificPbxDeviceParams{

		Context: ctx,
	}
}

// NewGetaspecificPbxDeviceParamsWithHTTPClient creates a new GetaspecificPbxDeviceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificPbxDeviceParamsWithHTTPClient(client *http.Client) *GetaspecificPbxDeviceParams {
	var ()
	return &GetaspecificPbxDeviceParams{
		HTTPClient: client,
	}
}

/*GetaspecificPbxDeviceParams contains all the parameters to send to the API endpoint
for the getaspecific pbx device operation typically these are written to a http.Request
*/
type GetaspecificPbxDeviceParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific pbx device params
func (o *GetaspecificPbxDeviceParams) WithTimeout(timeout time.Duration) *GetaspecificPbxDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific pbx device params
func (o *GetaspecificPbxDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific pbx device params
func (o *GetaspecificPbxDeviceParams) WithContext(ctx context.Context) *GetaspecificPbxDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific pbx device params
func (o *GetaspecificPbxDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific pbx device params
func (o *GetaspecificPbxDeviceParams) WithHTTPClient(client *http.Client) *GetaspecificPbxDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific pbx device params
func (o *GetaspecificPbxDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific pbx device params
func (o *GetaspecificPbxDeviceParams) WithID(id string) *GetaspecificPbxDeviceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific pbx device params
func (o *GetaspecificPbxDeviceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificPbxDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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