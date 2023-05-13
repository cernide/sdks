// Code generated by go-swagger; DO NOT EDIT.

package users_v1

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

// NewPatchTokenParams creates a new PatchTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchTokenParams() *PatchTokenParams {
	return &PatchTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchTokenParamsWithTimeout creates a new PatchTokenParams object
// with the ability to set a timeout on a request.
func NewPatchTokenParamsWithTimeout(timeout time.Duration) *PatchTokenParams {
	return &PatchTokenParams{
		timeout: timeout,
	}
}

// NewPatchTokenParamsWithContext creates a new PatchTokenParams object
// with the ability to set a context for a request.
func NewPatchTokenParamsWithContext(ctx context.Context) *PatchTokenParams {
	return &PatchTokenParams{
		Context: ctx,
	}
}

// NewPatchTokenParamsWithHTTPClient creates a new PatchTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchTokenParamsWithHTTPClient(client *http.Client) *PatchTokenParams {
	return &PatchTokenParams{
		HTTPClient: client,
	}
}

/*
PatchTokenParams contains all the parameters to send to the API endpoint

	for the patch token operation.

	Typically these are written to a http.Request.
*/
type PatchTokenParams struct {

	/* Body.

	   Token body
	*/
	Body *service_model.V1Token

	/* TokenUUID.

	   UUID
	*/
	TokenUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchTokenParams) WithDefaults() *PatchTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch token params
func (o *PatchTokenParams) WithTimeout(timeout time.Duration) *PatchTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch token params
func (o *PatchTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch token params
func (o *PatchTokenParams) WithContext(ctx context.Context) *PatchTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch token params
func (o *PatchTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch token params
func (o *PatchTokenParams) WithHTTPClient(client *http.Client) *PatchTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch token params
func (o *PatchTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch token params
func (o *PatchTokenParams) WithBody(body *service_model.V1Token) *PatchTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch token params
func (o *PatchTokenParams) SetBody(body *service_model.V1Token) {
	o.Body = body
}

// WithTokenUUID adds the tokenUUID to the patch token params
func (o *PatchTokenParams) WithTokenUUID(tokenUUID string) *PatchTokenParams {
	o.SetTokenUUID(tokenUUID)
	return o
}

// SetTokenUUID adds the tokenUuid to the patch token params
func (o *PatchTokenParams) SetTokenUUID(tokenUUID string) {
	o.TokenUUID = tokenUUID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param token.uuid
	if err := r.SetPathParam("token.uuid", o.TokenUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
