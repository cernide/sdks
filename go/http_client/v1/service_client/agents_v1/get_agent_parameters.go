// Copyright 2018-2023 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

// NewGetAgentParams creates a new GetAgentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAgentParams() *GetAgentParams {
	return &GetAgentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAgentParamsWithTimeout creates a new GetAgentParams object
// with the ability to set a timeout on a request.
func NewGetAgentParamsWithTimeout(timeout time.Duration) *GetAgentParams {
	return &GetAgentParams{
		timeout: timeout,
	}
}

// NewGetAgentParamsWithContext creates a new GetAgentParams object
// with the ability to set a context for a request.
func NewGetAgentParamsWithContext(ctx context.Context) *GetAgentParams {
	return &GetAgentParams{
		Context: ctx,
	}
}

// NewGetAgentParamsWithHTTPClient creates a new GetAgentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAgentParamsWithHTTPClient(client *http.Client) *GetAgentParams {
	return &GetAgentParams{
		HTTPClient: client,
	}
}

/* GetAgentParams contains all the parameters to send to the API endpoint
   for the get agent operation.

   Typically these are written to a http.Request.
*/
type GetAgentParams struct {

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

// WithDefaults hydrates default values in the get agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAgentParams) WithDefaults() *GetAgentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAgentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get agent params
func (o *GetAgentParams) WithTimeout(timeout time.Duration) *GetAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get agent params
func (o *GetAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get agent params
func (o *GetAgentParams) WithContext(ctx context.Context) *GetAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get agent params
func (o *GetAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get agent params
func (o *GetAgentParams) WithHTTPClient(client *http.Client) *GetAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get agent params
func (o *GetAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntity adds the entity to the get agent params
func (o *GetAgentParams) WithEntity(entity *string) *GetAgentParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the get agent params
func (o *GetAgentParams) SetEntity(entity *string) {
	o.Entity = entity
}

// WithOwner adds the owner to the get agent params
func (o *GetAgentParams) WithOwner(owner string) *GetAgentParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get agent params
func (o *GetAgentParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the get agent params
func (o *GetAgentParams) WithUUID(uuid string) *GetAgentParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get agent params
func (o *GetAgentParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
