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

// DeleteProjectReader is a Reader for the DeleteProject structure.
type DeleteProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteProjectNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteProjectOK creates a DeleteProjectOK with default headers values
func NewDeleteProjectOK() *DeleteProjectOK {
	return &DeleteProjectOK{}
}

/*
DeleteProjectOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteProjectOK struct {
}

// IsSuccess returns true when this delete project o k response has a 2xx status code
func (o *DeleteProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete project o k response has a 3xx status code
func (o *DeleteProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project o k response has a 4xx status code
func (o *DeleteProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete project o k response has a 5xx status code
func (o *DeleteProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project o k response a status code equal to that given
func (o *DeleteProjectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete project o k response
func (o *DeleteProjectOK) Code() int {
	return 200
}

func (o *DeleteProjectOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{name}][%d] deleteProjectOK ", 200)
}

func (o *DeleteProjectOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{name}][%d] deleteProjectOK ", 200)
}

func (o *DeleteProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectNoContent creates a DeleteProjectNoContent with default headers values
func NewDeleteProjectNoContent() *DeleteProjectNoContent {
	return &DeleteProjectNoContent{}
}

/*
DeleteProjectNoContent describes a response with status code 204, with default header values.

No content.
*/
type DeleteProjectNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this delete project no content response has a 2xx status code
func (o *DeleteProjectNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete project no content response has a 3xx status code
func (o *DeleteProjectNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project no content response has a 4xx status code
func (o *DeleteProjectNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete project no content response has a 5xx status code
func (o *DeleteProjectNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project no content response a status code equal to that given
func (o *DeleteProjectNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete project no content response
func (o *DeleteProjectNoContent) Code() int {
	return 204
}

func (o *DeleteProjectNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{name}][%d] deleteProjectNoContent  %+v", 204, o.Payload)
}

func (o *DeleteProjectNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{name}][%d] deleteProjectNoContent  %+v", 204, o.Payload)
}

func (o *DeleteProjectNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteProjectNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectForbidden creates a DeleteProjectForbidden with default headers values
func NewDeleteProjectForbidden() *DeleteProjectForbidden {
	return &DeleteProjectForbidden{}
}

/*
DeleteProjectForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type DeleteProjectForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this delete project forbidden response has a 2xx status code
func (o *DeleteProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete project forbidden response has a 3xx status code
func (o *DeleteProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project forbidden response has a 4xx status code
func (o *DeleteProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete project forbidden response has a 5xx status code
func (o *DeleteProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project forbidden response a status code equal to that given
func (o *DeleteProjectForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete project forbidden response
func (o *DeleteProjectForbidden) Code() int {
	return 403
}

func (o *DeleteProjectForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{name}][%d] deleteProjectForbidden  %+v", 403, o.Payload)
}

func (o *DeleteProjectForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{name}][%d] deleteProjectForbidden  %+v", 403, o.Payload)
}

func (o *DeleteProjectForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectNotFound creates a DeleteProjectNotFound with default headers values
func NewDeleteProjectNotFound() *DeleteProjectNotFound {
	return &DeleteProjectNotFound{}
}

/*
DeleteProjectNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type DeleteProjectNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this delete project not found response has a 2xx status code
func (o *DeleteProjectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete project not found response has a 3xx status code
func (o *DeleteProjectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project not found response has a 4xx status code
func (o *DeleteProjectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete project not found response has a 5xx status code
func (o *DeleteProjectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project not found response a status code equal to that given
func (o *DeleteProjectNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete project not found response
func (o *DeleteProjectNotFound) Code() int {
	return 404
}

func (o *DeleteProjectNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{name}][%d] deleteProjectNotFound  %+v", 404, o.Payload)
}

func (o *DeleteProjectNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{name}][%d] deleteProjectNotFound  %+v", 404, o.Payload)
}

func (o *DeleteProjectNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectDefault creates a DeleteProjectDefault with default headers values
func NewDeleteProjectDefault(code int) *DeleteProjectDefault {
	return &DeleteProjectDefault{
		_statusCode: code,
	}
}

/*
DeleteProjectDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteProjectDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this delete project default response has a 2xx status code
func (o *DeleteProjectDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete project default response has a 3xx status code
func (o *DeleteProjectDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete project default response has a 4xx status code
func (o *DeleteProjectDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete project default response has a 5xx status code
func (o *DeleteProjectDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete project default response a status code equal to that given
func (o *DeleteProjectDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete project default response
func (o *DeleteProjectDefault) Code() int {
	return o._statusCode
}

func (o *DeleteProjectDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{name}][%d] DeleteProject default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteProjectDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{name}][%d] DeleteProject default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteProjectDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *DeleteProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
