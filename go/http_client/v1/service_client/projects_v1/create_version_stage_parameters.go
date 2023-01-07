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

package projects_v1

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

// NewCreateVersionStageParams creates a new CreateVersionStageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVersionStageParams() *CreateVersionStageParams {
	return &CreateVersionStageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVersionStageParamsWithTimeout creates a new CreateVersionStageParams object
// with the ability to set a timeout on a request.
func NewCreateVersionStageParamsWithTimeout(timeout time.Duration) *CreateVersionStageParams {
	return &CreateVersionStageParams{
		timeout: timeout,
	}
}

// NewCreateVersionStageParamsWithContext creates a new CreateVersionStageParams object
// with the ability to set a context for a request.
func NewCreateVersionStageParamsWithContext(ctx context.Context) *CreateVersionStageParams {
	return &CreateVersionStageParams{
		Context: ctx,
	}
}

// NewCreateVersionStageParamsWithHTTPClient creates a new CreateVersionStageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVersionStageParamsWithHTTPClient(client *http.Client) *CreateVersionStageParams {
	return &CreateVersionStageParams{
		HTTPClient: client,
	}
}

/* CreateVersionStageParams contains all the parameters to send to the API endpoint
   for the create version stage operation.

   Typically these are written to a http.Request.
*/
type CreateVersionStageParams struct {

	// Body.
	Body *service_model.V1EntityStageBodyRequest

	/* Entity.

	   Entity namespace
	*/
	Entity string

	/* Kind.

	   Optional kind, only required for an version entity
	*/
	Kind string

	/* Name.

	   Name of the entity to apply the stage to
	*/
	Name string

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create version stage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVersionStageParams) WithDefaults() *CreateVersionStageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create version stage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVersionStageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create version stage params
func (o *CreateVersionStageParams) WithTimeout(timeout time.Duration) *CreateVersionStageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create version stage params
func (o *CreateVersionStageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create version stage params
func (o *CreateVersionStageParams) WithContext(ctx context.Context) *CreateVersionStageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create version stage params
func (o *CreateVersionStageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create version stage params
func (o *CreateVersionStageParams) WithHTTPClient(client *http.Client) *CreateVersionStageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create version stage params
func (o *CreateVersionStageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create version stage params
func (o *CreateVersionStageParams) WithBody(body *service_model.V1EntityStageBodyRequest) *CreateVersionStageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create version stage params
func (o *CreateVersionStageParams) SetBody(body *service_model.V1EntityStageBodyRequest) {
	o.Body = body
}

// WithEntity adds the entity to the create version stage params
func (o *CreateVersionStageParams) WithEntity(entity string) *CreateVersionStageParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the create version stage params
func (o *CreateVersionStageParams) SetEntity(entity string) {
	o.Entity = entity
}

// WithKind adds the kind to the create version stage params
func (o *CreateVersionStageParams) WithKind(kind string) *CreateVersionStageParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the create version stage params
func (o *CreateVersionStageParams) SetKind(kind string) {
	o.Kind = kind
}

// WithName adds the name to the create version stage params
func (o *CreateVersionStageParams) WithName(name string) *CreateVersionStageParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the create version stage params
func (o *CreateVersionStageParams) SetName(name string) {
	o.Name = name
}

// WithOwner adds the owner to the create version stage params
func (o *CreateVersionStageParams) WithOwner(owner string) *CreateVersionStageParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the create version stage params
func (o *CreateVersionStageParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVersionStageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param entity
	if err := r.SetPathParam("entity", o.Entity); err != nil {
		return err
	}

	// path param kind
	if err := r.SetPathParam("kind", o.Kind); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
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
