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

// ResendOrganizationInvitationReader is a Reader for the ResendOrganizationInvitation structure.
type ResendOrganizationInvitationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResendOrganizationInvitationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResendOrganizationInvitationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewResendOrganizationInvitationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewResendOrganizationInvitationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewResendOrganizationInvitationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewResendOrganizationInvitationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResendOrganizationInvitationOK creates a ResendOrganizationInvitationOK with default headers values
func NewResendOrganizationInvitationOK() *ResendOrganizationInvitationOK {
	return &ResendOrganizationInvitationOK{}
}

/*
ResendOrganizationInvitationOK describes a response with status code 200, with default header values.

A successful response.
*/
type ResendOrganizationInvitationOK struct {
	Payload *service_model.V1OrganizationMember
}

// IsSuccess returns true when this resend organization invitation o k response has a 2xx status code
func (o *ResendOrganizationInvitationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resend organization invitation o k response has a 3xx status code
func (o *ResendOrganizationInvitationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resend organization invitation o k response has a 4xx status code
func (o *ResendOrganizationInvitationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this resend organization invitation o k response has a 5xx status code
func (o *ResendOrganizationInvitationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this resend organization invitation o k response a status code equal to that given
func (o *ResendOrganizationInvitationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the resend organization invitation o k response
func (o *ResendOrganizationInvitationOK) Code() int {
	return 200
}

func (o *ResendOrganizationInvitationOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/invitations][%d] resendOrganizationInvitationOK  %+v", 200, o.Payload)
}

func (o *ResendOrganizationInvitationOK) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/invitations][%d] resendOrganizationInvitationOK  %+v", 200, o.Payload)
}

func (o *ResendOrganizationInvitationOK) GetPayload() *service_model.V1OrganizationMember {
	return o.Payload
}

func (o *ResendOrganizationInvitationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1OrganizationMember)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResendOrganizationInvitationNoContent creates a ResendOrganizationInvitationNoContent with default headers values
func NewResendOrganizationInvitationNoContent() *ResendOrganizationInvitationNoContent {
	return &ResendOrganizationInvitationNoContent{}
}

/*
ResendOrganizationInvitationNoContent describes a response with status code 204, with default header values.

No content.
*/
type ResendOrganizationInvitationNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this resend organization invitation no content response has a 2xx status code
func (o *ResendOrganizationInvitationNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resend organization invitation no content response has a 3xx status code
func (o *ResendOrganizationInvitationNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resend organization invitation no content response has a 4xx status code
func (o *ResendOrganizationInvitationNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this resend organization invitation no content response has a 5xx status code
func (o *ResendOrganizationInvitationNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this resend organization invitation no content response a status code equal to that given
func (o *ResendOrganizationInvitationNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the resend organization invitation no content response
func (o *ResendOrganizationInvitationNoContent) Code() int {
	return 204
}

func (o *ResendOrganizationInvitationNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/invitations][%d] resendOrganizationInvitationNoContent  %+v", 204, o.Payload)
}

func (o *ResendOrganizationInvitationNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/invitations][%d] resendOrganizationInvitationNoContent  %+v", 204, o.Payload)
}

func (o *ResendOrganizationInvitationNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ResendOrganizationInvitationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResendOrganizationInvitationForbidden creates a ResendOrganizationInvitationForbidden with default headers values
func NewResendOrganizationInvitationForbidden() *ResendOrganizationInvitationForbidden {
	return &ResendOrganizationInvitationForbidden{}
}

/*
ResendOrganizationInvitationForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ResendOrganizationInvitationForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this resend organization invitation forbidden response has a 2xx status code
func (o *ResendOrganizationInvitationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resend organization invitation forbidden response has a 3xx status code
func (o *ResendOrganizationInvitationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resend organization invitation forbidden response has a 4xx status code
func (o *ResendOrganizationInvitationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this resend organization invitation forbidden response has a 5xx status code
func (o *ResendOrganizationInvitationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this resend organization invitation forbidden response a status code equal to that given
func (o *ResendOrganizationInvitationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the resend organization invitation forbidden response
func (o *ResendOrganizationInvitationForbidden) Code() int {
	return 403
}

func (o *ResendOrganizationInvitationForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/invitations][%d] resendOrganizationInvitationForbidden  %+v", 403, o.Payload)
}

func (o *ResendOrganizationInvitationForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/invitations][%d] resendOrganizationInvitationForbidden  %+v", 403, o.Payload)
}

func (o *ResendOrganizationInvitationForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ResendOrganizationInvitationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResendOrganizationInvitationNotFound creates a ResendOrganizationInvitationNotFound with default headers values
func NewResendOrganizationInvitationNotFound() *ResendOrganizationInvitationNotFound {
	return &ResendOrganizationInvitationNotFound{}
}

/*
ResendOrganizationInvitationNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ResendOrganizationInvitationNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this resend organization invitation not found response has a 2xx status code
func (o *ResendOrganizationInvitationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resend organization invitation not found response has a 3xx status code
func (o *ResendOrganizationInvitationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resend organization invitation not found response has a 4xx status code
func (o *ResendOrganizationInvitationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this resend organization invitation not found response has a 5xx status code
func (o *ResendOrganizationInvitationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this resend organization invitation not found response a status code equal to that given
func (o *ResendOrganizationInvitationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the resend organization invitation not found response
func (o *ResendOrganizationInvitationNotFound) Code() int {
	return 404
}

func (o *ResendOrganizationInvitationNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/invitations][%d] resendOrganizationInvitationNotFound  %+v", 404, o.Payload)
}

func (o *ResendOrganizationInvitationNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/invitations][%d] resendOrganizationInvitationNotFound  %+v", 404, o.Payload)
}

func (o *ResendOrganizationInvitationNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ResendOrganizationInvitationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResendOrganizationInvitationDefault creates a ResendOrganizationInvitationDefault with default headers values
func NewResendOrganizationInvitationDefault(code int) *ResendOrganizationInvitationDefault {
	return &ResendOrganizationInvitationDefault{
		_statusCode: code,
	}
}

/*
ResendOrganizationInvitationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ResendOrganizationInvitationDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this resend organization invitation default response has a 2xx status code
func (o *ResendOrganizationInvitationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this resend organization invitation default response has a 3xx status code
func (o *ResendOrganizationInvitationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this resend organization invitation default response has a 4xx status code
func (o *ResendOrganizationInvitationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this resend organization invitation default response has a 5xx status code
func (o *ResendOrganizationInvitationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this resend organization invitation default response a status code equal to that given
func (o *ResendOrganizationInvitationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the resend organization invitation default response
func (o *ResendOrganizationInvitationDefault) Code() int {
	return o._statusCode
}

func (o *ResendOrganizationInvitationDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/invitations][%d] ResendOrganizationInvitation default  %+v", o._statusCode, o.Payload)
}

func (o *ResendOrganizationInvitationDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/invitations][%d] ResendOrganizationInvitation default  %+v", o._statusCode, o.Payload)
}

func (o *ResendOrganizationInvitationDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ResendOrganizationInvitationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
