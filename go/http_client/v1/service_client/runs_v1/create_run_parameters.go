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

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewCreateRunParams creates a new CreateRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRunParams() *CreateRunParams {
	return &CreateRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRunParamsWithTimeout creates a new CreateRunParams object
// with the ability to set a timeout on a request.
func NewCreateRunParamsWithTimeout(timeout time.Duration) *CreateRunParams {
	return &CreateRunParams{
		timeout: timeout,
	}
}

// NewCreateRunParamsWithContext creates a new CreateRunParams object
// with the ability to set a context for a request.
func NewCreateRunParamsWithContext(ctx context.Context) *CreateRunParams {
	return &CreateRunParams{
		Context: ctx,
	}
}

// NewCreateRunParamsWithHTTPClient creates a new CreateRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRunParamsWithHTTPClient(client *http.Client) *CreateRunParams {
	return &CreateRunParams{
		HTTPClient: client,
	}
}

/*
CreateRunParams contains all the parameters to send to the API endpoint

	for the create run operation.

	Typically these are written to a http.Request.
*/
type CreateRunParams struct {

	/* Body.

	   operation object
	*/
	Body *service_model.V1OperationBody

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* Project.

	   Project where the run will be assigned
	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRunParams) WithDefaults() *CreateRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create run params
func (o *CreateRunParams) WithTimeout(timeout time.Duration) *CreateRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create run params
func (o *CreateRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create run params
func (o *CreateRunParams) WithContext(ctx context.Context) *CreateRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create run params
func (o *CreateRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create run params
func (o *CreateRunParams) WithHTTPClient(client *http.Client) *CreateRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create run params
func (o *CreateRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create run params
func (o *CreateRunParams) WithBody(body *service_model.V1OperationBody) *CreateRunParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create run params
func (o *CreateRunParams) SetBody(body *service_model.V1OperationBody) {
	o.Body = body
}

// WithOwner adds the owner to the create run params
func (o *CreateRunParams) WithOwner(owner string) *CreateRunParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the create run params
func (o *CreateRunParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the create run params
func (o *CreateRunParams) WithProject(project string) *CreateRunParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the create run params
func (o *CreateRunParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
