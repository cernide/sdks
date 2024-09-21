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

// InspectRunReader is a Reader for the InspectRun structure.
type InspectRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InspectRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInspectRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewInspectRunNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewInspectRunForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInspectRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewInspectRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInspectRunOK creates a InspectRunOK with default headers values
func NewInspectRunOK() *InspectRunOK {
	return &InspectRunOK{}
}

/*
InspectRunOK describes a response with status code 200, with default header values.

A successful response.
*/
type InspectRunOK struct {
	Payload interface{}
}

// IsSuccess returns true when this inspect run o k response has a 2xx status code
func (o *InspectRunOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inspect run o k response has a 3xx status code
func (o *InspectRunOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inspect run o k response has a 4xx status code
func (o *InspectRunOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this inspect run o k response has a 5xx status code
func (o *InspectRunOK) IsServerError() bool {
	return false
}

// IsCode returns true when this inspect run o k response a status code equal to that given
func (o *InspectRunOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the inspect run o k response
func (o *InspectRunOK) Code() int {
	return 200
}

func (o *InspectRunOK) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/k8s_inspect][%d] inspectRunOK  %+v", 200, o.Payload)
}

func (o *InspectRunOK) String() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/k8s_inspect][%d] inspectRunOK  %+v", 200, o.Payload)
}

func (o *InspectRunOK) GetPayload() interface{} {
	return o.Payload
}

func (o *InspectRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInspectRunNoContent creates a InspectRunNoContent with default headers values
func NewInspectRunNoContent() *InspectRunNoContent {
	return &InspectRunNoContent{}
}

/*
InspectRunNoContent describes a response with status code 204, with default header values.

No content.
*/
type InspectRunNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this inspect run no content response has a 2xx status code
func (o *InspectRunNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inspect run no content response has a 3xx status code
func (o *InspectRunNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inspect run no content response has a 4xx status code
func (o *InspectRunNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this inspect run no content response has a 5xx status code
func (o *InspectRunNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this inspect run no content response a status code equal to that given
func (o *InspectRunNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the inspect run no content response
func (o *InspectRunNoContent) Code() int {
	return 204
}

func (o *InspectRunNoContent) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/k8s_inspect][%d] inspectRunNoContent  %+v", 204, o.Payload)
}

func (o *InspectRunNoContent) String() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/k8s_inspect][%d] inspectRunNoContent  %+v", 204, o.Payload)
}

func (o *InspectRunNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *InspectRunNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInspectRunForbidden creates a InspectRunForbidden with default headers values
func NewInspectRunForbidden() *InspectRunForbidden {
	return &InspectRunForbidden{}
}

/*
InspectRunForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type InspectRunForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this inspect run forbidden response has a 2xx status code
func (o *InspectRunForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this inspect run forbidden response has a 3xx status code
func (o *InspectRunForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inspect run forbidden response has a 4xx status code
func (o *InspectRunForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this inspect run forbidden response has a 5xx status code
func (o *InspectRunForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this inspect run forbidden response a status code equal to that given
func (o *InspectRunForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the inspect run forbidden response
func (o *InspectRunForbidden) Code() int {
	return 403
}

func (o *InspectRunForbidden) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/k8s_inspect][%d] inspectRunForbidden  %+v", 403, o.Payload)
}

func (o *InspectRunForbidden) String() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/k8s_inspect][%d] inspectRunForbidden  %+v", 403, o.Payload)
}

func (o *InspectRunForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *InspectRunForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInspectRunNotFound creates a InspectRunNotFound with default headers values
func NewInspectRunNotFound() *InspectRunNotFound {
	return &InspectRunNotFound{}
}

/*
InspectRunNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type InspectRunNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this inspect run not found response has a 2xx status code
func (o *InspectRunNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this inspect run not found response has a 3xx status code
func (o *InspectRunNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inspect run not found response has a 4xx status code
func (o *InspectRunNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this inspect run not found response has a 5xx status code
func (o *InspectRunNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this inspect run not found response a status code equal to that given
func (o *InspectRunNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the inspect run not found response
func (o *InspectRunNotFound) Code() int {
	return 404
}

func (o *InspectRunNotFound) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/k8s_inspect][%d] inspectRunNotFound  %+v", 404, o.Payload)
}

func (o *InspectRunNotFound) String() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/k8s_inspect][%d] inspectRunNotFound  %+v", 404, o.Payload)
}

func (o *InspectRunNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *InspectRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInspectRunDefault creates a InspectRunDefault with default headers values
func NewInspectRunDefault(code int) *InspectRunDefault {
	return &InspectRunDefault{
		_statusCode: code,
	}
}

/*
InspectRunDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type InspectRunDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this inspect run default response has a 2xx status code
func (o *InspectRunDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this inspect run default response has a 3xx status code
func (o *InspectRunDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this inspect run default response has a 4xx status code
func (o *InspectRunDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this inspect run default response has a 5xx status code
func (o *InspectRunDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this inspect run default response a status code equal to that given
func (o *InspectRunDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the inspect run default response
func (o *InspectRunDefault) Code() int {
	return o._statusCode
}

func (o *InspectRunDefault) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/k8s_inspect][%d] InspectRun default  %+v", o._statusCode, o.Payload)
}

func (o *InspectRunDefault) String() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/k8s_inspect][%d] InspectRun default  %+v", o._statusCode, o.Payload)
}

func (o *InspectRunDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *InspectRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
