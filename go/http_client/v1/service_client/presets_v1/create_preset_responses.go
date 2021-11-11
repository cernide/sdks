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

package presets_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// CreatePresetReader is a Reader for the CreatePreset structure.
type CreatePresetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePresetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePresetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCreatePresetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreatePresetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreatePresetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreatePresetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePresetOK creates a CreatePresetOK with default headers values
func NewCreatePresetOK() *CreatePresetOK {
	return &CreatePresetOK{}
}

/* CreatePresetOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreatePresetOK struct {
	Payload *service_model.V1Preset
}

func (o *CreatePresetOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/presets][%d] createPresetOK  %+v", 200, o.Payload)
}
func (o *CreatePresetOK) GetPayload() *service_model.V1Preset {
	return o.Payload
}

func (o *CreatePresetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Preset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePresetNoContent creates a CreatePresetNoContent with default headers values
func NewCreatePresetNoContent() *CreatePresetNoContent {
	return &CreatePresetNoContent{}
}

/* CreatePresetNoContent describes a response with status code 204, with default header values.

No content.
*/
type CreatePresetNoContent struct {
	Payload interface{}
}

func (o *CreatePresetNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/presets][%d] createPresetNoContent  %+v", 204, o.Payload)
}
func (o *CreatePresetNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CreatePresetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePresetForbidden creates a CreatePresetForbidden with default headers values
func NewCreatePresetForbidden() *CreatePresetForbidden {
	return &CreatePresetForbidden{}
}

/* CreatePresetForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type CreatePresetForbidden struct {
	Payload interface{}
}

func (o *CreatePresetForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/presets][%d] createPresetForbidden  %+v", 403, o.Payload)
}
func (o *CreatePresetForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreatePresetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePresetNotFound creates a CreatePresetNotFound with default headers values
func NewCreatePresetNotFound() *CreatePresetNotFound {
	return &CreatePresetNotFound{}
}

/* CreatePresetNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type CreatePresetNotFound struct {
	Payload interface{}
}

func (o *CreatePresetNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/presets][%d] createPresetNotFound  %+v", 404, o.Payload)
}
func (o *CreatePresetNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreatePresetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePresetDefault creates a CreatePresetDefault with default headers values
func NewCreatePresetDefault(code int) *CreatePresetDefault {
	return &CreatePresetDefault{
		_statusCode: code,
	}
}

/* CreatePresetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreatePresetDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the create preset default response
func (o *CreatePresetDefault) Code() int {
	return o._statusCode
}

func (o *CreatePresetDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/presets][%d] CreatePreset default  %+v", o._statusCode, o.Payload)
}
func (o *CreatePresetDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *CreatePresetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
