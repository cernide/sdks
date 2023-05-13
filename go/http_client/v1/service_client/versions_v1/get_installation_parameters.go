// Code generated by go-swagger; DO NOT EDIT.

package versions_v1

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
	"github.com/go-openapi/swag"
)

// NewGetInstallationParams creates a new GetInstallationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInstallationParams() *GetInstallationParams {
	return &GetInstallationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstallationParamsWithTimeout creates a new GetInstallationParams object
// with the ability to set a timeout on a request.
func NewGetInstallationParamsWithTimeout(timeout time.Duration) *GetInstallationParams {
	return &GetInstallationParams{
		timeout: timeout,
	}
}

// NewGetInstallationParamsWithContext creates a new GetInstallationParams object
// with the ability to set a context for a request.
func NewGetInstallationParamsWithContext(ctx context.Context) *GetInstallationParams {
	return &GetInstallationParams{
		Context: ctx,
	}
}

// NewGetInstallationParamsWithHTTPClient creates a new GetInstallationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInstallationParamsWithHTTPClient(client *http.Client) *GetInstallationParams {
	return &GetInstallationParams{
		HTTPClient: client,
	}
}

/*
GetInstallationParams contains all the parameters to send to the API endpoint

	for the get installation operation.

	Typically these are written to a http.Request.
*/
type GetInstallationParams struct {

	/* Auth.

	   auth.
	*/
	Auth *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get installation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInstallationParams) WithDefaults() *GetInstallationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get installation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInstallationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get installation params
func (o *GetInstallationParams) WithTimeout(timeout time.Duration) *GetInstallationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get installation params
func (o *GetInstallationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get installation params
func (o *GetInstallationParams) WithContext(ctx context.Context) *GetInstallationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get installation params
func (o *GetInstallationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get installation params
func (o *GetInstallationParams) WithHTTPClient(client *http.Client) *GetInstallationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get installation params
func (o *GetInstallationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuth adds the auth to the get installation params
func (o *GetInstallationParams) WithAuth(auth *bool) *GetInstallationParams {
	o.SetAuth(auth)
	return o
}

// SetAuth adds the auth to the get installation params
func (o *GetInstallationParams) SetAuth(auth *bool) {
	o.Auth = auth
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstallationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Auth != nil {

		// query param auth
		var qrAuth bool

		if o.Auth != nil {
			qrAuth = *o.Auth
		}
		qAuth := swag.FormatBool(qrAuth)
		if qAuth != "" {

			if err := r.SetQueryParam("auth", qAuth); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
