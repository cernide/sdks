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

package project_searches_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListProjectSearchNamesReader is a Reader for the ListProjectSearchNames structure.
type ListProjectSearchNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectSearchNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectSearchNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListProjectSearchNamesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListProjectSearchNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListProjectSearchNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListProjectSearchNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectSearchNamesOK creates a ListProjectSearchNamesOK with default headers values
func NewListProjectSearchNamesOK() *ListProjectSearchNamesOK {
	return &ListProjectSearchNamesOK{}
}

/* ListProjectSearchNamesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListProjectSearchNamesOK struct {
	Payload *service_model.V1ListSearchesResponse
}

// IsSuccess returns true when this list project search names o k response has a 2xx status code
func (o *ListProjectSearchNamesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project search names o k response has a 3xx status code
func (o *ListProjectSearchNamesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project search names o k response has a 4xx status code
func (o *ListProjectSearchNamesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project search names o k response has a 5xx status code
func (o *ListProjectSearchNamesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project search names o k response a status code equal to that given
func (o *ListProjectSearchNamesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListProjectSearchNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/searches/names][%d] listProjectSearchNamesOK  %+v", 200, o.Payload)
}

func (o *ListProjectSearchNamesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/searches/names][%d] listProjectSearchNamesOK  %+v", 200, o.Payload)
}

func (o *ListProjectSearchNamesOK) GetPayload() *service_model.V1ListSearchesResponse {
	return o.Payload
}

func (o *ListProjectSearchNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListSearchesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectSearchNamesNoContent creates a ListProjectSearchNamesNoContent with default headers values
func NewListProjectSearchNamesNoContent() *ListProjectSearchNamesNoContent {
	return &ListProjectSearchNamesNoContent{}
}

/* ListProjectSearchNamesNoContent describes a response with status code 204, with default header values.

No content.
*/
type ListProjectSearchNamesNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this list project search names no content response has a 2xx status code
func (o *ListProjectSearchNamesNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project search names no content response has a 3xx status code
func (o *ListProjectSearchNamesNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project search names no content response has a 4xx status code
func (o *ListProjectSearchNamesNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project search names no content response has a 5xx status code
func (o *ListProjectSearchNamesNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this list project search names no content response a status code equal to that given
func (o *ListProjectSearchNamesNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ListProjectSearchNamesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/searches/names][%d] listProjectSearchNamesNoContent  %+v", 204, o.Payload)
}

func (o *ListProjectSearchNamesNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/searches/names][%d] listProjectSearchNamesNoContent  %+v", 204, o.Payload)
}

func (o *ListProjectSearchNamesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListProjectSearchNamesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectSearchNamesForbidden creates a ListProjectSearchNamesForbidden with default headers values
func NewListProjectSearchNamesForbidden() *ListProjectSearchNamesForbidden {
	return &ListProjectSearchNamesForbidden{}
}

/* ListProjectSearchNamesForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ListProjectSearchNamesForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this list project search names forbidden response has a 2xx status code
func (o *ListProjectSearchNamesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project search names forbidden response has a 3xx status code
func (o *ListProjectSearchNamesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project search names forbidden response has a 4xx status code
func (o *ListProjectSearchNamesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project search names forbidden response has a 5xx status code
func (o *ListProjectSearchNamesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list project search names forbidden response a status code equal to that given
func (o *ListProjectSearchNamesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListProjectSearchNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/searches/names][%d] listProjectSearchNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListProjectSearchNamesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/searches/names][%d] listProjectSearchNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListProjectSearchNamesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListProjectSearchNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectSearchNamesNotFound creates a ListProjectSearchNamesNotFound with default headers values
func NewListProjectSearchNamesNotFound() *ListProjectSearchNamesNotFound {
	return &ListProjectSearchNamesNotFound{}
}

/* ListProjectSearchNamesNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ListProjectSearchNamesNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this list project search names not found response has a 2xx status code
func (o *ListProjectSearchNamesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project search names not found response has a 3xx status code
func (o *ListProjectSearchNamesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project search names not found response has a 4xx status code
func (o *ListProjectSearchNamesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project search names not found response has a 5xx status code
func (o *ListProjectSearchNamesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list project search names not found response a status code equal to that given
func (o *ListProjectSearchNamesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ListProjectSearchNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/searches/names][%d] listProjectSearchNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListProjectSearchNamesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/searches/names][%d] listProjectSearchNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListProjectSearchNamesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListProjectSearchNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectSearchNamesDefault creates a ListProjectSearchNamesDefault with default headers values
func NewListProjectSearchNamesDefault(code int) *ListProjectSearchNamesDefault {
	return &ListProjectSearchNamesDefault{
		_statusCode: code,
	}
}

/* ListProjectSearchNamesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListProjectSearchNamesDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the list project search names default response
func (o *ListProjectSearchNamesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list project search names default response has a 2xx status code
func (o *ListProjectSearchNamesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project search names default response has a 3xx status code
func (o *ListProjectSearchNamesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project search names default response has a 4xx status code
func (o *ListProjectSearchNamesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project search names default response has a 5xx status code
func (o *ListProjectSearchNamesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project search names default response a status code equal to that given
func (o *ListProjectSearchNamesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListProjectSearchNamesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/searches/names][%d] ListProjectSearchNames default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectSearchNamesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/searches/names][%d] ListProjectSearchNames default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectSearchNamesDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListProjectSearchNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
