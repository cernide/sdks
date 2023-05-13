// Code generated by go-swagger; DO NOT EDIT.

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// GetRunNamespaceReader is a Reader for the GetRunNamespace structure.
type GetRunNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetRunNamespaceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRunNamespaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRunNamespaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetRunNamespaceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRunNamespaceOK creates a GetRunNamespaceOK with default headers values
func NewGetRunNamespaceOK() *GetRunNamespaceOK {
	return &GetRunNamespaceOK{}
}

/*
GetRunNamespaceOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetRunNamespaceOK struct {
	Payload *service_model.V1RunSettings
}

// IsSuccess returns true when this get run namespace o k response has a 2xx status code
func (o *GetRunNamespaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get run namespace o k response has a 3xx status code
func (o *GetRunNamespaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run namespace o k response has a 4xx status code
func (o *GetRunNamespaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get run namespace o k response has a 5xx status code
func (o *GetRunNamespaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get run namespace o k response a status code equal to that given
func (o *GetRunNamespaceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get run namespace o k response
func (o *GetRunNamespaceOK) Code() int {
	return 200
}

func (o *GetRunNamespaceOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/namespace][%d] getRunNamespaceOK  %+v", 200, o.Payload)
}

func (o *GetRunNamespaceOK) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/namespace][%d] getRunNamespaceOK  %+v", 200, o.Payload)
}

func (o *GetRunNamespaceOK) GetPayload() *service_model.V1RunSettings {
	return o.Payload
}

func (o *GetRunNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1RunSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunNamespaceNoContent creates a GetRunNamespaceNoContent with default headers values
func NewGetRunNamespaceNoContent() *GetRunNamespaceNoContent {
	return &GetRunNamespaceNoContent{}
}

/*
GetRunNamespaceNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetRunNamespaceNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get run namespace no content response has a 2xx status code
func (o *GetRunNamespaceNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get run namespace no content response has a 3xx status code
func (o *GetRunNamespaceNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run namespace no content response has a 4xx status code
func (o *GetRunNamespaceNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get run namespace no content response has a 5xx status code
func (o *GetRunNamespaceNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get run namespace no content response a status code equal to that given
func (o *GetRunNamespaceNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get run namespace no content response
func (o *GetRunNamespaceNoContent) Code() int {
	return 204
}

func (o *GetRunNamespaceNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/namespace][%d] getRunNamespaceNoContent  %+v", 204, o.Payload)
}

func (o *GetRunNamespaceNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/namespace][%d] getRunNamespaceNoContent  %+v", 204, o.Payload)
}

func (o *GetRunNamespaceNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunNamespaceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunNamespaceForbidden creates a GetRunNamespaceForbidden with default headers values
func NewGetRunNamespaceForbidden() *GetRunNamespaceForbidden {
	return &GetRunNamespaceForbidden{}
}

/*
GetRunNamespaceForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetRunNamespaceForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get run namespace forbidden response has a 2xx status code
func (o *GetRunNamespaceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get run namespace forbidden response has a 3xx status code
func (o *GetRunNamespaceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run namespace forbidden response has a 4xx status code
func (o *GetRunNamespaceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get run namespace forbidden response has a 5xx status code
func (o *GetRunNamespaceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get run namespace forbidden response a status code equal to that given
func (o *GetRunNamespaceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get run namespace forbidden response
func (o *GetRunNamespaceForbidden) Code() int {
	return 403
}

func (o *GetRunNamespaceForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/namespace][%d] getRunNamespaceForbidden  %+v", 403, o.Payload)
}

func (o *GetRunNamespaceForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/namespace][%d] getRunNamespaceForbidden  %+v", 403, o.Payload)
}

func (o *GetRunNamespaceForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunNamespaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunNamespaceNotFound creates a GetRunNamespaceNotFound with default headers values
func NewGetRunNamespaceNotFound() *GetRunNamespaceNotFound {
	return &GetRunNamespaceNotFound{}
}

/*
GetRunNamespaceNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetRunNamespaceNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get run namespace not found response has a 2xx status code
func (o *GetRunNamespaceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get run namespace not found response has a 3xx status code
func (o *GetRunNamespaceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run namespace not found response has a 4xx status code
func (o *GetRunNamespaceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get run namespace not found response has a 5xx status code
func (o *GetRunNamespaceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get run namespace not found response a status code equal to that given
func (o *GetRunNamespaceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get run namespace not found response
func (o *GetRunNamespaceNotFound) Code() int {
	return 404
}

func (o *GetRunNamespaceNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/namespace][%d] getRunNamespaceNotFound  %+v", 404, o.Payload)
}

func (o *GetRunNamespaceNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/namespace][%d] getRunNamespaceNotFound  %+v", 404, o.Payload)
}

func (o *GetRunNamespaceNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunNamespaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunNamespaceDefault creates a GetRunNamespaceDefault with default headers values
func NewGetRunNamespaceDefault(code int) *GetRunNamespaceDefault {
	return &GetRunNamespaceDefault{
		_statusCode: code,
	}
}

/*
GetRunNamespaceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetRunNamespaceDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this get run namespace default response has a 2xx status code
func (o *GetRunNamespaceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get run namespace default response has a 3xx status code
func (o *GetRunNamespaceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get run namespace default response has a 4xx status code
func (o *GetRunNamespaceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get run namespace default response has a 5xx status code
func (o *GetRunNamespaceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get run namespace default response a status code equal to that given
func (o *GetRunNamespaceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get run namespace default response
func (o *GetRunNamespaceDefault) Code() int {
	return o._statusCode
}

func (o *GetRunNamespaceDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/namespace][%d] GetRunNamespace default  %+v", o._statusCode, o.Payload)
}

func (o *GetRunNamespaceDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/namespace][%d] GetRunNamespace default  %+v", o._statusCode, o.Payload)
}

func (o *GetRunNamespaceDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetRunNamespaceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
