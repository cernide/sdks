// Code generated by go-swagger; DO NOT EDIT.

package project_searches_v1

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

// NewGetProjectSearchParams creates a new GetProjectSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectSearchParams() *GetProjectSearchParams {
	return &GetProjectSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectSearchParamsWithTimeout creates a new GetProjectSearchParams object
// with the ability to set a timeout on a request.
func NewGetProjectSearchParamsWithTimeout(timeout time.Duration) *GetProjectSearchParams {
	return &GetProjectSearchParams{
		timeout: timeout,
	}
}

// NewGetProjectSearchParamsWithContext creates a new GetProjectSearchParams object
// with the ability to set a context for a request.
func NewGetProjectSearchParamsWithContext(ctx context.Context) *GetProjectSearchParams {
	return &GetProjectSearchParams{
		Context: ctx,
	}
}

// NewGetProjectSearchParamsWithHTTPClient creates a new GetProjectSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectSearchParamsWithHTTPClient(client *http.Client) *GetProjectSearchParams {
	return &GetProjectSearchParams{
		HTTPClient: client,
	}
}

/*
GetProjectSearchParams contains all the parameters to send to the API endpoint

	for the get project search operation.

	Typically these are written to a http.Request.
*/
type GetProjectSearchParams struct {

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

// WithDefaults hydrates default values in the get project search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectSearchParams) WithDefaults() *GetProjectSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectSearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get project search params
func (o *GetProjectSearchParams) WithTimeout(timeout time.Duration) *GetProjectSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project search params
func (o *GetProjectSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project search params
func (o *GetProjectSearchParams) WithContext(ctx context.Context) *GetProjectSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project search params
func (o *GetProjectSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project search params
func (o *GetProjectSearchParams) WithHTTPClient(client *http.Client) *GetProjectSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project search params
func (o *GetProjectSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntity adds the entity to the get project search params
func (o *GetProjectSearchParams) WithEntity(entity string) *GetProjectSearchParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the get project search params
func (o *GetProjectSearchParams) SetEntity(entity string) {
	o.Entity = entity
}

// WithOwner adds the owner to the get project search params
func (o *GetProjectSearchParams) WithOwner(owner string) *GetProjectSearchParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get project search params
func (o *GetProjectSearchParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the get project search params
func (o *GetProjectSearchParams) WithUUID(uuid string) *GetProjectSearchParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get project search params
func (o *GetProjectSearchParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
