// Code generated by go-swagger; DO NOT EDIT.

package emergencymappings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new emergencymappings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for emergencymappings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateanewEmergencyMapping(params *CreateanewEmergencyMappingParams) (*CreateanewEmergencyMappingCreated, error)

	GetEmergencyMappingitems(params *GetEmergencyMappingitemsParams) (*GetEmergencyMappingitemsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateanewEmergencyMapping creates a new emergency mapping
*/
func (a *Client) CreateanewEmergencyMapping(params *CreateanewEmergencyMappingParams) (*CreateanewEmergencyMappingCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateanewEmergencyMappingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateanewEmergencyMapping",
		Method:             "POST",
		PathPattern:        "/emergencymappings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateanewEmergencyMappingReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateanewEmergencyMappingCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateanewEmergencyMapping: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEmergencyMappingitems gets emergency mapping items
*/
func (a *Client) GetEmergencyMappingitems(params *GetEmergencyMappingitemsParams) (*GetEmergencyMappingitemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEmergencyMappingitemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEmergencyMappingitems",
		Method:             "GET",
		PathPattern:        "/emergencymappings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEmergencyMappingitemsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEmergencyMappingitemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEmergencyMappingitems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}