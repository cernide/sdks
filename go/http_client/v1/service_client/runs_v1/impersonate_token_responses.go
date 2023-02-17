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

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// ImpersonateTokenReader is a Reader for the ImpersonateToken structure.
type ImpersonateTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImpersonateTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImpersonateTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewImpersonateTokenNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewImpersonateTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImpersonateTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewImpersonateTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImpersonateTokenOK creates a ImpersonateTokenOK with default headers values
func NewImpersonateTokenOK() *ImpersonateTokenOK {
	return &ImpersonateTokenOK{}
}

/* ImpersonateTokenOK describes a response with status code 200, with default header values.

A successful response.
*/
type ImpersonateTokenOK struct {
	Payload *service_model.V1Auth
}

// IsSuccess returns true when this impersonate token o k response has a 2xx status code
func (o *ImpersonateTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this impersonate token o k response has a 3xx status code
func (o *ImpersonateTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this impersonate token o k response has a 4xx status code
func (o *ImpersonateTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this impersonate token o k response has a 5xx status code
func (o *ImpersonateTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this impersonate token o k response a status code equal to that given
func (o *ImpersonateTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the impersonate token o k response
func (o *ImpersonateTokenOK) Code() int {
	return 200
}

func (o *ImpersonateTokenOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/impersonate][%d] impersonateTokenOK  %+v", 200, o.Payload)
}

func (o *ImpersonateTokenOK) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/impersonate][%d] impersonateTokenOK  %+v", 200, o.Payload)
}

func (o *ImpersonateTokenOK) GetPayload() *service_model.V1Auth {
	return o.Payload
}

func (o *ImpersonateTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Auth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImpersonateTokenNoContent creates a ImpersonateTokenNoContent with default headers values
func NewImpersonateTokenNoContent() *ImpersonateTokenNoContent {
	return &ImpersonateTokenNoContent{}
}

/* ImpersonateTokenNoContent describes a response with status code 204, with default header values.

No content.
*/
type ImpersonateTokenNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this impersonate token no content response has a 2xx status code
func (o *ImpersonateTokenNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this impersonate token no content response has a 3xx status code
func (o *ImpersonateTokenNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this impersonate token no content response has a 4xx status code
func (o *ImpersonateTokenNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this impersonate token no content response has a 5xx status code
func (o *ImpersonateTokenNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this impersonate token no content response a status code equal to that given
func (o *ImpersonateTokenNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the impersonate token no content response
func (o *ImpersonateTokenNoContent) Code() int {
	return 204
}

func (o *ImpersonateTokenNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/impersonate][%d] impersonateTokenNoContent  %+v", 204, o.Payload)
}

func (o *ImpersonateTokenNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/impersonate][%d] impersonateTokenNoContent  %+v", 204, o.Payload)
}

func (o *ImpersonateTokenNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ImpersonateTokenNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImpersonateTokenForbidden creates a ImpersonateTokenForbidden with default headers values
func NewImpersonateTokenForbidden() *ImpersonateTokenForbidden {
	return &ImpersonateTokenForbidden{}
}

/* ImpersonateTokenForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ImpersonateTokenForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this impersonate token forbidden response has a 2xx status code
func (o *ImpersonateTokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this impersonate token forbidden response has a 3xx status code
func (o *ImpersonateTokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this impersonate token forbidden response has a 4xx status code
func (o *ImpersonateTokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this impersonate token forbidden response has a 5xx status code
func (o *ImpersonateTokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this impersonate token forbidden response a status code equal to that given
func (o *ImpersonateTokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the impersonate token forbidden response
func (o *ImpersonateTokenForbidden) Code() int {
	return 403
}

func (o *ImpersonateTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/impersonate][%d] impersonateTokenForbidden  %+v", 403, o.Payload)
}

func (o *ImpersonateTokenForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/impersonate][%d] impersonateTokenForbidden  %+v", 403, o.Payload)
}

func (o *ImpersonateTokenForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ImpersonateTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImpersonateTokenNotFound creates a ImpersonateTokenNotFound with default headers values
func NewImpersonateTokenNotFound() *ImpersonateTokenNotFound {
	return &ImpersonateTokenNotFound{}
}

/* ImpersonateTokenNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ImpersonateTokenNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this impersonate token not found response has a 2xx status code
func (o *ImpersonateTokenNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this impersonate token not found response has a 3xx status code
func (o *ImpersonateTokenNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this impersonate token not found response has a 4xx status code
func (o *ImpersonateTokenNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this impersonate token not found response has a 5xx status code
func (o *ImpersonateTokenNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this impersonate token not found response a status code equal to that given
func (o *ImpersonateTokenNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the impersonate token not found response
func (o *ImpersonateTokenNotFound) Code() int {
	return 404
}

func (o *ImpersonateTokenNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/impersonate][%d] impersonateTokenNotFound  %+v", 404, o.Payload)
}

func (o *ImpersonateTokenNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/impersonate][%d] impersonateTokenNotFound  %+v", 404, o.Payload)
}

func (o *ImpersonateTokenNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ImpersonateTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImpersonateTokenDefault creates a ImpersonateTokenDefault with default headers values
func NewImpersonateTokenDefault(code int) *ImpersonateTokenDefault {
	return &ImpersonateTokenDefault{
		_statusCode: code,
	}
}

/* ImpersonateTokenDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ImpersonateTokenDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this impersonate token default response has a 2xx status code
func (o *ImpersonateTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this impersonate token default response has a 3xx status code
func (o *ImpersonateTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this impersonate token default response has a 4xx status code
func (o *ImpersonateTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this impersonate token default response has a 5xx status code
func (o *ImpersonateTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this impersonate token default response a status code equal to that given
func (o *ImpersonateTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the impersonate token default response
func (o *ImpersonateTokenDefault) Code() int {
	return o._statusCode
}

func (o *ImpersonateTokenDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/impersonate][%d] ImpersonateToken default  %+v", o._statusCode, o.Payload)
}

func (o *ImpersonateTokenDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/impersonate][%d] ImpersonateToken default  %+v", o._statusCode, o.Payload)
}

func (o *ImpersonateTokenDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ImpersonateTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
