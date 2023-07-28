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

// V1Run Run specification
//
// swagger:model v1Run
type V1Run struct {

	// Optional if this entity was bookmarked
	Bookmarked bool `json:"bookmarked,omitempty"`

	// Optional content of the entity's spec
	Content string `json:"content,omitempty"`

	// Optional time when the entity was created
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Optional description
	Description string `json:"description,omitempty"`

	// Optional duration of the entity
	Duration int32 `json:"duration,omitempty"`

	// Optional last time the entity was started
	// Format: date-time
	FinishedAt strfmt.DateTime `json:"finished_at,omitempty"`

	// Optional graph definition
	Graph interface{} `json:"graph,omitempty"`

	// Optional inputs of this entity
	Inputs interface{} `json:"inputs,omitempty"`

	// Deprecated flag that was replaced by "pending", and it will be completely dropped after v1.15
	IsApproved bool `json:"is_approved,omitempty"`

	// Optional flag to tell if this entity is managed by the platform
	IsManaged bool `json:"is_managed,omitempty"`

	// Optional kind to tell the kind of this run
	Kind *V1RunKind `json:"kind,omitempty"`

	// Current live state
	LiveState int32 `json:"live_state,omitempty"`

	// Optional flag of the managing service
	ManagedBy *V1ManagedBy `json:"managed_by,omitempty"`

	// Optional merge flag
	Merge bool `json:"merge,omitempty"`

	// Optional run meta info
	MetaInfo interface{} `json:"meta_info,omitempty"`

	// Optional name
	Name string `json:"name,omitempty"`

	// Optional original run meta information
	Original *V1Cloning `json:"original,omitempty"`

	// Optional outputs of this entity
	Outputs interface{} `json:"outputs,omitempty"`

	// Required name of owner of this entity
	Owner string `json:"owner,omitempty"`

	// Optional to tell if this entity requires approval before it should be scheduled
	Pending *V1RunPending `json:"pending,omitempty"`

	// Optional pipeline run meta information
	Pipeline *V1Pipeline `json:"pipeline,omitempty"`

	// Required project name
	Project string `json:"project,omitempty"`

	// Optional content of the entity's spec
	RawContent string `json:"raw_content,omitempty"`

	// Optional Markdown description/readme
	Readme string `json:"readme,omitempty"`

	// Options resources
	Resources *V1RunResources `json:"resources,omitempty"`

	// Current user's role in this (org/teams)/project/runs
	Role string `json:"role,omitempty"`

	// Optional meta kind to tell the nature of this run
	Runtime *V1RunKind `json:"runtime,omitempty"`

	// Optional last time the entity was started
	// Format: date-time
	ScheduleAt strfmt.DateTime `json:"schedule_at,omitempty"`

	// Optional settings
	Settings *V1RunSettings `json:"settings,omitempty"`

	// Optional last time the entity was started
	// Format: date-time
	StartedAt strfmt.DateTime `json:"started_at,omitempty"`

	// Optional latest status of this entity
	Status *V1Statuses `json:"status,omitempty"`

	// The status conditions timeline
	StatusConditions []*V1StatusCondition `json:"status_conditions"`

	// Optional tags of this entity
	Tags []string `json:"tags"`

	// Optional last time the entity was updated
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// Required name of user started this entity
	User string `json:"user,omitempty"`

	// UUID
	UUID string `json:"uuid,omitempty"`

	// Optional wait time of the entity
	WaitTime int32 `json:"wait_time,omitempty"`
}

// Validate validates this v1 run
func (m *V1Run) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinishedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePending(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePipeline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuntime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduleAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusConditions(formats); err != nil {
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

func (m *V1Run) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Run) validateFinishedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.FinishedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("finished_at", "body", "date-time", m.FinishedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Run) validateKind(formats strfmt.Registry) error {
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

func (m *V1Run) validateManagedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.ManagedBy) { // not required
		return nil
	}

	if m.ManagedBy != nil {
		if err := m.ManagedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("managed_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("managed_by")
			}
			return err
		}
	}

	return nil
}

