// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new conversations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for conversations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ConversationsByIDGet gets a specific conversation

TODO: Add Description
*/
func (a *Client) ConversationsByIDGet(params *ConversationsByIDGetParams) (*ConversationsByIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConversationsByIDGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ConversationsByIdGet",
		Method:             "GET",
		PathPattern:        "/conversations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConversationsByIDGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ConversationsByIDGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ConversationsByIdGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ConversationsGet gets conversation items

TODO: Add Description
*/
func (a *Client) ConversationsGet(params *ConversationsGetParams) (*ConversationsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConversationsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ConversationsGet",
		Method:             "GET",
		PathPattern:        "/conversations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConversationsGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ConversationsGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ConversationsGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}