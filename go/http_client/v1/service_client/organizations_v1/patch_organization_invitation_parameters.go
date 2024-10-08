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

// NewPatchOrganizationInvitationParams creates a new PatchOrganizationInvitationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchOrganizationInvitationParams() *PatchOrganizationInvitationParams {
	return &PatchOrganizationInvitationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchOrganizationInvitationParamsWithTimeout creates a new PatchOrganizationInvitationParams object
// with the ability to set a timeout on a request.
func NewPatchOrganizationInvitationParamsWithTimeout(timeout time.Duration) *PatchOrganizationInvitationParams {
	return &PatchOrganizationInvitationParams{
		timeout: timeout,
	}
}

// NewPatchOrganizationInvitationParamsWithContext creates a new PatchOrganizationInvitationParams object
// with the ability to set a context for a request.
func NewPatchOrganizationInvitationParamsWithContext(ctx context.Context) *PatchOrganizationInvitationParams {
	return &PatchOrganizationInvitationParams{
		Context: ctx,
	}
}

// NewPatchOrganizationInvitationParamsWithHTTPClient creates a new PatchOrganizationInvitationParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchOrganizationInvitationParamsWithHTTPClient(client *http.Client) *PatchOrganizationInvitationParams {
	return &PatchOrganizationInvitationParams{
		HTTPClient: client,
	}
}

/*
PatchOrganizationInvitationParams contains all the parameters to send to the API endpoint

	for the patch organization invitation operation.

	Typically these are written to a http.Request.
*/
type PatchOrganizationInvitationParams struct {

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

// WithDefaults hydrates default values in the patch organization invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchOrganizationInvitationParams) WithDefaults() *PatchOrganizationInvitationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch organization invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchOrganizationInvitationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch organization invitation params
func (o *PatchOrganizationInvitationParams) WithTimeout(timeout time.Duration) *PatchOrganizationInvitationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch organization invitation params
func (o *PatchOrganizationInvitationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch organization invitation params
func (o *PatchOrganizationInvitationParams) WithContext(ctx context.Context) *PatchOrganizationInvitationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch organization invitation params
func (o *PatchOrganizationInvitationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch organization invitation params
func (o *PatchOrganizationInvitationParams) WithHTTPClient(client *http.Client) *PatchOrganizationInvitationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch organization invitation params
func (o *PatchOrganizationInvitationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch organization invitation params
func (o *PatchOrganizationInvitationParams) WithBody(body *service_model.V1OrganizationMember) *PatchOrganizationInvitationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch organization invitation params
func (o *PatchOrganizationInvitationParams) SetBody(body *service_model.V1OrganizationMember) {
	o.Body = body
}

// WithEmail adds the email to the patch organization invitation params
func (o *PatchOrganizationInvitationParams) WithEmail(email *string) *PatchOrganizationInvitationParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the patch organization invitation params
func (o *PatchOrganizationInvitationParams) SetEmail(email *string) {
	o.Email = email
}

// WithOwner adds the owner to the patch organization invitation params
func (o *PatchOrganizationInvitationParams) WithOwner(owner string) *PatchOrganizationInvitationParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the patch organization invitation params
func (o *PatchOrganizationInvitationParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *PatchOrganizationInvitationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
