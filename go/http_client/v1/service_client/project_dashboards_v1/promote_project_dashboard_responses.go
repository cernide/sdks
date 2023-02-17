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

// PromoteProjectDashboardReader is a Reader for the PromoteProjectDashboard structure.
type PromoteProjectDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PromoteProjectDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPromoteProjectDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPromoteProjectDashboardNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPromoteProjectDashboardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPromoteProjectDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPromoteProjectDashboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPromoteProjectDashboardOK creates a PromoteProjectDashboardOK with default headers values
func NewPromoteProjectDashboardOK() *PromoteProjectDashboardOK {
	return &PromoteProjectDashboardOK{}
}

/* PromoteProjectDashboardOK describes a response with status code 200, with default header values.

A successful response.
*/
type PromoteProjectDashboardOK struct {
}

// IsSuccess returns true when this promote project dashboard o k response has a 2xx status code
func (o *PromoteProjectDashboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this promote project dashboard o k response has a 3xx status code
func (o *PromoteProjectDashboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this promote project dashboard o k response has a 4xx status code
func (o *PromoteProjectDashboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this promote project dashboard o k response has a 5xx status code
func (o *PromoteProjectDashboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this promote project dashboard o k response a status code equal to that given
func (o *PromoteProjectDashboardOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the promote project dashboard o k response
func (o *PromoteProjectDashboardOK) Code() int {
	return 200
}

func (o *PromoteProjectDashboardOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/dashboards/{uuid}/promote][%d] promoteProjectDashboardOK ", 200)
}

func (o *PromoteProjectDashboardOK) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/dashboards/{uuid}/promote][%d] promoteProjectDashboardOK ", 200)
}

func (o *PromoteProjectDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPromoteProjectDashboardNoContent creates a PromoteProjectDashboardNoContent with default headers values
func NewPromoteProjectDashboardNoContent() *PromoteProjectDashboardNoContent {
	return &PromoteProjectDashboardNoContent{}
}

/* PromoteProjectDashboardNoContent describes a response with status code 204, with default header values.

No content.
*/
type PromoteProjectDashboardNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this promote project dashboard no content response has a 2xx status code
func (o *PromoteProjectDashboardNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this promote project dashboard no content response has a 3xx status code
func (o *PromoteProjectDashboardNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this promote project dashboard no content response has a 4xx status code
func (o *PromoteProjectDashboardNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this promote project dashboard no content response has a 5xx status code
func (o *PromoteProjectDashboardNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this promote project dashboard no content response a status code equal to that given
func (o *PromoteProjectDashboardNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the promote project dashboard no content response
func (o *PromoteProjectDashboardNoContent) Code() int {
	return 204
}

func (o *PromoteProjectDashboardNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/dashboards/{uuid}/promote][%d] promoteProjectDashboardNoContent  %+v", 204, o.Payload)
}

func (o *PromoteProjectDashboardNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/dashboards/{uuid}/promote][%d] promoteProjectDashboardNoContent  %+v", 204, o.Payload)
}

func (o *PromoteProjectDashboardNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PromoteProjectDashboardNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPromoteProjectDashboardForbidden creates a PromoteProjectDashboardForbidden with default headers values
func NewPromoteProjectDashboardForbidden() *PromoteProjectDashboardForbidden {
	return &PromoteProjectDashboardForbidden{}
}

/* PromoteProjectDashboardForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type PromoteProjectDashboardForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this promote project dashboard forbidden response has a 2xx status code
func (o *PromoteProjectDashboardForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this promote project dashboard forbidden response has a 3xx status code
func (o *PromoteProjectDashboardForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this promote project dashboard forbidden response has a 4xx status code
func (o *PromoteProjectDashboardForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this promote project dashboard forbidden response has a 5xx status code
func (o *PromoteProjectDashboardForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this promote project dashboard forbidden response a status code equal to that given
func (o *PromoteProjectDashboardForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the promote project dashboard forbidden response
func (o *PromoteProjectDashboardForbidden) Code() int {
	return 403
}

func (o *PromoteProjectDashboardForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/dashboards/{uuid}/promote][%d] promoteProjectDashboardForbidden  %+v", 403, o.Payload)
}

func (o *PromoteProjectDashboardForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/dashboards/{uuid}/promote][%d] promoteProjectDashboardForbidden  %+v", 403, o.Payload)
}

func (o *PromoteProjectDashboardForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PromoteProjectDashboardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPromoteProjectDashboardNotFound creates a PromoteProjectDashboardNotFound with default headers values
func NewPromoteProjectDashboardNotFound() *PromoteProjectDashboardNotFound {
	return &PromoteProjectDashboardNotFound{}
}

/* PromoteProjectDashboardNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type PromoteProjectDashboardNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this promote project dashboard not found response has a 2xx status code
func (o *PromoteProjectDashboardNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this promote project dashboard not found response has a 3xx status code
func (o *PromoteProjectDashboardNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this promote project dashboard not found response has a 4xx status code
func (o *PromoteProjectDashboardNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this promote project dashboard not found response has a 5xx status code
func (o *PromoteProjectDashboardNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this promote project dashboard not found response a status code equal to that given
func (o *PromoteProjectDashboardNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the promote project dashboard not found response
func (o *PromoteProjectDashboardNotFound) Code() int {
	return 404
}

func (o *PromoteProjectDashboardNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/dashboards/{uuid}/promote][%d] promoteProjectDashboardNotFound  %+v", 404, o.Payload)
}

func (o *PromoteProjectDashboardNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/dashboards/{uuid}/promote][%d] promoteProjectDashboardNotFound  %+v", 404, o.Payload)
}

func (o *PromoteProjectDashboardNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PromoteProjectDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPromoteProjectDashboardDefault creates a PromoteProjectDashboardDefault with default headers values
func NewPromoteProjectDashboardDefault(code int) *PromoteProjectDashboardDefault {
	return &PromoteProjectDashboardDefault{
		_statusCode: code,
	}
}

/* PromoteProjectDashboardDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PromoteProjectDashboardDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this promote project dashboard default response has a 2xx status code
func (o *PromoteProjectDashboardDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this promote project dashboard default response has a 3xx status code
func (o *PromoteProjectDashboardDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this promote project dashboard default response has a 4xx status code
func (o *PromoteProjectDashboardDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this promote project dashboard default response has a 5xx status code
func (o *PromoteProjectDashboardDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this promote project dashboard default response a status code equal to that given
func (o *PromoteProjectDashboardDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the promote project dashboard default response
func (o *PromoteProjectDashboardDefault) Code() int {
	return o._statusCode
}

func (o *PromoteProjectDashboardDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/dashboards/{uuid}/promote][%d] PromoteProjectDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *PromoteProjectDashboardDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/dashboards/{uuid}/promote][%d] PromoteProjectDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *PromoteProjectDashboardDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *PromoteProjectDashboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
