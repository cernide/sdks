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

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// NewUpdateAgentParams creates a new UpdateAgentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAgentParams() *UpdateAgentParams {
	return &UpdateAgentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAgentParamsWithTimeout creates a new UpdateAgentParams object
// with the ability to set a timeout on a request.
func NewUpdateAgentParamsWithTimeout(timeout time.Duration) *UpdateAgentParams {
	return &UpdateAgentParams{
		timeout: timeout,
	}
}

// NewUpdateAgentParamsWithContext creates a new UpdateAgentParams object
// with the ability to set a context for a request.
func NewUpdateAgentParamsWithContext(ctx context.Context) *UpdateAgentParams {
	return &UpdateAgentParams{
		Context: ctx,
	}
}

// NewUpdateAgentParamsWithHTTPClient creates a new UpdateAgentParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAgentParamsWithHTTPClient(client *http.Client) *UpdateAgentParams {
	return &UpdateAgentParams{
		HTTPClient: client,
	}
}

/*
UpdateAgentParams contains all the parameters to send to the API endpoint

	for the update agent operation.

	Typically these are written to a http.Request.
*/
type UpdateAgentParams struct {

	/* AgentUUID.

	   UUID
	*/
	AgentUUID string

	/* Body.

	   Agent body
	*/
	Body *service_model.V1Agent

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAgentParams) WithDefaults() *UpdateAgentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAgentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update agent params
func (o *UpdateAgentParams) WithTimeout(timeout time.Duration) *UpdateAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update agent params
func (o *UpdateAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update agent params
func (o *UpdateAgentParams) WithContext(ctx context.Context) *UpdateAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update agent params
func (o *UpdateAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update agent params
func (o *UpdateAgentParams) WithHTTPClient(client *http.Client) *UpdateAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update agent params
func (o *UpdateAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentUUID adds the agentUUID to the update agent params
func (o *UpdateAgentParams) WithAgentUUID(agentUUID string) *UpdateAgentParams {
	o.SetAgentUUID(agentUUID)
	return o
}

// SetAgentUUID adds the agentUuid to the update agent params
func (o *UpdateAgentParams) SetAgentUUID(agentUUID string) {
	o.AgentUUID = agentUUID
}

// WithBody adds the body to the update agent params
func (o *UpdateAgentParams) WithBody(body *service_model.V1Agent) *UpdateAgentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update agent params
func (o *UpdateAgentParams) SetBody(body *service_model.V1Agent) {
	o.Body = body
}

// WithOwner adds the owner to the update agent params
func (o *UpdateAgentParams) WithOwner(owner string) *UpdateAgentParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the update agent params
func (o *UpdateAgentParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agent.uuid
	if err := r.SetPathParam("agent.uuid", o.AgentUUID); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
