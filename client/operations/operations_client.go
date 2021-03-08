// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Cash(params *CashParams, opts ...ClientOption) (*CashOK, error)

	Dig(params *DigParams, opts ...ClientOption) (*DigOK, error)

	ExploreArea(params *ExploreAreaParams, opts ...ClientOption) (*ExploreAreaOK, error)

	GetBalance(params *GetBalanceParams, opts ...ClientOption) (*GetBalanceOK, error)

	HealthCheck(params *HealthCheckParams, opts ...ClientOption) (*HealthCheckOK, error)

	IssueLicense(params *IssueLicenseParams, opts ...ClientOption) (*IssueLicenseOK, error)

	ListLicenses(params *ListLicensesParams, opts ...ClientOption) (*ListLicensesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  Cash Exchange provided treasure for money.
*/
func (a *Client) Cash(params *CashParams, opts ...ClientOption) (*CashOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCashParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cash",
		Method:             "POST",
		PathPattern:        "/cash",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CashReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CashOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CashDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Dig Dig at given point and depth, returns found treasures.
*/
func (a *Client) Dig(params *DigParams, opts ...ClientOption) (*DigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dig",
		Method:             "POST",
		PathPattern:        "/dig",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ExploreArea Returns amount of treasures in the provided area at full depth.
*/
func (a *Client) ExploreArea(params *ExploreAreaParams, opts ...ClientOption) (*ExploreAreaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExploreAreaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "exploreArea",
		Method:             "POST",
		PathPattern:        "/explore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ExploreAreaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExploreAreaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ExploreAreaDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetBalance Returns a current balance.
*/
func (a *Client) GetBalance(params *GetBalanceParams, opts ...ClientOption) (*GetBalanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBalanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBalance",
		Method:             "GET",
		PathPattern:        "/balance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBalanceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBalanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBalanceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  HealthCheck Returns 200 if service works okay.
*/
func (a *Client) HealthCheck(params *HealthCheckParams, opts ...ClientOption) (*HealthCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHealthCheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "healthCheck",
		Method:             "GET",
		PathPattern:        "/health-check",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HealthCheckReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HealthCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HealthCheckDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  IssueLicense Issue a new license.
*/
func (a *Client) IssueLicense(params *IssueLicenseParams, opts ...ClientOption) (*IssueLicenseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIssueLicenseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "issueLicense",
		Method:             "POST",
		PathPattern:        "/licenses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IssueLicenseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IssueLicenseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IssueLicenseDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListLicenses Returns a list of issued licenses.
*/
func (a *Client) ListLicenses(params *ListLicensesParams, opts ...ClientOption) (*ListLicensesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListLicensesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listLicenses",
		Method:             "GET",
		PathPattern:        "/licenses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListLicensesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListLicensesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListLicensesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}