// Code generated by go-swagger; DO NOT EDIT.

package tags_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/go/http_client/v1/service_model"
)

// CreateTagReader is a Reader for the CreateTag structure.
type CreateTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCreateTagNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateTagForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateTagDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateTagOK creates a CreateTagOK with default headers values
func NewCreateTagOK() *CreateTagOK {
	return &CreateTagOK{}
}

/*
CreateTagOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateTagOK struct {
	Payload *service_model.V1Tag
}

// IsSuccess returns true when this create tag o k response has a 2xx status code
func (o *CreateTagOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create tag o k response has a 3xx status code
func (o *CreateTagOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tag o k response has a 4xx status code
func (o *CreateTagOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create tag o k response has a 5xx status code
func (o *CreateTagOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create tag o k response a status code equal to that given
func (o *CreateTagOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create tag o k response
func (o *CreateTagOK) Code() int {
	return 200
}

func (o *CreateTagOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/tags][%d] createTagOK  %+v", 200, o.Payload)
}

func (o *CreateTagOK) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/tags][%d] createTagOK  %+v", 200, o.Payload)
}

func (o *CreateTagOK) GetPayload() *service_model.V1Tag {
	return o.Payload
}

func (o *CreateTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTagNoContent creates a CreateTagNoContent with default headers values
func NewCreateTagNoContent() *CreateTagNoContent {
	return &CreateTagNoContent{}
}

/*
CreateTagNoContent describes a response with status code 204, with default header values.

No content.
*/
type CreateTagNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this create tag no content response has a 2xx status code
func (o *CreateTagNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create tag no content response has a 3xx status code
func (o *CreateTagNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tag no content response has a 4xx status code
func (o *CreateTagNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this create tag no content response has a 5xx status code
func (o *CreateTagNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this create tag no content response a status code equal to that given
func (o *CreateTagNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the create tag no content response
func (o *CreateTagNoContent) Code() int {
	return 204
}

func (o *CreateTagNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/tags][%d] createTagNoContent  %+v", 204, o.Payload)
}

func (o *CreateTagNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/tags][%d] createTagNoContent  %+v", 204, o.Payload)
}

func (o *CreateTagNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateTagNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTagForbidden creates a CreateTagForbidden with default headers values
func NewCreateTagForbidden() *CreateTagForbidden {
	return &CreateTagForbidden{}
}

/*
CreateTagForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type CreateTagForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this create tag forbidden response has a 2xx status code
func (o *CreateTagForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create tag forbidden response has a 3xx status code
func (o *CreateTagForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tag forbidden response has a 4xx status code
func (o *CreateTagForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create tag forbidden response has a 5xx status code
func (o *CreateTagForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create tag forbidden response a status code equal to that given
func (o *CreateTagForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create tag forbidden response
func (o *CreateTagForbidden) Code() int {
	return 403
}

func (o *CreateTagForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/tags][%d] createTagForbidden  %+v", 403, o.Payload)
}

func (o *CreateTagForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/tags][%d] createTagForbidden  %+v", 403, o.Payload)
}

func (o *CreateTagForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateTagForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTagNotFound creates a CreateTagNotFound with default headers values
func NewCreateTagNotFound() *CreateTagNotFound {
	return &CreateTagNotFound{}
}

/*
CreateTagNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type CreateTagNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this create tag not found response has a 2xx status code
func (o *CreateTagNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create tag not found response has a 3xx status code
func (o *CreateTagNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tag not found response has a 4xx status code
func (o *CreateTagNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create tag not found response has a 5xx status code
func (o *CreateTagNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create tag not found response a status code equal to that given
func (o *CreateTagNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create tag not found response
func (o *CreateTagNotFound) Code() int {
	return 404
}

func (o *CreateTagNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/tags][%d] createTagNotFound  %+v", 404, o.Payload)
}

func (o *CreateTagNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/tags][%d] createTagNotFound  %+v", 404, o.Payload)
}

func (o *CreateTagNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTagDefault creates a CreateTagDefault with default headers values
func NewCreateTagDefault(code int) *CreateTagDefault {
	return &CreateTagDefault{
		_statusCode: code,
	}
}

/*
CreateTagDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateTagDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this create tag default response has a 2xx status code
func (o *CreateTagDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create tag default response has a 3xx status code
func (o *CreateTagDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create tag default response has a 4xx status code
func (o *CreateTagDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create tag default response has a 5xx status code
func (o *CreateTagDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create tag default response a status code equal to that given
func (o *CreateTagDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create tag default response
func (o *CreateTagDefault) Code() int {
	return o._statusCode
}

func (o *CreateTagDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/tags][%d] CreateTag default  %+v", o._statusCode, o.Payload)
}

func (o *CreateTagDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/tags][%d] CreateTag default  %+v", o._statusCode, o.Payload)
}

func (o *CreateTagDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *CreateTagDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
