// Code generated by go-swagger; DO NOT EDIT.

package customercontacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new customercontacts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for customercontacts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateanewCustomerContact(params *CreateanewCustomerContactParams) (*CreateanewCustomerContactCreated, error)

	GetCustomerContactitems(params *GetCustomerContactitemsParams) (*GetCustomerContactitemsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateanewCustomerContact creates a new customer contact
*/
func (a *Client) CreateanewCustomerContact(params *CreateanewCustomerContactParams) (*CreateanewCustomerContactCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateanewCustomerContactParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateanewCustomerContact",
		Method:             "POST",
		PathPattern:        "/customercontacts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateanewCustomerContactReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateanewCustomerContactCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateanewCustomerContact: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCustomerContactitems gets customer contact items
*/
func (a *Client) GetCustomerContactitems(params *GetCustomerContactitemsParams) (*GetCustomerContactitemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomerContactitemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCustomerContactitems",
		Method:             "GET",
		PathPattern:        "/customercontacts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCustomerContactitemsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCustomerContactitemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCustomerContactitems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}