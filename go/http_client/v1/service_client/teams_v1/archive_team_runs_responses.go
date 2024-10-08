// Code generated by go-swagger; DO NOT EDIT.

package teams_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// ArchiveTeamRunsReader is a Reader for the ArchiveTeamRuns structure.
type ArchiveTeamRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArchiveTeamRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewArchiveTeamRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewArchiveTeamRunsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewArchiveTeamRunsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewArchiveTeamRunsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewArchiveTeamRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewArchiveTeamRunsOK creates a ArchiveTeamRunsOK with default headers values
func NewArchiveTeamRunsOK() *ArchiveTeamRunsOK {
	return &ArchiveTeamRunsOK{}
}

/*
ArchiveTeamRunsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ArchiveTeamRunsOK struct {
}

// IsSuccess returns true when this archive team runs o k response has a 2xx status code
func (o *ArchiveTeamRunsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this archive team runs o k response has a 3xx status code
func (o *ArchiveTeamRunsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive team runs o k response has a 4xx status code
func (o *ArchiveTeamRunsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this archive team runs o k response has a 5xx status code
func (o *ArchiveTeamRunsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this archive team runs o k response a status code equal to that given
func (o *ArchiveTeamRunsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the archive team runs o k response
func (o *ArchiveTeamRunsOK) Code() int {
	return 200
}

func (o *ArchiveTeamRunsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/archive][%d] archiveTeamRunsOK ", 200)
}

func (o *ArchiveTeamRunsOK) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/archive][%d] archiveTeamRunsOK ", 200)
}

func (o *ArchiveTeamRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewArchiveTeamRunsNoContent creates a ArchiveTeamRunsNoContent with default headers values
func NewArchiveTeamRunsNoContent() *ArchiveTeamRunsNoContent {
	return &ArchiveTeamRunsNoContent{}
}

/*
ArchiveTeamRunsNoContent describes a response with status code 204, with default header values.

No content.
*/
type ArchiveTeamRunsNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this archive team runs no content response has a 2xx status code
func (o *ArchiveTeamRunsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this archive team runs no content response has a 3xx status code
func (o *ArchiveTeamRunsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive team runs no content response has a 4xx status code
func (o *ArchiveTeamRunsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this archive team runs no content response has a 5xx status code
func (o *ArchiveTeamRunsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this archive team runs no content response a status code equal to that given
func (o *ArchiveTeamRunsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the archive team runs no content response
func (o *ArchiveTeamRunsNoContent) Code() int {
	return 204
}

func (o *ArchiveTeamRunsNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/archive][%d] archiveTeamRunsNoContent  %+v", 204, o.Payload)
}

func (o *ArchiveTeamRunsNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/archive][%d] archiveTeamRunsNoContent  %+v", 204, o.Payload)
}

func (o *ArchiveTeamRunsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ArchiveTeamRunsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArchiveTeamRunsForbidden creates a ArchiveTeamRunsForbidden with default headers values
func NewArchiveTeamRunsForbidden() *ArchiveTeamRunsForbidden {
	return &ArchiveTeamRunsForbidden{}
}

/*
ArchiveTeamRunsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ArchiveTeamRunsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this archive team runs forbidden response has a 2xx status code
func (o *ArchiveTeamRunsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this archive team runs forbidden response has a 3xx status code
func (o *ArchiveTeamRunsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive team runs forbidden response has a 4xx status code
func (o *ArchiveTeamRunsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this archive team runs forbidden response has a 5xx status code
func (o *ArchiveTeamRunsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this archive team runs forbidden response a status code equal to that given
func (o *ArchiveTeamRunsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the archive team runs forbidden response
func (o *ArchiveTeamRunsForbidden) Code() int {
	return 403
}

func (o *ArchiveTeamRunsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/archive][%d] archiveTeamRunsForbidden  %+v", 403, o.Payload)
}

func (o *ArchiveTeamRunsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/archive][%d] archiveTeamRunsForbidden  %+v", 403, o.Payload)
}

func (o *ArchiveTeamRunsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ArchiveTeamRunsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArchiveTeamRunsNotFound creates a ArchiveTeamRunsNotFound with default headers values
func NewArchiveTeamRunsNotFound() *ArchiveTeamRunsNotFound {
	return &ArchiveTeamRunsNotFound{}
}

/*
ArchiveTeamRunsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ArchiveTeamRunsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this archive team runs not found response has a 2xx status code
func (o *ArchiveTeamRunsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this archive team runs not found response has a 3xx status code
func (o *ArchiveTeamRunsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive team runs not found response has a 4xx status code
func (o *ArchiveTeamRunsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this archive team runs not found response has a 5xx status code
func (o *ArchiveTeamRunsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this archive team runs not found response a status code equal to that given
func (o *ArchiveTeamRunsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the archive team runs not found response
func (o *ArchiveTeamRunsNotFound) Code() int {
	return 404
}

func (o *ArchiveTeamRunsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/archive][%d] archiveTeamRunsNotFound  %+v", 404, o.Payload)
}

func (o *ArchiveTeamRunsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/archive][%d] archiveTeamRunsNotFound  %+v", 404, o.Payload)
}

func (o *ArchiveTeamRunsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ArchiveTeamRunsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArchiveTeamRunsDefault creates a ArchiveTeamRunsDefault with default headers values
func NewArchiveTeamRunsDefault(code int) *ArchiveTeamRunsDefault {
	return &ArchiveTeamRunsDefault{
		_statusCode: code,
	}
}

/*
ArchiveTeamRunsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ArchiveTeamRunsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this archive team runs default response has a 2xx status code
func (o *ArchiveTeamRunsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this archive team runs default response has a 3xx status code
func (o *ArchiveTeamRunsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this archive team runs default response has a 4xx status code
func (o *ArchiveTeamRunsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this archive team runs default response has a 5xx status code
func (o *ArchiveTeamRunsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this archive team runs default response a status code equal to that given
func (o *ArchiveTeamRunsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the archive team runs default response
func (o *ArchiveTeamRunsDefault) Code() int {
	return o._statusCode
}

func (o *ArchiveTeamRunsDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/archive][%d] ArchiveTeamRuns default  %+v", o._statusCode, o.Payload)
}

func (o *ArchiveTeamRunsDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/archive][%d] ArchiveTeamRuns default  %+v", o._statusCode, o.Payload)
}

func (o *ArchiveTeamRunsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ArchiveTeamRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
