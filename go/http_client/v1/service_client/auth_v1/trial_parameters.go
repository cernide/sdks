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

package auth_v1

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

// NewTrialParams creates a new TrialParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTrialParams() *TrialParams {
	return &TrialParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTrialParamsWithTimeout creates a new TrialParams object
// with the ability to set a timeout on a request.
func NewTrialParamsWithTimeout(timeout time.Duration) *TrialParams {
	return &TrialParams{
		timeout: timeout,
	}
}

// NewTrialParamsWithContext creates a new TrialParams object
// with the ability to set a context for a request.
func NewTrialParamsWithContext(ctx context.Context) *TrialParams {
	return &TrialParams{
		Context: ctx,
	}
}

// NewTrialParamsWithHTTPClient creates a new TrialParams object
// with the ability to set a custom HTTPClient for a request.
func NewTrialParamsWithHTTPClient(client *http.Client) *TrialParams {
	return &TrialParams{
		HTTPClient: client,
	}
}

/* TrialParams contains all the parameters to send to the API endpoint
   for the trial operation.

   Typically these are written to a http.Request.
*/
type TrialParams struct {

	// Body.
	Body *service_model.V1TrialStart

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the trial params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TrialParams) WithDefaults() *TrialParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the trial params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TrialParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the trial params
func (o *TrialParams) WithTimeout(timeout time.Duration) *TrialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the trial params
func (o *TrialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the trial params
func (o *TrialParams) WithContext(ctx context.Context) *TrialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the trial params
func (o *TrialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the trial params
func (o *TrialParams) WithHTTPClient(client *http.Client) *TrialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the trial params
func (o *TrialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the trial params
func (o *TrialParams) WithBody(body *service_model.V1TrialStart) *TrialParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the trial params
func (o *TrialParams) SetBody(body *service_model.V1TrialStart) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TrialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
