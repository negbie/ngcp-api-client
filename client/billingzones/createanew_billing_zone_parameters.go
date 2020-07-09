// Code generated by go-swagger; DO NOT EDIT.

package billingzones

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

// NewCreateanewBillingZoneParams creates a new CreateanewBillingZoneParams object
// with the default values initialized.
func NewCreateanewBillingZoneParams() *CreateanewBillingZoneParams {
	var ()
	return &CreateanewBillingZoneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewBillingZoneParamsWithTimeout creates a new CreateanewBillingZoneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewBillingZoneParamsWithTimeout(timeout time.Duration) *CreateanewBillingZoneParams {
	var ()
	return &CreateanewBillingZoneParams{

		timeout: timeout,
	}
}

// NewCreateanewBillingZoneParamsWithContext creates a new CreateanewBillingZoneParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewBillingZoneParamsWithContext(ctx context.Context) *CreateanewBillingZoneParams {
	var ()
	return &CreateanewBillingZoneParams{

		Context: ctx,
	}
}

// NewCreateanewBillingZoneParamsWithHTTPClient creates a new CreateanewBillingZoneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewBillingZoneParamsWithHTTPClient(client *http.Client) *CreateanewBillingZoneParams {
	var ()
	return &CreateanewBillingZoneParams{
		HTTPClient: client,
	}
}

/*CreateanewBillingZoneParams contains all the parameters to send to the API endpoint
for the createanew billing zone operation typically these are written to a http.Request
*/
type CreateanewBillingZoneParams struct {

	/*Body*/
	Body *models.CreateanewBillingZoneRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew billing zone params
func (o *CreateanewBillingZoneParams) WithTimeout(timeout time.Duration) *CreateanewBillingZoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew billing zone params
func (o *CreateanewBillingZoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew billing zone params
func (o *CreateanewBillingZoneParams) WithContext(ctx context.Context) *CreateanewBillingZoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew billing zone params
func (o *CreateanewBillingZoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew billing zone params
func (o *CreateanewBillingZoneParams) WithHTTPClient(client *http.Client) *CreateanewBillingZoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew billing zone params
func (o *CreateanewBillingZoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew billing zone params
func (o *CreateanewBillingZoneParams) WithBody(body *models.CreateanewBillingZoneRequest) *CreateanewBillingZoneParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew billing zone params
func (o *CreateanewBillingZoneParams) SetBody(body *models.CreateanewBillingZoneRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew billing zone params
func (o *CreateanewBillingZoneParams) WithContentType(contentType string) *CreateanewBillingZoneParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew billing zone params
func (o *CreateanewBillingZoneParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewBillingZoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
