// Code generated by go-swagger; DO NOT EDIT.

package groups

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

	"github.com/fugue/fugue-client/models"
)

// NewBasicAuthParams creates a new BasicAuthParams object
// with the default values initialized.
func NewBasicAuthParams() *BasicAuthParams {
	var ()
	return &BasicAuthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBasicAuthParamsWithTimeout creates a new BasicAuthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBasicAuthParamsWithTimeout(timeout time.Duration) *BasicAuthParams {
	var ()
	return &BasicAuthParams{

		timeout: timeout,
	}
}

// NewBasicAuthParamsWithContext creates a new BasicAuthParams object
// with the default values initialized, and the ability to set a context for a request
func NewBasicAuthParamsWithContext(ctx context.Context) *BasicAuthParams {
	var ()
	return &BasicAuthParams{

		Context: ctx,
	}
}

// NewBasicAuthParamsWithHTTPClient creates a new BasicAuthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBasicAuthParamsWithHTTPClient(client *http.Client) *BasicAuthParams {
	var ()
	return &BasicAuthParams{
		HTTPClient: client,
	}
}

/*BasicAuthParams contains all the parameters to send to the API endpoint
for the basic auth operation typically these are written to a http.Request
*/
type BasicAuthParams struct {

	/*EditGroupAssignments
	  User and Group IDs to be associated.

	*/
	EditGroupAssignments *models.EditUsersGroupAssignmentsInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the basic auth params
func (o *BasicAuthParams) WithTimeout(timeout time.Duration) *BasicAuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the basic auth params
func (o *BasicAuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the basic auth params
func (o *BasicAuthParams) WithContext(ctx context.Context) *BasicAuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the basic auth params
func (o *BasicAuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the basic auth params
func (o *BasicAuthParams) WithHTTPClient(client *http.Client) *BasicAuthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the basic auth params
func (o *BasicAuthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEditGroupAssignments adds the editGroupAssignments to the basic auth params
func (o *BasicAuthParams) WithEditGroupAssignments(editGroupAssignments *models.EditUsersGroupAssignmentsInput) *BasicAuthParams {
	o.SetEditGroupAssignments(editGroupAssignments)
	return o
}

// SetEditGroupAssignments adds the editGroupAssignments to the basic auth params
func (o *BasicAuthParams) SetEditGroupAssignments(editGroupAssignments *models.EditUsersGroupAssignmentsInput) {
	o.EditGroupAssignments = editGroupAssignments
}

// WriteToRequest writes these params to a swagger request
func (o *BasicAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EditGroupAssignments != nil {
		if err := r.SetBodyParam(o.EditGroupAssignments); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
