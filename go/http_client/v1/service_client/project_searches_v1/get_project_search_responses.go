// Code generated by go-swagger; DO NOT EDIT.

package project_searches_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// GetProjectSearchReader is a Reader for the GetProjectSearch structure.
type GetProjectSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetProjectSearchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetProjectSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetProjectSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProjectSearchOK creates a GetProjectSearchOK with default headers values
func NewGetProjectSearchOK() *GetProjectSearchOK {
	return &GetProjectSearchOK{}
}

/*
GetProjectSearchOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetProjectSearchOK struct {
	Payload *service_model.V1Search
}

// IsSuccess returns true when this get project search o k response has a 2xx status code
func (o *GetProjectSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get project search o k response has a 3xx status code
func (o *GetProjectSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project search o k response has a 4xx status code
func (o *GetProjectSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project search o k response has a 5xx status code
func (o *GetProjectSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get project search o k response a status code equal to that given
func (o *GetProjectSearchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get project search o k response
func (o *GetProjectSearchOK) Code() int {
	return 200
}

func (o *GetProjectSearchOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/searches/{uuid}][%d] getProjectSearchOK  %+v", 200, o.Payload)
}

func (o *GetProjectSearchOK) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/searches/{uuid}][%d] getProjectSearchOK  %+v", 200, o.Payload)
}

func (o *GetProjectSearchOK) GetPayload() *service_model.V1Search {
	return o.Payload
}

func (o *GetProjectSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Search)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectSearchNoContent creates a GetProjectSearchNoContent with default headers values
func NewGetProjectSearchNoContent() *GetProjectSearchNoContent {
	return &GetProjectSearchNoContent{}
}

/*
GetProjectSearchNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetProjectSearchNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get project search no content response has a 2xx status code
func (o *GetProjectSearchNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get project search no content response has a 3xx status code
func (o *GetProjectSearchNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project search no content response has a 4xx status code
func (o *GetProjectSearchNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project search no content response has a 5xx status code
func (o *GetProjectSearchNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get project search no content response a status code equal to that given
func (o *GetProjectSearchNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get project search no content response
func (o *GetProjectSearchNoContent) Code() int {
	return 204
}

func (o *GetProjectSearchNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/searches/{uuid}][%d] getProjectSearchNoContent  %+v", 204, o.Payload)
}

func (o *GetProjectSearchNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/searches/{uuid}][%d] getProjectSearchNoContent  %+v", 204, o.Payload)
}

func (o *GetProjectSearchNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetProjectSearchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectSearchForbidden creates a GetProjectSearchForbidden with default headers values
func NewGetProjectSearchForbidden() *GetProjectSearchForbidden {
	return &GetProjectSearchForbidden{}
}

/*
GetProjectSearchForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetProjectSearchForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get project search forbidden response has a 2xx status code
func (o *GetProjectSearchForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project search forbidden response has a 3xx status code
func (o *GetProjectSearchForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project search forbidden response has a 4xx status code
func (o *GetProjectSearchForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project search forbidden response has a 5xx status code
func (o *GetProjectSearchForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get project search forbidden response a status code equal to that given
func (o *GetProjectSearchForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get project search forbidden response
func (o *GetProjectSearchForbidden) Code() int {
	return 403
}

func (o *GetProjectSearchForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/searches/{uuid}][%d] getProjectSearchForbidden  %+v", 403, o.Payload)
}

func (o *GetProjectSearchForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/searches/{uuid}][%d] getProjectSearchForbidden  %+v", 403, o.Payload)
}

func (o *GetProjectSearchForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetProjectSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectSearchNotFound creates a GetProjectSearchNotFound with default headers values
func NewGetProjectSearchNotFound() *GetProjectSearchNotFound {
	return &GetProjectSearchNotFound{}
}

/*
GetProjectSearchNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetProjectSearchNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get project search not found response has a 2xx status code
func (o *GetProjectSearchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project search not found response has a 3xx status code
func (o *GetProjectSearchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project search not found response has a 4xx status code
func (o *GetProjectSearchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project search not found response has a 5xx status code
func (o *GetProjectSearchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get project search not found response a status code equal to that given
func (o *GetProjectSearchNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get project search not found response
func (o *GetProjectSearchNotFound) Code() int {
	return 404
}

func (o *GetProjectSearchNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/searches/{uuid}][%d] getProjectSearchNotFound  %+v", 404, o.Payload)
}

func (o *GetProjectSearchNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/searches/{uuid}][%d] getProjectSearchNotFound  %+v", 404, o.Payload)
}

func (o *GetProjectSearchNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetProjectSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectSearchDefault creates a GetProjectSearchDefault with default headers values
func NewGetProjectSearchDefault(code int) *GetProjectSearchDefault {
	return &GetProjectSearchDefault{
		_statusCode: code,
	}
}

/*
GetProjectSearchDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetProjectSearchDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this get project search default response has a 2xx status code
func (o *GetProjectSearchDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get project search default response has a 3xx status code
func (o *GetProjectSearchDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get project search default response has a 4xx status code
func (o *GetProjectSearchDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get project search default response has a 5xx status code
func (o *GetProjectSearchDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get project search default response a status code equal to that given
func (o *GetProjectSearchDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get project search default response
func (o *GetProjectSearchDefault) Code() int {
	return o._statusCode
}

func (o *GetProjectSearchDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/searches/{uuid}][%d] GetProjectSearch default  %+v", o._statusCode, o.Payload)
}

func (o *GetProjectSearchDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/searches/{uuid}][%d] GetProjectSearch default  %+v", o._statusCode, o.Payload)
}

func (o *GetProjectSearchDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetProjectSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
