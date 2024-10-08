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

// ListArchivedRunsReader is a Reader for the ListArchivedRuns structure.
type ListArchivedRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListArchivedRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListArchivedRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListArchivedRunsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListArchivedRunsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListArchivedRunsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListArchivedRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListArchivedRunsOK creates a ListArchivedRunsOK with default headers values
func NewListArchivedRunsOK() *ListArchivedRunsOK {
	return &ListArchivedRunsOK{}
}

/*
ListArchivedRunsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListArchivedRunsOK struct {
	Payload *service_model.V1ListRunsResponse
}

// IsSuccess returns true when this list archived runs o k response has a 2xx status code
func (o *ListArchivedRunsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list archived runs o k response has a 3xx status code
func (o *ListArchivedRunsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list archived runs o k response has a 4xx status code
func (o *ListArchivedRunsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list archived runs o k response has a 5xx status code
func (o *ListArchivedRunsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list archived runs o k response a status code equal to that given
func (o *ListArchivedRunsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list archived runs o k response
func (o *ListArchivedRunsOK) Code() int {
	return 200
}

func (o *ListArchivedRunsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/runs][%d] listArchivedRunsOK  %+v", 200, o.Payload)
}

func (o *ListArchivedRunsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/runs][%d] listArchivedRunsOK  %+v", 200, o.Payload)
}

func (o *ListArchivedRunsOK) GetPayload() *service_model.V1ListRunsResponse {
	return o.Payload
}

func (o *ListArchivedRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListRunsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArchivedRunsNoContent creates a ListArchivedRunsNoContent with default headers values
func NewListArchivedRunsNoContent() *ListArchivedRunsNoContent {
	return &ListArchivedRunsNoContent{}
}

/*
ListArchivedRunsNoContent describes a response with status code 204, with default header values.

No content.
*/
type ListArchivedRunsNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this list archived runs no content response has a 2xx status code
func (o *ListArchivedRunsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list archived runs no content response has a 3xx status code
func (o *ListArchivedRunsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list archived runs no content response has a 4xx status code
func (o *ListArchivedRunsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this list archived runs no content response has a 5xx status code
func (o *ListArchivedRunsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this list archived runs no content response a status code equal to that given
func (o *ListArchivedRunsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the list archived runs no content response
func (o *ListArchivedRunsNoContent) Code() int {
	return 204
}

func (o *ListArchivedRunsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/runs][%d] listArchivedRunsNoContent  %+v", 204, o.Payload)
}

func (o *ListArchivedRunsNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/runs][%d] listArchivedRunsNoContent  %+v", 204, o.Payload)
}

func (o *ListArchivedRunsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListArchivedRunsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArchivedRunsForbidden creates a ListArchivedRunsForbidden with default headers values
func NewListArchivedRunsForbidden() *ListArchivedRunsForbidden {
	return &ListArchivedRunsForbidden{}
}

/*
ListArchivedRunsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ListArchivedRunsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this list archived runs forbidden response has a 2xx status code
func (o *ListArchivedRunsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list archived runs forbidden response has a 3xx status code
func (o *ListArchivedRunsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list archived runs forbidden response has a 4xx status code
func (o *ListArchivedRunsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list archived runs forbidden response has a 5xx status code
func (o *ListArchivedRunsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list archived runs forbidden response a status code equal to that given
func (o *ListArchivedRunsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list archived runs forbidden response
func (o *ListArchivedRunsForbidden) Code() int {
	return 403
}

func (o *ListArchivedRunsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/runs][%d] listArchivedRunsForbidden  %+v", 403, o.Payload)
}

func (o *ListArchivedRunsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/runs][%d] listArchivedRunsForbidden  %+v", 403, o.Payload)
}

func (o *ListArchivedRunsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListArchivedRunsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArchivedRunsNotFound creates a ListArchivedRunsNotFound with default headers values
func NewListArchivedRunsNotFound() *ListArchivedRunsNotFound {
	return &ListArchivedRunsNotFound{}
}

/*
ListArchivedRunsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ListArchivedRunsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this list archived runs not found response has a 2xx status code
func (o *ListArchivedRunsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list archived runs not found response has a 3xx status code
func (o *ListArchivedRunsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list archived runs not found response has a 4xx status code
func (o *ListArchivedRunsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list archived runs not found response has a 5xx status code
func (o *ListArchivedRunsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list archived runs not found response a status code equal to that given
func (o *ListArchivedRunsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list archived runs not found response
func (o *ListArchivedRunsNotFound) Code() int {
	return 404
}

func (o *ListArchivedRunsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/runs][%d] listArchivedRunsNotFound  %+v", 404, o.Payload)
}

func (o *ListArchivedRunsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/runs][%d] listArchivedRunsNotFound  %+v", 404, o.Payload)
}

func (o *ListArchivedRunsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListArchivedRunsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArchivedRunsDefault creates a ListArchivedRunsDefault with default headers values
func NewListArchivedRunsDefault(code int) *ListArchivedRunsDefault {
	return &ListArchivedRunsDefault{
		_statusCode: code,
	}
}

/*
ListArchivedRunsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListArchivedRunsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this list archived runs default response has a 2xx status code
func (o *ListArchivedRunsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list archived runs default response has a 3xx status code
func (o *ListArchivedRunsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list archived runs default response has a 4xx status code
func (o *ListArchivedRunsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list archived runs default response has a 5xx status code
func (o *ListArchivedRunsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list archived runs default response a status code equal to that given
func (o *ListArchivedRunsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list archived runs default response
func (o *ListArchivedRunsDefault) Code() int {
	return o._statusCode
}

func (o *ListArchivedRunsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/runs][%d] ListArchivedRuns default  %+v", o._statusCode, o.Payload)
}

func (o *ListArchivedRunsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/runs][%d] ListArchivedRuns default  %+v", o._statusCode, o.Payload)
}

func (o *ListArchivedRunsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListArchivedRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
