// Code generated by go-swagger; DO NOT EDIT.

package teams_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// GetTeamRunsArtifactsLineageReader is a Reader for the GetTeamRunsArtifactsLineage structure.
type GetTeamRunsArtifactsLineageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamRunsArtifactsLineageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamRunsArtifactsLineageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetTeamRunsArtifactsLineageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetTeamRunsArtifactsLineageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTeamRunsArtifactsLineageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetTeamRunsArtifactsLineageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTeamRunsArtifactsLineageOK creates a GetTeamRunsArtifactsLineageOK with default headers values
func NewGetTeamRunsArtifactsLineageOK() *GetTeamRunsArtifactsLineageOK {
	return &GetTeamRunsArtifactsLineageOK{}
}

/*
GetTeamRunsArtifactsLineageOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetTeamRunsArtifactsLineageOK struct {
	Payload *service_model.V1ListRunArtifactsResponse
}

// IsSuccess returns true when this get team runs artifacts lineage o k response has a 2xx status code
func (o *GetTeamRunsArtifactsLineageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get team runs artifacts lineage o k response has a 3xx status code
func (o *GetTeamRunsArtifactsLineageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team runs artifacts lineage o k response has a 4xx status code
func (o *GetTeamRunsArtifactsLineageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team runs artifacts lineage o k response has a 5xx status code
func (o *GetTeamRunsArtifactsLineageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get team runs artifacts lineage o k response a status code equal to that given
func (o *GetTeamRunsArtifactsLineageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get team runs artifacts lineage o k response
func (o *GetTeamRunsArtifactsLineageOK) Code() int {
	return 200
}

func (o *GetTeamRunsArtifactsLineageOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{name}/runs/lineage/artifacts][%d] getTeamRunsArtifactsLineageOK  %+v", 200, o.Payload)
}

func (o *GetTeamRunsArtifactsLineageOK) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{name}/runs/lineage/artifacts][%d] getTeamRunsArtifactsLineageOK  %+v", 200, o.Payload)
}

func (o *GetTeamRunsArtifactsLineageOK) GetPayload() *service_model.V1ListRunArtifactsResponse {
	return o.Payload
}

func (o *GetTeamRunsArtifactsLineageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListRunArtifactsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamRunsArtifactsLineageNoContent creates a GetTeamRunsArtifactsLineageNoContent with default headers values
func NewGetTeamRunsArtifactsLineageNoContent() *GetTeamRunsArtifactsLineageNoContent {
	return &GetTeamRunsArtifactsLineageNoContent{}
}

/*
GetTeamRunsArtifactsLineageNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetTeamRunsArtifactsLineageNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get team runs artifacts lineage no content response has a 2xx status code
func (o *GetTeamRunsArtifactsLineageNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get team runs artifacts lineage no content response has a 3xx status code
func (o *GetTeamRunsArtifactsLineageNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team runs artifacts lineage no content response has a 4xx status code
func (o *GetTeamRunsArtifactsLineageNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team runs artifacts lineage no content response has a 5xx status code
func (o *GetTeamRunsArtifactsLineageNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get team runs artifacts lineage no content response a status code equal to that given
func (o *GetTeamRunsArtifactsLineageNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get team runs artifacts lineage no content response
func (o *GetTeamRunsArtifactsLineageNoContent) Code() int {
	return 204
}

func (o *GetTeamRunsArtifactsLineageNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{name}/runs/lineage/artifacts][%d] getTeamRunsArtifactsLineageNoContent  %+v", 204, o.Payload)
}

func (o *GetTeamRunsArtifactsLineageNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{name}/runs/lineage/artifacts][%d] getTeamRunsArtifactsLineageNoContent  %+v", 204, o.Payload)
}

func (o *GetTeamRunsArtifactsLineageNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetTeamRunsArtifactsLineageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamRunsArtifactsLineageForbidden creates a GetTeamRunsArtifactsLineageForbidden with default headers values
func NewGetTeamRunsArtifactsLineageForbidden() *GetTeamRunsArtifactsLineageForbidden {
	return &GetTeamRunsArtifactsLineageForbidden{}
}

/*
GetTeamRunsArtifactsLineageForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetTeamRunsArtifactsLineageForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get team runs artifacts lineage forbidden response has a 2xx status code
func (o *GetTeamRunsArtifactsLineageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team runs artifacts lineage forbidden response has a 3xx status code
func (o *GetTeamRunsArtifactsLineageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team runs artifacts lineage forbidden response has a 4xx status code
func (o *GetTeamRunsArtifactsLineageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team runs artifacts lineage forbidden response has a 5xx status code
func (o *GetTeamRunsArtifactsLineageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get team runs artifacts lineage forbidden response a status code equal to that given
func (o *GetTeamRunsArtifactsLineageForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get team runs artifacts lineage forbidden response
func (o *GetTeamRunsArtifactsLineageForbidden) Code() int {
	return 403
}

func (o *GetTeamRunsArtifactsLineageForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{name}/runs/lineage/artifacts][%d] getTeamRunsArtifactsLineageForbidden  %+v", 403, o.Payload)
}

func (o *GetTeamRunsArtifactsLineageForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{name}/runs/lineage/artifacts][%d] getTeamRunsArtifactsLineageForbidden  %+v", 403, o.Payload)
}

func (o *GetTeamRunsArtifactsLineageForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetTeamRunsArtifactsLineageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamRunsArtifactsLineageNotFound creates a GetTeamRunsArtifactsLineageNotFound with default headers values
func NewGetTeamRunsArtifactsLineageNotFound() *GetTeamRunsArtifactsLineageNotFound {
	return &GetTeamRunsArtifactsLineageNotFound{}
}

/*
GetTeamRunsArtifactsLineageNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetTeamRunsArtifactsLineageNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get team runs artifacts lineage not found response has a 2xx status code
func (o *GetTeamRunsArtifactsLineageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team runs artifacts lineage not found response has a 3xx status code
func (o *GetTeamRunsArtifactsLineageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team runs artifacts lineage not found response has a 4xx status code
func (o *GetTeamRunsArtifactsLineageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team runs artifacts lineage not found response has a 5xx status code
func (o *GetTeamRunsArtifactsLineageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get team runs artifacts lineage not found response a status code equal to that given
func (o *GetTeamRunsArtifactsLineageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get team runs artifacts lineage not found response
func (o *GetTeamRunsArtifactsLineageNotFound) Code() int {
	return 404
}

func (o *GetTeamRunsArtifactsLineageNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{name}/runs/lineage/artifacts][%d] getTeamRunsArtifactsLineageNotFound  %+v", 404, o.Payload)
}

func (o *GetTeamRunsArtifactsLineageNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{name}/runs/lineage/artifacts][%d] getTeamRunsArtifactsLineageNotFound  %+v", 404, o.Payload)
}

func (o *GetTeamRunsArtifactsLineageNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetTeamRunsArtifactsLineageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamRunsArtifactsLineageDefault creates a GetTeamRunsArtifactsLineageDefault with default headers values
func NewGetTeamRunsArtifactsLineageDefault(code int) *GetTeamRunsArtifactsLineageDefault {
	return &GetTeamRunsArtifactsLineageDefault{
		_statusCode: code,
	}
}

/*
GetTeamRunsArtifactsLineageDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetTeamRunsArtifactsLineageDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this get team runs artifacts lineage default response has a 2xx status code
func (o *GetTeamRunsArtifactsLineageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get team runs artifacts lineage default response has a 3xx status code
func (o *GetTeamRunsArtifactsLineageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get team runs artifacts lineage default response has a 4xx status code
func (o *GetTeamRunsArtifactsLineageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get team runs artifacts lineage default response has a 5xx status code
func (o *GetTeamRunsArtifactsLineageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get team runs artifacts lineage default response a status code equal to that given
func (o *GetTeamRunsArtifactsLineageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get team runs artifacts lineage default response
func (o *GetTeamRunsArtifactsLineageDefault) Code() int {
	return o._statusCode
}

func (o *GetTeamRunsArtifactsLineageDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{name}/runs/lineage/artifacts][%d] GetTeamRunsArtifactsLineage default  %+v", o._statusCode, o.Payload)
}

func (o *GetTeamRunsArtifactsLineageDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/{name}/runs/lineage/artifacts][%d] GetTeamRunsArtifactsLineage default  %+v", o._statusCode, o.Payload)
}

func (o *GetTeamRunsArtifactsLineageDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetTeamRunsArtifactsLineageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
