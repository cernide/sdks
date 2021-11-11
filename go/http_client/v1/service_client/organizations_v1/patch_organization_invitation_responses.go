// Copyright 2018-2021 Polyaxon, Inc.
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

// PatchOrganizationInvitationReader is a Reader for the PatchOrganizationInvitation structure.
type PatchOrganizationInvitationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchOrganizationInvitationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchOrganizationInvitationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchOrganizationInvitationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchOrganizationInvitationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchOrganizationInvitationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchOrganizationInvitationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchOrganizationInvitationOK creates a PatchOrganizationInvitationOK with default headers values
func NewPatchOrganizationInvitationOK() *PatchOrganizationInvitationOK {
	return &PatchOrganizationInvitationOK{}
}

/* PatchOrganizationInvitationOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchOrganizationInvitationOK struct {
	Payload *service_model.V1OrganizationMember
}

func (o *PatchOrganizationInvitationOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/invitations][%d] patchOrganizationInvitationOK  %+v", 200, o.Payload)
}
func (o *PatchOrganizationInvitationOK) GetPayload() *service_model.V1OrganizationMember {
	return o.Payload
}

func (o *PatchOrganizationInvitationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1OrganizationMember)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOrganizationInvitationNoContent creates a PatchOrganizationInvitationNoContent with default headers values
func NewPatchOrganizationInvitationNoContent() *PatchOrganizationInvitationNoContent {
	return &PatchOrganizationInvitationNoContent{}
}

/* PatchOrganizationInvitationNoContent describes a response with status code 204, with default header values.

No content.
*/
type PatchOrganizationInvitationNoContent struct {
	Payload interface{}
}

func (o *PatchOrganizationInvitationNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/invitations][%d] patchOrganizationInvitationNoContent  %+v", 204, o.Payload)
}
func (o *PatchOrganizationInvitationNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchOrganizationInvitationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOrganizationInvitationForbidden creates a PatchOrganizationInvitationForbidden with default headers values
func NewPatchOrganizationInvitationForbidden() *PatchOrganizationInvitationForbidden {
	return &PatchOrganizationInvitationForbidden{}
}

/* PatchOrganizationInvitationForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type PatchOrganizationInvitationForbidden struct {
	Payload interface{}
}

func (o *PatchOrganizationInvitationForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/invitations][%d] patchOrganizationInvitationForbidden  %+v", 403, o.Payload)
}
func (o *PatchOrganizationInvitationForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchOrganizationInvitationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOrganizationInvitationNotFound creates a PatchOrganizationInvitationNotFound with default headers values
func NewPatchOrganizationInvitationNotFound() *PatchOrganizationInvitationNotFound {
	return &PatchOrganizationInvitationNotFound{}
}

/* PatchOrganizationInvitationNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type PatchOrganizationInvitationNotFound struct {
	Payload interface{}
}

func (o *PatchOrganizationInvitationNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/invitations][%d] patchOrganizationInvitationNotFound  %+v", 404, o.Payload)
}
func (o *PatchOrganizationInvitationNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchOrganizationInvitationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOrganizationInvitationDefault creates a PatchOrganizationInvitationDefault with default headers values
func NewPatchOrganizationInvitationDefault(code int) *PatchOrganizationInvitationDefault {
	return &PatchOrganizationInvitationDefault{
		_statusCode: code,
	}
}

/* PatchOrganizationInvitationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PatchOrganizationInvitationDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the patch organization invitation default response
func (o *PatchOrganizationInvitationDefault) Code() int {
	return o._statusCode
}

func (o *PatchOrganizationInvitationDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/invitations][%d] PatchOrganizationInvitation default  %+v", o._statusCode, o.Payload)
}
func (o *PatchOrganizationInvitationDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *PatchOrganizationInvitationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
