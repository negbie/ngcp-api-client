// Code generated by go-swagger; DO NOT EDIT.

package customerzonecosts

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

// NewGetaspecificCustomerZoneCostParams creates a new GetaspecificCustomerZoneCostParams object
// with the default values initialized.
func NewGetaspecificCustomerZoneCostParams() *GetaspecificCustomerZoneCostParams {
	var ()
	return &GetaspecificCustomerZoneCostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificCustomerZoneCostParamsWithTimeout creates a new GetaspecificCustomerZoneCostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificCustomerZoneCostParamsWithTimeout(timeout time.Duration) *GetaspecificCustomerZoneCostParams {
	var ()
	return &GetaspecificCustomerZoneCostParams{

		timeout: timeout,
	}
}

// NewGetaspecificCustomerZoneCostParamsWithContext creates a new GetaspecificCustomerZoneCostParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificCustomerZoneCostParamsWithContext(ctx context.Context) *GetaspecificCustomerZoneCostParams {
	var ()
	return &GetaspecificCustomerZoneCostParams{

		Context: ctx,
	}
}

// NewGetaspecificCustomerZoneCostParamsWithHTTPClient creates a new GetaspecificCustomerZoneCostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificCustomerZoneCostParamsWithHTTPClient(client *http.Client) *GetaspecificCustomerZoneCostParams {
	var ()
	return &GetaspecificCustomerZoneCostParams{
		HTTPClient: client,
	}
}

/*GetaspecificCustomerZoneCostParams contains all the parameters to send to the API endpoint
for the getaspecific customer zone cost operation typically these are written to a http.Request
*/
type GetaspecificCustomerZoneCostParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific customer zone cost params
func (o *GetaspecificCustomerZoneCostParams) WithTimeout(timeout time.Duration) *GetaspecificCustomerZoneCostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific customer zone cost params
func (o *GetaspecificCustomerZoneCostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific customer zone cost params
func (o *GetaspecificCustomerZoneCostParams) WithContext(ctx context.Context) *GetaspecificCustomerZoneCostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific customer zone cost params
func (o *GetaspecificCustomerZoneCostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific customer zone cost params
func (o *GetaspecificCustomerZoneCostParams) WithHTTPClient(client *http.Client) *GetaspecificCustomerZoneCostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific customer zone cost params
func (o *GetaspecificCustomerZoneCostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific customer zone cost params
func (o *GetaspecificCustomerZoneCostParams) WithID(id string) *GetaspecificCustomerZoneCostParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific customer zone cost params
func (o *GetaspecificCustomerZoneCostParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificCustomerZoneCostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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