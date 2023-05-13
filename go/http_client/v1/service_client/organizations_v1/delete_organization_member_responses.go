// Code generated by go-swagger; DO NOT EDIT.

package organizations_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// DeleteOrganizationMemberReader is a Reader for the DeleteOrganizationMember structure.
type DeleteOrganizationMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOrganizationMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteOrganizationMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteOrganizationMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOrganizationMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteOrganizationMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteOrganizationMemberOK creates a DeleteOrganizationMemberOK with default headers values
func NewDeleteOrganizationMemberOK() *DeleteOrganizationMemberOK {
	return &DeleteOrganizationMemberOK{}
}

/*
DeleteOrganizationMemberOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteOrganizationMemberOK struct {
}

// IsSuccess returns true when this delete organization member o k response has a 2xx status code
func (o *DeleteOrganizationMemberOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete organization member o k response has a 3xx status code
func (o *DeleteOrganizationMemberOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organization member o k response has a 4xx status code
func (o *DeleteOrganizationMemberOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete organization member o k response has a 5xx status code
func (o *DeleteOrganizationMemberOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organization member o k response a status code equal to that given
func (o *DeleteOrganizationMemberOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete organization member o k response
func (o *DeleteOrganizationMemberOK) Code() int {
	return 200
}

func (o *DeleteOrganizationMemberOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/members/{name}][%d] deleteOrganizationMemberOK ", 200)
}

func (o *DeleteOrganizationMemberOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/members/{name}][%d] deleteOrganizationMemberOK ", 200)
}

func (o *DeleteOrganizationMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationMemberNoContent creates a DeleteOrganizationMemberNoContent with default headers values
func NewDeleteOrganizationMemberNoContent() *DeleteOrganizationMemberNoContent {
	return &DeleteOrganizationMemberNoContent{}
}

/*
DeleteOrganizationMemberNoContent describes a response with status code 204, with default header values.

No content.
*/
type DeleteOrganizationMemberNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this delete organization member no content response has a 2xx status code
func (o *DeleteOrganizationMemberNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete organization member no content response has a 3xx status code
func (o *DeleteOrganizationMemberNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organization member no content response has a 4xx status code
func (o *DeleteOrganizationMemberNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete organization member no content response has a 5xx status code
func (o *DeleteOrganizationMemberNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organization member no content response a status code equal to that given
func (o *DeleteOrganizationMemberNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete organization member no content response
func (o *DeleteOrganizationMemberNoContent) Code() int {
	return 204
}

func (o *DeleteOrganizationMemberNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/members/{name}][%d] deleteOrganizationMemberNoContent  %+v", 204, o.Payload)
}

func (o *DeleteOrganizationMemberNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/members/{name}][%d] deleteOrganizationMemberNoContent  %+v", 204, o.Payload)
}

func (o *DeleteOrganizationMemberNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteOrganizationMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationMemberForbidden creates a DeleteOrganizationMemberForbidden with default headers values
func NewDeleteOrganizationMemberForbidden() *DeleteOrganizationMemberForbidden {
	return &DeleteOrganizationMemberForbidden{}
}

/*
DeleteOrganizationMemberForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type DeleteOrganizationMemberForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this delete organization member forbidden response has a 2xx status code
func (o *DeleteOrganizationMemberForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete organization member forbidden response has a 3xx status code
func (o *DeleteOrganizationMemberForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organization member forbidden response has a 4xx status code
func (o *DeleteOrganizationMemberForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete organization member forbidden response has a 5xx status code
func (o *DeleteOrganizationMemberForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organization member forbidden response a status code equal to that given
func (o *DeleteOrganizationMemberForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete organization member forbidden response
func (o *DeleteOrganizationMemberForbidden) Code() int {
	return 403
}

func (o *DeleteOrganizationMemberForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/members/{name}][%d] deleteOrganizationMemberForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOrganizationMemberForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/members/{name}][%d] deleteOrganizationMemberForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOrganizationMemberForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteOrganizationMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationMemberNotFound creates a DeleteOrganizationMemberNotFound with default headers values
func NewDeleteOrganizationMemberNotFound() *DeleteOrganizationMemberNotFound {
	return &DeleteOrganizationMemberNotFound{}
}

/*
DeleteOrganizationMemberNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type DeleteOrganizationMemberNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this delete organization member not found response has a 2xx status code
func (o *DeleteOrganizationMemberNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete organization member not found response has a 3xx status code
func (o *DeleteOrganizationMemberNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organization member not found response has a 4xx status code
func (o *DeleteOrganizationMemberNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete organization member not found response has a 5xx status code
func (o *DeleteOrganizationMemberNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organization member not found response a status code equal to that given
func (o *DeleteOrganizationMemberNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete organization member not found response
func (o *DeleteOrganizationMemberNotFound) Code() int {
	return 404
}

func (o *DeleteOrganizationMemberNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/members/{name}][%d] deleteOrganizationMemberNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOrganizationMemberNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/members/{name}][%d] deleteOrganizationMemberNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOrganizationMemberNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteOrganizationMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationMemberDefault creates a DeleteOrganizationMemberDefault with default headers values
func NewDeleteOrganizationMemberDefault(code int) *DeleteOrganizationMemberDefault {
	return &DeleteOrganizationMemberDefault{
		_statusCode: code,
	}
}

/*
DeleteOrganizationMemberDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteOrganizationMemberDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// IsSuccess returns true when this delete organization member default response has a 2xx status code
func (o *DeleteOrganizationMemberDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete organization member default response has a 3xx status code
func (o *DeleteOrganizationMemberDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete organization member default response has a 4xx status code
func (o *DeleteOrganizationMemberDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete organization member default response has a 5xx status code
func (o *DeleteOrganizationMemberDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete organization member default response a status code equal to that given
func (o *DeleteOrganizationMemberDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete organization member default response
func (o *DeleteOrganizationMemberDefault) Code() int {
	return o._statusCode
}

func (o *DeleteOrganizationMemberDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/members/{name}][%d] DeleteOrganizationMember default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteOrganizationMemberDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/members/{name}][%d] DeleteOrganizationMember default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteOrganizationMemberDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *DeleteOrganizationMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
