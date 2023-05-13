// Code generated by go-swagger; DO NOT EDIT.

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetRunArtifactReader is a Reader for the GetRunArtifact structure.
type GetRunArtifactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunArtifactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunArtifactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetRunArtifactNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRunArtifactForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRunArtifactNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRunArtifactOK creates a GetRunArtifactOK with default headers values
func NewGetRunArtifactOK() *GetRunArtifactOK {
	return &GetRunArtifactOK{}
}

/*
GetRunArtifactOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetRunArtifactOK struct {
	Payload string
}

// IsSuccess returns true when this get run artifact o k response has a 2xx status code
func (o *GetRunArtifactOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get run artifact o k response has a 3xx status code
func (o *GetRunArtifactOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run artifact o k response has a 4xx status code
func (o *GetRunArtifactOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get run artifact o k response has a 5xx status code
func (o *GetRunArtifactOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get run artifact o k response a status code equal to that given
func (o *GetRunArtifactOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get run artifact o k response
func (o *GetRunArtifactOK) Code() int {
	return 200
}

func (o *GetRunArtifactOK) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/artifact][%d] getRunArtifactOK  %+v", 200, o.Payload)
}

func (o *GetRunArtifactOK) String() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/artifact][%d] getRunArtifactOK  %+v", 200, o.Payload)
}

func (o *GetRunArtifactOK) GetPayload() string {
	return o.Payload
}

func (o *GetRunArtifactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunArtifactNoContent creates a GetRunArtifactNoContent with default headers values
func NewGetRunArtifactNoContent() *GetRunArtifactNoContent {
	return &GetRunArtifactNoContent{}
}

/*
GetRunArtifactNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetRunArtifactNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this get run artifact no content response has a 2xx status code
func (o *GetRunArtifactNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get run artifact no content response has a 3xx status code
func (o *GetRunArtifactNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run artifact no content response has a 4xx status code
func (o *GetRunArtifactNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get run artifact no content response has a 5xx status code
func (o *GetRunArtifactNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get run artifact no content response a status code equal to that given
func (o *GetRunArtifactNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get run artifact no content response
func (o *GetRunArtifactNoContent) Code() int {
	return 204
}

func (o *GetRunArtifactNoContent) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/artifact][%d] getRunArtifactNoContent  %+v", 204, o.Payload)
}

func (o *GetRunArtifactNoContent) String() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/artifact][%d] getRunArtifactNoContent  %+v", 204, o.Payload)
}

func (o *GetRunArtifactNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunArtifactNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunArtifactForbidden creates a GetRunArtifactForbidden with default headers values
func NewGetRunArtifactForbidden() *GetRunArtifactForbidden {
	return &GetRunArtifactForbidden{}
}

/*
GetRunArtifactForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetRunArtifactForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this get run artifact forbidden response has a 2xx status code
func (o *GetRunArtifactForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get run artifact forbidden response has a 3xx status code
func (o *GetRunArtifactForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run artifact forbidden response has a 4xx status code
func (o *GetRunArtifactForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get run artifact forbidden response has a 5xx status code
func (o *GetRunArtifactForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get run artifact forbidden response a status code equal to that given
func (o *GetRunArtifactForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get run artifact forbidden response
func (o *GetRunArtifactForbidden) Code() int {
	return 403
}

func (o *GetRunArtifactForbidden) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/artifact][%d] getRunArtifactForbidden  %+v", 403, o.Payload)
}

func (o *GetRunArtifactForbidden) String() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/artifact][%d] getRunArtifactForbidden  %+v", 403, o.Payload)
}

func (o *GetRunArtifactForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunArtifactForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunArtifactNotFound creates a GetRunArtifactNotFound with default headers values
func NewGetRunArtifactNotFound() *GetRunArtifactNotFound {
	return &GetRunArtifactNotFound{}
}

/*
GetRunArtifactNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetRunArtifactNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get run artifact not found response has a 2xx status code
func (o *GetRunArtifactNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get run artifact not found response has a 3xx status code
func (o *GetRunArtifactNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run artifact not found response has a 4xx status code
func (o *GetRunArtifactNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get run artifact not found response has a 5xx status code
func (o *GetRunArtifactNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get run artifact not found response a status code equal to that given
func (o *GetRunArtifactNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get run artifact not found response
func (o *GetRunArtifactNotFound) Code() int {
	return 404
}

func (o *GetRunArtifactNotFound) Error() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/artifact][%d] getRunArtifactNotFound  %+v", 404, o.Payload)
}

func (o *GetRunArtifactNotFound) String() string {
	return fmt.Sprintf("[GET /streams/v1/{namespace}/{owner}/{project}/runs/{uuid}/artifact][%d] getRunArtifactNotFound  %+v", 404, o.Payload)
}

func (o *GetRunArtifactNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunArtifactNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
