// Code generated by go-swagger; DO NOT EDIT.

package teams_v1

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

// NewArchiveTeamRunsParams creates a new ArchiveTeamRunsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewArchiveTeamRunsParams() *ArchiveTeamRunsParams {
	return &ArchiveTeamRunsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewArchiveTeamRunsParamsWithTimeout creates a new ArchiveTeamRunsParams object
// with the ability to set a timeout on a request.
func NewArchiveTeamRunsParamsWithTimeout(timeout time.Duration) *ArchiveTeamRunsParams {
	return &ArchiveTeamRunsParams{
		timeout: timeout,
	}
}

// NewArchiveTeamRunsParamsWithContext creates a new ArchiveTeamRunsParams object
// with the ability to set a context for a request.
func NewArchiveTeamRunsParamsWithContext(ctx context.Context) *ArchiveTeamRunsParams {
	return &ArchiveTeamRunsParams{
		Context: ctx,
	}
}

// NewArchiveTeamRunsParamsWithHTTPClient creates a new ArchiveTeamRunsParams object
// with the ability to set a custom HTTPClient for a request.
func NewArchiveTeamRunsParamsWithHTTPClient(client *http.Client) *ArchiveTeamRunsParams {
	return &ArchiveTeamRunsParams{
		HTTPClient: client,
	}
}

/*
ArchiveTeamRunsParams contains all the parameters to send to the API endpoint

	for the archive team runs operation.

	Typically these are written to a http.Request.
*/
type ArchiveTeamRunsParams struct {

	/* Body.

	   Uuids of the entities
	*/
	Body *service_model.V1Uuids

	/* Name.

	   Entity under namespace
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

// WithDefaults hydrates default values in the archive team runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ArchiveTeamRunsParams) WithDefaults() *ArchiveTeamRunsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the archive team runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ArchiveTeamRunsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the archive team runs params
func (o *ArchiveTeamRunsParams) WithTimeout(timeout time.Duration) *ArchiveTeamRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the archive team runs params
func (o *ArchiveTeamRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the archive team runs params
func (o *ArchiveTeamRunsParams) WithContext(ctx context.Context) *ArchiveTeamRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the archive team runs params
func (o *ArchiveTeamRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the archive team runs params
func (o *ArchiveTeamRunsParams) WithHTTPClient(client *http.Client) *ArchiveTeamRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the archive team runs params
func (o *ArchiveTeamRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the archive team runs params
func (o *ArchiveTeamRunsParams) WithBody(body *service_model.V1Uuids) *ArchiveTeamRunsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the archive team runs params
func (o *ArchiveTeamRunsParams) SetBody(body *service_model.V1Uuids) {
	o.Body = body
}

// WithName adds the name to the archive team runs params
func (o *ArchiveTeamRunsParams) WithName(name string) *ArchiveTeamRunsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the archive team runs params
func (o *ArchiveTeamRunsParams) SetName(name string) {
	o.Name = name
}

// WithOwner adds the owner to the archive team runs params
func (o *ArchiveTeamRunsParams) WithOwner(owner string) *ArchiveTeamRunsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the archive team runs params
func (o *ArchiveTeamRunsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *ArchiveTeamRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
