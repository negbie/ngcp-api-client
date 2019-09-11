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

// NewCfsourcesetsPostParams creates a new CfsourcesetsPostParams object
// with the default values initialized.
func NewCfsourcesetsPostParams() *CfsourcesetsPostParams {
	var ()
	return &CfsourcesetsPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCfsourcesetsPostParamsWithTimeout creates a new CfsourcesetsPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCfsourcesetsPostParamsWithTimeout(timeout time.Duration) *CfsourcesetsPostParams {
	var ()
	return &CfsourcesetsPostParams{

		timeout: timeout,
	}
}

// NewCfsourcesetsPostParamsWithContext creates a new CfsourcesetsPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCfsourcesetsPostParamsWithContext(ctx context.Context) *CfsourcesetsPostParams {
	var ()
	return &CfsourcesetsPostParams{

		Context: ctx,
	}
}

// NewCfsourcesetsPostParamsWithHTTPClient creates a new CfsourcesetsPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCfsourcesetsPostParamsWithHTTPClient(client *http.Client) *CfsourcesetsPostParams {
	var ()
	return &CfsourcesetsPostParams{
		HTTPClient: client,
	}
}

/*CfsourcesetsPostParams contains all the parameters to send to the API endpoint
for the cfsourcesets post operation typically these are written to a http.Request
*/
type CfsourcesetsPostParams struct {

	/*Body*/
	Body *models.CreateanewCFSourceSetRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cfsourcesets post params
func (o *CfsourcesetsPostParams) WithTimeout(timeout time.Duration) *CfsourcesetsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cfsourcesets post params
func (o *CfsourcesetsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cfsourcesets post params
func (o *CfsourcesetsPostParams) WithContext(ctx context.Context) *CfsourcesetsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cfsourcesets post params
func (o *CfsourcesetsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cfsourcesets post params
func (o *CfsourcesetsPostParams) WithHTTPClient(client *http.Client) *CfsourcesetsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cfsourcesets post params
func (o *CfsourcesetsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cfsourcesets post params
func (o *CfsourcesetsPostParams) WithBody(body *models.CreateanewCFSourceSetRequest) *CfsourcesetsPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cfsourcesets post params
func (o *CfsourcesetsPostParams) SetBody(body *models.CreateanewCFSourceSetRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the cfsourcesets post params
func (o *CfsourcesetsPostParams) WithContentType(contentType string) *CfsourcesetsPostParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the cfsourcesets post params
func (o *CfsourcesetsPostParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CfsourcesetsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
