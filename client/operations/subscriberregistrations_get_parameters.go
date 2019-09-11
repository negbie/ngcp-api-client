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
)

// NewSubscriberregistrationsGetParams creates a new SubscriberregistrationsGetParams object
// with the default values initialized.
func NewSubscriberregistrationsGetParams() *SubscriberregistrationsGetParams {
	var ()
	return &SubscriberregistrationsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubscriberregistrationsGetParamsWithTimeout creates a new SubscriberregistrationsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubscriberregistrationsGetParamsWithTimeout(timeout time.Duration) *SubscriberregistrationsGetParams {
	var ()
	return &SubscriberregistrationsGetParams{

		timeout: timeout,
	}
}

// NewSubscriberregistrationsGetParamsWithContext creates a new SubscriberregistrationsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubscriberregistrationsGetParamsWithContext(ctx context.Context) *SubscriberregistrationsGetParams {
	var ()
	return &SubscriberregistrationsGetParams{

		Context: ctx,
	}
}

// NewSubscriberregistrationsGetParamsWithHTTPClient creates a new SubscriberregistrationsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSubscriberregistrationsGetParamsWithHTTPClient(client *http.Client) *SubscriberregistrationsGetParams {
	var ()
	return &SubscriberregistrationsGetParams{
		HTTPClient: client,
	}
}

/*SubscriberregistrationsGetParams contains all the parameters to send to the API endpoint
for the subscriberregistrations get operation typically these are written to a http.Request
*/
type SubscriberregistrationsGetParams struct {

	/*Page
	  Pagination page which should be displayed (default: 1)

	*/
	Page string
	/*Rows
	  Number of rows in one pagination page (default: 10)

	*/
	Rows string
	/*SubscriberID
	  Filter for registrations of a specific subscriber

	*/
	SubscriberID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the subscriberregistrations get params
func (o *SubscriberregistrationsGetParams) WithTimeout(timeout time.Duration) *SubscriberregistrationsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscriberregistrations get params
func (o *SubscriberregistrationsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscriberregistrations get params
func (o *SubscriberregistrationsGetParams) WithContext(ctx context.Context) *SubscriberregistrationsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscriberregistrations get params
func (o *SubscriberregistrationsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscriberregistrations get params
func (o *SubscriberregistrationsGetParams) WithHTTPClient(client *http.Client) *SubscriberregistrationsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscriberregistrations get params
func (o *SubscriberregistrationsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the subscriberregistrations get params
func (o *SubscriberregistrationsGetParams) WithPage(page string) *SubscriberregistrationsGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the subscriberregistrations get params
func (o *SubscriberregistrationsGetParams) SetPage(page string) {
	o.Page = page
}

// WithRows adds the rows to the subscriberregistrations get params
func (o *SubscriberregistrationsGetParams) WithRows(rows string) *SubscriberregistrationsGetParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the subscriberregistrations get params
func (o *SubscriberregistrationsGetParams) SetRows(rows string) {
	o.Rows = rows
}

// WithSubscriberID adds the subscriberID to the subscriberregistrations get params
func (o *SubscriberregistrationsGetParams) WithSubscriberID(subscriberID string) *SubscriberregistrationsGetParams {
	o.SetSubscriberID(subscriberID)
	return o
}

// SetSubscriberID adds the subscriberId to the subscriberregistrations get params
func (o *SubscriberregistrationsGetParams) SetSubscriberID(subscriberID string) {
	o.SubscriberID = subscriberID
}

// WriteToRequest writes these params to a swagger request
func (o *SubscriberregistrationsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param page
	qrPage := o.Page
	qPage := qrPage
	if qPage != "" {
		if err := r.SetQueryParam("page", qPage); err != nil {
			return err
		}
	}

	// query param rows
	qrRows := o.Rows
	qRows := qrRows
	if qRows != "" {
		if err := r.SetQueryParam("rows", qRows); err != nil {
			return err
		}
	}

	// query param subscriber_id
	qrSubscriberID := o.SubscriberID
	qSubscriberID := qrSubscriberID
	if qSubscriberID != "" {
		if err := r.SetQueryParam("subscriber_id", qSubscriberID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}