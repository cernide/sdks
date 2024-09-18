// Code generated by go-swagger; DO NOT EDIT.

package searches_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/go/http_client/v1/service_model"
)

// UpdateSearchReader is a Reader for the UpdateSearch structure.
type UpdateSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateSearchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSearchOK creates a UpdateSearchOK with default headers values
func NewUpdateSearchOK() *UpdateSearchOK {
	return &UpdateSearchOK{}
}

/*
UpdateSearchOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateSearchOK struct {
	Payload *service_model.V1Search
}

// IsSuccess returns true when this update search o k response has a 2xx status code
func (o *UpdateSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update search o k response has a 3xx status code
func (o *UpdateSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update search o k response has a 4xx status code
func (o *UpdateSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update search o k response has a 5xx status code
func (o *UpdateSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update search o k response a status code equal to that given
func (o *UpdateSearchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update search o k response
func (o *UpdateSearchOK) Code() int {
	return 200
}

func (o *UpdateSearchOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/searches/{search.uuid}][%d] updateSearchOK  %+v", 200, o.Payload)
}

func (o *UpdateSearchOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/searches/{search.uuid}][%d] updateSearchOK  %+v", 200, o.Payload)
}

func (o *UpdateSearchOK) GetPayload() *service_model.V1Search {
	return o.Payload
}

func (o *UpdateSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Search)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSearchNoContent creates a UpdateSearchNoContent with default headers values
func NewUpdateSearchNoContent() *UpdateSearchNoContent {
	return &UpdateSearchNoContent{}
}

/*
UpdateSearchNoContent describes a response with status code 204, with default header values.

No content.
*/
type UpdateSearchNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this update search no content response has a 2xx status code
func (o *UpdateSearchNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update search no content response has a 3xx status code
func (o *UpdateSearchNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update search no content response has a 4xx status code
func (o *UpdateSearchNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update search no content response has a 5xx status code
func (o *UpdateSearchNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update search no content response a status code equal to that given
func (o *UpdateSearchNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update search no content response
func (o *UpdateSearchNoContent) Code() int {
	return 204
}

func (o *UpdateSearchNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/searches/{search.uuid}][%d] updateSearchNoContent  %+v", 204, o.Payload)
}

func (o *UpdateSearchNoContent) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/searches/{search.uuid}][%d] updateSearchNoContent  %+v", 204, o.Payload)
}

func (o *UpdateSearchNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateSearchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSearchForbidden creates a UpdateSearchForbidden with default headers values
func NewUpdateSearchForbidden() *UpdateSearchForbidden {
	return &UpdateSearchForbidden{}
}

/*
UpdateSearchForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type UpdateSearchForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this update search forbidden response has a 2xx status code
func (o *UpdateSearchForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update search forbidden response has a 3xx status code
func (o *UpdateSearchForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update search forbidden response has a 4xx status code
func (o *UpdateSearchForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update search forbidden response has a 5xx status code
func (o *UpdateSearchForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update search forbidden response a status code equal to that given
func (o *UpdateSearchForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update search forbidden response
func (o *UpdateSearchForbidden) Code() int {
	return 403
}

func (o *UpdateSearchForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/searches/{search.uuid}][%d] updateSearchForbidden  %+v", 403, o.Payload)
}

func (o *UpdateSearchForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/searches/{search.uuid}][%d] updateSearchForbidden  %+v", 403, o.Payload)
}

func (o *UpdateSearchForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSearchNotFound creates a UpdateSearchNotFound with default headers values
func NewUpdateSearchNotFound() *UpdateSearchNotFound {
	return &UpdateSearchNotFound{}
}

/*
UpdateSearchNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type UpdateSearchNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this update search not found response has a 2xx status code
func (o *UpdateSearchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update search not found response has a 3xx status code
func (o *UpdateSearchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update search not found response has a 4xx status code
func (o *UpdateSearchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update search not found response has a 5xx status code
func (o *UpdateSearchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update search not found response a status code equal to that given
func (o *UpdateSearchNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update search not found response
func (o *UpdateSearchNotFound) Code() int {
	return 404
}

func (o *UpdateSearchNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/searches/{search.uuid}][%d] updateSearchNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSearchNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/searches/{search.uuid}][%d] updateSearchNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSearchNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSearchDefault creates a UpdateSearchDefault with default headers values
func NewUpdateSearchDefault(code int) *UpdateSearchDefault {
	return &UpdateSearchDefault{
		_statusCode: code,
	}
}

/*
UpdateSearchDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateSearchDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this update search default response has a 2xx status code
func (o *UpdateSearchDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update search default response has a 3xx status code
func (o *UpdateSearchDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update search default response has a 4xx status code
func (o *UpdateSearchDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update search default response has a 5xx status code
func (o *UpdateSearchDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update search default response a status code equal to that given
func (o *UpdateSearchDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update search default response
func (o *UpdateSearchDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSearchDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/searches/{search.uuid}][%d] UpdateSearch default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSearchDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/searches/{search.uuid}][%d] UpdateSearch default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSearchDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *UpdateSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
