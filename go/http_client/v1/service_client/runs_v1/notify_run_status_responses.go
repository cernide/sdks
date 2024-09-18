// Code generated by go-swagger; DO NOT EDIT.

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/go/http_client/v1/service_model"
)

// NotifyRunStatusReader is a Reader for the NotifyRunStatus structure.
type NotifyRunStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotifyRunStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotifyRunStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewNotifyRunStatusNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewNotifyRunStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNotifyRunStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewNotifyRunStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNotifyRunStatusOK creates a NotifyRunStatusOK with default headers values
func NewNotifyRunStatusOK() *NotifyRunStatusOK {
	return &NotifyRunStatusOK{}
}

/*
NotifyRunStatusOK describes a response with status code 200, with default header values.

A successful response.
*/
type NotifyRunStatusOK struct {
}

// IsSuccess returns true when this notify run status o k response has a 2xx status code
func (o *NotifyRunStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this notify run status o k response has a 3xx status code
func (o *NotifyRunStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notify run status o k response has a 4xx status code
func (o *NotifyRunStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this notify run status o k response has a 5xx status code
func (o *NotifyRunStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this notify run status o k response a status code equal to that given
func (o *NotifyRunStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the notify run status o k response
func (o *NotifyRunStatusOK) Code() int {
	return 200
}

func (o *NotifyRunStatusOK) Error() string {
	return fmt.Sprintf("[POST /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/notify][%d] notifyRunStatusOK ", 200)
}

func (o *NotifyRunStatusOK) String() string {
	return fmt.Sprintf("[POST /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/notify][%d] notifyRunStatusOK ", 200)
}

func (o *NotifyRunStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNotifyRunStatusNoContent creates a NotifyRunStatusNoContent with default headers values
func NewNotifyRunStatusNoContent() *NotifyRunStatusNoContent {
	return &NotifyRunStatusNoContent{}
}

/*
NotifyRunStatusNoContent describes a response with status code 204, with default header values.

No content.
*/
type NotifyRunStatusNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this notify run status no content response has a 2xx status code
func (o *NotifyRunStatusNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this notify run status no content response has a 3xx status code
func (o *NotifyRunStatusNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notify run status no content response has a 4xx status code
func (o *NotifyRunStatusNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this notify run status no content response has a 5xx status code
func (o *NotifyRunStatusNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this notify run status no content response a status code equal to that given
func (o *NotifyRunStatusNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the notify run status no content response
func (o *NotifyRunStatusNoContent) Code() int {
	return 204
}

func (o *NotifyRunStatusNoContent) Error() string {
	return fmt.Sprintf("[POST /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/notify][%d] notifyRunStatusNoContent  %+v", 204, o.Payload)
}

func (o *NotifyRunStatusNoContent) String() string {
	return fmt.Sprintf("[POST /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/notify][%d] notifyRunStatusNoContent  %+v", 204, o.Payload)
}

func (o *NotifyRunStatusNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *NotifyRunStatusNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotifyRunStatusForbidden creates a NotifyRunStatusForbidden with default headers values
func NewNotifyRunStatusForbidden() *NotifyRunStatusForbidden {
	return &NotifyRunStatusForbidden{}
}

/*
NotifyRunStatusForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type NotifyRunStatusForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this notify run status forbidden response has a 2xx status code
func (o *NotifyRunStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notify run status forbidden response has a 3xx status code
func (o *NotifyRunStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notify run status forbidden response has a 4xx status code
func (o *NotifyRunStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this notify run status forbidden response has a 5xx status code
func (o *NotifyRunStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this notify run status forbidden response a status code equal to that given
func (o *NotifyRunStatusForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the notify run status forbidden response
func (o *NotifyRunStatusForbidden) Code() int {
	return 403
}

func (o *NotifyRunStatusForbidden) Error() string {
	return fmt.Sprintf("[POST /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/notify][%d] notifyRunStatusForbidden  %+v", 403, o.Payload)
}

func (o *NotifyRunStatusForbidden) String() string {
	return fmt.Sprintf("[POST /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/notify][%d] notifyRunStatusForbidden  %+v", 403, o.Payload)
}

func (o *NotifyRunStatusForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *NotifyRunStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotifyRunStatusNotFound creates a NotifyRunStatusNotFound with default headers values
func NewNotifyRunStatusNotFound() *NotifyRunStatusNotFound {
	return &NotifyRunStatusNotFound{}
}

/*
NotifyRunStatusNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type NotifyRunStatusNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this notify run status not found response has a 2xx status code
func (o *NotifyRunStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notify run status not found response has a 3xx status code
func (o *NotifyRunStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notify run status not found response has a 4xx status code
func (o *NotifyRunStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this notify run status not found response has a 5xx status code
func (o *NotifyRunStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this notify run status not found response a status code equal to that given
func (o *NotifyRunStatusNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the notify run status not found response
func (o *NotifyRunStatusNotFound) Code() int {
	return 404
}

func (o *NotifyRunStatusNotFound) Error() string {
	return fmt.Sprintf("[POST /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/notify][%d] notifyRunStatusNotFound  %+v", 404, o.Payload)
}

func (o *NotifyRunStatusNotFound) String() string {
	return fmt.Sprintf("[POST /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/notify][%d] notifyRunStatusNotFound  %+v", 404, o.Payload)
}

func (o *NotifyRunStatusNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *NotifyRunStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotifyRunStatusDefault creates a NotifyRunStatusDefault with default headers values
func NewNotifyRunStatusDefault(code int) *NotifyRunStatusDefault {
	return &NotifyRunStatusDefault{
		_statusCode: code,
	}
}

/*
NotifyRunStatusDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type NotifyRunStatusDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this notify run status default response has a 2xx status code
func (o *NotifyRunStatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this notify run status default response has a 3xx status code
func (o *NotifyRunStatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this notify run status default response has a 4xx status code
func (o *NotifyRunStatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this notify run status default response has a 5xx status code
func (o *NotifyRunStatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this notify run status default response a status code equal to that given
func (o *NotifyRunStatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the notify run status default response
func (o *NotifyRunStatusDefault) Code() int {
	return o._statusCode
}

func (o *NotifyRunStatusDefault) Error() string {
	return fmt.Sprintf("[POST /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/notify][%d] NotifyRunStatus default  %+v", o._statusCode, o.Payload)
}

func (o *NotifyRunStatusDefault) String() string {
	return fmt.Sprintf("[POST /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/notify][%d] NotifyRunStatus default  %+v", o._statusCode, o.Payload)
}

func (o *NotifyRunStatusDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *NotifyRunStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
