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

// NewUpdateOrganizationMemberParams creates a new UpdateOrganizationMemberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOrganizationMemberParams() *UpdateOrganizationMemberParams {
	return &UpdateOrganizationMemberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrganizationMemberParamsWithTimeout creates a new UpdateOrganizationMemberParams object
// with the ability to set a timeout on a request.
func NewUpdateOrganizationMemberParamsWithTimeout(timeout time.Duration) *UpdateOrganizationMemberParams {
	return &UpdateOrganizationMemberParams{
		timeout: timeout,
	}
}

// NewUpdateOrganizationMemberParamsWithContext creates a new UpdateOrganizationMemberParams object
// with the ability to set a context for a request.
func NewUpdateOrganizationMemberParamsWithContext(ctx context.Context) *UpdateOrganizationMemberParams {
	return &UpdateOrganizationMemberParams{
		Context: ctx,
	}
}

// NewUpdateOrganizationMemberParamsWithHTTPClient creates a new UpdateOrganizationMemberParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOrganizationMemberParamsWithHTTPClient(client *http.Client) *UpdateOrganizationMemberParams {
	return &UpdateOrganizationMemberParams{
		HTTPClient: client,
	}
}

/*
UpdateOrganizationMemberParams contains all the parameters to send to the API endpoint

	for the update organization member operation.

	Typically these are written to a http.Request.
*/
type UpdateOrganizationMemberParams struct {

	/* Body.

	   Organization body
	*/
	Body *service_model.V1OrganizationMember

	/* Email.

	   Optional email.
	*/
	Email *string

	/* MemberUser.

	   User
	*/
	MemberUser string

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update organization member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationMemberParams) WithDefaults() *UpdateOrganizationMemberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update organization member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationMemberParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update organization member params
func (o *UpdateOrganizationMemberParams) WithTimeout(timeout time.Duration) *UpdateOrganizationMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update organization member params
func (o *UpdateOrganizationMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update organization member params
func (o *UpdateOrganizationMemberParams) WithContext(ctx context.Context) *UpdateOrganizationMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update organization member params
func (o *UpdateOrganizationMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update organization member params
func (o *UpdateOrganizationMemberParams) WithHTTPClient(client *http.Client) *UpdateOrganizationMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update organization member params
func (o *UpdateOrganizationMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update organization member params
func (o *UpdateOrganizationMemberParams) WithBody(body *service_model.V1OrganizationMember) *UpdateOrganizationMemberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update organization member params
func (o *UpdateOrganizationMemberParams) SetBody(body *service_model.V1OrganizationMember) {
	o.Body = body
}

// WithEmail adds the email to the update organization member params
func (o *UpdateOrganizationMemberParams) WithEmail(email *string) *UpdateOrganizationMemberParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the update organization member params
func (o *UpdateOrganizationMemberParams) SetEmail(email *string) {
	o.Email = email
}

// WithMemberUser adds the memberUser to the update organization member params
func (o *UpdateOrganizationMemberParams) WithMemberUser(memberUser string) *UpdateOrganizationMemberParams {
	o.SetMemberUser(memberUser)
	return o
}

// SetMemberUser adds the memberUser to the update organization member params
func (o *UpdateOrganizationMemberParams) SetMemberUser(memberUser string) {
	o.MemberUser = memberUser
}

// WithOwner adds the owner to the update organization member params
func (o *UpdateOrganizationMemberParams) WithOwner(owner string) *UpdateOrganizationMemberParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the update organization member params
func (o *UpdateOrganizationMemberParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param member.user
	if err := r.SetPathParam("member.user", o.MemberUser); err != nil {
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
