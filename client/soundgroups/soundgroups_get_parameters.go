// Code generated by go-swagger; DO NOT EDIT.

package soundgroups

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

// NewSoundgroupsGetParams creates a new SoundgroupsGetParams object
// with the default values initialized.
func NewSoundgroupsGetParams() *SoundgroupsGetParams {
	var ()
	return &SoundgroupsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSoundgroupsGetParamsWithTimeout creates a new SoundgroupsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSoundgroupsGetParamsWithTimeout(timeout time.Duration) *SoundgroupsGetParams {
	var ()
	return &SoundgroupsGetParams{

		timeout: timeout,
	}
}

// NewSoundgroupsGetParamsWithContext creates a new SoundgroupsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSoundgroupsGetParamsWithContext(ctx context.Context) *SoundgroupsGetParams {
	var ()
	return &SoundgroupsGetParams{

		Context: ctx,
	}
}

// NewSoundgroupsGetParamsWithHTTPClient creates a new SoundgroupsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSoundgroupsGetParamsWithHTTPClient(client *http.Client) *SoundgroupsGetParams {
	var ()
	return &SoundgroupsGetParams{
		HTTPClient: client,
	}
}

/*SoundgroupsGetParams contains all the parameters to send to the API endpoint
for the soundgroups get operation typically these are written to a http.Request
*/
type SoundgroupsGetParams struct {

	/*Name
	  Filter for sound groups with a specific name

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
	/*Rows
	  Number of rows in one pagination page (default: 10)

	*/
	Rows string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the soundgroups get params
func (o *SoundgroupsGetParams) WithTimeout(timeout time.Duration) *SoundgroupsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the soundgroups get params
func (o *SoundgroupsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the soundgroups get params
func (o *SoundgroupsGetParams) WithContext(ctx context.Context) *SoundgroupsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the soundgroups get params
func (o *SoundgroupsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the soundgroups get params
func (o *SoundgroupsGetParams) WithHTTPClient(client *http.Client) *SoundgroupsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the soundgroups get params
func (o *SoundgroupsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the soundgroups get params
func (o *SoundgroupsGetParams) WithName(name string) *SoundgroupsGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the soundgroups get params
func (o *SoundgroupsGetParams) SetName(name string) {
	o.Name = name
}

// WithOrderBy adds the orderBy to the soundgroups get params
func (o *SoundgroupsGetParams) WithOrderBy(orderBy string) *SoundgroupsGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the soundgroups get params
func (o *SoundgroupsGetParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the soundgroups get params
func (o *SoundgroupsGetParams) WithOrderByDirection(orderByDirection string) *SoundgroupsGetParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the soundgroups get params
func (o *SoundgroupsGetParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the soundgroups get params
func (o *SoundgroupsGetParams) WithPage(page string) *SoundgroupsGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the soundgroups get params
func (o *SoundgroupsGetParams) SetPage(page string) {
	o.Page = page
}

// WithRows adds the rows to the soundgroups get params
func (o *SoundgroupsGetParams) WithRows(rows string) *SoundgroupsGetParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the soundgroups get params
func (o *SoundgroupsGetParams) SetRows(rows string) {
	o.Rows = rows
}

// WriteToRequest writes these params to a swagger request
func (o *SoundgroupsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
