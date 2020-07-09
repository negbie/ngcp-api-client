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

// NewGetaspecificBillingZoneParams creates a new GetaspecificBillingZoneParams object
// with the default values initialized.
func NewGetaspecificBillingZoneParams() *GetaspecificBillingZoneParams {
	var ()
	return &GetaspecificBillingZoneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificBillingZoneParamsWithTimeout creates a new GetaspecificBillingZoneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificBillingZoneParamsWithTimeout(timeout time.Duration) *GetaspecificBillingZoneParams {
	var ()
	return &GetaspecificBillingZoneParams{

		timeout: timeout,
	}
}

// NewGetaspecificBillingZoneParamsWithContext creates a new GetaspecificBillingZoneParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificBillingZoneParamsWithContext(ctx context.Context) *GetaspecificBillingZoneParams {
	var ()
	return &GetaspecificBillingZoneParams{

		Context: ctx,
	}
}

// NewGetaspecificBillingZoneParamsWithHTTPClient creates a new GetaspecificBillingZoneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificBillingZoneParamsWithHTTPClient(client *http.Client) *GetaspecificBillingZoneParams {
	var ()
	return &GetaspecificBillingZoneParams{
		HTTPClient: client,
	}
}

/*GetaspecificBillingZoneParams contains all the parameters to send to the API endpoint
for the getaspecific billing zone operation typically these are written to a http.Request
*/
type GetaspecificBillingZoneParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific billing zone params
func (o *GetaspecificBillingZoneParams) WithTimeout(timeout time.Duration) *GetaspecificBillingZoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific billing zone params
func (o *GetaspecificBillingZoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific billing zone params
func (o *GetaspecificBillingZoneParams) WithContext(ctx context.Context) *GetaspecificBillingZoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific billing zone params
func (o *GetaspecificBillingZoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific billing zone params
func (o *GetaspecificBillingZoneParams) WithHTTPClient(client *http.Client) *GetaspecificBillingZoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific billing zone params
func (o *GetaspecificBillingZoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific billing zone params
func (o *GetaspecificBillingZoneParams) WithID(id string) *GetaspecificBillingZoneParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific billing zone params
func (o *GetaspecificBillingZoneParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificBillingZoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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