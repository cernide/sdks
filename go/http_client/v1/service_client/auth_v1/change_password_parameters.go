// Code generated by go-swagger; DO NOT EDIT.

package auth_v1

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

// NewChangePasswordParams creates a new ChangePasswordParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChangePasswordParams() *ChangePasswordParams {
	return &ChangePasswordParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChangePasswordParamsWithTimeout creates a new ChangePasswordParams object
// with the ability to set a timeout on a request.
func NewChangePasswordParamsWithTimeout(timeout time.Duration) *ChangePasswordParams {
	return &ChangePasswordParams{
		timeout: timeout,
	}
}

// NewChangePasswordParamsWithContext creates a new ChangePasswordParams object
// with the ability to set a context for a request.
func NewChangePasswordParamsWithContext(ctx context.Context) *ChangePasswordParams {
	return &ChangePasswordParams{
		Context: ctx,
	}
}

// NewChangePasswordParamsWithHTTPClient creates a new ChangePasswordParams object
// with the ability to set a custom HTTPClient for a request.
func NewChangePasswordParamsWithHTTPClient(client *http.Client) *ChangePasswordParams {
	return &ChangePasswordParams{
		HTTPClient: client,
	}
}

/*
ChangePasswordParams contains all the parameters to send to the API endpoint

	for the change password operation.

	Typically these are written to a http.Request.
*/
type ChangePasswordParams struct {

	// Body.
	Body *service_model.V1PasswordChange

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the change password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangePasswordParams) WithDefaults() *ChangePasswordParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the change password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangePasswordParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the change password params
func (o *ChangePasswordParams) WithTimeout(timeout time.Duration) *ChangePasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change password params
func (o *ChangePasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change password params
func (o *ChangePasswordParams) WithContext(ctx context.Context) *ChangePasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change password params
func (o *ChangePasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change password params
func (o *ChangePasswordParams) WithHTTPClient(client *http.Client) *ChangePasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change password params
func (o *ChangePasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change password params
func (o *ChangePasswordParams) WithBody(body *service_model.V1PasswordChange) *ChangePasswordParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change password params
func (o *ChangePasswordParams) SetBody(body *service_model.V1PasswordChange) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ChangePasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
