// Code generated by go-swagger; DO NOT EDIT.

package peeringserverpreferences

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

// NewPeeringserverpreferencesGetParams creates a new PeeringserverpreferencesGetParams object
// with the default values initialized.
func NewPeeringserverpreferencesGetParams() *PeeringserverpreferencesGetParams {
	var ()
	return &PeeringserverpreferencesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPeeringserverpreferencesGetParamsWithTimeout creates a new PeeringserverpreferencesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPeeringserverpreferencesGetParamsWithTimeout(timeout time.Duration) *PeeringserverpreferencesGetParams {
	var ()
	return &PeeringserverpreferencesGetParams{

		timeout: timeout,
	}
}

// NewPeeringserverpreferencesGetParamsWithContext creates a new PeeringserverpreferencesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPeeringserverpreferencesGetParamsWithContext(ctx context.Context) *PeeringserverpreferencesGetParams {
	var ()
	return &PeeringserverpreferencesGetParams{

		Context: ctx,
	}
}

// NewPeeringserverpreferencesGetParamsWithHTTPClient creates a new PeeringserverpreferencesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPeeringserverpreferencesGetParamsWithHTTPClient(client *http.Client) *PeeringserverpreferencesGetParams {
	var ()
	return &PeeringserverpreferencesGetParams{
		HTTPClient: client,
	}
}

/*PeeringserverpreferencesGetParams contains all the parameters to send to the API endpoint
for the peeringserverpreferences get operation typically these are written to a http.Request
*/
type PeeringserverpreferencesGetParams struct {

	/*OrderBy
	  Order collection by a specific attribute.

	*/
	OrderBy string
	/*OrderByDirection
	  Direction which the collection should be ordered by. Possible values are: asc (default), desc.

	*/
	OrderByDirection string
	/*Page
	  Pagination page which should be displayed (default: 1)

	*/
	Page string
	/*Rows
	  Number of rows in one pagination page (default: 10)

	*/
	Rows string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the peeringserverpreferences get params
func (o *PeeringserverpreferencesGetParams) WithTimeout(timeout time.Duration) *PeeringserverpreferencesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the peeringserverpreferences get params
func (o *PeeringserverpreferencesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the peeringserverpreferences get params
func (o *PeeringserverpreferencesGetParams) WithContext(ctx context.Context) *PeeringserverpreferencesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the peeringserverpreferences get params
func (o *PeeringserverpreferencesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the peeringserverpreferences get params
func (o *PeeringserverpreferencesGetParams) WithHTTPClient(client *http.Client) *PeeringserverpreferencesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the peeringserverpreferences get params
func (o *PeeringserverpreferencesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderBy adds the orderBy to the peeringserverpreferences get params
func (o *PeeringserverpreferencesGetParams) WithOrderBy(orderBy string) *PeeringserverpreferencesGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the peeringserverpreferences get params
func (o *PeeringserverpreferencesGetParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the peeringserverpreferences get params
func (o *PeeringserverpreferencesGetParams) WithOrderByDirection(orderByDirection string) *PeeringserverpreferencesGetParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the peeringserverpreferences get params
func (o *PeeringserverpreferencesGetParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the peeringserverpreferences get params
func (o *PeeringserverpreferencesGetParams) WithPage(page string) *PeeringserverpreferencesGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the peeringserverpreferences get params
func (o *PeeringserverpreferencesGetParams) SetPage(page string) {
	o.Page = page
}

// WithRows adds the rows to the peeringserverpreferences get params
func (o *PeeringserverpreferencesGetParams) WithRows(rows string) *PeeringserverpreferencesGetParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the peeringserverpreferences get params
func (o *PeeringserverpreferencesGetParams) SetRows(rows string) {
	o.Rows = rows
}

// WriteToRequest writes these params to a swagger request
func (o *PeeringserverpreferencesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param order_by
	qrOrderBy := o.OrderBy
	qOrderBy := qrOrderBy
	if qOrderBy != "" {
		if err := r.SetQueryParam("order_by", qOrderBy); err != nil {
			return err
		}
	}

	// query param order_by_direction
	qrOrderByDirection := o.OrderByDirection
	qOrderByDirection := qrOrderByDirection
	if qOrderByDirection != "" {
		if err := r.SetQueryParam("order_by_direction", qOrderByDirection); err != nil {
			return err
		}
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
