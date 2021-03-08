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
	"github.com/go-openapi/strfmt"

	"github.com/deissh/highloadcup-goldenrush/models"
)

// NewIssueLicenseParams creates a new IssueLicenseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssueLicenseParams() *IssueLicenseParams {
	return &IssueLicenseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssueLicenseParamsWithTimeout creates a new IssueLicenseParams object
// with the ability to set a timeout on a request.
func NewIssueLicenseParamsWithTimeout(timeout time.Duration) *IssueLicenseParams {
	return &IssueLicenseParams{
		timeout: timeout,
	}
}

// NewIssueLicenseParamsWithContext creates a new IssueLicenseParams object
// with the ability to set a context for a request.
func NewIssueLicenseParamsWithContext(ctx context.Context) *IssueLicenseParams {
	return &IssueLicenseParams{
		Context: ctx,
	}
}

// NewIssueLicenseParamsWithHTTPClient creates a new IssueLicenseParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssueLicenseParamsWithHTTPClient(client *http.Client) *IssueLicenseParams {
	return &IssueLicenseParams{
		HTTPClient: client,
	}
}

/* IssueLicenseParams contains all the parameters to send to the API endpoint
   for the issue license operation.

   Typically these are written to a http.Request.
*/
type IssueLicenseParams struct {

	/* Args.

	   Amount of money to spend for a license. Empty array for get free license. Maximum 10 active licenses
	*/
	Args models.Wallet

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the issue license params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueLicenseParams) WithDefaults() *IssueLicenseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issue license params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueLicenseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issue license params
func (o *IssueLicenseParams) WithTimeout(timeout time.Duration) *IssueLicenseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue license params
func (o *IssueLicenseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue license params
func (o *IssueLicenseParams) WithContext(ctx context.Context) *IssueLicenseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue license params
func (o *IssueLicenseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue license params
func (o *IssueLicenseParams) WithHTTPClient(client *http.Client) *IssueLicenseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue license params
func (o *IssueLicenseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArgs adds the args to the issue license params
func (o *IssueLicenseParams) WithArgs(args models.Wallet) *IssueLicenseParams {
	o.SetArgs(args)
	return o
}

// SetArgs adds the args to the issue license params
func (o *IssueLicenseParams) SetArgs(args models.Wallet) {
	o.Args = args
}

// WriteToRequest writes these params to a swagger request
func (o *IssueLicenseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Args != nil {
		if err := r.SetBodyParam(o.Args); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}