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

// NewBookmarkRunParams creates a new BookmarkRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBookmarkRunParams() *BookmarkRunParams {
	return &BookmarkRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBookmarkRunParamsWithTimeout creates a new BookmarkRunParams object
// with the ability to set a timeout on a request.
func NewBookmarkRunParamsWithTimeout(timeout time.Duration) *BookmarkRunParams {
	return &BookmarkRunParams{
		timeout: timeout,
	}
}

// NewBookmarkRunParamsWithContext creates a new BookmarkRunParams object
// with the ability to set a context for a request.
func NewBookmarkRunParamsWithContext(ctx context.Context) *BookmarkRunParams {
	return &BookmarkRunParams{
		Context: ctx,
	}
}

// NewBookmarkRunParamsWithHTTPClient creates a new BookmarkRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewBookmarkRunParamsWithHTTPClient(client *http.Client) *BookmarkRunParams {
	return &BookmarkRunParams{
		HTTPClient: client,
	}
}

/* BookmarkRunParams contains all the parameters to send to the API endpoint
   for the bookmark run operation.

   Typically these are written to a http.Request.
*/
type BookmarkRunParams struct {

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

// WithDefaults hydrates default values in the bookmark run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BookmarkRunParams) WithDefaults() *BookmarkRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bookmark run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BookmarkRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bookmark run params
func (o *BookmarkRunParams) WithTimeout(timeout time.Duration) *BookmarkRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bookmark run params
func (o *BookmarkRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bookmark run params
func (o *BookmarkRunParams) WithContext(ctx context.Context) *BookmarkRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bookmark run params
func (o *BookmarkRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bookmark run params
func (o *BookmarkRunParams) WithHTTPClient(client *http.Client) *BookmarkRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bookmark run params
func (o *BookmarkRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntity adds the entity to the bookmark run params
func (o *BookmarkRunParams) WithEntity(entity string) *BookmarkRunParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the bookmark run params
func (o *BookmarkRunParams) SetEntity(entity string) {
	o.Entity = entity
}

// WithOwner adds the owner to the bookmark run params
func (o *BookmarkRunParams) WithOwner(owner string) *BookmarkRunParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the bookmark run params
func (o *BookmarkRunParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the bookmark run params
func (o *BookmarkRunParams) WithUUID(uuid string) *BookmarkRunParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the bookmark run params
func (o *BookmarkRunParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *BookmarkRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
