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

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// NewSetRunEdgesLineageParams creates a new SetRunEdgesLineageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetRunEdgesLineageParams() *SetRunEdgesLineageParams {
	return &SetRunEdgesLineageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetRunEdgesLineageParamsWithTimeout creates a new SetRunEdgesLineageParams object
// with the ability to set a timeout on a request.
func NewSetRunEdgesLineageParamsWithTimeout(timeout time.Duration) *SetRunEdgesLineageParams {
	return &SetRunEdgesLineageParams{
		timeout: timeout,
	}
}

// NewSetRunEdgesLineageParamsWithContext creates a new SetRunEdgesLineageParams object
// with the ability to set a context for a request.
func NewSetRunEdgesLineageParamsWithContext(ctx context.Context) *SetRunEdgesLineageParams {
	return &SetRunEdgesLineageParams{
		Context: ctx,
	}
}

// NewSetRunEdgesLineageParamsWithHTTPClient creates a new SetRunEdgesLineageParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetRunEdgesLineageParamsWithHTTPClient(client *http.Client) *SetRunEdgesLineageParams {
	return &SetRunEdgesLineageParams{
		HTTPClient: client,
	}
}

/*
SetRunEdgesLineageParams contains all the parameters to send to the API endpoint

	for the set run edges lineage operation.

	Typically these are written to a http.Request.
*/
type SetRunEdgesLineageParams struct {

	/* Body.

	   Run edges graph
	*/
	Body *service_model.V1RunEdgesGraph

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* Project.

	   Project
	*/
	Project string

	/* UUID.

	   Run uuid
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set run edges lineage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetRunEdgesLineageParams) WithDefaults() *SetRunEdgesLineageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set run edges lineage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetRunEdgesLineageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set run edges lineage params
func (o *SetRunEdgesLineageParams) WithTimeout(timeout time.Duration) *SetRunEdgesLineageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set run edges lineage params
func (o *SetRunEdgesLineageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set run edges lineage params
func (o *SetRunEdgesLineageParams) WithContext(ctx context.Context) *SetRunEdgesLineageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set run edges lineage params
func (o *SetRunEdgesLineageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set run edges lineage params
func (o *SetRunEdgesLineageParams) WithHTTPClient(client *http.Client) *SetRunEdgesLineageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set run edges lineage params
func (o *SetRunEdgesLineageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set run edges lineage params
func (o *SetRunEdgesLineageParams) WithBody(body *service_model.V1RunEdgesGraph) *SetRunEdgesLineageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set run edges lineage params
func (o *SetRunEdgesLineageParams) SetBody(body *service_model.V1RunEdgesGraph) {
	o.Body = body
}

// WithOwner adds the owner to the set run edges lineage params
func (o *SetRunEdgesLineageParams) WithOwner(owner string) *SetRunEdgesLineageParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the set run edges lineage params
func (o *SetRunEdgesLineageParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the set run edges lineage params
func (o *SetRunEdgesLineageParams) WithProject(project string) *SetRunEdgesLineageParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the set run edges lineage params
func (o *SetRunEdgesLineageParams) SetProject(project string) {
	o.Project = project
}

// WithUUID adds the uuid to the set run edges lineage params
func (o *SetRunEdgesLineageParams) WithUUID(uuid string) *SetRunEdgesLineageParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the set run edges lineage params
func (o *SetRunEdgesLineageParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SetRunEdgesLineageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
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
