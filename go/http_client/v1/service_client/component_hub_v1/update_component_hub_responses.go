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

package component_hub_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// UpdateComponentHubReader is a Reader for the UpdateComponentHub structure.
type UpdateComponentHubReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateComponentHubReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateComponentHubOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateComponentHubNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateComponentHubForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateComponentHubNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateComponentHubDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateComponentHubOK creates a UpdateComponentHubOK with default headers values
func NewUpdateComponentHubOK() *UpdateComponentHubOK {
	return &UpdateComponentHubOK{}
}

/* UpdateComponentHubOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateComponentHubOK struct {
	Payload *service_model.V1ComponentHub
}

func (o *UpdateComponentHubOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/hub/{component.name}][%d] updateComponentHubOK  %+v", 200, o.Payload)
}
func (o *UpdateComponentHubOK) GetPayload() *service_model.V1ComponentHub {
	return o.Payload
}

func (o *UpdateComponentHubOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ComponentHub)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateComponentHubNoContent creates a UpdateComponentHubNoContent with default headers values
func NewUpdateComponentHubNoContent() *UpdateComponentHubNoContent {
	return &UpdateComponentHubNoContent{}
}

/* UpdateComponentHubNoContent describes a response with status code 204, with default header values.

No content.
*/
type UpdateComponentHubNoContent struct {
	Payload interface{}
}

func (o *UpdateComponentHubNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/hub/{component.name}][%d] updateComponentHubNoContent  %+v", 204, o.Payload)
}
func (o *UpdateComponentHubNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateComponentHubNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateComponentHubForbidden creates a UpdateComponentHubForbidden with default headers values
func NewUpdateComponentHubForbidden() *UpdateComponentHubForbidden {
	return &UpdateComponentHubForbidden{}
}

/* UpdateComponentHubForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type UpdateComponentHubForbidden struct {
	Payload interface{}
}

func (o *UpdateComponentHubForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/hub/{component.name}][%d] updateComponentHubForbidden  %+v", 403, o.Payload)
}
func (o *UpdateComponentHubForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateComponentHubForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateComponentHubNotFound creates a UpdateComponentHubNotFound with default headers values
func NewUpdateComponentHubNotFound() *UpdateComponentHubNotFound {
	return &UpdateComponentHubNotFound{}
}

/* UpdateComponentHubNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type UpdateComponentHubNotFound struct {
	Payload interface{}
}

func (o *UpdateComponentHubNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/hub/{component.name}][%d] updateComponentHubNotFound  %+v", 404, o.Payload)
}
func (o *UpdateComponentHubNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateComponentHubNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateComponentHubDefault creates a UpdateComponentHubDefault with default headers values
func NewUpdateComponentHubDefault(code int) *UpdateComponentHubDefault {
	return &UpdateComponentHubDefault{
		_statusCode: code,
	}
}

/* UpdateComponentHubDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateComponentHubDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the update component hub default response
func (o *UpdateComponentHubDefault) Code() int {
	return o._statusCode
}

func (o *UpdateComponentHubDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/hub/{component.name}][%d] UpdateComponentHub default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateComponentHubDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *UpdateComponentHubDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
