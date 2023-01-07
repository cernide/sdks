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

package presets_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// PatchPresetReader is a Reader for the PatchPreset structure.
type PatchPresetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchPresetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchPresetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchPresetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchPresetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchPresetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchPresetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchPresetOK creates a PatchPresetOK with default headers values
func NewPatchPresetOK() *PatchPresetOK {
	return &PatchPresetOK{}
}

/* PatchPresetOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchPresetOK struct {
	Payload *service_model.V1Preset
}

// IsSuccess returns true when this patch preset o k response has a 2xx status code
func (o *PatchPresetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch preset o k response has a 3xx status code
func (o *PatchPresetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch preset o k response has a 4xx status code
func (o *PatchPresetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch preset o k response has a 5xx status code
func (o *PatchPresetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch preset o k response a status code equal to that given
func (o *PatchPresetOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchPresetOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/presets/{preset.uuid}][%d] patchPresetOK  %+v", 200, o.Payload)
}

func (o *PatchPresetOK) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/presets/{preset.uuid}][%d] patchPresetOK  %+v", 200, o.Payload)
}

func (o *PatchPresetOK) GetPayload() *service_model.V1Preset {
	return o.Payload
}

func (o *PatchPresetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Preset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPresetNoContent creates a PatchPresetNoContent with default headers values
func NewPatchPresetNoContent() *PatchPresetNoContent {
	return &PatchPresetNoContent{}
}

/* PatchPresetNoContent describes a response with status code 204, with default header values.

No content.
*/
type PatchPresetNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this patch preset no content response has a 2xx status code
func (o *PatchPresetNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch preset no content response has a 3xx status code
func (o *PatchPresetNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch preset no content response has a 4xx status code
func (o *PatchPresetNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch preset no content response has a 5xx status code
func (o *PatchPresetNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this patch preset no content response a status code equal to that given
func (o *PatchPresetNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *PatchPresetNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/presets/{preset.uuid}][%d] patchPresetNoContent  %+v", 204, o.Payload)
}

func (o *PatchPresetNoContent) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/presets/{preset.uuid}][%d] patchPresetNoContent  %+v", 204, o.Payload)
}

func (o *PatchPresetNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchPresetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPresetForbidden creates a PatchPresetForbidden with default headers values
func NewPatchPresetForbidden() *PatchPresetForbidden {
	return &PatchPresetForbidden{}
}

/* PatchPresetForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type PatchPresetForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this patch preset forbidden response has a 2xx status code
func (o *PatchPresetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch preset forbidden response has a 3xx status code
func (o *PatchPresetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch preset forbidden response has a 4xx status code
func (o *PatchPresetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch preset forbidden response has a 5xx status code
func (o *PatchPresetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch preset forbidden response a status code equal to that given
func (o *PatchPresetForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchPresetForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/presets/{preset.uuid}][%d] patchPresetForbidden  %+v", 403, o.Payload)
}

func (o *PatchPresetForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/presets/{preset.uuid}][%d] patchPresetForbidden  %+v", 403, o.Payload)
}

func (o *PatchPresetForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchPresetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPresetNotFound creates a PatchPresetNotFound with default headers values
func NewPatchPresetNotFound() *PatchPresetNotFound {
	return &PatchPresetNotFound{}
}

/* PatchPresetNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type PatchPresetNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this patch preset not found response has a 2xx status code
func (o *PatchPresetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch preset not found response has a 3xx status code
func (o *PatchPresetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch preset not found response has a 4xx status code
func (o *PatchPresetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch preset not found response has a 5xx status code
func (o *PatchPresetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch preset not found response a status code equal to that given
func (o *PatchPresetNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchPresetNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/presets/{preset.uuid}][%d] patchPresetNotFound  %+v", 404, o.Payload)
}

func (o *PatchPresetNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/presets/{preset.uuid}][%d] patchPresetNotFound  %+v", 404, o.Payload)
}

func (o *PatchPresetNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchPresetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPresetDefault creates a PatchPresetDefault with default headers values
func NewPatchPresetDefault(code int) *PatchPresetDefault {
	return &PatchPresetDefault{
		_statusCode: code,
	}
}

/* PatchPresetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PatchPresetDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the patch preset default response
func (o *PatchPresetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this patch preset default response has a 2xx status code
func (o *PatchPresetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch preset default response has a 3xx status code
func (o *PatchPresetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch preset default response has a 4xx status code
func (o *PatchPresetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch preset default response has a 5xx status code
func (o *PatchPresetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch preset default response a status code equal to that given
func (o *PatchPresetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PatchPresetDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/presets/{preset.uuid}][%d] PatchPreset default  %+v", o._statusCode, o.Payload)
}

func (o *PatchPresetDefault) String() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/presets/{preset.uuid}][%d] PatchPreset default  %+v", o._statusCode, o.Payload)
}

func (o *PatchPresetDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *PatchPresetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
