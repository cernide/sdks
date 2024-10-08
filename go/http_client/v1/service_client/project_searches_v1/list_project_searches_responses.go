// Code generated by go-swagger; DO NOT EDIT.

package project_searches_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// ListProjectSearchesReader is a Reader for the ListProjectSearches structure.
type ListProjectSearchesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectSearchesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectSearchesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListProjectSearchesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListProjectSearchesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListProjectSearchesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListProjectSearchesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectSearchesOK creates a ListProjectSearchesOK with default headers values
func NewListProjectSearchesOK() *ListProjectSearchesOK {
	return &ListProjectSearchesOK{}
}

/*
ListProjectSearchesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListProjectSearchesOK struct {
	Payload *service_model.V1ListSearchesResponse
}

// IsSuccess returns true when this list project searches o k response has a 2xx status code
func (o *ListProjectSearchesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project searches o k response has a 3xx status code
func (o *ListProjectSearchesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project searches o k response has a 4xx status code
func (o *ListProjectSearchesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project searches o k response has a 5xx status code
func (o *ListProjectSearchesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project searches o k response a status code equal to that given
func (o *ListProjectSearchesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list project searches o k response
func (o *ListProjectSearchesOK) Code() int {
	return 200
}

func (o *ListProjectSearchesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/searches][%d] listProjectSearchesOK  %+v", 200, o.Payload)
}

func (o *ListProjectSearchesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/searches][%d] listProjectSearchesOK  %+v", 200, o.Payload)
}

func (o *ListProjectSearchesOK) GetPayload() *service_model.V1ListSearchesResponse {
	return o.Payload
}

func (o *ListProjectSearchesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListSearchesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectSearchesNoContent creates a ListProjectSearchesNoContent with default headers values
func NewListProjectSearchesNoContent() *ListProjectSearchesNoContent {
	return &ListProjectSearchesNoContent{}
}

/*
ListProjectSearchesNoContent describes a response with status code 204, with default header values.

No content.
*/
type ListProjectSearchesNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this list project searches no content response has a 2xx status code
func (o *ListProjectSearchesNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project searches no content response has a 3xx status code
func (o *ListProjectSearchesNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project searches no content response has a 4xx status code
func (o *ListProjectSearchesNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project searches no content response has a 5xx status code
func (o *ListProjectSearchesNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this list project searches no content response a status code equal to that given
func (o *ListProjectSearchesNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the list project searches no content response
func (o *ListProjectSearchesNoContent) Code() int {
	return 204
}

func (o *ListProjectSearchesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/searches][%d] listProjectSearchesNoContent  %+v", 204, o.Payload)
}

func (o *ListProjectSearchesNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/searches][%d] listProjectSearchesNoContent  %+v", 204, o.Payload)
}

func (o *ListProjectSearchesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListProjectSearchesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectSearchesForbidden creates a ListProjectSearchesForbidden with default headers values
func NewListProjectSearchesForbidden() *ListProjectSearchesForbidden {
	return &ListProjectSearchesForbidden{}
}

/*
ListProjectSearchesForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ListProjectSearchesForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this list project searches forbidden response has a 2xx status code
func (o *ListProjectSearchesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project searches forbidden response has a 3xx status code
func (o *ListProjectSearchesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project searches forbidden response has a 4xx status code
func (o *ListProjectSearchesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project searches forbidden response has a 5xx status code
func (o *ListProjectSearchesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list project searches forbidden response a status code equal to that given
func (o *ListProjectSearchesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list project searches forbidden response
func (o *ListProjectSearchesForbidden) Code() int {
	return 403
}

func (o *ListProjectSearchesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/searches][%d] listProjectSearchesForbidden  %+v", 403, o.Payload)
}

func (o *ListProjectSearchesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/searches][%d] listProjectSearchesForbidden  %+v", 403, o.Payload)
}

func (o *ListProjectSearchesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListProjectSearchesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectSearchesNotFound creates a ListProjectSearchesNotFound with default headers values
func NewListProjectSearchesNotFound() *ListProjectSearchesNotFound {
	return &ListProjectSearchesNotFound{}
}

/*
ListProjectSearchesNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ListProjectSearchesNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this list project searches not found response has a 2xx status code
func (o *ListProjectSearchesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project searches not found response has a 3xx status code
func (o *ListProjectSearchesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project searches not found response has a 4xx status code
func (o *ListProjectSearchesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project searches not found response has a 5xx status code
func (o *ListProjectSearchesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list project searches not found response a status code equal to that given
func (o *ListProjectSearchesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list project searches not found response
func (o *ListProjectSearchesNotFound) Code() int {
	return 404
}

func (o *ListProjectSearchesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/searches][%d] listProjectSearchesNotFound  %+v", 404, o.Payload)
}

func (o *ListProjectSearchesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/searches][%d] listProjectSearchesNotFound  %+v", 404, o.Payload)
}

func (o *ListProjectSearchesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListProjectSearchesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectSearchesDefault creates a ListProjectSearchesDefault with default headers values
func NewListProjectSearchesDefault(code int) *ListProjectSearchesDefault {
	return &ListProjectSearchesDefault{
		_statusCode: code,
	}
}

/*
ListProjectSearchesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListProjectSearchesDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this list project searches default response has a 2xx status code
func (o *ListProjectSearchesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project searches default response has a 3xx status code
func (o *ListProjectSearchesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project searches default response has a 4xx status code
func (o *ListProjectSearchesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project searches default response has a 5xx status code
func (o *ListProjectSearchesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project searches default response a status code equal to that given
func (o *ListProjectSearchesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list project searches default response
func (o *ListProjectSearchesDefault) Code() int {
	return o._statusCode
}

func (o *ListProjectSearchesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/searches][%d] ListProjectSearches default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectSearchesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/searches][%d] ListProjectSearches default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectSearchesDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListProjectSearchesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
