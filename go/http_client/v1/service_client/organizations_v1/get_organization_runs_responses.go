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

// GetOrganizationRunsReader is a Reader for the GetOrganizationRuns structure.
type GetOrganizationRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetOrganizationRunsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetOrganizationRunsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationRunsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetOrganizationRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOrganizationRunsOK creates a GetOrganizationRunsOK with default headers values
func NewGetOrganizationRunsOK() *GetOrganizationRunsOK {
	return &GetOrganizationRunsOK{}
}

/* GetOrganizationRunsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetOrganizationRunsOK struct {
	Payload *service_model.V1ListRunsResponse
}

// IsSuccess returns true when this get organization runs o k response has a 2xx status code
func (o *GetOrganizationRunsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization runs o k response has a 3xx status code
func (o *GetOrganizationRunsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization runs o k response has a 4xx status code
func (o *GetOrganizationRunsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization runs o k response has a 5xx status code
func (o *GetOrganizationRunsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization runs o k response a status code equal to that given
func (o *GetOrganizationRunsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization runs o k response
func (o *GetOrganizationRunsOK) Code() int {
	return 200
}

func (o *GetOrganizationRunsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/runs][%d] getOrganizationRunsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationRunsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/runs][%d] getOrganizationRunsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationRunsOK) GetPayload() *service_model.V1ListRunsResponse {
	return o.Payload
}

func (o *GetOrganizationRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListRunsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationRunsNoContent creates a GetOrganizationRunsNoContent with default headers values
func NewGetOrganizationRunsNoContent() *GetOrganizationRunsNoContent {
	return &GetOrganizationRunsNoContent{}
}

/* GetOrganizationRunsNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetOrganizationRunsNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get organization runs no content response has a 2xx status code
func (o *GetOrganizationRunsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization runs no content response has a 3xx status code
func (o *GetOrganizationRunsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization runs no content response has a 4xx status code
func (o *GetOrganizationRunsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization runs no content response has a 5xx status code
func (o *GetOrganizationRunsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization runs no content response a status code equal to that given
func (o *GetOrganizationRunsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get organization runs no content response
func (o *GetOrganizationRunsNoContent) Code() int {
	return 204
}

func (o *GetOrganizationRunsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/runs][%d] getOrganizationRunsNoContent  %+v", 204, o.Payload)
}

func (o *GetOrganizationRunsNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/runs][%d] getOrganizationRunsNoContent  %+v", 204, o.Payload)
}

func (o *GetOrganizationRunsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationRunsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationRunsForbidden creates a GetOrganizationRunsForbidden with default headers values
func NewGetOrganizationRunsForbidden() *GetOrganizationRunsForbidden {
	return &GetOrganizationRunsForbidden{}
}

/* GetOrganizationRunsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetOrganizationRunsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get organization runs forbidden response has a 2xx status code
func (o *GetOrganizationRunsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization runs forbidden response has a 3xx status code
func (o *GetOrganizationRunsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization runs forbidden response has a 4xx status code
func (o *GetOrganizationRunsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization runs forbidden response has a 5xx status code
func (o *GetOrganizationRunsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization runs forbidden response a status code equal to that given
func (o *GetOrganizationRunsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get organization runs forbidden response
func (o *GetOrganizationRunsForbidden) Code() int {
	return 403
}

func (o *GetOrganizationRunsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/runs][%d] getOrganizationRunsForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationRunsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/runs][%d] getOrganizationRunsForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationRunsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationRunsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationRunsNotFound creates a GetOrganizationRunsNotFound with default headers values
func NewGetOrganizationRunsNotFound() *GetOrganizationRunsNotFound {
	return &GetOrganizationRunsNotFound{}
}

/* GetOrganizationRunsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetOrganizationRunsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get organization runs not found response has a 2xx status code
func (o *GetOrganizationRunsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization runs not found response has a 3xx status code
func (o *GetOrganizationRunsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization runs not found response has a 4xx status code
func (o *GetOrganizationRunsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization runs not found response has a 5xx status code
func (o *GetOrganizationRunsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization runs not found response a status code equal to that given
func (o *GetOrganizationRunsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get organization runs not found response
func (o *GetOrganizationRunsNotFound) Code() int {
	return 404
}

func (o *GetOrganizationRunsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/runs][%d] getOrganizationRunsNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationRunsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/runs][%d] getOrganizationRunsNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationRunsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationRunsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationRunsDefault creates a GetOrganizationRunsDefault with default headers values
func NewGetOrganizationRunsDefault(code int) *GetOrganizationRunsDefault {
	return &GetOrganizationRunsDefault{
		_statusCode: code,
	}
}

/* GetOrganizationRunsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetOrganizationRunsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this get organization runs default response has a 2xx status code
func (o *GetOrganizationRunsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get organization runs default response has a 3xx status code
func (o *GetOrganizationRunsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get organization runs default response has a 4xx status code
func (o *GetOrganizationRunsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get organization runs default response has a 5xx status code
func (o *GetOrganizationRunsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get organization runs default response a status code equal to that given
func (o *GetOrganizationRunsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get organization runs default response
func (o *GetOrganizationRunsDefault) Code() int {
	return o._statusCode
}

func (o *GetOrganizationRunsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/runs][%d] GetOrganizationRuns default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrganizationRunsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/runs][%d] GetOrganizationRuns default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrganizationRunsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetOrganizationRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
