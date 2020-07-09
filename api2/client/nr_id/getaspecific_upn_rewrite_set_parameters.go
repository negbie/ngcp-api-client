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

// NewGetaspecificUpnRewriteSetParams creates a new GetaspecificUpnRewriteSetParams object
// with the default values initialized.
func NewGetaspecificUpnRewriteSetParams() *GetaspecificUpnRewriteSetParams {
	var ()
	return &GetaspecificUpnRewriteSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificUpnRewriteSetParamsWithTimeout creates a new GetaspecificUpnRewriteSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificUpnRewriteSetParamsWithTimeout(timeout time.Duration) *GetaspecificUpnRewriteSetParams {
	var ()
	return &GetaspecificUpnRewriteSetParams{

		timeout: timeout,
	}
}

// NewGetaspecificUpnRewriteSetParamsWithContext creates a new GetaspecificUpnRewriteSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificUpnRewriteSetParamsWithContext(ctx context.Context) *GetaspecificUpnRewriteSetParams {
	var ()
	return &GetaspecificUpnRewriteSetParams{

		Context: ctx,
	}
}

// NewGetaspecificUpnRewriteSetParamsWithHTTPClient creates a new GetaspecificUpnRewriteSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificUpnRewriteSetParamsWithHTTPClient(client *http.Client) *GetaspecificUpnRewriteSetParams {
	var ()
	return &GetaspecificUpnRewriteSetParams{
		HTTPClient: client,
	}
}

/*GetaspecificUpnRewriteSetParams contains all the parameters to send to the API endpoint
for the getaspecific upn rewrite set operation typically these are written to a http.Request
*/
type GetaspecificUpnRewriteSetParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific upn rewrite set params
func (o *GetaspecificUpnRewriteSetParams) WithTimeout(timeout time.Duration) *GetaspecificUpnRewriteSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific upn rewrite set params
func (o *GetaspecificUpnRewriteSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific upn rewrite set params
func (o *GetaspecificUpnRewriteSetParams) WithContext(ctx context.Context) *GetaspecificUpnRewriteSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific upn rewrite set params
func (o *GetaspecificUpnRewriteSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific upn rewrite set params
func (o *GetaspecificUpnRewriteSetParams) WithHTTPClient(client *http.Client) *GetaspecificUpnRewriteSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific upn rewrite set params
func (o *GetaspecificUpnRewriteSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific upn rewrite set params
func (o *GetaspecificUpnRewriteSetParams) WithID(id string) *GetaspecificUpnRewriteSetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific upn rewrite set params
func (o *GetaspecificUpnRewriteSetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificUpnRewriteSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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