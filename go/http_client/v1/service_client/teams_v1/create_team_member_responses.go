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

// CreateTeamMemberReader is a Reader for the CreateTeamMember structure.
type CreateTeamMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTeamMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTeamMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCreateTeamMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateTeamMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateTeamMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateTeamMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateTeamMemberOK creates a CreateTeamMemberOK with default headers values
func NewCreateTeamMemberOK() *CreateTeamMemberOK {
	return &CreateTeamMemberOK{}
}

/*
CreateTeamMemberOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateTeamMemberOK struct {
	Payload *service_model.V1TeamMember
}

// IsSuccess returns true when this create team member o k response has a 2xx status code
func (o *CreateTeamMemberOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create team member o k response has a 3xx status code
func (o *CreateTeamMemberOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create team member o k response has a 4xx status code
func (o *CreateTeamMemberOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create team member o k response has a 5xx status code
func (o *CreateTeamMemberOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create team member o k response a status code equal to that given
func (o *CreateTeamMemberOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create team member o k response
func (o *CreateTeamMemberOK) Code() int {
	return 200
}

func (o *CreateTeamMemberOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{team}/members][%d] createTeamMemberOK  %+v", 200, o.Payload)
}

func (o *CreateTeamMemberOK) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{team}/members][%d] createTeamMemberOK  %+v", 200, o.Payload)
}

func (o *CreateTeamMemberOK) GetPayload() *service_model.V1TeamMember {
	return o.Payload
}

func (o *CreateTeamMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1TeamMember)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamMemberNoContent creates a CreateTeamMemberNoContent with default headers values
func NewCreateTeamMemberNoContent() *CreateTeamMemberNoContent {
	return &CreateTeamMemberNoContent{}
}

/*
CreateTeamMemberNoContent describes a response with status code 204, with default header values.

No content.
*/
type CreateTeamMemberNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this create team member no content response has a 2xx status code
func (o *CreateTeamMemberNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create team member no content response has a 3xx status code
func (o *CreateTeamMemberNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create team member no content response has a 4xx status code
func (o *CreateTeamMemberNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this create team member no content response has a 5xx status code
func (o *CreateTeamMemberNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this create team member no content response a status code equal to that given
func (o *CreateTeamMemberNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the create team member no content response
func (o *CreateTeamMemberNoContent) Code() int {
	return 204
}

func (o *CreateTeamMemberNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{team}/members][%d] createTeamMemberNoContent  %+v", 204, o.Payload)
}

func (o *CreateTeamMemberNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{team}/members][%d] createTeamMemberNoContent  %+v", 204, o.Payload)
}

func (o *CreateTeamMemberNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateTeamMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamMemberForbidden creates a CreateTeamMemberForbidden with default headers values
func NewCreateTeamMemberForbidden() *CreateTeamMemberForbidden {
	return &CreateTeamMemberForbidden{}
}

/*
CreateTeamMemberForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type CreateTeamMemberForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this create team member forbidden response has a 2xx status code
func (o *CreateTeamMemberForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create team member forbidden response has a 3xx status code
func (o *CreateTeamMemberForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create team member forbidden response has a 4xx status code
func (o *CreateTeamMemberForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create team member forbidden response has a 5xx status code
func (o *CreateTeamMemberForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create team member forbidden response a status code equal to that given
func (o *CreateTeamMemberForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create team member forbidden response
func (o *CreateTeamMemberForbidden) Code() int {
	return 403
}

func (o *CreateTeamMemberForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{team}/members][%d] createTeamMemberForbidden  %+v", 403, o.Payload)
}

func (o *CreateTeamMemberForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{team}/members][%d] createTeamMemberForbidden  %+v", 403, o.Payload)
}

func (o *CreateTeamMemberForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateTeamMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamMemberNotFound creates a CreateTeamMemberNotFound with default headers values
func NewCreateTeamMemberNotFound() *CreateTeamMemberNotFound {
	return &CreateTeamMemberNotFound{}
}

/*
CreateTeamMemberNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type CreateTeamMemberNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this create team member not found response has a 2xx status code
func (o *CreateTeamMemberNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create team member not found response has a 3xx status code
func (o *CreateTeamMemberNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create team member not found response has a 4xx status code
func (o *CreateTeamMemberNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create team member not found response has a 5xx status code
func (o *CreateTeamMemberNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create team member not found response a status code equal to that given
func (o *CreateTeamMemberNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create team member not found response
func (o *CreateTeamMemberNotFound) Code() int {
	return 404
}

func (o *CreateTeamMemberNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{team}/members][%d] createTeamMemberNotFound  %+v", 404, o.Payload)
}

func (o *CreateTeamMemberNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{team}/members][%d] createTeamMemberNotFound  %+v", 404, o.Payload)
}

func (o *CreateTeamMemberNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateTeamMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamMemberDefault creates a CreateTeamMemberDefault with default headers values
func NewCreateTeamMemberDefault(code int) *CreateTeamMemberDefault {
	return &CreateTeamMemberDefault{
		_statusCode: code,
	}
}

/*
CreateTeamMemberDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateTeamMemberDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this create team member default response has a 2xx status code
func (o *CreateTeamMemberDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create team member default response has a 3xx status code
func (o *CreateTeamMemberDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create team member default response has a 4xx status code
func (o *CreateTeamMemberDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create team member default response has a 5xx status code
func (o *CreateTeamMemberDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create team member default response a status code equal to that given
func (o *CreateTeamMemberDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create team member default response
func (o *CreateTeamMemberDefault) Code() int {
	return o._statusCode
}

func (o *CreateTeamMemberDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{team}/members][%d] CreateTeamMember default  %+v", o._statusCode, o.Payload)
}

func (o *CreateTeamMemberDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{team}/members][%d] CreateTeamMember default  %+v", o._statusCode, o.Payload)
}

func (o *CreateTeamMemberDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *CreateTeamMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
