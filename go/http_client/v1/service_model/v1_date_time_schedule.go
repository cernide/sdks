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
	"github.com/go-openapi/validate"
)

// V1DateTimeSchedule DateTime schedule specification
//
// swagger:model v1DateTimeSchedule
type V1DateTimeSchedule struct {

	// Kind of schedule, should be equal to "datetime"
	Kind *string `json:"kind,omitempty"`

	// Whan to start this cron schedule
	// Format: date-time
	StartAt strfmt.DateTime `json:"startAt,omitempty"`
}

// Validate validates this v1 date time schedule
func (m *V1DateTimeSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStartAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1DateTimeSchedule) validateStartAt(formats strfmt.Registry) error {
	if swag.IsZero(m.StartAt) { // not required
		return nil
	}

	if err := validate.FormatOf("startAt", "body", "date-time", m.StartAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 date time schedule based on context it is used
func (m *V1DateTimeSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1DateTimeSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DateTimeSchedule) UnmarshalBinary(b []byte) error {
	var res V1DateTimeSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
