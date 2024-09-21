// Code generated by go-swagger; DO NOT EDIT.

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// GetRunArtifactsLineageReader is a Reader for the GetRunArtifactsLineage structure.
type GetRunArtifactsLineageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunArtifactsLineageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunArtifactsLineageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetRunArtifactsLineageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRunArtifactsLineageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRunArtifactsLineageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetRunArtifactsLineageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRunArtifactsLineageOK creates a GetRunArtifactsLineageOK with default headers values
func NewGetRunArtifactsLineageOK() *GetRunArtifactsLineageOK {
	return &GetRunArtifactsLineageOK{}
}

/*
GetRunArtifactsLineageOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetRunArtifactsLineageOK struct {
	Payload *service_model.V1ListRunArtifactsResponse
}

// IsSuccess returns true when this get run artifacts lineage o k response has a 2xx status code
func (o *GetRunArtifactsLineageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get run artifacts lineage o k response has a 3xx status code
func (o *GetRunArtifactsLineageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run artifacts lineage o k response has a 4xx status code
func (o *GetRunArtifactsLineageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get run artifacts lineage o k response has a 5xx status code
func (o *GetRunArtifactsLineageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get run artifacts lineage o k response a status code equal to that given
func (o *GetRunArtifactsLineageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get run artifacts lineage o k response
func (o *GetRunArtifactsLineageOK) Code() int {
	return 200
}

func (o *GetRunArtifactsLineageOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/artifacts][%d] getRunArtifactsLineageOK  %+v", 200, o.Payload)
}

func (o *GetRunArtifactsLineageOK) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/artifacts][%d] getRunArtifactsLineageOK  %+v", 200, o.Payload)
}

func (o *GetRunArtifactsLineageOK) GetPayload() *service_model.V1ListRunArtifactsResponse {
	return o.Payload
}

func (o *GetRunArtifactsLineageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListRunArtifactsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunArtifactsLineageNoContent creates a GetRunArtifactsLineageNoContent with default headers values
func NewGetRunArtifactsLineageNoContent() *GetRunArtifactsLineageNoContent {
	return &GetRunArtifactsLineageNoContent{}
}

/*
GetRunArtifactsLineageNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetRunArtifactsLineageNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get run artifacts lineage no content response has a 2xx status code
func (o *GetRunArtifactsLineageNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get run artifacts lineage no content response has a 3xx status code
func (o *GetRunArtifactsLineageNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run artifacts lineage no content response has a 4xx status code
func (o *GetRunArtifactsLineageNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get run artifacts lineage no content response has a 5xx status code
func (o *GetRunArtifactsLineageNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get run artifacts lineage no content response a status code equal to that given
func (o *GetRunArtifactsLineageNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get run artifacts lineage no content response
func (o *GetRunArtifactsLineageNoContent) Code() int {
	return 204
}

func (o *GetRunArtifactsLineageNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/artifacts][%d] getRunArtifactsLineageNoContent  %+v", 204, o.Payload)
}

func (o *GetRunArtifactsLineageNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/artifacts][%d] getRunArtifactsLineageNoContent  %+v", 204, o.Payload)
}

func (o *GetRunArtifactsLineageNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunArtifactsLineageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunArtifactsLineageForbidden creates a GetRunArtifactsLineageForbidden with default headers values
func NewGetRunArtifactsLineageForbidden() *GetRunArtifactsLineageForbidden {
	return &GetRunArtifactsLineageForbidden{}
}

/*
GetRunArtifactsLineageForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetRunArtifactsLineageForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get run artifacts lineage forbidden response has a 2xx status code
func (o *GetRunArtifactsLineageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get run artifacts lineage forbidden response has a 3xx status code
func (o *GetRunArtifactsLineageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run artifacts lineage forbidden response has a 4xx status code
func (o *GetRunArtifactsLineageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get run artifacts lineage forbidden response has a 5xx status code
func (o *GetRunArtifactsLineageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get run artifacts lineage forbidden response a status code equal to that given
func (o *GetRunArtifactsLineageForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get run artifacts lineage forbidden response
func (o *GetRunArtifactsLineageForbidden) Code() int {
	return 403
}

func (o *GetRunArtifactsLineageForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/artifacts][%d] getRunArtifactsLineageForbidden  %+v", 403, o.Payload)
}

func (o *GetRunArtifactsLineageForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/artifacts][%d] getRunArtifactsLineageForbidden  %+v", 403, o.Payload)
}

func (o *GetRunArtifactsLineageForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunArtifactsLineageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunArtifactsLineageNotFound creates a GetRunArtifactsLineageNotFound with default headers values
func NewGetRunArtifactsLineageNotFound() *GetRunArtifactsLineageNotFound {
	return &GetRunArtifactsLineageNotFound{}
}

/*
GetRunArtifactsLineageNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetRunArtifactsLineageNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get run artifacts lineage not found response has a 2xx status code
func (o *GetRunArtifactsLineageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get run artifacts lineage not found response has a 3xx status code
func (o *GetRunArtifactsLineageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run artifacts lineage not found response has a 4xx status code
func (o *GetRunArtifactsLineageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get run artifacts lineage not found response has a 5xx status code
func (o *GetRunArtifactsLineageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get run artifacts lineage not found response a status code equal to that given
func (o *GetRunArtifactsLineageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get run artifacts lineage not found response
func (o *GetRunArtifactsLineageNotFound) Code() int {
	return 404
}

func (o *GetRunArtifactsLineageNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/artifacts][%d] getRunArtifactsLineageNotFound  %+v", 404, o.Payload)
}

func (o *GetRunArtifactsLineageNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/artifacts][%d] getRunArtifactsLineageNotFound  %+v", 404, o.Payload)
}

func (o *GetRunArtifactsLineageNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunArtifactsLineageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunArtifactsLineageDefault creates a GetRunArtifactsLineageDefault with default headers values
func NewGetRunArtifactsLineageDefault(code int) *GetRunArtifactsLineageDefault {
	return &GetRunArtifactsLineageDefault{
		_statusCode: code,
	}
}

/*
GetRunArtifactsLineageDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetRunArtifactsLineageDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this get run artifacts lineage default response has a 2xx status code
func (o *GetRunArtifactsLineageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get run artifacts lineage default response has a 3xx status code
func (o *GetRunArtifactsLineageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get run artifacts lineage default response has a 4xx status code
func (o *GetRunArtifactsLineageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get run artifacts lineage default response has a 5xx status code
func (o *GetRunArtifactsLineageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get run artifacts lineage default response a status code equal to that given
func (o *GetRunArtifactsLineageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get run artifacts lineage default response
func (o *GetRunArtifactsLineageDefault) Code() int {
	return o._statusCode
}

func (o *GetRunArtifactsLineageDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/artifacts][%d] GetRunArtifactsLineage default  %+v", o._statusCode, o.Payload)
}

func (o *GetRunArtifactsLineageDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/artifacts][%d] GetRunArtifactsLineage default  %+v", o._statusCode, o.Payload)
}

func (o *GetRunArtifactsLineageDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetRunArtifactsLineageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
