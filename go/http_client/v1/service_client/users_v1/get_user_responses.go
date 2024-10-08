// Code generated by go-swagger; DO NOT EDIT.

package users_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// GetUserReader is a Reader for the GetUser structure.
type GetUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUserOK creates a GetUserOK with default headers values
func NewGetUserOK() *GetUserOK {
	return &GetUserOK{}
}

/*
GetUserOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetUserOK struct {
	Payload *service_model.V1User
}

// IsSuccess returns true when this get user o k response has a 2xx status code
func (o *GetUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user o k response has a 3xx status code
func (o *GetUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user o k response has a 4xx status code
func (o *GetUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user o k response has a 5xx status code
func (o *GetUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user o k response a status code equal to that given
func (o *GetUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get user o k response
func (o *GetUserOK) Code() int {
	return 200
}

func (o *GetUserOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/users][%d] getUserOK  %+v", 200, o.Payload)
}

func (o *GetUserOK) String() string {
	return fmt.Sprintf("[GET /api/v1/users][%d] getUserOK  %+v", 200, o.Payload)
}

func (o *GetUserOK) GetPayload() *service_model.V1User {
	return o.Payload
}

func (o *GetUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserNoContent creates a GetUserNoContent with default headers values
func NewGetUserNoContent() *GetUserNoContent {
	return &GetUserNoContent{}
}

/*
GetUserNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetUserNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get user no content response has a 2xx status code
func (o *GetUserNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user no content response has a 3xx status code
func (o *GetUserNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user no content response has a 4xx status code
func (o *GetUserNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user no content response has a 5xx status code
func (o *GetUserNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get user no content response a status code equal to that given
func (o *GetUserNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get user no content response
func (o *GetUserNoContent) Code() int {
	return 204
}

func (o *GetUserNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/users][%d] getUserNoContent  %+v", 204, o.Payload)
}

func (o *GetUserNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/users][%d] getUserNoContent  %+v", 204, o.Payload)
}

func (o *GetUserNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserForbidden creates a GetUserForbidden with default headers values
func NewGetUserForbidden() *GetUserForbidden {
	return &GetUserForbidden{}
}

/*
GetUserForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetUserForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get user forbidden response has a 2xx status code
func (o *GetUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user forbidden response has a 3xx status code
func (o *GetUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user forbidden response has a 4xx status code
func (o *GetUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user forbidden response has a 5xx status code
func (o *GetUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user forbidden response a status code equal to that given
func (o *GetUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get user forbidden response
func (o *GetUserForbidden) Code() int {
	return 403
}

func (o *GetUserForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/users][%d] getUserForbidden  %+v", 403, o.Payload)
}

func (o *GetUserForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/users][%d] getUserForbidden  %+v", 403, o.Payload)
}

func (o *GetUserForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserNotFound creates a GetUserNotFound with default headers values
func NewGetUserNotFound() *GetUserNotFound {
	return &GetUserNotFound{}
}

/*
GetUserNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetUserNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get user not found response has a 2xx status code
func (o *GetUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user not found response has a 3xx status code
func (o *GetUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user not found response has a 4xx status code
func (o *GetUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user not found response has a 5xx status code
func (o *GetUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get user not found response a status code equal to that given
func (o *GetUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get user not found response
func (o *GetUserNotFound) Code() int {
	return 404
}

func (o *GetUserNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/users][%d] getUserNotFound  %+v", 404, o.Payload)
}

func (o *GetUserNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/users][%d] getUserNotFound  %+v", 404, o.Payload)
}

func (o *GetUserNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserDefault creates a GetUserDefault with default headers values
func NewGetUserDefault(code int) *GetUserDefault {
	return &GetUserDefault{
		_statusCode: code,
	}
}

/*
GetUserDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetUserDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this get user default response has a 2xx status code
func (o *GetUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get user default response has a 3xx status code
func (o *GetUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get user default response has a 4xx status code
func (o *GetUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get user default response has a 5xx status code
func (o *GetUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get user default response a status code equal to that given
func (o *GetUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get user default response
func (o *GetUserDefault) Code() int {
	return o._statusCode
}

func (o *GetUserDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/users][%d] GetUser default  %+v", o._statusCode, o.Payload)
}

func (o *GetUserDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/users][%d] GetUser default  %+v", o._statusCode, o.Payload)
}

func (o *GetUserDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
