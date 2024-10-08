// Code generated by go-swagger; DO NOT EDIT.

package searches_v1

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

// NewPatchSearchParams creates a new PatchSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchSearchParams() *PatchSearchParams {
	return &PatchSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchSearchParamsWithTimeout creates a new PatchSearchParams object
// with the ability to set a timeout on a request.
func NewPatchSearchParamsWithTimeout(timeout time.Duration) *PatchSearchParams {
	return &PatchSearchParams{
		timeout: timeout,
	}
}

// NewPatchSearchParamsWithContext creates a new PatchSearchParams object
// with the ability to set a context for a request.
func NewPatchSearchParamsWithContext(ctx context.Context) *PatchSearchParams {
	return &PatchSearchParams{
		Context: ctx,
	}
}

// NewPatchSearchParamsWithHTTPClient creates a new PatchSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchSearchParamsWithHTTPClient(client *http.Client) *PatchSearchParams {
	return &PatchSearchParams{
		HTTPClient: client,
	}
}

/*
PatchSearchParams contains all the parameters to send to the API endpoint

	for the patch search operation.

	Typically these are written to a http.Request.
*/
type PatchSearchParams struct {

	/* Body.

	   Search body
	*/
	Body *service_model.V1Search

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* SearchUUID.

	   UUID
	*/
	SearchUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchSearchParams) WithDefaults() *PatchSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchSearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch search params
func (o *PatchSearchParams) WithTimeout(timeout time.Duration) *PatchSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch search params
func (o *PatchSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch search params
func (o *PatchSearchParams) WithContext(ctx context.Context) *PatchSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch search params
func (o *PatchSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch search params
func (o *PatchSearchParams) WithHTTPClient(client *http.Client) *PatchSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch search params
func (o *PatchSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch search params
func (o *PatchSearchParams) WithBody(body *service_model.V1Search) *PatchSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch search params
func (o *PatchSearchParams) SetBody(body *service_model.V1Search) {
	o.Body = body
}

// WithOwner adds the owner to the patch search params
func (o *PatchSearchParams) WithOwner(owner string) *PatchSearchParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the patch search params
func (o *PatchSearchParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithSearchUUID adds the searchUUID to the patch search params
func (o *PatchSearchParams) WithSearchUUID(searchUUID string) *PatchSearchParams {
	o.SetSearchUUID(searchUUID)
	return o
}

// SetSearchUUID adds the searchUuid to the patch search params
func (o *PatchSearchParams) SetSearchUUID(searchUUID string) {
	o.SearchUUID = searchUUID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param search.uuid
	if err := r.SetPathParam("search.uuid", o.SearchUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
