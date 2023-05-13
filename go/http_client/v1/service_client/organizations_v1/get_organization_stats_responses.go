// Code generated by go-swagger; DO NOT EDIT.

package organizations_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// GetOrganizationStatsReader is a Reader for the GetOrganizationStats structure.
type GetOrganizationStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetOrganizationStatsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetOrganizationStatsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationStatsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetOrganizationStatsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOrganizationStatsOK creates a GetOrganizationStatsOK with default headers values
func NewGetOrganizationStatsOK() *GetOrganizationStatsOK {
	return &GetOrganizationStatsOK{}
}

/*
GetOrganizationStatsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetOrganizationStatsOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get organization stats o k response has a 2xx status code
func (o *GetOrganizationStatsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization stats o k response has a 3xx status code
func (o *GetOrganizationStatsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization stats o k response has a 4xx status code
func (o *GetOrganizationStatsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization stats o k response has a 5xx status code
func (o *GetOrganizationStatsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization stats o k response a status code equal to that given
func (o *GetOrganizationStatsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization stats o k response
func (o *GetOrganizationStatsOK) Code() int {
	return 200
}

func (o *GetOrganizationStatsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/stats][%d] getOrganizationStatsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationStatsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/stats][%d] getOrganizationStatsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationStatsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationStatsNoContent creates a GetOrganizationStatsNoContent with default headers values
func NewGetOrganizationStatsNoContent() *GetOrganizationStatsNoContent {
	return &GetOrganizationStatsNoContent{}
}

/*
GetOrganizationStatsNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetOrganizationStatsNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get organization stats no content response has a 2xx status code
func (o *GetOrganizationStatsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization stats no content response has a 3xx status code
func (o *GetOrganizationStatsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization stats no content response has a 4xx status code
func (o *GetOrganizationStatsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization stats no content response has a 5xx status code
func (o *GetOrganizationStatsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization stats no content response a status code equal to that given
func (o *GetOrganizationStatsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get organization stats no content response
func (o *GetOrganizationStatsNoContent) Code() int {
	return 204
}

func (o *GetOrganizationStatsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/stats][%d] getOrganizationStatsNoContent  %+v", 204, o.Payload)
}

func (o *GetOrganizationStatsNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/stats][%d] getOrganizationStatsNoContent  %+v", 204, o.Payload)
}

func (o *GetOrganizationStatsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationStatsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationStatsForbidden creates a GetOrganizationStatsForbidden with default headers values
func NewGetOrganizationStatsForbidden() *GetOrganizationStatsForbidden {
	return &GetOrganizationStatsForbidden{}
}

/*
GetOrganizationStatsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetOrganizationStatsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get organization stats forbidden response has a 2xx status code
func (o *GetOrganizationStatsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization stats forbidden response has a 3xx status code
func (o *GetOrganizationStatsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization stats forbidden response has a 4xx status code
func (o *GetOrganizationStatsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization stats forbidden response has a 5xx status code
func (o *GetOrganizationStatsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization stats forbidden response a status code equal to that given
func (o *GetOrganizationStatsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get organization stats forbidden response
func (o *GetOrganizationStatsForbidden) Code() int {
	return 403
}

func (o *GetOrganizationStatsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/stats][%d] getOrganizationStatsForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationStatsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/stats][%d] getOrganizationStatsForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationStatsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationStatsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationStatsNotFound creates a GetOrganizationStatsNotFound with default headers values
func NewGetOrganizationStatsNotFound() *GetOrganizationStatsNotFound {
	return &GetOrganizationStatsNotFound{}
}

/*
GetOrganizationStatsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetOrganizationStatsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get organization stats not found response has a 2xx status code
func (o *GetOrganizationStatsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization stats not found response has a 3xx status code
func (o *GetOrganizationStatsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization stats not found response has a 4xx status code
func (o *GetOrganizationStatsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization stats not found response has a 5xx status code
func (o *GetOrganizationStatsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization stats not found response a status code equal to that given
func (o *GetOrganizationStatsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get organization stats not found response
func (o *GetOrganizationStatsNotFound) Code() int {
	return 404
}

func (o *GetOrganizationStatsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/stats][%d] getOrganizationStatsNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationStatsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/stats][%d] getOrganizationStatsNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationStatsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationStatsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationStatsDefault creates a GetOrganizationStatsDefault with default headers values
func NewGetOrganizationStatsDefault(code int) *GetOrganizationStatsDefault {
	return &GetOrganizationStatsDefault{
		_statusCode: code,
	}
}

/*
GetOrganizationStatsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetOrganizationStatsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this get organization stats default response has a 2xx status code
func (o *GetOrganizationStatsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get organization stats default response has a 3xx status code
func (o *GetOrganizationStatsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get organization stats default response has a 4xx status code
func (o *GetOrganizationStatsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get organization stats default response has a 5xx status code
func (o *GetOrganizationStatsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get organization stats default response a status code equal to that given
func (o *GetOrganizationStatsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get organization stats default response
func (o *GetOrganizationStatsDefault) Code() int {
	return o._statusCode
}

func (o *GetOrganizationStatsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/stats][%d] GetOrganizationStats default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrganizationStatsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/stats][%d] GetOrganizationStats default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrganizationStatsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetOrganizationStatsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
