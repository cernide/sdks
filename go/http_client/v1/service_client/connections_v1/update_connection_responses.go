// Code generated by go-swagger; DO NOT EDIT.

package connections_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// UpdateConnectionReader is a Reader for the UpdateConnection structure.
type UpdateConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateConnectionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateConnectionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateConnectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateConnectionOK creates a UpdateConnectionOK with default headers values
func NewUpdateConnectionOK() *UpdateConnectionOK {
	return &UpdateConnectionOK{}
}

/*
UpdateConnectionOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateConnectionOK struct {
	Payload *service_model.V1ConnectionResponse
}

// IsSuccess returns true when this update connection o k response has a 2xx status code
func (o *UpdateConnectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update connection o k response has a 3xx status code
func (o *UpdateConnectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update connection o k response has a 4xx status code
func (o *UpdateConnectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update connection o k response has a 5xx status code
func (o *UpdateConnectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update connection o k response a status code equal to that given
func (o *UpdateConnectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update connection o k response
func (o *UpdateConnectionOK) Code() int {
	return 200
}

func (o *UpdateConnectionOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/connections/{connection.uuid}][%d] updateConnectionOK  %+v", 200, o.Payload)
}

func (o *UpdateConnectionOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/connections/{connection.uuid}][%d] updateConnectionOK  %+v", 200, o.Payload)
}

func (o *UpdateConnectionOK) GetPayload() *service_model.V1ConnectionResponse {
	return o.Payload
}

func (o *UpdateConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ConnectionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConnectionNoContent creates a UpdateConnectionNoContent with default headers values
func NewUpdateConnectionNoContent() *UpdateConnectionNoContent {
	return &UpdateConnectionNoContent{}
}

/*
UpdateConnectionNoContent describes a response with status code 204, with default header values.

No content.
*/
type UpdateConnectionNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this update connection no content response has a 2xx status code
func (o *UpdateConnectionNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update connection no content response has a 3xx status code
func (o *UpdateConnectionNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update connection no content response has a 4xx status code
func (o *UpdateConnectionNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update connection no content response has a 5xx status code
func (o *UpdateConnectionNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update connection no content response a status code equal to that given
func (o *UpdateConnectionNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update connection no content response
func (o *UpdateConnectionNoContent) Code() int {
	return 204
}

func (o *UpdateConnectionNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/connections/{connection.uuid}][%d] updateConnectionNoContent  %+v", 204, o.Payload)
}

func (o *UpdateConnectionNoContent) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/connections/{connection.uuid}][%d] updateConnectionNoContent  %+v", 204, o.Payload)
}

func (o *UpdateConnectionNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateConnectionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConnectionForbidden creates a UpdateConnectionForbidden with default headers values
func NewUpdateConnectionForbidden() *UpdateConnectionForbidden {
	return &UpdateConnectionForbidden{}
}

/*
UpdateConnectionForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type UpdateConnectionForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this update connection forbidden response has a 2xx status code
func (o *UpdateConnectionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update connection forbidden response has a 3xx status code
func (o *UpdateConnectionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update connection forbidden response has a 4xx status code
func (o *UpdateConnectionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update connection forbidden response has a 5xx status code
func (o *UpdateConnectionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update connection forbidden response a status code equal to that given
func (o *UpdateConnectionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update connection forbidden response
func (o *UpdateConnectionForbidden) Code() int {
	return 403
}

func (o *UpdateConnectionForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/connections/{connection.uuid}][%d] updateConnectionForbidden  %+v", 403, o.Payload)
}

func (o *UpdateConnectionForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/connections/{connection.uuid}][%d] updateConnectionForbidden  %+v", 403, o.Payload)
}

func (o *UpdateConnectionForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateConnectionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConnectionNotFound creates a UpdateConnectionNotFound with default headers values
func NewUpdateConnectionNotFound() *UpdateConnectionNotFound {
	return &UpdateConnectionNotFound{}
}

/*
UpdateConnectionNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type UpdateConnectionNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this update connection not found response has a 2xx status code
func (o *UpdateConnectionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update connection not found response has a 3xx status code
func (o *UpdateConnectionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update connection not found response has a 4xx status code
func (o *UpdateConnectionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update connection not found response has a 5xx status code
func (o *UpdateConnectionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update connection not found response a status code equal to that given
func (o *UpdateConnectionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update connection not found response
func (o *UpdateConnectionNotFound) Code() int {
	return 404
}

func (o *UpdateConnectionNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/connections/{connection.uuid}][%d] updateConnectionNotFound  %+v", 404, o.Payload)
}

func (o *UpdateConnectionNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/connections/{connection.uuid}][%d] updateConnectionNotFound  %+v", 404, o.Payload)
}

func (o *UpdateConnectionNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConnectionDefault creates a UpdateConnectionDefault with default headers values
func NewUpdateConnectionDefault(code int) *UpdateConnectionDefault {
	return &UpdateConnectionDefault{
		_statusCode: code,
	}
}

/*
UpdateConnectionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateConnectionDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this update connection default response has a 2xx status code
func (o *UpdateConnectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update connection default response has a 3xx status code
func (o *UpdateConnectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update connection default response has a 4xx status code
func (o *UpdateConnectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update connection default response has a 5xx status code
func (o *UpdateConnectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update connection default response a status code equal to that given
func (o *UpdateConnectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update connection default response
func (o *UpdateConnectionDefault) Code() int {
	return o._statusCode
}

func (o *UpdateConnectionDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/connections/{connection.uuid}][%d] UpdateConnection default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateConnectionDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/connections/{connection.uuid}][%d] UpdateConnection default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateConnectionDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *UpdateConnectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
