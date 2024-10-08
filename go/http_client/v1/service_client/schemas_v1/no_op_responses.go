// Code generated by go-swagger; DO NOT EDIT.

package schemas_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// NoOpReader is a Reader for the NoOp structure.
type NoOpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NoOpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNoOpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewNoOpNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewNoOpForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNoOpNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewNoOpDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNoOpOK creates a NoOpOK with default headers values
func NewNoOpOK() *NoOpOK {
	return &NoOpOK{}
}

/*
NoOpOK describes a response with status code 200, with default header values.

A successful response.
*/
type NoOpOK struct {
	Payload *service_model.V1Schemas
}

// IsSuccess returns true when this no op o k response has a 2xx status code
func (o *NoOpOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this no op o k response has a 3xx status code
func (o *NoOpOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this no op o k response has a 4xx status code
func (o *NoOpOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this no op o k response has a 5xx status code
func (o *NoOpOK) IsServerError() bool {
	return false
}

// IsCode returns true when this no op o k response a status code equal to that given
func (o *NoOpOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the no op o k response
func (o *NoOpOK) Code() int {
	return 200
}

func (o *NoOpOK) Error() string {
	return fmt.Sprintf("[GET /schemas][%d] noOpOK  %+v", 200, o.Payload)
}

func (o *NoOpOK) String() string {
	return fmt.Sprintf("[GET /schemas][%d] noOpOK  %+v", 200, o.Payload)
}

func (o *NoOpOK) GetPayload() *service_model.V1Schemas {
	return o.Payload
}

func (o *NoOpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Schemas)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNoOpNoContent creates a NoOpNoContent with default headers values
func NewNoOpNoContent() *NoOpNoContent {
	return &NoOpNoContent{}
}

/*
NoOpNoContent describes a response with status code 204, with default header values.

No content.
*/
type NoOpNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this no op no content response has a 2xx status code
func (o *NoOpNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this no op no content response has a 3xx status code
func (o *NoOpNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this no op no content response has a 4xx status code
func (o *NoOpNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this no op no content response has a 5xx status code
func (o *NoOpNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this no op no content response a status code equal to that given
func (o *NoOpNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the no op no content response
func (o *NoOpNoContent) Code() int {
	return 204
}

func (o *NoOpNoContent) Error() string {
	return fmt.Sprintf("[GET /schemas][%d] noOpNoContent  %+v", 204, o.Payload)
}

func (o *NoOpNoContent) String() string {
	return fmt.Sprintf("[GET /schemas][%d] noOpNoContent  %+v", 204, o.Payload)
}

func (o *NoOpNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *NoOpNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNoOpForbidden creates a NoOpForbidden with default headers values
func NewNoOpForbidden() *NoOpForbidden {
	return &NoOpForbidden{}
}

/*
NoOpForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type NoOpForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this no op forbidden response has a 2xx status code
func (o *NoOpForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this no op forbidden response has a 3xx status code
func (o *NoOpForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this no op forbidden response has a 4xx status code
func (o *NoOpForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this no op forbidden response has a 5xx status code
func (o *NoOpForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this no op forbidden response a status code equal to that given
func (o *NoOpForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the no op forbidden response
func (o *NoOpForbidden) Code() int {
	return 403
}

func (o *NoOpForbidden) Error() string {
	return fmt.Sprintf("[GET /schemas][%d] noOpForbidden  %+v", 403, o.Payload)
}

func (o *NoOpForbidden) String() string {
	return fmt.Sprintf("[GET /schemas][%d] noOpForbidden  %+v", 403, o.Payload)
}

func (o *NoOpForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *NoOpForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNoOpNotFound creates a NoOpNotFound with default headers values
func NewNoOpNotFound() *NoOpNotFound {
	return &NoOpNotFound{}
}

/*
NoOpNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type NoOpNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this no op not found response has a 2xx status code
func (o *NoOpNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this no op not found response has a 3xx status code
func (o *NoOpNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this no op not found response has a 4xx status code
func (o *NoOpNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this no op not found response has a 5xx status code
func (o *NoOpNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this no op not found response a status code equal to that given
func (o *NoOpNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the no op not found response
func (o *NoOpNotFound) Code() int {
	return 404
}

func (o *NoOpNotFound) Error() string {
	return fmt.Sprintf("[GET /schemas][%d] noOpNotFound  %+v", 404, o.Payload)
}

func (o *NoOpNotFound) String() string {
	return fmt.Sprintf("[GET /schemas][%d] noOpNotFound  %+v", 404, o.Payload)
}

func (o *NoOpNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *NoOpNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNoOpDefault creates a NoOpDefault with default headers values
func NewNoOpDefault(code int) *NoOpDefault {
	return &NoOpDefault{
		_statusCode: code,
	}
}

/*
NoOpDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type NoOpDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this no op default response has a 2xx status code
func (o *NoOpDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this no op default response has a 3xx status code
func (o *NoOpDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this no op default response has a 4xx status code
func (o *NoOpDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this no op default response has a 5xx status code
func (o *NoOpDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this no op default response a status code equal to that given
func (o *NoOpDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the no op default response
func (o *NoOpDefault) Code() int {
	return o._statusCode
}

func (o *NoOpDefault) Error() string {
	return fmt.Sprintf("[GET /schemas][%d] NoOp default  %+v", o._statusCode, o.Payload)
}

func (o *NoOpDefault) String() string {
	return fmt.Sprintf("[GET /schemas][%d] NoOp default  %+v", o._statusCode, o.Payload)
}

func (o *NoOpDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *NoOpDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
