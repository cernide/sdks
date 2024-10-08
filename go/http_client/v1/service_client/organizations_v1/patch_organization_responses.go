// Code generated by go-swagger; DO NOT EDIT.

package organizations_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// PatchOrganizationReader is a Reader for the PatchOrganization structure.
type PatchOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchOrganizationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchOrganizationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchOrganizationOK creates a PatchOrganizationOK with default headers values
func NewPatchOrganizationOK() *PatchOrganizationOK {
	return &PatchOrganizationOK{}
}

/*
PatchOrganizationOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchOrganizationOK struct {
	Payload *service_model.V1Organization
}

// IsSuccess returns true when this patch organization o k response has a 2xx status code
func (o *PatchOrganizationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch organization o k response has a 3xx status code
func (o *PatchOrganizationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch organization o k response has a 4xx status code
func (o *PatchOrganizationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch organization o k response has a 5xx status code
func (o *PatchOrganizationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch organization o k response a status code equal to that given
func (o *PatchOrganizationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch organization o k response
func (o *PatchOrganizationOK) Code() int {
	return 200
}

func (o *PatchOrganizationOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}][%d] patchOrganizationOK  %+v", 200, o.Payload)
}

func (o *PatchOrganizationOK) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}][%d] patchOrganizationOK  %+v", 200, o.Payload)
}

func (o *PatchOrganizationOK) GetPayload() *service_model.V1Organization {
	return o.Payload
}

func (o *PatchOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Organization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOrganizationNoContent creates a PatchOrganizationNoContent with default headers values
func NewPatchOrganizationNoContent() *PatchOrganizationNoContent {
	return &PatchOrganizationNoContent{}
}

/*
PatchOrganizationNoContent describes a response with status code 204, with default header values.

No content.
*/
type PatchOrganizationNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this patch organization no content response has a 2xx status code
func (o *PatchOrganizationNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch organization no content response has a 3xx status code
func (o *PatchOrganizationNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch organization no content response has a 4xx status code
func (o *PatchOrganizationNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch organization no content response has a 5xx status code
func (o *PatchOrganizationNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this patch organization no content response a status code equal to that given
func (o *PatchOrganizationNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the patch organization no content response
func (o *PatchOrganizationNoContent) Code() int {
	return 204
}

func (o *PatchOrganizationNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}][%d] patchOrganizationNoContent  %+v", 204, o.Payload)
}

func (o *PatchOrganizationNoContent) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}][%d] patchOrganizationNoContent  %+v", 204, o.Payload)
}

func (o *PatchOrganizationNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchOrganizationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOrganizationForbidden creates a PatchOrganizationForbidden with default headers values
func NewPatchOrganizationForbidden() *PatchOrganizationForbidden {
	return &PatchOrganizationForbidden{}
}

/*
PatchOrganizationForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type PatchOrganizationForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this patch organization forbidden response has a 2xx status code
func (o *PatchOrganizationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch organization forbidden response has a 3xx status code
func (o *PatchOrganizationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch organization forbidden response has a 4xx status code
func (o *PatchOrganizationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch organization forbidden response has a 5xx status code
func (o *PatchOrganizationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch organization forbidden response a status code equal to that given
func (o *PatchOrganizationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the patch organization forbidden response
func (o *PatchOrganizationForbidden) Code() int {
	return 403
}

func (o *PatchOrganizationForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}][%d] patchOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *PatchOrganizationForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}][%d] patchOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *PatchOrganizationForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOrganizationNotFound creates a PatchOrganizationNotFound with default headers values
func NewPatchOrganizationNotFound() *PatchOrganizationNotFound {
	return &PatchOrganizationNotFound{}
}

/*
PatchOrganizationNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type PatchOrganizationNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this patch organization not found response has a 2xx status code
func (o *PatchOrganizationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch organization not found response has a 3xx status code
func (o *PatchOrganizationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch organization not found response has a 4xx status code
func (o *PatchOrganizationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch organization not found response has a 5xx status code
func (o *PatchOrganizationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch organization not found response a status code equal to that given
func (o *PatchOrganizationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the patch organization not found response
func (o *PatchOrganizationNotFound) Code() int {
	return 404
}

func (o *PatchOrganizationNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}][%d] patchOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *PatchOrganizationNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}][%d] patchOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *PatchOrganizationNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOrganizationDefault creates a PatchOrganizationDefault with default headers values
func NewPatchOrganizationDefault(code int) *PatchOrganizationDefault {
	return &PatchOrganizationDefault{
		_statusCode: code,
	}
}

/*
PatchOrganizationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PatchOrganizationDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this patch organization default response has a 2xx status code
func (o *PatchOrganizationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch organization default response has a 3xx status code
func (o *PatchOrganizationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch organization default response has a 4xx status code
func (o *PatchOrganizationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch organization default response has a 5xx status code
func (o *PatchOrganizationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch organization default response a status code equal to that given
func (o *PatchOrganizationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch organization default response
func (o *PatchOrganizationDefault) Code() int {
	return o._statusCode
}

func (o *PatchOrganizationDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}][%d] PatchOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *PatchOrganizationDefault) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}][%d] PatchOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *PatchOrganizationDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *PatchOrganizationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
