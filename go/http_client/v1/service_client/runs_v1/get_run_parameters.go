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
)

// NewGetRunParams creates a new GetRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRunParams() *GetRunParams {
	return &GetRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunParamsWithTimeout creates a new GetRunParams object
// with the ability to set a timeout on a request.
func NewGetRunParamsWithTimeout(timeout time.Duration) *GetRunParams {
	return &GetRunParams{
		timeout: timeout,
	}
}

// NewGetRunParamsWithContext creates a new GetRunParams object
// with the ability to set a context for a request.
func NewGetRunParamsWithContext(ctx context.Context) *GetRunParams {
	return &GetRunParams{
		Context: ctx,
	}
}

// NewGetRunParamsWithHTTPClient creates a new GetRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRunParamsWithHTTPClient(client *http.Client) *GetRunParams {
	return &GetRunParams{
		HTTPClient: client,
	}
}

/*
GetRunParams contains all the parameters to send to the API endpoint

	for the get run operation.

	Typically these are written to a http.Request.
*/
type GetRunParams struct {

	/* Entity.

	   Entity: project name, hub name, registry name, ...
	*/
	Entity string

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* UUID.

	   Uuid identifier of the sub-entity
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunParams) WithDefaults() *GetRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get run params
func (o *GetRunParams) WithTimeout(timeout time.Duration) *GetRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get run params
func (o *GetRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get run params
func (o *GetRunParams) WithContext(ctx context.Context) *GetRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get run params
func (o *GetRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get run params
func (o *GetRunParams) WithHTTPClient(client *http.Client) *GetRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get run params
func (o *GetRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntity adds the entity to the get run params
func (o *GetRunParams) WithEntity(entity string) *GetRunParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the get run params
func (o *GetRunParams) SetEntity(entity string) {
	o.Entity = entity
}

// WithOwner adds the owner to the get run params
func (o *GetRunParams) WithOwner(owner string) *GetRunParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get run params
func (o *GetRunParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the get run params
func (o *GetRunParams) WithUUID(uuid string) *GetRunParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get run params
func (o *GetRunParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entity
	if err := r.SetPathParam("entity", o.Entity); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
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
