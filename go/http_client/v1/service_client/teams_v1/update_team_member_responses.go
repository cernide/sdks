// Code generated by go-swagger; DO NOT EDIT.

package teams_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/go/http_client/v1/service_model"
)

// UpdateTeamMemberReader is a Reader for the UpdateTeamMember structure.
type UpdateTeamMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTeamMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTeamMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateTeamMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateTeamMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTeamMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateTeamMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateTeamMemberOK creates a UpdateTeamMemberOK with default headers values
func NewUpdateTeamMemberOK() *UpdateTeamMemberOK {
	return &UpdateTeamMemberOK{}
}

/*
UpdateTeamMemberOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateTeamMemberOK struct {
	Payload *service_model.V1TeamMember
}

// IsSuccess returns true when this update team member o k response has a 2xx status code
func (o *UpdateTeamMemberOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update team member o k response has a 3xx status code
func (o *UpdateTeamMemberOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update team member o k response has a 4xx status code
func (o *UpdateTeamMemberOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update team member o k response has a 5xx status code
func (o *UpdateTeamMemberOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update team member o k response a status code equal to that given
func (o *UpdateTeamMemberOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update team member o k response
func (o *UpdateTeamMemberOK) Code() int {
	return 200
}

func (o *UpdateTeamMemberOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team}/members/{member.user}][%d] updateTeamMemberOK  %+v", 200, o.Payload)
}

func (o *UpdateTeamMemberOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team}/members/{member.user}][%d] updateTeamMemberOK  %+v", 200, o.Payload)
}

func (o *UpdateTeamMemberOK) GetPayload() *service_model.V1TeamMember {
	return o.Payload
}

func (o *UpdateTeamMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1TeamMember)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamMemberNoContent creates a UpdateTeamMemberNoContent with default headers values
func NewUpdateTeamMemberNoContent() *UpdateTeamMemberNoContent {
	return &UpdateTeamMemberNoContent{}
}

/*
UpdateTeamMemberNoContent describes a response with status code 204, with default header values.

No content.
*/
type UpdateTeamMemberNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this update team member no content response has a 2xx status code
func (o *UpdateTeamMemberNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update team member no content response has a 3xx status code
func (o *UpdateTeamMemberNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update team member no content response has a 4xx status code
func (o *UpdateTeamMemberNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update team member no content response has a 5xx status code
func (o *UpdateTeamMemberNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update team member no content response a status code equal to that given
func (o *UpdateTeamMemberNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update team member no content response
func (o *UpdateTeamMemberNoContent) Code() int {
	return 204
}

func (o *UpdateTeamMemberNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team}/members/{member.user}][%d] updateTeamMemberNoContent  %+v", 204, o.Payload)
}

func (o *UpdateTeamMemberNoContent) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team}/members/{member.user}][%d] updateTeamMemberNoContent  %+v", 204, o.Payload)
}

func (o *UpdateTeamMemberNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateTeamMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamMemberForbidden creates a UpdateTeamMemberForbidden with default headers values
func NewUpdateTeamMemberForbidden() *UpdateTeamMemberForbidden {
	return &UpdateTeamMemberForbidden{}
}

/*
UpdateTeamMemberForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type UpdateTeamMemberForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this update team member forbidden response has a 2xx status code
func (o *UpdateTeamMemberForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update team member forbidden response has a 3xx status code
func (o *UpdateTeamMemberForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update team member forbidden response has a 4xx status code
func (o *UpdateTeamMemberForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update team member forbidden response has a 5xx status code
func (o *UpdateTeamMemberForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update team member forbidden response a status code equal to that given
func (o *UpdateTeamMemberForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update team member forbidden response
func (o *UpdateTeamMemberForbidden) Code() int {
	return 403
}

func (o *UpdateTeamMemberForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team}/members/{member.user}][%d] updateTeamMemberForbidden  %+v", 403, o.Payload)
}

func (o *UpdateTeamMemberForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team}/members/{member.user}][%d] updateTeamMemberForbidden  %+v", 403, o.Payload)
}

func (o *UpdateTeamMemberForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateTeamMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamMemberNotFound creates a UpdateTeamMemberNotFound with default headers values
func NewUpdateTeamMemberNotFound() *UpdateTeamMemberNotFound {
	return &UpdateTeamMemberNotFound{}
}

/*
UpdateTeamMemberNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type UpdateTeamMemberNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this update team member not found response has a 2xx status code
func (o *UpdateTeamMemberNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update team member not found response has a 3xx status code
func (o *UpdateTeamMemberNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update team member not found response has a 4xx status code
func (o *UpdateTeamMemberNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update team member not found response has a 5xx status code
func (o *UpdateTeamMemberNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update team member not found response a status code equal to that given
func (o *UpdateTeamMemberNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update team member not found response
func (o *UpdateTeamMemberNotFound) Code() int {
	return 404
}

func (o *UpdateTeamMemberNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team}/members/{member.user}][%d] updateTeamMemberNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTeamMemberNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team}/members/{member.user}][%d] updateTeamMemberNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTeamMemberNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateTeamMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamMemberDefault creates a UpdateTeamMemberDefault with default headers values
func NewUpdateTeamMemberDefault(code int) *UpdateTeamMemberDefault {
	return &UpdateTeamMemberDefault{
		_statusCode: code,
	}
}

/*
UpdateTeamMemberDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateTeamMemberDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this update team member default response has a 2xx status code
func (o *UpdateTeamMemberDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update team member default response has a 3xx status code
func (o *UpdateTeamMemberDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update team member default response has a 4xx status code
func (o *UpdateTeamMemberDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update team member default response has a 5xx status code
func (o *UpdateTeamMemberDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update team member default response a status code equal to that given
func (o *UpdateTeamMemberDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update team member default response
func (o *UpdateTeamMemberDefault) Code() int {
	return o._statusCode
}

func (o *UpdateTeamMemberDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team}/members/{member.user}][%d] UpdateTeamMember default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTeamMemberDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team}/members/{member.user}][%d] UpdateTeamMember default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTeamMemberDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *UpdateTeamMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
