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

// NewLnpcarriersPostParams creates a new LnpcarriersPostParams object
// with the default values initialized.
func NewLnpcarriersPostParams() *LnpcarriersPostParams {
	var ()
	return &LnpcarriersPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLnpcarriersPostParamsWithTimeout creates a new LnpcarriersPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLnpcarriersPostParamsWithTimeout(timeout time.Duration) *LnpcarriersPostParams {
	var ()
	return &LnpcarriersPostParams{

		timeout: timeout,
	}
}

// NewLnpcarriersPostParamsWithContext creates a new LnpcarriersPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewLnpcarriersPostParamsWithContext(ctx context.Context) *LnpcarriersPostParams {
	var ()
	return &LnpcarriersPostParams{

		Context: ctx,
	}
}

// NewLnpcarriersPostParamsWithHTTPClient creates a new LnpcarriersPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLnpcarriersPostParamsWithHTTPClient(client *http.Client) *LnpcarriersPostParams {
	var ()
	return &LnpcarriersPostParams{
		HTTPClient: client,
	}
}

/*LnpcarriersPostParams contains all the parameters to send to the API endpoint
for the lnpcarriers post operation typically these are written to a http.Request
*/
type LnpcarriersPostParams struct {

	/*Body*/
	Body *models.CreateanewLnpCarrierRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the lnpcarriers post params
func (o *LnpcarriersPostParams) WithTimeout(timeout time.Duration) *LnpcarriersPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lnpcarriers post params
func (o *LnpcarriersPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lnpcarriers post params
func (o *LnpcarriersPostParams) WithContext(ctx context.Context) *LnpcarriersPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lnpcarriers post params
func (o *LnpcarriersPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lnpcarriers post params
func (o *LnpcarriersPostParams) WithHTTPClient(client *http.Client) *LnpcarriersPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lnpcarriers post params
func (o *LnpcarriersPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the lnpcarriers post params
func (o *LnpcarriersPostParams) WithBody(body *models.CreateanewLnpCarrierRequest) *LnpcarriersPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the lnpcarriers post params
func (o *LnpcarriersPostParams) SetBody(body *models.CreateanewLnpCarrierRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the lnpcarriers post params
func (o *LnpcarriersPostParams) WithContentType(contentType string) *LnpcarriersPostParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the lnpcarriers post params
func (o *LnpcarriersPostParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *LnpcarriersPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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