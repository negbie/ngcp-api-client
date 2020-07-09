// Code generated by go-swagger; DO NOT EDIT.

package cfsourcesets

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

	"github.com/negbie/ngcp-api/models"
)

// NewCreateanewCFSourceSetParams creates a new CreateanewCFSourceSetParams object
// with the default values initialized.
func NewCreateanewCFSourceSetParams() *CreateanewCFSourceSetParams {
	var ()
	return &CreateanewCFSourceSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewCFSourceSetParamsWithTimeout creates a new CreateanewCFSourceSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewCFSourceSetParamsWithTimeout(timeout time.Duration) *CreateanewCFSourceSetParams {
	var ()
	return &CreateanewCFSourceSetParams{

		timeout: timeout,
	}
}

// NewCreateanewCFSourceSetParamsWithContext creates a new CreateanewCFSourceSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewCFSourceSetParamsWithContext(ctx context.Context) *CreateanewCFSourceSetParams {
	var ()
	return &CreateanewCFSourceSetParams{

		Context: ctx,
	}
}

// NewCreateanewCFSourceSetParamsWithHTTPClient creates a new CreateanewCFSourceSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewCFSourceSetParamsWithHTTPClient(client *http.Client) *CreateanewCFSourceSetParams {
	var ()
	return &CreateanewCFSourceSetParams{
		HTTPClient: client,
	}
}

/*CreateanewCFSourceSetParams contains all the parameters to send to the API endpoint
for the createanew c f source set operation typically these are written to a http.Request
*/
type CreateanewCFSourceSetParams struct {

	/*Body*/
	Body *models.CreateanewCFSourceSetRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew c f source set params
func (o *CreateanewCFSourceSetParams) WithTimeout(timeout time.Duration) *CreateanewCFSourceSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew c f source set params
func (o *CreateanewCFSourceSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew c f source set params
func (o *CreateanewCFSourceSetParams) WithContext(ctx context.Context) *CreateanewCFSourceSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew c f source set params
func (o *CreateanewCFSourceSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew c f source set params
func (o *CreateanewCFSourceSetParams) WithHTTPClient(client *http.Client) *CreateanewCFSourceSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew c f source set params
func (o *CreateanewCFSourceSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew c f source set params
func (o *CreateanewCFSourceSetParams) WithBody(body *models.CreateanewCFSourceSetRequest) *CreateanewCFSourceSetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew c f source set params
func (o *CreateanewCFSourceSetParams) SetBody(body *models.CreateanewCFSourceSetRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew c f source set params
func (o *CreateanewCFSourceSetParams) WithContentType(contentType string) *CreateanewCFSourceSetParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew c f source set params
func (o *CreateanewCFSourceSetParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewCFSourceSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// header param Content-Type
	if err := r.SetHeaderParam("Content-Type", o.ContentType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
