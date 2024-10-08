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

// NewTransferRunParams creates a new TransferRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTransferRunParams() *TransferRunParams {
	return &TransferRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTransferRunParamsWithTimeout creates a new TransferRunParams object
// with the ability to set a timeout on a request.
func NewTransferRunParamsWithTimeout(timeout time.Duration) *TransferRunParams {
	return &TransferRunParams{
		timeout: timeout,
	}
}

// NewTransferRunParamsWithContext creates a new TransferRunParams object
// with the ability to set a context for a request.
func NewTransferRunParamsWithContext(ctx context.Context) *TransferRunParams {
	return &TransferRunParams{
		Context: ctx,
	}
}

// NewTransferRunParamsWithHTTPClient creates a new TransferRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewTransferRunParamsWithHTTPClient(client *http.Client) *TransferRunParams {
	return &TransferRunParams{
		HTTPClient: client,
	}
}

/*
TransferRunParams contains all the parameters to send to the API endpoint

	for the transfer run operation.

	Typically these are written to a http.Request.
*/
type TransferRunParams struct {

	/* Body.

	   Run object
	*/
	Body *service_model.V1Run

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* Project.

	   Project where the run will be assigned
	*/
	Project string

	/* RunUUID.

	   UUID
	*/
	RunUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the transfer run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TransferRunParams) WithDefaults() *TransferRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the transfer run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TransferRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the transfer run params
func (o *TransferRunParams) WithTimeout(timeout time.Duration) *TransferRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the transfer run params
func (o *TransferRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the transfer run params
func (o *TransferRunParams) WithContext(ctx context.Context) *TransferRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the transfer run params
func (o *TransferRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the transfer run params
func (o *TransferRunParams) WithHTTPClient(client *http.Client) *TransferRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the transfer run params
func (o *TransferRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the transfer run params
func (o *TransferRunParams) WithBody(body *service_model.V1Run) *TransferRunParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the transfer run params
func (o *TransferRunParams) SetBody(body *service_model.V1Run) {
	o.Body = body
}

// WithOwner adds the owner to the transfer run params
func (o *TransferRunParams) WithOwner(owner string) *TransferRunParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the transfer run params
func (o *TransferRunParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the transfer run params
func (o *TransferRunParams) WithProject(project string) *TransferRunParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the transfer run params
func (o *TransferRunParams) SetProject(project string) {
	o.Project = project
}

// WithRunUUID adds the runUUID to the transfer run params
func (o *TransferRunParams) WithRunUUID(runUUID string) *TransferRunParams {
	o.SetRunUUID(runUUID)
	return o
}

// SetRunUUID adds the runUuid to the transfer run params
func (o *TransferRunParams) SetRunUUID(runUUID string) {
	o.RunUUID = runUUID
}

// WriteToRequest writes these params to a swagger request
func (o *TransferRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// path param run.uuid
	if err := r.SetPathParam("run.uuid", o.RunUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
