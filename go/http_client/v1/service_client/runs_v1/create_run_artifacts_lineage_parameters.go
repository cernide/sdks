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

// NewCreateRunArtifactsLineageParams creates a new CreateRunArtifactsLineageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRunArtifactsLineageParams() *CreateRunArtifactsLineageParams {
	return &CreateRunArtifactsLineageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRunArtifactsLineageParamsWithTimeout creates a new CreateRunArtifactsLineageParams object
// with the ability to set a timeout on a request.
func NewCreateRunArtifactsLineageParamsWithTimeout(timeout time.Duration) *CreateRunArtifactsLineageParams {
	return &CreateRunArtifactsLineageParams{
		timeout: timeout,
	}
}

// NewCreateRunArtifactsLineageParamsWithContext creates a new CreateRunArtifactsLineageParams object
// with the ability to set a context for a request.
func NewCreateRunArtifactsLineageParamsWithContext(ctx context.Context) *CreateRunArtifactsLineageParams {
	return &CreateRunArtifactsLineageParams{
		Context: ctx,
	}
}

// NewCreateRunArtifactsLineageParamsWithHTTPClient creates a new CreateRunArtifactsLineageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRunArtifactsLineageParamsWithHTTPClient(client *http.Client) *CreateRunArtifactsLineageParams {
	return &CreateRunArtifactsLineageParams{
		HTTPClient: client,
	}
}

/*
CreateRunArtifactsLineageParams contains all the parameters to send to the API endpoint

	for the create run artifacts lineage operation.

	Typically these are written to a http.Request.
*/
type CreateRunArtifactsLineageParams struct {

	/* Body.

	   Run Artifacts
	*/
	Body *service_model.V1RunArtifacts

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* Project.

	   Project where the run will be assigned
	*/
	Project string

	/* UUID.

	   Uuid identifier of the entity
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create run artifacts lineage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRunArtifactsLineageParams) WithDefaults() *CreateRunArtifactsLineageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create run artifacts lineage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRunArtifactsLineageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create run artifacts lineage params
func (o *CreateRunArtifactsLineageParams) WithTimeout(timeout time.Duration) *CreateRunArtifactsLineageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create run artifacts lineage params
func (o *CreateRunArtifactsLineageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create run artifacts lineage params
func (o *CreateRunArtifactsLineageParams) WithContext(ctx context.Context) *CreateRunArtifactsLineageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create run artifacts lineage params
func (o *CreateRunArtifactsLineageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create run artifacts lineage params
func (o *CreateRunArtifactsLineageParams) WithHTTPClient(client *http.Client) *CreateRunArtifactsLineageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create run artifacts lineage params
func (o *CreateRunArtifactsLineageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create run artifacts lineage params
func (o *CreateRunArtifactsLineageParams) WithBody(body *service_model.V1RunArtifacts) *CreateRunArtifactsLineageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create run artifacts lineage params
func (o *CreateRunArtifactsLineageParams) SetBody(body *service_model.V1RunArtifacts) {
	o.Body = body
}

// WithOwner adds the owner to the create run artifacts lineage params
func (o *CreateRunArtifactsLineageParams) WithOwner(owner string) *CreateRunArtifactsLineageParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the create run artifacts lineage params
func (o *CreateRunArtifactsLineageParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the create run artifacts lineage params
func (o *CreateRunArtifactsLineageParams) WithProject(project string) *CreateRunArtifactsLineageParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the create run artifacts lineage params
func (o *CreateRunArtifactsLineageParams) SetProject(project string) {
	o.Project = project
}

// WithUUID adds the uuid to the create run artifacts lineage params
func (o *CreateRunArtifactsLineageParams) WithUUID(uuid string) *CreateRunArtifactsLineageParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the create run artifacts lineage params
func (o *CreateRunArtifactsLineageParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRunArtifactsLineageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
