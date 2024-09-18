// Code generated by go-swagger; DO NOT EDIT.

package projects_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/go/http_client/v1/service_model"
)

// GetVersionReader is a Reader for the GetVersion structure.
type GetVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetVersionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVersionOK creates a GetVersionOK with default headers values
func NewGetVersionOK() *GetVersionOK {
	return &GetVersionOK{}
}

/*
GetVersionOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetVersionOK struct {
	Payload *service_model.V1ProjectVersion
}

// IsSuccess returns true when this get version o k response has a 2xx status code
func (o *GetVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get version o k response has a 3xx status code
func (o *GetVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get version o k response has a 4xx status code
func (o *GetVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get version o k response has a 5xx status code
func (o *GetVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get version o k response a status code equal to that given
func (o *GetVersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get version o k response
func (o *GetVersionOK) Code() int {
	return 200
}

func (o *GetVersionOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/versions/{kind}/{name}][%d] getVersionOK  %+v", 200, o.Payload)
}

func (o *GetVersionOK) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/versions/{kind}/{name}][%d] getVersionOK  %+v", 200, o.Payload)
}

func (o *GetVersionOK) GetPayload() *service_model.V1ProjectVersion {
	return o.Payload
}

func (o *GetVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ProjectVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVersionNoContent creates a GetVersionNoContent with default headers values
func NewGetVersionNoContent() *GetVersionNoContent {
	return &GetVersionNoContent{}
}

/*
GetVersionNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetVersionNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get version no content response has a 2xx status code
func (o *GetVersionNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get version no content response has a 3xx status code
func (o *GetVersionNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get version no content response has a 4xx status code
func (o *GetVersionNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get version no content response has a 5xx status code
func (o *GetVersionNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get version no content response a status code equal to that given
func (o *GetVersionNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get version no content response
func (o *GetVersionNoContent) Code() int {
	return 204
}

func (o *GetVersionNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/versions/{kind}/{name}][%d] getVersionNoContent  %+v", 204, o.Payload)
}

func (o *GetVersionNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/versions/{kind}/{name}][%d] getVersionNoContent  %+v", 204, o.Payload)
}

func (o *GetVersionNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetVersionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVersionForbidden creates a GetVersionForbidden with default headers values
func NewGetVersionForbidden() *GetVersionForbidden {
	return &GetVersionForbidden{}
}

/*
GetVersionForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetVersionForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get version forbidden response has a 2xx status code
func (o *GetVersionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get version forbidden response has a 3xx status code
func (o *GetVersionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get version forbidden response has a 4xx status code
func (o *GetVersionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get version forbidden response has a 5xx status code
func (o *GetVersionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get version forbidden response a status code equal to that given
func (o *GetVersionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get version forbidden response
func (o *GetVersionForbidden) Code() int {
	return 403
}

func (o *GetVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/versions/{kind}/{name}][%d] getVersionForbidden  %+v", 403, o.Payload)
}

func (o *GetVersionForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/versions/{kind}/{name}][%d] getVersionForbidden  %+v", 403, o.Payload)
}

func (o *GetVersionForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVersionNotFound creates a GetVersionNotFound with default headers values
func NewGetVersionNotFound() *GetVersionNotFound {
	return &GetVersionNotFound{}
}

/*
GetVersionNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetVersionNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get version not found response has a 2xx status code
func (o *GetVersionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get version not found response has a 3xx status code
func (o *GetVersionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get version not found response has a 4xx status code
func (o *GetVersionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get version not found response has a 5xx status code
func (o *GetVersionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get version not found response a status code equal to that given
func (o *GetVersionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get version not found response
func (o *GetVersionNotFound) Code() int {
	return 404
}

func (o *GetVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/versions/{kind}/{name}][%d] getVersionNotFound  %+v", 404, o.Payload)
}

func (o *GetVersionNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/versions/{kind}/{name}][%d] getVersionNotFound  %+v", 404, o.Payload)
}

func (o *GetVersionNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVersionDefault creates a GetVersionDefault with default headers values
func NewGetVersionDefault(code int) *GetVersionDefault {
	return &GetVersionDefault{
		_statusCode: code,
	}
}

/*
GetVersionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetVersionDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this get version default response has a 2xx status code
func (o *GetVersionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get version default response has a 3xx status code
func (o *GetVersionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get version default response has a 4xx status code
func (o *GetVersionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get version default response has a 5xx status code
func (o *GetVersionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get version default response a status code equal to that given
func (o *GetVersionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get version default response
func (o *GetVersionDefault) Code() int {
	return o._statusCode
}

func (o *GetVersionDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/versions/{kind}/{name}][%d] GetVersion default  %+v", o._statusCode, o.Payload)
}

func (o *GetVersionDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/versions/{kind}/{name}][%d] GetVersion default  %+v", o._statusCode, o.Payload)
}

func (o *GetVersionDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
