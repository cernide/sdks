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

// CopyRunReader is a Reader for the CopyRun structure.
type CopyRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CopyRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCopyRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCopyRunNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCopyRunForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCopyRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCopyRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCopyRunOK creates a CopyRunOK with default headers values
func NewCopyRunOK() *CopyRunOK {
	return &CopyRunOK{}
}

/* CopyRunOK describes a response with status code 200, with default header values.

A successful response.
*/
type CopyRunOK struct {
	Payload *service_model.V1Run
}

// IsSuccess returns true when this copy run o k response has a 2xx status code
func (o *CopyRunOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this copy run o k response has a 3xx status code
func (o *CopyRunOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy run o k response has a 4xx status code
func (o *CopyRunOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this copy run o k response has a 5xx status code
func (o *CopyRunOK) IsServerError() bool {
	return false
}

// IsCode returns true when this copy run o k response a status code equal to that given
func (o *CopyRunOK) IsCode(code int) bool {
	return code == 200
}

func (o *CopyRunOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run.uuid}/copy][%d] copyRunOK  %+v", 200, o.Payload)
}

func (o *CopyRunOK) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run.uuid}/copy][%d] copyRunOK  %+v", 200, o.Payload)
}

func (o *CopyRunOK) GetPayload() *service_model.V1Run {
	return o.Payload
}

func (o *CopyRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Run)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCopyRunNoContent creates a CopyRunNoContent with default headers values
func NewCopyRunNoContent() *CopyRunNoContent {
	return &CopyRunNoContent{}
}

/* CopyRunNoContent describes a response with status code 204, with default header values.

No content.
*/
type CopyRunNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this copy run no content response has a 2xx status code
func (o *CopyRunNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this copy run no content response has a 3xx status code
func (o *CopyRunNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy run no content response has a 4xx status code
func (o *CopyRunNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this copy run no content response has a 5xx status code
func (o *CopyRunNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this copy run no content response a status code equal to that given
func (o *CopyRunNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *CopyRunNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run.uuid}/copy][%d] copyRunNoContent  %+v", 204, o.Payload)
}

func (o *CopyRunNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run.uuid}/copy][%d] copyRunNoContent  %+v", 204, o.Payload)
}

func (o *CopyRunNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CopyRunNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCopyRunForbidden creates a CopyRunForbidden with default headers values
func NewCopyRunForbidden() *CopyRunForbidden {
	return &CopyRunForbidden{}
}

/* CopyRunForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type CopyRunForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this copy run forbidden response has a 2xx status code
func (o *CopyRunForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this copy run forbidden response has a 3xx status code
func (o *CopyRunForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy run forbidden response has a 4xx status code
func (o *CopyRunForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this copy run forbidden response has a 5xx status code
func (o *CopyRunForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this copy run forbidden response a status code equal to that given
func (o *CopyRunForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CopyRunForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run.uuid}/copy][%d] copyRunForbidden  %+v", 403, o.Payload)
}

func (o *CopyRunForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run.uuid}/copy][%d] copyRunForbidden  %+v", 403, o.Payload)
}

func (o *CopyRunForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CopyRunForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCopyRunNotFound creates a CopyRunNotFound with default headers values
func NewCopyRunNotFound() *CopyRunNotFound {
	return &CopyRunNotFound{}
}

/* CopyRunNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type CopyRunNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this copy run not found response has a 2xx status code
func (o *CopyRunNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this copy run not found response has a 3xx status code
func (o *CopyRunNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy run not found response has a 4xx status code
func (o *CopyRunNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this copy run not found response has a 5xx status code
func (o *CopyRunNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this copy run not found response a status code equal to that given
func (o *CopyRunNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CopyRunNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run.uuid}/copy][%d] copyRunNotFound  %+v", 404, o.Payload)
}

func (o *CopyRunNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run.uuid}/copy][%d] copyRunNotFound  %+v", 404, o.Payload)
}

func (o *CopyRunNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CopyRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCopyRunDefault creates a CopyRunDefault with default headers values
func NewCopyRunDefault(code int) *CopyRunDefault {
	return &CopyRunDefault{
		_statusCode: code,
	}
}

/* CopyRunDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CopyRunDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the copy run default response
func (o *CopyRunDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this copy run default response has a 2xx status code
func (o *CopyRunDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this copy run default response has a 3xx status code
func (o *CopyRunDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this copy run default response has a 4xx status code
func (o *CopyRunDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this copy run default response has a 5xx status code
func (o *CopyRunDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this copy run default response a status code equal to that given
func (o *CopyRunDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CopyRunDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run.uuid}/copy][%d] CopyRun default  %+v", o._statusCode, o.Payload)
}

func (o *CopyRunDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{run.uuid}/copy][%d] CopyRun default  %+v", o._statusCode, o.Payload)
}

func (o *CopyRunDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *CopyRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
