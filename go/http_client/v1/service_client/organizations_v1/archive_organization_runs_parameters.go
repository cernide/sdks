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

	"github.com/cernide/sdks/go/http_client/v1/service_model"
)

// NewArchiveOrganizationRunsParams creates a new ArchiveOrganizationRunsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewArchiveOrganizationRunsParams() *ArchiveOrganizationRunsParams {
	return &ArchiveOrganizationRunsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewArchiveOrganizationRunsParamsWithTimeout creates a new ArchiveOrganizationRunsParams object
// with the ability to set a timeout on a request.
func NewArchiveOrganizationRunsParamsWithTimeout(timeout time.Duration) *ArchiveOrganizationRunsParams {
	return &ArchiveOrganizationRunsParams{
		timeout: timeout,
	}
}

// NewArchiveOrganizationRunsParamsWithContext creates a new ArchiveOrganizationRunsParams object
// with the ability to set a context for a request.
func NewArchiveOrganizationRunsParamsWithContext(ctx context.Context) *ArchiveOrganizationRunsParams {
	return &ArchiveOrganizationRunsParams{
		Context: ctx,
	}
}

// NewArchiveOrganizationRunsParamsWithHTTPClient creates a new ArchiveOrganizationRunsParams object
// with the ability to set a custom HTTPClient for a request.
func NewArchiveOrganizationRunsParamsWithHTTPClient(client *http.Client) *ArchiveOrganizationRunsParams {
	return &ArchiveOrganizationRunsParams{
		HTTPClient: client,
	}
}

/*
ArchiveOrganizationRunsParams contains all the parameters to send to the API endpoint

	for the archive organization runs operation.

	Typically these are written to a http.Request.
*/
type ArchiveOrganizationRunsParams struct {

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

// WithDefaults hydrates default values in the archive organization runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ArchiveOrganizationRunsParams) WithDefaults() *ArchiveOrganizationRunsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the archive organization runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ArchiveOrganizationRunsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the archive organization runs params
func (o *ArchiveOrganizationRunsParams) WithTimeout(timeout time.Duration) *ArchiveOrganizationRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the archive organization runs params
func (o *ArchiveOrganizationRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the archive organization runs params
func (o *ArchiveOrganizationRunsParams) WithContext(ctx context.Context) *ArchiveOrganizationRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the archive organization runs params
func (o *ArchiveOrganizationRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the archive organization runs params
func (o *ArchiveOrganizationRunsParams) WithHTTPClient(client *http.Client) *ArchiveOrganizationRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the archive organization runs params
func (o *ArchiveOrganizationRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the archive organization runs params
func (o *ArchiveOrganizationRunsParams) WithBody(body *service_model.V1Uuids) *ArchiveOrganizationRunsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the archive organization runs params
func (o *ArchiveOrganizationRunsParams) SetBody(body *service_model.V1Uuids) {
	o.Body = body
}

// WithOwner adds the owner to the archive organization runs params
func (o *ArchiveOrganizationRunsParams) WithOwner(owner string) *ArchiveOrganizationRunsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the archive organization runs params
func (o *ArchiveOrganizationRunsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *ArchiveOrganizationRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
