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

// V1TFJob TF Job specification
//
// swagger:model v1TFJob
type V1TFJob struct {

	// Chief replicas definition
	Chief *V1KFReplica `json:"chief,omitempty"`

	// optional clean pod policy section
	CleanPodPolicy *V1CleanPodPolicy `json:"cleanPodPolicy,omitempty"`

	// optional flag to enable dynamic worker
	EnableDynamicWorker bool `json:"enableDynamicWorker,omitempty"`

	// Evaluator replicas definition
	Evaluator *V1KFReplica `json:"evaluator,omitempty"`

	// Optional component kind, should be equal to 'tfjob'
	Kind *string `json:"kind,omitempty"`

	// PS replicas definition
	Ps *V1KFReplica `json:"ps,omitempty"`

	// optional scheduling policy section
	SchedulingPolicy *V1SchedulingPolicy `json:"schedulingPolicy,omitempty"`

	// optiona success policy
	SuccessPolicy string `json:"successPolicy,omitempty"`

	// Worker replicas definition
	Worker *V1KFReplica `json:"worker,omitempty"`
}

// Validate validates this v1 t f job
func (m *V1TFJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChief(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCleanPodPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvaluator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedulingPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorker(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TFJob) validateChief(formats strfmt.Registry) error {
	if swag.IsZero(m.Chief) { // not required
		return nil
	}

	if m.Chief != nil {
		if err := m.Chief.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chief")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chief")
			}
			return err
		}
	}

	return nil
}

func (m *V1TFJob) validateCleanPodPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.CleanPodPolicy) { // not required
		return nil
	}

	if m.CleanPodPolicy != nil {
		if err := m.CleanPodPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cleanPodPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cleanPodPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *V1TFJob) validateEvaluator(formats strfmt.Registry) error {
	if swag.IsZero(m.Evaluator) { // not required
		return nil
	}

	if m.Evaluator != nil {
		if err := m.Evaluator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("evaluator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("evaluator")
			}
			return err
		}
	}

	return nil
}

func (m *V1TFJob) validatePs(formats strfmt.Registry) error {
	if swag.IsZero(m.Ps) { // not required
		return nil
	}

	if m.Ps != nil {
		if err := m.Ps.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ps")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ps")
			}
			return err
		}
	}

	return nil
}

func (m *V1TFJob) validateSchedulingPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.SchedulingPolicy) { // not required
		return nil
	}

	if m.SchedulingPolicy != nil {
		if err := m.SchedulingPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedulingPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedulingPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *V1TFJob) validateWorker(formats strfmt.Registry) error {
	if swag.IsZero(m.Worker) { // not required
		return nil
	}

	if m.Worker != nil {
		if err := m.Worker.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("worker")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("worker")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 t f job based on the context it is used
func (m *V1TFJob) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChief(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCleanPodPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEvaluator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchedulingPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorker(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TFJob) contextValidateChief(ctx context.Context, formats strfmt.Registry) error {

	if m.Chief != nil {

		if swag.IsZero(m.Chief) { // not required
			return nil
		}

		if err := m.Chief.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chief")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chief")
			}
			return err
		}
	}

	return nil
}

func (m *V1TFJob) contextValidateCleanPodPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.CleanPodPolicy != nil {

		if swag.IsZero(m.CleanPodPolicy) { // not required
			return nil
		}

		if err := m.CleanPodPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cleanPodPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cleanPodPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *V1TFJob) contextValidateEvaluator(ctx context.Context, formats strfmt.Registry) error {

	if m.Evaluator != nil {

		if swag.IsZero(m.Evaluator) { // not required
			return nil
		}

		if err := m.Evaluator.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("evaluator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("evaluator")
			}
			return err
		}
	}

	return nil
}

func (m *V1TFJob) contextValidatePs(ctx context.Context, formats strfmt.Registry) error {

	if m.Ps != nil {

		if swag.IsZero(m.Ps) { // not required
			return nil
		}

		if err := m.Ps.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ps")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ps")
			}
			return err
		}
	}

	return nil
}

func (m *V1TFJob) contextValidateSchedulingPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.SchedulingPolicy != nil {

		if swag.IsZero(m.SchedulingPolicy) { // not required
			return nil
		}

		if err := m.SchedulingPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedulingPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedulingPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *V1TFJob) contextValidateWorker(ctx context.Context, formats strfmt.Registry) error {

	if m.Worker != nil {

		if swag.IsZero(m.Worker) { // not required
			return nil
		}

		if err := m.Worker.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("worker")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("worker")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1TFJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TFJob) UnmarshalBinary(b []byte) error {
	var res V1TFJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
