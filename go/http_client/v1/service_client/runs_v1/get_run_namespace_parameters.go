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

package runs_v1

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

// NewGetRunNamespaceParams creates a new GetRunNamespaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRunNamespaceParams() *GetRunNamespaceParams {
	return &GetRunNamespaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunNamespaceParamsWithTimeout creates a new GetRunNamespaceParams object
// with the ability to set a timeout on a request.
func NewGetRunNamespaceParamsWithTimeout(timeout time.Duration) *GetRunNamespaceParams {
	return &GetRunNamespaceParams{
		timeout: timeout,
	}
}

// NewGetRunNamespaceParamsWithContext creates a new GetRunNamespaceParams object
// with the ability to set a context for a request.
func NewGetRunNamespaceParamsWithContext(ctx context.Context) *GetRunNamespaceParams {
	return &GetRunNamespaceParams{
		Context: ctx,
	}
}

// NewGetRunNamespaceParamsWithHTTPClient creates a new GetRunNamespaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRunNamespaceParamsWithHTTPClient(client *http.Client) *GetRunNamespaceParams {
	return &GetRunNamespaceParams{
		HTTPClient: client,
	}
}

/* GetRunNamespaceParams contains all the parameters to send to the API endpoint
   for the get run namespace operation.

   Typically these are written to a http.Request.
*/
type GetRunNamespaceParams struct {

	/* Entity.

	   Entity: project name, hub name, registry name, ...
	*/
	Entity string

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

// WithDefaults hydrates default values in the get run namespace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunNamespaceParams) WithDefaults() *GetRunNamespaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get run namespace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunNamespaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get run namespace params
func (o *GetRunNamespaceParams) WithTimeout(timeout time.Duration) *GetRunNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get run namespace params
func (o *GetRunNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get run namespace params
func (o *GetRunNamespaceParams) WithContext(ctx context.Context) *GetRunNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get run namespace params
func (o *GetRunNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get run namespace params
func (o *GetRunNamespaceParams) WithHTTPClient(client *http.Client) *GetRunNamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get run namespace params
func (o *GetRunNamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntity adds the entity to the get run namespace params
func (o *GetRunNamespaceParams) WithEntity(entity string) *GetRunNamespaceParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the get run namespace params
func (o *GetRunNamespaceParams) SetEntity(entity string) {
	o.Entity = entity
}

// WithOwner adds the owner to the get run namespace params
func (o *GetRunNamespaceParams) WithOwner(owner string) *GetRunNamespaceParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get run namespace params
func (o *GetRunNamespaceParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the get run namespace params
func (o *GetRunNamespaceParams) WithUUID(uuid string) *GetRunNamespaceParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get run namespace params
func (o *GetRunNamespaceParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entity
	if err := r.SetPathParam("entity", o.Entity); err != nil {
		return err
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
