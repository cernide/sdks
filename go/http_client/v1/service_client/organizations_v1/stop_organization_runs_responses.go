// Code generated by go-swagger; DO NOT EDIT.

package organizations_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// StopOrganizationRunsReader is a Reader for the StopOrganizationRuns structure.
type StopOrganizationRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopOrganizationRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopOrganizationRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewStopOrganizationRunsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewStopOrganizationRunsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopOrganizationRunsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewStopOrganizationRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStopOrganizationRunsOK creates a StopOrganizationRunsOK with default headers values
func NewStopOrganizationRunsOK() *StopOrganizationRunsOK {
	return &StopOrganizationRunsOK{}
}

/*
StopOrganizationRunsOK describes a response with status code 200, with default header values.

A successful response.
*/
type StopOrganizationRunsOK struct {
}

// IsSuccess returns true when this stop organization runs o k response has a 2xx status code
func (o *StopOrganizationRunsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop organization runs o k response has a 3xx status code
func (o *StopOrganizationRunsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop organization runs o k response has a 4xx status code
func (o *StopOrganizationRunsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop organization runs o k response has a 5xx status code
func (o *StopOrganizationRunsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stop organization runs o k response a status code equal to that given
func (o *StopOrganizationRunsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stop organization runs o k response
func (o *StopOrganizationRunsOK) Code() int {
	return 200
}

func (o *StopOrganizationRunsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/stop][%d] stopOrganizationRunsOK ", 200)
}

func (o *StopOrganizationRunsOK) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/stop][%d] stopOrganizationRunsOK ", 200)
}

func (o *StopOrganizationRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopOrganizationRunsNoContent creates a StopOrganizationRunsNoContent with default headers values
func NewStopOrganizationRunsNoContent() *StopOrganizationRunsNoContent {
	return &StopOrganizationRunsNoContent{}
}

/*
StopOrganizationRunsNoContent describes a response with status code 204, with default header values.

No content.
*/
type StopOrganizationRunsNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this stop organization runs no content response has a 2xx status code
func (o *StopOrganizationRunsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop organization runs no content response has a 3xx status code
func (o *StopOrganizationRunsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop organization runs no content response has a 4xx status code
func (o *StopOrganizationRunsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop organization runs no content response has a 5xx status code
func (o *StopOrganizationRunsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this stop organization runs no content response a status code equal to that given
func (o *StopOrganizationRunsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the stop organization runs no content response
func (o *StopOrganizationRunsNoContent) Code() int {
	return 204
}

func (o *StopOrganizationRunsNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/stop][%d] stopOrganizationRunsNoContent  %+v", 204, o.Payload)
}

func (o *StopOrganizationRunsNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/stop][%d] stopOrganizationRunsNoContent  %+v", 204, o.Payload)
}

func (o *StopOrganizationRunsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *StopOrganizationRunsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopOrganizationRunsForbidden creates a StopOrganizationRunsForbidden with default headers values
func NewStopOrganizationRunsForbidden() *StopOrganizationRunsForbidden {
	return &StopOrganizationRunsForbidden{}
}

/*
StopOrganizationRunsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type StopOrganizationRunsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this stop organization runs forbidden response has a 2xx status code
func (o *StopOrganizationRunsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop organization runs forbidden response has a 3xx status code
func (o *StopOrganizationRunsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop organization runs forbidden response has a 4xx status code
func (o *StopOrganizationRunsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop organization runs forbidden response has a 5xx status code
func (o *StopOrganizationRunsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stop organization runs forbidden response a status code equal to that given
func (o *StopOrganizationRunsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stop organization runs forbidden response
func (o *StopOrganizationRunsForbidden) Code() int {
	return 403
}

func (o *StopOrganizationRunsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/stop][%d] stopOrganizationRunsForbidden  %+v", 403, o.Payload)
}

func (o *StopOrganizationRunsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/stop][%d] stopOrganizationRunsForbidden  %+v", 403, o.Payload)
}

func (o *StopOrganizationRunsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *StopOrganizationRunsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopOrganizationRunsNotFound creates a StopOrganizationRunsNotFound with default headers values
func NewStopOrganizationRunsNotFound() *StopOrganizationRunsNotFound {
	return &StopOrganizationRunsNotFound{}
}

/*
StopOrganizationRunsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type StopOrganizationRunsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this stop organization runs not found response has a 2xx status code
func (o *StopOrganizationRunsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop organization runs not found response has a 3xx status code
func (o *StopOrganizationRunsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop organization runs not found response has a 4xx status code
func (o *StopOrganizationRunsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop organization runs not found response has a 5xx status code
func (o *StopOrganizationRunsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stop organization runs not found response a status code equal to that given
func (o *StopOrganizationRunsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stop organization runs not found response
func (o *StopOrganizationRunsNotFound) Code() int {
	return 404
}

func (o *StopOrganizationRunsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/stop][%d] stopOrganizationRunsNotFound  %+v", 404, o.Payload)
}

func (o *StopOrganizationRunsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/stop][%d] stopOrganizationRunsNotFound  %+v", 404, o.Payload)
}

func (o *StopOrganizationRunsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *StopOrganizationRunsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopOrganizationRunsDefault creates a StopOrganizationRunsDefault with default headers values
func NewStopOrganizationRunsDefault(code int) *StopOrganizationRunsDefault {
	return &StopOrganizationRunsDefault{
		_statusCode: code,
	}
}

/*
StopOrganizationRunsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StopOrganizationRunsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this stop organization runs default response has a 2xx status code
func (o *StopOrganizationRunsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this stop organization runs default response has a 3xx status code
func (o *StopOrganizationRunsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this stop organization runs default response has a 4xx status code
func (o *StopOrganizationRunsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this stop organization runs default response has a 5xx status code
func (o *StopOrganizationRunsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this stop organization runs default response a status code equal to that given
func (o *StopOrganizationRunsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the stop organization runs default response
func (o *StopOrganizationRunsDefault) Code() int {
	return o._statusCode
}

func (o *StopOrganizationRunsDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/stop][%d] StopOrganizationRuns default  %+v", o._statusCode, o.Payload)
}

func (o *StopOrganizationRunsDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/stop][%d] StopOrganizationRuns default  %+v", o._statusCode, o.Payload)
}

func (o *StopOrganizationRunsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *StopOrganizationRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
