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
)

// NewGetInviteByEmailParams creates a new GetInviteByEmailParams object
// with the default values initialized.
func NewGetInviteByEmailParams() *GetInviteByEmailParams {
	var ()
	return &GetInviteByEmailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInviteByEmailParamsWithTimeout creates a new GetInviteByEmailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInviteByEmailParamsWithTimeout(timeout time.Duration) *GetInviteByEmailParams {
	var ()
	return &GetInviteByEmailParams{

		timeout: timeout,
	}
}

// NewGetInviteByEmailParamsWithContext creates a new GetInviteByEmailParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInviteByEmailParamsWithContext(ctx context.Context) *GetInviteByEmailParams {
	var ()
	return &GetInviteByEmailParams{

		Context: ctx,
	}
}

// NewGetInviteByEmailParamsWithHTTPClient creates a new GetInviteByEmailParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInviteByEmailParamsWithHTTPClient(client *http.Client) *GetInviteByEmailParams {
	var ()
	return &GetInviteByEmailParams{
		HTTPClient: client,
	}
}

/*GetInviteByEmailParams contains all the parameters to send to the API endpoint
for the get invite by email operation typically these are written to a http.Request
*/
type GetInviteByEmailParams struct {

	/*Email
	  email address of the invite

	*/
	Email string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get invite by email params
func (o *GetInviteByEmailParams) WithTimeout(timeout time.Duration) *GetInviteByEmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invite by email params
func (o *GetInviteByEmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invite by email params
func (o *GetInviteByEmailParams) WithContext(ctx context.Context) *GetInviteByEmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invite by email params
func (o *GetInviteByEmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invite by email params
func (o *GetInviteByEmailParams) WithHTTPClient(client *http.Client) *GetInviteByEmailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invite by email params
func (o *GetInviteByEmailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmail adds the email to the get invite by email params
func (o *GetInviteByEmailParams) WithEmail(email string) *GetInviteByEmailParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the get invite by email params
func (o *GetInviteByEmailParams) SetEmail(email string) {
	o.Email = email
}

// WriteToRequest writes these params to a swagger request
func (o *GetInviteByEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param email
	qrEmail := o.Email
	qEmail := qrEmail
	if qEmail != "" {
		if err := r.SetQueryParam("email", qEmail); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
