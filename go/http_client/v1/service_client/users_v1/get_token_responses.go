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

// GetTokenReader is a Reader for the GetToken structure.
type GetTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetTokenNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTokenOK creates a GetTokenOK with default headers values
func NewGetTokenOK() *GetTokenOK {
	return &GetTokenOK{}
}

/* GetTokenOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetTokenOK struct {
	Payload *service_model.V1Token
}

// IsSuccess returns true when this get token o k response has a 2xx status code
func (o *GetTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get token o k response has a 3xx status code
func (o *GetTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get token o k response has a 4xx status code
func (o *GetTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get token o k response has a 5xx status code
func (o *GetTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get token o k response a status code equal to that given
func (o *GetTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get token o k response
func (o *GetTokenOK) Code() int {
	return 200
}

func (o *GetTokenOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/users/tokens/{uuid}][%d] getTokenOK  %+v", 200, o.Payload)
}

func (o *GetTokenOK) String() string {
	return fmt.Sprintf("[GET /api/v1/users/tokens/{uuid}][%d] getTokenOK  %+v", 200, o.Payload)
}

func (o *GetTokenOK) GetPayload() *service_model.V1Token {
	return o.Payload
}

func (o *GetTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokenNoContent creates a GetTokenNoContent with default headers values
func NewGetTokenNoContent() *GetTokenNoContent {
	return &GetTokenNoContent{}
}

/* GetTokenNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetTokenNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get token no content response has a 2xx status code
func (o *GetTokenNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get token no content response has a 3xx status code
func (o *GetTokenNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get token no content response has a 4xx status code
func (o *GetTokenNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get token no content response has a 5xx status code
func (o *GetTokenNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get token no content response a status code equal to that given
func (o *GetTokenNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get token no content response
func (o *GetTokenNoContent) Code() int {
	return 204
}

func (o *GetTokenNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/users/tokens/{uuid}][%d] getTokenNoContent  %+v", 204, o.Payload)
}

func (o *GetTokenNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/users/tokens/{uuid}][%d] getTokenNoContent  %+v", 204, o.Payload)
}

func (o *GetTokenNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetTokenNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokenForbidden creates a GetTokenForbidden with default headers values
func NewGetTokenForbidden() *GetTokenForbidden {
	return &GetTokenForbidden{}
}

/* GetTokenForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetTokenForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get token forbidden response has a 2xx status code
func (o *GetTokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get token forbidden response has a 3xx status code
func (o *GetTokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get token forbidden response has a 4xx status code
func (o *GetTokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get token forbidden response has a 5xx status code
func (o *GetTokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get token forbidden response a status code equal to that given
func (o *GetTokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get token forbidden response
func (o *GetTokenForbidden) Code() int {
	return 403
}

func (o *GetTokenForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/users/tokens/{uuid}][%d] getTokenForbidden  %+v", 403, o.Payload)
}

func (o *GetTokenForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/users/tokens/{uuid}][%d] getTokenForbidden  %+v", 403, o.Payload)
}

func (o *GetTokenForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokenNotFound creates a GetTokenNotFound with default headers values
func NewGetTokenNotFound() *GetTokenNotFound {
	return &GetTokenNotFound{}
}

/* GetTokenNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetTokenNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get token not found response has a 2xx status code
func (o *GetTokenNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get token not found response has a 3xx status code
func (o *GetTokenNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get token not found response has a 4xx status code
func (o *GetTokenNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get token not found response has a 5xx status code
func (o *GetTokenNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get token not found response a status code equal to that given
func (o *GetTokenNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get token not found response
func (o *GetTokenNotFound) Code() int {
	return 404
}

func (o *GetTokenNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/users/tokens/{uuid}][%d] getTokenNotFound  %+v", 404, o.Payload)
}

func (o *GetTokenNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/users/tokens/{uuid}][%d] getTokenNotFound  %+v", 404, o.Payload)
}

func (o *GetTokenNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokenDefault creates a GetTokenDefault with default headers values
func NewGetTokenDefault(code int) *GetTokenDefault {
	return &GetTokenDefault{
		_statusCode: code,
	}
}

/* GetTokenDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetTokenDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this get token default response has a 2xx status code
func (o *GetTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get token default response has a 3xx status code
func (o *GetTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get token default response has a 4xx status code
func (o *GetTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get token default response has a 5xx status code
func (o *GetTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get token default response a status code equal to that given
func (o *GetTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get token default response
func (o *GetTokenDefault) Code() int {
	return o._statusCode
}

func (o *GetTokenDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/users/tokens/{uuid}][%d] GetToken default  %+v", o._statusCode, o.Payload)
}

func (o *GetTokenDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/users/tokens/{uuid}][%d] GetToken default  %+v", o._statusCode, o.Payload)
}

func (o *GetTokenDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
