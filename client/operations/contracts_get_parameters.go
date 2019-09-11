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

// NewContractsGetParams creates a new ContractsGetParams object
// with the default values initialized.
func NewContractsGetParams() *ContractsGetParams {
	var ()
	return &ContractsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewContractsGetParamsWithTimeout creates a new ContractsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewContractsGetParamsWithTimeout(timeout time.Duration) *ContractsGetParams {
	var ()
	return &ContractsGetParams{

		timeout: timeout,
	}
}

// NewContractsGetParamsWithContext creates a new ContractsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewContractsGetParamsWithContext(ctx context.Context) *ContractsGetParams {
	var ()
	return &ContractsGetParams{

		Context: ctx,
	}
}

// NewContractsGetParamsWithHTTPClient creates a new ContractsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContractsGetParamsWithHTTPClient(client *http.Client) *ContractsGetParams {
	var ()
	return &ContractsGetParams{
		HTTPClient: client,
	}
}

/*ContractsGetParams contains all the parameters to send to the API endpoint
for the contracts get operation typically these are written to a http.Request
*/
type ContractsGetParams struct {

	/*ContactID
	  Filter for contracts with a specific contact id

	*/
	ContactID string
	/*ExternalID
	  Filter for contracts with a specific external id

	*/
	ExternalID string
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
	/*Status
	  Filter for contracts with a specific status (except "terminated")

	*/
	Status string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the contracts get params
func (o *ContractsGetParams) WithTimeout(timeout time.Duration) *ContractsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contracts get params
func (o *ContractsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contracts get params
func (o *ContractsGetParams) WithContext(ctx context.Context) *ContractsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contracts get params
func (o *ContractsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contracts get params
func (o *ContractsGetParams) WithHTTPClient(client *http.Client) *ContractsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contracts get params
func (o *ContractsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the contracts get params
func (o *ContractsGetParams) WithContactID(contactID string) *ContractsGetParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the contracts get params
func (o *ContractsGetParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithExternalID adds the externalID to the contracts get params
func (o *ContractsGetParams) WithExternalID(externalID string) *ContractsGetParams {
	o.SetExternalID(externalID)
	return o
}

// SetExternalID adds the externalId to the contracts get params
func (o *ContractsGetParams) SetExternalID(externalID string) {
	o.ExternalID = externalID
}

// WithOrderBy adds the orderBy to the contracts get params
func (o *ContractsGetParams) WithOrderBy(orderBy string) *ContractsGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the contracts get params
func (o *ContractsGetParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the contracts get params
func (o *ContractsGetParams) WithOrderByDirection(orderByDirection string) *ContractsGetParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the contracts get params
func (o *ContractsGetParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the contracts get params
func (o *ContractsGetParams) WithPage(page string) *ContractsGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the contracts get params
func (o *ContractsGetParams) SetPage(page string) {
	o.Page = page
}

// WithRows adds the rows to the contracts get params
func (o *ContractsGetParams) WithRows(rows string) *ContractsGetParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the contracts get params
func (o *ContractsGetParams) SetRows(rows string) {
	o.Rows = rows
}

// WithStatus adds the status to the contracts get params
func (o *ContractsGetParams) WithStatus(status string) *ContractsGetParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the contracts get params
func (o *ContractsGetParams) SetStatus(status string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *ContractsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param contact_id
	qrContactID := o.ContactID
	qContactID := qrContactID
	if qContactID != "" {
		if err := r.SetQueryParam("contact_id", qContactID); err != nil {
			return err
		}
	}

	// query param external_id
	qrExternalID := o.ExternalID
	qExternalID := qrExternalID
	if qExternalID != "" {
		if err := r.SetQueryParam("external_id", qExternalID); err != nil {
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

	// query param status
	qrStatus := o.Status
	qStatus := qrStatus
	if qStatus != "" {
		if err := r.SetQueryParam("status", qStatus); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
