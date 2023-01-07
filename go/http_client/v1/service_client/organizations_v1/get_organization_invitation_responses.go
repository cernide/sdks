// Copyright 2018-2023 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package organizations_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// GetOrganizationInvitationReader is a Reader for the GetOrganizationInvitation structure.
type GetOrganizationInvitationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationInvitationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationInvitationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetOrganizationInvitationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetOrganizationInvitationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationInvitationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetOrganizationInvitationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOrganizationInvitationOK creates a GetOrganizationInvitationOK with default headers values
func NewGetOrganizationInvitationOK() *GetOrganizationInvitationOK {
	return &GetOrganizationInvitationOK{}
}

/* GetOrganizationInvitationOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetOrganizationInvitationOK struct {
	Payload *service_model.V1OrganizationMember
}

// IsSuccess returns true when this get organization invitation o k response has a 2xx status code
func (o *GetOrganizationInvitationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization invitation o k response has a 3xx status code
func (o *GetOrganizationInvitationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization invitation o k response has a 4xx status code
func (o *GetOrganizationInvitationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization invitation o k response has a 5xx status code
func (o *GetOrganizationInvitationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization invitation o k response a status code equal to that given
func (o *GetOrganizationInvitationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrganizationInvitationOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/invitations][%d] getOrganizationInvitationOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationInvitationOK) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/invitations][%d] getOrganizationInvitationOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationInvitationOK) GetPayload() *service_model.V1OrganizationMember {
	return o.Payload
}

func (o *GetOrganizationInvitationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1OrganizationMember)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationInvitationNoContent creates a GetOrganizationInvitationNoContent with default headers values
func NewGetOrganizationInvitationNoContent() *GetOrganizationInvitationNoContent {
	return &GetOrganizationInvitationNoContent{}
}

/* GetOrganizationInvitationNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetOrganizationInvitationNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get organization invitation no content response has a 2xx status code
func (o *GetOrganizationInvitationNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization invitation no content response has a 3xx status code
func (o *GetOrganizationInvitationNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization invitation no content response has a 4xx status code
func (o *GetOrganizationInvitationNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization invitation no content response has a 5xx status code
func (o *GetOrganizationInvitationNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization invitation no content response a status code equal to that given
func (o *GetOrganizationInvitationNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *GetOrganizationInvitationNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/invitations][%d] getOrganizationInvitationNoContent  %+v", 204, o.Payload)
}

func (o *GetOrganizationInvitationNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/invitations][%d] getOrganizationInvitationNoContent  %+v", 204, o.Payload)
}

func (o *GetOrganizationInvitationNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationInvitationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationInvitationForbidden creates a GetOrganizationInvitationForbidden with default headers values
func NewGetOrganizationInvitationForbidden() *GetOrganizationInvitationForbidden {
	return &GetOrganizationInvitationForbidden{}
}

/* GetOrganizationInvitationForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetOrganizationInvitationForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get organization invitation forbidden response has a 2xx status code
func (o *GetOrganizationInvitationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization invitation forbidden response has a 3xx status code
func (o *GetOrganizationInvitationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization invitation forbidden response has a 4xx status code
func (o *GetOrganizationInvitationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization invitation forbidden response has a 5xx status code
func (o *GetOrganizationInvitationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization invitation forbidden response a status code equal to that given
func (o *GetOrganizationInvitationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOrganizationInvitationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/invitations][%d] getOrganizationInvitationForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationInvitationForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/invitations][%d] getOrganizationInvitationForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationInvitationForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationInvitationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationInvitationNotFound creates a GetOrganizationInvitationNotFound with default headers values
func NewGetOrganizationInvitationNotFound() *GetOrganizationInvitationNotFound {
	return &GetOrganizationInvitationNotFound{}
}

/* GetOrganizationInvitationNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetOrganizationInvitationNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get organization invitation not found response has a 2xx status code
func (o *GetOrganizationInvitationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization invitation not found response has a 3xx status code
func (o *GetOrganizationInvitationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization invitation not found response has a 4xx status code
func (o *GetOrganizationInvitationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization invitation not found response has a 5xx status code
func (o *GetOrganizationInvitationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization invitation not found response a status code equal to that given
func (o *GetOrganizationInvitationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOrganizationInvitationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/invitations][%d] getOrganizationInvitationNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationInvitationNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/invitations][%d] getOrganizationInvitationNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationInvitationNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationInvitationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationInvitationDefault creates a GetOrganizationInvitationDefault with default headers values
func NewGetOrganizationInvitationDefault(code int) *GetOrganizationInvitationDefault {
	return &GetOrganizationInvitationDefault{
		_statusCode: code,
	}
}

/* GetOrganizationInvitationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetOrganizationInvitationDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the get organization invitation default response
func (o *GetOrganizationInvitationDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get organization invitation default response has a 2xx status code
func (o *GetOrganizationInvitationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get organization invitation default response has a 3xx status code
func (o *GetOrganizationInvitationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get organization invitation default response has a 4xx status code
func (o *GetOrganizationInvitationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get organization invitation default response has a 5xx status code
func (o *GetOrganizationInvitationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get organization invitation default response a status code equal to that given
func (o *GetOrganizationInvitationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetOrganizationInvitationDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/invitations][%d] GetOrganizationInvitation default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrganizationInvitationDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/invitations][%d] GetOrganizationInvitation default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrganizationInvitationDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetOrganizationInvitationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
