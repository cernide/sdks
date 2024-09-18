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

// ListOrganizationMembersReader is a Reader for the ListOrganizationMembers structure.
type ListOrganizationMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOrganizationMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOrganizationMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListOrganizationMembersNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListOrganizationMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListOrganizationMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListOrganizationMembersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOrganizationMembersOK creates a ListOrganizationMembersOK with default headers values
func NewListOrganizationMembersOK() *ListOrganizationMembersOK {
	return &ListOrganizationMembersOK{}
}

/*
ListOrganizationMembersOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListOrganizationMembersOK struct {
	Payload *service_model.V1ListOrganizationMembersResponse
}

// IsSuccess returns true when this list organization members o k response has a 2xx status code
func (o *ListOrganizationMembersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list organization members o k response has a 3xx status code
func (o *ListOrganizationMembersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list organization members o k response has a 4xx status code
func (o *ListOrganizationMembersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list organization members o k response has a 5xx status code
func (o *ListOrganizationMembersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list organization members o k response a status code equal to that given
func (o *ListOrganizationMembersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list organization members o k response
func (o *ListOrganizationMembersOK) Code() int {
	return 200
}

func (o *ListOrganizationMembersOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/members][%d] listOrganizationMembersOK  %+v", 200, o.Payload)
}

func (o *ListOrganizationMembersOK) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/members][%d] listOrganizationMembersOK  %+v", 200, o.Payload)
}

func (o *ListOrganizationMembersOK) GetPayload() *service_model.V1ListOrganizationMembersResponse {
	return o.Payload
}

func (o *ListOrganizationMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListOrganizationMembersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationMembersNoContent creates a ListOrganizationMembersNoContent with default headers values
func NewListOrganizationMembersNoContent() *ListOrganizationMembersNoContent {
	return &ListOrganizationMembersNoContent{}
}

/*
ListOrganizationMembersNoContent describes a response with status code 204, with default header values.

No content.
*/
type ListOrganizationMembersNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this list organization members no content response has a 2xx status code
func (o *ListOrganizationMembersNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list organization members no content response has a 3xx status code
func (o *ListOrganizationMembersNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list organization members no content response has a 4xx status code
func (o *ListOrganizationMembersNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this list organization members no content response has a 5xx status code
func (o *ListOrganizationMembersNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this list organization members no content response a status code equal to that given
func (o *ListOrganizationMembersNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the list organization members no content response
func (o *ListOrganizationMembersNoContent) Code() int {
	return 204
}

func (o *ListOrganizationMembersNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/members][%d] listOrganizationMembersNoContent  %+v", 204, o.Payload)
}

func (o *ListOrganizationMembersNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/members][%d] listOrganizationMembersNoContent  %+v", 204, o.Payload)
}

func (o *ListOrganizationMembersNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListOrganizationMembersNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationMembersForbidden creates a ListOrganizationMembersForbidden with default headers values
func NewListOrganizationMembersForbidden() *ListOrganizationMembersForbidden {
	return &ListOrganizationMembersForbidden{}
}

/*
ListOrganizationMembersForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ListOrganizationMembersForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this list organization members forbidden response has a 2xx status code
func (o *ListOrganizationMembersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list organization members forbidden response has a 3xx status code
func (o *ListOrganizationMembersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list organization members forbidden response has a 4xx status code
func (o *ListOrganizationMembersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list organization members forbidden response has a 5xx status code
func (o *ListOrganizationMembersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list organization members forbidden response a status code equal to that given
func (o *ListOrganizationMembersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list organization members forbidden response
func (o *ListOrganizationMembersForbidden) Code() int {
	return 403
}

func (o *ListOrganizationMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/members][%d] listOrganizationMembersForbidden  %+v", 403, o.Payload)
}

func (o *ListOrganizationMembersForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/members][%d] listOrganizationMembersForbidden  %+v", 403, o.Payload)
}

func (o *ListOrganizationMembersForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListOrganizationMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationMembersNotFound creates a ListOrganizationMembersNotFound with default headers values
func NewListOrganizationMembersNotFound() *ListOrganizationMembersNotFound {
	return &ListOrganizationMembersNotFound{}
}

/*
ListOrganizationMembersNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ListOrganizationMembersNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this list organization members not found response has a 2xx status code
func (o *ListOrganizationMembersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list organization members not found response has a 3xx status code
func (o *ListOrganizationMembersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list organization members not found response has a 4xx status code
func (o *ListOrganizationMembersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list organization members not found response has a 5xx status code
func (o *ListOrganizationMembersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list organization members not found response a status code equal to that given
func (o *ListOrganizationMembersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list organization members not found response
func (o *ListOrganizationMembersNotFound) Code() int {
	return 404
}

func (o *ListOrganizationMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/members][%d] listOrganizationMembersNotFound  %+v", 404, o.Payload)
}

func (o *ListOrganizationMembersNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/members][%d] listOrganizationMembersNotFound  %+v", 404, o.Payload)
}

func (o *ListOrganizationMembersNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListOrganizationMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationMembersDefault creates a ListOrganizationMembersDefault with default headers values
func NewListOrganizationMembersDefault(code int) *ListOrganizationMembersDefault {
	return &ListOrganizationMembersDefault{
		_statusCode: code,
	}
}

/*
ListOrganizationMembersDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListOrganizationMembersDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this list organization members default response has a 2xx status code
func (o *ListOrganizationMembersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list organization members default response has a 3xx status code
func (o *ListOrganizationMembersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list organization members default response has a 4xx status code
func (o *ListOrganizationMembersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list organization members default response has a 5xx status code
func (o *ListOrganizationMembersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list organization members default response a status code equal to that given
func (o *ListOrganizationMembersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list organization members default response
func (o *ListOrganizationMembersDefault) Code() int {
	return o._statusCode
}

func (o *ListOrganizationMembersDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/members][%d] ListOrganizationMembers default  %+v", o._statusCode, o.Payload)
}

func (o *ListOrganizationMembersDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/members][%d] ListOrganizationMembers default  %+v", o._statusCode, o.Payload)
}

func (o *ListOrganizationMembersDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListOrganizationMembersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
