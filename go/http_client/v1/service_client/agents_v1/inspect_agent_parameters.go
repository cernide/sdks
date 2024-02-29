// Code generated by go-swagger; DO NOT EDIT.

package agents_v1

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

// NewInspectAgentParams creates a new InspectAgentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInspectAgentParams() *InspectAgentParams {
	return &InspectAgentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInspectAgentParamsWithTimeout creates a new InspectAgentParams object
// with the ability to set a timeout on a request.
func NewInspectAgentParamsWithTimeout(timeout time.Duration) *InspectAgentParams {
	return &InspectAgentParams{
		timeout: timeout,
	}
}

// NewInspectAgentParamsWithContext creates a new InspectAgentParams object
// with the ability to set a context for a request.
func NewInspectAgentParamsWithContext(ctx context.Context) *InspectAgentParams {
	return &InspectAgentParams{
		Context: ctx,
	}
}

// NewInspectAgentParamsWithHTTPClient creates a new InspectAgentParams object
// with the ability to set a custom HTTPClient for a request.
func NewInspectAgentParamsWithHTTPClient(client *http.Client) *InspectAgentParams {
	return &InspectAgentParams{
		HTTPClient: client,
	}
}

/*
InspectAgentParams contains all the parameters to send to the API endpoint

	for the inspect agent operation.

	Typically these are written to a http.Request.
*/
type InspectAgentParams struct {

	/* Connection.

	   Connection to use.
	*/
	Connection *string

	/* Force.

	   Force query param.
	*/
	Force *bool

	/* LastFile.

	   last_file.
	*/
	LastFile *string

	/* Namespace.

	   namespace
	*/
	Namespace string

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* Service.

	   Service.
	*/
	Service *string

	/* UUID.

	   Uuid identifier of the entity
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the inspect agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InspectAgentParams) WithDefaults() *InspectAgentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the inspect agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InspectAgentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the inspect agent params
func (o *InspectAgentParams) WithTimeout(timeout time.Duration) *InspectAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inspect agent params
func (o *InspectAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inspect agent params
func (o *InspectAgentParams) WithContext(ctx context.Context) *InspectAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inspect agent params
func (o *InspectAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inspect agent params
func (o *InspectAgentParams) WithHTTPClient(client *http.Client) *InspectAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inspect agent params
func (o *InspectAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnection adds the connection to the inspect agent params
func (o *InspectAgentParams) WithConnection(connection *string) *InspectAgentParams {
	o.SetConnection(connection)
	return o
}

// SetConnection adds the connection to the inspect agent params
func (o *InspectAgentParams) SetConnection(connection *string) {
	o.Connection = connection
}

// WithForce adds the force to the inspect agent params
func (o *InspectAgentParams) WithForce(force *bool) *InspectAgentParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the inspect agent params
func (o *InspectAgentParams) SetForce(force *bool) {
	o.Force = force
}

// WithLastFile adds the lastFile to the inspect agent params
func (o *InspectAgentParams) WithLastFile(lastFile *string) *InspectAgentParams {
	o.SetLastFile(lastFile)
	return o
}

// SetLastFile adds the lastFile to the inspect agent params
func (o *InspectAgentParams) SetLastFile(lastFile *string) {
	o.LastFile = lastFile
}

// WithNamespace adds the namespace to the inspect agent params
func (o *InspectAgentParams) WithNamespace(namespace string) *InspectAgentParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the inspect agent params
func (o *InspectAgentParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOwner adds the owner to the inspect agent params
func (o *InspectAgentParams) WithOwner(owner string) *InspectAgentParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the inspect agent params
func (o *InspectAgentParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithService adds the service to the inspect agent params
func (o *InspectAgentParams) WithService(service *string) *InspectAgentParams {
	o.SetService(service)
	return o
}

// SetService adds the service to the inspect agent params
func (o *InspectAgentParams) SetService(service *string) {
	o.Service = service
}

// WithUUID adds the uuid to the inspect agent params
func (o *InspectAgentParams) WithUUID(uuid string) *InspectAgentParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the inspect agent params
func (o *InspectAgentParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *InspectAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.LastFile != nil {

		// query param last_file
		var qrLastFile string

		if o.LastFile != nil {
			qrLastFile = *o.LastFile
		}
		qLastFile := qrLastFile
		if qLastFile != "" {

			if err := r.SetQueryParam("last_file", qLastFile); err != nil {
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

	if o.Service != nil {

		// query param service
		var qrService string

		if o.Service != nil {
			qrService = *o.Service
		}
		qService := qrService
		if qService != "" {

			if err := r.SetQueryParam("service", qService); err != nil {
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