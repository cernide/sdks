// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1LogHandler v1 log handler
//
// swagger:model v1LogHandler
type V1LogHandler struct {

	// dsn
	Dsn string `json:"dsn,omitempty"`

	// environment
	Environment string `json:"environment,omitempty"`
}

// Validate validates this v1 log handler
func (m *V1LogHandler) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 log handler based on context it is used
func (m *V1LogHandler) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1LogHandler) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1LogHandler) UnmarshalBinary(b []byte) error {
	var res V1LogHandler
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
