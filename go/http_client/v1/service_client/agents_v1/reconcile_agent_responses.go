// Code generated by go-swagger; DO NOT EDIT.

package agents_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// ReconcileAgentReader is a Reader for the ReconcileAgent structure.
type ReconcileAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReconcileAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReconcileAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewReconcileAgentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewReconcileAgentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReconcileAgentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReconcileAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReconcileAgentOK creates a ReconcileAgentOK with default headers values
func NewReconcileAgentOK() *ReconcileAgentOK {
	return &ReconcileAgentOK{}
}

/*
ReconcileAgentOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReconcileAgentOK struct {
	Payload interface{}
}

// IsSuccess returns true when this reconcile agent o k response has a 2xx status code
func (o *ReconcileAgentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this reconcile agent o k response has a 3xx status code
func (o *ReconcileAgentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reconcile agent o k response has a 4xx status code
func (o *ReconcileAgentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this reconcile agent o k response has a 5xx status code
func (o *ReconcileAgentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this reconcile agent o k response a status code equal to that given
func (o *ReconcileAgentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the reconcile agent o k response
func (o *ReconcileAgentOK) Code() int {
	return 200
}

func (o *ReconcileAgentOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{uuid}/reconcile][%d] reconcileAgentOK  %+v", 200, o.Payload)
}

func (o *ReconcileAgentOK) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{uuid}/reconcile][%d] reconcileAgentOK  %+v", 200, o.Payload)
}

func (o *ReconcileAgentOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ReconcileAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReconcileAgentNoContent creates a ReconcileAgentNoContent with default headers values
func NewReconcileAgentNoContent() *ReconcileAgentNoContent {
	return &ReconcileAgentNoContent{}
}

/*
ReconcileAgentNoContent describes a response with status code 204, with default header values.

No content.
*/
type ReconcileAgentNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this reconcile agent no content response has a 2xx status code
func (o *ReconcileAgentNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this reconcile agent no content response has a 3xx status code
func (o *ReconcileAgentNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reconcile agent no content response has a 4xx status code
func (o *ReconcileAgentNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this reconcile agent no content response has a 5xx status code
func (o *ReconcileAgentNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this reconcile agent no content response a status code equal to that given
func (o *ReconcileAgentNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the reconcile agent no content response
func (o *ReconcileAgentNoContent) Code() int {
	return 204
}

func (o *ReconcileAgentNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{uuid}/reconcile][%d] reconcileAgentNoContent  %+v", 204, o.Payload)
}

func (o *ReconcileAgentNoContent) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{uuid}/reconcile][%d] reconcileAgentNoContent  %+v", 204, o.Payload)
}

func (o *ReconcileAgentNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ReconcileAgentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReconcileAgentForbidden creates a ReconcileAgentForbidden with default headers values
func NewReconcileAgentForbidden() *ReconcileAgentForbidden {
	return &ReconcileAgentForbidden{}
}

/*
ReconcileAgentForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ReconcileAgentForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this reconcile agent forbidden response has a 2xx status code
func (o *ReconcileAgentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this reconcile agent forbidden response has a 3xx status code
func (o *ReconcileAgentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reconcile agent forbidden response has a 4xx status code
func (o *ReconcileAgentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this reconcile agent forbidden response has a 5xx status code
func (o *ReconcileAgentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this reconcile agent forbidden response a status code equal to that given
func (o *ReconcileAgentForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the reconcile agent forbidden response
func (o *ReconcileAgentForbidden) Code() int {
	return 403
}

func (o *ReconcileAgentForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{uuid}/reconcile][%d] reconcileAgentForbidden  %+v", 403, o.Payload)
}

func (o *ReconcileAgentForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{uuid}/reconcile][%d] reconcileAgentForbidden  %+v", 403, o.Payload)
}

func (o *ReconcileAgentForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ReconcileAgentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReconcileAgentNotFound creates a ReconcileAgentNotFound with default headers values
func NewReconcileAgentNotFound() *ReconcileAgentNotFound {
	return &ReconcileAgentNotFound{}
}

/*
ReconcileAgentNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ReconcileAgentNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this reconcile agent not found response has a 2xx status code
func (o *ReconcileAgentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this reconcile agent not found response has a 3xx status code
func (o *ReconcileAgentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reconcile agent not found response has a 4xx status code
func (o *ReconcileAgentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this reconcile agent not found response has a 5xx status code
func (o *ReconcileAgentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this reconcile agent not found response a status code equal to that given
func (o *ReconcileAgentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the reconcile agent not found response
func (o *ReconcileAgentNotFound) Code() int {
	return 404
}

func (o *ReconcileAgentNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{uuid}/reconcile][%d] reconcileAgentNotFound  %+v", 404, o.Payload)
}

func (o *ReconcileAgentNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{uuid}/reconcile][%d] reconcileAgentNotFound  %+v", 404, o.Payload)
}

func (o *ReconcileAgentNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ReconcileAgentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReconcileAgentDefault creates a ReconcileAgentDefault with default headers values
func NewReconcileAgentDefault(code int) *ReconcileAgentDefault {
	return &ReconcileAgentDefault{
		_statusCode: code,
	}
}

/*
ReconcileAgentDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ReconcileAgentDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this reconcile agent default response has a 2xx status code
func (o *ReconcileAgentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this reconcile agent default response has a 3xx status code
func (o *ReconcileAgentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this reconcile agent default response has a 4xx status code
func (o *ReconcileAgentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this reconcile agent default response has a 5xx status code
func (o *ReconcileAgentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this reconcile agent default response a status code equal to that given
func (o *ReconcileAgentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the reconcile agent default response
func (o *ReconcileAgentDefault) Code() int {
	return o._statusCode
}

func (o *ReconcileAgentDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{uuid}/reconcile][%d] ReconcileAgent default  %+v", o._statusCode, o.Payload)
}

func (o *ReconcileAgentDefault) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{uuid}/reconcile][%d] ReconcileAgent default  %+v", o._statusCode, o.Payload)
}

func (o *ReconcileAgentDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ReconcileAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
