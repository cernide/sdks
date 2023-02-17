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

package service_accounts_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// UpdateServiceAccountTokenReader is a Reader for the UpdateServiceAccountToken structure.
type UpdateServiceAccountTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateServiceAccountTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateServiceAccountTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateServiceAccountTokenNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateServiceAccountTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateServiceAccountTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateServiceAccountTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateServiceAccountTokenOK creates a UpdateServiceAccountTokenOK with default headers values
func NewUpdateServiceAccountTokenOK() *UpdateServiceAccountTokenOK {
	return &UpdateServiceAccountTokenOK{}
}

/* UpdateServiceAccountTokenOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateServiceAccountTokenOK struct {
	Payload *service_model.V1Token
}

// IsSuccess returns true when this update service account token o k response has a 2xx status code
func (o *UpdateServiceAccountTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update service account token o k response has a 3xx status code
func (o *UpdateServiceAccountTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service account token o k response has a 4xx status code
func (o *UpdateServiceAccountTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update service account token o k response has a 5xx status code
func (o *UpdateServiceAccountTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update service account token o k response a status code equal to that given
func (o *UpdateServiceAccountTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update service account token o k response
func (o *UpdateServiceAccountTokenOK) Code() int {
	return 200
}

func (o *UpdateServiceAccountTokenOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/sa/{entity}/tokens/{token.uuid}][%d] updateServiceAccountTokenOK  %+v", 200, o.Payload)
}

func (o *UpdateServiceAccountTokenOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/sa/{entity}/tokens/{token.uuid}][%d] updateServiceAccountTokenOK  %+v", 200, o.Payload)
}

func (o *UpdateServiceAccountTokenOK) GetPayload() *service_model.V1Token {
	return o.Payload
}

func (o *UpdateServiceAccountTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceAccountTokenNoContent creates a UpdateServiceAccountTokenNoContent with default headers values
func NewUpdateServiceAccountTokenNoContent() *UpdateServiceAccountTokenNoContent {
	return &UpdateServiceAccountTokenNoContent{}
}

/* UpdateServiceAccountTokenNoContent describes a response with status code 204, with default header values.

No content.
*/
type UpdateServiceAccountTokenNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this update service account token no content response has a 2xx status code
func (o *UpdateServiceAccountTokenNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update service account token no content response has a 3xx status code
func (o *UpdateServiceAccountTokenNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service account token no content response has a 4xx status code
func (o *UpdateServiceAccountTokenNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update service account token no content response has a 5xx status code
func (o *UpdateServiceAccountTokenNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update service account token no content response a status code equal to that given
func (o *UpdateServiceAccountTokenNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update service account token no content response
func (o *UpdateServiceAccountTokenNoContent) Code() int {
	return 204
}

func (o *UpdateServiceAccountTokenNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/sa/{entity}/tokens/{token.uuid}][%d] updateServiceAccountTokenNoContent  %+v", 204, o.Payload)
}

func (o *UpdateServiceAccountTokenNoContent) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/sa/{entity}/tokens/{token.uuid}][%d] updateServiceAccountTokenNoContent  %+v", 204, o.Payload)
}

func (o *UpdateServiceAccountTokenNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateServiceAccountTokenNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceAccountTokenForbidden creates a UpdateServiceAccountTokenForbidden with default headers values
func NewUpdateServiceAccountTokenForbidden() *UpdateServiceAccountTokenForbidden {
	return &UpdateServiceAccountTokenForbidden{}
}

/* UpdateServiceAccountTokenForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type UpdateServiceAccountTokenForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this update service account token forbidden response has a 2xx status code
func (o *UpdateServiceAccountTokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update service account token forbidden response has a 3xx status code
func (o *UpdateServiceAccountTokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service account token forbidden response has a 4xx status code
func (o *UpdateServiceAccountTokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update service account token forbidden response has a 5xx status code
func (o *UpdateServiceAccountTokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update service account token forbidden response a status code equal to that given
func (o *UpdateServiceAccountTokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update service account token forbidden response
func (o *UpdateServiceAccountTokenForbidden) Code() int {
	return 403
}

func (o *UpdateServiceAccountTokenForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/sa/{entity}/tokens/{token.uuid}][%d] updateServiceAccountTokenForbidden  %+v", 403, o.Payload)
}

func (o *UpdateServiceAccountTokenForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/sa/{entity}/tokens/{token.uuid}][%d] updateServiceAccountTokenForbidden  %+v", 403, o.Payload)
}

func (o *UpdateServiceAccountTokenForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateServiceAccountTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceAccountTokenNotFound creates a UpdateServiceAccountTokenNotFound with default headers values
func NewUpdateServiceAccountTokenNotFound() *UpdateServiceAccountTokenNotFound {
	return &UpdateServiceAccountTokenNotFound{}
}

/* UpdateServiceAccountTokenNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type UpdateServiceAccountTokenNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this update service account token not found response has a 2xx status code
func (o *UpdateServiceAccountTokenNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update service account token not found response has a 3xx status code
func (o *UpdateServiceAccountTokenNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service account token not found response has a 4xx status code
func (o *UpdateServiceAccountTokenNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update service account token not found response has a 5xx status code
func (o *UpdateServiceAccountTokenNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update service account token not found response a status code equal to that given
func (o *UpdateServiceAccountTokenNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update service account token not found response
func (o *UpdateServiceAccountTokenNotFound) Code() int {
	return 404
}

func (o *UpdateServiceAccountTokenNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/sa/{entity}/tokens/{token.uuid}][%d] updateServiceAccountTokenNotFound  %+v", 404, o.Payload)
}

func (o *UpdateServiceAccountTokenNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/sa/{entity}/tokens/{token.uuid}][%d] updateServiceAccountTokenNotFound  %+v", 404, o.Payload)
}

func (o *UpdateServiceAccountTokenNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateServiceAccountTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceAccountTokenDefault creates a UpdateServiceAccountTokenDefault with default headers values
func NewUpdateServiceAccountTokenDefault(code int) *UpdateServiceAccountTokenDefault {
	return &UpdateServiceAccountTokenDefault{
		_statusCode: code,
	}
}

/* UpdateServiceAccountTokenDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateServiceAccountTokenDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this update service account token default response has a 2xx status code
func (o *UpdateServiceAccountTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update service account token default response has a 3xx status code
func (o *UpdateServiceAccountTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update service account token default response has a 4xx status code
func (o *UpdateServiceAccountTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update service account token default response has a 5xx status code
func (o *UpdateServiceAccountTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update service account token default response a status code equal to that given
func (o *UpdateServiceAccountTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update service account token default response
func (o *UpdateServiceAccountTokenDefault) Code() int {
	return o._statusCode
}

func (o *UpdateServiceAccountTokenDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/sa/{entity}/tokens/{token.uuid}][%d] UpdateServiceAccountToken default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateServiceAccountTokenDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/sa/{entity}/tokens/{token.uuid}][%d] UpdateServiceAccountToken default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateServiceAccountTokenDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *UpdateServiceAccountTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
