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

package connections_v1

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

// NewPatchConnectionParams creates a new PatchConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchConnectionParams() *PatchConnectionParams {
	return &PatchConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchConnectionParamsWithTimeout creates a new PatchConnectionParams object
// with the ability to set a timeout on a request.
func NewPatchConnectionParamsWithTimeout(timeout time.Duration) *PatchConnectionParams {
	return &PatchConnectionParams{
		timeout: timeout,
	}
}

// NewPatchConnectionParamsWithContext creates a new PatchConnectionParams object
// with the ability to set a context for a request.
func NewPatchConnectionParamsWithContext(ctx context.Context) *PatchConnectionParams {
	return &PatchConnectionParams{
		Context: ctx,
	}
}

// NewPatchConnectionParamsWithHTTPClient creates a new PatchConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchConnectionParamsWithHTTPClient(client *http.Client) *PatchConnectionParams {
	return &PatchConnectionParams{
		HTTPClient: client,
	}
}

/* PatchConnectionParams contains all the parameters to send to the API endpoint
   for the patch connection operation.

   Typically these are written to a http.Request.
*/
type PatchConnectionParams struct {

	/* Body.

	   Connection body
	*/
	Body *service_model.V1ConnectionResponse

	/* ConnectionUUID.

	   UUID
	*/
	ConnectionUUID string

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchConnectionParams) WithDefaults() *PatchConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch connection params
func (o *PatchConnectionParams) WithTimeout(timeout time.Duration) *PatchConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch connection params
func (o *PatchConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch connection params
func (o *PatchConnectionParams) WithContext(ctx context.Context) *PatchConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch connection params
func (o *PatchConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch connection params
func (o *PatchConnectionParams) WithHTTPClient(client *http.Client) *PatchConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch connection params
func (o *PatchConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch connection params
func (o *PatchConnectionParams) WithBody(body *service_model.V1ConnectionResponse) *PatchConnectionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch connection params
func (o *PatchConnectionParams) SetBody(body *service_model.V1ConnectionResponse) {
	o.Body = body
}

// WithConnectionUUID adds the connectionUUID to the patch connection params
func (o *PatchConnectionParams) WithConnectionUUID(connectionUUID string) *PatchConnectionParams {
	o.SetConnectionUUID(connectionUUID)
	return o
}

// SetConnectionUUID adds the connectionUuid to the patch connection params
func (o *PatchConnectionParams) SetConnectionUUID(connectionUUID string) {
	o.ConnectionUUID = connectionUUID
}

// WithOwner adds the owner to the patch connection params
func (o *PatchConnectionParams) WithOwner(owner string) *PatchConnectionParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the patch connection params
func (o *PatchConnectionParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *PatchConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param connection.uuid
	if err := r.SetPathParam("connection.uuid", o.ConnectionUUID); err != nil {
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
