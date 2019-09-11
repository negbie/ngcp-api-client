// Code generated by go-swagger; DO NOT EDIT.

package id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewEmailtemplatesByIDPutParams creates a new EmailtemplatesByIDPutParams object
// with the default values initialized.
func NewEmailtemplatesByIDPutParams() *EmailtemplatesByIDPutParams {
	var ()
	return &EmailtemplatesByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmailtemplatesByIDPutParamsWithTimeout creates a new EmailtemplatesByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmailtemplatesByIDPutParamsWithTimeout(timeout time.Duration) *EmailtemplatesByIDPutParams {
	var ()
	return &EmailtemplatesByIDPutParams{

		timeout: timeout,
	}
}

// NewEmailtemplatesByIDPutParamsWithContext creates a new EmailtemplatesByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmailtemplatesByIDPutParamsWithContext(ctx context.Context) *EmailtemplatesByIDPutParams {
	var ()
	return &EmailtemplatesByIDPutParams{

		Context: ctx,
	}
}

// NewEmailtemplatesByIDPutParamsWithHTTPClient creates a new EmailtemplatesByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmailtemplatesByIDPutParamsWithHTTPClient(client *http.Client) *EmailtemplatesByIDPutParams {
	var ()
	return &EmailtemplatesByIDPutParams{
		HTTPClient: client,
	}
}

/*EmailtemplatesByIDPutParams contains all the parameters to send to the API endpoint
for the emailtemplates by Id put operation typically these are written to a http.Request
*/
type EmailtemplatesByIDPutParams struct {

	/*Body*/
	Body *Replace1changeaspecificEmailTemplateRequest
	/*ContentType*/
	ContentType string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the emailtemplates by Id put params
func (o *EmailtemplatesByIDPutParams) WithTimeout(timeout time.Duration) *EmailtemplatesByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the emailtemplates by Id put params
func (o *EmailtemplatesByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the emailtemplates by Id put params
func (o *EmailtemplatesByIDPutParams) WithContext(ctx context.Context) *EmailtemplatesByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the emailtemplates by Id put params
func (o *EmailtemplatesByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the emailtemplates by Id put params
func (o *EmailtemplatesByIDPutParams) WithHTTPClient(client *http.Client) *EmailtemplatesByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the emailtemplates by Id put params
func (o *EmailtemplatesByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the emailtemplates by Id put params
func (o *EmailtemplatesByIDPutParams) WithBody(body *Replace1changeaspecificEmailTemplateRequest) *EmailtemplatesByIDPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the emailtemplates by Id put params
func (o *EmailtemplatesByIDPutParams) SetBody(body *Replace1changeaspecificEmailTemplateRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the emailtemplates by Id put params
func (o *EmailtemplatesByIDPutParams) WithContentType(contentType string) *EmailtemplatesByIDPutParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the emailtemplates by Id put params
func (o *EmailtemplatesByIDPutParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the emailtemplates by Id put params
func (o *EmailtemplatesByIDPutParams) WithID(id string) *EmailtemplatesByIDPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the emailtemplates by Id put params
func (o *EmailtemplatesByIDPutParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EmailtemplatesByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}