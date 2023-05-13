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

// NewInspectRunParams creates a new InspectRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInspectRunParams() *InspectRunParams {
	return &InspectRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInspectRunParamsWithTimeout creates a new InspectRunParams object
// with the ability to set a timeout on a request.
func NewInspectRunParamsWithTimeout(timeout time.Duration) *InspectRunParams {
	return &InspectRunParams{
		timeout: timeout,
	}
}

// NewInspectRunParamsWithContext creates a new InspectRunParams object
// with the ability to set a context for a request.
func NewInspectRunParamsWithContext(ctx context.Context) *InspectRunParams {
	return &InspectRunParams{
		Context: ctx,
	}
}

// NewInspectRunParamsWithHTTPClient creates a new InspectRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewInspectRunParamsWithHTTPClient(client *http.Client) *InspectRunParams {
	return &InspectRunParams{
		HTTPClient: client,
	}
}

/*
InspectRunParams contains all the parameters to send to the API endpoint

	for the inspect run operation.

	Typically these are written to a http.Request.
*/
type InspectRunParams struct {

	/* Force.

	   Force query param.
	*/
	Force *bool

	/* Names.

	   Names query param.
	*/
	Names *string

	/* Namespace.

	   namespace
	*/
	Namespace string

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

	/* Tail.

	   Query param flag to tail the values.
	*/
	Tail *bool

	/* UUID.

	   Uuid identifier of the entity
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the inspect run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InspectRunParams) WithDefaults() *InspectRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the inspect run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InspectRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the inspect run params
func (o *InspectRunParams) WithTimeout(timeout time.Duration) *InspectRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inspect run params
func (o *InspectRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inspect run params
func (o *InspectRunParams) WithContext(ctx context.Context) *InspectRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inspect run params
func (o *InspectRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inspect run params
func (o *InspectRunParams) WithHTTPClient(client *http.Client) *InspectRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inspect run params
func (o *InspectRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the inspect run params
func (o *InspectRunParams) WithForce(force *bool) *InspectRunParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the inspect run params
func (o *InspectRunParams) SetForce(force *bool) {
	o.Force = force
}

// WithNames adds the names to the inspect run params
func (o *InspectRunParams) WithNames(names *string) *InspectRunParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the inspect run params
func (o *InspectRunParams) SetNames(names *string) {
	o.Names = names
}

// WithNamespace adds the namespace to the inspect run params
func (o *InspectRunParams) WithNamespace(namespace string) *InspectRunParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the inspect run params
func (o *InspectRunParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOwner adds the owner to the inspect run params
func (o *InspectRunParams) WithOwner(owner string) *InspectRunParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the inspect run params
func (o *InspectRunParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the inspect run params
func (o *InspectRunParams) WithProject(project string) *InspectRunParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the inspect run params
func (o *InspectRunParams) SetProject(project string) {
	o.Project = project
}

// WithSample adds the sample to the inspect run params
func (o *InspectRunParams) WithSample(sample *int32) *InspectRunParams {
	o.SetSample(sample)
	return o
}

// SetSample adds the sample to the inspect run params
func (o *InspectRunParams) SetSample(sample *int32) {
	o.Sample = sample
}

// WithTail adds the tail to the inspect run params
func (o *InspectRunParams) WithTail(tail *bool) *InspectRunParams {
	o.SetTail(tail)
	return o
}

// SetTail adds the tail to the inspect run params
func (o *InspectRunParams) SetTail(tail *bool) {
	o.Tail = tail
}

// WithUUID adds the uuid to the inspect run params
func (o *InspectRunParams) WithUUID(uuid string) *InspectRunParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the inspect run params
func (o *InspectRunParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *InspectRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Tail != nil {

		// query param tail
		var qrTail bool

		if o.Tail != nil {
			qrTail = *o.Tail
		}
		qTail := swag.FormatBool(qrTail)
		if qTail != "" {

			if err := r.SetQueryParam("tail", qTail); err != nil {
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
