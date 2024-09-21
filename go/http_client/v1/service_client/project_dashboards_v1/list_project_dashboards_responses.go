// Code generated by go-swagger; DO NOT EDIT.

package project_dashboards_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// ListProjectDashboardsReader is a Reader for the ListProjectDashboards structure.
type ListProjectDashboardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectDashboardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectDashboardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListProjectDashboardsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListProjectDashboardsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListProjectDashboardsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListProjectDashboardsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectDashboardsOK creates a ListProjectDashboardsOK with default headers values
func NewListProjectDashboardsOK() *ListProjectDashboardsOK {
	return &ListProjectDashboardsOK{}
}

/*
ListProjectDashboardsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListProjectDashboardsOK struct {
	Payload *service_model.V1ListDashboardsResponse
}

// IsSuccess returns true when this list project dashboards o k response has a 2xx status code
func (o *ListProjectDashboardsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project dashboards o k response has a 3xx status code
func (o *ListProjectDashboardsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project dashboards o k response has a 4xx status code
func (o *ListProjectDashboardsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project dashboards o k response has a 5xx status code
func (o *ListProjectDashboardsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project dashboards o k response a status code equal to that given
func (o *ListProjectDashboardsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list project dashboards o k response
func (o *ListProjectDashboardsOK) Code() int {
	return 200
}

func (o *ListProjectDashboardsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards][%d] listProjectDashboardsOK  %+v", 200, o.Payload)
}

func (o *ListProjectDashboardsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards][%d] listProjectDashboardsOK  %+v", 200, o.Payload)
}

func (o *ListProjectDashboardsOK) GetPayload() *service_model.V1ListDashboardsResponse {
	return o.Payload
}

func (o *ListProjectDashboardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListDashboardsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectDashboardsNoContent creates a ListProjectDashboardsNoContent with default headers values
func NewListProjectDashboardsNoContent() *ListProjectDashboardsNoContent {
	return &ListProjectDashboardsNoContent{}
}

/*
ListProjectDashboardsNoContent describes a response with status code 204, with default header values.

No content.
*/
type ListProjectDashboardsNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this list project dashboards no content response has a 2xx status code
func (o *ListProjectDashboardsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project dashboards no content response has a 3xx status code
func (o *ListProjectDashboardsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project dashboards no content response has a 4xx status code
func (o *ListProjectDashboardsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project dashboards no content response has a 5xx status code
func (o *ListProjectDashboardsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this list project dashboards no content response a status code equal to that given
func (o *ListProjectDashboardsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the list project dashboards no content response
func (o *ListProjectDashboardsNoContent) Code() int {
	return 204
}

func (o *ListProjectDashboardsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards][%d] listProjectDashboardsNoContent  %+v", 204, o.Payload)
}

func (o *ListProjectDashboardsNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards][%d] listProjectDashboardsNoContent  %+v", 204, o.Payload)
}

func (o *ListProjectDashboardsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListProjectDashboardsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectDashboardsForbidden creates a ListProjectDashboardsForbidden with default headers values
func NewListProjectDashboardsForbidden() *ListProjectDashboardsForbidden {
	return &ListProjectDashboardsForbidden{}
}

/*
ListProjectDashboardsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ListProjectDashboardsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this list project dashboards forbidden response has a 2xx status code
func (o *ListProjectDashboardsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project dashboards forbidden response has a 3xx status code
func (o *ListProjectDashboardsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project dashboards forbidden response has a 4xx status code
func (o *ListProjectDashboardsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project dashboards forbidden response has a 5xx status code
func (o *ListProjectDashboardsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list project dashboards forbidden response a status code equal to that given
func (o *ListProjectDashboardsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list project dashboards forbidden response
func (o *ListProjectDashboardsForbidden) Code() int {
	return 403
}

func (o *ListProjectDashboardsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards][%d] listProjectDashboardsForbidden  %+v", 403, o.Payload)
}

func (o *ListProjectDashboardsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards][%d] listProjectDashboardsForbidden  %+v", 403, o.Payload)
}

func (o *ListProjectDashboardsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListProjectDashboardsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectDashboardsNotFound creates a ListProjectDashboardsNotFound with default headers values
func NewListProjectDashboardsNotFound() *ListProjectDashboardsNotFound {
	return &ListProjectDashboardsNotFound{}
}

/*
ListProjectDashboardsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ListProjectDashboardsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this list project dashboards not found response has a 2xx status code
func (o *ListProjectDashboardsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project dashboards not found response has a 3xx status code
func (o *ListProjectDashboardsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project dashboards not found response has a 4xx status code
func (o *ListProjectDashboardsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project dashboards not found response has a 5xx status code
func (o *ListProjectDashboardsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list project dashboards not found response a status code equal to that given
func (o *ListProjectDashboardsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list project dashboards not found response
func (o *ListProjectDashboardsNotFound) Code() int {
	return 404
}

func (o *ListProjectDashboardsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards][%d] listProjectDashboardsNotFound  %+v", 404, o.Payload)
}

func (o *ListProjectDashboardsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards][%d] listProjectDashboardsNotFound  %+v", 404, o.Payload)
}

func (o *ListProjectDashboardsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListProjectDashboardsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectDashboardsDefault creates a ListProjectDashboardsDefault with default headers values
func NewListProjectDashboardsDefault(code int) *ListProjectDashboardsDefault {
	return &ListProjectDashboardsDefault{
		_statusCode: code,
	}
}

/*
ListProjectDashboardsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListProjectDashboardsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this list project dashboards default response has a 2xx status code
func (o *ListProjectDashboardsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project dashboards default response has a 3xx status code
func (o *ListProjectDashboardsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project dashboards default response has a 4xx status code
func (o *ListProjectDashboardsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project dashboards default response has a 5xx status code
func (o *ListProjectDashboardsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project dashboards default response a status code equal to that given
func (o *ListProjectDashboardsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list project dashboards default response
func (o *ListProjectDashboardsDefault) Code() int {
	return o._statusCode
}

func (o *ListProjectDashboardsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards][%d] ListProjectDashboards default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectDashboardsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards][%d] ListProjectDashboards default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectDashboardsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListProjectDashboardsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
