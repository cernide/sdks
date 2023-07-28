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

// V1PaddleJob Paddle Job specification
//
// swagger:model v1PaddleJob
type V1PaddleJob struct {

	// optional clean pod policy section
	CleanPodPolicy *V1CleanPodPolicy `json:"cleanPodPolicy,omitempty"`

	// Optional component kind, should be equal to 'paddlejob'
	Kind *string `json:"kind,omitempty"`

	// Master replicas definition
	Master *V1KFReplica `json:"master,omitempty"`

	// optional scheduling policy section
	SchedulingPolicy *V1SchedulingPolicy `json:"schedulingPolicy,omitempty"`

	// Worker replicas definition
	Worker *V1KFReplica `json:"worker,omitempty"`
}

// Validate validates this v1 paddle job
func (m *V1PaddleJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCleanPodPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaster(formats); err != nil {
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

func (m *V1PaddleJob) validateCleanPodPolicy(formats strfmt.Registry) error {
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

func (m *V1PaddleJob) validateMaster(formats strfmt.Registry) error {
	if swag.IsZero(m.Master) { // not required
		return nil
	}

	if m.Master != nil {
		if err := m.Master.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("master")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("master")
			}
			return err
		}
	}

	return nil
}

func (m *V1PaddleJob) validateSchedulingPolicy(formats strfmt.Registry) error {
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

func (m *V1PaddleJob) validateWorker(formats strfmt.Registry) error {
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

// ContextValidate validate this v1 paddle job based on the context it is used
func (m *V1PaddleJob) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCleanPodPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaster(ctx, formats); err != nil {
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

func (m *V1PaddleJob) contextValidateCleanPodPolicy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1PaddleJob) contextValidateMaster(ctx context.Context, formats strfmt.Registry) error {

	if m.Master != nil {

		if swag.IsZero(m.Master) { // not required
			return nil
		}

		if err := m.Master.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("master")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("master")
			}
			return err
		}
	}

	return nil
}

func (m *V1PaddleJob) contextValidateSchedulingPolicy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1PaddleJob) contextValidateWorker(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1PaddleJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PaddleJob) UnmarshalBinary(b []byte) error {
	var res V1PaddleJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
