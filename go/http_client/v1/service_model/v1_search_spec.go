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

// V1SearchSpec Search spec definition
//
// swagger:model v1SearchSpec
type V1SearchSpec struct {

	// Optional analytics specification
	Analytics *V1AnalyticsSpec `json:"analytics,omitempty"`

	// Search columns
	Columns string `json:"columns,omitempty"`

	// Compare Flags
	Compares string `json:"compares,omitempty"`

	// Optional events specification
	Events *V1DashboardSpec `json:"events,omitempty"`

	// Search group bys
	Groupby string `json:"groupby,omitempty"`

	// Optional heat fields
	Heat string `json:"heat,omitempty"`

	// Optional histograms specification
	Histograms interface{} `json:"histograms,omitempty"`

	// Search layout
	Layout string `json:"layout,omitempty"`

	// Limit size
	Limit int32 `json:"limit,omitempty"`

	// Offset value
	Offset int32 `json:"offset,omitempty"`

	// Search query
	Query string `json:"query,omitempty"`

	// Search sections
	Sections string `json:"sections,omitempty"`

	// Search sort
	Sort string `json:"sort,omitempty"`

	// Optional trends specification
	Trends interface{} `json:"trends,omitempty"`
}

// Validate validates this v1 search spec
func (m *V1SearchSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnalytics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SearchSpec) validateAnalytics(formats strfmt.Registry) error {
	if swag.IsZero(m.Analytics) { // not required
		return nil
	}

	if m.Analytics != nil {
		if err := m.Analytics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("analytics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("analytics")
			}
			return err
		}
	}

	return nil
}

func (m *V1SearchSpec) validateEvents(formats strfmt.Registry) error {
	if swag.IsZero(m.Events) { // not required
		return nil
	}

	if m.Events != nil {
		if err := m.Events.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("events")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("events")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 search spec based on the context it is used
func (m *V1SearchSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAnalytics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SearchSpec) contextValidateAnalytics(ctx context.Context, formats strfmt.Registry) error {

	if m.Analytics != nil {
		if err := m.Analytics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("analytics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("analytics")
			}
			return err
		}
	}

	return nil
}

func (m *V1SearchSpec) contextValidateEvents(ctx context.Context, formats strfmt.Registry) error {

	if m.Events != nil {
		if err := m.Events.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("events")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("events")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SearchSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SearchSpec) UnmarshalBinary(b []byte) error {
	var res V1SearchSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
