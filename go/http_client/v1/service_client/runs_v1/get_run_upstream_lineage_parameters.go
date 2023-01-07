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
	"github.com/go-openapi/swag"
)

// NewGetRunUpstreamLineageParams creates a new GetRunUpstreamLineageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRunUpstreamLineageParams() *GetRunUpstreamLineageParams {
	return &GetRunUpstreamLineageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunUpstreamLineageParamsWithTimeout creates a new GetRunUpstreamLineageParams object
// with the ability to set a timeout on a request.
func NewGetRunUpstreamLineageParamsWithTimeout(timeout time.Duration) *GetRunUpstreamLineageParams {
	return &GetRunUpstreamLineageParams{
		timeout: timeout,
	}
}

// NewGetRunUpstreamLineageParamsWithContext creates a new GetRunUpstreamLineageParams object
// with the ability to set a context for a request.
func NewGetRunUpstreamLineageParamsWithContext(ctx context.Context) *GetRunUpstreamLineageParams {
	return &GetRunUpstreamLineageParams{
		Context: ctx,
	}
}

// NewGetRunUpstreamLineageParamsWithHTTPClient creates a new GetRunUpstreamLineageParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRunUpstreamLineageParamsWithHTTPClient(client *http.Client) *GetRunUpstreamLineageParams {
	return &GetRunUpstreamLineageParams{
		HTTPClient: client,
	}
}

/* GetRunUpstreamLineageParams contains all the parameters to send to the API endpoint
   for the get run upstream lineage operation.

   Typically these are written to a http.Request.
*/
type GetRunUpstreamLineageParams struct {

	/* Entity.

	   Entity name under namesapce
	*/
	Entity string

	/* Limit.

	   Limit size.

	   Format: int32
	*/
	Limit *int32

	/* NoPage.

	   No pagination.
	*/
	NoPage *bool

	/* Offset.

	   Pagination offset.

	   Format: int32
	*/
	Offset *int32

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* Query.

	   Query filter the search.
	*/
	Query *string

	/* Sort.

	   Sort to order the search.
	*/
	Sort *string

	/* UUID.

	   SubEntity uuid
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get run upstream lineage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunUpstreamLineageParams) WithDefaults() *GetRunUpstreamLineageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get run upstream lineage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunUpstreamLineageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) WithTimeout(timeout time.Duration) *GetRunUpstreamLineageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) WithContext(ctx context.Context) *GetRunUpstreamLineageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) WithHTTPClient(client *http.Client) *GetRunUpstreamLineageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntity adds the entity to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) WithEntity(entity string) *GetRunUpstreamLineageParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) SetEntity(entity string) {
	o.Entity = entity
}

// WithLimit adds the limit to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) WithLimit(limit *int32) *GetRunUpstreamLineageParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithNoPage adds the noPage to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) WithNoPage(noPage *bool) *GetRunUpstreamLineageParams {
	o.SetNoPage(noPage)
	return o
}

// SetNoPage adds the noPage to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) SetNoPage(noPage *bool) {
	o.NoPage = noPage
}

// WithOffset adds the offset to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) WithOffset(offset *int32) *GetRunUpstreamLineageParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOwner adds the owner to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) WithOwner(owner string) *GetRunUpstreamLineageParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithQuery adds the query to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) WithQuery(query *string) *GetRunUpstreamLineageParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) SetQuery(query *string) {
	o.Query = query
}

// WithSort adds the sort to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) WithSort(sort *string) *GetRunUpstreamLineageParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithUUID adds the uuid to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) WithUUID(uuid string) *GetRunUpstreamLineageParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get run upstream lineage params
func (o *GetRunUpstreamLineageParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunUpstreamLineageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entity
	if err := r.SetPathParam("entity", o.Entity); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.NoPage != nil {

		// query param no_page
		var qrNoPage bool

		if o.NoPage != nil {
			qrNoPage = *o.NoPage
		}
		qNoPage := swag.FormatBool(qrNoPage)
		if qNoPage != "" {

			if err := r.SetQueryParam("no_page", qNoPage); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
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
