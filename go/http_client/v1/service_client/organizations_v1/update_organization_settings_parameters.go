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

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewUpdateOrganizationSettingsParams creates a new UpdateOrganizationSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOrganizationSettingsParams() *UpdateOrganizationSettingsParams {
	return &UpdateOrganizationSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrganizationSettingsParamsWithTimeout creates a new UpdateOrganizationSettingsParams object
// with the ability to set a timeout on a request.
func NewUpdateOrganizationSettingsParamsWithTimeout(timeout time.Duration) *UpdateOrganizationSettingsParams {
	return &UpdateOrganizationSettingsParams{
		timeout: timeout,
	}
}

// NewUpdateOrganizationSettingsParamsWithContext creates a new UpdateOrganizationSettingsParams object
// with the ability to set a context for a request.
func NewUpdateOrganizationSettingsParamsWithContext(ctx context.Context) *UpdateOrganizationSettingsParams {
	return &UpdateOrganizationSettingsParams{
		Context: ctx,
	}
}

// NewUpdateOrganizationSettingsParamsWithHTTPClient creates a new UpdateOrganizationSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOrganizationSettingsParamsWithHTTPClient(client *http.Client) *UpdateOrganizationSettingsParams {
	return &UpdateOrganizationSettingsParams{
		HTTPClient: client,
	}
}

/* UpdateOrganizationSettingsParams contains all the parameters to send to the API endpoint
   for the update organization settings operation.

   Typically these are written to a http.Request.
*/
type UpdateOrganizationSettingsParams struct {

	/* Body.

	   Organization body
	*/
	Body *service_model.V1Organization

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update organization settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationSettingsParams) WithDefaults() *UpdateOrganizationSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update organization settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update organization settings params
func (o *UpdateOrganizationSettingsParams) WithTimeout(timeout time.Duration) *UpdateOrganizationSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update organization settings params
func (o *UpdateOrganizationSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update organization settings params
func (o *UpdateOrganizationSettingsParams) WithContext(ctx context.Context) *UpdateOrganizationSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update organization settings params
func (o *UpdateOrganizationSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update organization settings params
func (o *UpdateOrganizationSettingsParams) WithHTTPClient(client *http.Client) *UpdateOrganizationSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update organization settings params
func (o *UpdateOrganizationSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update organization settings params
func (o *UpdateOrganizationSettingsParams) WithBody(body *service_model.V1Organization) *UpdateOrganizationSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update organization settings params
func (o *UpdateOrganizationSettingsParams) SetBody(body *service_model.V1Organization) {
	o.Body = body
}

// WithOwner adds the owner to the update organization settings params
func (o *UpdateOrganizationSettingsParams) WithOwner(owner string) *UpdateOrganizationSettingsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the update organization settings params
func (o *UpdateOrganizationSettingsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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