// Code generated by go-swagger; DO NOT EDIT.

package connections_v1

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

// NewGetConnectionParams creates a new GetConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConnectionParams() *GetConnectionParams {
	return &GetConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConnectionParamsWithTimeout creates a new GetConnectionParams object
// with the ability to set a timeout on a request.
func NewGetConnectionParamsWithTimeout(timeout time.Duration) *GetConnectionParams {
	return &GetConnectionParams{
		timeout: timeout,
	}
}

// NewGetConnectionParamsWithContext creates a new GetConnectionParams object
// with the ability to set a context for a request.
func NewGetConnectionParamsWithContext(ctx context.Context) *GetConnectionParams {
	return &GetConnectionParams{
		Context: ctx,
	}
}

// NewGetConnectionParamsWithHTTPClient creates a new GetConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConnectionParamsWithHTTPClient(client *http.Client) *GetConnectionParams {
	return &GetConnectionParams{
		HTTPClient: client,
	}
}

/*
GetConnectionParams contains all the parameters to send to the API endpoint

	for the get connection operation.

	Typically these are written to a http.Request.
*/
type GetConnectionParams struct {

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* UUID.

	   Uuid identifier of the entity
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConnectionParams) WithDefaults() *GetConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get connection params
func (o *GetConnectionParams) WithTimeout(timeout time.Duration) *GetConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get connection params
func (o *GetConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get connection params
func (o *GetConnectionParams) WithContext(ctx context.Context) *GetConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get connection params
func (o *GetConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get connection params
func (o *GetConnectionParams) WithHTTPClient(client *http.Client) *GetConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get connection params
func (o *GetConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the get connection params
func (o *GetConnectionParams) WithOwner(owner string) *GetConnectionParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get connection params
func (o *GetConnectionParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the get connection params
func (o *GetConnectionParams) WithUUID(uuid string) *GetConnectionParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get connection params
func (o *GetConnectionParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
