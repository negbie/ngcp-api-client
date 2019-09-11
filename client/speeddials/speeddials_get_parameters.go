// Code generated by go-swagger; DO NOT EDIT.

package speeddials

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

// NewSpeeddialsGetParams creates a new SpeeddialsGetParams object
// with the default values initialized.
func NewSpeeddialsGetParams() *SpeeddialsGetParams {
	var ()
	return &SpeeddialsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSpeeddialsGetParamsWithTimeout creates a new SpeeddialsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSpeeddialsGetParamsWithTimeout(timeout time.Duration) *SpeeddialsGetParams {
	var ()
	return &SpeeddialsGetParams{

		timeout: timeout,
	}
}

// NewSpeeddialsGetParamsWithContext creates a new SpeeddialsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSpeeddialsGetParamsWithContext(ctx context.Context) *SpeeddialsGetParams {
	var ()
	return &SpeeddialsGetParams{

		Context: ctx,
	}
}

// NewSpeeddialsGetParamsWithHTTPClient creates a new SpeeddialsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSpeeddialsGetParamsWithHTTPClient(client *http.Client) *SpeeddialsGetParams {
	var ()
	return &SpeeddialsGetParams{
		HTTPClient: client,
	}
}

/*SpeeddialsGetParams contains all the parameters to send to the API endpoint
for the speeddials get operation typically these are written to a http.Request
*/
type SpeeddialsGetParams struct {

	/*Nonempty
	  Filter for subscribers with nonempty speeddials

	*/
	Nonempty string
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

// WithTimeout adds the timeout to the speeddials get params
func (o *SpeeddialsGetParams) WithTimeout(timeout time.Duration) *SpeeddialsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the speeddials get params
func (o *SpeeddialsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the speeddials get params
func (o *SpeeddialsGetParams) WithContext(ctx context.Context) *SpeeddialsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the speeddials get params
func (o *SpeeddialsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the speeddials get params
func (o *SpeeddialsGetParams) WithHTTPClient(client *http.Client) *SpeeddialsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the speeddials get params
func (o *SpeeddialsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNonempty adds the nonempty to the speeddials get params
func (o *SpeeddialsGetParams) WithNonempty(nonempty string) *SpeeddialsGetParams {
	o.SetNonempty(nonempty)
	return o
}

// SetNonempty adds the nonempty to the speeddials get params
func (o *SpeeddialsGetParams) SetNonempty(nonempty string) {
	o.Nonempty = nonempty
}

// WithOrderBy adds the orderBy to the speeddials get params
func (o *SpeeddialsGetParams) WithOrderBy(orderBy string) *SpeeddialsGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the speeddials get params
func (o *SpeeddialsGetParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the speeddials get params
func (o *SpeeddialsGetParams) WithOrderByDirection(orderByDirection string) *SpeeddialsGetParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the speeddials get params
func (o *SpeeddialsGetParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the speeddials get params
func (o *SpeeddialsGetParams) WithPage(page string) *SpeeddialsGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the speeddials get params
func (o *SpeeddialsGetParams) SetPage(page string) {
	o.Page = page
}

// WithRows adds the rows to the speeddials get params
func (o *SpeeddialsGetParams) WithRows(rows string) *SpeeddialsGetParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the speeddials get params
func (o *SpeeddialsGetParams) SetRows(rows string) {
	o.Rows = rows
}

// WriteToRequest writes these params to a swagger request
func (o *SpeeddialsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param nonempty
	qrNonempty := o.Nonempty
	qNonempty := qrNonempty
	if qNonempty != "" {
		if err := r.SetQueryParam("nonempty", qNonempty); err != nil {
			return err
		}
	}

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