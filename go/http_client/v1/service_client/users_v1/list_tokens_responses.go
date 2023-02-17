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

// ListTokensReader is a Reader for the ListTokens structure.
type ListTokensReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTokensReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTokensOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListTokensNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListTokensForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListTokensNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListTokensDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListTokensOK creates a ListTokensOK with default headers values
func NewListTokensOK() *ListTokensOK {
	return &ListTokensOK{}
}

/* ListTokensOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListTokensOK struct {
	Payload *service_model.V1ListTokenResponse
}

// IsSuccess returns true when this list tokens o k response has a 2xx status code
func (o *ListTokensOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list tokens o k response has a 3xx status code
func (o *ListTokensOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list tokens o k response has a 4xx status code
func (o *ListTokensOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list tokens o k response has a 5xx status code
func (o *ListTokensOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list tokens o k response a status code equal to that given
func (o *ListTokensOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list tokens o k response
func (o *ListTokensOK) Code() int {
	return 200
}

func (o *ListTokensOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/users/tokens][%d] listTokensOK  %+v", 200, o.Payload)
}

func (o *ListTokensOK) String() string {
	return fmt.Sprintf("[GET /api/v1/users/tokens][%d] listTokensOK  %+v", 200, o.Payload)
}

func (o *ListTokensOK) GetPayload() *service_model.V1ListTokenResponse {
	return o.Payload
}

func (o *ListTokensOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTokensNoContent creates a ListTokensNoContent with default headers values
func NewListTokensNoContent() *ListTokensNoContent {
	return &ListTokensNoContent{}
}

/* ListTokensNoContent describes a response with status code 204, with default header values.

No content.
*/
type ListTokensNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this list tokens no content response has a 2xx status code
func (o *ListTokensNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list tokens no content response has a 3xx status code
func (o *ListTokensNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list tokens no content response has a 4xx status code
func (o *ListTokensNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this list tokens no content response has a 5xx status code
func (o *ListTokensNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this list tokens no content response a status code equal to that given
func (o *ListTokensNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the list tokens no content response
func (o *ListTokensNoContent) Code() int {
	return 204
}

func (o *ListTokensNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/users/tokens][%d] listTokensNoContent  %+v", 204, o.Payload)
}

func (o *ListTokensNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/users/tokens][%d] listTokensNoContent  %+v", 204, o.Payload)
}

func (o *ListTokensNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListTokensNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTokensForbidden creates a ListTokensForbidden with default headers values
func NewListTokensForbidden() *ListTokensForbidden {
	return &ListTokensForbidden{}
}

/* ListTokensForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ListTokensForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this list tokens forbidden response has a 2xx status code
func (o *ListTokensForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list tokens forbidden response has a 3xx status code
func (o *ListTokensForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list tokens forbidden response has a 4xx status code
func (o *ListTokensForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list tokens forbidden response has a 5xx status code
func (o *ListTokensForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list tokens forbidden response a status code equal to that given
func (o *ListTokensForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list tokens forbidden response
func (o *ListTokensForbidden) Code() int {
	return 403
}

func (o *ListTokensForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/users/tokens][%d] listTokensForbidden  %+v", 403, o.Payload)
}

func (o *ListTokensForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/users/tokens][%d] listTokensForbidden  %+v", 403, o.Payload)
}

func (o *ListTokensForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListTokensForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTokensNotFound creates a ListTokensNotFound with default headers values
func NewListTokensNotFound() *ListTokensNotFound {
	return &ListTokensNotFound{}
}

/* ListTokensNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ListTokensNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this list tokens not found response has a 2xx status code
func (o *ListTokensNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list tokens not found response has a 3xx status code
func (o *ListTokensNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list tokens not found response has a 4xx status code
func (o *ListTokensNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list tokens not found response has a 5xx status code
func (o *ListTokensNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list tokens not found response a status code equal to that given
func (o *ListTokensNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list tokens not found response
func (o *ListTokensNotFound) Code() int {
	return 404
}

func (o *ListTokensNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/users/tokens][%d] listTokensNotFound  %+v", 404, o.Payload)
}

func (o *ListTokensNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/users/tokens][%d] listTokensNotFound  %+v", 404, o.Payload)
}

func (o *ListTokensNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListTokensNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTokensDefault creates a ListTokensDefault with default headers values
func NewListTokensDefault(code int) *ListTokensDefault {
	return &ListTokensDefault{
		_statusCode: code,
	}
}

/* ListTokensDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListTokensDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this list tokens default response has a 2xx status code
func (o *ListTokensDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list tokens default response has a 3xx status code
func (o *ListTokensDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list tokens default response has a 4xx status code
func (o *ListTokensDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list tokens default response has a 5xx status code
func (o *ListTokensDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list tokens default response a status code equal to that given
func (o *ListTokensDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list tokens default response
func (o *ListTokensDefault) Code() int {
	return o._statusCode
}

func (o *ListTokensDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/users/tokens][%d] ListTokens default  %+v", o._statusCode, o.Payload)
}

func (o *ListTokensDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/users/tokens][%d] ListTokens default  %+v", o._statusCode, o.Payload)
}

func (o *ListTokensDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListTokensDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
