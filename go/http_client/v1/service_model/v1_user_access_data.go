// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1UserAccessData v1 user access data
//
// swagger:model v1UserAccessData
type V1UserAccessData struct {

	// Is service account
	IsSa bool `json:"is_sa,omitempty"`

	// Username
	Username string `json:"username,omitempty"`
}

// Validate validates this v1 user access data
func (m *V1UserAccessData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 user access data based on context it is used
func (m *V1UserAccessData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1UserAccessData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UserAccessData) UnmarshalBinary(b []byte) error {
	var res V1UserAccessData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}