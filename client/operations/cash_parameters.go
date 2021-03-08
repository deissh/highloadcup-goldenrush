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

// NewCashParams creates a new CashParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCashParams() *CashParams {
	return &CashParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCashParamsWithTimeout creates a new CashParams object
// with the ability to set a timeout on a request.
func NewCashParamsWithTimeout(timeout time.Duration) *CashParams {
	return &CashParams{
		timeout: timeout,
	}
}

// NewCashParamsWithContext creates a new CashParams object
// with the ability to set a context for a request.
func NewCashParamsWithContext(ctx context.Context) *CashParams {
	return &CashParams{
		Context: ctx,
	}
}

// NewCashParamsWithHTTPClient creates a new CashParams object
// with the ability to set a custom HTTPClient for a request.
func NewCashParamsWithHTTPClient(client *http.Client) *CashParams {
	return &CashParams{
		HTTPClient: client,
	}
}

/* CashParams contains all the parameters to send to the API endpoint
   for the cash operation.

   Typically these are written to a http.Request.
*/
type CashParams struct {

	/* Args.

	   Treasure for exchange.
	*/
	Args models.Treasure

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cash params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CashParams) WithDefaults() *CashParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cash params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CashParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cash params
func (o *CashParams) WithTimeout(timeout time.Duration) *CashParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cash params
func (o *CashParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cash params
func (o *CashParams) WithContext(ctx context.Context) *CashParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cash params
func (o *CashParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cash params
func (o *CashParams) WithHTTPClient(client *http.Client) *CashParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cash params
func (o *CashParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArgs adds the args to the cash params
func (o *CashParams) WithArgs(args models.Treasure) *CashParams {
	o.SetArgs(args)
	return o
}

// SetArgs adds the args to the cash params
func (o *CashParams) SetArgs(args models.Treasure) {
	o.Args = args
}

// WriteToRequest writes these params to a swagger request
func (o *CashParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Args); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
