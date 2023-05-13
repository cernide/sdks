// Code generated by go-swagger; DO NOT EDIT.

package auth_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// ResetPasswordConfirmReader is a Reader for the ResetPasswordConfirm structure.
type ResetPasswordConfirmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResetPasswordConfirmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResetPasswordConfirmOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewResetPasswordConfirmNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewResetPasswordConfirmForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewResetPasswordConfirmNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewResetPasswordConfirmDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResetPasswordConfirmOK creates a ResetPasswordConfirmOK with default headers values
func NewResetPasswordConfirmOK() *ResetPasswordConfirmOK {
	return &ResetPasswordConfirmOK{}
}

/*
ResetPasswordConfirmOK describes a response with status code 200, with default header values.

A successful response.
*/
type ResetPasswordConfirmOK struct {
	Payload *service_model.V1Auth
}

// IsSuccess returns true when this reset password confirm o k response has a 2xx status code
func (o *ResetPasswordConfirmOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this reset password confirm o k response has a 3xx status code
func (o *ResetPasswordConfirmOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reset password confirm o k response has a 4xx status code
func (o *ResetPasswordConfirmOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this reset password confirm o k response has a 5xx status code
func (o *ResetPasswordConfirmOK) IsServerError() bool {
	return false
}

// IsCode returns true when this reset password confirm o k response a status code equal to that given
func (o *ResetPasswordConfirmOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the reset password confirm o k response
func (o *ResetPasswordConfirmOK) Code() int {
	return 200
}

func (o *ResetPasswordConfirmOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/auth/reset-password-confirm][%d] resetPasswordConfirmOK  %+v", 200, o.Payload)
}

func (o *ResetPasswordConfirmOK) String() string {
	return fmt.Sprintf("[POST /api/v1/auth/reset-password-confirm][%d] resetPasswordConfirmOK  %+v", 200, o.Payload)
}

func (o *ResetPasswordConfirmOK) GetPayload() *service_model.V1Auth {
	return o.Payload
}

func (o *ResetPasswordConfirmOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Auth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetPasswordConfirmNoContent creates a ResetPasswordConfirmNoContent with default headers values
func NewResetPasswordConfirmNoContent() *ResetPasswordConfirmNoContent {
	return &ResetPasswordConfirmNoContent{}
}

/*
ResetPasswordConfirmNoContent describes a response with status code 204, with default header values.

No content.
*/
type ResetPasswordConfirmNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this reset password confirm no content response has a 2xx status code
func (o *ResetPasswordConfirmNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this reset password confirm no content response has a 3xx status code
func (o *ResetPasswordConfirmNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reset password confirm no content response has a 4xx status code
func (o *ResetPasswordConfirmNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this reset password confirm no content response has a 5xx status code
func (o *ResetPasswordConfirmNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this reset password confirm no content response a status code equal to that given
func (o *ResetPasswordConfirmNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the reset password confirm no content response
func (o *ResetPasswordConfirmNoContent) Code() int {
	return 204
}

func (o *ResetPasswordConfirmNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/auth/reset-password-confirm][%d] resetPasswordConfirmNoContent  %+v", 204, o.Payload)
}

func (o *ResetPasswordConfirmNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/auth/reset-password-confirm][%d] resetPasswordConfirmNoContent  %+v", 204, o.Payload)
}

func (o *ResetPasswordConfirmNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ResetPasswordConfirmNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetPasswordConfirmForbidden creates a ResetPasswordConfirmForbidden with default headers values
func NewResetPasswordConfirmForbidden() *ResetPasswordConfirmForbidden {
	return &ResetPasswordConfirmForbidden{}
}

/*
ResetPasswordConfirmForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ResetPasswordConfirmForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this reset password confirm forbidden response has a 2xx status code
func (o *ResetPasswordConfirmForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this reset password confirm forbidden response has a 3xx status code
func (o *ResetPasswordConfirmForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reset password confirm forbidden response has a 4xx status code
func (o *ResetPasswordConfirmForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this reset password confirm forbidden response has a 5xx status code
func (o *ResetPasswordConfirmForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this reset password confirm forbidden response a status code equal to that given
func (o *ResetPasswordConfirmForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the reset password confirm forbidden response
func (o *ResetPasswordConfirmForbidden) Code() int {
	return 403
}

func (o *ResetPasswordConfirmForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/auth/reset-password-confirm][%d] resetPasswordConfirmForbidden  %+v", 403, o.Payload)
}

func (o *ResetPasswordConfirmForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/auth/reset-password-confirm][%d] resetPasswordConfirmForbidden  %+v", 403, o.Payload)
}

func (o *ResetPasswordConfirmForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ResetPasswordConfirmForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetPasswordConfirmNotFound creates a ResetPasswordConfirmNotFound with default headers values
func NewResetPasswordConfirmNotFound() *ResetPasswordConfirmNotFound {
	return &ResetPasswordConfirmNotFound{}
}

/*
ResetPasswordConfirmNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ResetPasswordConfirmNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this reset password confirm not found response has a 2xx status code
func (o *ResetPasswordConfirmNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this reset password confirm not found response has a 3xx status code
func (o *ResetPasswordConfirmNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reset password confirm not found response has a 4xx status code
func (o *ResetPasswordConfirmNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this reset password confirm not found response has a 5xx status code
func (o *ResetPasswordConfirmNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this reset password confirm not found response a status code equal to that given
func (o *ResetPasswordConfirmNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the reset password confirm not found response
func (o *ResetPasswordConfirmNotFound) Code() int {
	return 404
}

func (o *ResetPasswordConfirmNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/auth/reset-password-confirm][%d] resetPasswordConfirmNotFound  %+v", 404, o.Payload)
}

func (o *ResetPasswordConfirmNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/auth/reset-password-confirm][%d] resetPasswordConfirmNotFound  %+v", 404, o.Payload)
}

func (o *ResetPasswordConfirmNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ResetPasswordConfirmNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetPasswordConfirmDefault creates a ResetPasswordConfirmDefault with default headers values
func NewResetPasswordConfirmDefault(code int) *ResetPasswordConfirmDefault {
	return &ResetPasswordConfirmDefault{
		_statusCode: code,
	}
}

/*
ResetPasswordConfirmDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ResetPasswordConfirmDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this reset password confirm default response has a 2xx status code
func (o *ResetPasswordConfirmDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this reset password confirm default response has a 3xx status code
func (o *ResetPasswordConfirmDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this reset password confirm default response has a 4xx status code
func (o *ResetPasswordConfirmDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this reset password confirm default response has a 5xx status code
func (o *ResetPasswordConfirmDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this reset password confirm default response a status code equal to that given
func (o *ResetPasswordConfirmDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the reset password confirm default response
func (o *ResetPasswordConfirmDefault) Code() int {
	return o._statusCode
}

func (o *ResetPasswordConfirmDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/auth/reset-password-confirm][%d] ResetPasswordConfirm default  %+v", o._statusCode, o.Payload)
}

func (o *ResetPasswordConfirmDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/auth/reset-password-confirm][%d] ResetPasswordConfirm default  %+v", o._statusCode, o.Payload)
}

func (o *ResetPasswordConfirmDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ResetPasswordConfirmDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
