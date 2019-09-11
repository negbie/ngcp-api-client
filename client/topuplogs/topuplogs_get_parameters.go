// Code generated by go-swagger; DO NOT EDIT.

package topuplogs

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

// NewTopuplogsGetParams creates a new TopuplogsGetParams object
// with the default values initialized.
func NewTopuplogsGetParams() *TopuplogsGetParams {
	var ()
	return &TopuplogsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTopuplogsGetParamsWithTimeout creates a new TopuplogsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTopuplogsGetParamsWithTimeout(timeout time.Duration) *TopuplogsGetParams {
	var ()
	return &TopuplogsGetParams{

		timeout: timeout,
	}
}

// NewTopuplogsGetParamsWithContext creates a new TopuplogsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewTopuplogsGetParamsWithContext(ctx context.Context) *TopuplogsGetParams {
	var ()
	return &TopuplogsGetParams{

		Context: ctx,
	}
}

// NewTopuplogsGetParamsWithHTTPClient creates a new TopuplogsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTopuplogsGetParamsWithHTTPClient(client *http.Client) *TopuplogsGetParams {
	var ()
	return &TopuplogsGetParams{
		HTTPClient: client,
	}
}

/*TopuplogsGetParams contains all the parameters to send to the API endpoint
for the topuplogs get operation typically these are written to a http.Request
*/
type TopuplogsGetParams struct {

	/*AmountAbove
	  Filter for top-up requests with an amount greater than or equal to the given value in USD/EUR/etc.

	*/
	AmountAbove string
	/*AmountBelow
	  Filter for top-up requests with an amount less than or equal to the given value in USD/EUR/etc.

	*/
	AmountBelow string
	/*ContractID
	  Filter for top-up requests of a specific customer contract.

	*/
	ContractID string
	/*OrderBy
	  Order collection by a specific attribute.

	*/
	OrderBy string
	/*OrderByDirection
	  Direction which the collection should be ordered by. Possible values are: asc (default), desc.

	*/
	OrderByDirection string
	/*Outcome
	  Filter for top-up requests by outcome.

	*/
	Outcome string
	/*Page
	  Pagination page which should be displayed (default: 1)

	*/
	Page string
	/*RequestToken
	  Filter for top-up requests with the given request_token.

	*/
	RequestToken string
	/*ResellerID
	  Filter for top-up requests for customers/subscribers of a specific reseller.

	*/
	ResellerID string
	/*Rows
	  Number of rows in one pagination page (default: 10)

	*/
	Rows string
	/*SubscriberID
	  Filter for top-up requests of a specific subscriber.

	*/
	SubscriberID string
	/*TimestampFrom
	  Filter for top-up requests performed after or at the given time stamp.

	*/
	TimestampFrom string
	/*TimestampTo
	  Filter for top-up requests performed before or at the given time stamp.

	*/
	TimestampTo string
	/*VoucherID
	  Filter for top-up requests with a specific voucher.

	*/
	VoucherID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the topuplogs get params
func (o *TopuplogsGetParams) WithTimeout(timeout time.Duration) *TopuplogsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the topuplogs get params
func (o *TopuplogsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the topuplogs get params
func (o *TopuplogsGetParams) WithContext(ctx context.Context) *TopuplogsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the topuplogs get params
func (o *TopuplogsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the topuplogs get params
func (o *TopuplogsGetParams) WithHTTPClient(client *http.Client) *TopuplogsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the topuplogs get params
func (o *TopuplogsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAmountAbove adds the amountAbove to the topuplogs get params
func (o *TopuplogsGetParams) WithAmountAbove(amountAbove string) *TopuplogsGetParams {
	o.SetAmountAbove(amountAbove)
	return o
}

// SetAmountAbove adds the amountAbove to the topuplogs get params
func (o *TopuplogsGetParams) SetAmountAbove(amountAbove string) {
	o.AmountAbove = amountAbove
}

// WithAmountBelow adds the amountBelow to the topuplogs get params
func (o *TopuplogsGetParams) WithAmountBelow(amountBelow string) *TopuplogsGetParams {
	o.SetAmountBelow(amountBelow)
	return o
}

// SetAmountBelow adds the amountBelow to the topuplogs get params
func (o *TopuplogsGetParams) SetAmountBelow(amountBelow string) {
	o.AmountBelow = amountBelow
}

// WithContractID adds the contractID to the topuplogs get params
func (o *TopuplogsGetParams) WithContractID(contractID string) *TopuplogsGetParams {
	o.SetContractID(contractID)
	return o
}

// SetContractID adds the contractId to the topuplogs get params
func (o *TopuplogsGetParams) SetContractID(contractID string) {
	o.ContractID = contractID
}

// WithOrderBy adds the orderBy to the topuplogs get params
func (o *TopuplogsGetParams) WithOrderBy(orderBy string) *TopuplogsGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the topuplogs get params
func (o *TopuplogsGetParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the topuplogs get params
func (o *TopuplogsGetParams) WithOrderByDirection(orderByDirection string) *TopuplogsGetParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the topuplogs get params
func (o *TopuplogsGetParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithOutcome adds the outcome to the topuplogs get params
func (o *TopuplogsGetParams) WithOutcome(outcome string) *TopuplogsGetParams {
	o.SetOutcome(outcome)
	return o
}

// SetOutcome adds the outcome to the topuplogs get params
func (o *TopuplogsGetParams) SetOutcome(outcome string) {
	o.Outcome = outcome
}

// WithPage adds the page to the topuplogs get params
func (o *TopuplogsGetParams) WithPage(page string) *TopuplogsGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the topuplogs get params
func (o *TopuplogsGetParams) SetPage(page string) {
	o.Page = page
}

// WithRequestToken adds the requestToken to the topuplogs get params
func (o *TopuplogsGetParams) WithRequestToken(requestToken string) *TopuplogsGetParams {
	o.SetRequestToken(requestToken)
	return o
}

// SetRequestToken adds the requestToken to the topuplogs get params
func (o *TopuplogsGetParams) SetRequestToken(requestToken string) {
	o.RequestToken = requestToken
}

// WithResellerID adds the resellerID to the topuplogs get params
func (o *TopuplogsGetParams) WithResellerID(resellerID string) *TopuplogsGetParams {
	o.SetResellerID(resellerID)
	return o
}

// SetResellerID adds the resellerId to the topuplogs get params
func (o *TopuplogsGetParams) SetResellerID(resellerID string) {
	o.ResellerID = resellerID
}

// WithRows adds the rows to the topuplogs get params
func (o *TopuplogsGetParams) WithRows(rows string) *TopuplogsGetParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the topuplogs get params
func (o *TopuplogsGetParams) SetRows(rows string) {
	o.Rows = rows
}

// WithSubscriberID adds the subscriberID to the topuplogs get params
func (o *TopuplogsGetParams) WithSubscriberID(subscriberID string) *TopuplogsGetParams {
	o.SetSubscriberID(subscriberID)
	return o
}

// SetSubscriberID adds the subscriberId to the topuplogs get params
func (o *TopuplogsGetParams) SetSubscriberID(subscriberID string) {
	o.SubscriberID = subscriberID
}

// WithTimestampFrom adds the timestampFrom to the topuplogs get params
func (o *TopuplogsGetParams) WithTimestampFrom(timestampFrom string) *TopuplogsGetParams {
	o.SetTimestampFrom(timestampFrom)
	return o
}

// SetTimestampFrom adds the timestampFrom to the topuplogs get params
func (o *TopuplogsGetParams) SetTimestampFrom(timestampFrom string) {
	o.TimestampFrom = timestampFrom
}

// WithTimestampTo adds the timestampTo to the topuplogs get params
func (o *TopuplogsGetParams) WithTimestampTo(timestampTo string) *TopuplogsGetParams {
	o.SetTimestampTo(timestampTo)
	return o
}

// SetTimestampTo adds the timestampTo to the topuplogs get params
func (o *TopuplogsGetParams) SetTimestampTo(timestampTo string) {
	o.TimestampTo = timestampTo
}

// WithVoucherID adds the voucherID to the topuplogs get params
func (o *TopuplogsGetParams) WithVoucherID(voucherID string) *TopuplogsGetParams {
	o.SetVoucherID(voucherID)
	return o
}

// SetVoucherID adds the voucherId to the topuplogs get params
func (o *TopuplogsGetParams) SetVoucherID(voucherID string) {
	o.VoucherID = voucherID
}

// WriteToRequest writes these params to a swagger request
func (o *TopuplogsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param amount_above
	qrAmountAbove := o.AmountAbove
	qAmountAbove := qrAmountAbove
	if qAmountAbove != "" {
		if err := r.SetQueryParam("amount_above", qAmountAbove); err != nil {
			return err
		}
	}

	// query param amount_below
	qrAmountBelow := o.AmountBelow
	qAmountBelow := qrAmountBelow
	if qAmountBelow != "" {
		if err := r.SetQueryParam("amount_below", qAmountBelow); err != nil {
			return err
		}
	}

	// query param contract_id
	qrContractID := o.ContractID
	qContractID := qrContractID
	if qContractID != "" {
		if err := r.SetQueryParam("contract_id", qContractID); err != nil {
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

	// query param outcome
	qrOutcome := o.Outcome
	qOutcome := qrOutcome
	if qOutcome != "" {
		if err := r.SetQueryParam("outcome", qOutcome); err != nil {
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

	// query param request_token
	qrRequestToken := o.RequestToken
	qRequestToken := qrRequestToken
	if qRequestToken != "" {
		if err := r.SetQueryParam("request_token", qRequestToken); err != nil {
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

	// query param voucher_id
	qrVoucherID := o.VoucherID
	qVoucherID := qrVoucherID
	if qVoucherID != "" {
		if err := r.SetQueryParam("voucher_id", qVoucherID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
