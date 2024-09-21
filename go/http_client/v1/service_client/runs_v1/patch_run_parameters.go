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

// NewPatchRunParams creates a new PatchRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchRunParams() *PatchRunParams {
	return &PatchRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchRunParamsWithTimeout creates a new PatchRunParams object
// with the ability to set a timeout on a request.
func NewPatchRunParamsWithTimeout(timeout time.Duration) *PatchRunParams {
	return &PatchRunParams{
		timeout: timeout,
	}
}

// NewPatchRunParamsWithContext creates a new PatchRunParams object
// with the ability to set a context for a request.
func NewPatchRunParamsWithContext(ctx context.Context) *PatchRunParams {
	return &PatchRunParams{
		Context: ctx,
	}
}

// NewPatchRunParamsWithHTTPClient creates a new PatchRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchRunParamsWithHTTPClient(client *http.Client) *PatchRunParams {
	return &PatchRunParams{
		HTTPClient: client,
	}
}

/*
PatchRunParams contains all the parameters to send to the API endpoint

	for the patch run operation.

	Typically these are written to a http.Request.
*/
type PatchRunParams struct {

	/* Body.

	   Run object
	*/
	Body *service_model.V1Run

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* Project.

	   Project where the run will be assigned
	*/
	Project string

	/* RunUUID.

	   UUID
	*/
	RunUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchRunParams) WithDefaults() *PatchRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch run params
func (o *PatchRunParams) WithTimeout(timeout time.Duration) *PatchRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch run params
func (o *PatchRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch run params
func (o *PatchRunParams) WithContext(ctx context.Context) *PatchRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch run params
func (o *PatchRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch run params
func (o *PatchRunParams) WithHTTPClient(client *http.Client) *PatchRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch run params
func (o *PatchRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch run params
func (o *PatchRunParams) WithBody(body *service_model.V1Run) *PatchRunParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch run params
func (o *PatchRunParams) SetBody(body *service_model.V1Run) {
	o.Body = body
}

// WithOwner adds the owner to the patch run params
func (o *PatchRunParams) WithOwner(owner string) *PatchRunParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the patch run params
func (o *PatchRunParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the patch run params
func (o *PatchRunParams) WithProject(project string) *PatchRunParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the patch run params
func (o *PatchRunParams) SetProject(project string) {
	o.Project = project
}

// WithRunUUID adds the runUUID to the patch run params
func (o *PatchRunParams) WithRunUUID(runUUID string) *PatchRunParams {
	o.SetRunUUID(runUUID)
	return o
}

// SetRunUUID adds the runUuid to the patch run params
func (o *PatchRunParams) SetRunUUID(runUUID string) {
	o.RunUUID = runUUID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param run.uuid
	if err := r.SetPathParam("run.uuid", o.RunUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
