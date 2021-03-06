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
)

// NewCreateScanParams creates a new CreateScanParams object
// with the default values initialized.
func NewCreateScanParams() *CreateScanParams {
	var ()
	return &CreateScanParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateScanParamsWithTimeout creates a new CreateScanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateScanParamsWithTimeout(timeout time.Duration) *CreateScanParams {
	var ()
	return &CreateScanParams{

		timeout: timeout,
	}
}

// NewCreateScanParamsWithContext creates a new CreateScanParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateScanParamsWithContext(ctx context.Context) *CreateScanParams {
	var ()
	return &CreateScanParams{

		Context: ctx,
	}
}

// NewCreateScanParamsWithHTTPClient creates a new CreateScanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateScanParamsWithHTTPClient(client *http.Client) *CreateScanParams {
	var ()
	return &CreateScanParams{
		HTTPClient: client,
	}
}

/*CreateScanParams contains all the parameters to send to the API endpoint
for the create scan operation typically these are written to a http.Request
*/
type CreateScanParams struct {

	/*EnvironmentID
	  ID of the environment to scan.

	*/
	EnvironmentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create scan params
func (o *CreateScanParams) WithTimeout(timeout time.Duration) *CreateScanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create scan params
func (o *CreateScanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create scan params
func (o *CreateScanParams) WithContext(ctx context.Context) *CreateScanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create scan params
func (o *CreateScanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create scan params
func (o *CreateScanParams) WithHTTPClient(client *http.Client) *CreateScanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create scan params
func (o *CreateScanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the create scan params
func (o *CreateScanParams) WithEnvironmentID(environmentID string) *CreateScanParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the create scan params
func (o *CreateScanParams) SetEnvironmentID(environmentID string) {
	o.EnvironmentID = environmentID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateScanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
