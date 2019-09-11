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

// NewBillingprofilesGetParams creates a new BillingprofilesGetParams object
// with the default values initialized.
func NewBillingprofilesGetParams() *BillingprofilesGetParams {
	var ()
	return &BillingprofilesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBillingprofilesGetParamsWithTimeout creates a new BillingprofilesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBillingprofilesGetParamsWithTimeout(timeout time.Duration) *BillingprofilesGetParams {
	var ()
	return &BillingprofilesGetParams{

		timeout: timeout,
	}
}

// NewBillingprofilesGetParamsWithContext creates a new BillingprofilesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewBillingprofilesGetParamsWithContext(ctx context.Context) *BillingprofilesGetParams {
	var ()
	return &BillingprofilesGetParams{

		Context: ctx,
	}
}

// NewBillingprofilesGetParamsWithHTTPClient creates a new BillingprofilesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBillingprofilesGetParamsWithHTTPClient(client *http.Client) *BillingprofilesGetParams {
	var ()
	return &BillingprofilesGetParams{
		HTTPClient: client,
	}
}

/*BillingprofilesGetParams contains all the parameters to send to the API endpoint
for the billingprofiles get operation typically these are written to a http.Request
*/
type BillingprofilesGetParams struct {

	/*Handle
	  Filter for billing profiles with a specific handle

	*/
	Handle string
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
	/*ResellerID
	  Filter for billing profiles belonging to a specific reseller

	*/
	ResellerID string
	/*Rows
	  Number of rows in one pagination page (default: 10)

	*/
	Rows string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the billingprofiles get params
func (o *BillingprofilesGetParams) WithTimeout(timeout time.Duration) *BillingprofilesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billingprofiles get params
func (o *BillingprofilesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billingprofiles get params
func (o *BillingprofilesGetParams) WithContext(ctx context.Context) *BillingprofilesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billingprofiles get params
func (o *BillingprofilesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billingprofiles get params
func (o *BillingprofilesGetParams) WithHTTPClient(client *http.Client) *BillingprofilesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billingprofiles get params
func (o *BillingprofilesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHandle adds the handle to the billingprofiles get params
func (o *BillingprofilesGetParams) WithHandle(handle string) *BillingprofilesGetParams {
	o.SetHandle(handle)
	return o
}

// SetHandle adds the handle to the billingprofiles get params
func (o *BillingprofilesGetParams) SetHandle(handle string) {
	o.Handle = handle
}

// WithOrderBy adds the orderBy to the billingprofiles get params
func (o *BillingprofilesGetParams) WithOrderBy(orderBy string) *BillingprofilesGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the billingprofiles get params
func (o *BillingprofilesGetParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the billingprofiles get params
func (o *BillingprofilesGetParams) WithOrderByDirection(orderByDirection string) *BillingprofilesGetParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the billingprofiles get params
func (o *BillingprofilesGetParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the billingprofiles get params
func (o *BillingprofilesGetParams) WithPage(page string) *BillingprofilesGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the billingprofiles get params
func (o *BillingprofilesGetParams) SetPage(page string) {
	o.Page = page
}

// WithResellerID adds the resellerID to the billingprofiles get params
func (o *BillingprofilesGetParams) WithResellerID(resellerID string) *BillingprofilesGetParams {
	o.SetResellerID(resellerID)
	return o
}

// SetResellerID adds the resellerId to the billingprofiles get params
func (o *BillingprofilesGetParams) SetResellerID(resellerID string) {
	o.ResellerID = resellerID
}

// WithRows adds the rows to the billingprofiles get params
func (o *BillingprofilesGetParams) WithRows(rows string) *BillingprofilesGetParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the billingprofiles get params
func (o *BillingprofilesGetParams) SetRows(rows string) {
	o.Rows = rows
}

// WriteToRequest writes these params to a swagger request
func (o *BillingprofilesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param handle
	qrHandle := o.Handle
	qHandle := qrHandle
	if qHandle != "" {
		if err := r.SetQueryParam("handle", qHandle); err != nil {
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

	// query param reseller_id
	qrResellerID := o.ResellerID
	qResellerID := qrResellerID
	if qResellerID != "" {
		if err := r.SetQueryParam("reseller_id", qResellerID); err != nil {
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