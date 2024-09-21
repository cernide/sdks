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

// NewDeleteRunsParams creates a new DeleteRunsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRunsParams() *DeleteRunsParams {
	return &DeleteRunsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRunsParamsWithTimeout creates a new DeleteRunsParams object
// with the ability to set a timeout on a request.
func NewDeleteRunsParamsWithTimeout(timeout time.Duration) *DeleteRunsParams {
	return &DeleteRunsParams{
		timeout: timeout,
	}
}

// NewDeleteRunsParamsWithContext creates a new DeleteRunsParams object
// with the ability to set a context for a request.
func NewDeleteRunsParamsWithContext(ctx context.Context) *DeleteRunsParams {
	return &DeleteRunsParams{
		Context: ctx,
	}
}

// NewDeleteRunsParamsWithHTTPClient creates a new DeleteRunsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRunsParamsWithHTTPClient(client *http.Client) *DeleteRunsParams {
	return &DeleteRunsParams{
		HTTPClient: client,
	}
}

/*
DeleteRunsParams contains all the parameters to send to the API endpoint

	for the delete runs operation.

	Typically these are written to a http.Request.
*/
type DeleteRunsParams struct {

	/* Body.

	   Uuids of the entities
	*/
	Body *service_model.V1Uuids

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

// WithDefaults hydrates default values in the delete runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRunsParams) WithDefaults() *DeleteRunsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRunsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete runs params
func (o *DeleteRunsParams) WithTimeout(timeout time.Duration) *DeleteRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete runs params
func (o *DeleteRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete runs params
func (o *DeleteRunsParams) WithContext(ctx context.Context) *DeleteRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete runs params
func (o *DeleteRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete runs params
func (o *DeleteRunsParams) WithHTTPClient(client *http.Client) *DeleteRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete runs params
func (o *DeleteRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete runs params
func (o *DeleteRunsParams) WithBody(body *service_model.V1Uuids) *DeleteRunsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete runs params
func (o *DeleteRunsParams) SetBody(body *service_model.V1Uuids) {
	o.Body = body
}

// WithName adds the name to the delete runs params
func (o *DeleteRunsParams) WithName(name string) *DeleteRunsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete runs params
func (o *DeleteRunsParams) SetName(name string) {
	o.Name = name
}

// WithOwner adds the owner to the delete runs params
func (o *DeleteRunsParams) WithOwner(owner string) *DeleteRunsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the delete runs params
func (o *DeleteRunsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
