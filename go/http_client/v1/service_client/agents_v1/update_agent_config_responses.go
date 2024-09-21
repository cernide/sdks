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

// UpdateAgentConfigReader is a Reader for the UpdateAgentConfig structure.
type UpdateAgentConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAgentConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAgentConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateAgentConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateAgentConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAgentConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateAgentConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAgentConfigOK creates a UpdateAgentConfigOK with default headers values
func NewUpdateAgentConfigOK() *UpdateAgentConfigOK {
	return &UpdateAgentConfigOK{}
}

/*
UpdateAgentConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateAgentConfigOK struct {
	Payload *service_model.V1Agent
}

// IsSuccess returns true when this update agent config o k response has a 2xx status code
func (o *UpdateAgentConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update agent config o k response has a 3xx status code
func (o *UpdateAgentConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update agent config o k response has a 4xx status code
func (o *UpdateAgentConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update agent config o k response has a 5xx status code
func (o *UpdateAgentConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update agent config o k response a status code equal to that given
func (o *UpdateAgentConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update agent config o k response
func (o *UpdateAgentConfigOK) Code() int {
	return 200
}

func (o *UpdateAgentConfigOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}/config][%d] updateAgentConfigOK  %+v", 200, o.Payload)
}

func (o *UpdateAgentConfigOK) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}/config][%d] updateAgentConfigOK  %+v", 200, o.Payload)
}

func (o *UpdateAgentConfigOK) GetPayload() *service_model.V1Agent {
	return o.Payload
}

func (o *UpdateAgentConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Agent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAgentConfigNoContent creates a UpdateAgentConfigNoContent with default headers values
func NewUpdateAgentConfigNoContent() *UpdateAgentConfigNoContent {
	return &UpdateAgentConfigNoContent{}
}

/*
UpdateAgentConfigNoContent describes a response with status code 204, with default header values.

No content.
*/
type UpdateAgentConfigNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this update agent config no content response has a 2xx status code
func (o *UpdateAgentConfigNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update agent config no content response has a 3xx status code
func (o *UpdateAgentConfigNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update agent config no content response has a 4xx status code
func (o *UpdateAgentConfigNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update agent config no content response has a 5xx status code
func (o *UpdateAgentConfigNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update agent config no content response a status code equal to that given
func (o *UpdateAgentConfigNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update agent config no content response
func (o *UpdateAgentConfigNoContent) Code() int {
	return 204
}

func (o *UpdateAgentConfigNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}/config][%d] updateAgentConfigNoContent  %+v", 204, o.Payload)
}

func (o *UpdateAgentConfigNoContent) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}/config][%d] updateAgentConfigNoContent  %+v", 204, o.Payload)
}

func (o *UpdateAgentConfigNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateAgentConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAgentConfigForbidden creates a UpdateAgentConfigForbidden with default headers values
func NewUpdateAgentConfigForbidden() *UpdateAgentConfigForbidden {
	return &UpdateAgentConfigForbidden{}
}

/*
UpdateAgentConfigForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type UpdateAgentConfigForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this update agent config forbidden response has a 2xx status code
func (o *UpdateAgentConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update agent config forbidden response has a 3xx status code
func (o *UpdateAgentConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update agent config forbidden response has a 4xx status code
func (o *UpdateAgentConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update agent config forbidden response has a 5xx status code
func (o *UpdateAgentConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update agent config forbidden response a status code equal to that given
func (o *UpdateAgentConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update agent config forbidden response
func (o *UpdateAgentConfigForbidden) Code() int {
	return 403
}

func (o *UpdateAgentConfigForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}/config][%d] updateAgentConfigForbidden  %+v", 403, o.Payload)
}

func (o *UpdateAgentConfigForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}/config][%d] updateAgentConfigForbidden  %+v", 403, o.Payload)
}

func (o *UpdateAgentConfigForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateAgentConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAgentConfigNotFound creates a UpdateAgentConfigNotFound with default headers values
func NewUpdateAgentConfigNotFound() *UpdateAgentConfigNotFound {
	return &UpdateAgentConfigNotFound{}
}

/*
UpdateAgentConfigNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type UpdateAgentConfigNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this update agent config not found response has a 2xx status code
func (o *UpdateAgentConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update agent config not found response has a 3xx status code
func (o *UpdateAgentConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update agent config not found response has a 4xx status code
func (o *UpdateAgentConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update agent config not found response has a 5xx status code
func (o *UpdateAgentConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update agent config not found response a status code equal to that given
func (o *UpdateAgentConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update agent config not found response
func (o *UpdateAgentConfigNotFound) Code() int {
	return 404
}

func (o *UpdateAgentConfigNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}/config][%d] updateAgentConfigNotFound  %+v", 404, o.Payload)
}

func (o *UpdateAgentConfigNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}/config][%d] updateAgentConfigNotFound  %+v", 404, o.Payload)
}

func (o *UpdateAgentConfigNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateAgentConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAgentConfigDefault creates a UpdateAgentConfigDefault with default headers values
func NewUpdateAgentConfigDefault(code int) *UpdateAgentConfigDefault {
	return &UpdateAgentConfigDefault{
		_statusCode: code,
	}
}

/*
UpdateAgentConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateAgentConfigDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this update agent config default response has a 2xx status code
func (o *UpdateAgentConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update agent config default response has a 3xx status code
func (o *UpdateAgentConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update agent config default response has a 4xx status code
func (o *UpdateAgentConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update agent config default response has a 5xx status code
func (o *UpdateAgentConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update agent config default response a status code equal to that given
func (o *UpdateAgentConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update agent config default response
func (o *UpdateAgentConfigDefault) Code() int {
	return o._statusCode
}

func (o *UpdateAgentConfigDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}/config][%d] UpdateAgentConfig default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateAgentConfigDefault) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}/config][%d] UpdateAgentConfig default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateAgentConfigDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *UpdateAgentConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
