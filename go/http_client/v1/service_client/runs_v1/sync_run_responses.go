// Code generated by go-swagger; DO NOT EDIT.

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// SyncRunReader is a Reader for the SyncRun structure.
type SyncRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSyncRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewSyncRunNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSyncRunForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSyncRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewSyncRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSyncRunOK creates a SyncRunOK with default headers values
func NewSyncRunOK() *SyncRunOK {
	return &SyncRunOK{}
}

/*
SyncRunOK describes a response with status code 200, with default header values.

A successful response.
*/
type SyncRunOK struct {
}

// IsSuccess returns true when this sync run o k response has a 2xx status code
func (o *SyncRunOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sync run o k response has a 3xx status code
func (o *SyncRunOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync run o k response has a 4xx status code
func (o *SyncRunOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this sync run o k response has a 5xx status code
func (o *SyncRunOK) IsServerError() bool {
	return false
}

// IsCode returns true when this sync run o k response a status code equal to that given
func (o *SyncRunOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the sync run o k response
func (o *SyncRunOK) Code() int {
	return 200
}

func (o *SyncRunOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/sync][%d] syncRunOK ", 200)
}

func (o *SyncRunOK) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/sync][%d] syncRunOK ", 200)
}

func (o *SyncRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSyncRunNoContent creates a SyncRunNoContent with default headers values
func NewSyncRunNoContent() *SyncRunNoContent {
	return &SyncRunNoContent{}
}

/*
SyncRunNoContent describes a response with status code 204, with default header values.

No content.
*/
type SyncRunNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this sync run no content response has a 2xx status code
func (o *SyncRunNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sync run no content response has a 3xx status code
func (o *SyncRunNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync run no content response has a 4xx status code
func (o *SyncRunNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this sync run no content response has a 5xx status code
func (o *SyncRunNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this sync run no content response a status code equal to that given
func (o *SyncRunNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the sync run no content response
func (o *SyncRunNoContent) Code() int {
	return 204
}

func (o *SyncRunNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/sync][%d] syncRunNoContent  %+v", 204, o.Payload)
}

func (o *SyncRunNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/sync][%d] syncRunNoContent  %+v", 204, o.Payload)
}

func (o *SyncRunNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *SyncRunNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncRunForbidden creates a SyncRunForbidden with default headers values
func NewSyncRunForbidden() *SyncRunForbidden {
	return &SyncRunForbidden{}
}

/*
SyncRunForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type SyncRunForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this sync run forbidden response has a 2xx status code
func (o *SyncRunForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this sync run forbidden response has a 3xx status code
func (o *SyncRunForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync run forbidden response has a 4xx status code
func (o *SyncRunForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this sync run forbidden response has a 5xx status code
func (o *SyncRunForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this sync run forbidden response a status code equal to that given
func (o *SyncRunForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the sync run forbidden response
func (o *SyncRunForbidden) Code() int {
	return 403
}

func (o *SyncRunForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/sync][%d] syncRunForbidden  %+v", 403, o.Payload)
}

func (o *SyncRunForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/sync][%d] syncRunForbidden  %+v", 403, o.Payload)
}

func (o *SyncRunForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *SyncRunForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncRunNotFound creates a SyncRunNotFound with default headers values
func NewSyncRunNotFound() *SyncRunNotFound {
	return &SyncRunNotFound{}
}

/*
SyncRunNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type SyncRunNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this sync run not found response has a 2xx status code
func (o *SyncRunNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this sync run not found response has a 3xx status code
func (o *SyncRunNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync run not found response has a 4xx status code
func (o *SyncRunNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this sync run not found response has a 5xx status code
func (o *SyncRunNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this sync run not found response a status code equal to that given
func (o *SyncRunNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the sync run not found response
func (o *SyncRunNotFound) Code() int {
	return 404
}

func (o *SyncRunNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/sync][%d] syncRunNotFound  %+v", 404, o.Payload)
}

func (o *SyncRunNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/sync][%d] syncRunNotFound  %+v", 404, o.Payload)
}

func (o *SyncRunNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *SyncRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncRunDefault creates a SyncRunDefault with default headers values
func NewSyncRunDefault(code int) *SyncRunDefault {
	return &SyncRunDefault{
		_statusCode: code,
	}
}

/*
SyncRunDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SyncRunDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this sync run default response has a 2xx status code
func (o *SyncRunDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this sync run default response has a 3xx status code
func (o *SyncRunDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this sync run default response has a 4xx status code
func (o *SyncRunDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this sync run default response has a 5xx status code
func (o *SyncRunDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this sync run default response a status code equal to that given
func (o *SyncRunDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the sync run default response
func (o *SyncRunDefault) Code() int {
	return o._statusCode
}

func (o *SyncRunDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/sync][%d] SyncRun default  %+v", o._statusCode, o.Payload)
}

func (o *SyncRunDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/sync][%d] SyncRun default  %+v", o._statusCode, o.Payload)
}

func (o *SyncRunDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *SyncRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
