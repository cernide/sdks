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

package project_searches_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// DeleteProjectSearchReader is a Reader for the DeleteProjectSearch structure.
type DeleteProjectSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProjectSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteProjectSearchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteProjectSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteProjectSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteProjectSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteProjectSearchOK creates a DeleteProjectSearchOK with default headers values
func NewDeleteProjectSearchOK() *DeleteProjectSearchOK {
	return &DeleteProjectSearchOK{}
}

/* DeleteProjectSearchOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteProjectSearchOK struct {
}

// IsSuccess returns true when this delete project search o k response has a 2xx status code
func (o *DeleteProjectSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete project search o k response has a 3xx status code
func (o *DeleteProjectSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project search o k response has a 4xx status code
func (o *DeleteProjectSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete project search o k response has a 5xx status code
func (o *DeleteProjectSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project search o k response a status code equal to that given
func (o *DeleteProjectSearchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete project search o k response
func (o *DeleteProjectSearchOK) Code() int {
	return 200
}

func (o *DeleteProjectSearchOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{entity}/searches/{uuid}][%d] deleteProjectSearchOK ", 200)
}

func (o *DeleteProjectSearchOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{entity}/searches/{uuid}][%d] deleteProjectSearchOK ", 200)
}

func (o *DeleteProjectSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectSearchNoContent creates a DeleteProjectSearchNoContent with default headers values
func NewDeleteProjectSearchNoContent() *DeleteProjectSearchNoContent {
	return &DeleteProjectSearchNoContent{}
}

/* DeleteProjectSearchNoContent describes a response with status code 204, with default header values.

No content.
*/
type DeleteProjectSearchNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this delete project search no content response has a 2xx status code
func (o *DeleteProjectSearchNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete project search no content response has a 3xx status code
func (o *DeleteProjectSearchNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project search no content response has a 4xx status code
func (o *DeleteProjectSearchNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete project search no content response has a 5xx status code
func (o *DeleteProjectSearchNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project search no content response a status code equal to that given
func (o *DeleteProjectSearchNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete project search no content response
func (o *DeleteProjectSearchNoContent) Code() int {
	return 204
}

func (o *DeleteProjectSearchNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{entity}/searches/{uuid}][%d] deleteProjectSearchNoContent  %+v", 204, o.Payload)
}

func (o *DeleteProjectSearchNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{entity}/searches/{uuid}][%d] deleteProjectSearchNoContent  %+v", 204, o.Payload)
}

func (o *DeleteProjectSearchNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteProjectSearchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectSearchForbidden creates a DeleteProjectSearchForbidden with default headers values
func NewDeleteProjectSearchForbidden() *DeleteProjectSearchForbidden {
	return &DeleteProjectSearchForbidden{}
}

/* DeleteProjectSearchForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type DeleteProjectSearchForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this delete project search forbidden response has a 2xx status code
func (o *DeleteProjectSearchForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete project search forbidden response has a 3xx status code
func (o *DeleteProjectSearchForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project search forbidden response has a 4xx status code
func (o *DeleteProjectSearchForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete project search forbidden response has a 5xx status code
func (o *DeleteProjectSearchForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project search forbidden response a status code equal to that given
func (o *DeleteProjectSearchForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete project search forbidden response
func (o *DeleteProjectSearchForbidden) Code() int {
	return 403
}

func (o *DeleteProjectSearchForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{entity}/searches/{uuid}][%d] deleteProjectSearchForbidden  %+v", 403, o.Payload)
}

func (o *DeleteProjectSearchForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{entity}/searches/{uuid}][%d] deleteProjectSearchForbidden  %+v", 403, o.Payload)
}

func (o *DeleteProjectSearchForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteProjectSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectSearchNotFound creates a DeleteProjectSearchNotFound with default headers values
func NewDeleteProjectSearchNotFound() *DeleteProjectSearchNotFound {
	return &DeleteProjectSearchNotFound{}
}

/* DeleteProjectSearchNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type DeleteProjectSearchNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this delete project search not found response has a 2xx status code
func (o *DeleteProjectSearchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete project search not found response has a 3xx status code
func (o *DeleteProjectSearchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project search not found response has a 4xx status code
func (o *DeleteProjectSearchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete project search not found response has a 5xx status code
func (o *DeleteProjectSearchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project search not found response a status code equal to that given
func (o *DeleteProjectSearchNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete project search not found response
func (o *DeleteProjectSearchNotFound) Code() int {
	return 404
}

func (o *DeleteProjectSearchNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{entity}/searches/{uuid}][%d] deleteProjectSearchNotFound  %+v", 404, o.Payload)
}

func (o *DeleteProjectSearchNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{entity}/searches/{uuid}][%d] deleteProjectSearchNotFound  %+v", 404, o.Payload)
}

func (o *DeleteProjectSearchNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteProjectSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectSearchDefault creates a DeleteProjectSearchDefault with default headers values
func NewDeleteProjectSearchDefault(code int) *DeleteProjectSearchDefault {
	return &DeleteProjectSearchDefault{
		_statusCode: code,
	}
}

/* DeleteProjectSearchDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteProjectSearchDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this delete project search default response has a 2xx status code
func (o *DeleteProjectSearchDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete project search default response has a 3xx status code
func (o *DeleteProjectSearchDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete project search default response has a 4xx status code
func (o *DeleteProjectSearchDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete project search default response has a 5xx status code
func (o *DeleteProjectSearchDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete project search default response a status code equal to that given
func (o *DeleteProjectSearchDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete project search default response
func (o *DeleteProjectSearchDefault) Code() int {
	return o._statusCode
}

func (o *DeleteProjectSearchDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{entity}/searches/{uuid}][%d] DeleteProjectSearch default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteProjectSearchDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{entity}/searches/{uuid}][%d] DeleteProjectSearch default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteProjectSearchDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *DeleteProjectSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
