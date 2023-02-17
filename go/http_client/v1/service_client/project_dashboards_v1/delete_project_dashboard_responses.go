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

// DeleteProjectDashboardReader is a Reader for the DeleteProjectDashboard structure.
type DeleteProjectDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProjectDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteProjectDashboardNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteProjectDashboardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteProjectDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteProjectDashboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteProjectDashboardOK creates a DeleteProjectDashboardOK with default headers values
func NewDeleteProjectDashboardOK() *DeleteProjectDashboardOK {
	return &DeleteProjectDashboardOK{}
}

/* DeleteProjectDashboardOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteProjectDashboardOK struct {
}

// IsSuccess returns true when this delete project dashboard o k response has a 2xx status code
func (o *DeleteProjectDashboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete project dashboard o k response has a 3xx status code
func (o *DeleteProjectDashboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project dashboard o k response has a 4xx status code
func (o *DeleteProjectDashboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete project dashboard o k response has a 5xx status code
func (o *DeleteProjectDashboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project dashboard o k response a status code equal to that given
func (o *DeleteProjectDashboardOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete project dashboard o k response
func (o *DeleteProjectDashboardOK) Code() int {
	return 200
}

func (o *DeleteProjectDashboardOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{entity}/dashboards/{uuid}][%d] deleteProjectDashboardOK ", 200)
}

func (o *DeleteProjectDashboardOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{entity}/dashboards/{uuid}][%d] deleteProjectDashboardOK ", 200)
}

func (o *DeleteProjectDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectDashboardNoContent creates a DeleteProjectDashboardNoContent with default headers values
func NewDeleteProjectDashboardNoContent() *DeleteProjectDashboardNoContent {
	return &DeleteProjectDashboardNoContent{}
}

/* DeleteProjectDashboardNoContent describes a response with status code 204, with default header values.

No content.
*/
type DeleteProjectDashboardNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this delete project dashboard no content response has a 2xx status code
func (o *DeleteProjectDashboardNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete project dashboard no content response has a 3xx status code
func (o *DeleteProjectDashboardNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project dashboard no content response has a 4xx status code
func (o *DeleteProjectDashboardNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete project dashboard no content response has a 5xx status code
func (o *DeleteProjectDashboardNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project dashboard no content response a status code equal to that given
func (o *DeleteProjectDashboardNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete project dashboard no content response
func (o *DeleteProjectDashboardNoContent) Code() int {
	return 204
}

func (o *DeleteProjectDashboardNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{entity}/dashboards/{uuid}][%d] deleteProjectDashboardNoContent  %+v", 204, o.Payload)
}

func (o *DeleteProjectDashboardNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{entity}/dashboards/{uuid}][%d] deleteProjectDashboardNoContent  %+v", 204, o.Payload)
}

func (o *DeleteProjectDashboardNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteProjectDashboardNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectDashboardForbidden creates a DeleteProjectDashboardForbidden with default headers values
func NewDeleteProjectDashboardForbidden() *DeleteProjectDashboardForbidden {
	return &DeleteProjectDashboardForbidden{}
}

/* DeleteProjectDashboardForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type DeleteProjectDashboardForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this delete project dashboard forbidden response has a 2xx status code
func (o *DeleteProjectDashboardForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete project dashboard forbidden response has a 3xx status code
func (o *DeleteProjectDashboardForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project dashboard forbidden response has a 4xx status code
func (o *DeleteProjectDashboardForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete project dashboard forbidden response has a 5xx status code
func (o *DeleteProjectDashboardForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project dashboard forbidden response a status code equal to that given
func (o *DeleteProjectDashboardForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete project dashboard forbidden response
func (o *DeleteProjectDashboardForbidden) Code() int {
	return 403
}

func (o *DeleteProjectDashboardForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{entity}/dashboards/{uuid}][%d] deleteProjectDashboardForbidden  %+v", 403, o.Payload)
}

func (o *DeleteProjectDashboardForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{entity}/dashboards/{uuid}][%d] deleteProjectDashboardForbidden  %+v", 403, o.Payload)
}

func (o *DeleteProjectDashboardForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteProjectDashboardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectDashboardNotFound creates a DeleteProjectDashboardNotFound with default headers values
func NewDeleteProjectDashboardNotFound() *DeleteProjectDashboardNotFound {
	return &DeleteProjectDashboardNotFound{}
}

/* DeleteProjectDashboardNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type DeleteProjectDashboardNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this delete project dashboard not found response has a 2xx status code
func (o *DeleteProjectDashboardNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete project dashboard not found response has a 3xx status code
func (o *DeleteProjectDashboardNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project dashboard not found response has a 4xx status code
func (o *DeleteProjectDashboardNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete project dashboard not found response has a 5xx status code
func (o *DeleteProjectDashboardNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project dashboard not found response a status code equal to that given
func (o *DeleteProjectDashboardNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete project dashboard not found response
func (o *DeleteProjectDashboardNotFound) Code() int {
	return 404
}

func (o *DeleteProjectDashboardNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{entity}/dashboards/{uuid}][%d] deleteProjectDashboardNotFound  %+v", 404, o.Payload)
}

func (o *DeleteProjectDashboardNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{entity}/dashboards/{uuid}][%d] deleteProjectDashboardNotFound  %+v", 404, o.Payload)
}

func (o *DeleteProjectDashboardNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteProjectDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectDashboardDefault creates a DeleteProjectDashboardDefault with default headers values
func NewDeleteProjectDashboardDefault(code int) *DeleteProjectDashboardDefault {
	return &DeleteProjectDashboardDefault{
		_statusCode: code,
	}
}

/* DeleteProjectDashboardDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteProjectDashboardDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this delete project dashboard default response has a 2xx status code
func (o *DeleteProjectDashboardDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete project dashboard default response has a 3xx status code
func (o *DeleteProjectDashboardDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete project dashboard default response has a 4xx status code
func (o *DeleteProjectDashboardDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete project dashboard default response has a 5xx status code
func (o *DeleteProjectDashboardDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete project dashboard default response a status code equal to that given
func (o *DeleteProjectDashboardDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete project dashboard default response
func (o *DeleteProjectDashboardDefault) Code() int {
	return o._statusCode
}

func (o *DeleteProjectDashboardDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{entity}/dashboards/{uuid}][%d] DeleteProjectDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteProjectDashboardDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{entity}/dashboards/{uuid}][%d] DeleteProjectDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteProjectDashboardDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *DeleteProjectDashboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
