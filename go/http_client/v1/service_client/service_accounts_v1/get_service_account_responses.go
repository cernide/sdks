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

// GetServiceAccountReader is a Reader for the GetServiceAccount structure.
type GetServiceAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetServiceAccountNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetServiceAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetServiceAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetServiceAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServiceAccountOK creates a GetServiceAccountOK with default headers values
func NewGetServiceAccountOK() *GetServiceAccountOK {
	return &GetServiceAccountOK{}
}

/*
GetServiceAccountOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetServiceAccountOK struct {
	Payload *service_model.V1ServiceAccount
}

// IsSuccess returns true when this get service account o k response has a 2xx status code
func (o *GetServiceAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get service account o k response has a 3xx status code
func (o *GetServiceAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get service account o k response has a 4xx status code
func (o *GetServiceAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get service account o k response has a 5xx status code
func (o *GetServiceAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get service account o k response a status code equal to that given
func (o *GetServiceAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get service account o k response
func (o *GetServiceAccountOK) Code() int {
	return 200
}

func (o *GetServiceAccountOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/sa/{uuid}][%d] getServiceAccountOK  %+v", 200, o.Payload)
}

func (o *GetServiceAccountOK) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/sa/{uuid}][%d] getServiceAccountOK  %+v", 200, o.Payload)
}

func (o *GetServiceAccountOK) GetPayload() *service_model.V1ServiceAccount {
	return o.Payload
}

func (o *GetServiceAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ServiceAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceAccountNoContent creates a GetServiceAccountNoContent with default headers values
func NewGetServiceAccountNoContent() *GetServiceAccountNoContent {
	return &GetServiceAccountNoContent{}
}

/*
GetServiceAccountNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetServiceAccountNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get service account no content response has a 2xx status code
func (o *GetServiceAccountNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get service account no content response has a 3xx status code
func (o *GetServiceAccountNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get service account no content response has a 4xx status code
func (o *GetServiceAccountNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get service account no content response has a 5xx status code
func (o *GetServiceAccountNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get service account no content response a status code equal to that given
func (o *GetServiceAccountNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get service account no content response
func (o *GetServiceAccountNoContent) Code() int {
	return 204
}

func (o *GetServiceAccountNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/sa/{uuid}][%d] getServiceAccountNoContent  %+v", 204, o.Payload)
}

func (o *GetServiceAccountNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/sa/{uuid}][%d] getServiceAccountNoContent  %+v", 204, o.Payload)
}

func (o *GetServiceAccountNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetServiceAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceAccountForbidden creates a GetServiceAccountForbidden with default headers values
func NewGetServiceAccountForbidden() *GetServiceAccountForbidden {
	return &GetServiceAccountForbidden{}
}

/*
GetServiceAccountForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetServiceAccountForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get service account forbidden response has a 2xx status code
func (o *GetServiceAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get service account forbidden response has a 3xx status code
func (o *GetServiceAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get service account forbidden response has a 4xx status code
func (o *GetServiceAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get service account forbidden response has a 5xx status code
func (o *GetServiceAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get service account forbidden response a status code equal to that given
func (o *GetServiceAccountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get service account forbidden response
func (o *GetServiceAccountForbidden) Code() int {
	return 403
}

func (o *GetServiceAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/sa/{uuid}][%d] getServiceAccountForbidden  %+v", 403, o.Payload)
}

func (o *GetServiceAccountForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/sa/{uuid}][%d] getServiceAccountForbidden  %+v", 403, o.Payload)
}

func (o *GetServiceAccountForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetServiceAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceAccountNotFound creates a GetServiceAccountNotFound with default headers values
func NewGetServiceAccountNotFound() *GetServiceAccountNotFound {
	return &GetServiceAccountNotFound{}
}

/*
GetServiceAccountNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetServiceAccountNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get service account not found response has a 2xx status code
func (o *GetServiceAccountNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get service account not found response has a 3xx status code
func (o *GetServiceAccountNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get service account not found response has a 4xx status code
func (o *GetServiceAccountNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get service account not found response has a 5xx status code
func (o *GetServiceAccountNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get service account not found response a status code equal to that given
func (o *GetServiceAccountNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get service account not found response
func (o *GetServiceAccountNotFound) Code() int {
	return 404
}

func (o *GetServiceAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/sa/{uuid}][%d] getServiceAccountNotFound  %+v", 404, o.Payload)
}

func (o *GetServiceAccountNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/sa/{uuid}][%d] getServiceAccountNotFound  %+v", 404, o.Payload)
}

func (o *GetServiceAccountNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetServiceAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceAccountDefault creates a GetServiceAccountDefault with default headers values
func NewGetServiceAccountDefault(code int) *GetServiceAccountDefault {
	return &GetServiceAccountDefault{
		_statusCode: code,
	}
}

/*
GetServiceAccountDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetServiceAccountDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this get service account default response has a 2xx status code
func (o *GetServiceAccountDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get service account default response has a 3xx status code
func (o *GetServiceAccountDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get service account default response has a 4xx status code
func (o *GetServiceAccountDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get service account default response has a 5xx status code
func (o *GetServiceAccountDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get service account default response a status code equal to that given
func (o *GetServiceAccountDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get service account default response
func (o *GetServiceAccountDefault) Code() int {
	return o._statusCode
}

func (o *GetServiceAccountDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/sa/{uuid}][%d] GetServiceAccount default  %+v", o._statusCode, o.Payload)
}

func (o *GetServiceAccountDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/sa/{uuid}][%d] GetServiceAccount default  %+v", o._statusCode, o.Payload)
}

func (o *GetServiceAccountDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetServiceAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
