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

// ListOrganizationsReader is a Reader for the ListOrganizations structure.
type ListOrganizationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOrganizationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOrganizationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListOrganizationsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListOrganizationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListOrganizationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListOrganizationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOrganizationsOK creates a ListOrganizationsOK with default headers values
func NewListOrganizationsOK() *ListOrganizationsOK {
	return &ListOrganizationsOK{}
}

/* ListOrganizationsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListOrganizationsOK struct {
	Payload *service_model.V1ListOrganizationsResponse
}

// IsSuccess returns true when this list organizations o k response has a 2xx status code
func (o *ListOrganizationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list organizations o k response has a 3xx status code
func (o *ListOrganizationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list organizations o k response has a 4xx status code
func (o *ListOrganizationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list organizations o k response has a 5xx status code
func (o *ListOrganizationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list organizations o k response a status code equal to that given
func (o *ListOrganizationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list organizations o k response
func (o *ListOrganizationsOK) Code() int {
	return 200
}

func (o *ListOrganizationsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/list][%d] listOrganizationsOK  %+v", 200, o.Payload)
}

func (o *ListOrganizationsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/list][%d] listOrganizationsOK  %+v", 200, o.Payload)
}

func (o *ListOrganizationsOK) GetPayload() *service_model.V1ListOrganizationsResponse {
	return o.Payload
}

func (o *ListOrganizationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListOrganizationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationsNoContent creates a ListOrganizationsNoContent with default headers values
func NewListOrganizationsNoContent() *ListOrganizationsNoContent {
	return &ListOrganizationsNoContent{}
}

/* ListOrganizationsNoContent describes a response with status code 204, with default header values.

No content.
*/
type ListOrganizationsNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this list organizations no content response has a 2xx status code
func (o *ListOrganizationsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list organizations no content response has a 3xx status code
func (o *ListOrganizationsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list organizations no content response has a 4xx status code
func (o *ListOrganizationsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this list organizations no content response has a 5xx status code
func (o *ListOrganizationsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this list organizations no content response a status code equal to that given
func (o *ListOrganizationsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the list organizations no content response
func (o *ListOrganizationsNoContent) Code() int {
	return 204
}

func (o *ListOrganizationsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/list][%d] listOrganizationsNoContent  %+v", 204, o.Payload)
}

func (o *ListOrganizationsNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/list][%d] listOrganizationsNoContent  %+v", 204, o.Payload)
}

func (o *ListOrganizationsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListOrganizationsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationsForbidden creates a ListOrganizationsForbidden with default headers values
func NewListOrganizationsForbidden() *ListOrganizationsForbidden {
	return &ListOrganizationsForbidden{}
}

/* ListOrganizationsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ListOrganizationsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this list organizations forbidden response has a 2xx status code
func (o *ListOrganizationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list organizations forbidden response has a 3xx status code
func (o *ListOrganizationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list organizations forbidden response has a 4xx status code
func (o *ListOrganizationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list organizations forbidden response has a 5xx status code
func (o *ListOrganizationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list organizations forbidden response a status code equal to that given
func (o *ListOrganizationsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list organizations forbidden response
func (o *ListOrganizationsForbidden) Code() int {
	return 403
}

func (o *ListOrganizationsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/list][%d] listOrganizationsForbidden  %+v", 403, o.Payload)
}

func (o *ListOrganizationsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/list][%d] listOrganizationsForbidden  %+v", 403, o.Payload)
}

func (o *ListOrganizationsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListOrganizationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationsNotFound creates a ListOrganizationsNotFound with default headers values
func NewListOrganizationsNotFound() *ListOrganizationsNotFound {
	return &ListOrganizationsNotFound{}
}

/* ListOrganizationsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ListOrganizationsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this list organizations not found response has a 2xx status code
func (o *ListOrganizationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list organizations not found response has a 3xx status code
func (o *ListOrganizationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list organizations not found response has a 4xx status code
func (o *ListOrganizationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list organizations not found response has a 5xx status code
func (o *ListOrganizationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list organizations not found response a status code equal to that given
func (o *ListOrganizationsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list organizations not found response
func (o *ListOrganizationsNotFound) Code() int {
	return 404
}

func (o *ListOrganizationsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/list][%d] listOrganizationsNotFound  %+v", 404, o.Payload)
}

func (o *ListOrganizationsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/list][%d] listOrganizationsNotFound  %+v", 404, o.Payload)
}

func (o *ListOrganizationsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListOrganizationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationsDefault creates a ListOrganizationsDefault with default headers values
func NewListOrganizationsDefault(code int) *ListOrganizationsDefault {
	return &ListOrganizationsDefault{
		_statusCode: code,
	}
}

/* ListOrganizationsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListOrganizationsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this list organizations default response has a 2xx status code
func (o *ListOrganizationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list organizations default response has a 3xx status code
func (o *ListOrganizationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list organizations default response has a 4xx status code
func (o *ListOrganizationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list organizations default response has a 5xx status code
func (o *ListOrganizationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list organizations default response a status code equal to that given
func (o *ListOrganizationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list organizations default response
func (o *ListOrganizationsDefault) Code() int {
	return o._statusCode
}

func (o *ListOrganizationsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/list][%d] ListOrganizations default  %+v", o._statusCode, o.Payload)
}

func (o *ListOrganizationsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/list][%d] ListOrganizations default  %+v", o._statusCode, o.Payload)
}

func (o *ListOrganizationsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListOrganizationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
