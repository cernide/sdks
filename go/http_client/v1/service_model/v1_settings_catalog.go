// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SettingsCatalog Settings catalog specification
//
// swagger:model v1SettingsCatalog
type V1SettingsCatalog struct {

	// Name
	Name string `json:"name,omitempty"`

	// Uuid
	UUID string `json:"uuid,omitempty"`

	// Version
	Version string `json:"version,omitempty"`
}

// Validate validates this v1 settings catalog
func (m *V1SettingsCatalog) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 settings catalog based on context it is used
func (m *V1SettingsCatalog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SettingsCatalog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SettingsCatalog) UnmarshalBinary(b []byte) error {
	var res V1SettingsCatalog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
