// Code generated by go-swagger; DO NOT EDIT.

package service_accounts_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/go/http_client/v1/service_model"
)

// DeleteServiceAccountTokenReader is a Reader for the DeleteServiceAccountToken structure.
type DeleteServiceAccountTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServiceAccountTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteServiceAccountTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteServiceAccountTokenNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteServiceAccountTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteServiceAccountTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteServiceAccountTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteServiceAccountTokenOK creates a DeleteServiceAccountTokenOK with default headers values
func NewDeleteServiceAccountTokenOK() *DeleteServiceAccountTokenOK {
	return &DeleteServiceAccountTokenOK{}
}

/*
DeleteServiceAccountTokenOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteServiceAccountTokenOK struct {
}

// IsSuccess returns true when this delete service account token o k response has a 2xx status code
func (o *DeleteServiceAccountTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete service account token o k response has a 3xx status code
func (o *DeleteServiceAccountTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service account token o k response has a 4xx status code
func (o *DeleteServiceAccountTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete service account token o k response has a 5xx status code
func (o *DeleteServiceAccountTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete service account token o k response a status code equal to that given
func (o *DeleteServiceAccountTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete service account token o k response
func (o *DeleteServiceAccountTokenOK) Code() int {
	return 200
}

func (o *DeleteServiceAccountTokenOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/sa/{entity}/tokens/{uuid}][%d] deleteServiceAccountTokenOK ", 200)
}

func (o *DeleteServiceAccountTokenOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/sa/{entity}/tokens/{uuid}][%d] deleteServiceAccountTokenOK ", 200)
}

func (o *DeleteServiceAccountTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServiceAccountTokenNoContent creates a DeleteServiceAccountTokenNoContent with default headers values
func NewDeleteServiceAccountTokenNoContent() *DeleteServiceAccountTokenNoContent {
	return &DeleteServiceAccountTokenNoContent{}
}

/*
DeleteServiceAccountTokenNoContent describes a response with status code 204, with default header values.

No content.
*/
type DeleteServiceAccountTokenNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this delete service account token no content response has a 2xx status code
func (o *DeleteServiceAccountTokenNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete service account token no content response has a 3xx status code
func (o *DeleteServiceAccountTokenNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service account token no content response has a 4xx status code
func (o *DeleteServiceAccountTokenNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete service account token no content response has a 5xx status code
func (o *DeleteServiceAccountTokenNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete service account token no content response a status code equal to that given
func (o *DeleteServiceAccountTokenNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete service account token no content response
func (o *DeleteServiceAccountTokenNoContent) Code() int {
	return 204
}

func (o *DeleteServiceAccountTokenNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/sa/{entity}/tokens/{uuid}][%d] deleteServiceAccountTokenNoContent  %+v", 204, o.Payload)
}

func (o *DeleteServiceAccountTokenNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/sa/{entity}/tokens/{uuid}][%d] deleteServiceAccountTokenNoContent  %+v", 204, o.Payload)
}

func (o *DeleteServiceAccountTokenNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteServiceAccountTokenNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServiceAccountTokenForbidden creates a DeleteServiceAccountTokenForbidden with default headers values
func NewDeleteServiceAccountTokenForbidden() *DeleteServiceAccountTokenForbidden {
	return &DeleteServiceAccountTokenForbidden{}
}

/*
DeleteServiceAccountTokenForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type DeleteServiceAccountTokenForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this delete service account token forbidden response has a 2xx status code
func (o *DeleteServiceAccountTokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete service account token forbidden response has a 3xx status code
func (o *DeleteServiceAccountTokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service account token forbidden response has a 4xx status code
func (o *DeleteServiceAccountTokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete service account token forbidden response has a 5xx status code
func (o *DeleteServiceAccountTokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete service account token forbidden response a status code equal to that given
func (o *DeleteServiceAccountTokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete service account token forbidden response
func (o *DeleteServiceAccountTokenForbidden) Code() int {
	return 403
}

func (o *DeleteServiceAccountTokenForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/sa/{entity}/tokens/{uuid}][%d] deleteServiceAccountTokenForbidden  %+v", 403, o.Payload)
}

func (o *DeleteServiceAccountTokenForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/sa/{entity}/tokens/{uuid}][%d] deleteServiceAccountTokenForbidden  %+v", 403, o.Payload)
}

func (o *DeleteServiceAccountTokenForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteServiceAccountTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServiceAccountTokenNotFound creates a DeleteServiceAccountTokenNotFound with default headers values
func NewDeleteServiceAccountTokenNotFound() *DeleteServiceAccountTokenNotFound {
	return &DeleteServiceAccountTokenNotFound{}
}

/*
DeleteServiceAccountTokenNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type DeleteServiceAccountTokenNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this delete service account token not found response has a 2xx status code
func (o *DeleteServiceAccountTokenNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete service account token not found response has a 3xx status code
func (o *DeleteServiceAccountTokenNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service account token not found response has a 4xx status code
func (o *DeleteServiceAccountTokenNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete service account token not found response has a 5xx status code
func (o *DeleteServiceAccountTokenNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete service account token not found response a status code equal to that given
func (o *DeleteServiceAccountTokenNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete service account token not found response
func (o *DeleteServiceAccountTokenNotFound) Code() int {
	return 404
}

func (o *DeleteServiceAccountTokenNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/sa/{entity}/tokens/{uuid}][%d] deleteServiceAccountTokenNotFound  %+v", 404, o.Payload)
}

func (o *DeleteServiceAccountTokenNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/sa/{entity}/tokens/{uuid}][%d] deleteServiceAccountTokenNotFound  %+v", 404, o.Payload)
}

func (o *DeleteServiceAccountTokenNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteServiceAccountTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServiceAccountTokenDefault creates a DeleteServiceAccountTokenDefault with default headers values
func NewDeleteServiceAccountTokenDefault(code int) *DeleteServiceAccountTokenDefault {
	return &DeleteServiceAccountTokenDefault{
		_statusCode: code,
	}
}

/*
DeleteServiceAccountTokenDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteServiceAccountTokenDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this delete service account token default response has a 2xx status code
func (o *DeleteServiceAccountTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete service account token default response has a 3xx status code
func (o *DeleteServiceAccountTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete service account token default response has a 4xx status code
func (o *DeleteServiceAccountTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete service account token default response has a 5xx status code
func (o *DeleteServiceAccountTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete service account token default response a status code equal to that given
func (o *DeleteServiceAccountTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete service account token default response
func (o *DeleteServiceAccountTokenDefault) Code() int {
	return o._statusCode
}

func (o *DeleteServiceAccountTokenDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/sa/{entity}/tokens/{uuid}][%d] DeleteServiceAccountToken default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteServiceAccountTokenDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/sa/{entity}/tokens/{uuid}][%d] DeleteServiceAccountToken default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteServiceAccountTokenDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *DeleteServiceAccountTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
