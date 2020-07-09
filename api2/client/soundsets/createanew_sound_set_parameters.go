// Code generated by go-swagger; DO NOT EDIT.

package soundsets

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

// NewCreateanewSoundSetParams creates a new CreateanewSoundSetParams object
// with the default values initialized.
func NewCreateanewSoundSetParams() *CreateanewSoundSetParams {
	var ()
	return &CreateanewSoundSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewSoundSetParamsWithTimeout creates a new CreateanewSoundSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewSoundSetParamsWithTimeout(timeout time.Duration) *CreateanewSoundSetParams {
	var ()
	return &CreateanewSoundSetParams{

		timeout: timeout,
	}
}

// NewCreateanewSoundSetParamsWithContext creates a new CreateanewSoundSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewSoundSetParamsWithContext(ctx context.Context) *CreateanewSoundSetParams {
	var ()
	return &CreateanewSoundSetParams{

		Context: ctx,
	}
}

// NewCreateanewSoundSetParamsWithHTTPClient creates a new CreateanewSoundSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewSoundSetParamsWithHTTPClient(client *http.Client) *CreateanewSoundSetParams {
	var ()
	return &CreateanewSoundSetParams{
		HTTPClient: client,
	}
}

/*CreateanewSoundSetParams contains all the parameters to send to the API endpoint
for the createanew sound set operation typically these are written to a http.Request
*/
type CreateanewSoundSetParams struct {

	/*Body*/
	Body *models.CreateanewSoundSetRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew sound set params
func (o *CreateanewSoundSetParams) WithTimeout(timeout time.Duration) *CreateanewSoundSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew sound set params
func (o *CreateanewSoundSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew sound set params
func (o *CreateanewSoundSetParams) WithContext(ctx context.Context) *CreateanewSoundSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew sound set params
func (o *CreateanewSoundSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew sound set params
func (o *CreateanewSoundSetParams) WithHTTPClient(client *http.Client) *CreateanewSoundSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew sound set params
func (o *CreateanewSoundSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew sound set params
func (o *CreateanewSoundSetParams) WithBody(body *models.CreateanewSoundSetRequest) *CreateanewSoundSetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew sound set params
func (o *CreateanewSoundSetParams) SetBody(body *models.CreateanewSoundSetRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew sound set params
func (o *CreateanewSoundSetParams) WithContentType(contentType string) *CreateanewSoundSetParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew sound set params
func (o *CreateanewSoundSetParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewSoundSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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