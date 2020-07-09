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

// NewGetaspecificBillingNetworkParams creates a new GetaspecificBillingNetworkParams object
// with the default values initialized.
func NewGetaspecificBillingNetworkParams() *GetaspecificBillingNetworkParams {
	var ()
	return &GetaspecificBillingNetworkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificBillingNetworkParamsWithTimeout creates a new GetaspecificBillingNetworkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificBillingNetworkParamsWithTimeout(timeout time.Duration) *GetaspecificBillingNetworkParams {
	var ()
	return &GetaspecificBillingNetworkParams{

		timeout: timeout,
	}
}

// NewGetaspecificBillingNetworkParamsWithContext creates a new GetaspecificBillingNetworkParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificBillingNetworkParamsWithContext(ctx context.Context) *GetaspecificBillingNetworkParams {
	var ()
	return &GetaspecificBillingNetworkParams{

		Context: ctx,
	}
}

// NewGetaspecificBillingNetworkParamsWithHTTPClient creates a new GetaspecificBillingNetworkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificBillingNetworkParamsWithHTTPClient(client *http.Client) *GetaspecificBillingNetworkParams {
	var ()
	return &GetaspecificBillingNetworkParams{
		HTTPClient: client,
	}
}

/*GetaspecificBillingNetworkParams contains all the parameters to send to the API endpoint
for the getaspecific billing network operation typically these are written to a http.Request
*/
type GetaspecificBillingNetworkParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific billing network params
func (o *GetaspecificBillingNetworkParams) WithTimeout(timeout time.Duration) *GetaspecificBillingNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific billing network params
func (o *GetaspecificBillingNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific billing network params
func (o *GetaspecificBillingNetworkParams) WithContext(ctx context.Context) *GetaspecificBillingNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific billing network params
func (o *GetaspecificBillingNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific billing network params
func (o *GetaspecificBillingNetworkParams) WithHTTPClient(client *http.Client) *GetaspecificBillingNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific billing network params
func (o *GetaspecificBillingNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific billing network params
func (o *GetaspecificBillingNetworkParams) WithID(id string) *GetaspecificBillingNetworkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific billing network params
func (o *GetaspecificBillingNetworkParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificBillingNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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