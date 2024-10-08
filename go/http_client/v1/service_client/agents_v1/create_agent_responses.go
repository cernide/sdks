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

// CreateAgentReader is a Reader for the CreateAgent structure.
type CreateAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCreateAgentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateAgentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateAgentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAgentOK creates a CreateAgentOK with default headers values
func NewCreateAgentOK() *CreateAgentOK {
	return &CreateAgentOK{}
}

/*
CreateAgentOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateAgentOK struct {
	Payload *service_model.V1Agent
}

// IsSuccess returns true when this create agent o k response has a 2xx status code
func (o *CreateAgentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create agent o k response has a 3xx status code
func (o *CreateAgentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create agent o k response has a 4xx status code
func (o *CreateAgentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create agent o k response has a 5xx status code
func (o *CreateAgentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create agent o k response a status code equal to that given
func (o *CreateAgentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create agent o k response
func (o *CreateAgentOK) Code() int {
	return 200
}

func (o *CreateAgentOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents][%d] createAgentOK  %+v", 200, o.Payload)
}

func (o *CreateAgentOK) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents][%d] createAgentOK  %+v", 200, o.Payload)
}

func (o *CreateAgentOK) GetPayload() *service_model.V1Agent {
	return o.Payload
}

func (o *CreateAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Agent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAgentNoContent creates a CreateAgentNoContent with default headers values
func NewCreateAgentNoContent() *CreateAgentNoContent {
	return &CreateAgentNoContent{}
}

/*
CreateAgentNoContent describes a response with status code 204, with default header values.

No content.
*/
type CreateAgentNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this create agent no content response has a 2xx status code
func (o *CreateAgentNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create agent no content response has a 3xx status code
func (o *CreateAgentNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create agent no content response has a 4xx status code
func (o *CreateAgentNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this create agent no content response has a 5xx status code
func (o *CreateAgentNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this create agent no content response a status code equal to that given
func (o *CreateAgentNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the create agent no content response
func (o *CreateAgentNoContent) Code() int {
	return 204
}

func (o *CreateAgentNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents][%d] createAgentNoContent  %+v", 204, o.Payload)
}

func (o *CreateAgentNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents][%d] createAgentNoContent  %+v", 204, o.Payload)
}

func (o *CreateAgentNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateAgentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAgentForbidden creates a CreateAgentForbidden with default headers values
func NewCreateAgentForbidden() *CreateAgentForbidden {
	return &CreateAgentForbidden{}
}

/*
CreateAgentForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type CreateAgentForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this create agent forbidden response has a 2xx status code
func (o *CreateAgentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create agent forbidden response has a 3xx status code
func (o *CreateAgentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create agent forbidden response has a 4xx status code
func (o *CreateAgentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create agent forbidden response has a 5xx status code
func (o *CreateAgentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create agent forbidden response a status code equal to that given
func (o *CreateAgentForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create agent forbidden response
func (o *CreateAgentForbidden) Code() int {
	return 403
}

func (o *CreateAgentForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents][%d] createAgentForbidden  %+v", 403, o.Payload)
}

func (o *CreateAgentForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents][%d] createAgentForbidden  %+v", 403, o.Payload)
}

func (o *CreateAgentForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateAgentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAgentNotFound creates a CreateAgentNotFound with default headers values
func NewCreateAgentNotFound() *CreateAgentNotFound {
	return &CreateAgentNotFound{}
}

/*
CreateAgentNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type CreateAgentNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this create agent not found response has a 2xx status code
func (o *CreateAgentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create agent not found response has a 3xx status code
func (o *CreateAgentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create agent not found response has a 4xx status code
func (o *CreateAgentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create agent not found response has a 5xx status code
func (o *CreateAgentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create agent not found response a status code equal to that given
func (o *CreateAgentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create agent not found response
func (o *CreateAgentNotFound) Code() int {
	return 404
}

func (o *CreateAgentNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents][%d] createAgentNotFound  %+v", 404, o.Payload)
}

func (o *CreateAgentNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents][%d] createAgentNotFound  %+v", 404, o.Payload)
}

func (o *CreateAgentNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateAgentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAgentDefault creates a CreateAgentDefault with default headers values
func NewCreateAgentDefault(code int) *CreateAgentDefault {
	return &CreateAgentDefault{
		_statusCode: code,
	}
}

/*
CreateAgentDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateAgentDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this create agent default response has a 2xx status code
func (o *CreateAgentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create agent default response has a 3xx status code
func (o *CreateAgentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create agent default response has a 4xx status code
func (o *CreateAgentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create agent default response has a 5xx status code
func (o *CreateAgentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create agent default response a status code equal to that given
func (o *CreateAgentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create agent default response
func (o *CreateAgentDefault) Code() int {
	return o._statusCode
}

func (o *CreateAgentDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents][%d] CreateAgent default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAgentDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents][%d] CreateAgent default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAgentDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *CreateAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
