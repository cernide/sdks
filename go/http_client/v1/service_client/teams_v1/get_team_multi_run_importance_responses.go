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

// GetTeamMultiRunImportanceReader is a Reader for the GetTeamMultiRunImportance structure.
type GetTeamMultiRunImportanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamMultiRunImportanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamMultiRunImportanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetTeamMultiRunImportanceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetTeamMultiRunImportanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTeamMultiRunImportanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetTeamMultiRunImportanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTeamMultiRunImportanceOK creates a GetTeamMultiRunImportanceOK with default headers values
func NewGetTeamMultiRunImportanceOK() *GetTeamMultiRunImportanceOK {
	return &GetTeamMultiRunImportanceOK{}
}

/*
GetTeamMultiRunImportanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetTeamMultiRunImportanceOK struct {
	Payload *service_model.V1MultiEventsResponse
}

// IsSuccess returns true when this get team multi run importance o k response has a 2xx status code
func (o *GetTeamMultiRunImportanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get team multi run importance o k response has a 3xx status code
func (o *GetTeamMultiRunImportanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team multi run importance o k response has a 4xx status code
func (o *GetTeamMultiRunImportanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team multi run importance o k response has a 5xx status code
func (o *GetTeamMultiRunImportanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get team multi run importance o k response a status code equal to that given
func (o *GetTeamMultiRunImportanceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get team multi run importance o k response
func (o *GetTeamMultiRunImportanceOK) Code() int {
	return 200
}

func (o *GetTeamMultiRunImportanceOK) Error() string {
	return fmt.Sprintf("[POST /streams/v1/{namespace}/orgs/{owner}/teams/{entity}/runs/multi/importance][%d] getTeamMultiRunImportanceOK  %+v", 200, o.Payload)
}

func (o *GetTeamMultiRunImportanceOK) String() string {
	return fmt.Sprintf("[POST /streams/v1/{namespace}/orgs/{owner}/teams/{entity}/runs/multi/importance][%d] getTeamMultiRunImportanceOK  %+v", 200, o.Payload)
}

func (o *GetTeamMultiRunImportanceOK) GetPayload() *service_model.V1MultiEventsResponse {
	return o.Payload
}

func (o *GetTeamMultiRunImportanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1MultiEventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMultiRunImportanceNoContent creates a GetTeamMultiRunImportanceNoContent with default headers values
func NewGetTeamMultiRunImportanceNoContent() *GetTeamMultiRunImportanceNoContent {
	return &GetTeamMultiRunImportanceNoContent{}
}

/*
GetTeamMultiRunImportanceNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetTeamMultiRunImportanceNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get team multi run importance no content response has a 2xx status code
func (o *GetTeamMultiRunImportanceNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get team multi run importance no content response has a 3xx status code
func (o *GetTeamMultiRunImportanceNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team multi run importance no content response has a 4xx status code
func (o *GetTeamMultiRunImportanceNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team multi run importance no content response has a 5xx status code
func (o *GetTeamMultiRunImportanceNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get team multi run importance no content response a status code equal to that given
func (o *GetTeamMultiRunImportanceNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get team multi run importance no content response
func (o *GetTeamMultiRunImportanceNoContent) Code() int {
	return 204
}

func (o *GetTeamMultiRunImportanceNoContent) Error() string {
	return fmt.Sprintf("[POST /streams/v1/{namespace}/orgs/{owner}/teams/{entity}/runs/multi/importance][%d] getTeamMultiRunImportanceNoContent  %+v", 204, o.Payload)
}

func (o *GetTeamMultiRunImportanceNoContent) String() string {
	return fmt.Sprintf("[POST /streams/v1/{namespace}/orgs/{owner}/teams/{entity}/runs/multi/importance][%d] getTeamMultiRunImportanceNoContent  %+v", 204, o.Payload)
}

func (o *GetTeamMultiRunImportanceNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetTeamMultiRunImportanceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMultiRunImportanceForbidden creates a GetTeamMultiRunImportanceForbidden with default headers values
func NewGetTeamMultiRunImportanceForbidden() *GetTeamMultiRunImportanceForbidden {
	return &GetTeamMultiRunImportanceForbidden{}
}

/*
GetTeamMultiRunImportanceForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetTeamMultiRunImportanceForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get team multi run importance forbidden response has a 2xx status code
func (o *GetTeamMultiRunImportanceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team multi run importance forbidden response has a 3xx status code
func (o *GetTeamMultiRunImportanceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team multi run importance forbidden response has a 4xx status code
func (o *GetTeamMultiRunImportanceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team multi run importance forbidden response has a 5xx status code
func (o *GetTeamMultiRunImportanceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get team multi run importance forbidden response a status code equal to that given
func (o *GetTeamMultiRunImportanceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get team multi run importance forbidden response
func (o *GetTeamMultiRunImportanceForbidden) Code() int {
	return 403
}

func (o *GetTeamMultiRunImportanceForbidden) Error() string {
	return fmt.Sprintf("[POST /streams/v1/{namespace}/orgs/{owner}/teams/{entity}/runs/multi/importance][%d] getTeamMultiRunImportanceForbidden  %+v", 403, o.Payload)
}

func (o *GetTeamMultiRunImportanceForbidden) String() string {
	return fmt.Sprintf("[POST /streams/v1/{namespace}/orgs/{owner}/teams/{entity}/runs/multi/importance][%d] getTeamMultiRunImportanceForbidden  %+v", 403, o.Payload)
}

func (o *GetTeamMultiRunImportanceForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetTeamMultiRunImportanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMultiRunImportanceNotFound creates a GetTeamMultiRunImportanceNotFound with default headers values
func NewGetTeamMultiRunImportanceNotFound() *GetTeamMultiRunImportanceNotFound {
	return &GetTeamMultiRunImportanceNotFound{}
}

/*
GetTeamMultiRunImportanceNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetTeamMultiRunImportanceNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get team multi run importance not found response has a 2xx status code
func (o *GetTeamMultiRunImportanceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team multi run importance not found response has a 3xx status code
func (o *GetTeamMultiRunImportanceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team multi run importance not found response has a 4xx status code
func (o *GetTeamMultiRunImportanceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team multi run importance not found response has a 5xx status code
func (o *GetTeamMultiRunImportanceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get team multi run importance not found response a status code equal to that given
func (o *GetTeamMultiRunImportanceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get team multi run importance not found response
func (o *GetTeamMultiRunImportanceNotFound) Code() int {
	return 404
}

func (o *GetTeamMultiRunImportanceNotFound) Error() string {
	return fmt.Sprintf("[POST /streams/v1/{namespace}/orgs/{owner}/teams/{entity}/runs/multi/importance][%d] getTeamMultiRunImportanceNotFound  %+v", 404, o.Payload)
}

func (o *GetTeamMultiRunImportanceNotFound) String() string {
	return fmt.Sprintf("[POST /streams/v1/{namespace}/orgs/{owner}/teams/{entity}/runs/multi/importance][%d] getTeamMultiRunImportanceNotFound  %+v", 404, o.Payload)
}

func (o *GetTeamMultiRunImportanceNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetTeamMultiRunImportanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMultiRunImportanceDefault creates a GetTeamMultiRunImportanceDefault with default headers values
func NewGetTeamMultiRunImportanceDefault(code int) *GetTeamMultiRunImportanceDefault {
	return &GetTeamMultiRunImportanceDefault{
		_statusCode: code,
	}
}

/*
GetTeamMultiRunImportanceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetTeamMultiRunImportanceDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this get team multi run importance default response has a 2xx status code
func (o *GetTeamMultiRunImportanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get team multi run importance default response has a 3xx status code
func (o *GetTeamMultiRunImportanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get team multi run importance default response has a 4xx status code
func (o *GetTeamMultiRunImportanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get team multi run importance default response has a 5xx status code
func (o *GetTeamMultiRunImportanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get team multi run importance default response a status code equal to that given
func (o *GetTeamMultiRunImportanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get team multi run importance default response
func (o *GetTeamMultiRunImportanceDefault) Code() int {
	return o._statusCode
}

func (o *GetTeamMultiRunImportanceDefault) Error() string {
	return fmt.Sprintf("[POST /streams/v1/{namespace}/orgs/{owner}/teams/{entity}/runs/multi/importance][%d] GetTeamMultiRunImportance default  %+v", o._statusCode, o.Payload)
}

func (o *GetTeamMultiRunImportanceDefault) String() string {
	return fmt.Sprintf("[POST /streams/v1/{namespace}/orgs/{owner}/teams/{entity}/runs/multi/importance][%d] GetTeamMultiRunImportance default  %+v", o._statusCode, o.Payload)
}

func (o *GetTeamMultiRunImportanceDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetTeamMultiRunImportanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
