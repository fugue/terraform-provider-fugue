// Code generated by go-swagger; DO NOT EDIT.

package scans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewListScansParams creates a new ListScansParams object
// with the default values initialized.
func NewListScansParams() *ListScansParams {
	var (
		maxItemsDefault       = int64(100)
		offsetDefault         = int64(0)
		orderByDefault        = string("created_at")
		orderDirectionDefault = string("desc")
	)
	return &ListScansParams{
		MaxItems:       &maxItemsDefault,
		Offset:         &offsetDefault,
		OrderBy:        &orderByDefault,
		OrderDirection: &orderDirectionDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListScansParamsWithTimeout creates a new ListScansParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListScansParamsWithTimeout(timeout time.Duration) *ListScansParams {
	var (
		maxItemsDefault       = int64(100)
		offsetDefault         = int64(0)
		orderByDefault        = string("created_at")
		orderDirectionDefault = string("desc")
	)
	return &ListScansParams{
		MaxItems:       &maxItemsDefault,
		Offset:         &offsetDefault,
		OrderBy:        &orderByDefault,
		OrderDirection: &orderDirectionDefault,

		timeout: timeout,
	}
}

// NewListScansParamsWithContext creates a new ListScansParams object
// with the default values initialized, and the ability to set a context for a request
func NewListScansParamsWithContext(ctx context.Context) *ListScansParams {
	var (
		maxItemsDefault       = int64(100)
		offsetDefault         = int64(0)
		orderByDefault        = string("created_at")
		orderDirectionDefault = string("desc")
	)
	return &ListScansParams{
		MaxItems:       &maxItemsDefault,
		Offset:         &offsetDefault,
		OrderBy:        &orderByDefault,
		OrderDirection: &orderDirectionDefault,

		Context: ctx,
	}
}

// NewListScansParamsWithHTTPClient creates a new ListScansParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListScansParamsWithHTTPClient(client *http.Client) *ListScansParams {
	var (
		maxItemsDefault       = int64(100)
		offsetDefault         = int64(0)
		orderByDefault        = string("created_at")
		orderDirectionDefault = string("desc")
	)
	return &ListScansParams{
		MaxItems:       &maxItemsDefault,
		Offset:         &offsetDefault,
		OrderBy:        &orderByDefault,
		OrderDirection: &orderDirectionDefault,
		HTTPClient:     client,
	}
}

/*ListScansParams contains all the parameters to send to the API endpoint
for the list scans operation typically these are written to a http.Request
*/
type ListScansParams struct {

	/*EnvironmentID
	  ID of the environment to retrieve scans for.

	*/
	EnvironmentID string
	/*MaxItems
	  Maximum number of items to return.

	*/
	MaxItems *int64
	/*Offset
	  Number of items to skip before returning. This parameter is used when the number of items spans multiple pages.

	*/
	Offset *int64
	/*OrderBy
	  Field to sort the items by.

	*/
	OrderBy *string
	/*OrderDirection
	  Direction to sort the items in.

	*/
	OrderDirection *string
	/*RangeFrom
	  Earliest created_at time to return scans from.

	*/
	RangeFrom *int64
	/*RangeTo
	  Latest created_at time to return scans from.

	*/
	RangeTo *int64
	/*Status
	  Status to filter by. When not specified, all statuses will be returned.

	*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list scans params
func (o *ListScansParams) WithTimeout(timeout time.Duration) *ListScansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list scans params
func (o *ListScansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list scans params
func (o *ListScansParams) WithContext(ctx context.Context) *ListScansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list scans params
func (o *ListScansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list scans params
func (o *ListScansParams) WithHTTPClient(client *http.Client) *ListScansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list scans params
func (o *ListScansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the list scans params
func (o *ListScansParams) WithEnvironmentID(environmentID string) *ListScansParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the list scans params
func (o *ListScansParams) SetEnvironmentID(environmentID string) {
	o.EnvironmentID = environmentID
}

// WithMaxItems adds the maxItems to the list scans params
func (o *ListScansParams) WithMaxItems(maxItems *int64) *ListScansParams {
	o.SetMaxItems(maxItems)
	return o
}

// SetMaxItems adds the maxItems to the list scans params
func (o *ListScansParams) SetMaxItems(maxItems *int64) {
	o.MaxItems = maxItems
}

// WithOffset adds the offset to the list scans params
func (o *ListScansParams) WithOffset(offset *int64) *ListScansParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list scans params
func (o *ListScansParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrderBy adds the orderBy to the list scans params
func (o *ListScansParams) WithOrderBy(orderBy *string) *ListScansParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the list scans params
func (o *ListScansParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithOrderDirection adds the orderDirection to the list scans params
func (o *ListScansParams) WithOrderDirection(orderDirection *string) *ListScansParams {
	o.SetOrderDirection(orderDirection)
	return o
}

// SetOrderDirection adds the orderDirection to the list scans params
func (o *ListScansParams) SetOrderDirection(orderDirection *string) {
	o.OrderDirection = orderDirection
}

// WithRangeFrom adds the rangeFrom to the list scans params
func (o *ListScansParams) WithRangeFrom(rangeFrom *int64) *ListScansParams {
	o.SetRangeFrom(rangeFrom)
	return o
}

// SetRangeFrom adds the rangeFrom to the list scans params
func (o *ListScansParams) SetRangeFrom(rangeFrom *int64) {
	o.RangeFrom = rangeFrom
}

// WithRangeTo adds the rangeTo to the list scans params
func (o *ListScansParams) WithRangeTo(rangeTo *int64) *ListScansParams {
	o.SetRangeTo(rangeTo)
	return o
}

// SetRangeTo adds the rangeTo to the list scans params
func (o *ListScansParams) SetRangeTo(rangeTo *int64) {
	o.RangeTo = rangeTo
}

// WithStatus adds the status to the list scans params
func (o *ListScansParams) WithStatus(status []string) *ListScansParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the list scans params
func (o *ListScansParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *ListScansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param environment_id
	qrEnvironmentID := o.EnvironmentID
	qEnvironmentID := qrEnvironmentID
	if qEnvironmentID != "" {
		if err := r.SetQueryParam("environment_id", qEnvironmentID); err != nil {
			return err
		}
	}

	if o.MaxItems != nil {

		// query param max_items
		var qrMaxItems int64
		if o.MaxItems != nil {
			qrMaxItems = *o.MaxItems
		}
		qMaxItems := swag.FormatInt64(qrMaxItems)
		if qMaxItems != "" {
			if err := r.SetQueryParam("max_items", qMaxItems); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.OrderBy != nil {

		// query param order_by
		var qrOrderBy string
		if o.OrderBy != nil {
			qrOrderBy = *o.OrderBy
		}
		qOrderBy := qrOrderBy
		if qOrderBy != "" {
			if err := r.SetQueryParam("order_by", qOrderBy); err != nil {
				return err
			}
		}

	}

	if o.OrderDirection != nil {

		// query param order_direction
		var qrOrderDirection string
		if o.OrderDirection != nil {
			qrOrderDirection = *o.OrderDirection
		}
		qOrderDirection := qrOrderDirection
		if qOrderDirection != "" {
			if err := r.SetQueryParam("order_direction", qOrderDirection); err != nil {
				return err
			}
		}

	}

	if o.RangeFrom != nil {

		// query param range_from
		var qrRangeFrom int64
		if o.RangeFrom != nil {
			qrRangeFrom = *o.RangeFrom
		}
		qRangeFrom := swag.FormatInt64(qrRangeFrom)
		if qRangeFrom != "" {
			if err := r.SetQueryParam("range_from", qRangeFrom); err != nil {
				return err
			}
		}

	}

	if o.RangeTo != nil {

		// query param range_to
		var qrRangeTo int64
		if o.RangeTo != nil {
			qrRangeTo = *o.RangeTo
		}
		qRangeTo := swag.FormatInt64(qrRangeTo)
		if qRangeTo != "" {
			if err := r.SetQueryParam("range_to", qRangeTo); err != nil {
				return err
			}
		}

	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "multi")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
