// Code generated by go-swagger; DO NOT EDIT.

package agents_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// ListAgentsReader is a Reader for the ListAgents structure.
type ListAgentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAgentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAgentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListAgentsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListAgentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListAgentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListAgentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAgentsOK creates a ListAgentsOK with default headers values
func NewListAgentsOK() *ListAgentsOK {
	return &ListAgentsOK{}
}

/*
ListAgentsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListAgentsOK struct {
	Payload *service_model.V1ListAgentsResponse
}

// IsSuccess returns true when this list agents o k response has a 2xx status code
func (o *ListAgentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list agents o k response has a 3xx status code
func (o *ListAgentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list agents o k response has a 4xx status code
func (o *ListAgentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list agents o k response has a 5xx status code
func (o *ListAgentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list agents o k response a status code equal to that given
func (o *ListAgentsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list agents o k response
func (o *ListAgentsOK) Code() int {
	return 200
}

func (o *ListAgentsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents][%d] listAgentsOK  %+v", 200, o.Payload)
}

func (o *ListAgentsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents][%d] listAgentsOK  %+v", 200, o.Payload)
}

func (o *ListAgentsOK) GetPayload() *service_model.V1ListAgentsResponse {
	return o.Payload
}

func (o *ListAgentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListAgentsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAgentsNoContent creates a ListAgentsNoContent with default headers values
func NewListAgentsNoContent() *ListAgentsNoContent {
	return &ListAgentsNoContent{}
}

/*
ListAgentsNoContent describes a response with status code 204, with default header values.

No content.
*/
type ListAgentsNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this list agents no content response has a 2xx status code
func (o *ListAgentsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list agents no content response has a 3xx status code
func (o *ListAgentsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list agents no content response has a 4xx status code
func (o *ListAgentsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this list agents no content response has a 5xx status code
func (o *ListAgentsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this list agents no content response a status code equal to that given
func (o *ListAgentsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the list agents no content response
func (o *ListAgentsNoContent) Code() int {
	return 204
}

func (o *ListAgentsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents][%d] listAgentsNoContent  %+v", 204, o.Payload)
}

func (o *ListAgentsNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents][%d] listAgentsNoContent  %+v", 204, o.Payload)
}

func (o *ListAgentsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListAgentsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAgentsForbidden creates a ListAgentsForbidden with default headers values
func NewListAgentsForbidden() *ListAgentsForbidden {
	return &ListAgentsForbidden{}
}

/*
ListAgentsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ListAgentsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this list agents forbidden response has a 2xx status code
func (o *ListAgentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list agents forbidden response has a 3xx status code
func (o *ListAgentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list agents forbidden response has a 4xx status code
func (o *ListAgentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list agents forbidden response has a 5xx status code
func (o *ListAgentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list agents forbidden response a status code equal to that given
func (o *ListAgentsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list agents forbidden response
func (o *ListAgentsForbidden) Code() int {
	return 403
}

func (o *ListAgentsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents][%d] listAgentsForbidden  %+v", 403, o.Payload)
}

func (o *ListAgentsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents][%d] listAgentsForbidden  %+v", 403, o.Payload)
}

func (o *ListAgentsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListAgentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAgentsNotFound creates a ListAgentsNotFound with default headers values
func NewListAgentsNotFound() *ListAgentsNotFound {
	return &ListAgentsNotFound{}
}

/*
ListAgentsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ListAgentsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this list agents not found response has a 2xx status code
func (o *ListAgentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list agents not found response has a 3xx status code
func (o *ListAgentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list agents not found response has a 4xx status code
func (o *ListAgentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list agents not found response has a 5xx status code
func (o *ListAgentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list agents not found response a status code equal to that given
func (o *ListAgentsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list agents not found response
func (o *ListAgentsNotFound) Code() int {
	return 404
}

func (o *ListAgentsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents][%d] listAgentsNotFound  %+v", 404, o.Payload)
}

func (o *ListAgentsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents][%d] listAgentsNotFound  %+v", 404, o.Payload)
}

func (o *ListAgentsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListAgentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAgentsDefault creates a ListAgentsDefault with default headers values
func NewListAgentsDefault(code int) *ListAgentsDefault {
	return &ListAgentsDefault{
		_statusCode: code,
	}
}

/*
ListAgentsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListAgentsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this list agents default response has a 2xx status code
func (o *ListAgentsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list agents default response has a 3xx status code
func (o *ListAgentsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list agents default response has a 4xx status code
func (o *ListAgentsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list agents default response has a 5xx status code
func (o *ListAgentsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list agents default response a status code equal to that given
func (o *ListAgentsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list agents default response
func (o *ListAgentsDefault) Code() int {
	return o._statusCode
}

func (o *ListAgentsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents][%d] ListAgents default  %+v", o._statusCode, o.Payload)
}

func (o *ListAgentsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents][%d] ListAgents default  %+v", o._statusCode, o.Payload)
}

func (o *ListAgentsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListAgentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
