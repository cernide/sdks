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

// GetRunEventsReader is a Reader for the GetRunEvents structure.
type GetRunEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetRunEventsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRunEventsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRunEventsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetRunEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRunEventsOK creates a GetRunEventsOK with default headers values
func NewGetRunEventsOK() *GetRunEventsOK {
	return &GetRunEventsOK{}
}

/* GetRunEventsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetRunEventsOK struct {
	Payload *service_model.V1EventsResponse
}

// IsSuccess returns true when this get run events o k response has a 2xx status code
func (o *GetRunEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get run events o k response has a 3xx status code
func (o *GetRunEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run events o k response has a 4xx status code
func (o *GetRunEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get run events o k response has a 5xx status code
func (o *GetRunEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get run events o k response a status code equal to that given
func (o *GetRunEventsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get run events o k response
func (o *GetRunEventsOK) Code() int {
	return 200
}

func (o *GetRunEventsOK) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/events/{kind}][%d] getRunEventsOK  %+v", 200, o.Payload)
}

func (o *GetRunEventsOK) String() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/events/{kind}][%d] getRunEventsOK  %+v", 200, o.Payload)
}

func (o *GetRunEventsOK) GetPayload() *service_model.V1EventsResponse {
	return o.Payload
}

func (o *GetRunEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1EventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunEventsNoContent creates a GetRunEventsNoContent with default headers values
func NewGetRunEventsNoContent() *GetRunEventsNoContent {
	return &GetRunEventsNoContent{}
}

/* GetRunEventsNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetRunEventsNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get run events no content response has a 2xx status code
func (o *GetRunEventsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get run events no content response has a 3xx status code
func (o *GetRunEventsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run events no content response has a 4xx status code
func (o *GetRunEventsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get run events no content response has a 5xx status code
func (o *GetRunEventsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get run events no content response a status code equal to that given
func (o *GetRunEventsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get run events no content response
func (o *GetRunEventsNoContent) Code() int {
	return 204
}

func (o *GetRunEventsNoContent) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/events/{kind}][%d] getRunEventsNoContent  %+v", 204, o.Payload)
}

func (o *GetRunEventsNoContent) String() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/events/{kind}][%d] getRunEventsNoContent  %+v", 204, o.Payload)
}

func (o *GetRunEventsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunEventsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunEventsForbidden creates a GetRunEventsForbidden with default headers values
func NewGetRunEventsForbidden() *GetRunEventsForbidden {
	return &GetRunEventsForbidden{}
}

/* GetRunEventsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetRunEventsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get run events forbidden response has a 2xx status code
func (o *GetRunEventsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get run events forbidden response has a 3xx status code
func (o *GetRunEventsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run events forbidden response has a 4xx status code
func (o *GetRunEventsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get run events forbidden response has a 5xx status code
func (o *GetRunEventsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get run events forbidden response a status code equal to that given
func (o *GetRunEventsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get run events forbidden response
func (o *GetRunEventsForbidden) Code() int {
	return 403
}

func (o *GetRunEventsForbidden) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/events/{kind}][%d] getRunEventsForbidden  %+v", 403, o.Payload)
}

func (o *GetRunEventsForbidden) String() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/events/{kind}][%d] getRunEventsForbidden  %+v", 403, o.Payload)
}

func (o *GetRunEventsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunEventsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunEventsNotFound creates a GetRunEventsNotFound with default headers values
func NewGetRunEventsNotFound() *GetRunEventsNotFound {
	return &GetRunEventsNotFound{}
}

/* GetRunEventsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetRunEventsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get run events not found response has a 2xx status code
func (o *GetRunEventsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get run events not found response has a 3xx status code
func (o *GetRunEventsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run events not found response has a 4xx status code
func (o *GetRunEventsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get run events not found response has a 5xx status code
func (o *GetRunEventsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get run events not found response a status code equal to that given
func (o *GetRunEventsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get run events not found response
func (o *GetRunEventsNotFound) Code() int {
	return 404
}

func (o *GetRunEventsNotFound) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/events/{kind}][%d] getRunEventsNotFound  %+v", 404, o.Payload)
}

func (o *GetRunEventsNotFound) String() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/events/{kind}][%d] getRunEventsNotFound  %+v", 404, o.Payload)
}

func (o *GetRunEventsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunEventsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunEventsDefault creates a GetRunEventsDefault with default headers values
func NewGetRunEventsDefault(code int) *GetRunEventsDefault {
	return &GetRunEventsDefault{
		_statusCode: code,
	}
}

/* GetRunEventsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetRunEventsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this get run events default response has a 2xx status code
func (o *GetRunEventsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get run events default response has a 3xx status code
func (o *GetRunEventsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get run events default response has a 4xx status code
func (o *GetRunEventsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get run events default response has a 5xx status code
func (o *GetRunEventsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get run events default response a status code equal to that given
func (o *GetRunEventsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get run events default response
func (o *GetRunEventsDefault) Code() int {
	return o._statusCode
}

func (o *GetRunEventsDefault) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/events/{kind}][%d] GetRunEvents default  %+v", o._statusCode, o.Payload)
}

func (o *GetRunEventsDefault) String() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/events/{kind}][%d] GetRunEvents default  %+v", o._statusCode, o.Payload)
}

func (o *GetRunEventsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetRunEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
