// Code generated by go-swagger; DO NOT EDIT.

package invites

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

// NewListInvitesParams creates a new ListInvitesParams object
// with the default values initialized.
func NewListInvitesParams() *ListInvitesParams {
	var (
		maxItemsDefault       = int64(100)
		offsetDefault         = int64(0)
		orderDirectionDefault = string("desc")
	)
	return &ListInvitesParams{
		MaxItems:       &maxItemsDefault,
		Offset:         &offsetDefault,
		OrderDirection: &orderDirectionDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListInvitesParamsWithTimeout creates a new ListInvitesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListInvitesParamsWithTimeout(timeout time.Duration) *ListInvitesParams {
	var (
		maxItemsDefault       = int64(100)
		offsetDefault         = int64(0)
		orderDirectionDefault = string("desc")
	)
	return &ListInvitesParams{
		MaxItems:       &maxItemsDefault,
		Offset:         &offsetDefault,
		OrderDirection: &orderDirectionDefault,

		timeout: timeout,
	}
}

// NewListInvitesParamsWithContext creates a new ListInvitesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListInvitesParamsWithContext(ctx context.Context) *ListInvitesParams {
	var (
		maxItemsDefault       = int64(100)
		offsetDefault         = int64(0)
		orderDirectionDefault = string("desc")
	)
	return &ListInvitesParams{
		MaxItems:       &maxItemsDefault,
		Offset:         &offsetDefault,
		OrderDirection: &orderDirectionDefault,

		Context: ctx,
	}
}

// NewListInvitesParamsWithHTTPClient creates a new ListInvitesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListInvitesParamsWithHTTPClient(client *http.Client) *ListInvitesParams {
	var (
		maxItemsDefault       = int64(100)
		offsetDefault         = int64(0)
		orderDirectionDefault = string("desc")
	)
	return &ListInvitesParams{
		MaxItems:       &maxItemsDefault,
		Offset:         &offsetDefault,
		OrderDirection: &orderDirectionDefault,
		HTTPClient:     client,
	}
}

/*ListInvitesParams contains all the parameters to send to the API endpoint
for the list invites operation typically these are written to a http.Request
*/
type ListInvitesParams struct {

	/*Email
	  Used to filter list to a single invite by email.

	*/
	Email *string
	/*MaxItems
	  Maximum number of items to return.

	*/
	MaxItems *int64
	/*Offset
	  Number of items to skip before returning. This parameter is used when the number of items spans multiple pages.

	*/
	Offset *int64
	/*OrderDirection
	  Direction to sort the items in.

	*/
	OrderDirection *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list invites params
func (o *ListInvitesParams) WithTimeout(timeout time.Duration) *ListInvitesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list invites params
func (o *ListInvitesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list invites params
func (o *ListInvitesParams) WithContext(ctx context.Context) *ListInvitesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list invites params
func (o *ListInvitesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list invites params
func (o *ListInvitesParams) WithHTTPClient(client *http.Client) *ListInvitesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list invites params
func (o *ListInvitesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmail adds the email to the list invites params
func (o *ListInvitesParams) WithEmail(email *string) *ListInvitesParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the list invites params
func (o *ListInvitesParams) SetEmail(email *string) {
	o.Email = email
}

// WithMaxItems adds the maxItems to the list invites params
func (o *ListInvitesParams) WithMaxItems(maxItems *int64) *ListInvitesParams {
	o.SetMaxItems(maxItems)
	return o
}

// SetMaxItems adds the maxItems to the list invites params
func (o *ListInvitesParams) SetMaxItems(maxItems *int64) {
	o.MaxItems = maxItems
}

// WithOffset adds the offset to the list invites params
func (o *ListInvitesParams) WithOffset(offset *int64) *ListInvitesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list invites params
func (o *ListInvitesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrderDirection adds the orderDirection to the list invites params
func (o *ListInvitesParams) WithOrderDirection(orderDirection *string) *ListInvitesParams {
	o.SetOrderDirection(orderDirection)
	return o
}

// SetOrderDirection adds the orderDirection to the list invites params
func (o *ListInvitesParams) SetOrderDirection(orderDirection *string) {
	o.OrderDirection = orderDirection
}

// WriteToRequest writes these params to a swagger request
func (o *ListInvitesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Email != nil {

		// query param email
		var qrEmail string
		if o.Email != nil {
			qrEmail = *o.Email
		}
		qEmail := qrEmail
		if qEmail != "" {
			if err := r.SetQueryParam("email", qEmail); err != nil {
				return err
			}
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
