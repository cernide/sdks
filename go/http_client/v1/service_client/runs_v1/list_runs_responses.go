// Code generated by go-swagger; DO NOT EDIT.

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// ListRunsReader is a Reader for the ListRuns structure.
type ListRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListRunsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListRunsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListRunsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListRunsOK creates a ListRunsOK with default headers values
func NewListRunsOK() *ListRunsOK {
	return &ListRunsOK{}
}

/*
ListRunsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListRunsOK struct {
	Payload *service_model.V1ListRunsResponse
}

// IsSuccess returns true when this list runs o k response has a 2xx status code
func (o *ListRunsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list runs o k response has a 3xx status code
func (o *ListRunsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list runs o k response has a 4xx status code
func (o *ListRunsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list runs o k response has a 5xx status code
func (o *ListRunsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list runs o k response a status code equal to that given
func (o *ListRunsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list runs o k response
func (o *ListRunsOK) Code() int {
	return 200
}

func (o *ListRunsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/runs][%d] listRunsOK  %+v", 200, o.Payload)
}

func (o *ListRunsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/runs][%d] listRunsOK  %+v", 200, o.Payload)
}

func (o *ListRunsOK) GetPayload() *service_model.V1ListRunsResponse {
	return o.Payload
}

func (o *ListRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListRunsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRunsNoContent creates a ListRunsNoContent with default headers values
func NewListRunsNoContent() *ListRunsNoContent {
	return &ListRunsNoContent{}
}

/*
ListRunsNoContent describes a response with status code 204, with default header values.

No content.
*/
type ListRunsNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this list runs no content response has a 2xx status code
func (o *ListRunsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list runs no content response has a 3xx status code
func (o *ListRunsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list runs no content response has a 4xx status code
func (o *ListRunsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this list runs no content response has a 5xx status code
func (o *ListRunsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this list runs no content response a status code equal to that given
func (o *ListRunsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the list runs no content response
func (o *ListRunsNoContent) Code() int {
	return 204
}

func (o *ListRunsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/runs][%d] listRunsNoContent  %+v", 204, o.Payload)
}

func (o *ListRunsNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/runs][%d] listRunsNoContent  %+v", 204, o.Payload)
}

func (o *ListRunsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListRunsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRunsForbidden creates a ListRunsForbidden with default headers values
func NewListRunsForbidden() *ListRunsForbidden {
	return &ListRunsForbidden{}
}

/*
ListRunsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ListRunsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this list runs forbidden response has a 2xx status code
func (o *ListRunsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list runs forbidden response has a 3xx status code
func (o *ListRunsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list runs forbidden response has a 4xx status code
func (o *ListRunsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list runs forbidden response has a 5xx status code
func (o *ListRunsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list runs forbidden response a status code equal to that given
func (o *ListRunsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list runs forbidden response
func (o *ListRunsForbidden) Code() int {
	return 403
}

func (o *ListRunsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/runs][%d] listRunsForbidden  %+v", 403, o.Payload)
}

func (o *ListRunsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/runs][%d] listRunsForbidden  %+v", 403, o.Payload)
}

func (o *ListRunsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListRunsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRunsNotFound creates a ListRunsNotFound with default headers values
func NewListRunsNotFound() *ListRunsNotFound {
	return &ListRunsNotFound{}
}

/*
ListRunsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ListRunsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this list runs not found response has a 2xx status code
func (o *ListRunsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list runs not found response has a 3xx status code
func (o *ListRunsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list runs not found response has a 4xx status code
func (o *ListRunsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list runs not found response has a 5xx status code
func (o *ListRunsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list runs not found response a status code equal to that given
func (o *ListRunsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list runs not found response
func (o *ListRunsNotFound) Code() int {
	return 404
}

func (o *ListRunsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/runs][%d] listRunsNotFound  %+v", 404, o.Payload)
}

func (o *ListRunsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/runs][%d] listRunsNotFound  %+v", 404, o.Payload)
}

func (o *ListRunsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListRunsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRunsDefault creates a ListRunsDefault with default headers values
func NewListRunsDefault(code int) *ListRunsDefault {
	return &ListRunsDefault{
		_statusCode: code,
	}
}

/*
ListRunsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListRunsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this list runs default response has a 2xx status code
func (o *ListRunsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list runs default response has a 3xx status code
func (o *ListRunsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list runs default response has a 4xx status code
func (o *ListRunsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list runs default response has a 5xx status code
func (o *ListRunsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list runs default response a status code equal to that given
func (o *ListRunsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list runs default response
func (o *ListRunsDefault) Code() int {
	return o._statusCode
}

func (o *ListRunsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/runs][%d] ListRuns default  %+v", o._statusCode, o.Payload)
}

func (o *ListRunsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/runs][%d] ListRuns default  %+v", o._statusCode, o.Payload)
}

func (o *ListRunsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
