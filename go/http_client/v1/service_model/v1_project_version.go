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
	"github.com/go-openapi/validate"
)

// V1ProjectVersion component hub specification
//
// swagger:model v1ProjectVersion
type V1ProjectVersion struct {

	// Artifacts lineage
	Artifacts []string `json:"artifacts"`

	// Connection name
	Connection string `json:"connection,omitempty"`

	// The metadata/content body
	Content string `json:"content,omitempty"`

	// Optional time when the entity was created
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Optional description
	Description string `json:"description,omitempty"`

	// Optional kind to tell the kind of this version
	Kind *V1ProjectVersionKind `json:"kind,omitempty"`

	// Extra information related to the run, lineage, artifacts, ...
	MetaInfo interface{} `json:"meta_info,omitempty"`

	// Optional component name, should be a valid fully qualified value: name[:version]
	Name string `json:"name,omitempty"`

	// Owner of this entity
	Owner string `json:"owner,omitempty"`

	// Project name
	Project string `json:"project,omitempty"`

	// Optional Markdown description/readme
	Readme string `json:"readme,omitempty"`

	// Current user's role in this (org/teams)/hub/version
	Role string `json:"role,omitempty"`

	// Run lineage
	Run string `json:"run,omitempty"`

	// Optional latest stage of this entity
	Stage *V1Stages `json:"stage,omitempty"`

	// The status conditions timeline
	StageConditions []*V1StageCondition `json:"stage_conditions"`

	// The version state read-only
	State string `json:"state,omitempty"`

	// Optional tags of this entity
	Tags []string `json:"tags"`

	// Optional last time the entity was updated
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// UUID
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this v1 project version
func (m *V1ProjectVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStageConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ProjectVersion) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ProjectVersion) validateKind(formats strfmt.Registry) error {
	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	if m.Kind != nil {
		if err := m.Kind.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kind")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kind")
			}
			return err
		}
	}

	return nil
}

func (m *V1ProjectVersion) validateStage(formats strfmt.Registry) error {
	if swag.IsZero(m.Stage) { // not required
		return nil
	}

	if m.Stage != nil {
		if err := m.Stage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stage")
			}
			return err
		}
	}

	return nil
}

func (m *V1ProjectVersion) validateStageConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.StageConditions) { // not required
		return nil
	}

	for i := 0; i < len(m.StageConditions); i++ {
		if swag.IsZero(m.StageConditions[i]) { // not required
			continue
		}

		if m.StageConditions[i] != nil {
			if err := m.StageConditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stage_conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("stage_conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ProjectVersion) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 project version based on the context it is used
func (m *V1ProjectVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKind(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStageConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ProjectVersion) contextValidateKind(ctx context.Context, formats strfmt.Registry) error {

	if m.Kind != nil {

		if swag.IsZero(m.Kind) { // not required
			return nil
		}

		if err := m.Kind.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kind")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kind")
			}
			return err
		}
	}

	return nil
}

func (m *V1ProjectVersion) contextValidateStage(ctx context.Context, formats strfmt.Registry) error {

	if m.Stage != nil {

		if swag.IsZero(m.Stage) { // not required
			return nil
		}

		if err := m.Stage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stage")
			}
			return err
		}
	}

	return nil
}

func (m *V1ProjectVersion) contextValidateStageConditions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StageConditions); i++ {

		if m.StageConditions[i] != nil {

			if swag.IsZero(m.StageConditions[i]) { // not required
				return nil
			}

			if err := m.StageConditions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stage_conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("stage_conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ProjectVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ProjectVersion) UnmarshalBinary(b []byte) error {
	var res V1ProjectVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
