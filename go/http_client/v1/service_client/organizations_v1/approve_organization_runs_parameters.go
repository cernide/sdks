// Code generated by go-swagger; DO NOT EDIT.

package organizations_v1

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

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewApproveOrganizationRunsParams creates a new ApproveOrganizationRunsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApproveOrganizationRunsParams() *ApproveOrganizationRunsParams {
	return &ApproveOrganizationRunsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApproveOrganizationRunsParamsWithTimeout creates a new ApproveOrganizationRunsParams object
// with the ability to set a timeout on a request.
func NewApproveOrganizationRunsParamsWithTimeout(timeout time.Duration) *ApproveOrganizationRunsParams {
	return &ApproveOrganizationRunsParams{
		timeout: timeout,
	}
}

// NewApproveOrganizationRunsParamsWithContext creates a new ApproveOrganizationRunsParams object
// with the ability to set a context for a request.
func NewApproveOrganizationRunsParamsWithContext(ctx context.Context) *ApproveOrganizationRunsParams {
	return &ApproveOrganizationRunsParams{
		Context: ctx,
	}
}

// NewApproveOrganizationRunsParamsWithHTTPClient creates a new ApproveOrganizationRunsParams object
// with the ability to set a custom HTTPClient for a request.
func NewApproveOrganizationRunsParamsWithHTTPClient(client *http.Client) *ApproveOrganizationRunsParams {
	return &ApproveOrganizationRunsParams{
		HTTPClient: client,
	}
}

/*
ApproveOrganizationRunsParams contains all the parameters to send to the API endpoint

	for the approve organization runs operation.

	Typically these are written to a http.Request.
*/
type ApproveOrganizationRunsParams struct {

	/* Body.

	   Uuids of the entities
	*/
	Body *service_model.V1Uuids

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the approve organization runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApproveOrganizationRunsParams) WithDefaults() *ApproveOrganizationRunsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the approve organization runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApproveOrganizationRunsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the approve organization runs params
func (o *ApproveOrganizationRunsParams) WithTimeout(timeout time.Duration) *ApproveOrganizationRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the approve organization runs params
func (o *ApproveOrganizationRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the approve organization runs params
func (o *ApproveOrganizationRunsParams) WithContext(ctx context.Context) *ApproveOrganizationRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the approve organization runs params
func (o *ApproveOrganizationRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the approve organization runs params
func (o *ApproveOrganizationRunsParams) WithHTTPClient(client *http.Client) *ApproveOrganizationRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the approve organization runs params
func (o *ApproveOrganizationRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the approve organization runs params
func (o *ApproveOrganizationRunsParams) WithBody(body *service_model.V1Uuids) *ApproveOrganizationRunsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the approve organization runs params
func (o *ApproveOrganizationRunsParams) SetBody(body *service_model.V1Uuids) {
	o.Body = body
}

// WithOwner adds the owner to the approve organization runs params
func (o *ApproveOrganizationRunsParams) WithOwner(owner string) *ApproveOrganizationRunsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the approve organization runs params
func (o *ApproveOrganizationRunsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *ApproveOrganizationRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
