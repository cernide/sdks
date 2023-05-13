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

// NewCreateOrganizationMemberParams creates a new CreateOrganizationMemberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOrganizationMemberParams() *CreateOrganizationMemberParams {
	return &CreateOrganizationMemberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrganizationMemberParamsWithTimeout creates a new CreateOrganizationMemberParams object
// with the ability to set a timeout on a request.
func NewCreateOrganizationMemberParamsWithTimeout(timeout time.Duration) *CreateOrganizationMemberParams {
	return &CreateOrganizationMemberParams{
		timeout: timeout,
	}
}

// NewCreateOrganizationMemberParamsWithContext creates a new CreateOrganizationMemberParams object
// with the ability to set a context for a request.
func NewCreateOrganizationMemberParamsWithContext(ctx context.Context) *CreateOrganizationMemberParams {
	return &CreateOrganizationMemberParams{
		Context: ctx,
	}
}

// NewCreateOrganizationMemberParamsWithHTTPClient creates a new CreateOrganizationMemberParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOrganizationMemberParamsWithHTTPClient(client *http.Client) *CreateOrganizationMemberParams {
	return &CreateOrganizationMemberParams{
		HTTPClient: client,
	}
}

/*
CreateOrganizationMemberParams contains all the parameters to send to the API endpoint

	for the create organization member operation.

	Typically these are written to a http.Request.
*/
type CreateOrganizationMemberParams struct {

	/* Body.

	   Organization body
	*/
	Body *service_model.V1OrganizationMember

	/* Email.

	   Optional email.
	*/
	Email *string

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create organization member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationMemberParams) WithDefaults() *CreateOrganizationMemberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create organization member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationMemberParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create organization member params
func (o *CreateOrganizationMemberParams) WithTimeout(timeout time.Duration) *CreateOrganizationMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create organization member params
func (o *CreateOrganizationMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create organization member params
func (o *CreateOrganizationMemberParams) WithContext(ctx context.Context) *CreateOrganizationMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create organization member params
func (o *CreateOrganizationMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create organization member params
func (o *CreateOrganizationMemberParams) WithHTTPClient(client *http.Client) *CreateOrganizationMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create organization member params
func (o *CreateOrganizationMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create organization member params
func (o *CreateOrganizationMemberParams) WithBody(body *service_model.V1OrganizationMember) *CreateOrganizationMemberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create organization member params
func (o *CreateOrganizationMemberParams) SetBody(body *service_model.V1OrganizationMember) {
	o.Body = body
}

// WithEmail adds the email to the create organization member params
func (o *CreateOrganizationMemberParams) WithEmail(email *string) *CreateOrganizationMemberParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the create organization member params
func (o *CreateOrganizationMemberParams) SetEmail(email *string) {
	o.Email = email
}

// WithOwner adds the owner to the create organization member params
func (o *CreateOrganizationMemberParams) WithOwner(owner string) *CreateOrganizationMemberParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the create organization member params
func (o *CreateOrganizationMemberParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrganizationMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Email != nil {

		// query param email
		var qrEmail string

		if o.Email != nil {
			qrEmail = *o.Email
		}
		qEmail := qrEmail
		if qEmail != "" {

			if err := r.SetQueryParam("email", qEmail); err != nil {
				return err
			}
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
