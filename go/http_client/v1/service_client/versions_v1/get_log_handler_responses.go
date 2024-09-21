// Code generated by go-swagger; DO NOT EDIT.

package versions_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// GetLogHandlerReader is a Reader for the GetLogHandler structure.
type GetLogHandlerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogHandlerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogHandlerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetLogHandlerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetLogHandlerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLogHandlerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetLogHandlerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLogHandlerOK creates a GetLogHandlerOK with default headers values
func NewGetLogHandlerOK() *GetLogHandlerOK {
	return &GetLogHandlerOK{}
}

/*
GetLogHandlerOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetLogHandlerOK struct {
	Payload *service_model.V1LogHandler
}

// IsSuccess returns true when this get log handler o k response has a 2xx status code
func (o *GetLogHandlerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get log handler o k response has a 3xx status code
func (o *GetLogHandlerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get log handler o k response has a 4xx status code
func (o *GetLogHandlerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get log handler o k response has a 5xx status code
func (o *GetLogHandlerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get log handler o k response a status code equal to that given
func (o *GetLogHandlerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get log handler o k response
func (o *GetLogHandlerOK) Code() int {
	return 200
}

func (o *GetLogHandlerOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/log_handler][%d] getLogHandlerOK  %+v", 200, o.Payload)
}

func (o *GetLogHandlerOK) String() string {
	return fmt.Sprintf("[GET /api/v1/log_handler][%d] getLogHandlerOK  %+v", 200, o.Payload)
}

func (o *GetLogHandlerOK) GetPayload() *service_model.V1LogHandler {
	return o.Payload
}

func (o *GetLogHandlerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1LogHandler)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogHandlerNoContent creates a GetLogHandlerNoContent with default headers values
func NewGetLogHandlerNoContent() *GetLogHandlerNoContent {
	return &GetLogHandlerNoContent{}
}

/*
GetLogHandlerNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetLogHandlerNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get log handler no content response has a 2xx status code
func (o *GetLogHandlerNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get log handler no content response has a 3xx status code
func (o *GetLogHandlerNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get log handler no content response has a 4xx status code
func (o *GetLogHandlerNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get log handler no content response has a 5xx status code
func (o *GetLogHandlerNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get log handler no content response a status code equal to that given
func (o *GetLogHandlerNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get log handler no content response
func (o *GetLogHandlerNoContent) Code() int {
	return 204
}

func (o *GetLogHandlerNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/log_handler][%d] getLogHandlerNoContent  %+v", 204, o.Payload)
}

func (o *GetLogHandlerNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/log_handler][%d] getLogHandlerNoContent  %+v", 204, o.Payload)
}

func (o *GetLogHandlerNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetLogHandlerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogHandlerForbidden creates a GetLogHandlerForbidden with default headers values
func NewGetLogHandlerForbidden() *GetLogHandlerForbidden {
	return &GetLogHandlerForbidden{}
}

/*
GetLogHandlerForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetLogHandlerForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get log handler forbidden response has a 2xx status code
func (o *GetLogHandlerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get log handler forbidden response has a 3xx status code
func (o *GetLogHandlerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get log handler forbidden response has a 4xx status code
func (o *GetLogHandlerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get log handler forbidden response has a 5xx status code
func (o *GetLogHandlerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get log handler forbidden response a status code equal to that given
func (o *GetLogHandlerForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get log handler forbidden response
func (o *GetLogHandlerForbidden) Code() int {
	return 403
}

func (o *GetLogHandlerForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/log_handler][%d] getLogHandlerForbidden  %+v", 403, o.Payload)
}

func (o *GetLogHandlerForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/log_handler][%d] getLogHandlerForbidden  %+v", 403, o.Payload)
}

func (o *GetLogHandlerForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetLogHandlerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogHandlerNotFound creates a GetLogHandlerNotFound with default headers values
func NewGetLogHandlerNotFound() *GetLogHandlerNotFound {
	return &GetLogHandlerNotFound{}
}

/*
GetLogHandlerNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetLogHandlerNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get log handler not found response has a 2xx status code
func (o *GetLogHandlerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get log handler not found response has a 3xx status code
func (o *GetLogHandlerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get log handler not found response has a 4xx status code
func (o *GetLogHandlerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get log handler not found response has a 5xx status code
func (o *GetLogHandlerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get log handler not found response a status code equal to that given
func (o *GetLogHandlerNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get log handler not found response
func (o *GetLogHandlerNotFound) Code() int {
	return 404
}

func (o *GetLogHandlerNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/log_handler][%d] getLogHandlerNotFound  %+v", 404, o.Payload)
}

func (o *GetLogHandlerNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/log_handler][%d] getLogHandlerNotFound  %+v", 404, o.Payload)
}

func (o *GetLogHandlerNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetLogHandlerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogHandlerDefault creates a GetLogHandlerDefault with default headers values
func NewGetLogHandlerDefault(code int) *GetLogHandlerDefault {
	return &GetLogHandlerDefault{
		_statusCode: code,
	}
}

/*
GetLogHandlerDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetLogHandlerDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this get log handler default response has a 2xx status code
func (o *GetLogHandlerDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get log handler default response has a 3xx status code
func (o *GetLogHandlerDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get log handler default response has a 4xx status code
func (o *GetLogHandlerDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get log handler default response has a 5xx status code
func (o *GetLogHandlerDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get log handler default response a status code equal to that given
func (o *GetLogHandlerDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get log handler default response
func (o *GetLogHandlerDefault) Code() int {
	return o._statusCode
}

func (o *GetLogHandlerDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/log_handler][%d] GetLogHandler default  %+v", o._statusCode, o.Payload)
}

func (o *GetLogHandlerDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/log_handler][%d] GetLogHandler default  %+v", o._statusCode, o.Payload)
}

func (o *GetLogHandlerDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetLogHandlerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
