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

// CreateServiceAccountReader is a Reader for the CreateServiceAccount structure.
type CreateServiceAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateServiceAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateServiceAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCreateServiceAccountNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateServiceAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateServiceAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateServiceAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateServiceAccountOK creates a CreateServiceAccountOK with default headers values
func NewCreateServiceAccountOK() *CreateServiceAccountOK {
	return &CreateServiceAccountOK{}
}

/*
CreateServiceAccountOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateServiceAccountOK struct {
	Payload *service_model.V1ServiceAccount
}

// IsSuccess returns true when this create service account o k response has a 2xx status code
func (o *CreateServiceAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create service account o k response has a 3xx status code
func (o *CreateServiceAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create service account o k response has a 4xx status code
func (o *CreateServiceAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create service account o k response has a 5xx status code
func (o *CreateServiceAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create service account o k response a status code equal to that given
func (o *CreateServiceAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create service account o k response
func (o *CreateServiceAccountOK) Code() int {
	return 200
}

func (o *CreateServiceAccountOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/sa][%d] createServiceAccountOK  %+v", 200, o.Payload)
}

func (o *CreateServiceAccountOK) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/sa][%d] createServiceAccountOK  %+v", 200, o.Payload)
}

func (o *CreateServiceAccountOK) GetPayload() *service_model.V1ServiceAccount {
	return o.Payload
}

func (o *CreateServiceAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ServiceAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceAccountNoContent creates a CreateServiceAccountNoContent with default headers values
func NewCreateServiceAccountNoContent() *CreateServiceAccountNoContent {
	return &CreateServiceAccountNoContent{}
}

/*
CreateServiceAccountNoContent describes a response with status code 204, with default header values.

No content.
*/
type CreateServiceAccountNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this create service account no content response has a 2xx status code
func (o *CreateServiceAccountNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create service account no content response has a 3xx status code
func (o *CreateServiceAccountNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create service account no content response has a 4xx status code
func (o *CreateServiceAccountNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this create service account no content response has a 5xx status code
func (o *CreateServiceAccountNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this create service account no content response a status code equal to that given
func (o *CreateServiceAccountNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the create service account no content response
func (o *CreateServiceAccountNoContent) Code() int {
	return 204
}

func (o *CreateServiceAccountNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/sa][%d] createServiceAccountNoContent  %+v", 204, o.Payload)
}

func (o *CreateServiceAccountNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/sa][%d] createServiceAccountNoContent  %+v", 204, o.Payload)
}

func (o *CreateServiceAccountNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateServiceAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceAccountForbidden creates a CreateServiceAccountForbidden with default headers values
func NewCreateServiceAccountForbidden() *CreateServiceAccountForbidden {
	return &CreateServiceAccountForbidden{}
}

/*
CreateServiceAccountForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type CreateServiceAccountForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this create service account forbidden response has a 2xx status code
func (o *CreateServiceAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create service account forbidden response has a 3xx status code
func (o *CreateServiceAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create service account forbidden response has a 4xx status code
func (o *CreateServiceAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create service account forbidden response has a 5xx status code
func (o *CreateServiceAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create service account forbidden response a status code equal to that given
func (o *CreateServiceAccountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create service account forbidden response
func (o *CreateServiceAccountForbidden) Code() int {
	return 403
}

func (o *CreateServiceAccountForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/sa][%d] createServiceAccountForbidden  %+v", 403, o.Payload)
}

func (o *CreateServiceAccountForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/sa][%d] createServiceAccountForbidden  %+v", 403, o.Payload)
}

func (o *CreateServiceAccountForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateServiceAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceAccountNotFound creates a CreateServiceAccountNotFound with default headers values
func NewCreateServiceAccountNotFound() *CreateServiceAccountNotFound {
	return &CreateServiceAccountNotFound{}
}

/*
CreateServiceAccountNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type CreateServiceAccountNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this create service account not found response has a 2xx status code
func (o *CreateServiceAccountNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create service account not found response has a 3xx status code
func (o *CreateServiceAccountNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create service account not found response has a 4xx status code
func (o *CreateServiceAccountNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create service account not found response has a 5xx status code
func (o *CreateServiceAccountNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create service account not found response a status code equal to that given
func (o *CreateServiceAccountNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create service account not found response
func (o *CreateServiceAccountNotFound) Code() int {
	return 404
}

func (o *CreateServiceAccountNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/sa][%d] createServiceAccountNotFound  %+v", 404, o.Payload)
}

func (o *CreateServiceAccountNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/sa][%d] createServiceAccountNotFound  %+v", 404, o.Payload)
}

func (o *CreateServiceAccountNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateServiceAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceAccountDefault creates a CreateServiceAccountDefault with default headers values
func NewCreateServiceAccountDefault(code int) *CreateServiceAccountDefault {
	return &CreateServiceAccountDefault{
		_statusCode: code,
	}
}

/*
CreateServiceAccountDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateServiceAccountDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this create service account default response has a 2xx status code
func (o *CreateServiceAccountDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create service account default response has a 3xx status code
func (o *CreateServiceAccountDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create service account default response has a 4xx status code
func (o *CreateServiceAccountDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create service account default response has a 5xx status code
func (o *CreateServiceAccountDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create service account default response a status code equal to that given
func (o *CreateServiceAccountDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create service account default response
func (o *CreateServiceAccountDefault) Code() int {
	return o._statusCode
}

func (o *CreateServiceAccountDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/sa][%d] CreateServiceAccount default  %+v", o._statusCode, o.Payload)
}

func (o *CreateServiceAccountDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/sa][%d] CreateServiceAccount default  %+v", o._statusCode, o.Payload)
}

func (o *CreateServiceAccountDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *CreateServiceAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
