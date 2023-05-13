// Code generated by go-swagger; DO NOT EDIT.

package project_searches_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// PatchProjectSearchReader is a Reader for the PatchProjectSearch structure.
type PatchProjectSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchProjectSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchProjectSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchProjectSearchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchProjectSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchProjectSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchProjectSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchProjectSearchOK creates a PatchProjectSearchOK with default headers values
func NewPatchProjectSearchOK() *PatchProjectSearchOK {
	return &PatchProjectSearchOK{}
}

/*
PatchProjectSearchOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchProjectSearchOK struct {
	Payload *service_model.V1Search
}

// IsSuccess returns true when this patch project search o k response has a 2xx status code
func (o *PatchProjectSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch project search o k response has a 3xx status code
func (o *PatchProjectSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch project search o k response has a 4xx status code
func (o *PatchProjectSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch project search o k response has a 5xx status code
func (o *PatchProjectSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch project search o k response a status code equal to that given
func (o *PatchProjectSearchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch project search o k response
func (o *PatchProjectSearchOK) Code() int {
	return 200
}

func (o *PatchProjectSearchOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/searches/{search.uuid}][%d] patchProjectSearchOK  %+v", 200, o.Payload)
}

func (o *PatchProjectSearchOK) String() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/searches/{search.uuid}][%d] patchProjectSearchOK  %+v", 200, o.Payload)
}

func (o *PatchProjectSearchOK) GetPayload() *service_model.V1Search {
	return o.Payload
}

func (o *PatchProjectSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Search)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchProjectSearchNoContent creates a PatchProjectSearchNoContent with default headers values
func NewPatchProjectSearchNoContent() *PatchProjectSearchNoContent {
	return &PatchProjectSearchNoContent{}
}

/*
PatchProjectSearchNoContent describes a response with status code 204, with default header values.

No content.
*/
type PatchProjectSearchNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this patch project search no content response has a 2xx status code
func (o *PatchProjectSearchNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch project search no content response has a 3xx status code
func (o *PatchProjectSearchNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch project search no content response has a 4xx status code
func (o *PatchProjectSearchNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch project search no content response has a 5xx status code
func (o *PatchProjectSearchNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this patch project search no content response a status code equal to that given
func (o *PatchProjectSearchNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the patch project search no content response
func (o *PatchProjectSearchNoContent) Code() int {
	return 204
}

func (o *PatchProjectSearchNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/searches/{search.uuid}][%d] patchProjectSearchNoContent  %+v", 204, o.Payload)
}

func (o *PatchProjectSearchNoContent) String() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/searches/{search.uuid}][%d] patchProjectSearchNoContent  %+v", 204, o.Payload)
}

func (o *PatchProjectSearchNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchProjectSearchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchProjectSearchForbidden creates a PatchProjectSearchForbidden with default headers values
func NewPatchProjectSearchForbidden() *PatchProjectSearchForbidden {
	return &PatchProjectSearchForbidden{}
}

/*
PatchProjectSearchForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type PatchProjectSearchForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this patch project search forbidden response has a 2xx status code
func (o *PatchProjectSearchForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch project search forbidden response has a 3xx status code
func (o *PatchProjectSearchForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch project search forbidden response has a 4xx status code
func (o *PatchProjectSearchForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch project search forbidden response has a 5xx status code
func (o *PatchProjectSearchForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch project search forbidden response a status code equal to that given
func (o *PatchProjectSearchForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the patch project search forbidden response
func (o *PatchProjectSearchForbidden) Code() int {
	return 403
}

func (o *PatchProjectSearchForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/searches/{search.uuid}][%d] patchProjectSearchForbidden  %+v", 403, o.Payload)
}

func (o *PatchProjectSearchForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/searches/{search.uuid}][%d] patchProjectSearchForbidden  %+v", 403, o.Payload)
}

func (o *PatchProjectSearchForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchProjectSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchProjectSearchNotFound creates a PatchProjectSearchNotFound with default headers values
func NewPatchProjectSearchNotFound() *PatchProjectSearchNotFound {
	return &PatchProjectSearchNotFound{}
}

/*
PatchProjectSearchNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type PatchProjectSearchNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this patch project search not found response has a 2xx status code
func (o *PatchProjectSearchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch project search not found response has a 3xx status code
func (o *PatchProjectSearchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch project search not found response has a 4xx status code
func (o *PatchProjectSearchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch project search not found response has a 5xx status code
func (o *PatchProjectSearchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch project search not found response a status code equal to that given
func (o *PatchProjectSearchNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the patch project search not found response
func (o *PatchProjectSearchNotFound) Code() int {
	return 404
}

func (o *PatchProjectSearchNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/searches/{search.uuid}][%d] patchProjectSearchNotFound  %+v", 404, o.Payload)
}

func (o *PatchProjectSearchNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/searches/{search.uuid}][%d] patchProjectSearchNotFound  %+v", 404, o.Payload)
}

func (o *PatchProjectSearchNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchProjectSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchProjectSearchDefault creates a PatchProjectSearchDefault with default headers values
func NewPatchProjectSearchDefault(code int) *PatchProjectSearchDefault {
	return &PatchProjectSearchDefault{
		_statusCode: code,
	}
}

/*
PatchProjectSearchDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PatchProjectSearchDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this patch project search default response has a 2xx status code
func (o *PatchProjectSearchDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch project search default response has a 3xx status code
func (o *PatchProjectSearchDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch project search default response has a 4xx status code
func (o *PatchProjectSearchDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch project search default response has a 5xx status code
func (o *PatchProjectSearchDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch project search default response a status code equal to that given
func (o *PatchProjectSearchDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch project search default response
func (o *PatchProjectSearchDefault) Code() int {
	return o._statusCode
}

func (o *PatchProjectSearchDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/searches/{search.uuid}][%d] PatchProjectSearch default  %+v", o._statusCode, o.Payload)
}

func (o *PatchProjectSearchDefault) String() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/searches/{search.uuid}][%d] PatchProjectSearch default  %+v", o._statusCode, o.Payload)
}

func (o *PatchProjectSearchDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *PatchProjectSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
