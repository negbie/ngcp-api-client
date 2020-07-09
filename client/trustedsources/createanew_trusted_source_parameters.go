// Code generated by go-swagger; DO NOT EDIT.

package trustedsources

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

// NewCreateanewTrustedSourceParams creates a new CreateanewTrustedSourceParams object
// with the default values initialized.
func NewCreateanewTrustedSourceParams() *CreateanewTrustedSourceParams {
	var ()
	return &CreateanewTrustedSourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewTrustedSourceParamsWithTimeout creates a new CreateanewTrustedSourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewTrustedSourceParamsWithTimeout(timeout time.Duration) *CreateanewTrustedSourceParams {
	var ()
	return &CreateanewTrustedSourceParams{

		timeout: timeout,
	}
}

// NewCreateanewTrustedSourceParamsWithContext creates a new CreateanewTrustedSourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewTrustedSourceParamsWithContext(ctx context.Context) *CreateanewTrustedSourceParams {
	var ()
	return &CreateanewTrustedSourceParams{

		Context: ctx,
	}
}

// NewCreateanewTrustedSourceParamsWithHTTPClient creates a new CreateanewTrustedSourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewTrustedSourceParamsWithHTTPClient(client *http.Client) *CreateanewTrustedSourceParams {
	var ()
	return &CreateanewTrustedSourceParams{
		HTTPClient: client,
	}
}

/*CreateanewTrustedSourceParams contains all the parameters to send to the API endpoint
for the createanew trusted source operation typically these are written to a http.Request
*/
type CreateanewTrustedSourceParams struct {

	/*Body*/
	Body *models.CreateanewTrustedSourceRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew trusted source params
func (o *CreateanewTrustedSourceParams) WithTimeout(timeout time.Duration) *CreateanewTrustedSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew trusted source params
func (o *CreateanewTrustedSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew trusted source params
func (o *CreateanewTrustedSourceParams) WithContext(ctx context.Context) *CreateanewTrustedSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew trusted source params
func (o *CreateanewTrustedSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew trusted source params
func (o *CreateanewTrustedSourceParams) WithHTTPClient(client *http.Client) *CreateanewTrustedSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew trusted source params
func (o *CreateanewTrustedSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew trusted source params
func (o *CreateanewTrustedSourceParams) WithBody(body *models.CreateanewTrustedSourceRequest) *CreateanewTrustedSourceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew trusted source params
func (o *CreateanewTrustedSourceParams) SetBody(body *models.CreateanewTrustedSourceRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew trusted source params
func (o *CreateanewTrustedSourceParams) WithContentType(contentType string) *CreateanewTrustedSourceParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew trusted source params
func (o *CreateanewTrustedSourceParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewTrustedSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
