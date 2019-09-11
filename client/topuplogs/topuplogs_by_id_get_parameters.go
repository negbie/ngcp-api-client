// Code generated by go-swagger; DO NOT EDIT.

package topuplogs

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

// NewTopuplogsByIDGetParams creates a new TopuplogsByIDGetParams object
// with the default values initialized.
func NewTopuplogsByIDGetParams() *TopuplogsByIDGetParams {
	var ()
	return &TopuplogsByIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTopuplogsByIDGetParamsWithTimeout creates a new TopuplogsByIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTopuplogsByIDGetParamsWithTimeout(timeout time.Duration) *TopuplogsByIDGetParams {
	var ()
	return &TopuplogsByIDGetParams{

		timeout: timeout,
	}
}

// NewTopuplogsByIDGetParamsWithContext creates a new TopuplogsByIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewTopuplogsByIDGetParamsWithContext(ctx context.Context) *TopuplogsByIDGetParams {
	var ()
	return &TopuplogsByIDGetParams{

		Context: ctx,
	}
}

// NewTopuplogsByIDGetParamsWithHTTPClient creates a new TopuplogsByIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTopuplogsByIDGetParamsWithHTTPClient(client *http.Client) *TopuplogsByIDGetParams {
	var ()
	return &TopuplogsByIDGetParams{
		HTTPClient: client,
	}
}

/*TopuplogsByIDGetParams contains all the parameters to send to the API endpoint
for the topuplogs by Id get operation typically these are written to a http.Request
*/
type TopuplogsByIDGetParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the topuplogs by Id get params
func (o *TopuplogsByIDGetParams) WithTimeout(timeout time.Duration) *TopuplogsByIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the topuplogs by Id get params
func (o *TopuplogsByIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the topuplogs by Id get params
func (o *TopuplogsByIDGetParams) WithContext(ctx context.Context) *TopuplogsByIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the topuplogs by Id get params
func (o *TopuplogsByIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the topuplogs by Id get params
func (o *TopuplogsByIDGetParams) WithHTTPClient(client *http.Client) *TopuplogsByIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the topuplogs by Id get params
func (o *TopuplogsByIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the topuplogs by Id get params
func (o *TopuplogsByIDGetParams) WithID(id string) *TopuplogsByIDGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the topuplogs by Id get params
func (o *TopuplogsByIDGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TopuplogsByIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
