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

// NewTagOrganizationRunsParams creates a new TagOrganizationRunsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTagOrganizationRunsParams() *TagOrganizationRunsParams {
	return &TagOrganizationRunsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTagOrganizationRunsParamsWithTimeout creates a new TagOrganizationRunsParams object
// with the ability to set a timeout on a request.
func NewTagOrganizationRunsParamsWithTimeout(timeout time.Duration) *TagOrganizationRunsParams {
	return &TagOrganizationRunsParams{
		timeout: timeout,
	}
}

// NewTagOrganizationRunsParamsWithContext creates a new TagOrganizationRunsParams object
// with the ability to set a context for a request.
func NewTagOrganizationRunsParamsWithContext(ctx context.Context) *TagOrganizationRunsParams {
	return &TagOrganizationRunsParams{
		Context: ctx,
	}
}

// NewTagOrganizationRunsParamsWithHTTPClient creates a new TagOrganizationRunsParams object
// with the ability to set a custom HTTPClient for a request.
func NewTagOrganizationRunsParamsWithHTTPClient(client *http.Client) *TagOrganizationRunsParams {
	return &TagOrganizationRunsParams{
		HTTPClient: client,
	}
}

/*
TagOrganizationRunsParams contains all the parameters to send to the API endpoint

	for the tag organization runs operation.

	Typically these are written to a http.Request.
*/
type TagOrganizationRunsParams struct {

	/* Body.

	   Data
	*/
	Body *service_model.V1EntitiesTags

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tag organization runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TagOrganizationRunsParams) WithDefaults() *TagOrganizationRunsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tag organization runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TagOrganizationRunsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tag organization runs params
func (o *TagOrganizationRunsParams) WithTimeout(timeout time.Duration) *TagOrganizationRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tag organization runs params
func (o *TagOrganizationRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tag organization runs params
func (o *TagOrganizationRunsParams) WithContext(ctx context.Context) *TagOrganizationRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tag organization runs params
func (o *TagOrganizationRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tag organization runs params
func (o *TagOrganizationRunsParams) WithHTTPClient(client *http.Client) *TagOrganizationRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tag organization runs params
func (o *TagOrganizationRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the tag organization runs params
func (o *TagOrganizationRunsParams) WithBody(body *service_model.V1EntitiesTags) *TagOrganizationRunsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the tag organization runs params
func (o *TagOrganizationRunsParams) SetBody(body *service_model.V1EntitiesTags) {
	o.Body = body
}

// WithOwner adds the owner to the tag organization runs params
func (o *TagOrganizationRunsParams) WithOwner(owner string) *TagOrganizationRunsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the tag organization runs params
func (o *TagOrganizationRunsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *TagOrganizationRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
