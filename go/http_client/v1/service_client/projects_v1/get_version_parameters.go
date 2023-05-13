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
)

// NewGetVersionParams creates a new GetVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVersionParams() *GetVersionParams {
	return &GetVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVersionParamsWithTimeout creates a new GetVersionParams object
// with the ability to set a timeout on a request.
func NewGetVersionParamsWithTimeout(timeout time.Duration) *GetVersionParams {
	return &GetVersionParams{
		timeout: timeout,
	}
}

// NewGetVersionParamsWithContext creates a new GetVersionParams object
// with the ability to set a context for a request.
func NewGetVersionParamsWithContext(ctx context.Context) *GetVersionParams {
	return &GetVersionParams{
		Context: ctx,
	}
}

// NewGetVersionParamsWithHTTPClient creates a new GetVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVersionParamsWithHTTPClient(client *http.Client) *GetVersionParams {
	return &GetVersionParams{
		HTTPClient: client,
	}
}

/*
GetVersionParams contains all the parameters to send to the API endpoint

	for the get version operation.

	Typically these are written to a http.Request.
*/
type GetVersionParams struct {

	/* Entity.

	   Entity: project name, hub name, registry name, ...
	*/
	Entity string

	/* Kind.

	   Version Kind
	*/
	Kind string

	/* Name.

	   Sub-entity name
	*/
	Name string

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVersionParams) WithDefaults() *GetVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get version params
func (o *GetVersionParams) WithTimeout(timeout time.Duration) *GetVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get version params
func (o *GetVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get version params
func (o *GetVersionParams) WithContext(ctx context.Context) *GetVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get version params
func (o *GetVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get version params
func (o *GetVersionParams) WithHTTPClient(client *http.Client) *GetVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get version params
func (o *GetVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntity adds the entity to the get version params
func (o *GetVersionParams) WithEntity(entity string) *GetVersionParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the get version params
func (o *GetVersionParams) SetEntity(entity string) {
	o.Entity = entity
}

// WithKind adds the kind to the get version params
func (o *GetVersionParams) WithKind(kind string) *GetVersionParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the get version params
func (o *GetVersionParams) SetKind(kind string) {
	o.Kind = kind
}

// WithName adds the name to the get version params
func (o *GetVersionParams) WithName(name string) *GetVersionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get version params
func (o *GetVersionParams) SetName(name string) {
	o.Name = name
}

// WithOwner adds the owner to the get version params
func (o *GetVersionParams) WithOwner(owner string) *GetVersionParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get version params
func (o *GetVersionParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *GetVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entity
	if err := r.SetPathParam("entity", o.Entity); err != nil {
		return err
	}

	// path param kind
	if err := r.SetPathParam("kind", o.Kind); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
