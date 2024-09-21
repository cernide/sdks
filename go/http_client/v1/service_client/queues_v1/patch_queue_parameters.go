// Code generated by go-swagger; DO NOT EDIT.

package queues_v1

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

// NewPatchQueueParams creates a new PatchQueueParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchQueueParams() *PatchQueueParams {
	return &PatchQueueParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchQueueParamsWithTimeout creates a new PatchQueueParams object
// with the ability to set a timeout on a request.
func NewPatchQueueParamsWithTimeout(timeout time.Duration) *PatchQueueParams {
	return &PatchQueueParams{
		timeout: timeout,
	}
}

// NewPatchQueueParamsWithContext creates a new PatchQueueParams object
// with the ability to set a context for a request.
func NewPatchQueueParamsWithContext(ctx context.Context) *PatchQueueParams {
	return &PatchQueueParams{
		Context: ctx,
	}
}

// NewPatchQueueParamsWithHTTPClient creates a new PatchQueueParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchQueueParamsWithHTTPClient(client *http.Client) *PatchQueueParams {
	return &PatchQueueParams{
		HTTPClient: client,
	}
}

/*
PatchQueueParams contains all the parameters to send to the API endpoint

	for the patch queue operation.

	Typically these are written to a http.Request.
*/
type PatchQueueParams struct {

	/* Agent.

	   Agent that consumes the queue
	*/
	Agent string

	/* Body.

	   Queue body
	*/
	Body *service_model.V1Queue

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* QueueUUID.

	   UUID
	*/
	QueueUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch queue params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchQueueParams) WithDefaults() *PatchQueueParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch queue params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchQueueParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch queue params
func (o *PatchQueueParams) WithTimeout(timeout time.Duration) *PatchQueueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch queue params
func (o *PatchQueueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch queue params
func (o *PatchQueueParams) WithContext(ctx context.Context) *PatchQueueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch queue params
func (o *PatchQueueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch queue params
func (o *PatchQueueParams) WithHTTPClient(client *http.Client) *PatchQueueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch queue params
func (o *PatchQueueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgent adds the agent to the patch queue params
func (o *PatchQueueParams) WithAgent(agent string) *PatchQueueParams {
	o.SetAgent(agent)
	return o
}

// SetAgent adds the agent to the patch queue params
func (o *PatchQueueParams) SetAgent(agent string) {
	o.Agent = agent
}

// WithBody adds the body to the patch queue params
func (o *PatchQueueParams) WithBody(body *service_model.V1Queue) *PatchQueueParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch queue params
func (o *PatchQueueParams) SetBody(body *service_model.V1Queue) {
	o.Body = body
}

// WithOwner adds the owner to the patch queue params
func (o *PatchQueueParams) WithOwner(owner string) *PatchQueueParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the patch queue params
func (o *PatchQueueParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithQueueUUID adds the queueUUID to the patch queue params
func (o *PatchQueueParams) WithQueueUUID(queueUUID string) *PatchQueueParams {
	o.SetQueueUUID(queueUUID)
	return o
}

// SetQueueUUID adds the queueUuid to the patch queue params
func (o *PatchQueueParams) SetQueueUUID(queueUUID string) {
	o.QueueUUID = queueUUID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchQueueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agent
	if err := r.SetPathParam("agent", o.Agent); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param queue.uuid
	if err := r.SetPathParam("queue.uuid", o.QueueUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
