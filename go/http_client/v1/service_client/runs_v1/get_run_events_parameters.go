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

// NewGetRunEventsParams creates a new GetRunEventsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRunEventsParams() *GetRunEventsParams {
	return &GetRunEventsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunEventsParamsWithTimeout creates a new GetRunEventsParams object
// with the ability to set a timeout on a request.
func NewGetRunEventsParamsWithTimeout(timeout time.Duration) *GetRunEventsParams {
	return &GetRunEventsParams{
		timeout: timeout,
	}
}

// NewGetRunEventsParamsWithContext creates a new GetRunEventsParams object
// with the ability to set a context for a request.
func NewGetRunEventsParamsWithContext(ctx context.Context) *GetRunEventsParams {
	return &GetRunEventsParams{
		Context: ctx,
	}
}

// NewGetRunEventsParamsWithHTTPClient creates a new GetRunEventsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRunEventsParamsWithHTTPClient(client *http.Client) *GetRunEventsParams {
	return &GetRunEventsParams{
		HTTPClient: client,
	}
}

/* GetRunEventsParams contains all the parameters to send to the API endpoint
   for the get run events operation.

   Typically these are written to a http.Request.
*/
type GetRunEventsParams struct {

	/* Force.

	   Force query param.
	*/
	Force *bool

	/* Kind.

	   The artifact kind
	*/
	Kind string

	/* Names.

	   Names query param.
	*/
	Names *string

	/* Namespace.

	   namespace
	*/
	Namespace string

	/* Orient.

	   Orient query param.
	*/
	Orient *string

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* Project.

	   Project where the run will be assigned
	*/
	Project string

	/* Sample.

	   Sample query param.

	   Format: int32
	*/
	Sample *int32

	/* UUID.

	   Uuid identifier of the entity
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get run events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunEventsParams) WithDefaults() *GetRunEventsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get run events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunEventsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get run events params
func (o *GetRunEventsParams) WithTimeout(timeout time.Duration) *GetRunEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get run events params
func (o *GetRunEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get run events params
func (o *GetRunEventsParams) WithContext(ctx context.Context) *GetRunEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get run events params
func (o *GetRunEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get run events params
func (o *GetRunEventsParams) WithHTTPClient(client *http.Client) *GetRunEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get run events params
func (o *GetRunEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the get run events params
func (o *GetRunEventsParams) WithForce(force *bool) *GetRunEventsParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the get run events params
func (o *GetRunEventsParams) SetForce(force *bool) {
	o.Force = force
}

// WithKind adds the kind to the get run events params
func (o *GetRunEventsParams) WithKind(kind string) *GetRunEventsParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the get run events params
func (o *GetRunEventsParams) SetKind(kind string) {
	o.Kind = kind
}

// WithNames adds the names to the get run events params
func (o *GetRunEventsParams) WithNames(names *string) *GetRunEventsParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the get run events params
func (o *GetRunEventsParams) SetNames(names *string) {
	o.Names = names
}

// WithNamespace adds the namespace to the get run events params
func (o *GetRunEventsParams) WithNamespace(namespace string) *GetRunEventsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get run events params
func (o *GetRunEventsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOrient adds the orient to the get run events params
func (o *GetRunEventsParams) WithOrient(orient *string) *GetRunEventsParams {
	o.SetOrient(orient)
	return o
}

// SetOrient adds the orient to the get run events params
func (o *GetRunEventsParams) SetOrient(orient *string) {
	o.Orient = orient
}

// WithOwner adds the owner to the get run events params
func (o *GetRunEventsParams) WithOwner(owner string) *GetRunEventsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get run events params
func (o *GetRunEventsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the get run events params
func (o *GetRunEventsParams) WithProject(project string) *GetRunEventsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get run events params
func (o *GetRunEventsParams) SetProject(project string) {
	o.Project = project
}

// WithSample adds the sample to the get run events params
func (o *GetRunEventsParams) WithSample(sample *int32) *GetRunEventsParams {
	o.SetSample(sample)
	return o
}

// SetSample adds the sample to the get run events params
func (o *GetRunEventsParams) SetSample(sample *int32) {
	o.Sample = sample
}

// WithUUID adds the uuid to the get run events params
func (o *GetRunEventsParams) WithUUID(uuid string) *GetRunEventsParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get run events params
func (o *GetRunEventsParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param kind
	if err := r.SetPathParam("kind", o.Kind); err != nil {
		return err
	}

	if o.Names != nil {

		// query param names
		var qrNames string

		if o.Names != nil {
			qrNames = *o.Names
		}
		qNames := qrNames
		if qNames != "" {

			if err := r.SetQueryParam("names", qNames); err != nil {
				return err
			}
		}
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Orient != nil {

		// query param orient
		var qrOrient string

		if o.Orient != nil {
			qrOrient = *o.Orient
		}
		qOrient := qrOrient
		if qOrient != "" {

			if err := r.SetQueryParam("orient", qOrient); err != nil {
				return err
			}
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if o.Sample != nil {

		// query param sample
		var qrSample int32

		if o.Sample != nil {
			qrSample = *o.Sample
		}
		qSample := swag.FormatInt32(qrSample)
		if qSample != "" {

			if err := r.SetQueryParam("sample", qSample); err != nil {
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
