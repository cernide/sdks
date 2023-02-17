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

package service_accounts_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListServiceAccountNamesReader is a Reader for the ListServiceAccountNames structure.
type ListServiceAccountNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServiceAccountNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServiceAccountNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListServiceAccountNamesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListServiceAccountNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListServiceAccountNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListServiceAccountNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListServiceAccountNamesOK creates a ListServiceAccountNamesOK with default headers values
func NewListServiceAccountNamesOK() *ListServiceAccountNamesOK {
	return &ListServiceAccountNamesOK{}
}

/* ListServiceAccountNamesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListServiceAccountNamesOK struct {
	Payload *service_model.V1ListServiceAccountsResponse
}

// IsSuccess returns true when this list service account names o k response has a 2xx status code
func (o *ListServiceAccountNamesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list service account names o k response has a 3xx status code
func (o *ListServiceAccountNamesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list service account names o k response has a 4xx status code
func (o *ListServiceAccountNamesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list service account names o k response has a 5xx status code
func (o *ListServiceAccountNamesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list service account names o k response a status code equal to that given
func (o *ListServiceAccountNamesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list service account names o k response
func (o *ListServiceAccountNamesOK) Code() int {
	return 200
}

func (o *ListServiceAccountNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/sa/names][%d] listServiceAccountNamesOK  %+v", 200, o.Payload)
}

func (o *ListServiceAccountNamesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/sa/names][%d] listServiceAccountNamesOK  %+v", 200, o.Payload)
}

func (o *ListServiceAccountNamesOK) GetPayload() *service_model.V1ListServiceAccountsResponse {
	return o.Payload
}

func (o *ListServiceAccountNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListServiceAccountsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceAccountNamesNoContent creates a ListServiceAccountNamesNoContent with default headers values
func NewListServiceAccountNamesNoContent() *ListServiceAccountNamesNoContent {
	return &ListServiceAccountNamesNoContent{}
}

/* ListServiceAccountNamesNoContent describes a response with status code 204, with default header values.

No content.
*/
type ListServiceAccountNamesNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this list service account names no content response has a 2xx status code
func (o *ListServiceAccountNamesNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list service account names no content response has a 3xx status code
func (o *ListServiceAccountNamesNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list service account names no content response has a 4xx status code
func (o *ListServiceAccountNamesNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this list service account names no content response has a 5xx status code
func (o *ListServiceAccountNamesNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this list service account names no content response a status code equal to that given
func (o *ListServiceAccountNamesNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the list service account names no content response
func (o *ListServiceAccountNamesNoContent) Code() int {
	return 204
}

func (o *ListServiceAccountNamesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/sa/names][%d] listServiceAccountNamesNoContent  %+v", 204, o.Payload)
}

func (o *ListServiceAccountNamesNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/sa/names][%d] listServiceAccountNamesNoContent  %+v", 204, o.Payload)
}

func (o *ListServiceAccountNamesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListServiceAccountNamesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceAccountNamesForbidden creates a ListServiceAccountNamesForbidden with default headers values
func NewListServiceAccountNamesForbidden() *ListServiceAccountNamesForbidden {
	return &ListServiceAccountNamesForbidden{}
}

/* ListServiceAccountNamesForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ListServiceAccountNamesForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this list service account names forbidden response has a 2xx status code
func (o *ListServiceAccountNamesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list service account names forbidden response has a 3xx status code
func (o *ListServiceAccountNamesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list service account names forbidden response has a 4xx status code
func (o *ListServiceAccountNamesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list service account names forbidden response has a 5xx status code
func (o *ListServiceAccountNamesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list service account names forbidden response a status code equal to that given
func (o *ListServiceAccountNamesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list service account names forbidden response
func (o *ListServiceAccountNamesForbidden) Code() int {
	return 403
}

func (o *ListServiceAccountNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/sa/names][%d] listServiceAccountNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListServiceAccountNamesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/sa/names][%d] listServiceAccountNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListServiceAccountNamesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListServiceAccountNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceAccountNamesNotFound creates a ListServiceAccountNamesNotFound with default headers values
func NewListServiceAccountNamesNotFound() *ListServiceAccountNamesNotFound {
	return &ListServiceAccountNamesNotFound{}
}

/* ListServiceAccountNamesNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ListServiceAccountNamesNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this list service account names not found response has a 2xx status code
func (o *ListServiceAccountNamesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list service account names not found response has a 3xx status code
func (o *ListServiceAccountNamesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list service account names not found response has a 4xx status code
func (o *ListServiceAccountNamesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list service account names not found response has a 5xx status code
func (o *ListServiceAccountNamesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list service account names not found response a status code equal to that given
func (o *ListServiceAccountNamesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list service account names not found response
func (o *ListServiceAccountNamesNotFound) Code() int {
	return 404
}

func (o *ListServiceAccountNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/sa/names][%d] listServiceAccountNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListServiceAccountNamesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/sa/names][%d] listServiceAccountNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListServiceAccountNamesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListServiceAccountNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceAccountNamesDefault creates a ListServiceAccountNamesDefault with default headers values
func NewListServiceAccountNamesDefault(code int) *ListServiceAccountNamesDefault {
	return &ListServiceAccountNamesDefault{
		_statusCode: code,
	}
}

/* ListServiceAccountNamesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListServiceAccountNamesDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this list service account names default response has a 2xx status code
func (o *ListServiceAccountNamesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list service account names default response has a 3xx status code
func (o *ListServiceAccountNamesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list service account names default response has a 4xx status code
func (o *ListServiceAccountNamesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list service account names default response has a 5xx status code
func (o *ListServiceAccountNamesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list service account names default response a status code equal to that given
func (o *ListServiceAccountNamesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list service account names default response
func (o *ListServiceAccountNamesDefault) Code() int {
	return o._statusCode
}

func (o *ListServiceAccountNamesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/sa/names][%d] ListServiceAccountNames default  %+v", o._statusCode, o.Payload)
}

func (o *ListServiceAccountNamesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/sa/names][%d] ListServiceAccountNames default  %+v", o._statusCode, o.Payload)
}

func (o *ListServiceAccountNamesDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListServiceAccountNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
