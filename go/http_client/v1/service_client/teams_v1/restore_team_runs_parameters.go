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

// NewRestoreTeamRunsParams creates a new RestoreTeamRunsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestoreTeamRunsParams() *RestoreTeamRunsParams {
	return &RestoreTeamRunsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestoreTeamRunsParamsWithTimeout creates a new RestoreTeamRunsParams object
// with the ability to set a timeout on a request.
func NewRestoreTeamRunsParamsWithTimeout(timeout time.Duration) *RestoreTeamRunsParams {
	return &RestoreTeamRunsParams{
		timeout: timeout,
	}
}

// NewRestoreTeamRunsParamsWithContext creates a new RestoreTeamRunsParams object
// with the ability to set a context for a request.
func NewRestoreTeamRunsParamsWithContext(ctx context.Context) *RestoreTeamRunsParams {
	return &RestoreTeamRunsParams{
		Context: ctx,
	}
}

// NewRestoreTeamRunsParamsWithHTTPClient creates a new RestoreTeamRunsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestoreTeamRunsParamsWithHTTPClient(client *http.Client) *RestoreTeamRunsParams {
	return &RestoreTeamRunsParams{
		HTTPClient: client,
	}
}

/*
RestoreTeamRunsParams contains all the parameters to send to the API endpoint

	for the restore team runs operation.

	Typically these are written to a http.Request.
*/
type RestoreTeamRunsParams struct {

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

// WithDefaults hydrates default values in the restore team runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoreTeamRunsParams) WithDefaults() *RestoreTeamRunsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the restore team runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoreTeamRunsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the restore team runs params
func (o *RestoreTeamRunsParams) WithTimeout(timeout time.Duration) *RestoreTeamRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restore team runs params
func (o *RestoreTeamRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restore team runs params
func (o *RestoreTeamRunsParams) WithContext(ctx context.Context) *RestoreTeamRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restore team runs params
func (o *RestoreTeamRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restore team runs params
func (o *RestoreTeamRunsParams) WithHTTPClient(client *http.Client) *RestoreTeamRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restore team runs params
func (o *RestoreTeamRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the restore team runs params
func (o *RestoreTeamRunsParams) WithBody(body *service_model.V1Uuids) *RestoreTeamRunsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the restore team runs params
func (o *RestoreTeamRunsParams) SetBody(body *service_model.V1Uuids) {
	o.Body = body
}

// WithName adds the name to the restore team runs params
func (o *RestoreTeamRunsParams) WithName(name string) *RestoreTeamRunsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the restore team runs params
func (o *RestoreTeamRunsParams) SetName(name string) {
	o.Name = name
}

// WithOwner adds the owner to the restore team runs params
func (o *RestoreTeamRunsParams) WithOwner(owner string) *RestoreTeamRunsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the restore team runs params
func (o *RestoreTeamRunsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *RestoreTeamRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
