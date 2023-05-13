// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1HpLinSpace HP Linear Space specification
//
// swagger:model v1HpLinSpace
type V1HpLinSpace struct {

	// Kind of hp, should be equal to "linspace"
	Kind *string `json:"kind,omitempty"`

	// Value
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this v1 hp lin space
func (m *V1HpLinSpace) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 hp lin space based on context it is used
func (m *V1HpLinSpace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1HpLinSpace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1HpLinSpace) UnmarshalBinary(b []byte) error {
	var res V1HpLinSpace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
