// Code generated by go-swagger; DO NOT EDIT.

package service_accounts_v1

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

// NewCreateServiceAccountTokenParams creates a new CreateServiceAccountTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateServiceAccountTokenParams() *CreateServiceAccountTokenParams {
	return &CreateServiceAccountTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateServiceAccountTokenParamsWithTimeout creates a new CreateServiceAccountTokenParams object
// with the ability to set a timeout on a request.
func NewCreateServiceAccountTokenParamsWithTimeout(timeout time.Duration) *CreateServiceAccountTokenParams {
	return &CreateServiceAccountTokenParams{
		timeout: timeout,
	}
}

// NewCreateServiceAccountTokenParamsWithContext creates a new CreateServiceAccountTokenParams object
// with the ability to set a context for a request.
func NewCreateServiceAccountTokenParamsWithContext(ctx context.Context) *CreateServiceAccountTokenParams {
	return &CreateServiceAccountTokenParams{
		Context: ctx,
	}
}

// NewCreateServiceAccountTokenParamsWithHTTPClient creates a new CreateServiceAccountTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateServiceAccountTokenParamsWithHTTPClient(client *http.Client) *CreateServiceAccountTokenParams {
	return &CreateServiceAccountTokenParams{
		HTTPClient: client,
	}
}

/*
CreateServiceAccountTokenParams contains all the parameters to send to the API endpoint

	for the create service account token operation.

	Typically these are written to a http.Request.
*/
type CreateServiceAccountTokenParams struct {

	/* Body.

	   Token body
	*/
	Body *service_model.V1Token

	/* Entity.

	   Entity
	*/
	Entity string

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create service account token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateServiceAccountTokenParams) WithDefaults() *CreateServiceAccountTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create service account token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateServiceAccountTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create service account token params
func (o *CreateServiceAccountTokenParams) WithTimeout(timeout time.Duration) *CreateServiceAccountTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create service account token params
func (o *CreateServiceAccountTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create service account token params
func (o *CreateServiceAccountTokenParams) WithContext(ctx context.Context) *CreateServiceAccountTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create service account token params
func (o *CreateServiceAccountTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create service account token params
func (o *CreateServiceAccountTokenParams) WithHTTPClient(client *http.Client) *CreateServiceAccountTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create service account token params
func (o *CreateServiceAccountTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create service account token params
func (o *CreateServiceAccountTokenParams) WithBody(body *service_model.V1Token) *CreateServiceAccountTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create service account token params
func (o *CreateServiceAccountTokenParams) SetBody(body *service_model.V1Token) {
	o.Body = body
}

// WithEntity adds the entity to the create service account token params
func (o *CreateServiceAccountTokenParams) WithEntity(entity string) *CreateServiceAccountTokenParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the create service account token params
func (o *CreateServiceAccountTokenParams) SetEntity(entity string) {
	o.Entity = entity
}

// WithOwner adds the owner to the create service account token params
func (o *CreateServiceAccountTokenParams) WithOwner(owner string) *CreateServiceAccountTokenParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the create service account token params
func (o *CreateServiceAccountTokenParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *CreateServiceAccountTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param entity
	if err := r.SetPathParam("entity", o.Entity); err != nil {
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
