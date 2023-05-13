// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Auth Auth specification
//
// swagger:model v1Auth
type V1Auth struct {

	// token hash
	Token string `json:"token,omitempty"`
}

// Validate validates this v1 auth
func (m *V1Auth) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 auth based on context it is used
func (m *V1Auth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1Auth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Auth) UnmarshalBinary(b []byte) error {
	var res V1Auth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
