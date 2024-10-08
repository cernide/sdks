// Code generated by go-swagger; DO NOT EDIT.

package teams_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// UpdateTeamReader is a Reader for the UpdateTeam structure.
type UpdateTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateTeamNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateTeamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateTeamOK creates a UpdateTeamOK with default headers values
func NewUpdateTeamOK() *UpdateTeamOK {
	return &UpdateTeamOK{}
}

/*
UpdateTeamOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateTeamOK struct {
	Payload *service_model.V1Team
}

// IsSuccess returns true when this update team o k response has a 2xx status code
func (o *UpdateTeamOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update team o k response has a 3xx status code
func (o *UpdateTeamOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update team o k response has a 4xx status code
func (o *UpdateTeamOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update team o k response has a 5xx status code
func (o *UpdateTeamOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update team o k response a status code equal to that given
func (o *UpdateTeamOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update team o k response
func (o *UpdateTeamOK) Code() int {
	return 200
}

func (o *UpdateTeamOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team.name}][%d] updateTeamOK  %+v", 200, o.Payload)
}

func (o *UpdateTeamOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team.name}][%d] updateTeamOK  %+v", 200, o.Payload)
}

func (o *UpdateTeamOK) GetPayload() *service_model.V1Team {
	return o.Payload
}

func (o *UpdateTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Team)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamNoContent creates a UpdateTeamNoContent with default headers values
func NewUpdateTeamNoContent() *UpdateTeamNoContent {
	return &UpdateTeamNoContent{}
}

/*
UpdateTeamNoContent describes a response with status code 204, with default header values.

No content.
*/
type UpdateTeamNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this update team no content response has a 2xx status code
func (o *UpdateTeamNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update team no content response has a 3xx status code
func (o *UpdateTeamNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update team no content response has a 4xx status code
func (o *UpdateTeamNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update team no content response has a 5xx status code
func (o *UpdateTeamNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update team no content response a status code equal to that given
func (o *UpdateTeamNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update team no content response
func (o *UpdateTeamNoContent) Code() int {
	return 204
}

func (o *UpdateTeamNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team.name}][%d] updateTeamNoContent  %+v", 204, o.Payload)
}

func (o *UpdateTeamNoContent) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team.name}][%d] updateTeamNoContent  %+v", 204, o.Payload)
}

func (o *UpdateTeamNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateTeamNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamForbidden creates a UpdateTeamForbidden with default headers values
func NewUpdateTeamForbidden() *UpdateTeamForbidden {
	return &UpdateTeamForbidden{}
}

/*
UpdateTeamForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type UpdateTeamForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this update team forbidden response has a 2xx status code
func (o *UpdateTeamForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update team forbidden response has a 3xx status code
func (o *UpdateTeamForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update team forbidden response has a 4xx status code
func (o *UpdateTeamForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update team forbidden response has a 5xx status code
func (o *UpdateTeamForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update team forbidden response a status code equal to that given
func (o *UpdateTeamForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update team forbidden response
func (o *UpdateTeamForbidden) Code() int {
	return 403
}

func (o *UpdateTeamForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team.name}][%d] updateTeamForbidden  %+v", 403, o.Payload)
}

func (o *UpdateTeamForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team.name}][%d] updateTeamForbidden  %+v", 403, o.Payload)
}

func (o *UpdateTeamForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamNotFound creates a UpdateTeamNotFound with default headers values
func NewUpdateTeamNotFound() *UpdateTeamNotFound {
	return &UpdateTeamNotFound{}
}

/*
UpdateTeamNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type UpdateTeamNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this update team not found response has a 2xx status code
func (o *UpdateTeamNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update team not found response has a 3xx status code
func (o *UpdateTeamNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update team not found response has a 4xx status code
func (o *UpdateTeamNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update team not found response has a 5xx status code
func (o *UpdateTeamNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update team not found response a status code equal to that given
func (o *UpdateTeamNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update team not found response
func (o *UpdateTeamNotFound) Code() int {
	return 404
}

func (o *UpdateTeamNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team.name}][%d] updateTeamNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTeamNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team.name}][%d] updateTeamNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTeamNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamDefault creates a UpdateTeamDefault with default headers values
func NewUpdateTeamDefault(code int) *UpdateTeamDefault {
	return &UpdateTeamDefault{
		_statusCode: code,
	}
}

/*
UpdateTeamDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateTeamDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this update team default response has a 2xx status code
func (o *UpdateTeamDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update team default response has a 3xx status code
func (o *UpdateTeamDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update team default response has a 4xx status code
func (o *UpdateTeamDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update team default response has a 5xx status code
func (o *UpdateTeamDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update team default response a status code equal to that given
func (o *UpdateTeamDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update team default response
func (o *UpdateTeamDefault) Code() int {
	return o._statusCode
}

func (o *UpdateTeamDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team.name}][%d] UpdateTeam default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTeamDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team.name}][%d] UpdateTeam default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTeamDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *UpdateTeamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
