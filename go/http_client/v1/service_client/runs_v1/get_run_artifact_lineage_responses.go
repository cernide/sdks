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

// GetRunArtifactLineageReader is a Reader for the GetRunArtifactLineage structure.
type GetRunArtifactLineageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunArtifactLineageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunArtifactLineageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetRunArtifactLineageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRunArtifactLineageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRunArtifactLineageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetRunArtifactLineageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRunArtifactLineageOK creates a GetRunArtifactLineageOK with default headers values
func NewGetRunArtifactLineageOK() *GetRunArtifactLineageOK {
	return &GetRunArtifactLineageOK{}
}

/*
GetRunArtifactLineageOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetRunArtifactLineageOK struct {
	Payload *service_model.V1RunArtifact
}

// IsSuccess returns true when this get run artifact lineage o k response has a 2xx status code
func (o *GetRunArtifactLineageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get run artifact lineage o k response has a 3xx status code
func (o *GetRunArtifactLineageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run artifact lineage o k response has a 4xx status code
func (o *GetRunArtifactLineageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get run artifact lineage o k response has a 5xx status code
func (o *GetRunArtifactLineageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get run artifact lineage o k response a status code equal to that given
func (o *GetRunArtifactLineageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get run artifact lineage o k response
func (o *GetRunArtifactLineageOK) Code() int {
	return 200
}

func (o *GetRunArtifactLineageOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/lineage/artifacts/{name}][%d] getRunArtifactLineageOK  %+v", 200, o.Payload)
}

func (o *GetRunArtifactLineageOK) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/lineage/artifacts/{name}][%d] getRunArtifactLineageOK  %+v", 200, o.Payload)
}

func (o *GetRunArtifactLineageOK) GetPayload() *service_model.V1RunArtifact {
	return o.Payload
}

func (o *GetRunArtifactLineageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1RunArtifact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunArtifactLineageNoContent creates a GetRunArtifactLineageNoContent with default headers values
func NewGetRunArtifactLineageNoContent() *GetRunArtifactLineageNoContent {
	return &GetRunArtifactLineageNoContent{}
}

/*
GetRunArtifactLineageNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetRunArtifactLineageNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get run artifact lineage no content response has a 2xx status code
func (o *GetRunArtifactLineageNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get run artifact lineage no content response has a 3xx status code
func (o *GetRunArtifactLineageNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run artifact lineage no content response has a 4xx status code
func (o *GetRunArtifactLineageNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get run artifact lineage no content response has a 5xx status code
func (o *GetRunArtifactLineageNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get run artifact lineage no content response a status code equal to that given
func (o *GetRunArtifactLineageNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get run artifact lineage no content response
func (o *GetRunArtifactLineageNoContent) Code() int {
	return 204
}

func (o *GetRunArtifactLineageNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/lineage/artifacts/{name}][%d] getRunArtifactLineageNoContent  %+v", 204, o.Payload)
}

func (o *GetRunArtifactLineageNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/lineage/artifacts/{name}][%d] getRunArtifactLineageNoContent  %+v", 204, o.Payload)
}

func (o *GetRunArtifactLineageNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunArtifactLineageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunArtifactLineageForbidden creates a GetRunArtifactLineageForbidden with default headers values
func NewGetRunArtifactLineageForbidden() *GetRunArtifactLineageForbidden {
	return &GetRunArtifactLineageForbidden{}
}

/*
GetRunArtifactLineageForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetRunArtifactLineageForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get run artifact lineage forbidden response has a 2xx status code
func (o *GetRunArtifactLineageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get run artifact lineage forbidden response has a 3xx status code
func (o *GetRunArtifactLineageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run artifact lineage forbidden response has a 4xx status code
func (o *GetRunArtifactLineageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get run artifact lineage forbidden response has a 5xx status code
func (o *GetRunArtifactLineageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get run artifact lineage forbidden response a status code equal to that given
func (o *GetRunArtifactLineageForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get run artifact lineage forbidden response
func (o *GetRunArtifactLineageForbidden) Code() int {
	return 403
}

func (o *GetRunArtifactLineageForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/lineage/artifacts/{name}][%d] getRunArtifactLineageForbidden  %+v", 403, o.Payload)
}

func (o *GetRunArtifactLineageForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/lineage/artifacts/{name}][%d] getRunArtifactLineageForbidden  %+v", 403, o.Payload)
}

func (o *GetRunArtifactLineageForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunArtifactLineageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunArtifactLineageNotFound creates a GetRunArtifactLineageNotFound with default headers values
func NewGetRunArtifactLineageNotFound() *GetRunArtifactLineageNotFound {
	return &GetRunArtifactLineageNotFound{}
}

/*
GetRunArtifactLineageNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetRunArtifactLineageNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get run artifact lineage not found response has a 2xx status code
func (o *GetRunArtifactLineageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get run artifact lineage not found response has a 3xx status code
func (o *GetRunArtifactLineageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run artifact lineage not found response has a 4xx status code
func (o *GetRunArtifactLineageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get run artifact lineage not found response has a 5xx status code
func (o *GetRunArtifactLineageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get run artifact lineage not found response a status code equal to that given
func (o *GetRunArtifactLineageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get run artifact lineage not found response
func (o *GetRunArtifactLineageNotFound) Code() int {
	return 404
}

func (o *GetRunArtifactLineageNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/lineage/artifacts/{name}][%d] getRunArtifactLineageNotFound  %+v", 404, o.Payload)
}

func (o *GetRunArtifactLineageNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/lineage/artifacts/{name}][%d] getRunArtifactLineageNotFound  %+v", 404, o.Payload)
}

func (o *GetRunArtifactLineageNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunArtifactLineageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunArtifactLineageDefault creates a GetRunArtifactLineageDefault with default headers values
func NewGetRunArtifactLineageDefault(code int) *GetRunArtifactLineageDefault {
	return &GetRunArtifactLineageDefault{
		_statusCode: code,
	}
}

/*
GetRunArtifactLineageDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetRunArtifactLineageDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this get run artifact lineage default response has a 2xx status code
func (o *GetRunArtifactLineageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get run artifact lineage default response has a 3xx status code
func (o *GetRunArtifactLineageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get run artifact lineage default response has a 4xx status code
func (o *GetRunArtifactLineageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get run artifact lineage default response has a 5xx status code
func (o *GetRunArtifactLineageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get run artifact lineage default response a status code equal to that given
func (o *GetRunArtifactLineageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get run artifact lineage default response
func (o *GetRunArtifactLineageDefault) Code() int {
	return o._statusCode
}

func (o *GetRunArtifactLineageDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/lineage/artifacts/{name}][%d] GetRunArtifactLineage default  %+v", o._statusCode, o.Payload)
}

func (o *GetRunArtifactLineageDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/lineage/artifacts/{name}][%d] GetRunArtifactLineage default  %+v", o._statusCode, o.Payload)
}

func (o *GetRunArtifactLineageDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetRunArtifactLineageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
