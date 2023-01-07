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

// UpdateProjectSettingsReader is a Reader for the UpdateProjectSettings structure.
type UpdateProjectSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateProjectSettingsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateProjectSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateProjectSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateProjectSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateProjectSettingsOK creates a UpdateProjectSettingsOK with default headers values
func NewUpdateProjectSettingsOK() *UpdateProjectSettingsOK {
	return &UpdateProjectSettingsOK{}
}

/* UpdateProjectSettingsOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateProjectSettingsOK struct {
	Payload *service_model.V1ProjectSettings
}

// IsSuccess returns true when this update project settings o k response has a 2xx status code
func (o *UpdateProjectSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update project settings o k response has a 3xx status code
func (o *UpdateProjectSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project settings o k response has a 4xx status code
func (o *UpdateProjectSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update project settings o k response has a 5xx status code
func (o *UpdateProjectSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update project settings o k response a status code equal to that given
func (o *UpdateProjectSettingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateProjectSettingsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/settings][%d] updateProjectSettingsOK  %+v", 200, o.Payload)
}

func (o *UpdateProjectSettingsOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/settings][%d] updateProjectSettingsOK  %+v", 200, o.Payload)
}

func (o *UpdateProjectSettingsOK) GetPayload() *service_model.V1ProjectSettings {
	return o.Payload
}

func (o *UpdateProjectSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ProjectSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectSettingsNoContent creates a UpdateProjectSettingsNoContent with default headers values
func NewUpdateProjectSettingsNoContent() *UpdateProjectSettingsNoContent {
	return &UpdateProjectSettingsNoContent{}
}

/* UpdateProjectSettingsNoContent describes a response with status code 204, with default header values.

No content.
*/
type UpdateProjectSettingsNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this update project settings no content response has a 2xx status code
func (o *UpdateProjectSettingsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update project settings no content response has a 3xx status code
func (o *UpdateProjectSettingsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project settings no content response has a 4xx status code
func (o *UpdateProjectSettingsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update project settings no content response has a 5xx status code
func (o *UpdateProjectSettingsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update project settings no content response a status code equal to that given
func (o *UpdateProjectSettingsNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateProjectSettingsNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/settings][%d] updateProjectSettingsNoContent  %+v", 204, o.Payload)
}

func (o *UpdateProjectSettingsNoContent) String() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/settings][%d] updateProjectSettingsNoContent  %+v", 204, o.Payload)
}

func (o *UpdateProjectSettingsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateProjectSettingsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectSettingsForbidden creates a UpdateProjectSettingsForbidden with default headers values
func NewUpdateProjectSettingsForbidden() *UpdateProjectSettingsForbidden {
	return &UpdateProjectSettingsForbidden{}
}

/* UpdateProjectSettingsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type UpdateProjectSettingsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this update project settings forbidden response has a 2xx status code
func (o *UpdateProjectSettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project settings forbidden response has a 3xx status code
func (o *UpdateProjectSettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project settings forbidden response has a 4xx status code
func (o *UpdateProjectSettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project settings forbidden response has a 5xx status code
func (o *UpdateProjectSettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update project settings forbidden response a status code equal to that given
func (o *UpdateProjectSettingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateProjectSettingsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/settings][%d] updateProjectSettingsForbidden  %+v", 403, o.Payload)
}

func (o *UpdateProjectSettingsForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/settings][%d] updateProjectSettingsForbidden  %+v", 403, o.Payload)
}

func (o *UpdateProjectSettingsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateProjectSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectSettingsNotFound creates a UpdateProjectSettingsNotFound with default headers values
func NewUpdateProjectSettingsNotFound() *UpdateProjectSettingsNotFound {
	return &UpdateProjectSettingsNotFound{}
}

/* UpdateProjectSettingsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type UpdateProjectSettingsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this update project settings not found response has a 2xx status code
func (o *UpdateProjectSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project settings not found response has a 3xx status code
func (o *UpdateProjectSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project settings not found response has a 4xx status code
func (o *UpdateProjectSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project settings not found response has a 5xx status code
func (o *UpdateProjectSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update project settings not found response a status code equal to that given
func (o *UpdateProjectSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateProjectSettingsNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/settings][%d] updateProjectSettingsNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProjectSettingsNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/settings][%d] updateProjectSettingsNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProjectSettingsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateProjectSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectSettingsDefault creates a UpdateProjectSettingsDefault with default headers values
func NewUpdateProjectSettingsDefault(code int) *UpdateProjectSettingsDefault {
	return &UpdateProjectSettingsDefault{
		_statusCode: code,
	}
}

/* UpdateProjectSettingsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateProjectSettingsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the update project settings default response
func (o *UpdateProjectSettingsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update project settings default response has a 2xx status code
func (o *UpdateProjectSettingsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update project settings default response has a 3xx status code
func (o *UpdateProjectSettingsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update project settings default response has a 4xx status code
func (o *UpdateProjectSettingsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update project settings default response has a 5xx status code
func (o *UpdateProjectSettingsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update project settings default response a status code equal to that given
func (o *UpdateProjectSettingsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateProjectSettingsDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/settings][%d] UpdateProjectSettings default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateProjectSettingsDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/settings][%d] UpdateProjectSettings default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateProjectSettingsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *UpdateProjectSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
