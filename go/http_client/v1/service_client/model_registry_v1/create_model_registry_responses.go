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

package model_registry_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// CreateModelRegistryReader is a Reader for the CreateModelRegistry structure.
type CreateModelRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateModelRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateModelRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCreateModelRegistryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateModelRegistryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateModelRegistryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateModelRegistryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateModelRegistryOK creates a CreateModelRegistryOK with default headers values
func NewCreateModelRegistryOK() *CreateModelRegistryOK {
	return &CreateModelRegistryOK{}
}

/* CreateModelRegistryOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateModelRegistryOK struct {
	Payload *service_model.V1ModelRegistry
}

func (o *CreateModelRegistryOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/registry/create][%d] createModelRegistryOK  %+v", 200, o.Payload)
}
func (o *CreateModelRegistryOK) GetPayload() *service_model.V1ModelRegistry {
	return o.Payload
}

func (o *CreateModelRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ModelRegistry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateModelRegistryNoContent creates a CreateModelRegistryNoContent with default headers values
func NewCreateModelRegistryNoContent() *CreateModelRegistryNoContent {
	return &CreateModelRegistryNoContent{}
}

/* CreateModelRegistryNoContent describes a response with status code 204, with default header values.

No content.
*/
type CreateModelRegistryNoContent struct {
	Payload interface{}
}

func (o *CreateModelRegistryNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/registry/create][%d] createModelRegistryNoContent  %+v", 204, o.Payload)
}
func (o *CreateModelRegistryNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateModelRegistryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateModelRegistryForbidden creates a CreateModelRegistryForbidden with default headers values
func NewCreateModelRegistryForbidden() *CreateModelRegistryForbidden {
	return &CreateModelRegistryForbidden{}
}

/* CreateModelRegistryForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type CreateModelRegistryForbidden struct {
	Payload interface{}
}

func (o *CreateModelRegistryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/registry/create][%d] createModelRegistryForbidden  %+v", 403, o.Payload)
}
func (o *CreateModelRegistryForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateModelRegistryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateModelRegistryNotFound creates a CreateModelRegistryNotFound with default headers values
func NewCreateModelRegistryNotFound() *CreateModelRegistryNotFound {
	return &CreateModelRegistryNotFound{}
}

/* CreateModelRegistryNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type CreateModelRegistryNotFound struct {
	Payload interface{}
}

func (o *CreateModelRegistryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/registry/create][%d] createModelRegistryNotFound  %+v", 404, o.Payload)
}
func (o *CreateModelRegistryNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateModelRegistryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateModelRegistryDefault creates a CreateModelRegistryDefault with default headers values
func NewCreateModelRegistryDefault(code int) *CreateModelRegistryDefault {
	return &CreateModelRegistryDefault{
		_statusCode: code,
	}
}

/* CreateModelRegistryDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateModelRegistryDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the create model registry default response
func (o *CreateModelRegistryDefault) Code() int {
	return o._statusCode
}

func (o *CreateModelRegistryDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/registry/create][%d] CreateModelRegistry default  %+v", o._statusCode, o.Payload)
}
func (o *CreateModelRegistryDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *CreateModelRegistryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
