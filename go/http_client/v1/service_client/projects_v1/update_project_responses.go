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

// UpdateProjectReader is a Reader for the UpdateProject structure.
type UpdateProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateProjectNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateProjectOK creates a UpdateProjectOK with default headers values
func NewUpdateProjectOK() *UpdateProjectOK {
	return &UpdateProjectOK{}
}

/* UpdateProjectOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateProjectOK struct {
	Payload *service_model.V1Project
}

func (o *UpdateProjectOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project.name}][%d] updateProjectOK  %+v", 200, o.Payload)
}
func (o *UpdateProjectOK) GetPayload() *service_model.V1Project {
	return o.Payload
}

func (o *UpdateProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectNoContent creates a UpdateProjectNoContent with default headers values
func NewUpdateProjectNoContent() *UpdateProjectNoContent {
	return &UpdateProjectNoContent{}
}

/* UpdateProjectNoContent describes a response with status code 204, with default header values.

No content.
*/
type UpdateProjectNoContent struct {
	Payload interface{}
}

func (o *UpdateProjectNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project.name}][%d] updateProjectNoContent  %+v", 204, o.Payload)
}
func (o *UpdateProjectNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateProjectNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectForbidden creates a UpdateProjectForbidden with default headers values
func NewUpdateProjectForbidden() *UpdateProjectForbidden {
	return &UpdateProjectForbidden{}
}

/* UpdateProjectForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type UpdateProjectForbidden struct {
	Payload interface{}
}

func (o *UpdateProjectForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project.name}][%d] updateProjectForbidden  %+v", 403, o.Payload)
}
func (o *UpdateProjectForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectNotFound creates a UpdateProjectNotFound with default headers values
func NewUpdateProjectNotFound() *UpdateProjectNotFound {
	return &UpdateProjectNotFound{}
}

/* UpdateProjectNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type UpdateProjectNotFound struct {
	Payload interface{}
}

func (o *UpdateProjectNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project.name}][%d] updateProjectNotFound  %+v", 404, o.Payload)
}
func (o *UpdateProjectNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectDefault creates a UpdateProjectDefault with default headers values
func NewUpdateProjectDefault(code int) *UpdateProjectDefault {
	return &UpdateProjectDefault{
		_statusCode: code,
	}
}

/* UpdateProjectDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateProjectDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the update project default response
func (o *UpdateProjectDefault) Code() int {
	return o._statusCode
}

func (o *UpdateProjectDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project.name}][%d] UpdateProject default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateProjectDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *UpdateProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
