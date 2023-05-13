// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1HpQNormal HP Quantized Normal Dist specification
//
// swagger:model v1HpQNormal
type V1HpQNormal struct {

	// Kind of hp, should be equal to "qnormal"
	Kind *string `json:"kind,omitempty"`

	// Value
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this v1 hp q normal
func (m *V1HpQNormal) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 hp q normal based on context it is used
func (m *V1HpQNormal) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1HpQNormal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1HpQNormal) UnmarshalBinary(b []byte) error {
	var res V1HpQNormal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
