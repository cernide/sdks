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

// NewUpdateVersionParams creates a new UpdateVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVersionParams() *UpdateVersionParams {
	return &UpdateVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVersionParamsWithTimeout creates a new UpdateVersionParams object
// with the ability to set a timeout on a request.
func NewUpdateVersionParamsWithTimeout(timeout time.Duration) *UpdateVersionParams {
	return &UpdateVersionParams{
		timeout: timeout,
	}
}

// NewUpdateVersionParamsWithContext creates a new UpdateVersionParams object
// with the ability to set a context for a request.
func NewUpdateVersionParamsWithContext(ctx context.Context) *UpdateVersionParams {
	return &UpdateVersionParams{
		Context: ctx,
	}
}

// NewUpdateVersionParamsWithHTTPClient creates a new UpdateVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVersionParamsWithHTTPClient(client *http.Client) *UpdateVersionParams {
	return &UpdateVersionParams{
		HTTPClient: client,
	}
}

/*
UpdateVersionParams contains all the parameters to send to the API endpoint

	for the update version operation.

	Typically these are written to a http.Request.
*/
type UpdateVersionParams struct {

	/* Body.

	   Project version body
	*/
	Body *service_model.V1ProjectVersion

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* Project.

	   Project name
	*/
	Project string

	/* VersionKind.

	   Optional kind to tell the kind of this version
	*/
	VersionKind string

	/* VersionName.

	   Optional component name, should be a valid fully qualified value: name[:version]
	*/
	VersionName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVersionParams) WithDefaults() *UpdateVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update version params
func (o *UpdateVersionParams) WithTimeout(timeout time.Duration) *UpdateVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update version params
func (o *UpdateVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update version params
func (o *UpdateVersionParams) WithContext(ctx context.Context) *UpdateVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update version params
func (o *UpdateVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update version params
func (o *UpdateVersionParams) WithHTTPClient(client *http.Client) *UpdateVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update version params
func (o *UpdateVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update version params
func (o *UpdateVersionParams) WithBody(body *service_model.V1ProjectVersion) *UpdateVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update version params
func (o *UpdateVersionParams) SetBody(body *service_model.V1ProjectVersion) {
	o.Body = body
}

// WithOwner adds the owner to the update version params
func (o *UpdateVersionParams) WithOwner(owner string) *UpdateVersionParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the update version params
func (o *UpdateVersionParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the update version params
func (o *UpdateVersionParams) WithProject(project string) *UpdateVersionParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the update version params
func (o *UpdateVersionParams) SetProject(project string) {
	o.Project = project
}

// WithVersionKind adds the versionKind to the update version params
func (o *UpdateVersionParams) WithVersionKind(versionKind string) *UpdateVersionParams {
	o.SetVersionKind(versionKind)
	return o
}

// SetVersionKind adds the versionKind to the update version params
func (o *UpdateVersionParams) SetVersionKind(versionKind string) {
	o.VersionKind = versionKind
}

// WithVersionName adds the versionName to the update version params
func (o *UpdateVersionParams) WithVersionName(versionName string) *UpdateVersionParams {
	o.SetVersionName(versionName)
	return o
}

// SetVersionName adds the versionName to the update version params
func (o *UpdateVersionParams) SetVersionName(versionName string) {
	o.VersionName = versionName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param version.kind
	if err := r.SetPathParam("version.kind", o.VersionKind); err != nil {
		return err
	}

	// path param version.name
	if err := r.SetPathParam("version.name", o.VersionName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
