// Code generated by go-swagger; DO NOT EDIT.

package queues_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/go/http_client/v1/service_model"
)

// UpdateQueueReader is a Reader for the UpdateQueue structure.
type UpdateQueueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateQueueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateQueueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateQueueNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateQueueForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateQueueNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateQueueDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateQueueOK creates a UpdateQueueOK with default headers values
func NewUpdateQueueOK() *UpdateQueueOK {
	return &UpdateQueueOK{}
}

/*
UpdateQueueOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateQueueOK struct {
	Payload *service_model.V1Queue
}

// IsSuccess returns true when this update queue o k response has a 2xx status code
func (o *UpdateQueueOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update queue o k response has a 3xx status code
func (o *UpdateQueueOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update queue o k response has a 4xx status code
func (o *UpdateQueueOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update queue o k response has a 5xx status code
func (o *UpdateQueueOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update queue o k response a status code equal to that given
func (o *UpdateQueueOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update queue o k response
func (o *UpdateQueueOK) Code() int {
	return 200
}

func (o *UpdateQueueOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/agents/{agent}/queues/{queue.uuid}][%d] updateQueueOK  %+v", 200, o.Payload)
}

func (o *UpdateQueueOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/agents/{agent}/queues/{queue.uuid}][%d] updateQueueOK  %+v", 200, o.Payload)
}

func (o *UpdateQueueOK) GetPayload() *service_model.V1Queue {
	return o.Payload
}

func (o *UpdateQueueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Queue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateQueueNoContent creates a UpdateQueueNoContent with default headers values
func NewUpdateQueueNoContent() *UpdateQueueNoContent {
	return &UpdateQueueNoContent{}
}

/*
UpdateQueueNoContent describes a response with status code 204, with default header values.

No content.
*/
type UpdateQueueNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this update queue no content response has a 2xx status code
func (o *UpdateQueueNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update queue no content response has a 3xx status code
func (o *UpdateQueueNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update queue no content response has a 4xx status code
func (o *UpdateQueueNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update queue no content response has a 5xx status code
func (o *UpdateQueueNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update queue no content response a status code equal to that given
func (o *UpdateQueueNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update queue no content response
func (o *UpdateQueueNoContent) Code() int {
	return 204
}

func (o *UpdateQueueNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/agents/{agent}/queues/{queue.uuid}][%d] updateQueueNoContent  %+v", 204, o.Payload)
}

func (o *UpdateQueueNoContent) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/agents/{agent}/queues/{queue.uuid}][%d] updateQueueNoContent  %+v", 204, o.Payload)
}

func (o *UpdateQueueNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateQueueNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateQueueForbidden creates a UpdateQueueForbidden with default headers values
func NewUpdateQueueForbidden() *UpdateQueueForbidden {
	return &UpdateQueueForbidden{}
}

/*
UpdateQueueForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type UpdateQueueForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this update queue forbidden response has a 2xx status code
func (o *UpdateQueueForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update queue forbidden response has a 3xx status code
func (o *UpdateQueueForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update queue forbidden response has a 4xx status code
func (o *UpdateQueueForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update queue forbidden response has a 5xx status code
func (o *UpdateQueueForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update queue forbidden response a status code equal to that given
func (o *UpdateQueueForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update queue forbidden response
func (o *UpdateQueueForbidden) Code() int {
	return 403
}

func (o *UpdateQueueForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/agents/{agent}/queues/{queue.uuid}][%d] updateQueueForbidden  %+v", 403, o.Payload)
}

func (o *UpdateQueueForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/agents/{agent}/queues/{queue.uuid}][%d] updateQueueForbidden  %+v", 403, o.Payload)
}

func (o *UpdateQueueForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateQueueForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateQueueNotFound creates a UpdateQueueNotFound with default headers values
func NewUpdateQueueNotFound() *UpdateQueueNotFound {
	return &UpdateQueueNotFound{}
}

/*
UpdateQueueNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type UpdateQueueNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this update queue not found response has a 2xx status code
func (o *UpdateQueueNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update queue not found response has a 3xx status code
func (o *UpdateQueueNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update queue not found response has a 4xx status code
func (o *UpdateQueueNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update queue not found response has a 5xx status code
func (o *UpdateQueueNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update queue not found response a status code equal to that given
func (o *UpdateQueueNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update queue not found response
func (o *UpdateQueueNotFound) Code() int {
	return 404
}

func (o *UpdateQueueNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/agents/{agent}/queues/{queue.uuid}][%d] updateQueueNotFound  %+v", 404, o.Payload)
}

func (o *UpdateQueueNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/agents/{agent}/queues/{queue.uuid}][%d] updateQueueNotFound  %+v", 404, o.Payload)
}

func (o *UpdateQueueNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateQueueNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateQueueDefault creates a UpdateQueueDefault with default headers values
func NewUpdateQueueDefault(code int) *UpdateQueueDefault {
	return &UpdateQueueDefault{
		_statusCode: code,
	}
}

/*
UpdateQueueDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateQueueDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this update queue default response has a 2xx status code
func (o *UpdateQueueDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update queue default response has a 3xx status code
func (o *UpdateQueueDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update queue default response has a 4xx status code
func (o *UpdateQueueDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update queue default response has a 5xx status code
func (o *UpdateQueueDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update queue default response a status code equal to that given
func (o *UpdateQueueDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update queue default response
func (o *UpdateQueueDefault) Code() int {
	return o._statusCode
}

func (o *UpdateQueueDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/agents/{agent}/queues/{queue.uuid}][%d] UpdateQueue default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateQueueDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/agents/{agent}/queues/{queue.uuid}][%d] UpdateQueue default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateQueueDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *UpdateQueueDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
