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
	"github.com/go-openapi/swag"
)

// NewGetRunArtifactParams creates a new GetRunArtifactParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRunArtifactParams() *GetRunArtifactParams {
	return &GetRunArtifactParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunArtifactParamsWithTimeout creates a new GetRunArtifactParams object
// with the ability to set a timeout on a request.
func NewGetRunArtifactParamsWithTimeout(timeout time.Duration) *GetRunArtifactParams {
	return &GetRunArtifactParams{
		timeout: timeout,
	}
}

// NewGetRunArtifactParamsWithContext creates a new GetRunArtifactParams object
// with the ability to set a context for a request.
func NewGetRunArtifactParamsWithContext(ctx context.Context) *GetRunArtifactParams {
	return &GetRunArtifactParams{
		Context: ctx,
	}
}

// NewGetRunArtifactParamsWithHTTPClient creates a new GetRunArtifactParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRunArtifactParamsWithHTTPClient(client *http.Client) *GetRunArtifactParams {
	return &GetRunArtifactParams{
		HTTPClient: client,
	}
}

/* GetRunArtifactParams contains all the parameters to send to the API endpoint
   for the get run artifact operation.

   Typically these are written to a http.Request.
*/
type GetRunArtifactParams struct {

	/* Force.

	   Whether to force reload.
	*/
	Force *bool

	/* Namespace.

	   namespace
	*/
	Namespace string

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* Path.

	   Artifact filepath.
	*/
	Path *string

	/* Project.

	   Project where the experiement will be assigned
	*/
	Project string

	/* Stream.

	   Whether to stream the file.
	*/
	Stream *bool

	/* UUID.

	   Unique integer identifier of the entity
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get run artifact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunArtifactParams) WithDefaults() *GetRunArtifactParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get run artifact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunArtifactParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get run artifact params
func (o *GetRunArtifactParams) WithTimeout(timeout time.Duration) *GetRunArtifactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get run artifact params
func (o *GetRunArtifactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get run artifact params
func (o *GetRunArtifactParams) WithContext(ctx context.Context) *GetRunArtifactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get run artifact params
func (o *GetRunArtifactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get run artifact params
func (o *GetRunArtifactParams) WithHTTPClient(client *http.Client) *GetRunArtifactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get run artifact params
func (o *GetRunArtifactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the get run artifact params
func (o *GetRunArtifactParams) WithForce(force *bool) *GetRunArtifactParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the get run artifact params
func (o *GetRunArtifactParams) SetForce(force *bool) {
	o.Force = force
}

// WithNamespace adds the namespace to the get run artifact params
func (o *GetRunArtifactParams) WithNamespace(namespace string) *GetRunArtifactParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get run artifact params
func (o *GetRunArtifactParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOwner adds the owner to the get run artifact params
func (o *GetRunArtifactParams) WithOwner(owner string) *GetRunArtifactParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get run artifact params
func (o *GetRunArtifactParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithPath adds the path to the get run artifact params
func (o *GetRunArtifactParams) WithPath(path *string) *GetRunArtifactParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the get run artifact params
func (o *GetRunArtifactParams) SetPath(path *string) {
	o.Path = path
}

// WithProject adds the project to the get run artifact params
func (o *GetRunArtifactParams) WithProject(project string) *GetRunArtifactParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get run artifact params
func (o *GetRunArtifactParams) SetProject(project string) {
	o.Project = project
}

// WithStream adds the stream to the get run artifact params
func (o *GetRunArtifactParams) WithStream(stream *bool) *GetRunArtifactParams {
	o.SetStream(stream)
	return o
}

// SetStream adds the stream to the get run artifact params
func (o *GetRunArtifactParams) SetStream(stream *bool) {
	o.Stream = stream
}

// WithUUID adds the uuid to the get run artifact params
func (o *GetRunArtifactParams) WithUUID(uuid string) *GetRunArtifactParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get run artifact params
func (o *GetRunArtifactParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunArtifactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Force != nil {

		// query param force
		var qrForce bool

		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if o.Path != nil {

		// query param path
		var qrPath string

		if o.Path != nil {
			qrPath = *o.Path
		}
		qPath := qrPath
		if qPath != "" {

			if err := r.SetQueryParam("path", qPath); err != nil {
				return err
			}
		}
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if o.Stream != nil {

		// query param stream
		var qrStream bool

		if o.Stream != nil {
			qrStream = *o.Stream
		}
		qStream := swag.FormatBool(qrStream)
		if qStream != "" {

			if err := r.SetQueryParam("stream", qStream); err != nil {
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