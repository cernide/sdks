// Code generated by go-swagger; DO NOT EDIT.

package organizations_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// UpdateOrganizationInvitationReader is a Reader for the UpdateOrganizationInvitation structure.
type UpdateOrganizationInvitationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrganizationInvitationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrganizationInvitationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateOrganizationInvitationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateOrganizationInvitationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateOrganizationInvitationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateOrganizationInvitationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateOrganizationInvitationOK creates a UpdateOrganizationInvitationOK with default headers values
func NewUpdateOrganizationInvitationOK() *UpdateOrganizationInvitationOK {
	return &UpdateOrganizationInvitationOK{}
}

/*
UpdateOrganizationInvitationOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateOrganizationInvitationOK struct {
	Payload *service_model.V1OrganizationMember
}

// IsSuccess returns true when this update organization invitation o k response has a 2xx status code
func (o *UpdateOrganizationInvitationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update organization invitation o k response has a 3xx status code
func (o *UpdateOrganizationInvitationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization invitation o k response has a 4xx status code
func (o *UpdateOrganizationInvitationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update organization invitation o k response has a 5xx status code
func (o *UpdateOrganizationInvitationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization invitation o k response a status code equal to that given
func (o *UpdateOrganizationInvitationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update organization invitation o k response
func (o *UpdateOrganizationInvitationOK) Code() int {
	return 200
}

func (o *UpdateOrganizationInvitationOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/invitations][%d] updateOrganizationInvitationOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationInvitationOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/invitations][%d] updateOrganizationInvitationOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationInvitationOK) GetPayload() *service_model.V1OrganizationMember {
	return o.Payload
}

func (o *UpdateOrganizationInvitationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1OrganizationMember)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrganizationInvitationNoContent creates a UpdateOrganizationInvitationNoContent with default headers values
func NewUpdateOrganizationInvitationNoContent() *UpdateOrganizationInvitationNoContent {
	return &UpdateOrganizationInvitationNoContent{}
}

/*
UpdateOrganizationInvitationNoContent describes a response with status code 204, with default header values.

No content.
*/
type UpdateOrganizationInvitationNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this update organization invitation no content response has a 2xx status code
func (o *UpdateOrganizationInvitationNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update organization invitation no content response has a 3xx status code
func (o *UpdateOrganizationInvitationNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization invitation no content response has a 4xx status code
func (o *UpdateOrganizationInvitationNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update organization invitation no content response has a 5xx status code
func (o *UpdateOrganizationInvitationNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization invitation no content response a status code equal to that given
func (o *UpdateOrganizationInvitationNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update organization invitation no content response
func (o *UpdateOrganizationInvitationNoContent) Code() int {
	return 204
}

func (o *UpdateOrganizationInvitationNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/invitations][%d] updateOrganizationInvitationNoContent  %+v", 204, o.Payload)
}

func (o *UpdateOrganizationInvitationNoContent) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/invitations][%d] updateOrganizationInvitationNoContent  %+v", 204, o.Payload)
}

func (o *UpdateOrganizationInvitationNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateOrganizationInvitationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrganizationInvitationForbidden creates a UpdateOrganizationInvitationForbidden with default headers values
func NewUpdateOrganizationInvitationForbidden() *UpdateOrganizationInvitationForbidden {
	return &UpdateOrganizationInvitationForbidden{}
}

/*
UpdateOrganizationInvitationForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type UpdateOrganizationInvitationForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this update organization invitation forbidden response has a 2xx status code
func (o *UpdateOrganizationInvitationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update organization invitation forbidden response has a 3xx status code
func (o *UpdateOrganizationInvitationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization invitation forbidden response has a 4xx status code
func (o *UpdateOrganizationInvitationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update organization invitation forbidden response has a 5xx status code
func (o *UpdateOrganizationInvitationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization invitation forbidden response a status code equal to that given
func (o *UpdateOrganizationInvitationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update organization invitation forbidden response
func (o *UpdateOrganizationInvitationForbidden) Code() int {
	return 403
}

func (o *UpdateOrganizationInvitationForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/invitations][%d] updateOrganizationInvitationForbidden  %+v", 403, o.Payload)
}

func (o *UpdateOrganizationInvitationForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/invitations][%d] updateOrganizationInvitationForbidden  %+v", 403, o.Payload)
}

func (o *UpdateOrganizationInvitationForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateOrganizationInvitationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrganizationInvitationNotFound creates a UpdateOrganizationInvitationNotFound with default headers values
func NewUpdateOrganizationInvitationNotFound() *UpdateOrganizationInvitationNotFound {
	return &UpdateOrganizationInvitationNotFound{}
}

/*
UpdateOrganizationInvitationNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type UpdateOrganizationInvitationNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this update organization invitation not found response has a 2xx status code
func (o *UpdateOrganizationInvitationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update organization invitation not found response has a 3xx status code
func (o *UpdateOrganizationInvitationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization invitation not found response has a 4xx status code
func (o *UpdateOrganizationInvitationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update organization invitation not found response has a 5xx status code
func (o *UpdateOrganizationInvitationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization invitation not found response a status code equal to that given
func (o *UpdateOrganizationInvitationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update organization invitation not found response
func (o *UpdateOrganizationInvitationNotFound) Code() int {
	return 404
}

func (o *UpdateOrganizationInvitationNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/invitations][%d] updateOrganizationInvitationNotFound  %+v", 404, o.Payload)
}

func (o *UpdateOrganizationInvitationNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/invitations][%d] updateOrganizationInvitationNotFound  %+v", 404, o.Payload)
}

func (o *UpdateOrganizationInvitationNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateOrganizationInvitationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrganizationInvitationDefault creates a UpdateOrganizationInvitationDefault with default headers values
func NewUpdateOrganizationInvitationDefault(code int) *UpdateOrganizationInvitationDefault {
	return &UpdateOrganizationInvitationDefault{
		_statusCode: code,
	}
}

/*
UpdateOrganizationInvitationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateOrganizationInvitationDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this update organization invitation default response has a 2xx status code
func (o *UpdateOrganizationInvitationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update organization invitation default response has a 3xx status code
func (o *UpdateOrganizationInvitationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update organization invitation default response has a 4xx status code
func (o *UpdateOrganizationInvitationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update organization invitation default response has a 5xx status code
func (o *UpdateOrganizationInvitationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update organization invitation default response a status code equal to that given
func (o *UpdateOrganizationInvitationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update organization invitation default response
func (o *UpdateOrganizationInvitationDefault) Code() int {
	return o._statusCode
}

func (o *UpdateOrganizationInvitationDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/invitations][%d] UpdateOrganizationInvitation default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateOrganizationInvitationDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/invitations][%d] UpdateOrganizationInvitation default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateOrganizationInvitationDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *UpdateOrganizationInvitationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
