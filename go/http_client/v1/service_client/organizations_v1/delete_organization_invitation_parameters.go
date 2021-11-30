// Copyright 2018-2021 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
)

// NewDeleteOrganizationInvitationParams creates a new DeleteOrganizationInvitationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOrganizationInvitationParams() *DeleteOrganizationInvitationParams {
	return &DeleteOrganizationInvitationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrganizationInvitationParamsWithTimeout creates a new DeleteOrganizationInvitationParams object
// with the ability to set a timeout on a request.
func NewDeleteOrganizationInvitationParamsWithTimeout(timeout time.Duration) *DeleteOrganizationInvitationParams {
	return &DeleteOrganizationInvitationParams{
		timeout: timeout,
	}
}

// NewDeleteOrganizationInvitationParamsWithContext creates a new DeleteOrganizationInvitationParams object
// with the ability to set a context for a request.
func NewDeleteOrganizationInvitationParamsWithContext(ctx context.Context) *DeleteOrganizationInvitationParams {
	return &DeleteOrganizationInvitationParams{
		Context: ctx,
	}
}

// NewDeleteOrganizationInvitationParamsWithHTTPClient creates a new DeleteOrganizationInvitationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOrganizationInvitationParamsWithHTTPClient(client *http.Client) *DeleteOrganizationInvitationParams {
	return &DeleteOrganizationInvitationParams{
		HTTPClient: client,
	}
}

/* DeleteOrganizationInvitationParams contains all the parameters to send to the API endpoint
   for the delete organization invitation operation.

   Typically these are written to a http.Request.
*/
type DeleteOrganizationInvitationParams struct {

	/* Email.

	   Optional email.
	*/
	Email *string

	/* MemberCreatedAt.

	   Optional time when the entity was created.

	   Format: date-time
	*/
	MemberCreatedAt *strfmt.DateTime

	/* MemberKind.

	   Kind.
	*/
	MemberKind *string

	/* MemberRole.

	   Role.
	*/
	MemberRole *string

	/* MemberUpdatedAt.

	   Optional last time the entity was updated.

	   Format: date-time
	*/
	MemberUpdatedAt *strfmt.DateTime

	/* MemberUser.

	   User.
	*/
	MemberUser *string

	/* MemberUserEmail.

	   Read-only User email.
	*/
	MemberUserEmail *string

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete organization invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationInvitationParams) WithDefaults() *DeleteOrganizationInvitationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete organization invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationInvitationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) WithTimeout(timeout time.Duration) *DeleteOrganizationInvitationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) WithContext(ctx context.Context) *DeleteOrganizationInvitationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) WithHTTPClient(client *http.Client) *DeleteOrganizationInvitationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmail adds the email to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) WithEmail(email *string) *DeleteOrganizationInvitationParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) SetEmail(email *string) {
	o.Email = email
}

// WithMemberCreatedAt adds the memberCreatedAt to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) WithMemberCreatedAt(memberCreatedAt *strfmt.DateTime) *DeleteOrganizationInvitationParams {
	o.SetMemberCreatedAt(memberCreatedAt)
	return o
}

// SetMemberCreatedAt adds the memberCreatedAt to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) SetMemberCreatedAt(memberCreatedAt *strfmt.DateTime) {
	o.MemberCreatedAt = memberCreatedAt
}

// WithMemberKind adds the memberKind to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) WithMemberKind(memberKind *string) *DeleteOrganizationInvitationParams {
	o.SetMemberKind(memberKind)
	return o
}

// SetMemberKind adds the memberKind to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) SetMemberKind(memberKind *string) {
	o.MemberKind = memberKind
}

// WithMemberRole adds the memberRole to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) WithMemberRole(memberRole *string) *DeleteOrganizationInvitationParams {
	o.SetMemberRole(memberRole)
	return o
}

// SetMemberRole adds the memberRole to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) SetMemberRole(memberRole *string) {
	o.MemberRole = memberRole
}

