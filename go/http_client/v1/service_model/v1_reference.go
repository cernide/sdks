// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Reference Reference specification
//
// swagger:model v1Reference
type V1Reference struct {

	// dag ref
	DagRef *V1DagRef `json:"dagRef,omitempty"`

	// hub ref
	HubRef *V1HubRef `json:"hubRef,omitempty"`

	// path ref
	PathRef *V1PathRef `json:"pathRef,omitempty"`

	// url ref
	URLRef *V1URLRef `json:"urlRef,omitempty"`
}

// Validate validates this v1 reference
func (m *V1Reference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDagRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHubRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePathRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURLRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Reference) validateDagRef(formats strfmt.Registry) error {
	if swag.IsZero(m.DagRef) { // not required
		return nil
	}

	if m.DagRef != nil {
		if err := m.DagRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dagRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dagRef")
			}
			return err
		}
	}

	return nil
}

func (m *V1Reference) validateHubRef(formats strfmt.Registry) error {
	if swag.IsZero(m.HubRef) { // not required
		return nil
	}

	if m.HubRef != nil {
		if err := m.HubRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hubRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hubRef")
			}
			return err
		}
	}

	return nil
}

func (m *V1Reference) validatePathRef(formats strfmt.Registry) error {
	if swag.IsZero(m.PathRef) { // not required
		return nil
	}

	if m.PathRef != nil {
		if err := m.PathRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pathRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pathRef")
			}
			return err
		}
	}

	return nil
}

func (m *V1Reference) validateURLRef(formats strfmt.Registry) error {
	if swag.IsZero(m.URLRef) { // not required
		return nil
	}

	if m.URLRef != nil {
		if err := m.URLRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("urlRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("urlRef")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 reference based on the context it is used
func (m *V1Reference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDagRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHubRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePathRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURLRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Reference) contextValidateDagRef(ctx context.Context, formats strfmt.Registry) error {

	if m.DagRef != nil {

		if swag.IsZero(m.DagRef) { // not required
			return nil
		}

		if err := m.DagRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dagRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dagRef")
			}
			return err
		}
	}

	return nil
}

func (m *V1Reference) contextValidateHubRef(ctx context.Context, formats strfmt.Registry) error {

	if m.HubRef != nil {

		if swag.IsZero(m.HubRef) { // not required
			return nil
		}

		if err := m.HubRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hubRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hubRef")
			}
			return err
		}
	}

	return nil
}

func (m *V1Reference) contextValidatePathRef(ctx context.Context, formats strfmt.Registry) error {

	if m.PathRef != nil {

		if swag.IsZero(m.PathRef) { // not required
			return nil
		}

		if err := m.PathRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pathRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pathRef")
			}
			return err
		}
	}

	return nil
}

func (m *V1Reference) contextValidateURLRef(ctx context.Context, formats strfmt.Registry) error {

	if m.URLRef != nil {

		if swag.IsZero(m.URLRef) { // not required
			return nil
		}

		if err := m.URLRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("urlRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("urlRef")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Reference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Reference) UnmarshalBinary(b []byte) error {
	var res V1Reference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
