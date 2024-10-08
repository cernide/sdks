// Code generated by go-swagger; DO NOT EDIT.

package presets_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// CreatePresetReader is a Reader for the CreatePreset structure.
type CreatePresetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePresetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePresetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCreatePresetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreatePresetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreatePresetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreatePresetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePresetOK creates a CreatePresetOK with default headers values
func NewCreatePresetOK() *CreatePresetOK {
	return &CreatePresetOK{}
}

/*
CreatePresetOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreatePresetOK struct {
	Payload *service_model.V1Preset
}

// IsSuccess returns true when this create preset o k response has a 2xx status code
func (o *CreatePresetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create preset o k response has a 3xx status code
func (o *CreatePresetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create preset o k response has a 4xx status code
func (o *CreatePresetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create preset o k response has a 5xx status code
func (o *CreatePresetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create preset o k response a status code equal to that given
func (o *CreatePresetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create preset o k response
func (o *CreatePresetOK) Code() int {
	return 200
}

func (o *CreatePresetOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/presets][%d] createPresetOK  %+v", 200, o.Payload)
}

func (o *CreatePresetOK) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/presets][%d] createPresetOK  %+v", 200, o.Payload)
}

func (o *CreatePresetOK) GetPayload() *service_model.V1Preset {
	return o.Payload
}

func (o *CreatePresetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Preset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePresetNoContent creates a CreatePresetNoContent with default headers values
func NewCreatePresetNoContent() *CreatePresetNoContent {
	return &CreatePresetNoContent{}
}

/*
CreatePresetNoContent describes a response with status code 204, with default header values.

No content.
*/
type CreatePresetNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this create preset no content response has a 2xx status code
func (o *CreatePresetNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create preset no content response has a 3xx status code
func (o *CreatePresetNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create preset no content response has a 4xx status code
func (o *CreatePresetNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this create preset no content response has a 5xx status code
func (o *CreatePresetNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this create preset no content response a status code equal to that given
func (o *CreatePresetNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the create preset no content response
func (o *CreatePresetNoContent) Code() int {
	return 204
}

func (o *CreatePresetNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/presets][%d] createPresetNoContent  %+v", 204, o.Payload)
}

func (o *CreatePresetNoContent) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/presets][%d] createPresetNoContent  %+v", 204, o.Payload)
}

func (o *CreatePresetNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CreatePresetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePresetForbidden creates a CreatePresetForbidden with default headers values
func NewCreatePresetForbidden() *CreatePresetForbidden {
	return &CreatePresetForbidden{}
}

/*
CreatePresetForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type CreatePresetForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this create preset forbidden response has a 2xx status code
func (o *CreatePresetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create preset forbidden response has a 3xx status code
func (o *CreatePresetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create preset forbidden response has a 4xx status code
func (o *CreatePresetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create preset forbidden response has a 5xx status code
func (o *CreatePresetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create preset forbidden response a status code equal to that given
func (o *CreatePresetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create preset forbidden response
func (o *CreatePresetForbidden) Code() int {
	return 403
}

func (o *CreatePresetForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/presets][%d] createPresetForbidden  %+v", 403, o.Payload)
}

func (o *CreatePresetForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/presets][%d] createPresetForbidden  %+v", 403, o.Payload)
}

func (o *CreatePresetForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreatePresetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePresetNotFound creates a CreatePresetNotFound with default headers values
func NewCreatePresetNotFound() *CreatePresetNotFound {
	return &CreatePresetNotFound{}
}

/*
CreatePresetNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type CreatePresetNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this create preset not found response has a 2xx status code
func (o *CreatePresetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create preset not found response has a 3xx status code
func (o *CreatePresetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create preset not found response has a 4xx status code
func (o *CreatePresetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create preset not found response has a 5xx status code
func (o *CreatePresetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create preset not found response a status code equal to that given
func (o *CreatePresetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create preset not found response
func (o *CreatePresetNotFound) Code() int {
	return 404
}

func (o *CreatePresetNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/presets][%d] createPresetNotFound  %+v", 404, o.Payload)
}

func (o *CreatePresetNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/presets][%d] createPresetNotFound  %+v", 404, o.Payload)
}

func (o *CreatePresetNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreatePresetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePresetDefault creates a CreatePresetDefault with default headers values
func NewCreatePresetDefault(code int) *CreatePresetDefault {
	return &CreatePresetDefault{
		_statusCode: code,
	}
}

/*
CreatePresetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreatePresetDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this create preset default response has a 2xx status code
func (o *CreatePresetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create preset default response has a 3xx status code
func (o *CreatePresetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create preset default response has a 4xx status code
func (o *CreatePresetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create preset default response has a 5xx status code
func (o *CreatePresetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create preset default response a status code equal to that given
func (o *CreatePresetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create preset default response
func (o *CreatePresetDefault) Code() int {
	return o._statusCode
}

func (o *CreatePresetDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/presets][%d] CreatePreset default  %+v", o._statusCode, o.Payload)
}

func (o *CreatePresetDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/presets][%d] CreatePreset default  %+v", o._statusCode, o.Payload)
}

func (o *CreatePresetDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *CreatePresetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
