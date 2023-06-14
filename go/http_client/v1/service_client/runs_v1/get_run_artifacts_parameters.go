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

// NewGetRunArtifactsParams creates a new GetRunArtifactsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRunArtifactsParams() *GetRunArtifactsParams {
	return &GetRunArtifactsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunArtifactsParamsWithTimeout creates a new GetRunArtifactsParams object
// with the ability to set a timeout on a request.
func NewGetRunArtifactsParamsWithTimeout(timeout time.Duration) *GetRunArtifactsParams {
	return &GetRunArtifactsParams{
		timeout: timeout,
	}
}

// NewGetRunArtifactsParamsWithContext creates a new GetRunArtifactsParams object
// with the ability to set a context for a request.
func NewGetRunArtifactsParamsWithContext(ctx context.Context) *GetRunArtifactsParams {
	return &GetRunArtifactsParams{
		Context: ctx,
	}
}

// NewGetRunArtifactsParamsWithHTTPClient creates a new GetRunArtifactsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRunArtifactsParamsWithHTTPClient(client *http.Client) *GetRunArtifactsParams {
	return &GetRunArtifactsParams{
		HTTPClient: client,
	}
}

/*
GetRunArtifactsParams contains all the parameters to send to the API endpoint

	for the get run artifacts operation.

	Typically these are written to a http.Request.
*/
type GetRunArtifactsParams struct {

	/* Connection.

	   Connection to use.
	*/
	Connection *string

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

	   Project where the entity will be assigned
	*/
	Project string

	/* UUID.

	   Unique integer identifier of the entity
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get run artifacts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunArtifactsParams) WithDefaults() *GetRunArtifactsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get run artifacts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunArtifactsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get run artifacts params
func (o *GetRunArtifactsParams) WithTimeout(timeout time.Duration) *GetRunArtifactsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get run artifacts params
func (o *GetRunArtifactsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get run artifacts params
func (o *GetRunArtifactsParams) WithContext(ctx context.Context) *GetRunArtifactsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get run artifacts params
func (o *GetRunArtifactsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get run artifacts params
func (o *GetRunArtifactsParams) WithHTTPClient(client *http.Client) *GetRunArtifactsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get run artifacts params
func (o *GetRunArtifactsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnection adds the connection to the get run artifacts params
func (o *GetRunArtifactsParams) WithConnection(connection *string) *GetRunArtifactsParams {
	o.SetConnection(connection)
	return o
}

// SetConnection adds the connection to the get run artifacts params
func (o *GetRunArtifactsParams) SetConnection(connection *string) {
	o.Connection = connection
}

// WithForce adds the force to the get run artifacts params
func (o *GetRunArtifactsParams) WithForce(force *bool) *GetRunArtifactsParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the get run artifacts params
func (o *GetRunArtifactsParams) SetForce(force *bool) {
	o.Force = force
}

// WithNamespace adds the namespace to the get run artifacts params
func (o *GetRunArtifactsParams) WithNamespace(namespace string) *GetRunArtifactsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get run artifacts params
func (o *GetRunArtifactsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOwner adds the owner to the get run artifacts params
func (o *GetRunArtifactsParams) WithOwner(owner string) *GetRunArtifactsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get run artifacts params
func (o *GetRunArtifactsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithPath adds the path to the get run artifacts params
func (o *GetRunArtifactsParams) WithPath(path *string) *GetRunArtifactsParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the get run artifacts params
func (o *GetRunArtifactsParams) SetPath(path *string) {
	o.Path = path
}

// WithProject adds the project to the get run artifacts params
func (o *GetRunArtifactsParams) WithProject(project string) *GetRunArtifactsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get run artifacts params
func (o *GetRunArtifactsParams) SetProject(project string) {
	o.Project = project
}

// WithUUID adds the uuid to the get run artifacts params
func (o *GetRunArtifactsParams) WithUUID(uuid string) *GetRunArtifactsParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get run artifacts params
func (o *GetRunArtifactsParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunArtifactsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Connection != nil {

		// query param connection
		var qrConnection string

		if o.Connection != nil {
			qrConnection = *o.Connection
		}
		qConnection := qrConnection
		if qConnection != "" {

			if err := r.SetQueryParam("connection", qConnection); err != nil {
				return err
			}
		}
	}

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

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
