// Code generated by go-swagger; DO NOT EDIT.

package agents_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// DeleteAgentReader is a Reader for the DeleteAgent structure.
type DeleteAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteAgentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteAgentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAgentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAgentOK creates a DeleteAgentOK with default headers values
func NewDeleteAgentOK() *DeleteAgentOK {
	return &DeleteAgentOK{}
}

/*
DeleteAgentOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteAgentOK struct {
}

// IsSuccess returns true when this delete agent o k response has a 2xx status code
func (o *DeleteAgentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete agent o k response has a 3xx status code
func (o *DeleteAgentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete agent o k response has a 4xx status code
func (o *DeleteAgentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete agent o k response has a 5xx status code
func (o *DeleteAgentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete agent o k response a status code equal to that given
func (o *DeleteAgentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete agent o k response
func (o *DeleteAgentOK) Code() int {
	return 200
}

func (o *DeleteAgentOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/agents/{uuid}][%d] deleteAgentOK ", 200)
}

func (o *DeleteAgentOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/agents/{uuid}][%d] deleteAgentOK ", 200)
}

func (o *DeleteAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAgentNoContent creates a DeleteAgentNoContent with default headers values
func NewDeleteAgentNoContent() *DeleteAgentNoContent {
	return &DeleteAgentNoContent{}
}

/*
DeleteAgentNoContent describes a response with status code 204, with default header values.

No content.
*/
type DeleteAgentNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this delete agent no content response has a 2xx status code
func (o *DeleteAgentNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete agent no content response has a 3xx status code
func (o *DeleteAgentNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete agent no content response has a 4xx status code
func (o *DeleteAgentNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete agent no content response has a 5xx status code
func (o *DeleteAgentNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete agent no content response a status code equal to that given
func (o *DeleteAgentNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete agent no content response
func (o *DeleteAgentNoContent) Code() int {
	return 204
}

func (o *DeleteAgentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/agents/{uuid}][%d] deleteAgentNoContent  %+v", 204, o.Payload)
}

func (o *DeleteAgentNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/agents/{uuid}][%d] deleteAgentNoContent  %+v", 204, o.Payload)
}

func (o *DeleteAgentNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteAgentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAgentForbidden creates a DeleteAgentForbidden with default headers values
func NewDeleteAgentForbidden() *DeleteAgentForbidden {
	return &DeleteAgentForbidden{}
}

/*
DeleteAgentForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type DeleteAgentForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this delete agent forbidden response has a 2xx status code
func (o *DeleteAgentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete agent forbidden response has a 3xx status code
func (o *DeleteAgentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete agent forbidden response has a 4xx status code
func (o *DeleteAgentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete agent forbidden response has a 5xx status code
func (o *DeleteAgentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete agent forbidden response a status code equal to that given
func (o *DeleteAgentForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete agent forbidden response
func (o *DeleteAgentForbidden) Code() int {
	return 403
}

func (o *DeleteAgentForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/agents/{uuid}][%d] deleteAgentForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAgentForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/agents/{uuid}][%d] deleteAgentForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAgentForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteAgentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAgentNotFound creates a DeleteAgentNotFound with default headers values
func NewDeleteAgentNotFound() *DeleteAgentNotFound {
	return &DeleteAgentNotFound{}
}

/*
DeleteAgentNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type DeleteAgentNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this delete agent not found response has a 2xx status code
func (o *DeleteAgentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete agent not found response has a 3xx status code
func (o *DeleteAgentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete agent not found response has a 4xx status code
func (o *DeleteAgentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete agent not found response has a 5xx status code
func (o *DeleteAgentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete agent not found response a status code equal to that given
func (o *DeleteAgentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete agent not found response
func (o *DeleteAgentNotFound) Code() int {
	return 404
}

func (o *DeleteAgentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/agents/{uuid}][%d] deleteAgentNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAgentNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/agents/{uuid}][%d] deleteAgentNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAgentNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteAgentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAgentDefault creates a DeleteAgentDefault with default headers values
func NewDeleteAgentDefault(code int) *DeleteAgentDefault {
	return &DeleteAgentDefault{
		_statusCode: code,
	}
}

/*
DeleteAgentDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteAgentDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this delete agent default response has a 2xx status code
func (o *DeleteAgentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete agent default response has a 3xx status code
func (o *DeleteAgentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete agent default response has a 4xx status code
func (o *DeleteAgentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete agent default response has a 5xx status code
func (o *DeleteAgentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete agent default response a status code equal to that given
func (o *DeleteAgentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete agent default response
func (o *DeleteAgentDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAgentDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/agents/{uuid}][%d] DeleteAgent default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAgentDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/agents/{uuid}][%d] DeleteAgent default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAgentDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *DeleteAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
