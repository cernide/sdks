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

package agents_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// PatchAgentReader is a Reader for the PatchAgent structure.
type PatchAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchAgentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchAgentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchAgentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchAgentOK creates a PatchAgentOK with default headers values
func NewPatchAgentOK() *PatchAgentOK {
	return &PatchAgentOK{}
}

/* PatchAgentOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchAgentOK struct {
	Payload *service_model.V1Agent
}

// IsSuccess returns true when this patch agent o k response has a 2xx status code
func (o *PatchAgentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch agent o k response has a 3xx status code
func (o *PatchAgentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch agent o k response has a 4xx status code
func (o *PatchAgentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch agent o k response has a 5xx status code
func (o *PatchAgentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch agent o k response a status code equal to that given
func (o *PatchAgentOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchAgentOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}][%d] patchAgentOK  %+v", 200, o.Payload)
}

func (o *PatchAgentOK) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}][%d] patchAgentOK  %+v", 200, o.Payload)
}

func (o *PatchAgentOK) GetPayload() *service_model.V1Agent {
	return o.Payload
}

func (o *PatchAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Agent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAgentNoContent creates a PatchAgentNoContent with default headers values
func NewPatchAgentNoContent() *PatchAgentNoContent {
	return &PatchAgentNoContent{}
}

/* PatchAgentNoContent describes a response with status code 204, with default header values.

No content.
*/
type PatchAgentNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this patch agent no content response has a 2xx status code
func (o *PatchAgentNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch agent no content response has a 3xx status code
func (o *PatchAgentNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch agent no content response has a 4xx status code
func (o *PatchAgentNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch agent no content response has a 5xx status code
func (o *PatchAgentNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this patch agent no content response a status code equal to that given
func (o *PatchAgentNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *PatchAgentNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}][%d] patchAgentNoContent  %+v", 204, o.Payload)
}

func (o *PatchAgentNoContent) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}][%d] patchAgentNoContent  %+v", 204, o.Payload)
}

func (o *PatchAgentNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchAgentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAgentForbidden creates a PatchAgentForbidden with default headers values
func NewPatchAgentForbidden() *PatchAgentForbidden {
	return &PatchAgentForbidden{}
}

/* PatchAgentForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type PatchAgentForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this patch agent forbidden response has a 2xx status code
func (o *PatchAgentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch agent forbidden response has a 3xx status code
func (o *PatchAgentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch agent forbidden response has a 4xx status code
func (o *PatchAgentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch agent forbidden response has a 5xx status code
func (o *PatchAgentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch agent forbidden response a status code equal to that given
func (o *PatchAgentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchAgentForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}][%d] patchAgentForbidden  %+v", 403, o.Payload)
}

func (o *PatchAgentForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}][%d] patchAgentForbidden  %+v", 403, o.Payload)
}

func (o *PatchAgentForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchAgentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAgentNotFound creates a PatchAgentNotFound with default headers values
func NewPatchAgentNotFound() *PatchAgentNotFound {
	return &PatchAgentNotFound{}
}

/* PatchAgentNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type PatchAgentNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this patch agent not found response has a 2xx status code
func (o *PatchAgentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch agent not found response has a 3xx status code
func (o *PatchAgentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch agent not found response has a 4xx status code
func (o *PatchAgentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch agent not found response has a 5xx status code
func (o *PatchAgentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch agent not found response a status code equal to that given
func (o *PatchAgentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchAgentNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}][%d] patchAgentNotFound  %+v", 404, o.Payload)
}

func (o *PatchAgentNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}][%d] patchAgentNotFound  %+v", 404, o.Payload)
}

func (o *PatchAgentNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchAgentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAgentDefault creates a PatchAgentDefault with default headers values
func NewPatchAgentDefault(code int) *PatchAgentDefault {
	return &PatchAgentDefault{
		_statusCode: code,
	}
}

/* PatchAgentDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PatchAgentDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the patch agent default response
func (o *PatchAgentDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this patch agent default response has a 2xx status code
func (o *PatchAgentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch agent default response has a 3xx status code
func (o *PatchAgentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch agent default response has a 4xx status code
func (o *PatchAgentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch agent default response has a 5xx status code
func (o *PatchAgentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch agent default response a status code equal to that given
func (o *PatchAgentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PatchAgentDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}][%d] PatchAgent default  %+v", o._statusCode, o.Payload)
}

func (o *PatchAgentDefault) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}][%d] PatchAgent default  %+v", o._statusCode, o.Payload)
}

func (o *PatchAgentDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *PatchAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
