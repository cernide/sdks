// Code generated by go-swagger; DO NOT EDIT.

package connections_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/go/http_client/v1/service_model"
)

// ListConnectionsReader is a Reader for the ListConnections structure.
type ListConnectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListConnectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListConnectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListConnectionsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListConnectionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListConnectionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListConnectionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListConnectionsOK creates a ListConnectionsOK with default headers values
func NewListConnectionsOK() *ListConnectionsOK {
	return &ListConnectionsOK{}
}

/*
ListConnectionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListConnectionsOK struct {
	Payload *service_model.V1ListConnectionsResponse
}

// IsSuccess returns true when this list connections o k response has a 2xx status code
func (o *ListConnectionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list connections o k response has a 3xx status code
func (o *ListConnectionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list connections o k response has a 4xx status code
func (o *ListConnectionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list connections o k response has a 5xx status code
func (o *ListConnectionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list connections o k response a status code equal to that given
func (o *ListConnectionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list connections o k response
func (o *ListConnectionsOK) Code() int {
	return 200
}

func (o *ListConnectionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections][%d] listConnectionsOK  %+v", 200, o.Payload)
}

func (o *ListConnectionsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections][%d] listConnectionsOK  %+v", 200, o.Payload)
}

func (o *ListConnectionsOK) GetPayload() *service_model.V1ListConnectionsResponse {
	return o.Payload
}

func (o *ListConnectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListConnectionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConnectionsNoContent creates a ListConnectionsNoContent with default headers values
func NewListConnectionsNoContent() *ListConnectionsNoContent {
	return &ListConnectionsNoContent{}
}

/*
ListConnectionsNoContent describes a response with status code 204, with default header values.

No content.
*/
type ListConnectionsNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this list connections no content response has a 2xx status code
func (o *ListConnectionsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list connections no content response has a 3xx status code
func (o *ListConnectionsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list connections no content response has a 4xx status code
func (o *ListConnectionsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this list connections no content response has a 5xx status code
func (o *ListConnectionsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this list connections no content response a status code equal to that given
func (o *ListConnectionsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the list connections no content response
func (o *ListConnectionsNoContent) Code() int {
	return 204
}

func (o *ListConnectionsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections][%d] listConnectionsNoContent  %+v", 204, o.Payload)
}

func (o *ListConnectionsNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections][%d] listConnectionsNoContent  %+v", 204, o.Payload)
}

func (o *ListConnectionsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListConnectionsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConnectionsForbidden creates a ListConnectionsForbidden with default headers values
func NewListConnectionsForbidden() *ListConnectionsForbidden {
	return &ListConnectionsForbidden{}
}

/*
ListConnectionsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ListConnectionsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this list connections forbidden response has a 2xx status code
func (o *ListConnectionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list connections forbidden response has a 3xx status code
func (o *ListConnectionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list connections forbidden response has a 4xx status code
func (o *ListConnectionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list connections forbidden response has a 5xx status code
func (o *ListConnectionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list connections forbidden response a status code equal to that given
func (o *ListConnectionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list connections forbidden response
func (o *ListConnectionsForbidden) Code() int {
	return 403
}

func (o *ListConnectionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections][%d] listConnectionsForbidden  %+v", 403, o.Payload)
}

func (o *ListConnectionsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections][%d] listConnectionsForbidden  %+v", 403, o.Payload)
}

func (o *ListConnectionsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListConnectionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConnectionsNotFound creates a ListConnectionsNotFound with default headers values
func NewListConnectionsNotFound() *ListConnectionsNotFound {
	return &ListConnectionsNotFound{}
}

/*
ListConnectionsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ListConnectionsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this list connections not found response has a 2xx status code
func (o *ListConnectionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list connections not found response has a 3xx status code
func (o *ListConnectionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list connections not found response has a 4xx status code
func (o *ListConnectionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list connections not found response has a 5xx status code
func (o *ListConnectionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list connections not found response a status code equal to that given
func (o *ListConnectionsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list connections not found response
func (o *ListConnectionsNotFound) Code() int {
	return 404
}

func (o *ListConnectionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections][%d] listConnectionsNotFound  %+v", 404, o.Payload)
}

func (o *ListConnectionsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections][%d] listConnectionsNotFound  %+v", 404, o.Payload)
}

func (o *ListConnectionsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListConnectionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConnectionsDefault creates a ListConnectionsDefault with default headers values
func NewListConnectionsDefault(code int) *ListConnectionsDefault {
	return &ListConnectionsDefault{
		_statusCode: code,
	}
}

/*
ListConnectionsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListConnectionsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this list connections default response has a 2xx status code
func (o *ListConnectionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list connections default response has a 3xx status code
func (o *ListConnectionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list connections default response has a 4xx status code
func (o *ListConnectionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list connections default response has a 5xx status code
func (o *ListConnectionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list connections default response a status code equal to that given
func (o *ListConnectionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list connections default response
func (o *ListConnectionsDefault) Code() int {
	return o._statusCode
}

func (o *ListConnectionsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections][%d] ListConnections default  %+v", o._statusCode, o.Payload)
}

func (o *ListConnectionsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections][%d] ListConnections default  %+v", o._statusCode, o.Payload)
}

func (o *ListConnectionsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListConnectionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
