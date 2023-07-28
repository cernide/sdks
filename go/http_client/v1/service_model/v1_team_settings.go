// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1TeamSettings Run Settings specification
//
// swagger:model v1TeamSettings
type V1TeamSettings struct {

	// Component hubs
	Hubs []*V1SettingsCatalog `json:"hubs"`

	// Projects
	Projects []*V1SettingsCatalog `json:"projects"`

	// Model registries
	Registries []*V1SettingsCatalog `json:"registries"`
}

// Validate validates this v1 team settings
func (m *V1TeamSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHubs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TeamSettings) validateHubs(formats strfmt.Registry) error {
	if swag.IsZero(m.Hubs) { // not required
		return nil
	}

	for i := 0; i < len(m.Hubs); i++ {
		if swag.IsZero(m.Hubs[i]) { // not required
			continue
		}

		if m.Hubs[i] != nil {
			if err := m.Hubs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hubs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hubs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1TeamSettings) validateProjects(formats strfmt.Registry) error {
	if swag.IsZero(m.Projects) { // not required
		return nil
	}

	for i := 0; i < len(m.Projects); i++ {
		if swag.IsZero(m.Projects[i]) { // not required
			continue
		}

		if m.Projects[i] != nil {
			if err := m.Projects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("projects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1TeamSettings) validateRegistries(formats strfmt.Registry) error {
	if swag.IsZero(m.Registries) { // not required
		return nil
	}

	for i := 0; i < len(m.Registries); i++ {
		if swag.IsZero(m.Registries[i]) { // not required
			continue
		}

		if m.Registries[i] != nil {
			if err := m.Registries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("registries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("registries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 team settings based on the context it is used
func (m *V1TeamSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHubs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegistries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TeamSettings) contextValidateHubs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Hubs); i++ {

		if m.Hubs[i] != nil {

			if swag.IsZero(m.Hubs[i]) { // not required
				return nil
			}

			if err := m.Hubs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hubs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hubs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1TeamSettings) contextValidateProjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Projects); i++ {

		if m.Projects[i] != nil {

			if swag.IsZero(m.Projects[i]) { // not required
				return nil
			}

			if err := m.Projects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("projects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1TeamSettings) contextValidateRegistries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Registries); i++ {

		if m.Registries[i] != nil {

			if swag.IsZero(m.Registries[i]) { // not required
				return nil
			}

			if err := m.Registries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("registries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("registries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1TeamSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TeamSettings) UnmarshalBinary(b []byte) error {
	var res V1TeamSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
