// Code generated by go-swagger; DO NOT EDIT.

package agents_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// GetAgentReader is a Reader for the GetAgent structure.
type GetAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetAgentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetAgentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAgentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAgentOK creates a GetAgentOK with default headers values
func NewGetAgentOK() *GetAgentOK {
	return &GetAgentOK{}
}

/*
GetAgentOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetAgentOK struct {
	Payload *service_model.V1Agent
}

// IsSuccess returns true when this get agent o k response has a 2xx status code
func (o *GetAgentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get agent o k response has a 3xx status code
func (o *GetAgentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get agent o k response has a 4xx status code
func (o *GetAgentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get agent o k response has a 5xx status code
func (o *GetAgentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get agent o k response a status code equal to that given
func (o *GetAgentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get agent o k response
func (o *GetAgentOK) Code() int {
	return 200
}

func (o *GetAgentOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{uuid}][%d] getAgentOK  %+v", 200, o.Payload)
}

func (o *GetAgentOK) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{uuid}][%d] getAgentOK  %+v", 200, o.Payload)
}

func (o *GetAgentOK) GetPayload() *service_model.V1Agent {
	return o.Payload
}

func (o *GetAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Agent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAgentNoContent creates a GetAgentNoContent with default headers values
func NewGetAgentNoContent() *GetAgentNoContent {
	return &GetAgentNoContent{}
}

/*
GetAgentNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetAgentNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get agent no content response has a 2xx status code
func (o *GetAgentNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get agent no content response has a 3xx status code
func (o *GetAgentNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get agent no content response has a 4xx status code
func (o *GetAgentNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get agent no content response has a 5xx status code
func (o *GetAgentNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get agent no content response a status code equal to that given
func (o *GetAgentNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get agent no content response
func (o *GetAgentNoContent) Code() int {
	return 204
}

func (o *GetAgentNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{uuid}][%d] getAgentNoContent  %+v", 204, o.Payload)
}

func (o *GetAgentNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{uuid}][%d] getAgentNoContent  %+v", 204, o.Payload)
}

func (o *GetAgentNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetAgentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAgentForbidden creates a GetAgentForbidden with default headers values
func NewGetAgentForbidden() *GetAgentForbidden {
	return &GetAgentForbidden{}
}

/*
GetAgentForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetAgentForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get agent forbidden response has a 2xx status code
func (o *GetAgentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get agent forbidden response has a 3xx status code
func (o *GetAgentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get agent forbidden response has a 4xx status code
func (o *GetAgentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get agent forbidden response has a 5xx status code
func (o *GetAgentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get agent forbidden response a status code equal to that given
func (o *GetAgentForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get agent forbidden response
func (o *GetAgentForbidden) Code() int {
	return 403
}

func (o *GetAgentForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{uuid}][%d] getAgentForbidden  %+v", 403, o.Payload)
}

func (o *GetAgentForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{uuid}][%d] getAgentForbidden  %+v", 403, o.Payload)
}

func (o *GetAgentForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetAgentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAgentNotFound creates a GetAgentNotFound with default headers values
func NewGetAgentNotFound() *GetAgentNotFound {
	return &GetAgentNotFound{}
}

/*
GetAgentNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetAgentNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get agent not found response has a 2xx status code
func (o *GetAgentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get agent not found response has a 3xx status code
func (o *GetAgentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get agent not found response has a 4xx status code
func (o *GetAgentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get agent not found response has a 5xx status code
func (o *GetAgentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get agent not found response a status code equal to that given
func (o *GetAgentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get agent not found response
func (o *GetAgentNotFound) Code() int {
	return 404
}

func (o *GetAgentNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{uuid}][%d] getAgentNotFound  %+v", 404, o.Payload)
}

func (o *GetAgentNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{uuid}][%d] getAgentNotFound  %+v", 404, o.Payload)
}

func (o *GetAgentNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetAgentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAgentDefault creates a GetAgentDefault with default headers values
func NewGetAgentDefault(code int) *GetAgentDefault {
	return &GetAgentDefault{
		_statusCode: code,
	}
}

/*
GetAgentDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetAgentDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this get agent default response has a 2xx status code
func (o *GetAgentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get agent default response has a 3xx status code
func (o *GetAgentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get agent default response has a 4xx status code
func (o *GetAgentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get agent default response has a 5xx status code
func (o *GetAgentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get agent default response a status code equal to that given
func (o *GetAgentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get agent default response
func (o *GetAgentDefault) Code() int {
	return o._statusCode
}

func (o *GetAgentDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{uuid}][%d] GetAgent default  %+v", o._statusCode, o.Payload)
}

func (o *GetAgentDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{uuid}][%d] GetAgent default  %+v", o._statusCode, o.Payload)
}

func (o *GetAgentDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
