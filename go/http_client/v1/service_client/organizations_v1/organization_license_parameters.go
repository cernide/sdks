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

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// NewOrganizationLicenseParams creates a new OrganizationLicenseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrganizationLicenseParams() *OrganizationLicenseParams {
	return &OrganizationLicenseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationLicenseParamsWithTimeout creates a new OrganizationLicenseParams object
// with the ability to set a timeout on a request.
func NewOrganizationLicenseParamsWithTimeout(timeout time.Duration) *OrganizationLicenseParams {
	return &OrganizationLicenseParams{
		timeout: timeout,
	}
}

// NewOrganizationLicenseParamsWithContext creates a new OrganizationLicenseParams object
// with the ability to set a context for a request.
func NewOrganizationLicenseParamsWithContext(ctx context.Context) *OrganizationLicenseParams {
	return &OrganizationLicenseParams{
		Context: ctx,
	}
}

// NewOrganizationLicenseParamsWithHTTPClient creates a new OrganizationLicenseParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrganizationLicenseParamsWithHTTPClient(client *http.Client) *OrganizationLicenseParams {
	return &OrganizationLicenseParams{
		HTTPClient: client,
	}
}

/*
OrganizationLicenseParams contains all the parameters to send to the API endpoint

	for the organization license operation.

	Typically these are written to a http.Request.
*/
type OrganizationLicenseParams struct {

	/* Body.

	   Organization body
	*/
	Body *service_model.V1Organization

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the organization license params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationLicenseParams) WithDefaults() *OrganizationLicenseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the organization license params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationLicenseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the organization license params
func (o *OrganizationLicenseParams) WithTimeout(timeout time.Duration) *OrganizationLicenseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organization license params
func (o *OrganizationLicenseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organization license params
func (o *OrganizationLicenseParams) WithContext(ctx context.Context) *OrganizationLicenseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organization license params
func (o *OrganizationLicenseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organization license params
func (o *OrganizationLicenseParams) WithHTTPClient(client *http.Client) *OrganizationLicenseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organization license params
func (o *OrganizationLicenseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the organization license params
func (o *OrganizationLicenseParams) WithBody(body *service_model.V1Organization) *OrganizationLicenseParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the organization license params
func (o *OrganizationLicenseParams) SetBody(body *service_model.V1Organization) {
	o.Body = body
}

// WithOwner adds the owner to the organization license params
func (o *OrganizationLicenseParams) WithOwner(owner string) *OrganizationLicenseParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the organization license params
func (o *OrganizationLicenseParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationLicenseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
