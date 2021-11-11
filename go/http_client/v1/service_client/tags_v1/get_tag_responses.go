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

package tags_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// GetTagReader is a Reader for the GetTag structure.
type GetTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetTagNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetTagForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetTagDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTagOK creates a GetTagOK with default headers values
func NewGetTagOK() *GetTagOK {
	return &GetTagOK{}
}

/* GetTagOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetTagOK struct {
	Payload *service_model.V1Tag
}

func (o *GetTagOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/tags/{uuid}][%d] getTagOK  %+v", 200, o.Payload)
}
func (o *GetTagOK) GetPayload() *service_model.V1Tag {
	return o.Payload
}

func (o *GetTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTagNoContent creates a GetTagNoContent with default headers values
func NewGetTagNoContent() *GetTagNoContent {
	return &GetTagNoContent{}
}

/* GetTagNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetTagNoContent struct {
	Payload interface{}
}

func (o *GetTagNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/tags/{uuid}][%d] getTagNoContent  %+v", 204, o.Payload)
}
func (o *GetTagNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetTagNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTagForbidden creates a GetTagForbidden with default headers values
func NewGetTagForbidden() *GetTagForbidden {
	return &GetTagForbidden{}
}

/* GetTagForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetTagForbidden struct {
	Payload interface{}
}

func (o *GetTagForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/tags/{uuid}][%d] getTagForbidden  %+v", 403, o.Payload)
}
func (o *GetTagForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetTagForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTagNotFound creates a GetTagNotFound with default headers values
func NewGetTagNotFound() *GetTagNotFound {
	return &GetTagNotFound{}
}

/* GetTagNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetTagNotFound struct {
	Payload interface{}
}

func (o *GetTagNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/tags/{uuid}][%d] getTagNotFound  %+v", 404, o.Payload)
}
func (o *GetTagNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTagDefault creates a GetTagDefault with default headers values
func NewGetTagDefault(code int) *GetTagDefault {
	return &GetTagDefault{
		_statusCode: code,
	}
}

/* GetTagDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetTagDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the get tag default response
func (o *GetTagDefault) Code() int {
	return o._statusCode
}

func (o *GetTagDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/tags/{uuid}][%d] GetTag default  %+v", o._statusCode, o.Payload)
}
func (o *GetTagDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetTagDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
