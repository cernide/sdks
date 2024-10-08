// Code generated by go-swagger; DO NOT EDIT.

package projects_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// ListVersionNamesReader is a Reader for the ListVersionNames structure.
type ListVersionNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVersionNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVersionNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListVersionNamesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListVersionNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListVersionNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListVersionNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListVersionNamesOK creates a ListVersionNamesOK with default headers values
func NewListVersionNamesOK() *ListVersionNamesOK {
	return &ListVersionNamesOK{}
}

/*
ListVersionNamesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListVersionNamesOK struct {
	Payload *service_model.V1ListProjectVersionsResponse
}

// IsSuccess returns true when this list version names o k response has a 2xx status code
func (o *ListVersionNamesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list version names o k response has a 3xx status code
func (o *ListVersionNamesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list version names o k response has a 4xx status code
func (o *ListVersionNamesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list version names o k response has a 5xx status code
func (o *ListVersionNamesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list version names o k response a status code equal to that given
func (o *ListVersionNamesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list version names o k response
func (o *ListVersionNamesOK) Code() int {
	return 200
}

func (o *ListVersionNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/versions/{kind}/names][%d] listVersionNamesOK  %+v", 200, o.Payload)
}

func (o *ListVersionNamesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/versions/{kind}/names][%d] listVersionNamesOK  %+v", 200, o.Payload)
}

func (o *ListVersionNamesOK) GetPayload() *service_model.V1ListProjectVersionsResponse {
	return o.Payload
}

func (o *ListVersionNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListProjectVersionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVersionNamesNoContent creates a ListVersionNamesNoContent with default headers values
func NewListVersionNamesNoContent() *ListVersionNamesNoContent {
	return &ListVersionNamesNoContent{}
}

/*
ListVersionNamesNoContent describes a response with status code 204, with default header values.

No content.
*/
type ListVersionNamesNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this list version names no content response has a 2xx status code
func (o *ListVersionNamesNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list version names no content response has a 3xx status code
func (o *ListVersionNamesNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list version names no content response has a 4xx status code
func (o *ListVersionNamesNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this list version names no content response has a 5xx status code
func (o *ListVersionNamesNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this list version names no content response a status code equal to that given
func (o *ListVersionNamesNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the list version names no content response
func (o *ListVersionNamesNoContent) Code() int {
	return 204
}

func (o *ListVersionNamesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/versions/{kind}/names][%d] listVersionNamesNoContent  %+v", 204, o.Payload)
}

func (o *ListVersionNamesNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/versions/{kind}/names][%d] listVersionNamesNoContent  %+v", 204, o.Payload)
}

func (o *ListVersionNamesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListVersionNamesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVersionNamesForbidden creates a ListVersionNamesForbidden with default headers values
func NewListVersionNamesForbidden() *ListVersionNamesForbidden {
	return &ListVersionNamesForbidden{}
}

/*
ListVersionNamesForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ListVersionNamesForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this list version names forbidden response has a 2xx status code
func (o *ListVersionNamesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list version names forbidden response has a 3xx status code
func (o *ListVersionNamesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list version names forbidden response has a 4xx status code
func (o *ListVersionNamesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list version names forbidden response has a 5xx status code
func (o *ListVersionNamesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list version names forbidden response a status code equal to that given
func (o *ListVersionNamesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list version names forbidden response
func (o *ListVersionNamesForbidden) Code() int {
	return 403
}

func (o *ListVersionNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/versions/{kind}/names][%d] listVersionNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListVersionNamesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/versions/{kind}/names][%d] listVersionNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListVersionNamesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListVersionNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVersionNamesNotFound creates a ListVersionNamesNotFound with default headers values
func NewListVersionNamesNotFound() *ListVersionNamesNotFound {
	return &ListVersionNamesNotFound{}
}

/*
ListVersionNamesNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ListVersionNamesNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this list version names not found response has a 2xx status code
func (o *ListVersionNamesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list version names not found response has a 3xx status code
func (o *ListVersionNamesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list version names not found response has a 4xx status code
func (o *ListVersionNamesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list version names not found response has a 5xx status code
func (o *ListVersionNamesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list version names not found response a status code equal to that given
func (o *ListVersionNamesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list version names not found response
func (o *ListVersionNamesNotFound) Code() int {
	return 404
}

func (o *ListVersionNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/versions/{kind}/names][%d] listVersionNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListVersionNamesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/versions/{kind}/names][%d] listVersionNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListVersionNamesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListVersionNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVersionNamesDefault creates a ListVersionNamesDefault with default headers values
func NewListVersionNamesDefault(code int) *ListVersionNamesDefault {
	return &ListVersionNamesDefault{
		_statusCode: code,
	}
}

/*
ListVersionNamesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListVersionNamesDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this list version names default response has a 2xx status code
func (o *ListVersionNamesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list version names default response has a 3xx status code
func (o *ListVersionNamesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list version names default response has a 4xx status code
func (o *ListVersionNamesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list version names default response has a 5xx status code
func (o *ListVersionNamesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list version names default response a status code equal to that given
func (o *ListVersionNamesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list version names default response
func (o *ListVersionNamesDefault) Code() int {
	return o._statusCode
}

func (o *ListVersionNamesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/versions/{kind}/names][%d] ListVersionNames default  %+v", o._statusCode, o.Payload)
}

func (o *ListVersionNamesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/versions/{kind}/names][%d] ListVersionNames default  %+v", o._statusCode, o.Payload)
}

func (o *ListVersionNamesDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListVersionNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
