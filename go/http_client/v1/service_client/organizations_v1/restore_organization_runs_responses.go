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

// RestoreOrganizationRunsReader is a Reader for the RestoreOrganizationRuns structure.
type RestoreOrganizationRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestoreOrganizationRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestoreOrganizationRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewRestoreOrganizationRunsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRestoreOrganizationRunsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRestoreOrganizationRunsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRestoreOrganizationRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRestoreOrganizationRunsOK creates a RestoreOrganizationRunsOK with default headers values
func NewRestoreOrganizationRunsOK() *RestoreOrganizationRunsOK {
	return &RestoreOrganizationRunsOK{}
}

/* RestoreOrganizationRunsOK describes a response with status code 200, with default header values.

A successful response.
*/
type RestoreOrganizationRunsOK struct {
}

func (o *RestoreOrganizationRunsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/restore][%d] restoreOrganizationRunsOK ", 200)
}

func (o *RestoreOrganizationRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRestoreOrganizationRunsNoContent creates a RestoreOrganizationRunsNoContent with default headers values
func NewRestoreOrganizationRunsNoContent() *RestoreOrganizationRunsNoContent {
	return &RestoreOrganizationRunsNoContent{}
}

/* RestoreOrganizationRunsNoContent describes a response with status code 204, with default header values.

No content.
*/
type RestoreOrganizationRunsNoContent struct {
	Payload interface{}
}

func (o *RestoreOrganizationRunsNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/restore][%d] restoreOrganizationRunsNoContent  %+v", 204, o.Payload)
}
func (o *RestoreOrganizationRunsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *RestoreOrganizationRunsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreOrganizationRunsForbidden creates a RestoreOrganizationRunsForbidden with default headers values
func NewRestoreOrganizationRunsForbidden() *RestoreOrganizationRunsForbidden {
	return &RestoreOrganizationRunsForbidden{}
}

/* RestoreOrganizationRunsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type RestoreOrganizationRunsForbidden struct {
	Payload interface{}
}

func (o *RestoreOrganizationRunsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/restore][%d] restoreOrganizationRunsForbidden  %+v", 403, o.Payload)
}
func (o *RestoreOrganizationRunsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *RestoreOrganizationRunsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreOrganizationRunsNotFound creates a RestoreOrganizationRunsNotFound with default headers values
func NewRestoreOrganizationRunsNotFound() *RestoreOrganizationRunsNotFound {
	return &RestoreOrganizationRunsNotFound{}
}

/* RestoreOrganizationRunsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type RestoreOrganizationRunsNotFound struct {
	Payload interface{}
}

func (o *RestoreOrganizationRunsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/restore][%d] restoreOrganizationRunsNotFound  %+v", 404, o.Payload)
}
func (o *RestoreOrganizationRunsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *RestoreOrganizationRunsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreOrganizationRunsDefault creates a RestoreOrganizationRunsDefault with default headers values
func NewRestoreOrganizationRunsDefault(code int) *RestoreOrganizationRunsDefault {
	return &RestoreOrganizationRunsDefault{
		_statusCode: code,
	}
}

/* RestoreOrganizationRunsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RestoreOrganizationRunsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the restore organization runs default response
func (o *RestoreOrganizationRunsDefault) Code() int {
	return o._statusCode
}

func (o *RestoreOrganizationRunsDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/runs/restore][%d] RestoreOrganizationRuns default  %+v", o._statusCode, o.Payload)
}
func (o *RestoreOrganizationRunsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *RestoreOrganizationRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}