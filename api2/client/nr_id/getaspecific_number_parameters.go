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

// NewGetaspecificNumberParams creates a new GetaspecificNumberParams object
// with the default values initialized.
func NewGetaspecificNumberParams() *GetaspecificNumberParams {
	var ()
	return &GetaspecificNumberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificNumberParamsWithTimeout creates a new GetaspecificNumberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificNumberParamsWithTimeout(timeout time.Duration) *GetaspecificNumberParams {
	var ()
	return &GetaspecificNumberParams{

		timeout: timeout,
	}
}

// NewGetaspecificNumberParamsWithContext creates a new GetaspecificNumberParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificNumberParamsWithContext(ctx context.Context) *GetaspecificNumberParams {
	var ()
	return &GetaspecificNumberParams{

		Context: ctx,
	}
}

// NewGetaspecificNumberParamsWithHTTPClient creates a new GetaspecificNumberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificNumberParamsWithHTTPClient(client *http.Client) *GetaspecificNumberParams {
	var ()
	return &GetaspecificNumberParams{
		HTTPClient: client,
	}
}

/*GetaspecificNumberParams contains all the parameters to send to the API endpoint
for the getaspecific number operation typically these are written to a http.Request
*/
type GetaspecificNumberParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific number params
func (o *GetaspecificNumberParams) WithTimeout(timeout time.Duration) *GetaspecificNumberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific number params
func (o *GetaspecificNumberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific number params
func (o *GetaspecificNumberParams) WithContext(ctx context.Context) *GetaspecificNumberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific number params
func (o *GetaspecificNumberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific number params
func (o *GetaspecificNumberParams) WithHTTPClient(client *http.Client) *GetaspecificNumberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific number params
func (o *GetaspecificNumberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific number params
func (o *GetaspecificNumberParams) WithID(id string) *GetaspecificNumberParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific number params
func (o *GetaspecificNumberParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificNumberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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