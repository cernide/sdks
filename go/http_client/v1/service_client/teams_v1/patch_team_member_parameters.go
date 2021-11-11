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

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewPatchTeamMemberParams creates a new PatchTeamMemberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchTeamMemberParams() *PatchTeamMemberParams {
	return &PatchTeamMemberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchTeamMemberParamsWithTimeout creates a new PatchTeamMemberParams object
// with the ability to set a timeout on a request.
func NewPatchTeamMemberParamsWithTimeout(timeout time.Duration) *PatchTeamMemberParams {
	return &PatchTeamMemberParams{
		timeout: timeout,
	}
}

// NewPatchTeamMemberParamsWithContext creates a new PatchTeamMemberParams object
// with the ability to set a context for a request.
func NewPatchTeamMemberParamsWithContext(ctx context.Context) *PatchTeamMemberParams {
	return &PatchTeamMemberParams{
		Context: ctx,
	}
}

// NewPatchTeamMemberParamsWithHTTPClient creates a new PatchTeamMemberParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchTeamMemberParamsWithHTTPClient(client *http.Client) *PatchTeamMemberParams {
	return &PatchTeamMemberParams{
		HTTPClient: client,
	}
}

/* PatchTeamMemberParams contains all the parameters to send to the API endpoint
   for the patch team member operation.

   Typically these are written to a http.Request.
*/
type PatchTeamMemberParams struct {

	/* Body.

	   Team body
	*/
	Body *service_model.V1TeamMember

	/* MemberUser.

	   User
	*/
	MemberUser string

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* Team.

	   Team
	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch team member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchTeamMemberParams) WithDefaults() *PatchTeamMemberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch team member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchTeamMemberParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch team member params
func (o *PatchTeamMemberParams) WithTimeout(timeout time.Duration) *PatchTeamMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch team member params
func (o *PatchTeamMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch team member params
func (o *PatchTeamMemberParams) WithContext(ctx context.Context) *PatchTeamMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch team member params
func (o *PatchTeamMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch team member params
func (o *PatchTeamMemberParams) WithHTTPClient(client *http.Client) *PatchTeamMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch team member params
func (o *PatchTeamMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch team member params
func (o *PatchTeamMemberParams) WithBody(body *service_model.V1TeamMember) *PatchTeamMemberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch team member params
func (o *PatchTeamMemberParams) SetBody(body *service_model.V1TeamMember) {
	o.Body = body
}

// WithMemberUser adds the memberUser to the patch team member params
func (o *PatchTeamMemberParams) WithMemberUser(memberUser string) *PatchTeamMemberParams {
	o.SetMemberUser(memberUser)
	return o
}

// SetMemberUser adds the memberUser to the patch team member params
func (o *PatchTeamMemberParams) SetMemberUser(memberUser string) {
	o.MemberUser = memberUser
}

// WithOwner adds the owner to the patch team member params
func (o *PatchTeamMemberParams) WithOwner(owner string) *PatchTeamMemberParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the patch team member params
func (o *PatchTeamMemberParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithTeam adds the team to the patch team member params
func (o *PatchTeamMemberParams) WithTeam(team string) *PatchTeamMemberParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the patch team member params
func (o *PatchTeamMemberParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *PatchTeamMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
