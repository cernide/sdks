// Copyright 2018-2023 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

// V1RunSchema v1 run schema
//
// swagger:model v1RunSchema
type V1RunSchema struct {

	// dag
	Dag *V1Dag `json:"dag,omitempty"`

	// dask
	Dask *V1Dask `json:"dask,omitempty"`

	// flink
	Flink *V1Flink `json:"flink,omitempty"`

	// job
	Job *V1Job `json:"job,omitempty"`

	// mpi job
	MpiJob *V1MPIJob `json:"mpiJob,omitempty"`

	// mx job
	MxJob *V1MXJob `json:"mxJob,omitempty"`

	// pytorch job
	PytorchJob *V1PytorchJob `json:"pytorchJob,omitempty"`

	// ruy
	Ruy *V1Ray `json:"ruy,omitempty"`

	// service
	Service *V1Service `json:"service,omitempty"`

	// spark
	Spark *V1Spark `json:"spark,omitempty"`

	// tf job
	TfJob *V1TFJob `json:"tfJob,omitempty"`

	// xgboost job
	XgboostJob *V1XGBoostJob `json:"xgboostJob,omitempty"`
}

// Validate validates this v1 run schema
func (m *V1RunSchema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDask(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMpiJob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMxJob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePytorchJob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpark(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTfJob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateXgboostJob(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1RunSchema) validateDag(formats strfmt.Registry) error {
	if swag.IsZero(m.Dag) { // not required
		return nil
	}

	if m.Dag != nil {
		if err := m.Dag.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dag")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dag")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) validateDask(formats strfmt.Registry) error {
	if swag.IsZero(m.Dask) { // not required
		return nil
	}

	if m.Dask != nil {
		if err := m.Dask.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dask")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dask")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) validateFlink(formats strfmt.Registry) error {
	if swag.IsZero(m.Flink) { // not required
		return nil
	}

	if m.Flink != nil {
		if err := m.Flink.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flink")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flink")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) validateJob(formats strfmt.Registry) error {
	if swag.IsZero(m.Job) { // not required
		return nil
	}

	if m.Job != nil {
		if err := m.Job.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("job")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("job")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) validateMpiJob(formats strfmt.Registry) error {
	if swag.IsZero(m.MpiJob) { // not required
		return nil
	}

	if m.MpiJob != nil {
		if err := m.MpiJob.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mpiJob")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mpiJob")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) validateMxJob(formats strfmt.Registry) error {
	if swag.IsZero(m.MxJob) { // not required
		return nil
	}

	if m.MxJob != nil {
		if err := m.MxJob.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mxJob")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mxJob")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) validatePytorchJob(formats strfmt.Registry) error {
	if swag.IsZero(m.PytorchJob) { // not required
		return nil
	}

	if m.PytorchJob != nil {
		if err := m.PytorchJob.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pytorchJob")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pytorchJob")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) validateRuy(formats strfmt.Registry) error {
	if swag.IsZero(m.Ruy) { // not required
		return nil
	}

	if m.Ruy != nil {
		if err := m.Ruy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ruy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ruy")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) validateService(formats strfmt.Registry) error {
	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) validateSpark(formats strfmt.Registry) error {
	if swag.IsZero(m.Spark) { // not required
		return nil
	}

	if m.Spark != nil {
		if err := m.Spark.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spark")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spark")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) validateTfJob(formats strfmt.Registry) error {
	if swag.IsZero(m.TfJob) { // not required
		return nil
	}

	if m.TfJob != nil {
		if err := m.TfJob.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tfJob")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tfJob")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) validateXgboostJob(formats strfmt.Registry) error {
	if swag.IsZero(m.XgboostJob) { // not required
		return nil
	}

	if m.XgboostJob != nil {
		if err := m.XgboostJob.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("xgboostJob")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("xgboostJob")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 run schema based on the context it is used
func (m *V1RunSchema) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDag(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDask(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFlink(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateJob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMpiJob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMxJob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePytorchJob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRuy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpark(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTfJob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateXgboostJob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1RunSchema) contextValidateDag(ctx context.Context, formats strfmt.Registry) error {

	if m.Dag != nil {
		if err := m.Dag.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dag")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dag")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) contextValidateDask(ctx context.Context, formats strfmt.Registry) error {

	if m.Dask != nil {
		if err := m.Dask.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dask")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dask")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) contextValidateFlink(ctx context.Context, formats strfmt.Registry) error {

	if m.Flink != nil {
		if err := m.Flink.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flink")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flink")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) contextValidateJob(ctx context.Context, formats strfmt.Registry) error {

	if m.Job != nil {
		if err := m.Job.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("job")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("job")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) contextValidateMpiJob(ctx context.Context, formats strfmt.Registry) error {

	if m.MpiJob != nil {
		if err := m.MpiJob.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mpiJob")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mpiJob")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) contextValidateMxJob(ctx context.Context, formats strfmt.Registry) error {

	if m.MxJob != nil {
		if err := m.MxJob.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mxJob")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mxJob")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) contextValidatePytorchJob(ctx context.Context, formats strfmt.Registry) error {

	if m.PytorchJob != nil {
		if err := m.PytorchJob.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pytorchJob")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pytorchJob")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) contextValidateRuy(ctx context.Context, formats strfmt.Registry) error {

	if m.Ruy != nil {
		if err := m.Ruy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ruy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ruy")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) contextValidateService(ctx context.Context, formats strfmt.Registry) error {

	if m.Service != nil {
		if err := m.Service.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) contextValidateSpark(ctx context.Context, formats strfmt.Registry) error {

	if m.Spark != nil {
		if err := m.Spark.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spark")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spark")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) contextValidateTfJob(ctx context.Context, formats strfmt.Registry) error {

	if m.TfJob != nil {
		if err := m.TfJob.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tfJob")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tfJob")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSchema) contextValidateXgboostJob(ctx context.Context, formats strfmt.Registry) error {

	if m.XgboostJob != nil {
		if err := m.XgboostJob.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("xgboostJob")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("xgboostJob")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1RunSchema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RunSchema) UnmarshalBinary(b []byte) error {
	var res V1RunSchema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
