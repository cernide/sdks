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

// DeleteRunsReader is a Reader for the DeleteRuns structure.
type DeleteRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteRunsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteRunsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRunsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteRunsOK creates a DeleteRunsOK with default headers values
func NewDeleteRunsOK() *DeleteRunsOK {
	return &DeleteRunsOK{}
}

/*
DeleteRunsOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteRunsOK struct {
}

// IsSuccess returns true when this delete runs o k response has a 2xx status code
func (o *DeleteRunsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete runs o k response has a 3xx status code
func (o *DeleteRunsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete runs o k response has a 4xx status code
func (o *DeleteRunsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete runs o k response has a 5xx status code
func (o *DeleteRunsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete runs o k response a status code equal to that given
func (o *DeleteRunsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete runs o k response
func (o *DeleteRunsOK) Code() int {
	return 200
}

func (o *DeleteRunsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/delete][%d] deleteRunsOK ", 200)
}

func (o *DeleteRunsOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/delete][%d] deleteRunsOK ", 200)
}

func (o *DeleteRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRunsNoContent creates a DeleteRunsNoContent with default headers values
func NewDeleteRunsNoContent() *DeleteRunsNoContent {
	return &DeleteRunsNoContent{}
}

/*
DeleteRunsNoContent describes a response with status code 204, with default header values.

No content.
*/
type DeleteRunsNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this delete runs no content response has a 2xx status code
func (o *DeleteRunsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete runs no content response has a 3xx status code
func (o *DeleteRunsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete runs no content response has a 4xx status code
func (o *DeleteRunsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete runs no content response has a 5xx status code
func (o *DeleteRunsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete runs no content response a status code equal to that given
func (o *DeleteRunsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete runs no content response
func (o *DeleteRunsNoContent) Code() int {
	return 204
}

func (o *DeleteRunsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/delete][%d] deleteRunsNoContent  %+v", 204, o.Payload)
}

func (o *DeleteRunsNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/delete][%d] deleteRunsNoContent  %+v", 204, o.Payload)
}

func (o *DeleteRunsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteRunsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRunsForbidden creates a DeleteRunsForbidden with default headers values
func NewDeleteRunsForbidden() *DeleteRunsForbidden {
	return &DeleteRunsForbidden{}
}

/*
DeleteRunsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type DeleteRunsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this delete runs forbidden response has a 2xx status code
func (o *DeleteRunsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete runs forbidden response has a 3xx status code
func (o *DeleteRunsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete runs forbidden response has a 4xx status code
func (o *DeleteRunsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete runs forbidden response has a 5xx status code
func (o *DeleteRunsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete runs forbidden response a status code equal to that given
func (o *DeleteRunsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete runs forbidden response
func (o *DeleteRunsForbidden) Code() int {
	return 403
}

func (o *DeleteRunsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/delete][%d] deleteRunsForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRunsForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/delete][%d] deleteRunsForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRunsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteRunsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRunsNotFound creates a DeleteRunsNotFound with default headers values
func NewDeleteRunsNotFound() *DeleteRunsNotFound {
	return &DeleteRunsNotFound{}
}

/*
DeleteRunsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type DeleteRunsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this delete runs not found response has a 2xx status code
func (o *DeleteRunsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete runs not found response has a 3xx status code
func (o *DeleteRunsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete runs not found response has a 4xx status code
func (o *DeleteRunsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete runs not found response has a 5xx status code
func (o *DeleteRunsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete runs not found response a status code equal to that given
func (o *DeleteRunsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete runs not found response
func (o *DeleteRunsNotFound) Code() int {
	return 404
}

func (o *DeleteRunsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/delete][%d] deleteRunsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRunsNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/delete][%d] deleteRunsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRunsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteRunsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRunsDefault creates a DeleteRunsDefault with default headers values
func NewDeleteRunsDefault(code int) *DeleteRunsDefault {
	return &DeleteRunsDefault{
		_statusCode: code,
	}
}

/*
DeleteRunsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteRunsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this delete runs default response has a 2xx status code
func (o *DeleteRunsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete runs default response has a 3xx status code
func (o *DeleteRunsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete runs default response has a 4xx status code
func (o *DeleteRunsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete runs default response has a 5xx status code
func (o *DeleteRunsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete runs default response a status code equal to that given
func (o *DeleteRunsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete runs default response
func (o *DeleteRunsDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRunsDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/delete][%d] DeleteRuns default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRunsDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/delete][%d] DeleteRuns default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRunsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *DeleteRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
