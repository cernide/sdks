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

// GetTeamMemberReader is a Reader for the GetTeamMember structure.
type GetTeamMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetTeamMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetTeamMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTeamMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetTeamMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTeamMemberOK creates a GetTeamMemberOK with default headers values
func NewGetTeamMemberOK() *GetTeamMemberOK {
	return &GetTeamMemberOK{}
}

/*
GetTeamMemberOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetTeamMemberOK struct {
	Payload *service_model.V1TeamMember
}

// IsSuccess returns true when this get team member o k response has a 2xx status code
func (o *GetTeamMemberOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get team member o k response has a 3xx status code
func (o *GetTeamMemberOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team member o k response has a 4xx status code
func (o *GetTeamMemberOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team member o k response has a 5xx status code
func (o *GetTeamMemberOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get team member o k response a status code equal to that given
func (o *GetTeamMemberOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get team member o k response
func (o *GetTeamMemberOK) Code() int {
	return 200
}

func (o *GetTeamMemberOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{team}/members/{user}][%d] getTeamMemberOK  %+v", 200, o.Payload)
}

func (o *GetTeamMemberOK) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{team}/members/{user}][%d] getTeamMemberOK  %+v", 200, o.Payload)
}

func (o *GetTeamMemberOK) GetPayload() *service_model.V1TeamMember {
	return o.Payload
}

func (o *GetTeamMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1TeamMember)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMemberNoContent creates a GetTeamMemberNoContent with default headers values
func NewGetTeamMemberNoContent() *GetTeamMemberNoContent {
	return &GetTeamMemberNoContent{}
}

/*
GetTeamMemberNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetTeamMemberNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get team member no content response has a 2xx status code
func (o *GetTeamMemberNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get team member no content response has a 3xx status code
func (o *GetTeamMemberNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team member no content response has a 4xx status code
func (o *GetTeamMemberNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team member no content response has a 5xx status code
func (o *GetTeamMemberNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get team member no content response a status code equal to that given
func (o *GetTeamMemberNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get team member no content response
func (o *GetTeamMemberNoContent) Code() int {
	return 204
}

func (o *GetTeamMemberNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{team}/members/{user}][%d] getTeamMemberNoContent  %+v", 204, o.Payload)
}

func (o *GetTeamMemberNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{team}/members/{user}][%d] getTeamMemberNoContent  %+v", 204, o.Payload)
}

func (o *GetTeamMemberNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetTeamMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMemberForbidden creates a GetTeamMemberForbidden with default headers values
func NewGetTeamMemberForbidden() *GetTeamMemberForbidden {
	return &GetTeamMemberForbidden{}
}

/*
GetTeamMemberForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetTeamMemberForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get team member forbidden response has a 2xx status code
func (o *GetTeamMemberForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team member forbidden response has a 3xx status code
func (o *GetTeamMemberForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team member forbidden response has a 4xx status code
func (o *GetTeamMemberForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team member forbidden response has a 5xx status code
func (o *GetTeamMemberForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get team member forbidden response a status code equal to that given
func (o *GetTeamMemberForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get team member forbidden response
func (o *GetTeamMemberForbidden) Code() int {
	return 403
}

func (o *GetTeamMemberForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{team}/members/{user}][%d] getTeamMemberForbidden  %+v", 403, o.Payload)
}

func (o *GetTeamMemberForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{team}/members/{user}][%d] getTeamMemberForbidden  %+v", 403, o.Payload)
}

func (o *GetTeamMemberForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetTeamMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMemberNotFound creates a GetTeamMemberNotFound with default headers values
func NewGetTeamMemberNotFound() *GetTeamMemberNotFound {
	return &GetTeamMemberNotFound{}
}

/*
GetTeamMemberNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetTeamMemberNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get team member not found response has a 2xx status code
func (o *GetTeamMemberNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team member not found response has a 3xx status code
func (o *GetTeamMemberNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team member not found response has a 4xx status code
func (o *GetTeamMemberNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team member not found response has a 5xx status code
func (o *GetTeamMemberNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get team member not found response a status code equal to that given
func (o *GetTeamMemberNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get team member not found response
func (o *GetTeamMemberNotFound) Code() int {
	return 404
}

func (o *GetTeamMemberNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{team}/members/{user}][%d] getTeamMemberNotFound  %+v", 404, o.Payload)
}

func (o *GetTeamMemberNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{team}/members/{user}][%d] getTeamMemberNotFound  %+v", 404, o.Payload)
}

func (o *GetTeamMemberNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetTeamMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMemberDefault creates a GetTeamMemberDefault with default headers values
func NewGetTeamMemberDefault(code int) *GetTeamMemberDefault {
	return &GetTeamMemberDefault{
		_statusCode: code,
	}
}

/*
GetTeamMemberDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetTeamMemberDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this get team member default response has a 2xx status code
func (o *GetTeamMemberDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get team member default response has a 3xx status code
func (o *GetTeamMemberDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get team member default response has a 4xx status code
func (o *GetTeamMemberDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get team member default response has a 5xx status code
func (o *GetTeamMemberDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get team member default response a status code equal to that given
func (o *GetTeamMemberDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get team member default response
func (o *GetTeamMemberDefault) Code() int {
	return o._statusCode
}

func (o *GetTeamMemberDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{team}/members/{user}][%d] GetTeamMember default  %+v", o._statusCode, o.Payload)
}

func (o *GetTeamMemberDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{team}/members/{user}][%d] GetTeamMember default  %+v", o._statusCode, o.Payload)
}

func (o *GetTeamMemberDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetTeamMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
