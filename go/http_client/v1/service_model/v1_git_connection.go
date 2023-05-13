// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1GitConnection Git connection schema validation connection
//
// swagger:model v1GitConnection
type V1GitConnection struct {

	// Additional command flag
	Flags []string `json:"flags"`

	// revision
	Revision bool `json:"revision,omitempty"`

	// Url
	URL string `json:"url,omitempty"`
}

// Validate validates this v1 git connection
func (m *V1GitConnection) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 git connection based on context it is used
func (m *V1GitConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1GitConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GitConnection) UnmarshalBinary(b []byte) error {
	var res V1GitConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
