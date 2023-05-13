// Code generated by go-swagger; DO NOT EDIT.

package presets_v1

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

// NewGetPresetParams creates a new GetPresetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPresetParams() *GetPresetParams {
	return &GetPresetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPresetParamsWithTimeout creates a new GetPresetParams object
// with the ability to set a timeout on a request.
func NewGetPresetParamsWithTimeout(timeout time.Duration) *GetPresetParams {
	return &GetPresetParams{
		timeout: timeout,
	}
}

// NewGetPresetParamsWithContext creates a new GetPresetParams object
// with the ability to set a context for a request.
func NewGetPresetParamsWithContext(ctx context.Context) *GetPresetParams {
	return &GetPresetParams{
		Context: ctx,
	}
}

// NewGetPresetParamsWithHTTPClient creates a new GetPresetParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPresetParamsWithHTTPClient(client *http.Client) *GetPresetParams {
	return &GetPresetParams{
		HTTPClient: client,
	}
}

/*
GetPresetParams contains all the parameters to send to the API endpoint

	for the get preset operation.

	Typically these are written to a http.Request.
*/
type GetPresetParams struct {

	/* Entity.

	   Entity: project name, hub name, registry name, ...
	*/
	Entity *string

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

// WithDefaults hydrates default values in the get preset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPresetParams) WithDefaults() *GetPresetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get preset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPresetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get preset params
func (o *GetPresetParams) WithTimeout(timeout time.Duration) *GetPresetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get preset params
func (o *GetPresetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get preset params
func (o *GetPresetParams) WithContext(ctx context.Context) *GetPresetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get preset params
func (o *GetPresetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get preset params
func (o *GetPresetParams) WithHTTPClient(client *http.Client) *GetPresetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get preset params
func (o *GetPresetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntity adds the entity to the get preset params
func (o *GetPresetParams) WithEntity(entity *string) *GetPresetParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the get preset params
func (o *GetPresetParams) SetEntity(entity *string) {
	o.Entity = entity
}

// WithOwner adds the owner to the get preset params
func (o *GetPresetParams) WithOwner(owner string) *GetPresetParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get preset params
func (o *GetPresetParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the get preset params
func (o *GetPresetParams) WithUUID(uuid string) *GetPresetParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get preset params
func (o *GetPresetParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetPresetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Entity != nil {

		// query param entity
		var qrEntity string

		if o.Entity != nil {
			qrEntity = *o.Entity
		}
		qEntity := qrEntity
		if qEntity != "" {

			if err := r.SetQueryParam("entity", qEntity); err != nil {
				return err
			}
		}
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
