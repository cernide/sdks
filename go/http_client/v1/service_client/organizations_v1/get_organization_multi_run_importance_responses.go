// Code generated by go-swagger; DO NOT EDIT.

package organizations_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// GetOrganizationMultiRunImportanceReader is a Reader for the GetOrganizationMultiRunImportance structure.
type GetOrganizationMultiRunImportanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationMultiRunImportanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationMultiRunImportanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetOrganizationMultiRunImportanceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetOrganizationMultiRunImportanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationMultiRunImportanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetOrganizationMultiRunImportanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOrganizationMultiRunImportanceOK creates a GetOrganizationMultiRunImportanceOK with default headers values
func NewGetOrganizationMultiRunImportanceOK() *GetOrganizationMultiRunImportanceOK {
	return &GetOrganizationMultiRunImportanceOK{}
}

/*
GetOrganizationMultiRunImportanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetOrganizationMultiRunImportanceOK struct {
	Payload *service_model.V1MultiEventsResponse
}

// IsSuccess returns true when this get organization multi run importance o k response has a 2xx status code
func (o *GetOrganizationMultiRunImportanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization multi run importance o k response has a 3xx status code
func (o *GetOrganizationMultiRunImportanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization multi run importance o k response has a 4xx status code
func (o *GetOrganizationMultiRunImportanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization multi run importance o k response has a 5xx status code
func (o *GetOrganizationMultiRunImportanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization multi run importance o k response a status code equal to that given
func (o *GetOrganizationMultiRunImportanceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization multi run importance o k response
func (o *GetOrganizationMultiRunImportanceOK) Code() int {
	return 200
}

func (o *GetOrganizationMultiRunImportanceOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/multi/importance][%d] getOrganizationMultiRunImportanceOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationMultiRunImportanceOK) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/multi/importance][%d] getOrganizationMultiRunImportanceOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationMultiRunImportanceOK) GetPayload() *service_model.V1MultiEventsResponse {
	return o.Payload
}

func (o *GetOrganizationMultiRunImportanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1MultiEventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationMultiRunImportanceNoContent creates a GetOrganizationMultiRunImportanceNoContent with default headers values
func NewGetOrganizationMultiRunImportanceNoContent() *GetOrganizationMultiRunImportanceNoContent {
	return &GetOrganizationMultiRunImportanceNoContent{}
}

/*
GetOrganizationMultiRunImportanceNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetOrganizationMultiRunImportanceNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get organization multi run importance no content response has a 2xx status code
func (o *GetOrganizationMultiRunImportanceNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization multi run importance no content response has a 3xx status code
func (o *GetOrganizationMultiRunImportanceNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization multi run importance no content response has a 4xx status code
func (o *GetOrganizationMultiRunImportanceNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization multi run importance no content response has a 5xx status code
func (o *GetOrganizationMultiRunImportanceNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization multi run importance no content response a status code equal to that given
func (o *GetOrganizationMultiRunImportanceNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get organization multi run importance no content response
func (o *GetOrganizationMultiRunImportanceNoContent) Code() int {
	return 204
}

func (o *GetOrganizationMultiRunImportanceNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/multi/importance][%d] getOrganizationMultiRunImportanceNoContent  %+v", 204, o.Payload)
}

func (o *GetOrganizationMultiRunImportanceNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/multi/importance][%d] getOrganizationMultiRunImportanceNoContent  %+v", 204, o.Payload)
}

func (o *GetOrganizationMultiRunImportanceNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationMultiRunImportanceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationMultiRunImportanceForbidden creates a GetOrganizationMultiRunImportanceForbidden with default headers values
func NewGetOrganizationMultiRunImportanceForbidden() *GetOrganizationMultiRunImportanceForbidden {
	return &GetOrganizationMultiRunImportanceForbidden{}
}

/*
GetOrganizationMultiRunImportanceForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetOrganizationMultiRunImportanceForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get organization multi run importance forbidden response has a 2xx status code
func (o *GetOrganizationMultiRunImportanceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization multi run importance forbidden response has a 3xx status code
func (o *GetOrganizationMultiRunImportanceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization multi run importance forbidden response has a 4xx status code
func (o *GetOrganizationMultiRunImportanceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization multi run importance forbidden response has a 5xx status code
func (o *GetOrganizationMultiRunImportanceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization multi run importance forbidden response a status code equal to that given
func (o *GetOrganizationMultiRunImportanceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get organization multi run importance forbidden response
func (o *GetOrganizationMultiRunImportanceForbidden) Code() int {
	return 403
}

func (o *GetOrganizationMultiRunImportanceForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/multi/importance][%d] getOrganizationMultiRunImportanceForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationMultiRunImportanceForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/multi/importance][%d] getOrganizationMultiRunImportanceForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationMultiRunImportanceForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationMultiRunImportanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationMultiRunImportanceNotFound creates a GetOrganizationMultiRunImportanceNotFound with default headers values
func NewGetOrganizationMultiRunImportanceNotFound() *GetOrganizationMultiRunImportanceNotFound {
	return &GetOrganizationMultiRunImportanceNotFound{}
}

/*
GetOrganizationMultiRunImportanceNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetOrganizationMultiRunImportanceNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get organization multi run importance not found response has a 2xx status code
func (o *GetOrganizationMultiRunImportanceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization multi run importance not found response has a 3xx status code
func (o *GetOrganizationMultiRunImportanceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization multi run importance not found response has a 4xx status code
func (o *GetOrganizationMultiRunImportanceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization multi run importance not found response has a 5xx status code
func (o *GetOrganizationMultiRunImportanceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization multi run importance not found response a status code equal to that given
func (o *GetOrganizationMultiRunImportanceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get organization multi run importance not found response
func (o *GetOrganizationMultiRunImportanceNotFound) Code() int {
	return 404
}

func (o *GetOrganizationMultiRunImportanceNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/multi/importance][%d] getOrganizationMultiRunImportanceNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationMultiRunImportanceNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/multi/importance][%d] getOrganizationMultiRunImportanceNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationMultiRunImportanceNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationMultiRunImportanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationMultiRunImportanceDefault creates a GetOrganizationMultiRunImportanceDefault with default headers values
func NewGetOrganizationMultiRunImportanceDefault(code int) *GetOrganizationMultiRunImportanceDefault {
	return &GetOrganizationMultiRunImportanceDefault{
		_statusCode: code,
	}
}

/*
GetOrganizationMultiRunImportanceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetOrganizationMultiRunImportanceDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this get organization multi run importance default response has a 2xx status code
func (o *GetOrganizationMultiRunImportanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get organization multi run importance default response has a 3xx status code
func (o *GetOrganizationMultiRunImportanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get organization multi run importance default response has a 4xx status code
func (o *GetOrganizationMultiRunImportanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get organization multi run importance default response has a 5xx status code
func (o *GetOrganizationMultiRunImportanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get organization multi run importance default response a status code equal to that given
func (o *GetOrganizationMultiRunImportanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get organization multi run importance default response
func (o *GetOrganizationMultiRunImportanceDefault) Code() int {
	return o._statusCode
}

func (o *GetOrganizationMultiRunImportanceDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/multi/importance][%d] GetOrganizationMultiRunImportance default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrganizationMultiRunImportanceDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/multi/importance][%d] GetOrganizationMultiRunImportance default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrganizationMultiRunImportanceDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetOrganizationMultiRunImportanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
