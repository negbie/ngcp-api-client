// Code generated by go-swagger; DO NOT EDIT.

package topuplogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new topuplogs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for topuplogs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
TopuplogsByIDGet gets a specific topup log

TODO: Add Description
*/
func (a *Client) TopuplogsByIDGet(params *TopuplogsByIDGetParams) (*TopuplogsByIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTopuplogsByIDGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TopuplogsByIdGet",
		Method:             "GET",
		PathPattern:        "/topuplogs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TopuplogsByIDGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TopuplogsByIDGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TopuplogsByIdGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TopuplogsGet gets topup log items

TODO: Add Description
*/
func (a *Client) TopuplogsGet(params *TopuplogsGetParams) (*TopuplogsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTopuplogsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TopuplogsGet",
		Method:             "GET",
		PathPattern:        "/topuplogs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TopuplogsGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TopuplogsGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TopuplogsGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
