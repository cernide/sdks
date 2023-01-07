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

package teams_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListTeamNamesReader is a Reader for the ListTeamNames structure.
type ListTeamNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTeamNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTeamNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListTeamNamesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListTeamNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListTeamNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListTeamNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListTeamNamesOK creates a ListTeamNamesOK with default headers values
func NewListTeamNamesOK() *ListTeamNamesOK {
	return &ListTeamNamesOK{}
}

/* ListTeamNamesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListTeamNamesOK struct {
	Payload *service_model.V1ListTeamsResponse
}

// IsSuccess returns true when this list team names o k response has a 2xx status code
func (o *ListTeamNamesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list team names o k response has a 3xx status code
func (o *ListTeamNamesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list team names o k response has a 4xx status code
func (o *ListTeamNamesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list team names o k response has a 5xx status code
func (o *ListTeamNamesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list team names o k response a status code equal to that given
func (o *ListTeamNamesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListTeamNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/names][%d] listTeamNamesOK  %+v", 200, o.Payload)
}

func (o *ListTeamNamesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/names][%d] listTeamNamesOK  %+v", 200, o.Payload)
}

func (o *ListTeamNamesOK) GetPayload() *service_model.V1ListTeamsResponse {
	return o.Payload
}

func (o *ListTeamNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListTeamsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamNamesNoContent creates a ListTeamNamesNoContent with default headers values
func NewListTeamNamesNoContent() *ListTeamNamesNoContent {
	return &ListTeamNamesNoContent{}
}

/* ListTeamNamesNoContent describes a response with status code 204, with default header values.

No content.
*/
type ListTeamNamesNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this list team names no content response has a 2xx status code
func (o *ListTeamNamesNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list team names no content response has a 3xx status code
func (o *ListTeamNamesNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list team names no content response has a 4xx status code
func (o *ListTeamNamesNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this list team names no content response has a 5xx status code
func (o *ListTeamNamesNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this list team names no content response a status code equal to that given
func (o *ListTeamNamesNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ListTeamNamesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/names][%d] listTeamNamesNoContent  %+v", 204, o.Payload)
}

func (o *ListTeamNamesNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/names][%d] listTeamNamesNoContent  %+v", 204, o.Payload)
}

func (o *ListTeamNamesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListTeamNamesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamNamesForbidden creates a ListTeamNamesForbidden with default headers values
func NewListTeamNamesForbidden() *ListTeamNamesForbidden {
	return &ListTeamNamesForbidden{}
}

/* ListTeamNamesForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ListTeamNamesForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this list team names forbidden response has a 2xx status code
func (o *ListTeamNamesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list team names forbidden response has a 3xx status code
func (o *ListTeamNamesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list team names forbidden response has a 4xx status code
func (o *ListTeamNamesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list team names forbidden response has a 5xx status code
func (o *ListTeamNamesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list team names forbidden response a status code equal to that given
func (o *ListTeamNamesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListTeamNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/names][%d] listTeamNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListTeamNamesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/names][%d] listTeamNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListTeamNamesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListTeamNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamNamesNotFound creates a ListTeamNamesNotFound with default headers values
func NewListTeamNamesNotFound() *ListTeamNamesNotFound {
	return &ListTeamNamesNotFound{}
}

/* ListTeamNamesNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ListTeamNamesNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this list team names not found response has a 2xx status code
func (o *ListTeamNamesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list team names not found response has a 3xx status code
func (o *ListTeamNamesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list team names not found response has a 4xx status code
func (o *ListTeamNamesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list team names not found response has a 5xx status code
func (o *ListTeamNamesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list team names not found response a status code equal to that given
func (o *ListTeamNamesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ListTeamNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/names][%d] listTeamNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListTeamNamesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/names][%d] listTeamNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListTeamNamesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListTeamNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamNamesDefault creates a ListTeamNamesDefault with default headers values
func NewListTeamNamesDefault(code int) *ListTeamNamesDefault {
	return &ListTeamNamesDefault{
		_statusCode: code,
	}
}

/* ListTeamNamesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListTeamNamesDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the list team names default response
func (o *ListTeamNamesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list team names default response has a 2xx status code
func (o *ListTeamNamesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list team names default response has a 3xx status code
func (o *ListTeamNamesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list team names default response has a 4xx status code
func (o *ListTeamNamesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list team names default response has a 5xx status code
func (o *ListTeamNamesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list team names default response a status code equal to that given
func (o *ListTeamNamesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListTeamNamesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/names][%d] ListTeamNames default  %+v", o._statusCode, o.Payload)
}

func (o *ListTeamNamesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams/names][%d] ListTeamNames default  %+v", o._statusCode, o.Payload)
}

func (o *ListTeamNamesDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListTeamNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
