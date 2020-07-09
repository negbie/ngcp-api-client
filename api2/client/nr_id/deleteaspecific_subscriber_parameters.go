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

// NewDeleteaspecificSubscriberParams creates a new DeleteaspecificSubscriberParams object
// with the default values initialized.
func NewDeleteaspecificSubscriberParams() *DeleteaspecificSubscriberParams {
	var ()
	return &DeleteaspecificSubscriberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteaspecificSubscriberParamsWithTimeout creates a new DeleteaspecificSubscriberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteaspecificSubscriberParamsWithTimeout(timeout time.Duration) *DeleteaspecificSubscriberParams {
	var ()
	return &DeleteaspecificSubscriberParams{

		timeout: timeout,
	}
}

// NewDeleteaspecificSubscriberParamsWithContext creates a new DeleteaspecificSubscriberParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteaspecificSubscriberParamsWithContext(ctx context.Context) *DeleteaspecificSubscriberParams {
	var ()
	return &DeleteaspecificSubscriberParams{

		Context: ctx,
	}
}

// NewDeleteaspecificSubscriberParamsWithHTTPClient creates a new DeleteaspecificSubscriberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteaspecificSubscriberParamsWithHTTPClient(client *http.Client) *DeleteaspecificSubscriberParams {
	var ()
	return &DeleteaspecificSubscriberParams{
		HTTPClient: client,
	}
}

/*DeleteaspecificSubscriberParams contains all the parameters to send to the API endpoint
for the deleteaspecific subscriber operation typically these are written to a http.Request
*/
type DeleteaspecificSubscriberParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the deleteaspecific subscriber params
func (o *DeleteaspecificSubscriberParams) WithTimeout(timeout time.Duration) *DeleteaspecificSubscriberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deleteaspecific subscriber params
func (o *DeleteaspecificSubscriberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deleteaspecific subscriber params
func (o *DeleteaspecificSubscriberParams) WithContext(ctx context.Context) *DeleteaspecificSubscriberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deleteaspecific subscriber params
func (o *DeleteaspecificSubscriberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deleteaspecific subscriber params
func (o *DeleteaspecificSubscriberParams) WithHTTPClient(client *http.Client) *DeleteaspecificSubscriberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deleteaspecific subscriber params
func (o *DeleteaspecificSubscriberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the deleteaspecific subscriber params
func (o *DeleteaspecificSubscriberParams) WithID(id string) *DeleteaspecificSubscriberParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the deleteaspecific subscriber params
func (o *DeleteaspecificSubscriberParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteaspecificSubscriberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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