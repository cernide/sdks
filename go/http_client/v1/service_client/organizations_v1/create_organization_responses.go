// Code generated by go-swagger; DO NOT EDIT.

package organizations_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/go/http_client/v1/service_model"
)

// CreateOrganizationReader is a Reader for the CreateOrganization structure.
type CreateOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCreateOrganizationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateOrganizationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateOrganizationOK creates a CreateOrganizationOK with default headers values
func NewCreateOrganizationOK() *CreateOrganizationOK {
	return &CreateOrganizationOK{}
}

/*
CreateOrganizationOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateOrganizationOK struct {
	Payload *service_model.V1Organization
}

// IsSuccess returns true when this create organization o k response has a 2xx status code
func (o *CreateOrganizationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create organization o k response has a 3xx status code
func (o *CreateOrganizationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create organization o k response has a 4xx status code
func (o *CreateOrganizationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create organization o k response has a 5xx status code
func (o *CreateOrganizationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create organization o k response a status code equal to that given
func (o *CreateOrganizationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create organization o k response
func (o *CreateOrganizationOK) Code() int {
	return 200
}

func (o *CreateOrganizationOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/create][%d] createOrganizationOK  %+v", 200, o.Payload)
}

func (o *CreateOrganizationOK) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/create][%d] createOrganizationOK  %+v", 200, o.Payload)
}

func (o *CreateOrganizationOK) GetPayload() *service_model.V1Organization {
	return o.Payload
}

func (o *CreateOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Organization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationNoContent creates a CreateOrganizationNoContent with default headers values
func NewCreateOrganizationNoContent() *CreateOrganizationNoContent {
	return &CreateOrganizationNoContent{}
}

/*
CreateOrganizationNoContent describes a response with status code 204, with default header values.

No content.
*/
type CreateOrganizationNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this create organization no content response has a 2xx status code
func (o *CreateOrganizationNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create organization no content response has a 3xx status code
func (o *CreateOrganizationNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create organization no content response has a 4xx status code
func (o *CreateOrganizationNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this create organization no content response has a 5xx status code
func (o *CreateOrganizationNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this create organization no content response a status code equal to that given
func (o *CreateOrganizationNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the create organization no content response
func (o *CreateOrganizationNoContent) Code() int {
	return 204
}

func (o *CreateOrganizationNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/create][%d] createOrganizationNoContent  %+v", 204, o.Payload)
}

func (o *CreateOrganizationNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/create][%d] createOrganizationNoContent  %+v", 204, o.Payload)
}

func (o *CreateOrganizationNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateOrganizationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationForbidden creates a CreateOrganizationForbidden with default headers values
func NewCreateOrganizationForbidden() *CreateOrganizationForbidden {
	return &CreateOrganizationForbidden{}
}

/*
CreateOrganizationForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type CreateOrganizationForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this create organization forbidden response has a 2xx status code
func (o *CreateOrganizationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create organization forbidden response has a 3xx status code
func (o *CreateOrganizationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create organization forbidden response has a 4xx status code
func (o *CreateOrganizationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create organization forbidden response has a 5xx status code
func (o *CreateOrganizationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create organization forbidden response a status code equal to that given
func (o *CreateOrganizationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create organization forbidden response
func (o *CreateOrganizationForbidden) Code() int {
	return 403
}

func (o *CreateOrganizationForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/create][%d] createOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *CreateOrganizationForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/create][%d] createOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *CreateOrganizationForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationNotFound creates a CreateOrganizationNotFound with default headers values
func NewCreateOrganizationNotFound() *CreateOrganizationNotFound {
	return &CreateOrganizationNotFound{}
}

/*
CreateOrganizationNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type CreateOrganizationNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this create organization not found response has a 2xx status code
func (o *CreateOrganizationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create organization not found response has a 3xx status code
func (o *CreateOrganizationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create organization not found response has a 4xx status code
func (o *CreateOrganizationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create organization not found response has a 5xx status code
func (o *CreateOrganizationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create organization not found response a status code equal to that given
func (o *CreateOrganizationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create organization not found response
func (o *CreateOrganizationNotFound) Code() int {
	return 404
}

func (o *CreateOrganizationNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/create][%d] createOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *CreateOrganizationNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/create][%d] createOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *CreateOrganizationNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationDefault creates a CreateOrganizationDefault with default headers values
func NewCreateOrganizationDefault(code int) *CreateOrganizationDefault {
	return &CreateOrganizationDefault{
		_statusCode: code,
	}
}

/*
CreateOrganizationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateOrganizationDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this create organization default response has a 2xx status code
func (o *CreateOrganizationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create organization default response has a 3xx status code
func (o *CreateOrganizationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create organization default response has a 4xx status code
func (o *CreateOrganizationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create organization default response has a 5xx status code
func (o *CreateOrganizationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create organization default response a status code equal to that given
func (o *CreateOrganizationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create organization default response
func (o *CreateOrganizationDefault) Code() int {
	return o._statusCode
}

func (o *CreateOrganizationDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/create][%d] CreateOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *CreateOrganizationDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/create][%d] CreateOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *CreateOrganizationDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *CreateOrganizationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
