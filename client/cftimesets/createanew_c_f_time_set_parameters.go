// Code generated by go-swagger; DO NOT EDIT.

package cftimesets

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

	"github.com/negbie/ngcp-api1/models"
)

// NewCreateanewCFTimeSetParams creates a new CreateanewCFTimeSetParams object
// with the default values initialized.
func NewCreateanewCFTimeSetParams() *CreateanewCFTimeSetParams {
	var ()
	return &CreateanewCFTimeSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewCFTimeSetParamsWithTimeout creates a new CreateanewCFTimeSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewCFTimeSetParamsWithTimeout(timeout time.Duration) *CreateanewCFTimeSetParams {
	var ()
	return &CreateanewCFTimeSetParams{

		timeout: timeout,
	}
}

// NewCreateanewCFTimeSetParamsWithContext creates a new CreateanewCFTimeSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewCFTimeSetParamsWithContext(ctx context.Context) *CreateanewCFTimeSetParams {
	var ()
	return &CreateanewCFTimeSetParams{

		Context: ctx,
	}
}

// NewCreateanewCFTimeSetParamsWithHTTPClient creates a new CreateanewCFTimeSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewCFTimeSetParamsWithHTTPClient(client *http.Client) *CreateanewCFTimeSetParams {
	var ()
	return &CreateanewCFTimeSetParams{
		HTTPClient: client,
	}
}

/*CreateanewCFTimeSetParams contains all the parameters to send to the API endpoint
for the createanew c f time set operation typically these are written to a http.Request
*/
type CreateanewCFTimeSetParams struct {

	/*Body*/
	Body *models.CreateanewCFTimeSetRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew c f time set params
func (o *CreateanewCFTimeSetParams) WithTimeout(timeout time.Duration) *CreateanewCFTimeSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew c f time set params
func (o *CreateanewCFTimeSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew c f time set params
func (o *CreateanewCFTimeSetParams) WithContext(ctx context.Context) *CreateanewCFTimeSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew c f time set params
func (o *CreateanewCFTimeSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew c f time set params
func (o *CreateanewCFTimeSetParams) WithHTTPClient(client *http.Client) *CreateanewCFTimeSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew c f time set params
func (o *CreateanewCFTimeSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew c f time set params
func (o *CreateanewCFTimeSetParams) WithBody(body *models.CreateanewCFTimeSetRequest) *CreateanewCFTimeSetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew c f time set params
func (o *CreateanewCFTimeSetParams) SetBody(body *models.CreateanewCFTimeSetRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew c f time set params
func (o *CreateanewCFTimeSetParams) WithContentType(contentType string) *CreateanewCFTimeSetParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew c f time set params
func (o *CreateanewCFTimeSetParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewCFTimeSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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