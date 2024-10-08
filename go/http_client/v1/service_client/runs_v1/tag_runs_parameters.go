// Code generated by go-swagger; DO NOT EDIT.

package runs_v1

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

// NewTagRunsParams creates a new TagRunsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTagRunsParams() *TagRunsParams {
	return &TagRunsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTagRunsParamsWithTimeout creates a new TagRunsParams object
// with the ability to set a timeout on a request.
func NewTagRunsParamsWithTimeout(timeout time.Duration) *TagRunsParams {
	return &TagRunsParams{
		timeout: timeout,
	}
}

// NewTagRunsParamsWithContext creates a new TagRunsParams object
// with the ability to set a context for a request.
func NewTagRunsParamsWithContext(ctx context.Context) *TagRunsParams {
	return &TagRunsParams{
		Context: ctx,
	}
}

// NewTagRunsParamsWithHTTPClient creates a new TagRunsParams object
// with the ability to set a custom HTTPClient for a request.
func NewTagRunsParamsWithHTTPClient(client *http.Client) *TagRunsParams {
	return &TagRunsParams{
		HTTPClient: client,
	}
}

/*
TagRunsParams contains all the parameters to send to the API endpoint

	for the tag runs operation.

	Typically these are written to a http.Request.
*/
type TagRunsParams struct {

	/* Body.

	   Data
	*/
	Body *service_model.V1EntitiesTags

	/* Name.

	   Entity under namespace
	*/
	Name string

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tag runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TagRunsParams) WithDefaults() *TagRunsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tag runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TagRunsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tag runs params
func (o *TagRunsParams) WithTimeout(timeout time.Duration) *TagRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tag runs params
func (o *TagRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tag runs params
func (o *TagRunsParams) WithContext(ctx context.Context) *TagRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tag runs params
func (o *TagRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tag runs params
func (o *TagRunsParams) WithHTTPClient(client *http.Client) *TagRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tag runs params
func (o *TagRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the tag runs params
func (o *TagRunsParams) WithBody(body *service_model.V1EntitiesTags) *TagRunsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the tag runs params
func (o *TagRunsParams) SetBody(body *service_model.V1EntitiesTags) {
	o.Body = body
}

// WithName adds the name to the tag runs params
func (o *TagRunsParams) WithName(name string) *TagRunsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the tag runs params
func (o *TagRunsParams) SetName(name string) {
	o.Name = name
}

// WithOwner adds the owner to the tag runs params
func (o *TagRunsParams) WithOwner(owner string) *TagRunsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the tag runs params
func (o *TagRunsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *TagRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
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
