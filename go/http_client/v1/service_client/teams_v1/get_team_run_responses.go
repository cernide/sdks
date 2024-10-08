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

// GetTeamRunReader is a Reader for the GetTeamRun structure.
type GetTeamRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetTeamRunNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetTeamRunForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTeamRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetTeamRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTeamRunOK creates a GetTeamRunOK with default headers values
func NewGetTeamRunOK() *GetTeamRunOK {
	return &GetTeamRunOK{}
}

/*
GetTeamRunOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetTeamRunOK struct {
	Payload *service_model.V1Run
}

// IsSuccess returns true when this get team run o k response has a 2xx status code
func (o *GetTeamRunOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get team run o k response has a 3xx status code
func (o *GetTeamRunOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team run o k response has a 4xx status code
func (o *GetTeamRunOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team run o k response has a 5xx status code
func (o *GetTeamRunOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get team run o k response a status code equal to that given
func (o *GetTeamRunOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get team run o k response
func (o *GetTeamRunOK) Code() int {
	return 200
}

func (o *GetTeamRunOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{entity}/runs/{uuid}][%d] getTeamRunOK  %+v", 200, o.Payload)
}

func (o *GetTeamRunOK) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{entity}/runs/{uuid}][%d] getTeamRunOK  %+v", 200, o.Payload)
}

func (o *GetTeamRunOK) GetPayload() *service_model.V1Run {
	return o.Payload
}

func (o *GetTeamRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Run)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamRunNoContent creates a GetTeamRunNoContent with default headers values
func NewGetTeamRunNoContent() *GetTeamRunNoContent {
	return &GetTeamRunNoContent{}
}

/*
GetTeamRunNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetTeamRunNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get team run no content response has a 2xx status code
func (o *GetTeamRunNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get team run no content response has a 3xx status code
func (o *GetTeamRunNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team run no content response has a 4xx status code
func (o *GetTeamRunNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team run no content response has a 5xx status code
func (o *GetTeamRunNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get team run no content response a status code equal to that given
func (o *GetTeamRunNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get team run no content response
func (o *GetTeamRunNoContent) Code() int {
	return 204
}

func (o *GetTeamRunNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{entity}/runs/{uuid}][%d] getTeamRunNoContent  %+v", 204, o.Payload)
}

func (o *GetTeamRunNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{entity}/runs/{uuid}][%d] getTeamRunNoContent  %+v", 204, o.Payload)
}

func (o *GetTeamRunNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetTeamRunNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamRunForbidden creates a GetTeamRunForbidden with default headers values
func NewGetTeamRunForbidden() *GetTeamRunForbidden {
	return &GetTeamRunForbidden{}
}

/*
GetTeamRunForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetTeamRunForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get team run forbidden response has a 2xx status code
func (o *GetTeamRunForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team run forbidden response has a 3xx status code
func (o *GetTeamRunForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team run forbidden response has a 4xx status code
func (o *GetTeamRunForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team run forbidden response has a 5xx status code
func (o *GetTeamRunForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get team run forbidden response a status code equal to that given
func (o *GetTeamRunForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get team run forbidden response
func (o *GetTeamRunForbidden) Code() int {
	return 403
}

func (o *GetTeamRunForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{entity}/runs/{uuid}][%d] getTeamRunForbidden  %+v", 403, o.Payload)
}

func (o *GetTeamRunForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{entity}/runs/{uuid}][%d] getTeamRunForbidden  %+v", 403, o.Payload)
}

func (o *GetTeamRunForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetTeamRunForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamRunNotFound creates a GetTeamRunNotFound with default headers values
func NewGetTeamRunNotFound() *GetTeamRunNotFound {
	return &GetTeamRunNotFound{}
}

/*
GetTeamRunNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetTeamRunNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get team run not found response has a 2xx status code
func (o *GetTeamRunNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team run not found response has a 3xx status code
func (o *GetTeamRunNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team run not found response has a 4xx status code
func (o *GetTeamRunNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team run not found response has a 5xx status code
func (o *GetTeamRunNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get team run not found response a status code equal to that given
func (o *GetTeamRunNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get team run not found response
func (o *GetTeamRunNotFound) Code() int {
	return 404
}

func (o *GetTeamRunNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{entity}/runs/{uuid}][%d] getTeamRunNotFound  %+v", 404, o.Payload)
}

func (o *GetTeamRunNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{entity}/runs/{uuid}][%d] getTeamRunNotFound  %+v", 404, o.Payload)
}

func (o *GetTeamRunNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetTeamRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamRunDefault creates a GetTeamRunDefault with default headers values
func NewGetTeamRunDefault(code int) *GetTeamRunDefault {
	return &GetTeamRunDefault{
		_statusCode: code,
	}
}

/*
GetTeamRunDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetTeamRunDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this get team run default response has a 2xx status code
func (o *GetTeamRunDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get team run default response has a 3xx status code
func (o *GetTeamRunDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get team run default response has a 4xx status code
func (o *GetTeamRunDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get team run default response has a 5xx status code
func (o *GetTeamRunDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get team run default response a status code equal to that given
func (o *GetTeamRunDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get team run default response
func (o *GetTeamRunDefault) Code() int {
	return o._statusCode
}

func (o *GetTeamRunDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{entity}/runs/{uuid}][%d] GetTeamRun default  %+v", o._statusCode, o.Payload)
}

func (o *GetTeamRunDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{entity}/runs/{uuid}][%d] GetTeamRun default  %+v", o._statusCode, o.Payload)
}

func (o *GetTeamRunDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetTeamRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
