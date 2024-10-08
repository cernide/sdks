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

// NewUpdateProjectSettingsParams creates a new UpdateProjectSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateProjectSettingsParams() *UpdateProjectSettingsParams {
	return &UpdateProjectSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProjectSettingsParamsWithTimeout creates a new UpdateProjectSettingsParams object
// with the ability to set a timeout on a request.
func NewUpdateProjectSettingsParamsWithTimeout(timeout time.Duration) *UpdateProjectSettingsParams {
	return &UpdateProjectSettingsParams{
		timeout: timeout,
	}
}

// NewUpdateProjectSettingsParamsWithContext creates a new UpdateProjectSettingsParams object
// with the ability to set a context for a request.
func NewUpdateProjectSettingsParamsWithContext(ctx context.Context) *UpdateProjectSettingsParams {
	return &UpdateProjectSettingsParams{
		Context: ctx,
	}
}

// NewUpdateProjectSettingsParamsWithHTTPClient creates a new UpdateProjectSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateProjectSettingsParamsWithHTTPClient(client *http.Client) *UpdateProjectSettingsParams {
	return &UpdateProjectSettingsParams{
		HTTPClient: client,
	}
}

/*
UpdateProjectSettingsParams contains all the parameters to send to the API endpoint

	for the update project settings operation.

	Typically these are written to a http.Request.
*/
type UpdateProjectSettingsParams struct {

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

// WithDefaults hydrates default values in the update project settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProjectSettingsParams) WithDefaults() *UpdateProjectSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update project settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProjectSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update project settings params
func (o *UpdateProjectSettingsParams) WithTimeout(timeout time.Duration) *UpdateProjectSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update project settings params
func (o *UpdateProjectSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update project settings params
func (o *UpdateProjectSettingsParams) WithContext(ctx context.Context) *UpdateProjectSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update project settings params
func (o *UpdateProjectSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update project settings params
func (o *UpdateProjectSettingsParams) WithHTTPClient(client *http.Client) *UpdateProjectSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update project settings params
func (o *UpdateProjectSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update project settings params
func (o *UpdateProjectSettingsParams) WithBody(body *service_model.V1ProjectSettings) *UpdateProjectSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update project settings params
func (o *UpdateProjectSettingsParams) SetBody(body *service_model.V1ProjectSettings) {
	o.Body = body
}

// WithOwner adds the owner to the update project settings params
func (o *UpdateProjectSettingsParams) WithOwner(owner string) *UpdateProjectSettingsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the update project settings params
func (o *UpdateProjectSettingsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the update project settings params
func (o *UpdateProjectSettingsParams) WithProject(project string) *UpdateProjectSettingsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the update project settings params
func (o *UpdateProjectSettingsParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProjectSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
