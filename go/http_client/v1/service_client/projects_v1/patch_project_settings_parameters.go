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

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// NewPatchProjectSettingsParams creates a new PatchProjectSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchProjectSettingsParams() *PatchProjectSettingsParams {
	return &PatchProjectSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchProjectSettingsParamsWithTimeout creates a new PatchProjectSettingsParams object
// with the ability to set a timeout on a request.
func NewPatchProjectSettingsParamsWithTimeout(timeout time.Duration) *PatchProjectSettingsParams {
	return &PatchProjectSettingsParams{
		timeout: timeout,
	}
}

// NewPatchProjectSettingsParamsWithContext creates a new PatchProjectSettingsParams object
// with the ability to set a context for a request.
func NewPatchProjectSettingsParamsWithContext(ctx context.Context) *PatchProjectSettingsParams {
	return &PatchProjectSettingsParams{
		Context: ctx,
	}
}

// NewPatchProjectSettingsParamsWithHTTPClient creates a new PatchProjectSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchProjectSettingsParamsWithHTTPClient(client *http.Client) *PatchProjectSettingsParams {
	return &PatchProjectSettingsParams{
		HTTPClient: client,
	}
}

/*
PatchProjectSettingsParams contains all the parameters to send to the API endpoint

	for the patch project settings operation.

	Typically these are written to a http.Request.
*/
type PatchProjectSettingsParams struct {

	/* Body.

	   Project settings body
	*/
	Body *service_model.V1ProjectSettings

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* Project.

	   Project name
	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch project settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchProjectSettingsParams) WithDefaults() *PatchProjectSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch project settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchProjectSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch project settings params
func (o *PatchProjectSettingsParams) WithTimeout(timeout time.Duration) *PatchProjectSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch project settings params
func (o *PatchProjectSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch project settings params
func (o *PatchProjectSettingsParams) WithContext(ctx context.Context) *PatchProjectSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch project settings params
func (o *PatchProjectSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch project settings params
func (o *PatchProjectSettingsParams) WithHTTPClient(client *http.Client) *PatchProjectSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch project settings params
func (o *PatchProjectSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch project settings params
func (o *PatchProjectSettingsParams) WithBody(body *service_model.V1ProjectSettings) *PatchProjectSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch project settings params
func (o *PatchProjectSettingsParams) SetBody(body *service_model.V1ProjectSettings) {
	o.Body = body
}

// WithOwner adds the owner to the patch project settings params
func (o *PatchProjectSettingsParams) WithOwner(owner string) *PatchProjectSettingsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the patch project settings params
func (o *PatchProjectSettingsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the patch project settings params
func (o *PatchProjectSettingsParams) WithProject(project string) *PatchProjectSettingsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the patch project settings params
func (o *PatchProjectSettingsParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *PatchProjectSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
