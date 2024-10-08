// Code generated by go-swagger; DO NOT EDIT.

package projects_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// BookmarkProjectReader is a Reader for the BookmarkProject structure.
type BookmarkProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BookmarkProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBookmarkProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewBookmarkProjectNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewBookmarkProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBookmarkProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBookmarkProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBookmarkProjectOK creates a BookmarkProjectOK with default headers values
func NewBookmarkProjectOK() *BookmarkProjectOK {
	return &BookmarkProjectOK{}
}

/*
BookmarkProjectOK describes a response with status code 200, with default header values.

A successful response.
*/
type BookmarkProjectOK struct {
}

// IsSuccess returns true when this bookmark project o k response has a 2xx status code
func (o *BookmarkProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this bookmark project o k response has a 3xx status code
func (o *BookmarkProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bookmark project o k response has a 4xx status code
func (o *BookmarkProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this bookmark project o k response has a 5xx status code
func (o *BookmarkProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this bookmark project o k response a status code equal to that given
func (o *BookmarkProjectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the bookmark project o k response
func (o *BookmarkProjectOK) Code() int {
	return 200
}

func (o *BookmarkProjectOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/bookmark][%d] bookmarkProjectOK ", 200)
}

func (o *BookmarkProjectOK) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/bookmark][%d] bookmarkProjectOK ", 200)
}

func (o *BookmarkProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBookmarkProjectNoContent creates a BookmarkProjectNoContent with default headers values
func NewBookmarkProjectNoContent() *BookmarkProjectNoContent {
	return &BookmarkProjectNoContent{}
}

/*
BookmarkProjectNoContent describes a response with status code 204, with default header values.

No content.
*/
type BookmarkProjectNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this bookmark project no content response has a 2xx status code
func (o *BookmarkProjectNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this bookmark project no content response has a 3xx status code
func (o *BookmarkProjectNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bookmark project no content response has a 4xx status code
func (o *BookmarkProjectNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this bookmark project no content response has a 5xx status code
func (o *BookmarkProjectNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this bookmark project no content response a status code equal to that given
func (o *BookmarkProjectNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the bookmark project no content response
func (o *BookmarkProjectNoContent) Code() int {
	return 204
}

func (o *BookmarkProjectNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/bookmark][%d] bookmarkProjectNoContent  %+v", 204, o.Payload)
}

func (o *BookmarkProjectNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/bookmark][%d] bookmarkProjectNoContent  %+v", 204, o.Payload)
}

func (o *BookmarkProjectNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *BookmarkProjectNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarkProjectForbidden creates a BookmarkProjectForbidden with default headers values
func NewBookmarkProjectForbidden() *BookmarkProjectForbidden {
	return &BookmarkProjectForbidden{}
}

/*
BookmarkProjectForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type BookmarkProjectForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this bookmark project forbidden response has a 2xx status code
func (o *BookmarkProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bookmark project forbidden response has a 3xx status code
func (o *BookmarkProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bookmark project forbidden response has a 4xx status code
func (o *BookmarkProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this bookmark project forbidden response has a 5xx status code
func (o *BookmarkProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this bookmark project forbidden response a status code equal to that given
func (o *BookmarkProjectForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the bookmark project forbidden response
func (o *BookmarkProjectForbidden) Code() int {
	return 403
}

func (o *BookmarkProjectForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/bookmark][%d] bookmarkProjectForbidden  %+v", 403, o.Payload)
}

func (o *BookmarkProjectForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/bookmark][%d] bookmarkProjectForbidden  %+v", 403, o.Payload)
}

func (o *BookmarkProjectForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *BookmarkProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarkProjectNotFound creates a BookmarkProjectNotFound with default headers values
func NewBookmarkProjectNotFound() *BookmarkProjectNotFound {
	return &BookmarkProjectNotFound{}
}

/*
BookmarkProjectNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type BookmarkProjectNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this bookmark project not found response has a 2xx status code
func (o *BookmarkProjectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bookmark project not found response has a 3xx status code
func (o *BookmarkProjectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bookmark project not found response has a 4xx status code
func (o *BookmarkProjectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this bookmark project not found response has a 5xx status code
func (o *BookmarkProjectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this bookmark project not found response a status code equal to that given
func (o *BookmarkProjectNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the bookmark project not found response
func (o *BookmarkProjectNotFound) Code() int {
	return 404
}

func (o *BookmarkProjectNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/bookmark][%d] bookmarkProjectNotFound  %+v", 404, o.Payload)
}

func (o *BookmarkProjectNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/bookmark][%d] bookmarkProjectNotFound  %+v", 404, o.Payload)
}

func (o *BookmarkProjectNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *BookmarkProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarkProjectDefault creates a BookmarkProjectDefault with default headers values
func NewBookmarkProjectDefault(code int) *BookmarkProjectDefault {
	return &BookmarkProjectDefault{
		_statusCode: code,
	}
}

/*
BookmarkProjectDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BookmarkProjectDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this bookmark project default response has a 2xx status code
func (o *BookmarkProjectDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this bookmark project default response has a 3xx status code
func (o *BookmarkProjectDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this bookmark project default response has a 4xx status code
func (o *BookmarkProjectDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this bookmark project default response has a 5xx status code
func (o *BookmarkProjectDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this bookmark project default response a status code equal to that given
func (o *BookmarkProjectDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the bookmark project default response
func (o *BookmarkProjectDefault) Code() int {
	return o._statusCode
}

func (o *BookmarkProjectDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/bookmark][%d] BookmarkProject default  %+v", o._statusCode, o.Payload)
}

func (o *BookmarkProjectDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/bookmark][%d] BookmarkProject default  %+v", o._statusCode, o.Payload)
}

func (o *BookmarkProjectDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *BookmarkProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
