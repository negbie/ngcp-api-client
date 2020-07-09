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

// NewGetaspecificSoundSetParams creates a new GetaspecificSoundSetParams object
// with the default values initialized.
func NewGetaspecificSoundSetParams() *GetaspecificSoundSetParams {
	var ()
	return &GetaspecificSoundSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificSoundSetParamsWithTimeout creates a new GetaspecificSoundSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificSoundSetParamsWithTimeout(timeout time.Duration) *GetaspecificSoundSetParams {
	var ()
	return &GetaspecificSoundSetParams{

		timeout: timeout,
	}
}

// NewGetaspecificSoundSetParamsWithContext creates a new GetaspecificSoundSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificSoundSetParamsWithContext(ctx context.Context) *GetaspecificSoundSetParams {
	var ()
	return &GetaspecificSoundSetParams{

		Context: ctx,
	}
}

// NewGetaspecificSoundSetParamsWithHTTPClient creates a new GetaspecificSoundSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificSoundSetParamsWithHTTPClient(client *http.Client) *GetaspecificSoundSetParams {
	var ()
	return &GetaspecificSoundSetParams{
		HTTPClient: client,
	}
}

/*GetaspecificSoundSetParams contains all the parameters to send to the API endpoint
for the getaspecific sound set operation typically these are written to a http.Request
*/
type GetaspecificSoundSetParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific sound set params
func (o *GetaspecificSoundSetParams) WithTimeout(timeout time.Duration) *GetaspecificSoundSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific sound set params
func (o *GetaspecificSoundSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific sound set params
func (o *GetaspecificSoundSetParams) WithContext(ctx context.Context) *GetaspecificSoundSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific sound set params
func (o *GetaspecificSoundSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific sound set params
func (o *GetaspecificSoundSetParams) WithHTTPClient(client *http.Client) *GetaspecificSoundSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific sound set params
func (o *GetaspecificSoundSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific sound set params
func (o *GetaspecificSoundSetParams) WithID(id string) *GetaspecificSoundSetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific sound set params
func (o *GetaspecificSoundSetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificSoundSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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