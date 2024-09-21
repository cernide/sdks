// Code generated by go-swagger; DO NOT EDIT.

package service_accounts_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// DeleteServiceAccountReader is a Reader for the DeleteServiceAccount structure.
type DeleteServiceAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServiceAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteServiceAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteServiceAccountNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteServiceAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteServiceAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteServiceAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteServiceAccountOK creates a DeleteServiceAccountOK with default headers values
func NewDeleteServiceAccountOK() *DeleteServiceAccountOK {
	return &DeleteServiceAccountOK{}
}

/*
DeleteServiceAccountOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteServiceAccountOK struct {
}

// IsSuccess returns true when this delete service account o k response has a 2xx status code
func (o *DeleteServiceAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete service account o k response has a 3xx status code
func (o *DeleteServiceAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service account o k response has a 4xx status code
func (o *DeleteServiceAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete service account o k response has a 5xx status code
func (o *DeleteServiceAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete service account o k response a status code equal to that given
func (o *DeleteServiceAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete service account o k response
func (o *DeleteServiceAccountOK) Code() int {
	return 200
}

func (o *DeleteServiceAccountOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/sa/{uuid}][%d] deleteServiceAccountOK ", 200)
}

func (o *DeleteServiceAccountOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/sa/{uuid}][%d] deleteServiceAccountOK ", 200)
}

func (o *DeleteServiceAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServiceAccountNoContent creates a DeleteServiceAccountNoContent with default headers values
func NewDeleteServiceAccountNoContent() *DeleteServiceAccountNoContent {
	return &DeleteServiceAccountNoContent{}
}

/*
DeleteServiceAccountNoContent describes a response with status code 204, with default header values.

No content.
*/
type DeleteServiceAccountNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this delete service account no content response has a 2xx status code
func (o *DeleteServiceAccountNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete service account no content response has a 3xx status code
func (o *DeleteServiceAccountNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service account no content response has a 4xx status code
func (o *DeleteServiceAccountNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete service account no content response has a 5xx status code
func (o *DeleteServiceAccountNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete service account no content response a status code equal to that given
func (o *DeleteServiceAccountNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete service account no content response
func (o *DeleteServiceAccountNoContent) Code() int {
	return 204
}

func (o *DeleteServiceAccountNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/sa/{uuid}][%d] deleteServiceAccountNoContent  %+v", 204, o.Payload)
}

func (o *DeleteServiceAccountNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/sa/{uuid}][%d] deleteServiceAccountNoContent  %+v", 204, o.Payload)
}

func (o *DeleteServiceAccountNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteServiceAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServiceAccountForbidden creates a DeleteServiceAccountForbidden with default headers values
func NewDeleteServiceAccountForbidden() *DeleteServiceAccountForbidden {
	return &DeleteServiceAccountForbidden{}
}

/*
DeleteServiceAccountForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type DeleteServiceAccountForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this delete service account forbidden response has a 2xx status code
func (o *DeleteServiceAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete service account forbidden response has a 3xx status code
func (o *DeleteServiceAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service account forbidden response has a 4xx status code
func (o *DeleteServiceAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete service account forbidden response has a 5xx status code
func (o *DeleteServiceAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete service account forbidden response a status code equal to that given
func (o *DeleteServiceAccountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete service account forbidden response
func (o *DeleteServiceAccountForbidden) Code() int {
	return 403
}

func (o *DeleteServiceAccountForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/sa/{uuid}][%d] deleteServiceAccountForbidden  %+v", 403, o.Payload)
}

func (o *DeleteServiceAccountForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/sa/{uuid}][%d] deleteServiceAccountForbidden  %+v", 403, o.Payload)
}

func (o *DeleteServiceAccountForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteServiceAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServiceAccountNotFound creates a DeleteServiceAccountNotFound with default headers values
func NewDeleteServiceAccountNotFound() *DeleteServiceAccountNotFound {
	return &DeleteServiceAccountNotFound{}
}

/*
DeleteServiceAccountNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type DeleteServiceAccountNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this delete service account not found response has a 2xx status code
func (o *DeleteServiceAccountNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete service account not found response has a 3xx status code
func (o *DeleteServiceAccountNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service account not found response has a 4xx status code
func (o *DeleteServiceAccountNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete service account not found response has a 5xx status code
func (o *DeleteServiceAccountNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete service account not found response a status code equal to that given
func (o *DeleteServiceAccountNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete service account not found response
func (o *DeleteServiceAccountNotFound) Code() int {
	return 404
}

func (o *DeleteServiceAccountNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/sa/{uuid}][%d] deleteServiceAccountNotFound  %+v", 404, o.Payload)
}

func (o *DeleteServiceAccountNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/sa/{uuid}][%d] deleteServiceAccountNotFound  %+v", 404, o.Payload)
}

func (o *DeleteServiceAccountNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteServiceAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServiceAccountDefault creates a DeleteServiceAccountDefault with default headers values
func NewDeleteServiceAccountDefault(code int) *DeleteServiceAccountDefault {
	return &DeleteServiceAccountDefault{
		_statusCode: code,
	}
}

/*
DeleteServiceAccountDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteServiceAccountDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this delete service account default response has a 2xx status code
func (o *DeleteServiceAccountDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete service account default response has a 3xx status code
func (o *DeleteServiceAccountDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete service account default response has a 4xx status code
func (o *DeleteServiceAccountDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete service account default response has a 5xx status code
func (o *DeleteServiceAccountDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete service account default response a status code equal to that given
func (o *DeleteServiceAccountDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete service account default response
func (o *DeleteServiceAccountDefault) Code() int {
	return o._statusCode
}

func (o *DeleteServiceAccountDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/sa/{uuid}][%d] DeleteServiceAccount default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteServiceAccountDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/sa/{uuid}][%d] DeleteServiceAccount default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteServiceAccountDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *DeleteServiceAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
