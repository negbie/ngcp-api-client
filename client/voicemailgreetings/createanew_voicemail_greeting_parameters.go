// Code generated by go-swagger; DO NOT EDIT.

package voicemailgreetings

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

// NewCreateanewVoicemailGreetingParams creates a new CreateanewVoicemailGreetingParams object
// with the default values initialized.
func NewCreateanewVoicemailGreetingParams() *CreateanewVoicemailGreetingParams {
	var ()
	return &CreateanewVoicemailGreetingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewVoicemailGreetingParamsWithTimeout creates a new CreateanewVoicemailGreetingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewVoicemailGreetingParamsWithTimeout(timeout time.Duration) *CreateanewVoicemailGreetingParams {
	var ()
	return &CreateanewVoicemailGreetingParams{

		timeout: timeout,
	}
}

// NewCreateanewVoicemailGreetingParamsWithContext creates a new CreateanewVoicemailGreetingParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewVoicemailGreetingParamsWithContext(ctx context.Context) *CreateanewVoicemailGreetingParams {
	var ()
	return &CreateanewVoicemailGreetingParams{

		Context: ctx,
	}
}

// NewCreateanewVoicemailGreetingParamsWithHTTPClient creates a new CreateanewVoicemailGreetingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewVoicemailGreetingParamsWithHTTPClient(client *http.Client) *CreateanewVoicemailGreetingParams {
	var ()
	return &CreateanewVoicemailGreetingParams{
		HTTPClient: client,
	}
}

/*CreateanewVoicemailGreetingParams contains all the parameters to send to the API endpoint
for the createanew voicemail greeting operation typically these are written to a http.Request
*/
type CreateanewVoicemailGreetingParams struct {

	/*Body*/
	Body *models.CreateanewVoicemailGreetingRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew voicemail greeting params
func (o *CreateanewVoicemailGreetingParams) WithTimeout(timeout time.Duration) *CreateanewVoicemailGreetingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew voicemail greeting params
func (o *CreateanewVoicemailGreetingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew voicemail greeting params
func (o *CreateanewVoicemailGreetingParams) WithContext(ctx context.Context) *CreateanewVoicemailGreetingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew voicemail greeting params
func (o *CreateanewVoicemailGreetingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew voicemail greeting params
func (o *CreateanewVoicemailGreetingParams) WithHTTPClient(client *http.Client) *CreateanewVoicemailGreetingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew voicemail greeting params
func (o *CreateanewVoicemailGreetingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew voicemail greeting params
func (o *CreateanewVoicemailGreetingParams) WithBody(body *models.CreateanewVoicemailGreetingRequest) *CreateanewVoicemailGreetingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew voicemail greeting params
func (o *CreateanewVoicemailGreetingParams) SetBody(body *models.CreateanewVoicemailGreetingRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew voicemail greeting params
func (o *CreateanewVoicemailGreetingParams) WithContentType(contentType string) *CreateanewVoicemailGreetingParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew voicemail greeting params
func (o *CreateanewVoicemailGreetingParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewVoicemailGreetingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
