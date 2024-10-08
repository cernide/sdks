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

// PatchOrganizationMemberReader is a Reader for the PatchOrganizationMember structure.
type PatchOrganizationMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchOrganizationMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchOrganizationMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchOrganizationMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchOrganizationMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchOrganizationMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchOrganizationMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchOrganizationMemberOK creates a PatchOrganizationMemberOK with default headers values
func NewPatchOrganizationMemberOK() *PatchOrganizationMemberOK {
	return &PatchOrganizationMemberOK{}
}

/*
PatchOrganizationMemberOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchOrganizationMemberOK struct {
	Payload *service_model.V1OrganizationMember
}

// IsSuccess returns true when this patch organization member o k response has a 2xx status code
func (o *PatchOrganizationMemberOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch organization member o k response has a 3xx status code
func (o *PatchOrganizationMemberOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch organization member o k response has a 4xx status code
func (o *PatchOrganizationMemberOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch organization member o k response has a 5xx status code
func (o *PatchOrganizationMemberOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch organization member o k response a status code equal to that given
func (o *PatchOrganizationMemberOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch organization member o k response
func (o *PatchOrganizationMemberOK) Code() int {
	return 200
}

func (o *PatchOrganizationMemberOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/members/{member.user}][%d] patchOrganizationMemberOK  %+v", 200, o.Payload)
}

func (o *PatchOrganizationMemberOK) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/members/{member.user}][%d] patchOrganizationMemberOK  %+v", 200, o.Payload)
}

func (o *PatchOrganizationMemberOK) GetPayload() *service_model.V1OrganizationMember {
	return o.Payload
}

func (o *PatchOrganizationMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1OrganizationMember)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOrganizationMemberNoContent creates a PatchOrganizationMemberNoContent with default headers values
func NewPatchOrganizationMemberNoContent() *PatchOrganizationMemberNoContent {
	return &PatchOrganizationMemberNoContent{}
}

/*
PatchOrganizationMemberNoContent describes a response with status code 204, with default header values.

No content.
*/
type PatchOrganizationMemberNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this patch organization member no content response has a 2xx status code
func (o *PatchOrganizationMemberNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch organization member no content response has a 3xx status code
func (o *PatchOrganizationMemberNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch organization member no content response has a 4xx status code
func (o *PatchOrganizationMemberNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch organization member no content response has a 5xx status code
func (o *PatchOrganizationMemberNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this patch organization member no content response a status code equal to that given
func (o *PatchOrganizationMemberNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the patch organization member no content response
func (o *PatchOrganizationMemberNoContent) Code() int {
	return 204
}

func (o *PatchOrganizationMemberNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/members/{member.user}][%d] patchOrganizationMemberNoContent  %+v", 204, o.Payload)
}

func (o *PatchOrganizationMemberNoContent) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/members/{member.user}][%d] patchOrganizationMemberNoContent  %+v", 204, o.Payload)
}

func (o *PatchOrganizationMemberNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchOrganizationMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOrganizationMemberForbidden creates a PatchOrganizationMemberForbidden with default headers values
func NewPatchOrganizationMemberForbidden() *PatchOrganizationMemberForbidden {
	return &PatchOrganizationMemberForbidden{}
}

/*
PatchOrganizationMemberForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type PatchOrganizationMemberForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this patch organization member forbidden response has a 2xx status code
func (o *PatchOrganizationMemberForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch organization member forbidden response has a 3xx status code
func (o *PatchOrganizationMemberForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch organization member forbidden response has a 4xx status code
func (o *PatchOrganizationMemberForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch organization member forbidden response has a 5xx status code
func (o *PatchOrganizationMemberForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch organization member forbidden response a status code equal to that given
func (o *PatchOrganizationMemberForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the patch organization member forbidden response
func (o *PatchOrganizationMemberForbidden) Code() int {
	return 403
}

func (o *PatchOrganizationMemberForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/members/{member.user}][%d] patchOrganizationMemberForbidden  %+v", 403, o.Payload)
}

func (o *PatchOrganizationMemberForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/members/{member.user}][%d] patchOrganizationMemberForbidden  %+v", 403, o.Payload)
}

func (o *PatchOrganizationMemberForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchOrganizationMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOrganizationMemberNotFound creates a PatchOrganizationMemberNotFound with default headers values
func NewPatchOrganizationMemberNotFound() *PatchOrganizationMemberNotFound {
	return &PatchOrganizationMemberNotFound{}
}

/*
PatchOrganizationMemberNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type PatchOrganizationMemberNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this patch organization member not found response has a 2xx status code
func (o *PatchOrganizationMemberNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch organization member not found response has a 3xx status code
func (o *PatchOrganizationMemberNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch organization member not found response has a 4xx status code
func (o *PatchOrganizationMemberNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch organization member not found response has a 5xx status code
func (o *PatchOrganizationMemberNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch organization member not found response a status code equal to that given
func (o *PatchOrganizationMemberNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the patch organization member not found response
func (o *PatchOrganizationMemberNotFound) Code() int {
	return 404
}

func (o *PatchOrganizationMemberNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/members/{member.user}][%d] patchOrganizationMemberNotFound  %+v", 404, o.Payload)
}

func (o *PatchOrganizationMemberNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/members/{member.user}][%d] patchOrganizationMemberNotFound  %+v", 404, o.Payload)
}

func (o *PatchOrganizationMemberNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchOrganizationMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOrganizationMemberDefault creates a PatchOrganizationMemberDefault with default headers values
func NewPatchOrganizationMemberDefault(code int) *PatchOrganizationMemberDefault {
	return &PatchOrganizationMemberDefault{
		_statusCode: code,
	}
}

/*
PatchOrganizationMemberDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PatchOrganizationMemberDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this patch organization member default response has a 2xx status code
func (o *PatchOrganizationMemberDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch organization member default response has a 3xx status code
func (o *PatchOrganizationMemberDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch organization member default response has a 4xx status code
func (o *PatchOrganizationMemberDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch organization member default response has a 5xx status code
func (o *PatchOrganizationMemberDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch organization member default response a status code equal to that given
func (o *PatchOrganizationMemberDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch organization member default response
func (o *PatchOrganizationMemberDefault) Code() int {
	return o._statusCode
}

func (o *PatchOrganizationMemberDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/members/{member.user}][%d] PatchOrganizationMember default  %+v", o._statusCode, o.Payload)
}

func (o *PatchOrganizationMemberDefault) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/members/{member.user}][%d] PatchOrganizationMember default  %+v", o._statusCode, o.Payload)
}

func (o *PatchOrganizationMemberDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *PatchOrganizationMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
