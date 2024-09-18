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

	"github.com/cernide/sdks/go/http_client/v1/service_model"
)

// NewSignupParams creates a new SignupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSignupParams() *SignupParams {
	return &SignupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSignupParamsWithTimeout creates a new SignupParams object
// with the ability to set a timeout on a request.
func NewSignupParamsWithTimeout(timeout time.Duration) *SignupParams {
	return &SignupParams{
		timeout: timeout,
	}
}

// NewSignupParamsWithContext creates a new SignupParams object
// with the ability to set a context for a request.
func NewSignupParamsWithContext(ctx context.Context) *SignupParams {
	return &SignupParams{
		Context: ctx,
	}
}

// NewSignupParamsWithHTTPClient creates a new SignupParams object
// with the ability to set a custom HTTPClient for a request.
func NewSignupParamsWithHTTPClient(client *http.Client) *SignupParams {
	return &SignupParams{
		HTTPClient: client,
	}
}

/*
SignupParams contains all the parameters to send to the API endpoint

	for the signup operation.

	Typically these are written to a http.Request.
*/
type SignupParams struct {

	// Body.
	Body *service_model.V1UserSingup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the signup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SignupParams) WithDefaults() *SignupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the signup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SignupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the signup params
func (o *SignupParams) WithTimeout(timeout time.Duration) *SignupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the signup params
func (o *SignupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the signup params
func (o *SignupParams) WithContext(ctx context.Context) *SignupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the signup params
func (o *SignupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the signup params
func (o *SignupParams) WithHTTPClient(client *http.Client) *SignupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the signup params
func (o *SignupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the signup params
func (o *SignupParams) WithBody(body *service_model.V1UserSingup) *SignupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the signup params
func (o *SignupParams) SetBody(body *service_model.V1UserSingup) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SignupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
