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

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewStopOrganizationRunsParams creates a new StopOrganizationRunsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStopOrganizationRunsParams() *StopOrganizationRunsParams {
	return &StopOrganizationRunsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStopOrganizationRunsParamsWithTimeout creates a new StopOrganizationRunsParams object
// with the ability to set a timeout on a request.
func NewStopOrganizationRunsParamsWithTimeout(timeout time.Duration) *StopOrganizationRunsParams {
	return &StopOrganizationRunsParams{
		timeout: timeout,
	}
}

// NewStopOrganizationRunsParamsWithContext creates a new StopOrganizationRunsParams object
// with the ability to set a context for a request.
func NewStopOrganizationRunsParamsWithContext(ctx context.Context) *StopOrganizationRunsParams {
	return &StopOrganizationRunsParams{
		Context: ctx,
	}
}

// NewStopOrganizationRunsParamsWithHTTPClient creates a new StopOrganizationRunsParams object
// with the ability to set a custom HTTPClient for a request.
func NewStopOrganizationRunsParamsWithHTTPClient(client *http.Client) *StopOrganizationRunsParams {
	return &StopOrganizationRunsParams{
		HTTPClient: client,
	}
}

/* StopOrganizationRunsParams contains all the parameters to send to the API endpoint
   for the stop organization runs operation.

   Typically these are written to a http.Request.
*/
type StopOrganizationRunsParams struct {

	/* Body.

	   Uuids of the entities
	*/
	Body *service_model.V1Uuids

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stop organization runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopOrganizationRunsParams) WithDefaults() *StopOrganizationRunsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stop organization runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopOrganizationRunsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stop organization runs params
func (o *StopOrganizationRunsParams) WithTimeout(timeout time.Duration) *StopOrganizationRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop organization runs params
func (o *StopOrganizationRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop organization runs params
func (o *StopOrganizationRunsParams) WithContext(ctx context.Context) *StopOrganizationRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop organization runs params
func (o *StopOrganizationRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop organization runs params
func (o *StopOrganizationRunsParams) WithHTTPClient(client *http.Client) *StopOrganizationRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop organization runs params
func (o *StopOrganizationRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stop organization runs params
func (o *StopOrganizationRunsParams) WithBody(body *service_model.V1Uuids) *StopOrganizationRunsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stop organization runs params
func (o *StopOrganizationRunsParams) SetBody(body *service_model.V1Uuids) {
	o.Body = body
}

// WithOwner adds the owner to the stop organization runs params
func (o *StopOrganizationRunsParams) WithOwner(owner string) *StopOrganizationRunsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the stop organization runs params
func (o *StopOrganizationRunsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *StopOrganizationRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
