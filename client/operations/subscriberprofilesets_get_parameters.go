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

// NewSubscriberprofilesetsGetParams creates a new SubscriberprofilesetsGetParams object
// with the default values initialized.
func NewSubscriberprofilesetsGetParams() *SubscriberprofilesetsGetParams {
	var ()
	return &SubscriberprofilesetsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubscriberprofilesetsGetParamsWithTimeout creates a new SubscriberprofilesetsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubscriberprofilesetsGetParamsWithTimeout(timeout time.Duration) *SubscriberprofilesetsGetParams {
	var ()
	return &SubscriberprofilesetsGetParams{

		timeout: timeout,
	}
}

// NewSubscriberprofilesetsGetParamsWithContext creates a new SubscriberprofilesetsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubscriberprofilesetsGetParamsWithContext(ctx context.Context) *SubscriberprofilesetsGetParams {
	var ()
	return &SubscriberprofilesetsGetParams{

		Context: ctx,
	}
}

// NewSubscriberprofilesetsGetParamsWithHTTPClient creates a new SubscriberprofilesetsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSubscriberprofilesetsGetParamsWithHTTPClient(client *http.Client) *SubscriberprofilesetsGetParams {
	var ()
	return &SubscriberprofilesetsGetParams{
		HTTPClient: client,
	}
}

/*SubscriberprofilesetsGetParams contains all the parameters to send to the API endpoint
for the subscriberprofilesets get operation typically these are written to a http.Request
*/
type SubscriberprofilesetsGetParams struct {

	/*Name
	  Filter for profile sets with a specific name

	*/
	Name string
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
	  Filter for profile sets belonging to a specific reseller

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

// WithTimeout adds the timeout to the subscriberprofilesets get params
func (o *SubscriberprofilesetsGetParams) WithTimeout(timeout time.Duration) *SubscriberprofilesetsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscriberprofilesets get params
func (o *SubscriberprofilesetsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscriberprofilesets get params
func (o *SubscriberprofilesetsGetParams) WithContext(ctx context.Context) *SubscriberprofilesetsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscriberprofilesets get params
func (o *SubscriberprofilesetsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscriberprofilesets get params
func (o *SubscriberprofilesetsGetParams) WithHTTPClient(client *http.Client) *SubscriberprofilesetsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscriberprofilesets get params
func (o *SubscriberprofilesetsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the subscriberprofilesets get params
func (o *SubscriberprofilesetsGetParams) WithName(name string) *SubscriberprofilesetsGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the subscriberprofilesets get params
func (o *SubscriberprofilesetsGetParams) SetName(name string) {
	o.Name = name
}

// WithOrderBy adds the orderBy to the subscriberprofilesets get params
func (o *SubscriberprofilesetsGetParams) WithOrderBy(orderBy string) *SubscriberprofilesetsGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the subscriberprofilesets get params
func (o *SubscriberprofilesetsGetParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the subscriberprofilesets get params
func (o *SubscriberprofilesetsGetParams) WithOrderByDirection(orderByDirection string) *SubscriberprofilesetsGetParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the subscriberprofilesets get params
func (o *SubscriberprofilesetsGetParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the subscriberprofilesets get params
func (o *SubscriberprofilesetsGetParams) WithPage(page string) *SubscriberprofilesetsGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the subscriberprofilesets get params
func (o *SubscriberprofilesetsGetParams) SetPage(page string) {
	o.Page = page
}

// WithResellerID adds the resellerID to the subscriberprofilesets get params
func (o *SubscriberprofilesetsGetParams) WithResellerID(resellerID string) *SubscriberprofilesetsGetParams {
	o.SetResellerID(resellerID)
	return o
}

// SetResellerID adds the resellerId to the subscriberprofilesets get params
func (o *SubscriberprofilesetsGetParams) SetResellerID(resellerID string) {
	o.ResellerID = resellerID
}

// WithRows adds the rows to the subscriberprofilesets get params
func (o *SubscriberprofilesetsGetParams) WithRows(rows string) *SubscriberprofilesetsGetParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the subscriberprofilesets get params
func (o *SubscriberprofilesetsGetParams) SetRows(rows string) {
	o.Rows = rows
}

// WriteToRequest writes these params to a swagger request
func (o *SubscriberprofilesetsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param name
	qrName := o.Name
	qName := qrName
	if qName != "" {
		if err := r.SetQueryParam("name", qName); err != nil {
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