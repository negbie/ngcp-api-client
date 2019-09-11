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

// NewManagersecretaryByIDDeleteParams creates a new ManagersecretaryByIDDeleteParams object
// with the default values initialized.
func NewManagersecretaryByIDDeleteParams() *ManagersecretaryByIDDeleteParams {
	var ()
	return &ManagersecretaryByIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewManagersecretaryByIDDeleteParamsWithTimeout creates a new ManagersecretaryByIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewManagersecretaryByIDDeleteParamsWithTimeout(timeout time.Duration) *ManagersecretaryByIDDeleteParams {
	var ()
	return &ManagersecretaryByIDDeleteParams{

		timeout: timeout,
	}
}

// NewManagersecretaryByIDDeleteParamsWithContext creates a new ManagersecretaryByIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewManagersecretaryByIDDeleteParamsWithContext(ctx context.Context) *ManagersecretaryByIDDeleteParams {
	var ()
	return &ManagersecretaryByIDDeleteParams{

		Context: ctx,
	}
}

// NewManagersecretaryByIDDeleteParamsWithHTTPClient creates a new ManagersecretaryByIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewManagersecretaryByIDDeleteParamsWithHTTPClient(client *http.Client) *ManagersecretaryByIDDeleteParams {
	var ()
	return &ManagersecretaryByIDDeleteParams{
		HTTPClient: client,
	}
}

/*ManagersecretaryByIDDeleteParams contains all the parameters to send to the API endpoint
for the managersecretary by Id delete operation typically these are written to a http.Request
*/
type ManagersecretaryByIDDeleteParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the managersecretary by Id delete params
func (o *ManagersecretaryByIDDeleteParams) WithTimeout(timeout time.Duration) *ManagersecretaryByIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the managersecretary by Id delete params
func (o *ManagersecretaryByIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the managersecretary by Id delete params
func (o *ManagersecretaryByIDDeleteParams) WithContext(ctx context.Context) *ManagersecretaryByIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the managersecretary by Id delete params
func (o *ManagersecretaryByIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the managersecretary by Id delete params
func (o *ManagersecretaryByIDDeleteParams) WithHTTPClient(client *http.Client) *ManagersecretaryByIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the managersecretary by Id delete params
func (o *ManagersecretaryByIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the managersecretary by Id delete params
func (o *ManagersecretaryByIDDeleteParams) WithID(id string) *ManagersecretaryByIDDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the managersecretary by Id delete params
func (o *ManagersecretaryByIDDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ManagersecretaryByIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}