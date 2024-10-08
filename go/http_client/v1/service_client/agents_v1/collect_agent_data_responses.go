// Code generated by go-swagger; DO NOT EDIT.

package agents_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/v2/go/http_client/v1/service_model"
)

// CollectAgentDataReader is a Reader for the CollectAgentData structure.
type CollectAgentDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CollectAgentDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCollectAgentDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCollectAgentDataNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCollectAgentDataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCollectAgentDataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCollectAgentDataDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCollectAgentDataOK creates a CollectAgentDataOK with default headers values
func NewCollectAgentDataOK() *CollectAgentDataOK {
	return &CollectAgentDataOK{}
}

/*
CollectAgentDataOK describes a response with status code 200, with default header values.

A successful response.
*/
type CollectAgentDataOK struct {
	Payload interface{}
}

// IsSuccess returns true when this collect agent data o k response has a 2xx status code
func (o *CollectAgentDataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this collect agent data o k response has a 3xx status code
func (o *CollectAgentDataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this collect agent data o k response has a 4xx status code
func (o *CollectAgentDataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this collect agent data o k response has a 5xx status code
func (o *CollectAgentDataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this collect agent data o k response a status code equal to that given
func (o *CollectAgentDataOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the collect agent data o k response
func (o *CollectAgentDataOK) Code() int {
	return 200
}

func (o *CollectAgentDataOK) Error() string {
	return fmt.Sprintf("[POST /internal/v1/{namespace}/{owner}/agents/{uuid}/collect][%d] collectAgentDataOK  %+v", 200, o.Payload)
}

func (o *CollectAgentDataOK) String() string {
	return fmt.Sprintf("[POST /internal/v1/{namespace}/{owner}/agents/{uuid}/collect][%d] collectAgentDataOK  %+v", 200, o.Payload)
}

func (o *CollectAgentDataOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CollectAgentDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCollectAgentDataNoContent creates a CollectAgentDataNoContent with default headers values
func NewCollectAgentDataNoContent() *CollectAgentDataNoContent {
	return &CollectAgentDataNoContent{}
}

/*
CollectAgentDataNoContent describes a response with status code 204, with default header values.

No content.
*/
type CollectAgentDataNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this collect agent data no content response has a 2xx status code
func (o *CollectAgentDataNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this collect agent data no content response has a 3xx status code
func (o *CollectAgentDataNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this collect agent data no content response has a 4xx status code
func (o *CollectAgentDataNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this collect agent data no content response has a 5xx status code
func (o *CollectAgentDataNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this collect agent data no content response a status code equal to that given
func (o *CollectAgentDataNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the collect agent data no content response
func (o *CollectAgentDataNoContent) Code() int {
	return 204
}

func (o *CollectAgentDataNoContent) Error() string {
	return fmt.Sprintf("[POST /internal/v1/{namespace}/{owner}/agents/{uuid}/collect][%d] collectAgentDataNoContent  %+v", 204, o.Payload)
}

func (o *CollectAgentDataNoContent) String() string {
	return fmt.Sprintf("[POST /internal/v1/{namespace}/{owner}/agents/{uuid}/collect][%d] collectAgentDataNoContent  %+v", 204, o.Payload)
}

func (o *CollectAgentDataNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CollectAgentDataNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCollectAgentDataForbidden creates a CollectAgentDataForbidden with default headers values
func NewCollectAgentDataForbidden() *CollectAgentDataForbidden {
	return &CollectAgentDataForbidden{}
}

/*
CollectAgentDataForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type CollectAgentDataForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this collect agent data forbidden response has a 2xx status code
func (o *CollectAgentDataForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this collect agent data forbidden response has a 3xx status code
func (o *CollectAgentDataForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this collect agent data forbidden response has a 4xx status code
func (o *CollectAgentDataForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this collect agent data forbidden response has a 5xx status code
func (o *CollectAgentDataForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this collect agent data forbidden response a status code equal to that given
func (o *CollectAgentDataForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the collect agent data forbidden response
func (o *CollectAgentDataForbidden) Code() int {
	return 403
}

func (o *CollectAgentDataForbidden) Error() string {
	return fmt.Sprintf("[POST /internal/v1/{namespace}/{owner}/agents/{uuid}/collect][%d] collectAgentDataForbidden  %+v", 403, o.Payload)
}

func (o *CollectAgentDataForbidden) String() string {
	return fmt.Sprintf("[POST /internal/v1/{namespace}/{owner}/agents/{uuid}/collect][%d] collectAgentDataForbidden  %+v", 403, o.Payload)
}

func (o *CollectAgentDataForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CollectAgentDataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCollectAgentDataNotFound creates a CollectAgentDataNotFound with default headers values
func NewCollectAgentDataNotFound() *CollectAgentDataNotFound {
	return &CollectAgentDataNotFound{}
}

/*
CollectAgentDataNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type CollectAgentDataNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this collect agent data not found response has a 2xx status code
func (o *CollectAgentDataNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this collect agent data not found response has a 3xx status code
func (o *CollectAgentDataNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this collect agent data not found response has a 4xx status code
func (o *CollectAgentDataNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this collect agent data not found response has a 5xx status code
func (o *CollectAgentDataNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this collect agent data not found response a status code equal to that given
func (o *CollectAgentDataNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the collect agent data not found response
func (o *CollectAgentDataNotFound) Code() int {
	return 404
}

func (o *CollectAgentDataNotFound) Error() string {
	return fmt.Sprintf("[POST /internal/v1/{namespace}/{owner}/agents/{uuid}/collect][%d] collectAgentDataNotFound  %+v", 404, o.Payload)
}

func (o *CollectAgentDataNotFound) String() string {
	return fmt.Sprintf("[POST /internal/v1/{namespace}/{owner}/agents/{uuid}/collect][%d] collectAgentDataNotFound  %+v", 404, o.Payload)
}

func (o *CollectAgentDataNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CollectAgentDataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCollectAgentDataDefault creates a CollectAgentDataDefault with default headers values
func NewCollectAgentDataDefault(code int) *CollectAgentDataDefault {
	return &CollectAgentDataDefault{
		_statusCode: code,
	}
}

/*
CollectAgentDataDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CollectAgentDataDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this collect agent data default response has a 2xx status code
func (o *CollectAgentDataDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this collect agent data default response has a 3xx status code
func (o *CollectAgentDataDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this collect agent data default response has a 4xx status code
func (o *CollectAgentDataDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this collect agent data default response has a 5xx status code
func (o *CollectAgentDataDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this collect agent data default response a status code equal to that given
func (o *CollectAgentDataDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the collect agent data default response
func (o *CollectAgentDataDefault) Code() int {
	return o._statusCode
}

func (o *CollectAgentDataDefault) Error() string {
	return fmt.Sprintf("[POST /internal/v1/{namespace}/{owner}/agents/{uuid}/collect][%d] CollectAgentData default  %+v", o._statusCode, o.Payload)
}

func (o *CollectAgentDataDefault) String() string {
	return fmt.Sprintf("[POST /internal/v1/{namespace}/{owner}/agents/{uuid}/collect][%d] CollectAgentData default  %+v", o._statusCode, o.Payload)
}

func (o *CollectAgentDataDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *CollectAgentDataDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
