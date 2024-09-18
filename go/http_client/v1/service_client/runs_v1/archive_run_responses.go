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

// ArchiveRunReader is a Reader for the ArchiveRun structure.
type ArchiveRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArchiveRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewArchiveRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewArchiveRunNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewArchiveRunForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewArchiveRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewArchiveRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewArchiveRunOK creates a ArchiveRunOK with default headers values
func NewArchiveRunOK() *ArchiveRunOK {
	return &ArchiveRunOK{}
}

/*
ArchiveRunOK describes a response with status code 200, with default header values.

A successful response.
*/
type ArchiveRunOK struct {
}

// IsSuccess returns true when this archive run o k response has a 2xx status code
func (o *ArchiveRunOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this archive run o k response has a 3xx status code
func (o *ArchiveRunOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive run o k response has a 4xx status code
func (o *ArchiveRunOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this archive run o k response has a 5xx status code
func (o *ArchiveRunOK) IsServerError() bool {
	return false
}

// IsCode returns true when this archive run o k response a status code equal to that given
func (o *ArchiveRunOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the archive run o k response
func (o *ArchiveRunOK) Code() int {
	return 200
}

func (o *ArchiveRunOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/archive][%d] archiveRunOK ", 200)
}

func (o *ArchiveRunOK) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/archive][%d] archiveRunOK ", 200)
}

func (o *ArchiveRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewArchiveRunNoContent creates a ArchiveRunNoContent with default headers values
func NewArchiveRunNoContent() *ArchiveRunNoContent {
	return &ArchiveRunNoContent{}
}

/*
ArchiveRunNoContent describes a response with status code 204, with default header values.

No content.
*/
type ArchiveRunNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this archive run no content response has a 2xx status code
func (o *ArchiveRunNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this archive run no content response has a 3xx status code
func (o *ArchiveRunNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive run no content response has a 4xx status code
func (o *ArchiveRunNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this archive run no content response has a 5xx status code
func (o *ArchiveRunNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this archive run no content response a status code equal to that given
func (o *ArchiveRunNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the archive run no content response
func (o *ArchiveRunNoContent) Code() int {
	return 204
}

func (o *ArchiveRunNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/archive][%d] archiveRunNoContent  %+v", 204, o.Payload)
}

func (o *ArchiveRunNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/archive][%d] archiveRunNoContent  %+v", 204, o.Payload)
}

func (o *ArchiveRunNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ArchiveRunNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArchiveRunForbidden creates a ArchiveRunForbidden with default headers values
func NewArchiveRunForbidden() *ArchiveRunForbidden {
	return &ArchiveRunForbidden{}
}

/*
ArchiveRunForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ArchiveRunForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this archive run forbidden response has a 2xx status code
func (o *ArchiveRunForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this archive run forbidden response has a 3xx status code
func (o *ArchiveRunForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive run forbidden response has a 4xx status code
func (o *ArchiveRunForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this archive run forbidden response has a 5xx status code
func (o *ArchiveRunForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this archive run forbidden response a status code equal to that given
func (o *ArchiveRunForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the archive run forbidden response
func (o *ArchiveRunForbidden) Code() int {
	return 403
}

func (o *ArchiveRunForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/archive][%d] archiveRunForbidden  %+v", 403, o.Payload)
}

func (o *ArchiveRunForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/archive][%d] archiveRunForbidden  %+v", 403, o.Payload)
}

func (o *ArchiveRunForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ArchiveRunForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArchiveRunNotFound creates a ArchiveRunNotFound with default headers values
func NewArchiveRunNotFound() *ArchiveRunNotFound {
	return &ArchiveRunNotFound{}
}

/*
ArchiveRunNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ArchiveRunNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this archive run not found response has a 2xx status code
func (o *ArchiveRunNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this archive run not found response has a 3xx status code
func (o *ArchiveRunNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive run not found response has a 4xx status code
func (o *ArchiveRunNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this archive run not found response has a 5xx status code
func (o *ArchiveRunNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this archive run not found response a status code equal to that given
func (o *ArchiveRunNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the archive run not found response
func (o *ArchiveRunNotFound) Code() int {
	return 404
}

func (o *ArchiveRunNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/archive][%d] archiveRunNotFound  %+v", 404, o.Payload)
}

func (o *ArchiveRunNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/archive][%d] archiveRunNotFound  %+v", 404, o.Payload)
}

func (o *ArchiveRunNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ArchiveRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArchiveRunDefault creates a ArchiveRunDefault with default headers values
func NewArchiveRunDefault(code int) *ArchiveRunDefault {
	return &ArchiveRunDefault{
		_statusCode: code,
	}
}

/*
ArchiveRunDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ArchiveRunDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this archive run default response has a 2xx status code
func (o *ArchiveRunDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this archive run default response has a 3xx status code
func (o *ArchiveRunDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this archive run default response has a 4xx status code
func (o *ArchiveRunDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this archive run default response has a 5xx status code
func (o *ArchiveRunDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this archive run default response a status code equal to that given
func (o *ArchiveRunDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the archive run default response
func (o *ArchiveRunDefault) Code() int {
	return o._statusCode
}

func (o *ArchiveRunDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/archive][%d] ArchiveRun default  %+v", o._statusCode, o.Payload)
}

func (o *ArchiveRunDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/archive][%d] ArchiveRun default  %+v", o._statusCode, o.Payload)
}

func (o *ArchiveRunDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ArchiveRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
