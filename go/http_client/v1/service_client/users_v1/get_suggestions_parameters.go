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

package users_v1

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

// NewGetSuggestionsParams creates a new GetSuggestionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSuggestionsParams() *GetSuggestionsParams {
	return &GetSuggestionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSuggestionsParamsWithTimeout creates a new GetSuggestionsParams object
// with the ability to set a timeout on a request.
func NewGetSuggestionsParamsWithTimeout(timeout time.Duration) *GetSuggestionsParams {
	return &GetSuggestionsParams{
		timeout: timeout,
	}
}

// NewGetSuggestionsParamsWithContext creates a new GetSuggestionsParams object
// with the ability to set a context for a request.
func NewGetSuggestionsParamsWithContext(ctx context.Context) *GetSuggestionsParams {
	return &GetSuggestionsParams{
		Context: ctx,
	}
}

// NewGetSuggestionsParamsWithHTTPClient creates a new GetSuggestionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSuggestionsParamsWithHTTPClient(client *http.Client) *GetSuggestionsParams {
	return &GetSuggestionsParams{
		HTTPClient: client,
	}
}

/* GetSuggestionsParams contains all the parameters to send to the API endpoint
   for the get suggestions operation.

   Typically these are written to a http.Request.
*/
type GetSuggestionsParams struct {

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

	/* Query.

	   Query filter the search.
	*/
	Query *string

	/* Sort.

	   Sort to order the search.
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get suggestions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSuggestionsParams) WithDefaults() *GetSuggestionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get suggestions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSuggestionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get suggestions params
func (o *GetSuggestionsParams) WithTimeout(timeout time.Duration) *GetSuggestionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get suggestions params
func (o *GetSuggestionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get suggestions params
func (o *GetSuggestionsParams) WithContext(ctx context.Context) *GetSuggestionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get suggestions params
func (o *GetSuggestionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get suggestions params
func (o *GetSuggestionsParams) WithHTTPClient(client *http.Client) *GetSuggestionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get suggestions params
func (o *GetSuggestionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get suggestions params
func (o *GetSuggestionsParams) WithLimit(limit *int32) *GetSuggestionsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get suggestions params
func (o *GetSuggestionsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithNoPage adds the noPage to the get suggestions params
func (o *GetSuggestionsParams) WithNoPage(noPage *bool) *GetSuggestionsParams {
	o.SetNoPage(noPage)
	return o
}

// SetNoPage adds the noPage to the get suggestions params
func (o *GetSuggestionsParams) SetNoPage(noPage *bool) {
	o.NoPage = noPage
}

// WithOffset adds the offset to the get suggestions params
func (o *GetSuggestionsParams) WithOffset(offset *int32) *GetSuggestionsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get suggestions params
func (o *GetSuggestionsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithQuery adds the query to the get suggestions params
func (o *GetSuggestionsParams) WithQuery(query *string) *GetSuggestionsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get suggestions params
func (o *GetSuggestionsParams) SetQuery(query *string) {
	o.Query = query
}

// WithSort adds the sort to the get suggestions params
func (o *GetSuggestionsParams) WithSort(sort *string) *GetSuggestionsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get suggestions params
func (o *GetSuggestionsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetSuggestionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}