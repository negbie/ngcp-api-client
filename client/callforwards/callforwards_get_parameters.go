// Code generated by go-swagger; DO NOT EDIT.

package callforwards

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

// NewCallforwardsGetParams creates a new CallforwardsGetParams object
// with the default values initialized.
func NewCallforwardsGetParams() *CallforwardsGetParams {
	var ()
	return &CallforwardsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCallforwardsGetParamsWithTimeout creates a new CallforwardsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCallforwardsGetParamsWithTimeout(timeout time.Duration) *CallforwardsGetParams {
	var ()
	return &CallforwardsGetParams{

		timeout: timeout,
	}
}

// NewCallforwardsGetParamsWithContext creates a new CallforwardsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCallforwardsGetParamsWithContext(ctx context.Context) *CallforwardsGetParams {
	var ()
	return &CallforwardsGetParams{

		Context: ctx,
	}
}

// NewCallforwardsGetParamsWithHTTPClient creates a new CallforwardsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCallforwardsGetParamsWithHTTPClient(client *http.Client) *CallforwardsGetParams {
	var ()
	return &CallforwardsGetParams{
		HTTPClient: client,
	}
}

/*CallforwardsGetParams contains all the parameters to send to the API endpoint
for the callforwards get operation typically these are written to a http.Request
*/
type CallforwardsGetParams struct {

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

// WithTimeout adds the timeout to the callforwards get params
func (o *CallforwardsGetParams) WithTimeout(timeout time.Duration) *CallforwardsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the callforwards get params
func (o *CallforwardsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the callforwards get params
func (o *CallforwardsGetParams) WithContext(ctx context.Context) *CallforwardsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the callforwards get params
func (o *CallforwardsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the callforwards get params
func (o *CallforwardsGetParams) WithHTTPClient(client *http.Client) *CallforwardsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the callforwards get params
func (o *CallforwardsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderBy adds the orderBy to the callforwards get params
func (o *CallforwardsGetParams) WithOrderBy(orderBy string) *CallforwardsGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the callforwards get params
func (o *CallforwardsGetParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the callforwards get params
func (o *CallforwardsGetParams) WithOrderByDirection(orderByDirection string) *CallforwardsGetParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the callforwards get params
func (o *CallforwardsGetParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the callforwards get params
func (o *CallforwardsGetParams) WithPage(page string) *CallforwardsGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the callforwards get params
func (o *CallforwardsGetParams) SetPage(page string) {
	o.Page = page
}

// WithRows adds the rows to the callforwards get params
func (o *CallforwardsGetParams) WithRows(rows string) *CallforwardsGetParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the callforwards get params
func (o *CallforwardsGetParams) SetRows(rows string) {
	o.Rows = rows
}

// WriteToRequest writes these params to a swagger request
func (o *CallforwardsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
