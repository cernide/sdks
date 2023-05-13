// Code generated by go-swagger; DO NOT EDIT.

package agents_v1

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

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewCreateAgentStatusParams creates a new CreateAgentStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAgentStatusParams() *CreateAgentStatusParams {
	return &CreateAgentStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAgentStatusParamsWithTimeout creates a new CreateAgentStatusParams object
// with the ability to set a timeout on a request.
func NewCreateAgentStatusParamsWithTimeout(timeout time.Duration) *CreateAgentStatusParams {
	return &CreateAgentStatusParams{
		timeout: timeout,
	}
}

// NewCreateAgentStatusParamsWithContext creates a new CreateAgentStatusParams object
// with the ability to set a context for a request.
func NewCreateAgentStatusParamsWithContext(ctx context.Context) *CreateAgentStatusParams {
	return &CreateAgentStatusParams{
		Context: ctx,
	}
}

// NewCreateAgentStatusParamsWithHTTPClient creates a new CreateAgentStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAgentStatusParamsWithHTTPClient(client *http.Client) *CreateAgentStatusParams {
	return &CreateAgentStatusParams{
		HTTPClient: client,
	}
}

/*
CreateAgentStatusParams contains all the parameters to send to the API endpoint

	for the create agent status operation.

	Typically these are written to a http.Request.
*/
type CreateAgentStatusParams struct {

	// Body.
	Body *service_model.V1AgentStatusBodyRequest

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* UUID.

	   Uuid identifier of the entity
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create agent status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAgentStatusParams) WithDefaults() *CreateAgentStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create agent status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAgentStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create agent status params
func (o *CreateAgentStatusParams) WithTimeout(timeout time.Duration) *CreateAgentStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create agent status params
func (o *CreateAgentStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create agent status params
func (o *CreateAgentStatusParams) WithContext(ctx context.Context) *CreateAgentStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create agent status params
func (o *CreateAgentStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create agent status params
func (o *CreateAgentStatusParams) WithHTTPClient(client *http.Client) *CreateAgentStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create agent status params
func (o *CreateAgentStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create agent status params
func (o *CreateAgentStatusParams) WithBody(body *service_model.V1AgentStatusBodyRequest) *CreateAgentStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create agent status params
func (o *CreateAgentStatusParams) SetBody(body *service_model.V1AgentStatusBodyRequest) {
	o.Body = body
}

// WithOwner adds the owner to the create agent status params
func (o *CreateAgentStatusParams) WithOwner(owner string) *CreateAgentStatusParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the create agent status params
func (o *CreateAgentStatusParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the create agent status params
func (o *CreateAgentStatusParams) WithUUID(uuid string) *CreateAgentStatusParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the create agent status params
func (o *CreateAgentStatusParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAgentStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
