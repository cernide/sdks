// Code generated by go-swagger; DO NOT EDIT.

package searches_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// CreateSearchReader is a Reader for the CreateSearch structure.
type CreateSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCreateSearchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSearchOK creates a CreateSearchOK with default headers values
func NewCreateSearchOK() *CreateSearchOK {
	return &CreateSearchOK{}
}

/*
CreateSearchOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateSearchOK struct {
	Payload *service_model.V1Search
}

// IsSuccess returns true when this create search o k response has a 2xx status code
func (o *CreateSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create search o k response has a 3xx status code
func (o *CreateSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create search o k response has a 4xx status code
func (o *CreateSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create search o k response has a 5xx status code
func (o *CreateSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create search o k response a status code equal to that given
func (o *CreateSearchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create search o k response
func (o *CreateSearchOK) Code() int {
	return 200
}

func (o *CreateSearchOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/searches][%d] createSearchOK  %+v", 200, o.Payload)
}

func (o *CreateSearchOK) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/searches][%d] createSearchOK  %+v", 200, o.Payload)
}

func (o *CreateSearchOK) GetPayload() *service_model.V1Search {
	return o.Payload
}

func (o *CreateSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Search)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSearchNoContent creates a CreateSearchNoContent with default headers values
func NewCreateSearchNoContent() *CreateSearchNoContent {
	return &CreateSearchNoContent{}
}

/*
CreateSearchNoContent describes a response with status code 204, with default header values.

No content.
*/
type CreateSearchNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this create search no content response has a 2xx status code
func (o *CreateSearchNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create search no content response has a 3xx status code
func (o *CreateSearchNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create search no content response has a 4xx status code
func (o *CreateSearchNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this create search no content response has a 5xx status code
func (o *CreateSearchNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this create search no content response a status code equal to that given
func (o *CreateSearchNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the create search no content response
func (o *CreateSearchNoContent) Code() int {
	return 204
}

func (o *CreateSearchNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/searches][%d] createSearchNoContent  %+v", 204, o.Payload)
}

func (o *CreateSearchNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/searches][%d] createSearchNoContent  %+v", 204, o.Payload)
}

func (o *CreateSearchNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateSearchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSearchForbidden creates a CreateSearchForbidden with default headers values
func NewCreateSearchForbidden() *CreateSearchForbidden {
	return &CreateSearchForbidden{}
}

/*
CreateSearchForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type CreateSearchForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this create search forbidden response has a 2xx status code
func (o *CreateSearchForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create search forbidden response has a 3xx status code
func (o *CreateSearchForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create search forbidden response has a 4xx status code
func (o *CreateSearchForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create search forbidden response has a 5xx status code
func (o *CreateSearchForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create search forbidden response a status code equal to that given
func (o *CreateSearchForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create search forbidden response
func (o *CreateSearchForbidden) Code() int {
	return 403
}

func (o *CreateSearchForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/searches][%d] createSearchForbidden  %+v", 403, o.Payload)
}

func (o *CreateSearchForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/searches][%d] createSearchForbidden  %+v", 403, o.Payload)
}

func (o *CreateSearchForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSearchNotFound creates a CreateSearchNotFound with default headers values
func NewCreateSearchNotFound() *CreateSearchNotFound {
	return &CreateSearchNotFound{}
}

/*
CreateSearchNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type CreateSearchNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this create search not found response has a 2xx status code
func (o *CreateSearchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create search not found response has a 3xx status code
func (o *CreateSearchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create search not found response has a 4xx status code
func (o *CreateSearchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create search not found response has a 5xx status code
func (o *CreateSearchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create search not found response a status code equal to that given
func (o *CreateSearchNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create search not found response
func (o *CreateSearchNotFound) Code() int {
	return 404
}

func (o *CreateSearchNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/searches][%d] createSearchNotFound  %+v", 404, o.Payload)
}

func (o *CreateSearchNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/searches][%d] createSearchNotFound  %+v", 404, o.Payload)
}

func (o *CreateSearchNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSearchDefault creates a CreateSearchDefault with default headers values
func NewCreateSearchDefault(code int) *CreateSearchDefault {
	return &CreateSearchDefault{
		_statusCode: code,
	}
}

/*
CreateSearchDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateSearchDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this create search default response has a 2xx status code
func (o *CreateSearchDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create search default response has a 3xx status code
func (o *CreateSearchDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create search default response has a 4xx status code
func (o *CreateSearchDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create search default response has a 5xx status code
func (o *CreateSearchDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create search default response a status code equal to that given
func (o *CreateSearchDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create search default response
func (o *CreateSearchDefault) Code() int {
	return o._statusCode
}

func (o *CreateSearchDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/searches][%d] CreateSearch default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSearchDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/searches][%d] CreateSearch default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSearchDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *CreateSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
