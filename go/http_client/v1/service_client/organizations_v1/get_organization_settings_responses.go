// Code generated by go-swagger; DO NOT EDIT.

package organizations_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// GetOrganizationSettingsReader is a Reader for the GetOrganizationSettings structure.
type GetOrganizationSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetOrganizationSettingsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetOrganizationSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetOrganizationSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOrganizationSettingsOK creates a GetOrganizationSettingsOK with default headers values
func NewGetOrganizationSettingsOK() *GetOrganizationSettingsOK {
	return &GetOrganizationSettingsOK{}
}

/*
GetOrganizationSettingsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetOrganizationSettingsOK struct {
	Payload *service_model.V1Organization
}

// IsSuccess returns true when this get organization settings o k response has a 2xx status code
func (o *GetOrganizationSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization settings o k response has a 3xx status code
func (o *GetOrganizationSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization settings o k response has a 4xx status code
func (o *GetOrganizationSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization settings o k response has a 5xx status code
func (o *GetOrganizationSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization settings o k response a status code equal to that given
func (o *GetOrganizationSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization settings o k response
func (o *GetOrganizationSettingsOK) Code() int {
	return 200
}

func (o *GetOrganizationSettingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/settings][%d] getOrganizationSettingsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationSettingsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/settings][%d] getOrganizationSettingsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationSettingsOK) GetPayload() *service_model.V1Organization {
	return o.Payload
}

func (o *GetOrganizationSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Organization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationSettingsNoContent creates a GetOrganizationSettingsNoContent with default headers values
func NewGetOrganizationSettingsNoContent() *GetOrganizationSettingsNoContent {
	return &GetOrganizationSettingsNoContent{}
}

/*
GetOrganizationSettingsNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetOrganizationSettingsNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get organization settings no content response has a 2xx status code
func (o *GetOrganizationSettingsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization settings no content response has a 3xx status code
func (o *GetOrganizationSettingsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization settings no content response has a 4xx status code
func (o *GetOrganizationSettingsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization settings no content response has a 5xx status code
func (o *GetOrganizationSettingsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization settings no content response a status code equal to that given
func (o *GetOrganizationSettingsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get organization settings no content response
func (o *GetOrganizationSettingsNoContent) Code() int {
	return 204
}

func (o *GetOrganizationSettingsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/settings][%d] getOrganizationSettingsNoContent  %+v", 204, o.Payload)
}

func (o *GetOrganizationSettingsNoContent) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/settings][%d] getOrganizationSettingsNoContent  %+v", 204, o.Payload)
}

func (o *GetOrganizationSettingsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationSettingsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationSettingsForbidden creates a GetOrganizationSettingsForbidden with default headers values
func NewGetOrganizationSettingsForbidden() *GetOrganizationSettingsForbidden {
	return &GetOrganizationSettingsForbidden{}
}

/*
GetOrganizationSettingsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetOrganizationSettingsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get organization settings forbidden response has a 2xx status code
func (o *GetOrganizationSettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization settings forbidden response has a 3xx status code
func (o *GetOrganizationSettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization settings forbidden response has a 4xx status code
func (o *GetOrganizationSettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization settings forbidden response has a 5xx status code
func (o *GetOrganizationSettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization settings forbidden response a status code equal to that given
func (o *GetOrganizationSettingsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get organization settings forbidden response
func (o *GetOrganizationSettingsForbidden) Code() int {
	return 403
}

func (o *GetOrganizationSettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/settings][%d] getOrganizationSettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationSettingsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/settings][%d] getOrganizationSettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationSettingsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationSettingsNotFound creates a GetOrganizationSettingsNotFound with default headers values
func NewGetOrganizationSettingsNotFound() *GetOrganizationSettingsNotFound {
	return &GetOrganizationSettingsNotFound{}
}

/*
GetOrganizationSettingsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetOrganizationSettingsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get organization settings not found response has a 2xx status code
func (o *GetOrganizationSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization settings not found response has a 3xx status code
func (o *GetOrganizationSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization settings not found response has a 4xx status code
func (o *GetOrganizationSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization settings not found response has a 5xx status code
func (o *GetOrganizationSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization settings not found response a status code equal to that given
func (o *GetOrganizationSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get organization settings not found response
func (o *GetOrganizationSettingsNotFound) Code() int {
	return 404
}

func (o *GetOrganizationSettingsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/settings][%d] getOrganizationSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationSettingsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/settings][%d] getOrganizationSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationSettingsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationSettingsDefault creates a GetOrganizationSettingsDefault with default headers values
func NewGetOrganizationSettingsDefault(code int) *GetOrganizationSettingsDefault {
	return &GetOrganizationSettingsDefault{
		_statusCode: code,
	}
}

/*
GetOrganizationSettingsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetOrganizationSettingsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this get organization settings default response has a 2xx status code
func (o *GetOrganizationSettingsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get organization settings default response has a 3xx status code
func (o *GetOrganizationSettingsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get organization settings default response has a 4xx status code
func (o *GetOrganizationSettingsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get organization settings default response has a 5xx status code
func (o *GetOrganizationSettingsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get organization settings default response a status code equal to that given
func (o *GetOrganizationSettingsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get organization settings default response
func (o *GetOrganizationSettingsDefault) Code() int {
	return o._statusCode
}

func (o *GetOrganizationSettingsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/settings][%d] GetOrganizationSettings default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrganizationSettingsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/settings][%d] GetOrganizationSettings default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrganizationSettingsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetOrganizationSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
