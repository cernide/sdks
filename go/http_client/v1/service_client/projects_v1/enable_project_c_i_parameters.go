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
)

// NewEnableProjectCIParams creates a new EnableProjectCIParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnableProjectCIParams() *EnableProjectCIParams {
	return &EnableProjectCIParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnableProjectCIParamsWithTimeout creates a new EnableProjectCIParams object
// with the ability to set a timeout on a request.
func NewEnableProjectCIParamsWithTimeout(timeout time.Duration) *EnableProjectCIParams {
	return &EnableProjectCIParams{
		timeout: timeout,
	}
}

// NewEnableProjectCIParamsWithContext creates a new EnableProjectCIParams object
// with the ability to set a context for a request.
func NewEnableProjectCIParamsWithContext(ctx context.Context) *EnableProjectCIParams {
	return &EnableProjectCIParams{
		Context: ctx,
	}
}

// NewEnableProjectCIParamsWithHTTPClient creates a new EnableProjectCIParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnableProjectCIParamsWithHTTPClient(client *http.Client) *EnableProjectCIParams {
	return &EnableProjectCIParams{
		HTTPClient: client,
	}
}

/* EnableProjectCIParams contains all the parameters to send to the API endpoint
   for the enable project c i operation.

   Typically these are written to a http.Request.
*/
type EnableProjectCIParams struct {

	/* Name.

	   Component under namesapce
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

// WithDefaults hydrates default values in the enable project c i params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableProjectCIParams) WithDefaults() *EnableProjectCIParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enable project c i params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableProjectCIParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enable project c i params
func (o *EnableProjectCIParams) WithTimeout(timeout time.Duration) *EnableProjectCIParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable project c i params
func (o *EnableProjectCIParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable project c i params
func (o *EnableProjectCIParams) WithContext(ctx context.Context) *EnableProjectCIParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable project c i params
func (o *EnableProjectCIParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable project c i params
func (o *EnableProjectCIParams) WithHTTPClient(client *http.Client) *EnableProjectCIParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable project c i params
func (o *EnableProjectCIParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the enable project c i params
func (o *EnableProjectCIParams) WithName(name string) *EnableProjectCIParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the enable project c i params
func (o *EnableProjectCIParams) SetName(name string) {
	o.Name = name
}

// WithOwner adds the owner to the enable project c i params
func (o *EnableProjectCIParams) WithOwner(owner string) *EnableProjectCIParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the enable project c i params
func (o *EnableProjectCIParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *EnableProjectCIParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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