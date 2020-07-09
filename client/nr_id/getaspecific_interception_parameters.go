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

// NewGetaspecificInterceptionParams creates a new GetaspecificInterceptionParams object
// with the default values initialized.
func NewGetaspecificInterceptionParams() *GetaspecificInterceptionParams {
	var ()
	return &GetaspecificInterceptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificInterceptionParamsWithTimeout creates a new GetaspecificInterceptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificInterceptionParamsWithTimeout(timeout time.Duration) *GetaspecificInterceptionParams {
	var ()
	return &GetaspecificInterceptionParams{

		timeout: timeout,
	}
}

// NewGetaspecificInterceptionParamsWithContext creates a new GetaspecificInterceptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificInterceptionParamsWithContext(ctx context.Context) *GetaspecificInterceptionParams {
	var ()
	return &GetaspecificInterceptionParams{

		Context: ctx,
	}
}

// NewGetaspecificInterceptionParamsWithHTTPClient creates a new GetaspecificInterceptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificInterceptionParamsWithHTTPClient(client *http.Client) *GetaspecificInterceptionParams {
	var ()
	return &GetaspecificInterceptionParams{
		HTTPClient: client,
	}
}

/*GetaspecificInterceptionParams contains all the parameters to send to the API endpoint
for the getaspecific interception operation typically these are written to a http.Request
*/
type GetaspecificInterceptionParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific interception params
func (o *GetaspecificInterceptionParams) WithTimeout(timeout time.Duration) *GetaspecificInterceptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific interception params
func (o *GetaspecificInterceptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific interception params
func (o *GetaspecificInterceptionParams) WithContext(ctx context.Context) *GetaspecificInterceptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific interception params
func (o *GetaspecificInterceptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific interception params
func (o *GetaspecificInterceptionParams) WithHTTPClient(client *http.Client) *GetaspecificInterceptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific interception params
func (o *GetaspecificInterceptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific interception params
func (o *GetaspecificInterceptionParams) WithID(id string) *GetaspecificInterceptionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific interception params
func (o *GetaspecificInterceptionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificInterceptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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