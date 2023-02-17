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

package users_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// UpdateTokenReader is a Reader for the UpdateToken structure.
type UpdateTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateTokenNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateTokenOK creates a UpdateTokenOK with default headers values
func NewUpdateTokenOK() *UpdateTokenOK {
	return &UpdateTokenOK{}
}

/* UpdateTokenOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateTokenOK struct {
	Payload *service_model.V1Token
}

// IsSuccess returns true when this update token o k response has a 2xx status code
func (o *UpdateTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update token o k response has a 3xx status code
func (o *UpdateTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update token o k response has a 4xx status code
func (o *UpdateTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update token o k response has a 5xx status code
func (o *UpdateTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update token o k response a status code equal to that given
func (o *UpdateTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update token o k response
func (o *UpdateTokenOK) Code() int {
	return 200
}

func (o *UpdateTokenOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/users/tokens/{token.uuid}][%d] updateTokenOK  %+v", 200, o.Payload)
}

func (o *UpdateTokenOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/users/tokens/{token.uuid}][%d] updateTokenOK  %+v", 200, o.Payload)
}

func (o *UpdateTokenOK) GetPayload() *service_model.V1Token {
	return o.Payload
}

func (o *UpdateTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTokenNoContent creates a UpdateTokenNoContent with default headers values
func NewUpdateTokenNoContent() *UpdateTokenNoContent {
	return &UpdateTokenNoContent{}
}

/* UpdateTokenNoContent describes a response with status code 204, with default header values.

No content.
*/
type UpdateTokenNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this update token no content response has a 2xx status code
func (o *UpdateTokenNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update token no content response has a 3xx status code
func (o *UpdateTokenNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update token no content response has a 4xx status code
func (o *UpdateTokenNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update token no content response has a 5xx status code
func (o *UpdateTokenNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update token no content response a status code equal to that given
func (o *UpdateTokenNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update token no content response
func (o *UpdateTokenNoContent) Code() int {
	return 204
}

func (o *UpdateTokenNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/users/tokens/{token.uuid}][%d] updateTokenNoContent  %+v", 204, o.Payload)
}

func (o *UpdateTokenNoContent) String() string {
	return fmt.Sprintf("[PUT /api/v1/users/tokens/{token.uuid}][%d] updateTokenNoContent  %+v", 204, o.Payload)
}

func (o *UpdateTokenNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateTokenNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTokenForbidden creates a UpdateTokenForbidden with default headers values
func NewUpdateTokenForbidden() *UpdateTokenForbidden {
	return &UpdateTokenForbidden{}
}

/* UpdateTokenForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type UpdateTokenForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this update token forbidden response has a 2xx status code
func (o *UpdateTokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update token forbidden response has a 3xx status code
func (o *UpdateTokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update token forbidden response has a 4xx status code
func (o *UpdateTokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update token forbidden response has a 5xx status code
func (o *UpdateTokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update token forbidden response a status code equal to that given
func (o *UpdateTokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update token forbidden response
func (o *UpdateTokenForbidden) Code() int {
	return 403
}

func (o *UpdateTokenForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/users/tokens/{token.uuid}][%d] updateTokenForbidden  %+v", 403, o.Payload)
}

func (o *UpdateTokenForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v1/users/tokens/{token.uuid}][%d] updateTokenForbidden  %+v", 403, o.Payload)
}

func (o *UpdateTokenForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTokenNotFound creates a UpdateTokenNotFound with default headers values
func NewUpdateTokenNotFound() *UpdateTokenNotFound {
	return &UpdateTokenNotFound{}
}

/* UpdateTokenNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type UpdateTokenNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this update token not found response has a 2xx status code
func (o *UpdateTokenNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update token not found response has a 3xx status code
func (o *UpdateTokenNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update token not found response has a 4xx status code
func (o *UpdateTokenNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update token not found response has a 5xx status code
func (o *UpdateTokenNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update token not found response a status code equal to that given
func (o *UpdateTokenNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update token not found response
func (o *UpdateTokenNotFound) Code() int {
	return 404
}

func (o *UpdateTokenNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/users/tokens/{token.uuid}][%d] updateTokenNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTokenNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/users/tokens/{token.uuid}][%d] updateTokenNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTokenNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTokenDefault creates a UpdateTokenDefault with default headers values
func NewUpdateTokenDefault(code int) *UpdateTokenDefault {
	return &UpdateTokenDefault{
		_statusCode: code,
	}
}

/* UpdateTokenDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateTokenDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this update token default response has a 2xx status code
func (o *UpdateTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update token default response has a 3xx status code
func (o *UpdateTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update token default response has a 4xx status code
func (o *UpdateTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update token default response has a 5xx status code
func (o *UpdateTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update token default response a status code equal to that given
func (o *UpdateTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update token default response
func (o *UpdateTokenDefault) Code() int {
	return o._statusCode
}

func (o *UpdateTokenDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/users/tokens/{token.uuid}][%d] UpdateToken default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTokenDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/users/tokens/{token.uuid}][%d] UpdateToken default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTokenDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *UpdateTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
