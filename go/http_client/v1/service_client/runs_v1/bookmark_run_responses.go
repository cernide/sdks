// Code generated by go-swagger; DO NOT EDIT.

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// BookmarkRunReader is a Reader for the BookmarkRun structure.
type BookmarkRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BookmarkRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBookmarkRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewBookmarkRunNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewBookmarkRunForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBookmarkRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBookmarkRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBookmarkRunOK creates a BookmarkRunOK with default headers values
func NewBookmarkRunOK() *BookmarkRunOK {
	return &BookmarkRunOK{}
}

/*
BookmarkRunOK describes a response with status code 200, with default header values.

A successful response.
*/
type BookmarkRunOK struct {
}

// IsSuccess returns true when this bookmark run o k response has a 2xx status code
func (o *BookmarkRunOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this bookmark run o k response has a 3xx status code
func (o *BookmarkRunOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bookmark run o k response has a 4xx status code
func (o *BookmarkRunOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this bookmark run o k response has a 5xx status code
func (o *BookmarkRunOK) IsServerError() bool {
	return false
}

// IsCode returns true when this bookmark run o k response a status code equal to that given
func (o *BookmarkRunOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the bookmark run o k response
func (o *BookmarkRunOK) Code() int {
	return 200
}

func (o *BookmarkRunOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/bookmark][%d] bookmarkRunOK ", 200)
}

func (o *BookmarkRunOK) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/bookmark][%d] bookmarkRunOK ", 200)
}

func (o *BookmarkRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBookmarkRunNoContent creates a BookmarkRunNoContent with default headers values
func NewBookmarkRunNoContent() *BookmarkRunNoContent {
	return &BookmarkRunNoContent{}
}

/*
BookmarkRunNoContent describes a response with status code 204, with default header values.

No content.
*/
type BookmarkRunNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this bookmark run no content response has a 2xx status code
func (o *BookmarkRunNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this bookmark run no content response has a 3xx status code
func (o *BookmarkRunNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bookmark run no content response has a 4xx status code
func (o *BookmarkRunNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this bookmark run no content response has a 5xx status code
func (o *BookmarkRunNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this bookmark run no content response a status code equal to that given
func (o *BookmarkRunNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the bookmark run no content response
func (o *BookmarkRunNoContent) Code() int {
	return 204
}

func (o *BookmarkRunNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/bookmark][%d] bookmarkRunNoContent  %+v", 204, o.Payload)
}

func (o *BookmarkRunNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/bookmark][%d] bookmarkRunNoContent  %+v", 204, o.Payload)
}

func (o *BookmarkRunNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *BookmarkRunNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarkRunForbidden creates a BookmarkRunForbidden with default headers values
func NewBookmarkRunForbidden() *BookmarkRunForbidden {
	return &BookmarkRunForbidden{}
}

/*
BookmarkRunForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type BookmarkRunForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this bookmark run forbidden response has a 2xx status code
func (o *BookmarkRunForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bookmark run forbidden response has a 3xx status code
func (o *BookmarkRunForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bookmark run forbidden response has a 4xx status code
func (o *BookmarkRunForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this bookmark run forbidden response has a 5xx status code
func (o *BookmarkRunForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this bookmark run forbidden response a status code equal to that given
func (o *BookmarkRunForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the bookmark run forbidden response
func (o *BookmarkRunForbidden) Code() int {
	return 403
}

func (o *BookmarkRunForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/bookmark][%d] bookmarkRunForbidden  %+v", 403, o.Payload)
}

func (o *BookmarkRunForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/bookmark][%d] bookmarkRunForbidden  %+v", 403, o.Payload)
}

func (o *BookmarkRunForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *BookmarkRunForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarkRunNotFound creates a BookmarkRunNotFound with default headers values
func NewBookmarkRunNotFound() *BookmarkRunNotFound {
	return &BookmarkRunNotFound{}
}

/*
BookmarkRunNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type BookmarkRunNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this bookmark run not found response has a 2xx status code
func (o *BookmarkRunNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bookmark run not found response has a 3xx status code
func (o *BookmarkRunNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bookmark run not found response has a 4xx status code
func (o *BookmarkRunNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this bookmark run not found response has a 5xx status code
func (o *BookmarkRunNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this bookmark run not found response a status code equal to that given
func (o *BookmarkRunNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the bookmark run not found response
func (o *BookmarkRunNotFound) Code() int {
	return 404
}

func (o *BookmarkRunNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/bookmark][%d] bookmarkRunNotFound  %+v", 404, o.Payload)
}

func (o *BookmarkRunNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/bookmark][%d] bookmarkRunNotFound  %+v", 404, o.Payload)
}

func (o *BookmarkRunNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *BookmarkRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarkRunDefault creates a BookmarkRunDefault with default headers values
func NewBookmarkRunDefault(code int) *BookmarkRunDefault {
	return &BookmarkRunDefault{
		_statusCode: code,
	}
}

/*
BookmarkRunDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BookmarkRunDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this bookmark run default response has a 2xx status code
func (o *BookmarkRunDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this bookmark run default response has a 3xx status code
func (o *BookmarkRunDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this bookmark run default response has a 4xx status code
func (o *BookmarkRunDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this bookmark run default response has a 5xx status code
func (o *BookmarkRunDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this bookmark run default response a status code equal to that given
func (o *BookmarkRunDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the bookmark run default response
func (o *BookmarkRunDefault) Code() int {
	return o._statusCode
}

func (o *BookmarkRunDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/bookmark][%d] BookmarkRun default  %+v", o._statusCode, o.Payload)
}

func (o *BookmarkRunDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/bookmark][%d] BookmarkRun default  %+v", o._statusCode, o.Payload)
}

func (o *BookmarkRunDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *BookmarkRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
