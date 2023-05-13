// Code generated by go-swagger; DO NOT EDIT.

package connections_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new connections v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for connections v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateConnection(params *CreateConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateConnectionOK, *CreateConnectionNoContent, error)

	DeleteConnection(params *DeleteConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteConnectionOK, *DeleteConnectionNoContent, error)

	GetConnection(params *GetConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetConnectionOK, *GetConnectionNoContent, error)

	ListConnectionNames(params *ListConnectionNamesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListConnectionNamesOK, *ListConnectionNamesNoContent, error)

	ListConnections(params *ListConnectionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListConnectionsOK, *ListConnectionsNoContent, error)

	PatchConnection(params *PatchConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchConnectionOK, *PatchConnectionNoContent, error)

	UpdateConnection(params *UpdateConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateConnectionOK, *UpdateConnectionNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateConnection creates connection
*/
func (a *Client) CreateConnection(params *CreateConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateConnectionOK, *CreateConnectionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateConnection",
		Method:             "POST",
		PathPattern:        "/api/v1/orgs/{owner}/connections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateConnectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateConnectionOK:
		return value, nil, nil
	case *CreateConnectionNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateConnectionDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteConnection deletes connection
*/
func (a *Client) DeleteConnection(params *DeleteConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteConnectionOK, *DeleteConnectionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteConnection",
		Method:             "DELETE",
		PathPattern:        "/api/v1/orgs/{owner}/connections/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteConnectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteConnectionOK:
		return value, nil, nil
	case *DeleteConnectionNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteConnectionDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetConnection gets connection
*/
func (a *Client) GetConnection(params *GetConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetConnectionOK, *GetConnectionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetConnection",
		Method:             "GET",
		PathPattern:        "/api/v1/orgs/{owner}/connections/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetConnectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetConnectionOK:
		return value, nil, nil
	case *GetConnectionNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetConnectionDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListConnectionNames lists connections names
*/
func (a *Client) ListConnectionNames(params *ListConnectionNamesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListConnectionNamesOK, *ListConnectionNamesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListConnectionNamesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListConnectionNames",
		Method:             "GET",
		PathPattern:        "/api/v1/orgs/{owner}/connections/names",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListConnectionNamesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ListConnectionNamesOK:
		return value, nil, nil
	case *ListConnectionNamesNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListConnectionNamesDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListConnections lists connections
*/
func (a *Client) ListConnections(params *ListConnectionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListConnectionsOK, *ListConnectionsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListConnectionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListConnections",
		Method:             "GET",
		PathPattern:        "/api/v1/orgs/{owner}/connections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListConnectionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ListConnectionsOK:
		return value, nil, nil
	case *ListConnectionsNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListConnectionsDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchConnection patches connection
*/
func (a *Client) PatchConnection(params *PatchConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchConnectionOK, *PatchConnectionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchConnection",
		Method:             "PATCH",
		PathPattern:        "/api/v1/orgs/{owner}/connections/{connection.uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchConnectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PatchConnectionOK:
		return value, nil, nil
	case *PatchConnectionNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchConnectionDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateConnection updates connection
*/
func (a *Client) UpdateConnection(params *UpdateConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateConnectionOK, *UpdateConnectionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateConnection",
		Method:             "PUT",
		PathPattern:        "/api/v1/orgs/{owner}/connections/{connection.uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateConnectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateConnectionOK:
		return value, nil, nil
	case *UpdateConnectionNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateConnectionDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
