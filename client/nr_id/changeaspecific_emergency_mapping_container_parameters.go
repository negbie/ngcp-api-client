// Code generated by go-swagger; DO NOT EDIT.

package nr_id

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
)

// NewChangeaspecificEmergencyMappingContainerParams creates a new ChangeaspecificEmergencyMappingContainerParams object
// with the default values initialized.
func NewChangeaspecificEmergencyMappingContainerParams() *ChangeaspecificEmergencyMappingContainerParams {
	var ()
	return &ChangeaspecificEmergencyMappingContainerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeaspecificEmergencyMappingContainerParamsWithTimeout creates a new ChangeaspecificEmergencyMappingContainerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeaspecificEmergencyMappingContainerParamsWithTimeout(timeout time.Duration) *ChangeaspecificEmergencyMappingContainerParams {
	var ()
	return &ChangeaspecificEmergencyMappingContainerParams{

		timeout: timeout,
	}
}

// NewChangeaspecificEmergencyMappingContainerParamsWithContext creates a new ChangeaspecificEmergencyMappingContainerParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeaspecificEmergencyMappingContainerParamsWithContext(ctx context.Context) *ChangeaspecificEmergencyMappingContainerParams {
	var ()
	return &ChangeaspecificEmergencyMappingContainerParams{

		Context: ctx,
	}
}

// NewChangeaspecificEmergencyMappingContainerParamsWithHTTPClient creates a new ChangeaspecificEmergencyMappingContainerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeaspecificEmergencyMappingContainerParamsWithHTTPClient(client *http.Client) *ChangeaspecificEmergencyMappingContainerParams {
	var ()
	return &ChangeaspecificEmergencyMappingContainerParams{
		HTTPClient: client,
	}
}

/*ChangeaspecificEmergencyMappingContainerParams contains all the parameters to send to the API endpoint
for the changeaspecific emergency mapping container operation typically these are written to a http.Request
*/
type ChangeaspecificEmergencyMappingContainerParams struct {

	/*Body*/
	Body string
	/*ContentType*/
	ContentType string
	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the changeaspecific emergency mapping container params
func (o *ChangeaspecificEmergencyMappingContainerParams) WithTimeout(timeout time.Duration) *ChangeaspecificEmergencyMappingContainerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the changeaspecific emergency mapping container params
func (o *ChangeaspecificEmergencyMappingContainerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the changeaspecific emergency mapping container params
func (o *ChangeaspecificEmergencyMappingContainerParams) WithContext(ctx context.Context) *ChangeaspecificEmergencyMappingContainerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the changeaspecific emergency mapping container params
func (o *ChangeaspecificEmergencyMappingContainerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the changeaspecific emergency mapping container params
func (o *ChangeaspecificEmergencyMappingContainerParams) WithHTTPClient(client *http.Client) *ChangeaspecificEmergencyMappingContainerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the changeaspecific emergency mapping container params
func (o *ChangeaspecificEmergencyMappingContainerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the changeaspecific emergency mapping container params
func (o *ChangeaspecificEmergencyMappingContainerParams) WithBody(body string) *ChangeaspecificEmergencyMappingContainerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the changeaspecific emergency mapping container params
func (o *ChangeaspecificEmergencyMappingContainerParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the changeaspecific emergency mapping container params
func (o *ChangeaspecificEmergencyMappingContainerParams) WithContentType(contentType string) *ChangeaspecificEmergencyMappingContainerParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the changeaspecific emergency mapping container params
func (o *ChangeaspecificEmergencyMappingContainerParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the changeaspecific emergency mapping container params
func (o *ChangeaspecificEmergencyMappingContainerParams) WithID(id string) *ChangeaspecificEmergencyMappingContainerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the changeaspecific emergency mapping container params
func (o *ChangeaspecificEmergencyMappingContainerParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeaspecificEmergencyMappingContainerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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