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

package projects_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListArchivedProjectsReader is a Reader for the ListArchivedProjects structure.
type ListArchivedProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListArchivedProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListArchivedProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListArchivedProjectsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListArchivedProjectsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListArchivedProjectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListArchivedProjectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListArchivedProjectsOK creates a ListArchivedProjectsOK with default headers values
func NewListArchivedProjectsOK() *ListArchivedProjectsOK {
	return &ListArchivedProjectsOK{}
}

/* ListArchivedProjectsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListArchivedProjectsOK struct {
	Payload *service_model.V1ListProjectsResponse
}

func (o *ListArchivedProjectsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/projects][%d] listArchivedProjectsOK  %+v", 200, o.Payload)
}
func (o *ListArchivedProjectsOK) GetPayload() *service_model.V1ListProjectsResponse {
	return o.Payload
}

func (o *ListArchivedProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListProjectsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArchivedProjectsNoContent creates a ListArchivedProjectsNoContent with default headers values
func NewListArchivedProjectsNoContent() *ListArchivedProjectsNoContent {
	return &ListArchivedProjectsNoContent{}
}

/* ListArchivedProjectsNoContent describes a response with status code 204, with default header values.

No content.
*/
type ListArchivedProjectsNoContent struct {
	Payload interface{}
}

func (o *ListArchivedProjectsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/projects][%d] listArchivedProjectsNoContent  %+v", 204, o.Payload)
}
func (o *ListArchivedProjectsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListArchivedProjectsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArchivedProjectsForbidden creates a ListArchivedProjectsForbidden with default headers values
func NewListArchivedProjectsForbidden() *ListArchivedProjectsForbidden {
	return &ListArchivedProjectsForbidden{}
}

/* ListArchivedProjectsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ListArchivedProjectsForbidden struct {
	Payload interface{}
}

func (o *ListArchivedProjectsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/projects][%d] listArchivedProjectsForbidden  %+v", 403, o.Payload)
}
func (o *ListArchivedProjectsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListArchivedProjectsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArchivedProjectsNotFound creates a ListArchivedProjectsNotFound with default headers values
func NewListArchivedProjectsNotFound() *ListArchivedProjectsNotFound {
	return &ListArchivedProjectsNotFound{}
}

/* ListArchivedProjectsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ListArchivedProjectsNotFound struct {
	Payload interface{}
}

func (o *ListArchivedProjectsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/projects][%d] listArchivedProjectsNotFound  %+v", 404, o.Payload)
}
func (o *ListArchivedProjectsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListArchivedProjectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArchivedProjectsDefault creates a ListArchivedProjectsDefault with default headers values
func NewListArchivedProjectsDefault(code int) *ListArchivedProjectsDefault {
	return &ListArchivedProjectsDefault{
		_statusCode: code,
	}
}

/* ListArchivedProjectsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListArchivedProjectsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the list archived projects default response
func (o *ListArchivedProjectsDefault) Code() int {
	return o._statusCode
}

func (o *ListArchivedProjectsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/archives/{user}/projects][%d] ListArchivedProjects default  %+v", o._statusCode, o.Payload)
}
func (o *ListArchivedProjectsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListArchivedProjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