func (m *V1Run) validateOriginal(formats strfmt.Registry) error {
	if swag.IsZero(m.Original) { // not required
		return nil
	}

	if m.Original != nil {
		if err := m.Original.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("original")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("original")
			}
			return err
		}
	}

	return nil
}

func (m *V1Run) validatePending(formats strfmt.Registry) error {
	if swag.IsZero(m.Pending) { // not required
		return nil
	}

	if m.Pending != nil {
		if err := m.Pending.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pending")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pending")
			}
			return err
		}
	}

	return nil
}

func (m *V1Run) validatePipeline(formats strfmt.Registry) error {
	if swag.IsZero(m.Pipeline) { // not required
		return nil
	}

	if m.Pipeline != nil {
		if err := m.Pipeline.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pipeline")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pipeline")
			}
			return err
		}
	}

	return nil
}

func (m *V1Run) validateResources(formats strfmt.Registry) error {
	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	if m.Resources != nil {
		if err := m.Resources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

func (m *V1Run) validateRuntime(formats strfmt.Registry) error {
	if swag.IsZero(m.Runtime) { // not required
		return nil
	}

	if m.Runtime != nil {
		if err := m.Runtime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runtime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runtime")
			}
			return err
		}
	}

	return nil
}

func (m *V1Run) validateScheduleAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ScheduleAt) { // not required
		return nil
	}

	if err := validate.FormatOf("schedule_at", "body", "date-time", m.ScheduleAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Run) validateSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	if m.Settings != nil {
		if err := m.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

func (m *V1Run) validateStartedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Run) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *V1Run) validateStatusConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusConditions) { // not required
		return nil
	}

	for i := 0; i < len(m.StatusConditions); i++ {
		if swag.IsZero(m.StatusConditions[i]) { // not required
			continue
		}

		if m.StatusConditions[i] != nil {
			if err := m.StatusConditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status_conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("status_conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Run) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 run based on the context it is used
func (m *V1Run) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKind(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateManagedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOriginal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePending(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePipeline(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRuntime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Run) contextValidateKind(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1Run) contextValidateManagedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.ManagedBy != nil {

		if swag.IsZero(m.ManagedBy) { // not required
			return nil
		}

		if err := m.ManagedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("managed_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("managed_by")
			}
			return err
		}
	}

	return nil
}

func (m *V1Run) contextValidateOriginal(ctx context.Context, formats strfmt.Registry) error {

	if m.Original != nil {

		if swag.IsZero(m.Original) { // not required
			return nil
		}

		if err := m.Original.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("original")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("original")
			}
			return err
		}
	}

	return nil
}

func (m *V1Run) contextValidatePending(ctx context.Context, formats strfmt.Registry) error {

	if m.Pending != nil {

		if swag.IsZero(m.Pending) { // not required
			return nil
		}

		if err := m.Pending.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pending")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pending")
			}
			return err
		}
	}

	return nil
}

func (m *V1Run) contextValidatePipeline(ctx context.Context, formats strfmt.Registry) error {

	if m.Pipeline != nil {

		if swag.IsZero(m.Pipeline) { // not required
			return nil
		}

		if err := m.Pipeline.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pipeline")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pipeline")
			}
			return err
		}
	}

	return nil
}

func (m *V1Run) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	if m.Resources != nil {

		if swag.IsZero(m.Resources) { // not required
			return nil
		}

		if err := m.Resources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

func (m *V1Run) contextValidateRuntime(ctx context.Context, formats strfmt.Registry) error {

	if m.Runtime != nil {

		if swag.IsZero(m.Runtime) { // not required
			return nil
		}

		if err := m.Runtime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runtime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runtime")
			}
			return err
		}
	}

	return nil
}

func (m *V1Run) contextValidateSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.Settings != nil {

		if swag.IsZero(m.Settings) { // not required
			return nil
		}

		if err := m.Settings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

func (m *V1Run) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *V1Run) contextValidateStatusConditions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatusConditions); i++ {

		if m.StatusConditions[i] != nil {

			if swag.IsZero(m.StatusConditions[i]) { // not required
				return nil
			}

			if err := m.StatusConditions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status_conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("status_conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Run) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Run) UnmarshalBinary(b []byte) error {
	var res V1Run
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
