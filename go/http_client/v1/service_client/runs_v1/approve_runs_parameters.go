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

// NewApproveRunsParams creates a new ApproveRunsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApproveRunsParams() *ApproveRunsParams {
	return &ApproveRunsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApproveRunsParamsWithTimeout creates a new ApproveRunsParams object
// with the ability to set a timeout on a request.
func NewApproveRunsParamsWithTimeout(timeout time.Duration) *ApproveRunsParams {
	return &ApproveRunsParams{
		timeout: timeout,
	}
}

// NewApproveRunsParamsWithContext creates a new ApproveRunsParams object
// with the ability to set a context for a request.
func NewApproveRunsParamsWithContext(ctx context.Context) *ApproveRunsParams {
	return &ApproveRunsParams{
		Context: ctx,
	}
}

// NewApproveRunsParamsWithHTTPClient creates a new ApproveRunsParams object
// with the ability to set a custom HTTPClient for a request.
func NewApproveRunsParamsWithHTTPClient(client *http.Client) *ApproveRunsParams {
	return &ApproveRunsParams{
		HTTPClient: client,
	}
}

/*
ApproveRunsParams contains all the parameters to send to the API endpoint

	for the approve runs operation.

	Typically these are written to a http.Request.
*/
type ApproveRunsParams struct {

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

// WithDefaults hydrates default values in the approve runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApproveRunsParams) WithDefaults() *ApproveRunsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the approve runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApproveRunsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the approve runs params
func (o *ApproveRunsParams) WithTimeout(timeout time.Duration) *ApproveRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the approve runs params
func (o *ApproveRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the approve runs params
func (o *ApproveRunsParams) WithContext(ctx context.Context) *ApproveRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the approve runs params
func (o *ApproveRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the approve runs params
func (o *ApproveRunsParams) WithHTTPClient(client *http.Client) *ApproveRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the approve runs params
func (o *ApproveRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the approve runs params
func (o *ApproveRunsParams) WithBody(body *service_model.V1Uuids) *ApproveRunsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the approve runs params
func (o *ApproveRunsParams) SetBody(body *service_model.V1Uuids) {
	o.Body = body
}

// WithName adds the name to the approve runs params
func (o *ApproveRunsParams) WithName(name string) *ApproveRunsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the approve runs params
func (o *ApproveRunsParams) SetName(name string) {
	o.Name = name
}

// WithOwner adds the owner to the approve runs params
func (o *ApproveRunsParams) WithOwner(owner string) *ApproveRunsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the approve runs params
func (o *ApproveRunsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *ApproveRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
