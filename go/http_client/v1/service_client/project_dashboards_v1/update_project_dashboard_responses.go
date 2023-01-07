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

package project_dashboards_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// UpdateProjectDashboardReader is a Reader for the UpdateProjectDashboard structure.
type UpdateProjectDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateProjectDashboardNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateProjectDashboardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateProjectDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateProjectDashboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateProjectDashboardOK creates a UpdateProjectDashboardOK with default headers values
func NewUpdateProjectDashboardOK() *UpdateProjectDashboardOK {
	return &UpdateProjectDashboardOK{}
}

/* UpdateProjectDashboardOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateProjectDashboardOK struct {
	Payload *service_model.V1Dashboard
}

// IsSuccess returns true when this update project dashboard o k response has a 2xx status code
func (o *UpdateProjectDashboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update project dashboard o k response has a 3xx status code
func (o *UpdateProjectDashboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project dashboard o k response has a 4xx status code
func (o *UpdateProjectDashboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update project dashboard o k response has a 5xx status code
func (o *UpdateProjectDashboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update project dashboard o k response a status code equal to that given
func (o *UpdateProjectDashboardOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateProjectDashboardOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/dashboards/{dashboard.uuid}][%d] updateProjectDashboardOK  %+v", 200, o.Payload)
}

func (o *UpdateProjectDashboardOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/dashboards/{dashboard.uuid}][%d] updateProjectDashboardOK  %+v", 200, o.Payload)
}

func (o *UpdateProjectDashboardOK) GetPayload() *service_model.V1Dashboard {
	return o.Payload
}

func (o *UpdateProjectDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Dashboard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectDashboardNoContent creates a UpdateProjectDashboardNoContent with default headers values
func NewUpdateProjectDashboardNoContent() *UpdateProjectDashboardNoContent {
	return &UpdateProjectDashboardNoContent{}
}

/* UpdateProjectDashboardNoContent describes a response with status code 204, with default header values.

No content.
*/
type UpdateProjectDashboardNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this update project dashboard no content response has a 2xx status code
func (o *UpdateProjectDashboardNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update project dashboard no content response has a 3xx status code
func (o *UpdateProjectDashboardNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project dashboard no content response has a 4xx status code
func (o *UpdateProjectDashboardNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update project dashboard no content response has a 5xx status code
func (o *UpdateProjectDashboardNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update project dashboard no content response a status code equal to that given
func (o *UpdateProjectDashboardNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateProjectDashboardNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/dashboards/{dashboard.uuid}][%d] updateProjectDashboardNoContent  %+v", 204, o.Payload)
}

func (o *UpdateProjectDashboardNoContent) String() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/dashboards/{dashboard.uuid}][%d] updateProjectDashboardNoContent  %+v", 204, o.Payload)
}

func (o *UpdateProjectDashboardNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateProjectDashboardNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectDashboardForbidden creates a UpdateProjectDashboardForbidden with default headers values
func NewUpdateProjectDashboardForbidden() *UpdateProjectDashboardForbidden {
	return &UpdateProjectDashboardForbidden{}
}

/* UpdateProjectDashboardForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type UpdateProjectDashboardForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this update project dashboard forbidden response has a 2xx status code
func (o *UpdateProjectDashboardForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project dashboard forbidden response has a 3xx status code
func (o *UpdateProjectDashboardForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project dashboard forbidden response has a 4xx status code
func (o *UpdateProjectDashboardForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project dashboard forbidden response has a 5xx status code
func (o *UpdateProjectDashboardForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update project dashboard forbidden response a status code equal to that given
func (o *UpdateProjectDashboardForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateProjectDashboardForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/dashboards/{dashboard.uuid}][%d] updateProjectDashboardForbidden  %+v", 403, o.Payload)
}

func (o *UpdateProjectDashboardForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/dashboards/{dashboard.uuid}][%d] updateProjectDashboardForbidden  %+v", 403, o.Payload)
}

func (o *UpdateProjectDashboardForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateProjectDashboardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectDashboardNotFound creates a UpdateProjectDashboardNotFound with default headers values
func NewUpdateProjectDashboardNotFound() *UpdateProjectDashboardNotFound {
	return &UpdateProjectDashboardNotFound{}
}

/* UpdateProjectDashboardNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type UpdateProjectDashboardNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this update project dashboard not found response has a 2xx status code
func (o *UpdateProjectDashboardNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project dashboard not found response has a 3xx status code
func (o *UpdateProjectDashboardNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project dashboard not found response has a 4xx status code
func (o *UpdateProjectDashboardNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project dashboard not found response has a 5xx status code
func (o *UpdateProjectDashboardNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update project dashboard not found response a status code equal to that given
func (o *UpdateProjectDashboardNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateProjectDashboardNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/dashboards/{dashboard.uuid}][%d] updateProjectDashboardNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProjectDashboardNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/dashboards/{dashboard.uuid}][%d] updateProjectDashboardNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProjectDashboardNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateProjectDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectDashboardDefault creates a UpdateProjectDashboardDefault with default headers values
func NewUpdateProjectDashboardDefault(code int) *UpdateProjectDashboardDefault {
	return &UpdateProjectDashboardDefault{
		_statusCode: code,
	}
}

/* UpdateProjectDashboardDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateProjectDashboardDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the update project dashboard default response
func (o *UpdateProjectDashboardDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update project dashboard default response has a 2xx status code
func (o *UpdateProjectDashboardDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update project dashboard default response has a 3xx status code
func (o *UpdateProjectDashboardDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update project dashboard default response has a 4xx status code
func (o *UpdateProjectDashboardDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update project dashboard default response has a 5xx status code
func (o *UpdateProjectDashboardDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update project dashboard default response a status code equal to that given
func (o *UpdateProjectDashboardDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateProjectDashboardDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/dashboards/{dashboard.uuid}][%d] UpdateProjectDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateProjectDashboardDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/dashboards/{dashboard.uuid}][%d] UpdateProjectDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateProjectDashboardDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *UpdateProjectDashboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
