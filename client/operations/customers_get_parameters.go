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

// NewCustomersGetParams creates a new CustomersGetParams object
// with the default values initialized.
func NewCustomersGetParams() *CustomersGetParams {
	var ()
	return &CustomersGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomersGetParamsWithTimeout creates a new CustomersGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomersGetParamsWithTimeout(timeout time.Duration) *CustomersGetParams {
	var ()
	return &CustomersGetParams{

		timeout: timeout,
	}
}

// NewCustomersGetParamsWithContext creates a new CustomersGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomersGetParamsWithContext(ctx context.Context) *CustomersGetParams {
	var ()
	return &CustomersGetParams{

		Context: ctx,
	}
}

// NewCustomersGetParamsWithHTTPClient creates a new CustomersGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomersGetParamsWithHTTPClient(client *http.Client) *CustomersGetParams {
	var ()
	return &CustomersGetParams{
		HTTPClient: client,
	}
}

/*CustomersGetParams contains all the parameters to send to the API endpoint
for the customers get operation typically these are written to a http.Request
*/
type CustomersGetParams struct {

	/*ActivateTimestampGt
	  Filter for customers with activate_timestamp greater then specified value

	*/
	ActivateTimestampGt string
	/*ActivateTimestampLt
	  Filter for customers with activate_timestamp less then specified value

	*/
	ActivateTimestampLt string
	/*ContactID
	  Filter for customers belonging to a specific contact

	*/
	ContactID string
	/*CreateTimestampGt
	  Filter for customers with create_timestamp greater then specified value

	*/
	CreateTimestampGt string
	/*CreateTimestampLt
	  Filter for customers with create_timestamp less then specified value

	*/
	CreateTimestampLt string
	/*ExternalID
	  Filter for customer with specific external_id

	*/
	ExternalID string
	/*ModifyTimestampGt
	  Filter for customers with modify_timestamp greater then specified value

	*/
	ModifyTimestampGt string
	/*ModifyTimestampLt
	  Filter for customers with modify_timestamp less then specified value

	*/
	ModifyTimestampLt string
	/*NotStatus
	  Filter for customers not having a specific status (comma-separated list of statuses to exclude possible)

	*/
	NotStatus string
	/*OrderBy
	  Order collection by a specific attribute.

	*/
	OrderBy string
	/*OrderByDirection
	  Direction which the collection should be ordered by. Possible values are: asc (default), desc.

	*/
	OrderByDirection string
	/*PackageID
	  Filter for customers with specific profile package id

	*/
	PackageID string
	/*Page
	  Pagination page which should be displayed (default: 1)

	*/
	Page string
	/*ResellerID
	  Filter for customers belonging to a specific reseller

	*/
	ResellerID string
	/*Rows
	  Number of rows in one pagination page (default: 10)

	*/
	Rows string
	/*Status
	  Filter for customers with a specific status (comma-separated list of statuses to include possible)

	*/
	Status string
	/*TerminateTimestampGt
	  Filter for customers with terminate_timestamp greater then specified value

	*/
	TerminateTimestampGt string
	/*TerminateTimestampLt
	  Filter for customers with terminate_timestamp less then specified value

	*/
	TerminateTimestampLt string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customers get params
func (o *CustomersGetParams) WithTimeout(timeout time.Duration) *CustomersGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customers get params
func (o *CustomersGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customers get params
func (o *CustomersGetParams) WithContext(ctx context.Context) *CustomersGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customers get params
func (o *CustomersGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customers get params
func (o *CustomersGetParams) WithHTTPClient(client *http.Client) *CustomersGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customers get params
func (o *CustomersGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActivateTimestampGt adds the activateTimestampGt to the customers get params
func (o *CustomersGetParams) WithActivateTimestampGt(activateTimestampGt string) *CustomersGetParams {
	o.SetActivateTimestampGt(activateTimestampGt)
	return o
}

// SetActivateTimestampGt adds the activateTimestampGt to the customers get params
func (o *CustomersGetParams) SetActivateTimestampGt(activateTimestampGt string) {
	o.ActivateTimestampGt = activateTimestampGt
}

// WithActivateTimestampLt adds the activateTimestampLt to the customers get params
func (o *CustomersGetParams) WithActivateTimestampLt(activateTimestampLt string) *CustomersGetParams {
	o.SetActivateTimestampLt(activateTimestampLt)
	return o
}

// SetActivateTimestampLt adds the activateTimestampLt to the customers get params
func (o *CustomersGetParams) SetActivateTimestampLt(activateTimestampLt string) {
	o.ActivateTimestampLt = activateTimestampLt
}

// WithContactID adds the contactID to the customers get params
func (o *CustomersGetParams) WithContactID(contactID string) *CustomersGetParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the customers get params
func (o *CustomersGetParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithCreateTimestampGt adds the createTimestampGt to the customers get params
func (o *CustomersGetParams) WithCreateTimestampGt(createTimestampGt string) *CustomersGetParams {
	o.SetCreateTimestampGt(createTimestampGt)
	return o
}

// SetCreateTimestampGt adds the createTimestampGt to the customers get params
func (o *CustomersGetParams) SetCreateTimestampGt(createTimestampGt string) {
	o.CreateTimestampGt = createTimestampGt
}

// WithCreateTimestampLt adds the createTimestampLt to the customers get params
func (o *CustomersGetParams) WithCreateTimestampLt(createTimestampLt string) *CustomersGetParams {
	o.SetCreateTimestampLt(createTimestampLt)
	return o
}

// SetCreateTimestampLt adds the createTimestampLt to the customers get params
func (o *CustomersGetParams) SetCreateTimestampLt(createTimestampLt string) {
	o.CreateTimestampLt = createTimestampLt
}

// WithExternalID adds the externalID to the customers get params
func (o *CustomersGetParams) WithExternalID(externalID string) *CustomersGetParams {
	o.SetExternalID(externalID)
	return o
}

// SetExternalID adds the externalId to the customers get params
func (o *CustomersGetParams) SetExternalID(externalID string) {
	o.ExternalID = externalID
}

// WithModifyTimestampGt adds the modifyTimestampGt to the customers get params
func (o *CustomersGetParams) WithModifyTimestampGt(modifyTimestampGt string) *CustomersGetParams {
	o.SetModifyTimestampGt(modifyTimestampGt)
	return o
}

// SetModifyTimestampGt adds the modifyTimestampGt to the customers get params
func (o *CustomersGetParams) SetModifyTimestampGt(modifyTimestampGt string) {
	o.ModifyTimestampGt = modifyTimestampGt
}

// WithModifyTimestampLt adds the modifyTimestampLt to the customers get params
func (o *CustomersGetParams) WithModifyTimestampLt(modifyTimestampLt string) *CustomersGetParams {
	o.SetModifyTimestampLt(modifyTimestampLt)
	return o
}

// SetModifyTimestampLt adds the modifyTimestampLt to the customers get params
func (o *CustomersGetParams) SetModifyTimestampLt(modifyTimestampLt string) {
	o.ModifyTimestampLt = modifyTimestampLt
}

// WithNotStatus adds the notStatus to the customers get params
func (o *CustomersGetParams) WithNotStatus(notStatus string) *CustomersGetParams {
	o.SetNotStatus(notStatus)
	return o
}

// SetNotStatus adds the notStatus to the customers get params
func (o *CustomersGetParams) SetNotStatus(notStatus string) {
	o.NotStatus = notStatus
}

// WithOrderBy adds the orderBy to the customers get params
func (o *CustomersGetParams) WithOrderBy(orderBy string) *CustomersGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the customers get params
func (o *CustomersGetParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the customers get params
func (o *CustomersGetParams) WithOrderByDirection(orderByDirection string) *CustomersGetParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the customers get params
func (o *CustomersGetParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPackageID adds the packageID to the customers get params
func (o *CustomersGetParams) WithPackageID(packageID string) *CustomersGetParams {
	o.SetPackageID(packageID)
	return o
}

// SetPackageID adds the packageId to the customers get params
func (o *CustomersGetParams) SetPackageID(packageID string) {
	o.PackageID = packageID
}

// WithPage adds the page to the customers get params
func (o *CustomersGetParams) WithPage(page string) *CustomersGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the customers get params
func (o *CustomersGetParams) SetPage(page string) {
	o.Page = page
}

// WithResellerID adds the resellerID to the customers get params
func (o *CustomersGetParams) WithResellerID(resellerID string) *CustomersGetParams {
	o.SetResellerID(resellerID)
	return o
}

// SetResellerID adds the resellerId to the customers get params
func (o *CustomersGetParams) SetResellerID(resellerID string) {
	o.ResellerID = resellerID
}

// WithRows adds the rows to the customers get params
func (o *CustomersGetParams) WithRows(rows string) *CustomersGetParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the customers get params
func (o *CustomersGetParams) SetRows(rows string) {
	o.Rows = rows
}

// WithStatus adds the status to the customers get params
func (o *CustomersGetParams) WithStatus(status string) *CustomersGetParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the customers get params
func (o *CustomersGetParams) SetStatus(status string) {
	o.Status = status
}

// WithTerminateTimestampGt adds the terminateTimestampGt to the customers get params
func (o *CustomersGetParams) WithTerminateTimestampGt(terminateTimestampGt string) *CustomersGetParams {
	o.SetTerminateTimestampGt(terminateTimestampGt)
	return o
}

// SetTerminateTimestampGt adds the terminateTimestampGt to the customers get params
func (o *CustomersGetParams) SetTerminateTimestampGt(terminateTimestampGt string) {
	o.TerminateTimestampGt = terminateTimestampGt
}

// WithTerminateTimestampLt adds the terminateTimestampLt to the customers get params
func (o *CustomersGetParams) WithTerminateTimestampLt(terminateTimestampLt string) *CustomersGetParams {
	o.SetTerminateTimestampLt(terminateTimestampLt)
	return o
}

// SetTerminateTimestampLt adds the terminateTimestampLt to the customers get params
func (o *CustomersGetParams) SetTerminateTimestampLt(terminateTimestampLt string) {
	o.TerminateTimestampLt = terminateTimestampLt
}

// WriteToRequest writes these params to a swagger request
func (o *CustomersGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param activate_timestamp_gt
	qrActivateTimestampGt := o.ActivateTimestampGt
	qActivateTimestampGt := qrActivateTimestampGt
	if qActivateTimestampGt != "" {
		if err := r.SetQueryParam("activate_timestamp_gt", qActivateTimestampGt); err != nil {
			return err
		}
	}

	// query param activate_timestamp_lt
	qrActivateTimestampLt := o.ActivateTimestampLt
	qActivateTimestampLt := qrActivateTimestampLt
	if qActivateTimestampLt != "" {
		if err := r.SetQueryParam("activate_timestamp_lt", qActivateTimestampLt); err != nil {
			return err
		}
	}

	// query param contact_id
	qrContactID := o.ContactID
	qContactID := qrContactID
	if qContactID != "" {
		if err := r.SetQueryParam("contact_id", qContactID); err != nil {
			return err
		}
	}

	// query param create_timestamp_gt
	qrCreateTimestampGt := o.CreateTimestampGt
	qCreateTimestampGt := qrCreateTimestampGt
	if qCreateTimestampGt != "" {
		if err := r.SetQueryParam("create_timestamp_gt", qCreateTimestampGt); err != nil {
			return err
		}
	}

	// query param create_timestamp_lt
	qrCreateTimestampLt := o.CreateTimestampLt
	qCreateTimestampLt := qrCreateTimestampLt
	if qCreateTimestampLt != "" {
		if err := r.SetQueryParam("create_timestamp_lt", qCreateTimestampLt); err != nil {
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

	// query param modify_timestamp_gt
	qrModifyTimestampGt := o.ModifyTimestampGt
	qModifyTimestampGt := qrModifyTimestampGt
	if qModifyTimestampGt != "" {
		if err := r.SetQueryParam("modify_timestamp_gt", qModifyTimestampGt); err != nil {
			return err
		}
	}

	// query param modify_timestamp_lt
	qrModifyTimestampLt := o.ModifyTimestampLt
	qModifyTimestampLt := qrModifyTimestampLt
	if qModifyTimestampLt != "" {
		if err := r.SetQueryParam("modify_timestamp_lt", qModifyTimestampLt); err != nil {
			return err
		}
	}

	// query param not_status
	qrNotStatus := o.NotStatus
	qNotStatus := qrNotStatus
	if qNotStatus != "" {
		if err := r.SetQueryParam("not_status", qNotStatus); err != nil {
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

	// query param package_id
	qrPackageID := o.PackageID
	qPackageID := qrPackageID
	if qPackageID != "" {
		if err := r.SetQueryParam("package_id", qPackageID); err != nil {
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

	// query param status
	qrStatus := o.Status
	qStatus := qrStatus
	if qStatus != "" {
		if err := r.SetQueryParam("status", qStatus); err != nil {
			return err
		}
	}

	// query param terminate_timestamp_gt
	qrTerminateTimestampGt := o.TerminateTimestampGt
	qTerminateTimestampGt := qrTerminateTimestampGt
	if qTerminateTimestampGt != "" {
		if err := r.SetQueryParam("terminate_timestamp_gt", qTerminateTimestampGt); err != nil {
			return err
		}
	}

	// query param terminate_timestamp_lt
	qrTerminateTimestampLt := o.TerminateTimestampLt
	qTerminateTimestampLt := qrTerminateTimestampLt
	if qTerminateTimestampLt != "" {
		if err := r.SetQueryParam("terminate_timestamp_lt", qTerminateTimestampLt); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
