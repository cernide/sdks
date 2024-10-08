// Code generated by go-swagger; DO NOT EDIT.

package tags_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// PatchTagReader is a Reader for the PatchTag structure.
type PatchTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchTagNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchTagForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchTagDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchTagOK creates a PatchTagOK with default headers values
func NewPatchTagOK() *PatchTagOK {
	return &PatchTagOK{}
}

/*
PatchTagOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchTagOK struct {
	Payload *service_model.V1Tag
}

// IsSuccess returns true when this patch tag o k response has a 2xx status code
func (o *PatchTagOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch tag o k response has a 3xx status code
func (o *PatchTagOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch tag o k response has a 4xx status code
func (o *PatchTagOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch tag o k response has a 5xx status code
func (o *PatchTagOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch tag o k response a status code equal to that given
func (o *PatchTagOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch tag o k response
func (o *PatchTagOK) Code() int {
	return 200
}

func (o *PatchTagOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/tags/{tag.uuid}][%d] patchTagOK  %+v", 200, o.Payload)
}

func (o *PatchTagOK) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/tags/{tag.uuid}][%d] patchTagOK  %+v", 200, o.Payload)
}

func (o *PatchTagOK) GetPayload() *service_model.V1Tag {
	return o.Payload
}

func (o *PatchTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTagNoContent creates a PatchTagNoContent with default headers values
func NewPatchTagNoContent() *PatchTagNoContent {
	return &PatchTagNoContent{}
}

/*
PatchTagNoContent describes a response with status code 204, with default header values.

No content.
*/
type PatchTagNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this patch tag no content response has a 2xx status code
func (o *PatchTagNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch tag no content response has a 3xx status code
func (o *PatchTagNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch tag no content response has a 4xx status code
func (o *PatchTagNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch tag no content response has a 5xx status code
func (o *PatchTagNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this patch tag no content response a status code equal to that given
func (o *PatchTagNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the patch tag no content response
func (o *PatchTagNoContent) Code() int {
	return 204
}

func (o *PatchTagNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/tags/{tag.uuid}][%d] patchTagNoContent  %+v", 204, o.Payload)
}

func (o *PatchTagNoContent) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/tags/{tag.uuid}][%d] patchTagNoContent  %+v", 204, o.Payload)
}

func (o *PatchTagNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchTagNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTagForbidden creates a PatchTagForbidden with default headers values
func NewPatchTagForbidden() *PatchTagForbidden {
	return &PatchTagForbidden{}
}

/*
PatchTagForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type PatchTagForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this patch tag forbidden response has a 2xx status code
func (o *PatchTagForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch tag forbidden response has a 3xx status code
func (o *PatchTagForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch tag forbidden response has a 4xx status code
func (o *PatchTagForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch tag forbidden response has a 5xx status code
func (o *PatchTagForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch tag forbidden response a status code equal to that given
func (o *PatchTagForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the patch tag forbidden response
func (o *PatchTagForbidden) Code() int {
	return 403
}

func (o *PatchTagForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/tags/{tag.uuid}][%d] patchTagForbidden  %+v", 403, o.Payload)
}

func (o *PatchTagForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/tags/{tag.uuid}][%d] patchTagForbidden  %+v", 403, o.Payload)
}

func (o *PatchTagForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchTagForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTagNotFound creates a PatchTagNotFound with default headers values
func NewPatchTagNotFound() *PatchTagNotFound {
	return &PatchTagNotFound{}
}

/*
PatchTagNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type PatchTagNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this patch tag not found response has a 2xx status code
func (o *PatchTagNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch tag not found response has a 3xx status code
func (o *PatchTagNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch tag not found response has a 4xx status code
func (o *PatchTagNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch tag not found response has a 5xx status code
func (o *PatchTagNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch tag not found response a status code equal to that given
func (o *PatchTagNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the patch tag not found response
func (o *PatchTagNotFound) Code() int {
	return 404
}

func (o *PatchTagNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/tags/{tag.uuid}][%d] patchTagNotFound  %+v", 404, o.Payload)
}

func (o *PatchTagNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/tags/{tag.uuid}][%d] patchTagNotFound  %+v", 404, o.Payload)
}

func (o *PatchTagNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTagDefault creates a PatchTagDefault with default headers values
func NewPatchTagDefault(code int) *PatchTagDefault {
	return &PatchTagDefault{
		_statusCode: code,
	}
}

/*
PatchTagDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PatchTagDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this patch tag default response has a 2xx status code
func (o *PatchTagDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch tag default response has a 3xx status code
func (o *PatchTagDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch tag default response has a 4xx status code
func (o *PatchTagDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch tag default response has a 5xx status code
func (o *PatchTagDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch tag default response a status code equal to that given
func (o *PatchTagDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch tag default response
func (o *PatchTagDefault) Code() int {
	return o._statusCode
}

func (o *PatchTagDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/tags/{tag.uuid}][%d] PatchTag default  %+v", o._statusCode, o.Payload)
}

func (o *PatchTagDefault) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/tags/{tag.uuid}][%d] PatchTag default  %+v", o._statusCode, o.Payload)
}

func (o *PatchTagDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *PatchTagDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
