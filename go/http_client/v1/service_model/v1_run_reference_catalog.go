// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1RunReferenceCatalog Run Settings catalog specification
//
// swagger:model v1RunReferenceCatalog
type V1RunReferenceCatalog struct {

	// Name
	Name string `json:"name,omitempty"`

	// Owner
	Owner string `json:"owner,omitempty"`

	// Project
	Project string `json:"project,omitempty"`
}

// Validate validates this v1 run reference catalog
func (m *V1RunReferenceCatalog) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 run reference catalog based on context it is used
func (m *V1RunReferenceCatalog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1RunReferenceCatalog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RunReferenceCatalog) UnmarshalBinary(b []byte) error {
	var res V1RunReferenceCatalog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
