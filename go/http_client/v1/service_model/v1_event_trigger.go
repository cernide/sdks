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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1EventTrigger Event definition
//
// swagger:model v1EventTrigger
type V1EventTrigger struct {

	// The event kinds to subscribe to for the current reference
	Kinds []*V1EventKind `json:"kinds"`

	// Ref corresponds to a reference of an object
	Ref string `json:"ref,omitempty"`
}

// Validate validates this v1 event trigger
func (m *V1EventTrigger) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKinds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EventTrigger) validateKinds(formats strfmt.Registry) error {
	if swag.IsZero(m.Kinds) { // not required
		return nil
	}

	for i := 0; i < len(m.Kinds); i++ {
		if swag.IsZero(m.Kinds[i]) { // not required
			continue
		}

		if m.Kinds[i] != nil {
			if err := m.Kinds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kinds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("kinds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 event trigger based on the context it is used
func (m *V1EventTrigger) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKinds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EventTrigger) contextValidateKinds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Kinds); i++ {

		if m.Kinds[i] != nil {
			if err := m.Kinds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kinds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("kinds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1EventTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EventTrigger) UnmarshalBinary(b []byte) error {
	var res V1EventTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
