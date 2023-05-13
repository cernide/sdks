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
)

// NewGetAgentConfigParams creates a new GetAgentConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAgentConfigParams() *GetAgentConfigParams {
	return &GetAgentConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAgentConfigParamsWithTimeout creates a new GetAgentConfigParams object
// with the ability to set a timeout on a request.
func NewGetAgentConfigParamsWithTimeout(timeout time.Duration) *GetAgentConfigParams {
	return &GetAgentConfigParams{
		timeout: timeout,
	}
}

// NewGetAgentConfigParamsWithContext creates a new GetAgentConfigParams object
// with the ability to set a context for a request.
func NewGetAgentConfigParamsWithContext(ctx context.Context) *GetAgentConfigParams {
	return &GetAgentConfigParams{
		Context: ctx,
	}
}

// NewGetAgentConfigParamsWithHTTPClient creates a new GetAgentConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAgentConfigParamsWithHTTPClient(client *http.Client) *GetAgentConfigParams {
	return &GetAgentConfigParams{
		HTTPClient: client,
	}
}

/*
GetAgentConfigParams contains all the parameters to send to the API endpoint

	for the get agent config operation.

	Typically these are written to a http.Request.
*/
type GetAgentConfigParams struct {

	/* Entity.

	   Entity: project name, hub name, registry name, ...
	*/
	Entity *string

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* UUID.

	   Uuid identifier of the sub-entity
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get agent config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAgentConfigParams) WithDefaults() *GetAgentConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get agent config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAgentConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get agent config params
func (o *GetAgentConfigParams) WithTimeout(timeout time.Duration) *GetAgentConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get agent config params
func (o *GetAgentConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get agent config params
func (o *GetAgentConfigParams) WithContext(ctx context.Context) *GetAgentConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get agent config params
func (o *GetAgentConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get agent config params
func (o *GetAgentConfigParams) WithHTTPClient(client *http.Client) *GetAgentConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get agent config params
func (o *GetAgentConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntity adds the entity to the get agent config params
func (o *GetAgentConfigParams) WithEntity(entity *string) *GetAgentConfigParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the get agent config params
func (o *GetAgentConfigParams) SetEntity(entity *string) {
	o.Entity = entity
}

// WithOwner adds the owner to the get agent config params
func (o *GetAgentConfigParams) WithOwner(owner string) *GetAgentConfigParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get agent config params
func (o *GetAgentConfigParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the get agent config params
func (o *GetAgentConfigParams) WithUUID(uuid string) *GetAgentConfigParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get agent config params
func (o *GetAgentConfigParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetAgentConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Entity != nil {

		// query param entity
		var qrEntity string

		if o.Entity != nil {
			qrEntity = *o.Entity
		}
		qEntity := qrEntity
		if qEntity != "" {

			if err := r.SetQueryParam("entity", qEntity); err != nil {
				return err
			}
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
