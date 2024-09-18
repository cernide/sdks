// Code generated by go-swagger; DO NOT EDIT.

package teams_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/go/http_client/v1/service_model"
)

// TransferTeamRunsReader is a Reader for the TransferTeamRuns structure.
type TransferTeamRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TransferTeamRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTransferTeamRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewTransferTeamRunsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewTransferTeamRunsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTransferTeamRunsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewTransferTeamRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTransferTeamRunsOK creates a TransferTeamRunsOK with default headers values
func NewTransferTeamRunsOK() *TransferTeamRunsOK {
	return &TransferTeamRunsOK{}
}

/*
TransferTeamRunsOK describes a response with status code 200, with default header values.

A successful response.
*/
type TransferTeamRunsOK struct {
}

// IsSuccess returns true when this transfer team runs o k response has a 2xx status code
func (o *TransferTeamRunsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this transfer team runs o k response has a 3xx status code
func (o *TransferTeamRunsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this transfer team runs o k response has a 4xx status code
func (o *TransferTeamRunsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this transfer team runs o k response has a 5xx status code
func (o *TransferTeamRunsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this transfer team runs o k response a status code equal to that given
func (o *TransferTeamRunsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the transfer team runs o k response
func (o *TransferTeamRunsOK) Code() int {
	return 200
}

func (o *TransferTeamRunsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/transfer][%d] transferTeamRunsOK ", 200)
}

func (o *TransferTeamRunsOK) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/transfer][%d] transferTeamRunsOK ", 200)
}

func (o *TransferTeamRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTransferTeamRunsNoContent creates a TransferTeamRunsNoContent with default headers values
func NewTransferTeamRunsNoContent() *TransferTeamRunsNoContent {
	return &TransferTeamRunsNoContent{}
}

/*
TransferTeamRunsNoContent describes a response with status code 204, with default header values.

No content.
*/
type TransferTeamRunsNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this transfer team runs no content response has a 2xx status code
func (o *TransferTeamRunsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this transfer team runs no content response has a 3xx status code
func (o *TransferTeamRunsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this transfer team runs no content response has a 4xx status code
func (o *TransferTeamRunsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this transfer team runs no content response has a 5xx status code
func (o *TransferTeamRunsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this transfer team runs no content response a status code equal to that given
func (o *TransferTeamRunsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the transfer team runs no content response
func (o *TransferTeamRunsNoContent) Code() int {
	return 204
}

func (o *TransferTeamRunsNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/transfer][%d] transferTeamRunsNoContent  %+v", 204, o.Payload)
}

func (o *TransferTeamRunsNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/transfer][%d] transferTeamRunsNoContent  %+v", 204, o.Payload)
}

func (o *TransferTeamRunsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *TransferTeamRunsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransferTeamRunsForbidden creates a TransferTeamRunsForbidden with default headers values
func NewTransferTeamRunsForbidden() *TransferTeamRunsForbidden {
	return &TransferTeamRunsForbidden{}
}

/*
TransferTeamRunsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type TransferTeamRunsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this transfer team runs forbidden response has a 2xx status code
func (o *TransferTeamRunsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this transfer team runs forbidden response has a 3xx status code
func (o *TransferTeamRunsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this transfer team runs forbidden response has a 4xx status code
func (o *TransferTeamRunsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this transfer team runs forbidden response has a 5xx status code
func (o *TransferTeamRunsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this transfer team runs forbidden response a status code equal to that given
func (o *TransferTeamRunsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the transfer team runs forbidden response
func (o *TransferTeamRunsForbidden) Code() int {
	return 403
}

func (o *TransferTeamRunsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/transfer][%d] transferTeamRunsForbidden  %+v", 403, o.Payload)
}

func (o *TransferTeamRunsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/transfer][%d] transferTeamRunsForbidden  %+v", 403, o.Payload)
}

func (o *TransferTeamRunsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *TransferTeamRunsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransferTeamRunsNotFound creates a TransferTeamRunsNotFound with default headers values
func NewTransferTeamRunsNotFound() *TransferTeamRunsNotFound {
	return &TransferTeamRunsNotFound{}
}

/*
TransferTeamRunsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type TransferTeamRunsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this transfer team runs not found response has a 2xx status code
func (o *TransferTeamRunsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this transfer team runs not found response has a 3xx status code
func (o *TransferTeamRunsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this transfer team runs not found response has a 4xx status code
func (o *TransferTeamRunsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this transfer team runs not found response has a 5xx status code
func (o *TransferTeamRunsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this transfer team runs not found response a status code equal to that given
func (o *TransferTeamRunsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the transfer team runs not found response
func (o *TransferTeamRunsNotFound) Code() int {
	return 404
}

func (o *TransferTeamRunsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/transfer][%d] transferTeamRunsNotFound  %+v", 404, o.Payload)
}

func (o *TransferTeamRunsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/transfer][%d] transferTeamRunsNotFound  %+v", 404, o.Payload)
}

func (o *TransferTeamRunsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *TransferTeamRunsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransferTeamRunsDefault creates a TransferTeamRunsDefault with default headers values
func NewTransferTeamRunsDefault(code int) *TransferTeamRunsDefault {
	return &TransferTeamRunsDefault{
		_statusCode: code,
	}
}

/*
TransferTeamRunsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type TransferTeamRunsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this transfer team runs default response has a 2xx status code
func (o *TransferTeamRunsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this transfer team runs default response has a 3xx status code
func (o *TransferTeamRunsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this transfer team runs default response has a 4xx status code
func (o *TransferTeamRunsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this transfer team runs default response has a 5xx status code
func (o *TransferTeamRunsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this transfer team runs default response a status code equal to that given
func (o *TransferTeamRunsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the transfer team runs default response
func (o *TransferTeamRunsDefault) Code() int {
	return o._statusCode
}

func (o *TransferTeamRunsDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/transfer][%d] TransferTeamRuns default  %+v", o._statusCode, o.Payload)
}

func (o *TransferTeamRunsDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/teams/{name}/runs/transfer][%d] TransferTeamRuns default  %+v", o._statusCode, o.Payload)
}

func (o *TransferTeamRunsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *TransferTeamRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
