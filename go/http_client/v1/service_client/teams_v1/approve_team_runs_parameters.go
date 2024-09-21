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

// NewApproveTeamRunsParams creates a new ApproveTeamRunsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApproveTeamRunsParams() *ApproveTeamRunsParams {
	return &ApproveTeamRunsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApproveTeamRunsParamsWithTimeout creates a new ApproveTeamRunsParams object
// with the ability to set a timeout on a request.
func NewApproveTeamRunsParamsWithTimeout(timeout time.Duration) *ApproveTeamRunsParams {
	return &ApproveTeamRunsParams{
		timeout: timeout,
	}
}

// NewApproveTeamRunsParamsWithContext creates a new ApproveTeamRunsParams object
// with the ability to set a context for a request.
func NewApproveTeamRunsParamsWithContext(ctx context.Context) *ApproveTeamRunsParams {
	return &ApproveTeamRunsParams{
		Context: ctx,
	}
}

// NewApproveTeamRunsParamsWithHTTPClient creates a new ApproveTeamRunsParams object
// with the ability to set a custom HTTPClient for a request.
func NewApproveTeamRunsParamsWithHTTPClient(client *http.Client) *ApproveTeamRunsParams {
	return &ApproveTeamRunsParams{
		HTTPClient: client,
	}
}

/*
ApproveTeamRunsParams contains all the parameters to send to the API endpoint

	for the approve team runs operation.

	Typically these are written to a http.Request.
*/
type ApproveTeamRunsParams struct {

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

// WithDefaults hydrates default values in the approve team runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApproveTeamRunsParams) WithDefaults() *ApproveTeamRunsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the approve team runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApproveTeamRunsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the approve team runs params
func (o *ApproveTeamRunsParams) WithTimeout(timeout time.Duration) *ApproveTeamRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the approve team runs params
func (o *ApproveTeamRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the approve team runs params
func (o *ApproveTeamRunsParams) WithContext(ctx context.Context) *ApproveTeamRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the approve team runs params
func (o *ApproveTeamRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the approve team runs params
func (o *ApproveTeamRunsParams) WithHTTPClient(client *http.Client) *ApproveTeamRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the approve team runs params
func (o *ApproveTeamRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the approve team runs params
func (o *ApproveTeamRunsParams) WithBody(body *service_model.V1Uuids) *ApproveTeamRunsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the approve team runs params
func (o *ApproveTeamRunsParams) SetBody(body *service_model.V1Uuids) {
	o.Body = body
}

// WithName adds the name to the approve team runs params
func (o *ApproveTeamRunsParams) WithName(name string) *ApproveTeamRunsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the approve team runs params
func (o *ApproveTeamRunsParams) SetName(name string) {
	o.Name = name
}

// WithOwner adds the owner to the approve team runs params
func (o *ApproveTeamRunsParams) WithOwner(owner string) *ApproveTeamRunsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the approve team runs params
func (o *ApproveTeamRunsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *ApproveTeamRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
