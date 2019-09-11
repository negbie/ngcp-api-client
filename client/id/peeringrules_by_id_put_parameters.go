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

// NewPeeringrulesByIDPutParams creates a new PeeringrulesByIDPutParams object
// with the default values initialized.
func NewPeeringrulesByIDPutParams() *PeeringrulesByIDPutParams {
	var ()
	return &PeeringrulesByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPeeringrulesByIDPutParamsWithTimeout creates a new PeeringrulesByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPeeringrulesByIDPutParamsWithTimeout(timeout time.Duration) *PeeringrulesByIDPutParams {
	var ()
	return &PeeringrulesByIDPutParams{

		timeout: timeout,
	}
}

// NewPeeringrulesByIDPutParamsWithContext creates a new PeeringrulesByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewPeeringrulesByIDPutParamsWithContext(ctx context.Context) *PeeringrulesByIDPutParams {
	var ()
	return &PeeringrulesByIDPutParams{

		Context: ctx,
	}
}

// NewPeeringrulesByIDPutParamsWithHTTPClient creates a new PeeringrulesByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPeeringrulesByIDPutParamsWithHTTPClient(client *http.Client) *PeeringrulesByIDPutParams {
	var ()
	return &PeeringrulesByIDPutParams{
		HTTPClient: client,
	}
}

/*PeeringrulesByIDPutParams contains all the parameters to send to the API endpoint
for the peeringrules by Id put operation typically these are written to a http.Request
*/
type PeeringrulesByIDPutParams struct {

	/*Body*/
	Body *Replace1changeaspecificPeeringRuleRequest
	/*ContentType*/
	ContentType string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the peeringrules by Id put params
func (o *PeeringrulesByIDPutParams) WithTimeout(timeout time.Duration) *PeeringrulesByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the peeringrules by Id put params
func (o *PeeringrulesByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the peeringrules by Id put params
func (o *PeeringrulesByIDPutParams) WithContext(ctx context.Context) *PeeringrulesByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the peeringrules by Id put params
func (o *PeeringrulesByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the peeringrules by Id put params
func (o *PeeringrulesByIDPutParams) WithHTTPClient(client *http.Client) *PeeringrulesByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the peeringrules by Id put params
func (o *PeeringrulesByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the peeringrules by Id put params
func (o *PeeringrulesByIDPutParams) WithBody(body *Replace1changeaspecificPeeringRuleRequest) *PeeringrulesByIDPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the peeringrules by Id put params
func (o *PeeringrulesByIDPutParams) SetBody(body *Replace1changeaspecificPeeringRuleRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the peeringrules by Id put params
func (o *PeeringrulesByIDPutParams) WithContentType(contentType string) *PeeringrulesByIDPutParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the peeringrules by Id put params
func (o *PeeringrulesByIDPutParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the peeringrules by Id put params
func (o *PeeringrulesByIDPutParams) WithID(id string) *PeeringrulesByIDPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the peeringrules by Id put params
func (o *PeeringrulesByIDPutParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PeeringrulesByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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