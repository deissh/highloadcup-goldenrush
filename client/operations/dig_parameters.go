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

// NewDigParams creates a new DigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDigParams() *DigParams {
	return &DigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDigParamsWithTimeout creates a new DigParams object
// with the ability to set a timeout on a request.
func NewDigParamsWithTimeout(timeout time.Duration) *DigParams {
	return &DigParams{
		timeout: timeout,
	}
}

// NewDigParamsWithContext creates a new DigParams object
// with the ability to set a context for a request.
func NewDigParamsWithContext(ctx context.Context) *DigParams {
	return &DigParams{
		Context: ctx,
	}
}

// NewDigParamsWithHTTPClient creates a new DigParams object
// with the ability to set a custom HTTPClient for a request.
func NewDigParamsWithHTTPClient(client *http.Client) *DigParams {
	return &DigParams{
		HTTPClient: client,
	}
}

/* DigParams contains all the parameters to send to the API endpoint
   for the dig operation.

   Typically these are written to a http.Request.
*/
type DigParams struct {

	/* Args.

	   License, place and depth to dig.
	*/
	Args *models.Dig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dig params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DigParams) WithDefaults() *DigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dig params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dig params
func (o *DigParams) WithTimeout(timeout time.Duration) *DigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dig params
func (o *DigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dig params
func (o *DigParams) WithContext(ctx context.Context) *DigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dig params
func (o *DigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dig params
func (o *DigParams) WithHTTPClient(client *http.Client) *DigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dig params
func (o *DigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArgs adds the args to the dig params
func (o *DigParams) WithArgs(args *models.Dig) *DigParams {
	o.SetArgs(args)
	return o
}

// SetArgs adds the args to the dig params
func (o *DigParams) SetArgs(args *models.Dig) {
	o.Args = args
}

// WriteToRequest writes these params to a swagger request
func (o *DigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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