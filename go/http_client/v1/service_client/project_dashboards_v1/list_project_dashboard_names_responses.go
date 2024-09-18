// Code generated by go-swagger; DO NOT EDIT.

package project_dashboards_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/go/http_client/v1/service_model"
)

// ListProjectDashboardNamesReader is a Reader for the ListProjectDashboardNames structure.
type ListProjectDashboardNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectDashboardNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectDashboardNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListProjectDashboardNamesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListProjectDashboardNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListProjectDashboardNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListProjectDashboardNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectDashboardNamesOK creates a ListProjectDashboardNamesOK with default headers values
func NewListProjectDashboardNamesOK() *ListProjectDashboardNamesOK {
	return &ListProjectDashboardNamesOK{}
}

/*
ListProjectDashboardNamesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListProjectDashboardNamesOK struct {
	Payload *service_model.V1ListDashboardsResponse
}

// IsSuccess returns true when this list project dashboard names o k response has a 2xx status code
func (o *ListProjectDashboardNamesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project dashboard names o k response has a 3xx status code
func (o *ListProjectDashboardNamesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project dashboard names o k response has a 4xx status code
func (o *ListProjectDashboardNamesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project dashboard names o k response has a 5xx status code
func (o *ListProjectDashboardNamesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project dashboard names o k response a status code equal to that given
func (o *ListProjectDashboardNamesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list project dashboard names o k response
func (o *ListProjectDashboardNamesOK) Code() int {
	return 200
}

func (o *ListProjectDashboardNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards/names][%d] listProjectDashboardNamesOK  %+v", 200, o.Payload)
}

func (o *ListProjectDashboardNamesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards/names][%d] listProjectDashboardNamesOK  %+v", 200, o.Payload)
}

func (o *ListProjectDashboardNamesOK) GetPayload() *service_model.V1ListDashboardsResponse {
	return o.Payload
}

func (o *ListProjectDashboardNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListDashboardsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectDashboardNamesNoContent creates a ListProjectDashboardNamesNoContent with default headers values
func NewListProjectDashboardNamesNoContent() *ListProjectDashboardNamesNoContent {
	return &ListProjectDashboardNamesNoContent{}
}

/*
ListProjectDashboardNamesNoContent describes a response with status code 204, with default header values.

No content.
*/
type ListProjectDashboardNamesNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this list project dashboard names no content response has a 2xx status code
func (o *ListProjectDashboardNamesNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project dashboard names no content response has a 3xx status code
func (o *ListProjectDashboardNamesNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project dashboard names no content response has a 4xx status code
func (o *ListProjectDashboardNamesNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project dashboard names no content response has a 5xx status code
func (o *ListProjectDashboardNamesNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this list project dashboard names no content response a status code equal to that given
func (o *ListProjectDashboardNamesNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the list project dashboard names no content response
func (o *ListProjectDashboardNamesNoContent) Code() int {
	return 204
}

func (o *ListProjectDashboardNamesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards/names][%d] listProjectDashboardNamesNoContent  %+v", 204, o.Payload)
}

func (o *ListProjectDashboardNamesNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards/names][%d] listProjectDashboardNamesNoContent  %+v", 204, o.Payload)
}

func (o *ListProjectDashboardNamesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListProjectDashboardNamesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectDashboardNamesForbidden creates a ListProjectDashboardNamesForbidden with default headers values
func NewListProjectDashboardNamesForbidden() *ListProjectDashboardNamesForbidden {
	return &ListProjectDashboardNamesForbidden{}
}

/*
ListProjectDashboardNamesForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ListProjectDashboardNamesForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this list project dashboard names forbidden response has a 2xx status code
func (o *ListProjectDashboardNamesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project dashboard names forbidden response has a 3xx status code
func (o *ListProjectDashboardNamesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project dashboard names forbidden response has a 4xx status code
func (o *ListProjectDashboardNamesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project dashboard names forbidden response has a 5xx status code
func (o *ListProjectDashboardNamesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list project dashboard names forbidden response a status code equal to that given
func (o *ListProjectDashboardNamesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list project dashboard names forbidden response
func (o *ListProjectDashboardNamesForbidden) Code() int {
	return 403
}

func (o *ListProjectDashboardNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards/names][%d] listProjectDashboardNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListProjectDashboardNamesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards/names][%d] listProjectDashboardNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListProjectDashboardNamesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListProjectDashboardNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectDashboardNamesNotFound creates a ListProjectDashboardNamesNotFound with default headers values
func NewListProjectDashboardNamesNotFound() *ListProjectDashboardNamesNotFound {
	return &ListProjectDashboardNamesNotFound{}
}

/*
ListProjectDashboardNamesNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ListProjectDashboardNamesNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this list project dashboard names not found response has a 2xx status code
func (o *ListProjectDashboardNamesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project dashboard names not found response has a 3xx status code
func (o *ListProjectDashboardNamesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project dashboard names not found response has a 4xx status code
func (o *ListProjectDashboardNamesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project dashboard names not found response has a 5xx status code
func (o *ListProjectDashboardNamesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list project dashboard names not found response a status code equal to that given
func (o *ListProjectDashboardNamesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list project dashboard names not found response
func (o *ListProjectDashboardNamesNotFound) Code() int {
	return 404
}

func (o *ListProjectDashboardNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards/names][%d] listProjectDashboardNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListProjectDashboardNamesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards/names][%d] listProjectDashboardNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListProjectDashboardNamesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListProjectDashboardNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectDashboardNamesDefault creates a ListProjectDashboardNamesDefault with default headers values
func NewListProjectDashboardNamesDefault(code int) *ListProjectDashboardNamesDefault {
	return &ListProjectDashboardNamesDefault{
		_statusCode: code,
	}
}

/*
ListProjectDashboardNamesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListProjectDashboardNamesDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this list project dashboard names default response has a 2xx status code
func (o *ListProjectDashboardNamesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project dashboard names default response has a 3xx status code
func (o *ListProjectDashboardNamesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project dashboard names default response has a 4xx status code
func (o *ListProjectDashboardNamesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project dashboard names default response has a 5xx status code
func (o *ListProjectDashboardNamesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project dashboard names default response a status code equal to that given
func (o *ListProjectDashboardNamesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list project dashboard names default response
func (o *ListProjectDashboardNamesDefault) Code() int {
	return o._statusCode
}

func (o *ListProjectDashboardNamesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards/names][%d] ListProjectDashboardNames default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectDashboardNamesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/dashboards/names][%d] ListProjectDashboardNames default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectDashboardNamesDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListProjectDashboardNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
