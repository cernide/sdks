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

package dashboards_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// CreateDashboardReader is a Reader for the CreateDashboard structure.
type CreateDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCreateDashboardNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateDashboardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateDashboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateDashboardOK creates a CreateDashboardOK with default headers values
func NewCreateDashboardOK() *CreateDashboardOK {
	return &CreateDashboardOK{}
}

/* CreateDashboardOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateDashboardOK struct {
	Payload *service_model.V1Dashboard
}

// IsSuccess returns true when this create dashboard o k response has a 2xx status code
func (o *CreateDashboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create dashboard o k response has a 3xx status code
func (o *CreateDashboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create dashboard o k response has a 4xx status code
func (o *CreateDashboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create dashboard o k response has a 5xx status code
func (o *CreateDashboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create dashboard o k response a status code equal to that given
func (o *CreateDashboardOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create dashboard o k response
func (o *CreateDashboardOK) Code() int {
	return 200
}

func (o *CreateDashboardOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/dashboards][%d] createDashboardOK  %+v", 200, o.Payload)
}

func (o *CreateDashboardOK) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/dashboards][%d] createDashboardOK  %+v", 200, o.Payload)
}

func (o *CreateDashboardOK) GetPayload() *service_model.V1Dashboard {
	return o.Payload
}

func (o *CreateDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Dashboard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardNoContent creates a CreateDashboardNoContent with default headers values
func NewCreateDashboardNoContent() *CreateDashboardNoContent {
	return &CreateDashboardNoContent{}
}

/* CreateDashboardNoContent describes a response with status code 204, with default header values.

No content.
*/
type CreateDashboardNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this create dashboard no content response has a 2xx status code
func (o *CreateDashboardNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create dashboard no content response has a 3xx status code
func (o *CreateDashboardNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create dashboard no content response has a 4xx status code
func (o *CreateDashboardNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this create dashboard no content response has a 5xx status code
func (o *CreateDashboardNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this create dashboard no content response a status code equal to that given
func (o *CreateDashboardNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the create dashboard no content response
func (o *CreateDashboardNoContent) Code() int {
	return 204
}

func (o *CreateDashboardNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/dashboards][%d] createDashboardNoContent  %+v", 204, o.Payload)
}

func (o *CreateDashboardNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/dashboards][%d] createDashboardNoContent  %+v", 204, o.Payload)
}

func (o *CreateDashboardNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateDashboardNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardForbidden creates a CreateDashboardForbidden with default headers values
func NewCreateDashboardForbidden() *CreateDashboardForbidden {
	return &CreateDashboardForbidden{}
}

/* CreateDashboardForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type CreateDashboardForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this create dashboard forbidden response has a 2xx status code
func (o *CreateDashboardForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create dashboard forbidden response has a 3xx status code
func (o *CreateDashboardForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create dashboard forbidden response has a 4xx status code
func (o *CreateDashboardForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create dashboard forbidden response has a 5xx status code
func (o *CreateDashboardForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create dashboard forbidden response a status code equal to that given
func (o *CreateDashboardForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create dashboard forbidden response
func (o *CreateDashboardForbidden) Code() int {
	return 403
}

func (o *CreateDashboardForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/dashboards][%d] createDashboardForbidden  %+v", 403, o.Payload)
}

func (o *CreateDashboardForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/dashboards][%d] createDashboardForbidden  %+v", 403, o.Payload)
}

func (o *CreateDashboardForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateDashboardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardNotFound creates a CreateDashboardNotFound with default headers values
func NewCreateDashboardNotFound() *CreateDashboardNotFound {
	return &CreateDashboardNotFound{}
}

/* CreateDashboardNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type CreateDashboardNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this create dashboard not found response has a 2xx status code
func (o *CreateDashboardNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create dashboard not found response has a 3xx status code
func (o *CreateDashboardNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create dashboard not found response has a 4xx status code
func (o *CreateDashboardNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create dashboard not found response has a 5xx status code
func (o *CreateDashboardNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create dashboard not found response a status code equal to that given
func (o *CreateDashboardNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create dashboard not found response
func (o *CreateDashboardNotFound) Code() int {
	return 404
}

func (o *CreateDashboardNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/dashboards][%d] createDashboardNotFound  %+v", 404, o.Payload)
}

func (o *CreateDashboardNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/dashboards][%d] createDashboardNotFound  %+v", 404, o.Payload)
}

func (o *CreateDashboardNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardDefault creates a CreateDashboardDefault with default headers values
func NewCreateDashboardDefault(code int) *CreateDashboardDefault {
	return &CreateDashboardDefault{
		_statusCode: code,
	}
}

/* CreateDashboardDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateDashboardDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this create dashboard default response has a 2xx status code
func (o *CreateDashboardDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create dashboard default response has a 3xx status code
func (o *CreateDashboardDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create dashboard default response has a 4xx status code
func (o *CreateDashboardDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create dashboard default response has a 5xx status code
func (o *CreateDashboardDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create dashboard default response a status code equal to that given
func (o *CreateDashboardDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create dashboard default response
func (o *CreateDashboardDefault) Code() int {
	return o._statusCode
}

func (o *CreateDashboardDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/dashboards][%d] CreateDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDashboardDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/dashboards][%d] CreateDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDashboardDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *CreateDashboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