// WithMemberUpdatedAt adds the memberUpdatedAt to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) WithMemberUpdatedAt(memberUpdatedAt *strfmt.DateTime) *DeleteOrganizationInvitationParams {
	o.SetMemberUpdatedAt(memberUpdatedAt)
	return o
}

// SetMemberUpdatedAt adds the memberUpdatedAt to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) SetMemberUpdatedAt(memberUpdatedAt *strfmt.DateTime) {
	o.MemberUpdatedAt = memberUpdatedAt
}

// WithMemberUser adds the memberUser to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) WithMemberUser(memberUser *string) *DeleteOrganizationInvitationParams {
	o.SetMemberUser(memberUser)
	return o
}

// SetMemberUser adds the memberUser to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) SetMemberUser(memberUser *string) {
	o.MemberUser = memberUser
}

// WithMemberUserEmail adds the memberUserEmail to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) WithMemberUserEmail(memberUserEmail *string) *DeleteOrganizationInvitationParams {
	o.SetMemberUserEmail(memberUserEmail)
	return o
}

// SetMemberUserEmail adds the memberUserEmail to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) SetMemberUserEmail(memberUserEmail *string) {
	o.MemberUserEmail = memberUserEmail
}

// WithOwner adds the owner to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) WithOwner(owner string) *DeleteOrganizationInvitationParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the delete organization invitation params
func (o *DeleteOrganizationInvitationParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrganizationInvitationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.MemberCreatedAt != nil {

		// query param member.created_at
		var qrMemberCreatedAt strfmt.DateTime

		if o.MemberCreatedAt != nil {
			qrMemberCreatedAt = *o.MemberCreatedAt
		}
		qMemberCreatedAt := qrMemberCreatedAt.String()
		if qMemberCreatedAt != "" {

			if err := r.SetQueryParam("member.created_at", qMemberCreatedAt); err != nil {
				return err
			}
		}
	}

	if o.MemberKind != nil {

		// query param member.kind
		var qrMemberKind string

		if o.MemberKind != nil {
			qrMemberKind = *o.MemberKind
		}
		qMemberKind := qrMemberKind
		if qMemberKind != "" {

			if err := r.SetQueryParam("member.kind", qMemberKind); err != nil {
				return err
			}
		}
	}

	if o.MemberRole != nil {

		// query param member.role
		var qrMemberRole string

		if o.MemberRole != nil {
			qrMemberRole = *o.MemberRole
		}
		qMemberRole := qrMemberRole
		if qMemberRole != "" {

			if err := r.SetQueryParam("member.role", qMemberRole); err != nil {
				return err
			}
		}
	}

	if o.MemberUpdatedAt != nil {

		// query param member.updated_at
		var qrMemberUpdatedAt strfmt.DateTime

		if o.MemberUpdatedAt != nil {
			qrMemberUpdatedAt = *o.MemberUpdatedAt
		}
		qMemberUpdatedAt := qrMemberUpdatedAt.String()
		if qMemberUpdatedAt != "" {

			if err := r.SetQueryParam("member.updated_at", qMemberUpdatedAt); err != nil {
				return err
			}
		}
	}

	if o.MemberUser != nil {

		// query param member.user
		var qrMemberUser string

		if o.MemberUser != nil {
			qrMemberUser = *o.MemberUser
		}
		qMemberUser := qrMemberUser
		if qMemberUser != "" {

			if err := r.SetQueryParam("member.user", qMemberUser); err != nil {
				return err
			}
		}
	}

	if o.MemberUserEmail != nil {

		// query param member.user_email
		var qrMemberUserEmail string

		if o.MemberUserEmail != nil {
			qrMemberUserEmail = *o.MemberUserEmail
		}
		qMemberUserEmail := qrMemberUserEmail
		if qMemberUserEmail != "" {

			if err := r.SetQueryParam("member.user_email", qMemberUserEmail); err != nil {
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
