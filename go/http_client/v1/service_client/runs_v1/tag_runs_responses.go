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

// TagRunsReader is a Reader for the TagRuns structure.
type TagRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TagRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTagRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewTagRunsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewTagRunsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTagRunsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewTagRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTagRunsOK creates a TagRunsOK with default headers values
func NewTagRunsOK() *TagRunsOK {
	return &TagRunsOK{}
}

/*
TagRunsOK describes a response with status code 200, with default header values.

A successful response.
*/
type TagRunsOK struct {
}

// IsSuccess returns true when this tag runs o k response has a 2xx status code
func (o *TagRunsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tag runs o k response has a 3xx status code
func (o *TagRunsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tag runs o k response has a 4xx status code
func (o *TagRunsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tag runs o k response has a 5xx status code
func (o *TagRunsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tag runs o k response a status code equal to that given
func (o *TagRunsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tag runs o k response
func (o *TagRunsOK) Code() int {
	return 200
}

func (o *TagRunsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/runs/tag][%d] tagRunsOK ", 200)
}

func (o *TagRunsOK) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/runs/tag][%d] tagRunsOK ", 200)
}

func (o *TagRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTagRunsNoContent creates a TagRunsNoContent with default headers values
func NewTagRunsNoContent() *TagRunsNoContent {
	return &TagRunsNoContent{}
}

/*
TagRunsNoContent describes a response with status code 204, with default header values.

No content.
*/
type TagRunsNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this tag runs no content response has a 2xx status code
func (o *TagRunsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tag runs no content response has a 3xx status code
func (o *TagRunsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tag runs no content response has a 4xx status code
func (o *TagRunsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this tag runs no content response has a 5xx status code
func (o *TagRunsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this tag runs no content response a status code equal to that given
func (o *TagRunsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the tag runs no content response
func (o *TagRunsNoContent) Code() int {
	return 204
}

func (o *TagRunsNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/runs/tag][%d] tagRunsNoContent  %+v", 204, o.Payload)
}

func (o *TagRunsNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/runs/tag][%d] tagRunsNoContent  %+v", 204, o.Payload)
}

func (o *TagRunsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *TagRunsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTagRunsForbidden creates a TagRunsForbidden with default headers values
func NewTagRunsForbidden() *TagRunsForbidden {
	return &TagRunsForbidden{}
}

/*
TagRunsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type TagRunsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this tag runs forbidden response has a 2xx status code
func (o *TagRunsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tag runs forbidden response has a 3xx status code
func (o *TagRunsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tag runs forbidden response has a 4xx status code
func (o *TagRunsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this tag runs forbidden response has a 5xx status code
func (o *TagRunsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this tag runs forbidden response a status code equal to that given
func (o *TagRunsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the tag runs forbidden response
func (o *TagRunsForbidden) Code() int {
	return 403
}

func (o *TagRunsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/runs/tag][%d] tagRunsForbidden  %+v", 403, o.Payload)
}

func (o *TagRunsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/runs/tag][%d] tagRunsForbidden  %+v", 403, o.Payload)
}

func (o *TagRunsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *TagRunsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTagRunsNotFound creates a TagRunsNotFound with default headers values
func NewTagRunsNotFound() *TagRunsNotFound {
	return &TagRunsNotFound{}
}

/*
TagRunsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type TagRunsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this tag runs not found response has a 2xx status code
func (o *TagRunsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tag runs not found response has a 3xx status code
func (o *TagRunsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tag runs not found response has a 4xx status code
func (o *TagRunsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this tag runs not found response has a 5xx status code
func (o *TagRunsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this tag runs not found response a status code equal to that given
func (o *TagRunsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the tag runs not found response
func (o *TagRunsNotFound) Code() int {
	return 404
}

func (o *TagRunsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/runs/tag][%d] tagRunsNotFound  %+v", 404, o.Payload)
}

func (o *TagRunsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/runs/tag][%d] tagRunsNotFound  %+v", 404, o.Payload)
}

func (o *TagRunsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *TagRunsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTagRunsDefault creates a TagRunsDefault with default headers values
func NewTagRunsDefault(code int) *TagRunsDefault {
	return &TagRunsDefault{
		_statusCode: code,
	}
}

/*
TagRunsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type TagRunsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this tag runs default response has a 2xx status code
func (o *TagRunsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tag runs default response has a 3xx status code
func (o *TagRunsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tag runs default response has a 4xx status code
func (o *TagRunsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tag runs default response has a 5xx status code
func (o *TagRunsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tag runs default response a status code equal to that given
func (o *TagRunsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the tag runs default response
func (o *TagRunsDefault) Code() int {
	return o._statusCode
}

func (o *TagRunsDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/runs/tag][%d] TagRuns default  %+v", o._statusCode, o.Payload)
}

func (o *TagRunsDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/runs/tag][%d] TagRuns default  %+v", o._statusCode, o.Payload)
}

func (o *TagRunsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *TagRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
