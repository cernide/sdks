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

// SyncAgentReader is a Reader for the SyncAgent structure.
type SyncAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSyncAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewSyncAgentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSyncAgentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSyncAgentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewSyncAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSyncAgentOK creates a SyncAgentOK with default headers values
func NewSyncAgentOK() *SyncAgentOK {
	return &SyncAgentOK{}
}

/* SyncAgentOK describes a response with status code 200, with default header values.

A successful response.
*/
type SyncAgentOK struct {
}

// IsSuccess returns true when this sync agent o k response has a 2xx status code
func (o *SyncAgentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sync agent o k response has a 3xx status code
func (o *SyncAgentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync agent o k response has a 4xx status code
func (o *SyncAgentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this sync agent o k response has a 5xx status code
func (o *SyncAgentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this sync agent o k response a status code equal to that given
func (o *SyncAgentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the sync agent o k response
func (o *SyncAgentOK) Code() int {
	return 200
}

func (o *SyncAgentOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}/sync][%d] syncAgentOK ", 200)
}

func (o *SyncAgentOK) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}/sync][%d] syncAgentOK ", 200)
}

func (o *SyncAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSyncAgentNoContent creates a SyncAgentNoContent with default headers values
func NewSyncAgentNoContent() *SyncAgentNoContent {
	return &SyncAgentNoContent{}
}

/* SyncAgentNoContent describes a response with status code 204, with default header values.

No content.
*/
type SyncAgentNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this sync agent no content response has a 2xx status code
func (o *SyncAgentNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sync agent no content response has a 3xx status code
func (o *SyncAgentNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync agent no content response has a 4xx status code
func (o *SyncAgentNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this sync agent no content response has a 5xx status code
func (o *SyncAgentNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this sync agent no content response a status code equal to that given
func (o *SyncAgentNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the sync agent no content response
func (o *SyncAgentNoContent) Code() int {
	return 204
}

func (o *SyncAgentNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}/sync][%d] syncAgentNoContent  %+v", 204, o.Payload)
}

func (o *SyncAgentNoContent) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}/sync][%d] syncAgentNoContent  %+v", 204, o.Payload)
}

func (o *SyncAgentNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *SyncAgentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncAgentForbidden creates a SyncAgentForbidden with default headers values
func NewSyncAgentForbidden() *SyncAgentForbidden {
	return &SyncAgentForbidden{}
}

/* SyncAgentForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type SyncAgentForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this sync agent forbidden response has a 2xx status code
func (o *SyncAgentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this sync agent forbidden response has a 3xx status code
func (o *SyncAgentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync agent forbidden response has a 4xx status code
func (o *SyncAgentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this sync agent forbidden response has a 5xx status code
func (o *SyncAgentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this sync agent forbidden response a status code equal to that given
func (o *SyncAgentForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the sync agent forbidden response
func (o *SyncAgentForbidden) Code() int {
	return 403
}

func (o *SyncAgentForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}/sync][%d] syncAgentForbidden  %+v", 403, o.Payload)
}

func (o *SyncAgentForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}/sync][%d] syncAgentForbidden  %+v", 403, o.Payload)
}

func (o *SyncAgentForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *SyncAgentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncAgentNotFound creates a SyncAgentNotFound with default headers values
func NewSyncAgentNotFound() *SyncAgentNotFound {
	return &SyncAgentNotFound{}
}

/* SyncAgentNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type SyncAgentNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this sync agent not found response has a 2xx status code
func (o *SyncAgentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this sync agent not found response has a 3xx status code
func (o *SyncAgentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync agent not found response has a 4xx status code
func (o *SyncAgentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this sync agent not found response has a 5xx status code
func (o *SyncAgentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this sync agent not found response a status code equal to that given
func (o *SyncAgentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the sync agent not found response
func (o *SyncAgentNotFound) Code() int {
	return 404
}

func (o *SyncAgentNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}/sync][%d] syncAgentNotFound  %+v", 404, o.Payload)
}

func (o *SyncAgentNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}/sync][%d] syncAgentNotFound  %+v", 404, o.Payload)
}

func (o *SyncAgentNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *SyncAgentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncAgentDefault creates a SyncAgentDefault with default headers values
func NewSyncAgentDefault(code int) *SyncAgentDefault {
	return &SyncAgentDefault{
		_statusCode: code,
	}
}

/* SyncAgentDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SyncAgentDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this sync agent default response has a 2xx status code
func (o *SyncAgentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this sync agent default response has a 3xx status code
func (o *SyncAgentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this sync agent default response has a 4xx status code
func (o *SyncAgentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this sync agent default response has a 5xx status code
func (o *SyncAgentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this sync agent default response a status code equal to that given
func (o *SyncAgentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the sync agent default response
func (o *SyncAgentDefault) Code() int {
	return o._statusCode
}

func (o *SyncAgentDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}/sync][%d] SyncAgent default  %+v", o._statusCode, o.Payload)
}

func (o *SyncAgentDefault) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent.uuid}/sync][%d] SyncAgent default  %+v", o._statusCode, o.Payload)
}

func (o *SyncAgentDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *SyncAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
