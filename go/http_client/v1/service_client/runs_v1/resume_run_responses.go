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

// ResumeRunReader is a Reader for the ResumeRun structure.
type ResumeRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResumeRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResumeRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewResumeRunNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewResumeRunForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewResumeRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewResumeRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResumeRunOK creates a ResumeRunOK with default headers values
func NewResumeRunOK() *ResumeRunOK {
	return &ResumeRunOK{}
}

/*
ResumeRunOK describes a response with status code 200, with default header values.

A successful response.
*/
type ResumeRunOK struct {
	Payload *service_model.V1Run
}

// IsSuccess returns true when this resume run o k response has a 2xx status code
func (o *ResumeRunOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resume run o k response has a 3xx status code
func (o *ResumeRunOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resume run o k response has a 4xx status code
func (o *ResumeRunOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this resume run o k response has a 5xx status code
func (o *ResumeRunOK) IsServerError() bool {
	return false
}

// IsCode returns true when this resume run o k response a status code equal to that given
func (o *ResumeRunOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the resume run o k response
func (o *ResumeRunOK) Code() int {
	return 200
}

func (o *ResumeRunOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run.uuid}/resume][%d] resumeRunOK  %+v", 200, o.Payload)
}

func (o *ResumeRunOK) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run.uuid}/resume][%d] resumeRunOK  %+v", 200, o.Payload)
}

func (o *ResumeRunOK) GetPayload() *service_model.V1Run {
	return o.Payload
}

func (o *ResumeRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Run)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResumeRunNoContent creates a ResumeRunNoContent with default headers values
func NewResumeRunNoContent() *ResumeRunNoContent {
	return &ResumeRunNoContent{}
}

/*
ResumeRunNoContent describes a response with status code 204, with default header values.

No content.
*/
type ResumeRunNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this resume run no content response has a 2xx status code
func (o *ResumeRunNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resume run no content response has a 3xx status code
func (o *ResumeRunNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resume run no content response has a 4xx status code
func (o *ResumeRunNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this resume run no content response has a 5xx status code
func (o *ResumeRunNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this resume run no content response a status code equal to that given
func (o *ResumeRunNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the resume run no content response
func (o *ResumeRunNoContent) Code() int {
	return 204
}

func (o *ResumeRunNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run.uuid}/resume][%d] resumeRunNoContent  %+v", 204, o.Payload)
}

func (o *ResumeRunNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run.uuid}/resume][%d] resumeRunNoContent  %+v", 204, o.Payload)
}

func (o *ResumeRunNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ResumeRunNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResumeRunForbidden creates a ResumeRunForbidden with default headers values
func NewResumeRunForbidden() *ResumeRunForbidden {
	return &ResumeRunForbidden{}
}

/*
ResumeRunForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ResumeRunForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this resume run forbidden response has a 2xx status code
func (o *ResumeRunForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resume run forbidden response has a 3xx status code
func (o *ResumeRunForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resume run forbidden response has a 4xx status code
func (o *ResumeRunForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this resume run forbidden response has a 5xx status code
func (o *ResumeRunForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this resume run forbidden response a status code equal to that given
func (o *ResumeRunForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the resume run forbidden response
func (o *ResumeRunForbidden) Code() int {
	return 403
}

func (o *ResumeRunForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run.uuid}/resume][%d] resumeRunForbidden  %+v", 403, o.Payload)
}

func (o *ResumeRunForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run.uuid}/resume][%d] resumeRunForbidden  %+v", 403, o.Payload)
}

func (o *ResumeRunForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ResumeRunForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResumeRunNotFound creates a ResumeRunNotFound with default headers values
func NewResumeRunNotFound() *ResumeRunNotFound {
	return &ResumeRunNotFound{}
}

/*
ResumeRunNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ResumeRunNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this resume run not found response has a 2xx status code
func (o *ResumeRunNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resume run not found response has a 3xx status code
func (o *ResumeRunNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resume run not found response has a 4xx status code
func (o *ResumeRunNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this resume run not found response has a 5xx status code
func (o *ResumeRunNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this resume run not found response a status code equal to that given
func (o *ResumeRunNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the resume run not found response
func (o *ResumeRunNotFound) Code() int {
	return 404
}

func (o *ResumeRunNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run.uuid}/resume][%d] resumeRunNotFound  %+v", 404, o.Payload)
}

func (o *ResumeRunNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run.uuid}/resume][%d] resumeRunNotFound  %+v", 404, o.Payload)
}

func (o *ResumeRunNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ResumeRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResumeRunDefault creates a ResumeRunDefault with default headers values
func NewResumeRunDefault(code int) *ResumeRunDefault {
	return &ResumeRunDefault{
		_statusCode: code,
	}
}

/*
ResumeRunDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ResumeRunDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this resume run default response has a 2xx status code
func (o *ResumeRunDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this resume run default response has a 3xx status code
func (o *ResumeRunDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this resume run default response has a 4xx status code
func (o *ResumeRunDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this resume run default response has a 5xx status code
func (o *ResumeRunDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this resume run default response a status code equal to that given
func (o *ResumeRunDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the resume run default response
func (o *ResumeRunDefault) Code() int {
	return o._statusCode
}

func (o *ResumeRunDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run.uuid}/resume][%d] ResumeRun default  %+v", o._statusCode, o.Payload)
}

func (o *ResumeRunDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run.uuid}/resume][%d] ResumeRun default  %+v", o._statusCode, o.Payload)
}

func (o *ResumeRunDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ResumeRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
