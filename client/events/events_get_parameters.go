// Code generated by go-swagger; DO NOT EDIT.

package events

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

// NewEventsGetParams creates a new EventsGetParams object
// with the default values initialized.
func NewEventsGetParams() *EventsGetParams {
	var ()
	return &EventsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEventsGetParamsWithTimeout creates a new EventsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEventsGetParamsWithTimeout(timeout time.Duration) *EventsGetParams {
	var ()
	return &EventsGetParams{

		timeout: timeout,
	}
}

// NewEventsGetParamsWithContext creates a new EventsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewEventsGetParamsWithContext(ctx context.Context) *EventsGetParams {
	var ()
	return &EventsGetParams{

		Context: ctx,
	}
}

// NewEventsGetParamsWithHTTPClient creates a new EventsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEventsGetParamsWithHTTPClient(client *http.Client) *EventsGetParams {
	var ()
	return &EventsGetParams{
		HTTPClient: client,
	}
}

/*EventsGetParams contains all the parameters to send to the API endpoint
for the events get operation typically these are written to a http.Request
*/
type EventsGetParams struct {

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
	  Filter for events for customers/subscribers of a specific reseller.

	*/
	ResellerID string
	/*Rows
	  Number of rows in one pagination page (default: 10)

	*/
	Rows string
	/*SubscriberID
	  Filter for events of a specific subscriber.

	*/
	SubscriberID string
	/*TimestampFrom
	  Filter for events occurred after or at the given time stamp.

	*/
	TimestampFrom string
	/*TimestampTo
	  Filter for events occurred before or at the given time stamp.

	*/
	TimestampTo string
	/*Type
	  Filter for events of a specific type.

	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the events get params
func (o *EventsGetParams) WithTimeout(timeout time.Duration) *EventsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the events get params
func (o *EventsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the events get params
func (o *EventsGetParams) WithContext(ctx context.Context) *EventsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the events get params
func (o *EventsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the events get params
func (o *EventsGetParams) WithHTTPClient(client *http.Client) *EventsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the events get params
func (o *EventsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderBy adds the orderBy to the events get params
func (o *EventsGetParams) WithOrderBy(orderBy string) *EventsGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the events get params
func (o *EventsGetParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the events get params
func (o *EventsGetParams) WithOrderByDirection(orderByDirection string) *EventsGetParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the events get params
func (o *EventsGetParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the events get params
func (o *EventsGetParams) WithPage(page string) *EventsGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the events get params
func (o *EventsGetParams) SetPage(page string) {
	o.Page = page
}

// WithResellerID adds the resellerID to the events get params
func (o *EventsGetParams) WithResellerID(resellerID string) *EventsGetParams {
	o.SetResellerID(resellerID)
	return o
}

// SetResellerID adds the resellerId to the events get params
func (o *EventsGetParams) SetResellerID(resellerID string) {
	o.ResellerID = resellerID
}

// WithRows adds the rows to the events get params
func (o *EventsGetParams) WithRows(rows string) *EventsGetParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the events get params
func (o *EventsGetParams) SetRows(rows string) {
	o.Rows = rows
}

// WithSubscriberID adds the subscriberID to the events get params
func (o *EventsGetParams) WithSubscriberID(subscriberID string) *EventsGetParams {
	o.SetSubscriberID(subscriberID)
	return o
}

// SetSubscriberID adds the subscriberId to the events get params
func (o *EventsGetParams) SetSubscriberID(subscriberID string) {
	o.SubscriberID = subscriberID
}

// WithTimestampFrom adds the timestampFrom to the events get params
func (o *EventsGetParams) WithTimestampFrom(timestampFrom string) *EventsGetParams {
	o.SetTimestampFrom(timestampFrom)
	return o
}

// SetTimestampFrom adds the timestampFrom to the events get params
func (o *EventsGetParams) SetTimestampFrom(timestampFrom string) {
	o.TimestampFrom = timestampFrom
}

// WithTimestampTo adds the timestampTo to the events get params
func (o *EventsGetParams) WithTimestampTo(timestampTo string) *EventsGetParams {
	o.SetTimestampTo(timestampTo)
	return o
}

// SetTimestampTo adds the timestampTo to the events get params
func (o *EventsGetParams) SetTimestampTo(timestampTo string) {
	o.TimestampTo = timestampTo
}

// WithType adds the typeVar to the events get params
func (o *EventsGetParams) WithType(typeVar string) *EventsGetParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the events get params
func (o *EventsGetParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *EventsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param subscriber_id
	qrSubscriberID := o.SubscriberID
	qSubscriberID := qrSubscriberID
	if qSubscriberID != "" {
		if err := r.SetQueryParam("subscriber_id", qSubscriberID); err != nil {
			return err
		}
	}

	// query param timestamp_from
	qrTimestampFrom := o.TimestampFrom
	qTimestampFrom := qrTimestampFrom
	if qTimestampFrom != "" {
		if err := r.SetQueryParam("timestamp_from", qTimestampFrom); err != nil {
			return err
		}
	}

	// query param timestamp_to
	qrTimestampTo := o.TimestampTo
	qTimestampTo := qrTimestampTo
	if qTimestampTo != "" {
		if err := r.SetQueryParam("timestamp_to", qTimestampTo); err != nil {
			return err
		}
	}

	// query param type
	qrType := o.Type
	qType := qrType
	if qType != "" {
		if err := r.SetQueryParam("type", qType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
