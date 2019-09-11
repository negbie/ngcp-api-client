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

// NewPbxdeviceprofilesByIDPutParams creates a new PbxdeviceprofilesByIDPutParams object
// with the default values initialized.
func NewPbxdeviceprofilesByIDPutParams() *PbxdeviceprofilesByIDPutParams {
	var ()
	return &PbxdeviceprofilesByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPbxdeviceprofilesByIDPutParamsWithTimeout creates a new PbxdeviceprofilesByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPbxdeviceprofilesByIDPutParamsWithTimeout(timeout time.Duration) *PbxdeviceprofilesByIDPutParams {
	var ()
	return &PbxdeviceprofilesByIDPutParams{

		timeout: timeout,
	}
}

// NewPbxdeviceprofilesByIDPutParamsWithContext creates a new PbxdeviceprofilesByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewPbxdeviceprofilesByIDPutParamsWithContext(ctx context.Context) *PbxdeviceprofilesByIDPutParams {
	var ()
	return &PbxdeviceprofilesByIDPutParams{

		Context: ctx,
	}
}

// NewPbxdeviceprofilesByIDPutParamsWithHTTPClient creates a new PbxdeviceprofilesByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPbxdeviceprofilesByIDPutParamsWithHTTPClient(client *http.Client) *PbxdeviceprofilesByIDPutParams {
	var ()
	return &PbxdeviceprofilesByIDPutParams{
		HTTPClient: client,
	}
}

/*PbxdeviceprofilesByIDPutParams contains all the parameters to send to the API endpoint
for the pbxdeviceprofiles by Id put operation typically these are written to a http.Request
*/
type PbxdeviceprofilesByIDPutParams struct {

	/*Body*/
	Body *Replace1changeaspecificPbxDeviceProfileRequest
	/*ContentType*/
	ContentType string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pbxdeviceprofiles by Id put params
func (o *PbxdeviceprofilesByIDPutParams) WithTimeout(timeout time.Duration) *PbxdeviceprofilesByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pbxdeviceprofiles by Id put params
func (o *PbxdeviceprofilesByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pbxdeviceprofiles by Id put params
func (o *PbxdeviceprofilesByIDPutParams) WithContext(ctx context.Context) *PbxdeviceprofilesByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pbxdeviceprofiles by Id put params
func (o *PbxdeviceprofilesByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pbxdeviceprofiles by Id put params
func (o *PbxdeviceprofilesByIDPutParams) WithHTTPClient(client *http.Client) *PbxdeviceprofilesByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pbxdeviceprofiles by Id put params
func (o *PbxdeviceprofilesByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pbxdeviceprofiles by Id put params
func (o *PbxdeviceprofilesByIDPutParams) WithBody(body *Replace1changeaspecificPbxDeviceProfileRequest) *PbxdeviceprofilesByIDPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pbxdeviceprofiles by Id put params
func (o *PbxdeviceprofilesByIDPutParams) SetBody(body *Replace1changeaspecificPbxDeviceProfileRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the pbxdeviceprofiles by Id put params
func (o *PbxdeviceprofilesByIDPutParams) WithContentType(contentType string) *PbxdeviceprofilesByIDPutParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the pbxdeviceprofiles by Id put params
func (o *PbxdeviceprofilesByIDPutParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the pbxdeviceprofiles by Id put params
func (o *PbxdeviceprofilesByIDPutParams) WithID(id string) *PbxdeviceprofilesByIDPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the pbxdeviceprofiles by Id put params
func (o *PbxdeviceprofilesByIDPutParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PbxdeviceprofilesByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
