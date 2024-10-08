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

// UpdateRunReader is a Reader for the UpdateRun structure.
type UpdateRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateRunNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateRunForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateRunOK creates a UpdateRunOK with default headers values
func NewUpdateRunOK() *UpdateRunOK {
	return &UpdateRunOK{}
}

/*
UpdateRunOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateRunOK struct {
	Payload *service_model.V1Run
}

// IsSuccess returns true when this update run o k response has a 2xx status code
func (o *UpdateRunOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update run o k response has a 3xx status code
func (o *UpdateRunOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update run o k response has a 4xx status code
func (o *UpdateRunOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update run o k response has a 5xx status code
func (o *UpdateRunOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update run o k response a status code equal to that given
func (o *UpdateRunOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update run o k response
func (o *UpdateRunOK) Code() int {
	return 200
}

func (o *UpdateRunOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/runs/{run.uuid}][%d] updateRunOK  %+v", 200, o.Payload)
}

func (o *UpdateRunOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/runs/{run.uuid}][%d] updateRunOK  %+v", 200, o.Payload)
}

func (o *UpdateRunOK) GetPayload() *service_model.V1Run {
	return o.Payload
}

func (o *UpdateRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Run)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunNoContent creates a UpdateRunNoContent with default headers values
func NewUpdateRunNoContent() *UpdateRunNoContent {
	return &UpdateRunNoContent{}
}

/*
UpdateRunNoContent describes a response with status code 204, with default header values.

No content.
*/
type UpdateRunNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this update run no content response has a 2xx status code
func (o *UpdateRunNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update run no content response has a 3xx status code
func (o *UpdateRunNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update run no content response has a 4xx status code
func (o *UpdateRunNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update run no content response has a 5xx status code
func (o *UpdateRunNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update run no content response a status code equal to that given
func (o *UpdateRunNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update run no content response
func (o *UpdateRunNoContent) Code() int {
	return 204
}

func (o *UpdateRunNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/runs/{run.uuid}][%d] updateRunNoContent  %+v", 204, o.Payload)
}

func (o *UpdateRunNoContent) String() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/runs/{run.uuid}][%d] updateRunNoContent  %+v", 204, o.Payload)
}

func (o *UpdateRunNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateRunNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunForbidden creates a UpdateRunForbidden with default headers values
func NewUpdateRunForbidden() *UpdateRunForbidden {
	return &UpdateRunForbidden{}
}

/*
UpdateRunForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type UpdateRunForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this update run forbidden response has a 2xx status code
func (o *UpdateRunForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update run forbidden response has a 3xx status code
func (o *UpdateRunForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update run forbidden response has a 4xx status code
func (o *UpdateRunForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update run forbidden response has a 5xx status code
func (o *UpdateRunForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update run forbidden response a status code equal to that given
func (o *UpdateRunForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update run forbidden response
func (o *UpdateRunForbidden) Code() int {
	return 403
}

func (o *UpdateRunForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/runs/{run.uuid}][%d] updateRunForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRunForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/runs/{run.uuid}][%d] updateRunForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRunForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateRunForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunNotFound creates a UpdateRunNotFound with default headers values
func NewUpdateRunNotFound() *UpdateRunNotFound {
	return &UpdateRunNotFound{}
}

/*
UpdateRunNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type UpdateRunNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this update run not found response has a 2xx status code
func (o *UpdateRunNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update run not found response has a 3xx status code
func (o *UpdateRunNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update run not found response has a 4xx status code
func (o *UpdateRunNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update run not found response has a 5xx status code
func (o *UpdateRunNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update run not found response a status code equal to that given
func (o *UpdateRunNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update run not found response
func (o *UpdateRunNotFound) Code() int {
	return 404
}

func (o *UpdateRunNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/runs/{run.uuid}][%d] updateRunNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRunNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/runs/{run.uuid}][%d] updateRunNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRunNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunDefault creates a UpdateRunDefault with default headers values
func NewUpdateRunDefault(code int) *UpdateRunDefault {
	return &UpdateRunDefault{
		_statusCode: code,
	}
}

/*
UpdateRunDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateRunDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this update run default response has a 2xx status code
func (o *UpdateRunDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update run default response has a 3xx status code
func (o *UpdateRunDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update run default response has a 4xx status code
func (o *UpdateRunDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update run default response has a 5xx status code
func (o *UpdateRunDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update run default response a status code equal to that given
func (o *UpdateRunDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update run default response
func (o *UpdateRunDefault) Code() int {
	return o._statusCode
}

func (o *UpdateRunDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/runs/{run.uuid}][%d] UpdateRun default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateRunDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/runs/{run.uuid}][%d] UpdateRun default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateRunDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *UpdateRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
