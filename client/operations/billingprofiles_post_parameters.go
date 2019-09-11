// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	models "github.com/negbie/ngcp-api/models"
)

// NewBillingprofilesPostParams creates a new BillingprofilesPostParams object
// with the default values initialized.
func NewBillingprofilesPostParams() *BillingprofilesPostParams {
	var ()
	return &BillingprofilesPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBillingprofilesPostParamsWithTimeout creates a new BillingprofilesPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBillingprofilesPostParamsWithTimeout(timeout time.Duration) *BillingprofilesPostParams {
	var ()
	return &BillingprofilesPostParams{

		timeout: timeout,
	}
}

// NewBillingprofilesPostParamsWithContext creates a new BillingprofilesPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewBillingprofilesPostParamsWithContext(ctx context.Context) *BillingprofilesPostParams {
	var ()
	return &BillingprofilesPostParams{

		Context: ctx,
	}
}

// NewBillingprofilesPostParamsWithHTTPClient creates a new BillingprofilesPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBillingprofilesPostParamsWithHTTPClient(client *http.Client) *BillingprofilesPostParams {
	var ()
	return &BillingprofilesPostParams{
		HTTPClient: client,
	}
}

/*BillingprofilesPostParams contains all the parameters to send to the API endpoint
for the billingprofiles post operation typically these are written to a http.Request
*/
type BillingprofilesPostParams struct {

	/*Body*/
	Body *models.CreateanewBillingProfileRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the billingprofiles post params
func (o *BillingprofilesPostParams) WithTimeout(timeout time.Duration) *BillingprofilesPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billingprofiles post params
func (o *BillingprofilesPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billingprofiles post params
func (o *BillingprofilesPostParams) WithContext(ctx context.Context) *BillingprofilesPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billingprofiles post params
func (o *BillingprofilesPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billingprofiles post params
func (o *BillingprofilesPostParams) WithHTTPClient(client *http.Client) *BillingprofilesPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billingprofiles post params
func (o *BillingprofilesPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the billingprofiles post params
func (o *BillingprofilesPostParams) WithBody(body *models.CreateanewBillingProfileRequest) *BillingprofilesPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the billingprofiles post params
func (o *BillingprofilesPostParams) SetBody(body *models.CreateanewBillingProfileRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the billingprofiles post params
func (o *BillingprofilesPostParams) WithContentType(contentType string) *BillingprofilesPostParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the billingprofiles post params
func (o *BillingprofilesPostParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *BillingprofilesPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
