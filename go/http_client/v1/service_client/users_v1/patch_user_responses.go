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

// PatchUserReader is a Reader for the PatchUser structure.
type PatchUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchUserOK creates a PatchUserOK with default headers values
func NewPatchUserOK() *PatchUserOK {
	return &PatchUserOK{}
}

/* PatchUserOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchUserOK struct {
	Payload *service_model.V1User
}

// IsSuccess returns true when this patch user o k response has a 2xx status code
func (o *PatchUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch user o k response has a 3xx status code
func (o *PatchUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch user o k response has a 4xx status code
func (o *PatchUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch user o k response has a 5xx status code
func (o *PatchUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch user o k response a status code equal to that given
func (o *PatchUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchUserOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/users][%d] patchUserOK  %+v", 200, o.Payload)
}

func (o *PatchUserOK) String() string {
	return fmt.Sprintf("[PATCH /api/v1/users][%d] patchUserOK  %+v", 200, o.Payload)
}

func (o *PatchUserOK) GetPayload() *service_model.V1User {
	return o.Payload
}

func (o *PatchUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserNoContent creates a PatchUserNoContent with default headers values
func NewPatchUserNoContent() *PatchUserNoContent {
	return &PatchUserNoContent{}
}

/* PatchUserNoContent describes a response with status code 204, with default header values.

No content.
*/
type PatchUserNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this patch user no content response has a 2xx status code
func (o *PatchUserNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch user no content response has a 3xx status code
func (o *PatchUserNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch user no content response has a 4xx status code
func (o *PatchUserNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch user no content response has a 5xx status code
func (o *PatchUserNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this patch user no content response a status code equal to that given
func (o *PatchUserNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *PatchUserNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/users][%d] patchUserNoContent  %+v", 204, o.Payload)
}

func (o *PatchUserNoContent) String() string {
	return fmt.Sprintf("[PATCH /api/v1/users][%d] patchUserNoContent  %+v", 204, o.Payload)
}

func (o *PatchUserNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserForbidden creates a PatchUserForbidden with default headers values
func NewPatchUserForbidden() *PatchUserForbidden {
	return &PatchUserForbidden{}
}

/* PatchUserForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type PatchUserForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this patch user forbidden response has a 2xx status code
func (o *PatchUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch user forbidden response has a 3xx status code
func (o *PatchUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch user forbidden response has a 4xx status code
func (o *PatchUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch user forbidden response has a 5xx status code
func (o *PatchUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch user forbidden response a status code equal to that given
func (o *PatchUserForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchUserForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/users][%d] patchUserForbidden  %+v", 403, o.Payload)
}

func (o *PatchUserForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v1/users][%d] patchUserForbidden  %+v", 403, o.Payload)
}

func (o *PatchUserForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserNotFound creates a PatchUserNotFound with default headers values
func NewPatchUserNotFound() *PatchUserNotFound {
	return &PatchUserNotFound{}
}

/* PatchUserNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type PatchUserNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this patch user not found response has a 2xx status code
func (o *PatchUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch user not found response has a 3xx status code
func (o *PatchUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch user not found response has a 4xx status code
func (o *PatchUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch user not found response has a 5xx status code
func (o *PatchUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch user not found response a status code equal to that given
func (o *PatchUserNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchUserNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/users][%d] patchUserNotFound  %+v", 404, o.Payload)
}

func (o *PatchUserNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v1/users][%d] patchUserNotFound  %+v", 404, o.Payload)
}

func (o *PatchUserNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserDefault creates a PatchUserDefault with default headers values
func NewPatchUserDefault(code int) *PatchUserDefault {
	return &PatchUserDefault{
		_statusCode: code,
	}
}

/* PatchUserDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PatchUserDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the patch user default response
func (o *PatchUserDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this patch user default response has a 2xx status code
func (o *PatchUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch user default response has a 3xx status code
func (o *PatchUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch user default response has a 4xx status code
func (o *PatchUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch user default response has a 5xx status code
func (o *PatchUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch user default response a status code equal to that given
func (o *PatchUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PatchUserDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/users][%d] PatchUser default  %+v", o._statusCode, o.Payload)
}

func (o *PatchUserDefault) String() string {
	return fmt.Sprintf("[PATCH /api/v1/users][%d] PatchUser default  %+v", o._statusCode, o.Payload)
}

func (o *PatchUserDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *PatchUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
