// Code generated by go-swagger; DO NOT EDIT.

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/go/http_client/v1/service_model"
)

// TransferRunsReader is a Reader for the TransferRuns structure.
type TransferRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TransferRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTransferRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewTransferRunsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewTransferRunsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTransferRunsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewTransferRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTransferRunsOK creates a TransferRunsOK with default headers values
func NewTransferRunsOK() *TransferRunsOK {
	return &TransferRunsOK{}
}

/*
TransferRunsOK describes a response with status code 200, with default header values.

A successful response.
*/
type TransferRunsOK struct {
}

// IsSuccess returns true when this transfer runs o k response has a 2xx status code
func (o *TransferRunsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this transfer runs o k response has a 3xx status code
func (o *TransferRunsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this transfer runs o k response has a 4xx status code
func (o *TransferRunsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this transfer runs o k response has a 5xx status code
func (o *TransferRunsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this transfer runs o k response a status code equal to that given
func (o *TransferRunsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the transfer runs o k response
func (o *TransferRunsOK) Code() int {
	return 200
}

func (o *TransferRunsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/runs/transfer][%d] transferRunsOK ", 200)
}

func (o *TransferRunsOK) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/runs/transfer][%d] transferRunsOK ", 200)
}

func (o *TransferRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTransferRunsNoContent creates a TransferRunsNoContent with default headers values
func NewTransferRunsNoContent() *TransferRunsNoContent {
	return &TransferRunsNoContent{}
}

/*
TransferRunsNoContent describes a response with status code 204, with default header values.

No content.
*/
type TransferRunsNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this transfer runs no content response has a 2xx status code
func (o *TransferRunsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this transfer runs no content response has a 3xx status code
func (o *TransferRunsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this transfer runs no content response has a 4xx status code
func (o *TransferRunsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this transfer runs no content response has a 5xx status code
func (o *TransferRunsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this transfer runs no content response a status code equal to that given
func (o *TransferRunsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the transfer runs no content response
func (o *TransferRunsNoContent) Code() int {
	return 204
}

func (o *TransferRunsNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/runs/transfer][%d] transferRunsNoContent  %+v", 204, o.Payload)
}

func (o *TransferRunsNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/runs/transfer][%d] transferRunsNoContent  %+v", 204, o.Payload)
}

func (o *TransferRunsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *TransferRunsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransferRunsForbidden creates a TransferRunsForbidden with default headers values
func NewTransferRunsForbidden() *TransferRunsForbidden {
	return &TransferRunsForbidden{}
}

/*
TransferRunsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type TransferRunsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this transfer runs forbidden response has a 2xx status code
func (o *TransferRunsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this transfer runs forbidden response has a 3xx status code
func (o *TransferRunsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this transfer runs forbidden response has a 4xx status code
func (o *TransferRunsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this transfer runs forbidden response has a 5xx status code
func (o *TransferRunsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this transfer runs forbidden response a status code equal to that given
func (o *TransferRunsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the transfer runs forbidden response
func (o *TransferRunsForbidden) Code() int {
	return 403
}

func (o *TransferRunsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/runs/transfer][%d] transferRunsForbidden  %+v", 403, o.Payload)
}

func (o *TransferRunsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/runs/transfer][%d] transferRunsForbidden  %+v", 403, o.Payload)
}

func (o *TransferRunsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *TransferRunsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransferRunsNotFound creates a TransferRunsNotFound with default headers values
func NewTransferRunsNotFound() *TransferRunsNotFound {
	return &TransferRunsNotFound{}
}

/*
TransferRunsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type TransferRunsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this transfer runs not found response has a 2xx status code
func (o *TransferRunsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this transfer runs not found response has a 3xx status code
func (o *TransferRunsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this transfer runs not found response has a 4xx status code
func (o *TransferRunsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this transfer runs not found response has a 5xx status code
func (o *TransferRunsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this transfer runs not found response a status code equal to that given
func (o *TransferRunsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the transfer runs not found response
func (o *TransferRunsNotFound) Code() int {
	return 404
}

func (o *TransferRunsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/runs/transfer][%d] transferRunsNotFound  %+v", 404, o.Payload)
}

func (o *TransferRunsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/runs/transfer][%d] transferRunsNotFound  %+v", 404, o.Payload)
}

func (o *TransferRunsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *TransferRunsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransferRunsDefault creates a TransferRunsDefault with default headers values
func NewTransferRunsDefault(code int) *TransferRunsDefault {
	return &TransferRunsDefault{
		_statusCode: code,
	}
}

/*
TransferRunsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type TransferRunsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this transfer runs default response has a 2xx status code
func (o *TransferRunsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this transfer runs default response has a 3xx status code
func (o *TransferRunsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this transfer runs default response has a 4xx status code
func (o *TransferRunsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this transfer runs default response has a 5xx status code
func (o *TransferRunsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this transfer runs default response a status code equal to that given
func (o *TransferRunsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the transfer runs default response
func (o *TransferRunsDefault) Code() int {
	return o._statusCode
}

func (o *TransferRunsDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/runs/transfer][%d] TransferRuns default  %+v", o._statusCode, o.Payload)
}

func (o *TransferRunsDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/runs/transfer][%d] TransferRuns default  %+v", o._statusCode, o.Payload)
}

func (o *TransferRunsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *TransferRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
