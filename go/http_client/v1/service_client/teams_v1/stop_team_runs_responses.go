// Code generated by go-swagger; DO NOT EDIT.

package teams_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/go/http_client/v1/service_model"
)

// StopTeamRunsReader is a Reader for the StopTeamRuns structure.
type StopTeamRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopTeamRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopTeamRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewStopTeamRunsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewStopTeamRunsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopTeamRunsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewStopTeamRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStopTeamRunsOK creates a StopTeamRunsOK with default headers values
func NewStopTeamRunsOK() *StopTeamRunsOK {
	return &StopTeamRunsOK{}
}

/*
StopTeamRunsOK describes a response with status code 200, with default header values.

A successful response.
*/
type StopTeamRunsOK struct {
}

// IsSuccess returns true when this stop team runs o k response has a 2xx status code
func (o *StopTeamRunsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop team runs o k response has a 3xx status code
func (o *StopTeamRunsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop team runs o k response has a 4xx status code
func (o *StopTeamRunsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop team runs o k response has a 5xx status code
func (o *StopTeamRunsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stop team runs o k response a status code equal to that given
func (o *StopTeamRunsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stop team runs o k response
func (o *StopTeamRunsOK) Code() int {
	return 200
}

func (o *StopTeamRunsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/stop][%d] stopTeamRunsOK ", 200)
}

func (o *StopTeamRunsOK) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/stop][%d] stopTeamRunsOK ", 200)
}

func (o *StopTeamRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopTeamRunsNoContent creates a StopTeamRunsNoContent with default headers values
func NewStopTeamRunsNoContent() *StopTeamRunsNoContent {
	return &StopTeamRunsNoContent{}
}

/*
StopTeamRunsNoContent describes a response with status code 204, with default header values.

No content.
*/
type StopTeamRunsNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this stop team runs no content response has a 2xx status code
func (o *StopTeamRunsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop team runs no content response has a 3xx status code
func (o *StopTeamRunsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop team runs no content response has a 4xx status code
func (o *StopTeamRunsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop team runs no content response has a 5xx status code
func (o *StopTeamRunsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this stop team runs no content response a status code equal to that given
func (o *StopTeamRunsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the stop team runs no content response
func (o *StopTeamRunsNoContent) Code() int {
	return 204
}

func (o *StopTeamRunsNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/stop][%d] stopTeamRunsNoContent  %+v", 204, o.Payload)
}

func (o *StopTeamRunsNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/stop][%d] stopTeamRunsNoContent  %+v", 204, o.Payload)
}

func (o *StopTeamRunsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *StopTeamRunsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopTeamRunsForbidden creates a StopTeamRunsForbidden with default headers values
func NewStopTeamRunsForbidden() *StopTeamRunsForbidden {
	return &StopTeamRunsForbidden{}
}

/*
StopTeamRunsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type StopTeamRunsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this stop team runs forbidden response has a 2xx status code
func (o *StopTeamRunsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop team runs forbidden response has a 3xx status code
func (o *StopTeamRunsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop team runs forbidden response has a 4xx status code
func (o *StopTeamRunsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop team runs forbidden response has a 5xx status code
func (o *StopTeamRunsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stop team runs forbidden response a status code equal to that given
func (o *StopTeamRunsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stop team runs forbidden response
func (o *StopTeamRunsForbidden) Code() int {
	return 403
}

func (o *StopTeamRunsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/stop][%d] stopTeamRunsForbidden  %+v", 403, o.Payload)
}

func (o *StopTeamRunsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/stop][%d] stopTeamRunsForbidden  %+v", 403, o.Payload)
}

func (o *StopTeamRunsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *StopTeamRunsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopTeamRunsNotFound creates a StopTeamRunsNotFound with default headers values
func NewStopTeamRunsNotFound() *StopTeamRunsNotFound {
	return &StopTeamRunsNotFound{}
}

/*
StopTeamRunsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type StopTeamRunsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this stop team runs not found response has a 2xx status code
func (o *StopTeamRunsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop team runs not found response has a 3xx status code
func (o *StopTeamRunsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop team runs not found response has a 4xx status code
func (o *StopTeamRunsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop team runs not found response has a 5xx status code
func (o *StopTeamRunsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stop team runs not found response a status code equal to that given
func (o *StopTeamRunsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stop team runs not found response
func (o *StopTeamRunsNotFound) Code() int {
	return 404
}

func (o *StopTeamRunsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/stop][%d] stopTeamRunsNotFound  %+v", 404, o.Payload)
}

func (o *StopTeamRunsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/stop][%d] stopTeamRunsNotFound  %+v", 404, o.Payload)
}

func (o *StopTeamRunsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *StopTeamRunsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopTeamRunsDefault creates a StopTeamRunsDefault with default headers values
func NewStopTeamRunsDefault(code int) *StopTeamRunsDefault {
	return &StopTeamRunsDefault{
		_statusCode: code,
	}
}

/*
StopTeamRunsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StopTeamRunsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this stop team runs default response has a 2xx status code
func (o *StopTeamRunsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this stop team runs default response has a 3xx status code
func (o *StopTeamRunsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this stop team runs default response has a 4xx status code
func (o *StopTeamRunsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this stop team runs default response has a 5xx status code
func (o *StopTeamRunsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this stop team runs default response a status code equal to that given
func (o *StopTeamRunsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the stop team runs default response
func (o *StopTeamRunsDefault) Code() int {
	return o._statusCode
}

func (o *StopTeamRunsDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/stop][%d] StopTeamRuns default  %+v", o._statusCode, o.Payload)
}

func (o *StopTeamRunsDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/stop][%d] StopTeamRuns default  %+v", o._statusCode, o.Payload)
}

func (o *StopTeamRunsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *StopTeamRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
