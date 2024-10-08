// Code generated by go-swagger; DO NOT EDIT.

package tags_v1

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

// NewUpdateTagParams creates a new UpdateTagParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateTagParams() *UpdateTagParams {
	return &UpdateTagParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTagParamsWithTimeout creates a new UpdateTagParams object
// with the ability to set a timeout on a request.
func NewUpdateTagParamsWithTimeout(timeout time.Duration) *UpdateTagParams {
	return &UpdateTagParams{
		timeout: timeout,
	}
}

// NewUpdateTagParamsWithContext creates a new UpdateTagParams object
// with the ability to set a context for a request.
func NewUpdateTagParamsWithContext(ctx context.Context) *UpdateTagParams {
	return &UpdateTagParams{
		Context: ctx,
	}
}

// NewUpdateTagParamsWithHTTPClient creates a new UpdateTagParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateTagParamsWithHTTPClient(client *http.Client) *UpdateTagParams {
	return &UpdateTagParams{
		HTTPClient: client,
	}
}

/*
UpdateTagParams contains all the parameters to send to the API endpoint

	for the update tag operation.

	Typically these are written to a http.Request.
*/
type UpdateTagParams struct {

	/* Body.

	   Tag body
	*/
	Body *service_model.V1Tag

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* TagUUID.

	   UUID
	*/
	TagUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTagParams) WithDefaults() *UpdateTagParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTagParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update tag params
func (o *UpdateTagParams) WithTimeout(timeout time.Duration) *UpdateTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update tag params
func (o *UpdateTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update tag params
func (o *UpdateTagParams) WithContext(ctx context.Context) *UpdateTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update tag params
func (o *UpdateTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update tag params
func (o *UpdateTagParams) WithHTTPClient(client *http.Client) *UpdateTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update tag params
func (o *UpdateTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update tag params
func (o *UpdateTagParams) WithBody(body *service_model.V1Tag) *UpdateTagParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update tag params
func (o *UpdateTagParams) SetBody(body *service_model.V1Tag) {
	o.Body = body
}

// WithOwner adds the owner to the update tag params
func (o *UpdateTagParams) WithOwner(owner string) *UpdateTagParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the update tag params
func (o *UpdateTagParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithTagUUID adds the tagUUID to the update tag params
func (o *UpdateTagParams) WithTagUUID(tagUUID string) *UpdateTagParams {
	o.SetTagUUID(tagUUID)
	return o
}

// SetTagUUID adds the tagUuid to the update tag params
func (o *UpdateTagParams) SetTagUUID(tagUUID string) {
	o.TagUUID = tagUUID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param tag.uuid
	if err := r.SetPathParam("tag.uuid", o.TagUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
