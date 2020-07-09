// Code generated by go-swagger; DO NOT EDIT.

package admins

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

// NewCreateanewAdminParams creates a new CreateanewAdminParams object
// with the default values initialized.
func NewCreateanewAdminParams() *CreateanewAdminParams {
	var ()
	return &CreateanewAdminParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewAdminParamsWithTimeout creates a new CreateanewAdminParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewAdminParamsWithTimeout(timeout time.Duration) *CreateanewAdminParams {
	var ()
	return &CreateanewAdminParams{

		timeout: timeout,
	}
}

// NewCreateanewAdminParamsWithContext creates a new CreateanewAdminParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewAdminParamsWithContext(ctx context.Context) *CreateanewAdminParams {
	var ()
	return &CreateanewAdminParams{

		Context: ctx,
	}
}

// NewCreateanewAdminParamsWithHTTPClient creates a new CreateanewAdminParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewAdminParamsWithHTTPClient(client *http.Client) *CreateanewAdminParams {
	var ()
	return &CreateanewAdminParams{
		HTTPClient: client,
	}
}

/*CreateanewAdminParams contains all the parameters to send to the API endpoint
for the createanew admin operation typically these are written to a http.Request
*/
type CreateanewAdminParams struct {

	/*Body*/
	Body *models.CreateanewAdminRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew admin params
func (o *CreateanewAdminParams) WithTimeout(timeout time.Duration) *CreateanewAdminParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew admin params
func (o *CreateanewAdminParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew admin params
func (o *CreateanewAdminParams) WithContext(ctx context.Context) *CreateanewAdminParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew admin params
func (o *CreateanewAdminParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew admin params
func (o *CreateanewAdminParams) WithHTTPClient(client *http.Client) *CreateanewAdminParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew admin params
func (o *CreateanewAdminParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew admin params
func (o *CreateanewAdminParams) WithBody(body *models.CreateanewAdminRequest) *CreateanewAdminParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew admin params
func (o *CreateanewAdminParams) SetBody(body *models.CreateanewAdminRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew admin params
func (o *CreateanewAdminParams) WithContentType(contentType string) *CreateanewAdminParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew admin params
func (o *CreateanewAdminParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewAdminParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
