// Code generated by go-swagger; DO NOT EDIT.

package auth_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// TrialReader is a Reader for the Trial structure.
type TrialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TrialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTrialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewTrialNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewTrialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTrialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewTrialDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTrialOK creates a TrialOK with default headers values
func NewTrialOK() *TrialOK {
	return &TrialOK{}
}

/*
TrialOK describes a response with status code 200, with default header values.

A successful response.
*/
type TrialOK struct {
}

// IsSuccess returns true when this trial o k response has a 2xx status code
func (o *TrialOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this trial o k response has a 3xx status code
func (o *TrialOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trial o k response has a 4xx status code
func (o *TrialOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this trial o k response has a 5xx status code
func (o *TrialOK) IsServerError() bool {
	return false
}

// IsCode returns true when this trial o k response a status code equal to that given
func (o *TrialOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the trial o k response
func (o *TrialOK) Code() int {
	return 200
}

func (o *TrialOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/auth/trial][%d] trialOK ", 200)
}

func (o *TrialOK) String() string {
	return fmt.Sprintf("[POST /api/v1/auth/trial][%d] trialOK ", 200)
}

func (o *TrialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTrialNoContent creates a TrialNoContent with default headers values
func NewTrialNoContent() *TrialNoContent {
	return &TrialNoContent{}
}

/*
TrialNoContent describes a response with status code 204, with default header values.

No content.
*/
type TrialNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this trial no content response has a 2xx status code
func (o *TrialNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this trial no content response has a 3xx status code
func (o *TrialNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trial no content response has a 4xx status code
func (o *TrialNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this trial no content response has a 5xx status code
func (o *TrialNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this trial no content response a status code equal to that given
func (o *TrialNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the trial no content response
func (o *TrialNoContent) Code() int {
	return 204
}

func (o *TrialNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/auth/trial][%d] trialNoContent  %+v", 204, o.Payload)
}

func (o *TrialNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/auth/trial][%d] trialNoContent  %+v", 204, o.Payload)
}

func (o *TrialNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *TrialNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTrialForbidden creates a TrialForbidden with default headers values
func NewTrialForbidden() *TrialForbidden {
	return &TrialForbidden{}
}

/*
TrialForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type TrialForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this trial forbidden response has a 2xx status code
func (o *TrialForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this trial forbidden response has a 3xx status code
func (o *TrialForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trial forbidden response has a 4xx status code
func (o *TrialForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this trial forbidden response has a 5xx status code
func (o *TrialForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this trial forbidden response a status code equal to that given
func (o *TrialForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the trial forbidden response
func (o *TrialForbidden) Code() int {
	return 403
}

func (o *TrialForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/auth/trial][%d] trialForbidden  %+v", 403, o.Payload)
}

func (o *TrialForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/auth/trial][%d] trialForbidden  %+v", 403, o.Payload)
}

func (o *TrialForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *TrialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTrialNotFound creates a TrialNotFound with default headers values
func NewTrialNotFound() *TrialNotFound {
	return &TrialNotFound{}
}

/*
TrialNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type TrialNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this trial not found response has a 2xx status code
func (o *TrialNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this trial not found response has a 3xx status code
func (o *TrialNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trial not found response has a 4xx status code
func (o *TrialNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this trial not found response has a 5xx status code
func (o *TrialNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this trial not found response a status code equal to that given
func (o *TrialNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the trial not found response
func (o *TrialNotFound) Code() int {
	return 404
}

func (o *TrialNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/auth/trial][%d] trialNotFound  %+v", 404, o.Payload)
}

func (o *TrialNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/auth/trial][%d] trialNotFound  %+v", 404, o.Payload)
}

func (o *TrialNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *TrialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTrialDefault creates a TrialDefault with default headers values
func NewTrialDefault(code int) *TrialDefault {
	return &TrialDefault{
		_statusCode: code,
	}
}

/*
TrialDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type TrialDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this trial default response has a 2xx status code
func (o *TrialDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this trial default response has a 3xx status code
func (o *TrialDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this trial default response has a 4xx status code
func (o *TrialDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this trial default response has a 5xx status code
func (o *TrialDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this trial default response a status code equal to that given
func (o *TrialDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the trial default response
func (o *TrialDefault) Code() int {
	return o._statusCode
}

func (o *TrialDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/auth/trial][%d] Trial default  %+v", o._statusCode, o.Payload)
}

func (o *TrialDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/auth/trial][%d] Trial default  %+v", o._statusCode, o.Payload)
}

func (o *TrialDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *TrialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
