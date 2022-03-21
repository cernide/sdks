// Copyright 2018-2021 Polyaxon, Inc.
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

package organizations_v1

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

// NewGetOrganizationRunParams creates a new GetOrganizationRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationRunParams() *GetOrganizationRunParams {
	return &GetOrganizationRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationRunParamsWithTimeout creates a new GetOrganizationRunParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationRunParamsWithTimeout(timeout time.Duration) *GetOrganizationRunParams {
	return &GetOrganizationRunParams{
		timeout: timeout,
	}
}

// NewGetOrganizationRunParamsWithContext creates a new GetOrganizationRunParams object
// with the ability to set a context for a request.
func NewGetOrganizationRunParamsWithContext(ctx context.Context) *GetOrganizationRunParams {
	return &GetOrganizationRunParams{
		Context: ctx,
	}
}

// NewGetOrganizationRunParamsWithHTTPClient creates a new GetOrganizationRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationRunParamsWithHTTPClient(client *http.Client) *GetOrganizationRunParams {
	return &GetOrganizationRunParams{
		HTTPClient: client,
	}
}

/* GetOrganizationRunParams contains all the parameters to send to the API endpoint
   for the get organization run operation.

   Typically these are written to a http.Request.
*/
type GetOrganizationRunParams struct {

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

// WithDefaults hydrates default values in the get organization run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationRunParams) WithDefaults() *GetOrganizationRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization run params
func (o *GetOrganizationRunParams) WithTimeout(timeout time.Duration) *GetOrganizationRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization run params
func (o *GetOrganizationRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization run params
func (o *GetOrganizationRunParams) WithContext(ctx context.Context) *GetOrganizationRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization run params
func (o *GetOrganizationRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization run params
func (o *GetOrganizationRunParams) WithHTTPClient(client *http.Client) *GetOrganizationRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization run params
func (o *GetOrganizationRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the get organization run params
func (o *GetOrganizationRunParams) WithOwner(owner string) *GetOrganizationRunParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get organization run params
func (o *GetOrganizationRunParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the get organization run params
func (o *GetOrganizationRunParams) WithUUID(uuid string) *GetOrganizationRunParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get organization run params
func (o *GetOrganizationRunParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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