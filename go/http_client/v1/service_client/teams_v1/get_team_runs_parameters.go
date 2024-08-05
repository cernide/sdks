// Code generated by go-swagger; DO NOT EDIT.

package teams_v1

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

// NewGetTeamRunsParams creates a new GetTeamRunsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTeamRunsParams() *GetTeamRunsParams {
	return &GetTeamRunsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTeamRunsParamsWithTimeout creates a new GetTeamRunsParams object
// with the ability to set a timeout on a request.
func NewGetTeamRunsParamsWithTimeout(timeout time.Duration) *GetTeamRunsParams {
	return &GetTeamRunsParams{
		timeout: timeout,
	}
}

// NewGetTeamRunsParamsWithContext creates a new GetTeamRunsParams object
// with the ability to set a context for a request.
func NewGetTeamRunsParamsWithContext(ctx context.Context) *GetTeamRunsParams {
	return &GetTeamRunsParams{
		Context: ctx,
	}
}

// NewGetTeamRunsParamsWithHTTPClient creates a new GetTeamRunsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTeamRunsParamsWithHTTPClient(client *http.Client) *GetTeamRunsParams {
	return &GetTeamRunsParams{
		HTTPClient: client,
	}
}

/*
GetTeamRunsParams contains all the parameters to send to the API endpoint

	for the get team runs operation.

	Typically these are written to a http.Request.
*/
type GetTeamRunsParams struct {

	/* Bookmarks.

	   Filter by bookmarks.
	*/
	Bookmarks *bool

	/* Limit.

	   Limit size.

	   Format: int32
	*/
	Limit *int32

	/* Mode.

	   Mode of the search.
	*/
	Mode *string

	/* Name.

	   Entity managing the resource
	*/
	Name string

	/* NoPage.

	   No pagination.
	*/
	NoPage *bool

	/* Offset.

	   Pagination offset.

	   Format: int32
	*/
	Offset *int32

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* Query.

	   Query filter the search.
	*/
	Query *string

	/* Sort.

	   Sort to order the search.
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get team runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTeamRunsParams) WithDefaults() *GetTeamRunsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get team runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTeamRunsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get team runs params
func (o *GetTeamRunsParams) WithTimeout(timeout time.Duration) *GetTeamRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get team runs params
func (o *GetTeamRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get team runs params
func (o *GetTeamRunsParams) WithContext(ctx context.Context) *GetTeamRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get team runs params
func (o *GetTeamRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get team runs params
func (o *GetTeamRunsParams) WithHTTPClient(client *http.Client) *GetTeamRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get team runs params
func (o *GetTeamRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBookmarks adds the bookmarks to the get team runs params
func (o *GetTeamRunsParams) WithBookmarks(bookmarks *bool) *GetTeamRunsParams {
	o.SetBookmarks(bookmarks)
	return o
}

// SetBookmarks adds the bookmarks to the get team runs params
func (o *GetTeamRunsParams) SetBookmarks(bookmarks *bool) {
	o.Bookmarks = bookmarks
}

// WithLimit adds the limit to the get team runs params
func (o *GetTeamRunsParams) WithLimit(limit *int32) *GetTeamRunsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get team runs params
func (o *GetTeamRunsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithMode adds the mode to the get team runs params
func (o *GetTeamRunsParams) WithMode(mode *string) *GetTeamRunsParams {
	o.SetMode(mode)
	return o
}

// SetMode adds the mode to the get team runs params
func (o *GetTeamRunsParams) SetMode(mode *string) {
	o.Mode = mode
}

// WithName adds the name to the get team runs params
func (o *GetTeamRunsParams) WithName(name string) *GetTeamRunsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get team runs params
func (o *GetTeamRunsParams) SetName(name string) {
	o.Name = name
}

// WithNoPage adds the noPage to the get team runs params
func (o *GetTeamRunsParams) WithNoPage(noPage *bool) *GetTeamRunsParams {
	o.SetNoPage(noPage)
	return o
}

// SetNoPage adds the noPage to the get team runs params
func (o *GetTeamRunsParams) SetNoPage(noPage *bool) {
	o.NoPage = noPage
}

// WithOffset adds the offset to the get team runs params
func (o *GetTeamRunsParams) WithOffset(offset *int32) *GetTeamRunsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get team runs params
func (o *GetTeamRunsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOwner adds the owner to the get team runs params
func (o *GetTeamRunsParams) WithOwner(owner string) *GetTeamRunsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get team runs params
func (o *GetTeamRunsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithQuery adds the query to the get team runs params
func (o *GetTeamRunsParams) WithQuery(query *string) *GetTeamRunsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get team runs params
func (o *GetTeamRunsParams) SetQuery(query *string) {
	o.Query = query
}

// WithSort adds the sort to the get team runs params
func (o *GetTeamRunsParams) WithSort(sort *string) *GetTeamRunsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get team runs params
func (o *GetTeamRunsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetTeamRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Bookmarks != nil {

		// query param bookmarks
		var qrBookmarks bool

		if o.Bookmarks != nil {
			qrBookmarks = *o.Bookmarks
		}
		qBookmarks := swag.FormatBool(qrBookmarks)
		if qBookmarks != "" {

			if err := r.SetQueryParam("bookmarks", qBookmarks); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Mode != nil {

		// query param mode
		var qrMode string

		if o.Mode != nil {
			qrMode = *o.Mode
		}
		qMode := qrMode
		if qMode != "" {

			if err := r.SetQueryParam("mode", qMode); err != nil {
				return err
			}
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.NoPage != nil {

		// query param no_page
		var qrNoPage bool

		if o.NoPage != nil {
			qrNoPage = *o.NoPage
		}
		qNoPage := swag.FormatBool(qrNoPage)
		if qNoPage != "" {

			if err := r.SetQueryParam("no_page", qNoPage); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}