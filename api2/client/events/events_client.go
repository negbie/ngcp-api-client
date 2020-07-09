// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new events API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for events API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetEventitems(params *GetEventitemsParams) (*GetEventitemsOK, error)

	GetaspecificEvent(params *GetaspecificEventParams) (*GetaspecificEventOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetEventitems gets event items
*/
func (a *Client) GetEventitems(params *GetEventitemsParams) (*GetEventitemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEventitemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEventitems",
		Method:             "GET",
		PathPattern:        "/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEventitemsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEventitemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEventitems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetaspecificEvent gets a specific event
*/
func (a *Client) GetaspecificEvent(params *GetaspecificEventParams) (*GetaspecificEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetaspecificEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetaspecificEvent",
		Method:             "GET",
		PathPattern:        "/events/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetaspecificEventReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetaspecificEventOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetaspecificEvent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}