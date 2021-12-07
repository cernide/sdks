// Copyright 2018-2021 Polyaxon, Inc.
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

// V1OptimizationMetric OptimizationMetric specification
//
// swagger:model v1OptimizationMetric
type V1OptimizationMetric struct {

	// Name of the metric to optimize
	Name string `json:"name,omitempty"`

	// Optimization to use fot the metric
	Optimization *V1Optimization `json:"optimization,omitempty"`
}

// Validate validates this v1 optimization metric
func (m *V1OptimizationMetric) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOptimization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OptimizationMetric) validateOptimization(formats strfmt.Registry) error {
	if swag.IsZero(m.Optimization) { // not required
		return nil
	}

	if m.Optimization != nil {
		if err := m.Optimization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("optimization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("optimization")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 optimization metric based on the context it is used
func (m *V1OptimizationMetric) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOptimization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OptimizationMetric) contextValidateOptimization(ctx context.Context, formats strfmt.Registry) error {

	if m.Optimization != nil {
		if err := m.Optimization.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("optimization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("optimization")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1OptimizationMetric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OptimizationMetric) UnmarshalBinary(b []byte) error {
	var res V1OptimizationMetric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}