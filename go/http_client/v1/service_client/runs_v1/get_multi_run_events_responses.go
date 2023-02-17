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

// GetMultiRunEventsReader is a Reader for the GetMultiRunEvents structure.
type GetMultiRunEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMultiRunEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMultiRunEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetMultiRunEventsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetMultiRunEventsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMultiRunEventsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetMultiRunEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMultiRunEventsOK creates a GetMultiRunEventsOK with default headers values
func NewGetMultiRunEventsOK() *GetMultiRunEventsOK {
	return &GetMultiRunEventsOK{}
}

/* GetMultiRunEventsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetMultiRunEventsOK struct {
	Payload *service_model.V1EventsResponse
}

// IsSuccess returns true when this get multi run events o k response has a 2xx status code
func (o *GetMultiRunEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get multi run events o k response has a 3xx status code
func (o *GetMultiRunEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get multi run events o k response has a 4xx status code
func (o *GetMultiRunEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get multi run events o k response has a 5xx status code
func (o *GetMultiRunEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get multi run events o k response a status code equal to that given
func (o *GetMultiRunEventsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get multi run events o k response
func (o *GetMultiRunEventsOK) Code() int {
	return 200
}

func (o *GetMultiRunEventsOK) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/multi/events/{kind}][%d] getMultiRunEventsOK  %+v", 200, o.Payload)
}

func (o *GetMultiRunEventsOK) String() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/multi/events/{kind}][%d] getMultiRunEventsOK  %+v", 200, o.Payload)
}

func (o *GetMultiRunEventsOK) GetPayload() *service_model.V1EventsResponse {
	return o.Payload
}

func (o *GetMultiRunEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1EventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMultiRunEventsNoContent creates a GetMultiRunEventsNoContent with default headers values
func NewGetMultiRunEventsNoContent() *GetMultiRunEventsNoContent {
	return &GetMultiRunEventsNoContent{}
}

/* GetMultiRunEventsNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetMultiRunEventsNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get multi run events no content response has a 2xx status code
func (o *GetMultiRunEventsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get multi run events no content response has a 3xx status code
func (o *GetMultiRunEventsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get multi run events no content response has a 4xx status code
func (o *GetMultiRunEventsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get multi run events no content response has a 5xx status code
func (o *GetMultiRunEventsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get multi run events no content response a status code equal to that given
func (o *GetMultiRunEventsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get multi run events no content response
func (o *GetMultiRunEventsNoContent) Code() int {
	return 204
}

func (o *GetMultiRunEventsNoContent) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/multi/events/{kind}][%d] getMultiRunEventsNoContent  %+v", 204, o.Payload)
}

func (o *GetMultiRunEventsNoContent) String() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/multi/events/{kind}][%d] getMultiRunEventsNoContent  %+v", 204, o.Payload)
}

func (o *GetMultiRunEventsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetMultiRunEventsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMultiRunEventsForbidden creates a GetMultiRunEventsForbidden with default headers values
func NewGetMultiRunEventsForbidden() *GetMultiRunEventsForbidden {
	return &GetMultiRunEventsForbidden{}
}

/* GetMultiRunEventsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetMultiRunEventsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get multi run events forbidden response has a 2xx status code
func (o *GetMultiRunEventsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get multi run events forbidden response has a 3xx status code
func (o *GetMultiRunEventsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get multi run events forbidden response has a 4xx status code
func (o *GetMultiRunEventsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get multi run events forbidden response has a 5xx status code
func (o *GetMultiRunEventsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get multi run events forbidden response a status code equal to that given
func (o *GetMultiRunEventsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get multi run events forbidden response
func (o *GetMultiRunEventsForbidden) Code() int {
	return 403
}

func (o *GetMultiRunEventsForbidden) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/multi/events/{kind}][%d] getMultiRunEventsForbidden  %+v", 403, o.Payload)
}

func (o *GetMultiRunEventsForbidden) String() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/multi/events/{kind}][%d] getMultiRunEventsForbidden  %+v", 403, o.Payload)
}

func (o *GetMultiRunEventsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetMultiRunEventsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMultiRunEventsNotFound creates a GetMultiRunEventsNotFound with default headers values
func NewGetMultiRunEventsNotFound() *GetMultiRunEventsNotFound {
	return &GetMultiRunEventsNotFound{}
}

/* GetMultiRunEventsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetMultiRunEventsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get multi run events not found response has a 2xx status code
func (o *GetMultiRunEventsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get multi run events not found response has a 3xx status code
func (o *GetMultiRunEventsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get multi run events not found response has a 4xx status code
func (o *GetMultiRunEventsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get multi run events not found response has a 5xx status code
func (o *GetMultiRunEventsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get multi run events not found response a status code equal to that given
func (o *GetMultiRunEventsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get multi run events not found response
func (o *GetMultiRunEventsNotFound) Code() int {
	return 404
}

func (o *GetMultiRunEventsNotFound) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/multi/events/{kind}][%d] getMultiRunEventsNotFound  %+v", 404, o.Payload)
}

func (o *GetMultiRunEventsNotFound) String() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/multi/events/{kind}][%d] getMultiRunEventsNotFound  %+v", 404, o.Payload)
}

func (o *GetMultiRunEventsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetMultiRunEventsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMultiRunEventsDefault creates a GetMultiRunEventsDefault with default headers values
func NewGetMultiRunEventsDefault(code int) *GetMultiRunEventsDefault {
	return &GetMultiRunEventsDefault{
		_statusCode: code,
	}
}

/* GetMultiRunEventsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetMultiRunEventsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this get multi run events default response has a 2xx status code
func (o *GetMultiRunEventsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get multi run events default response has a 3xx status code
func (o *GetMultiRunEventsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get multi run events default response has a 4xx status code
func (o *GetMultiRunEventsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get multi run events default response has a 5xx status code
func (o *GetMultiRunEventsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get multi run events default response a status code equal to that given
func (o *GetMultiRunEventsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get multi run events default response
func (o *GetMultiRunEventsDefault) Code() int {
	return o._statusCode
}

func (o *GetMultiRunEventsDefault) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/multi/events/{kind}][%d] GetMultiRunEvents default  %+v", o._statusCode, o.Payload)
}

func (o *GetMultiRunEventsDefault) String() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/multi/events/{kind}][%d] GetMultiRunEvents default  %+v", o._statusCode, o.Payload)
}

func (o *GetMultiRunEventsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetMultiRunEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
