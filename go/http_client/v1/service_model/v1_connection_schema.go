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

// V1ConnectionSchema v1 connection schema
//
// swagger:model v1ConnectionSchema
type V1ConnectionSchema struct {

	// bucket connection
	BucketConnection *V1BucketConnection `json:"bucketConnection,omitempty"`

	// claim connection
	ClaimConnection *V1ClaimConnection `json:"claimConnection,omitempty"`

	// git connection
	GitConnection *V1GitConnection `json:"gitConnection,omitempty"`

	// host connection
	HostConnection *V1HostConnection `json:"hostConnection,omitempty"`

	// host path connection
	HostPathConnection *V1HostPathConnection `json:"hostPathConnection,omitempty"`
}

// Validate validates this v1 connection schema
func (m *V1ConnectionSchema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBucketConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClaimConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGitConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostPathConnection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ConnectionSchema) validateBucketConnection(formats strfmt.Registry) error {
	if swag.IsZero(m.BucketConnection) { // not required
		return nil
	}

	if m.BucketConnection != nil {
		if err := m.BucketConnection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bucketConnection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bucketConnection")
			}
			return err
		}
	}

	return nil
}

func (m *V1ConnectionSchema) validateClaimConnection(formats strfmt.Registry) error {
	if swag.IsZero(m.ClaimConnection) { // not required
		return nil
	}

	if m.ClaimConnection != nil {
		if err := m.ClaimConnection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("claimConnection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("claimConnection")
			}
			return err
		}
	}

	return nil
}

func (m *V1ConnectionSchema) validateGitConnection(formats strfmt.Registry) error {
	if swag.IsZero(m.GitConnection) { // not required
		return nil
	}

	if m.GitConnection != nil {
		if err := m.GitConnection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gitConnection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gitConnection")
			}
			return err
		}
	}

	return nil
}

func (m *V1ConnectionSchema) validateHostConnection(formats strfmt.Registry) error {
	if swag.IsZero(m.HostConnection) { // not required
		return nil
	}

	if m.HostConnection != nil {
		if err := m.HostConnection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostConnection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hostConnection")
			}
			return err
		}
	}

	return nil
}

func (m *V1ConnectionSchema) validateHostPathConnection(formats strfmt.Registry) error {
	if swag.IsZero(m.HostPathConnection) { // not required
		return nil
	}

	if m.HostPathConnection != nil {
		if err := m.HostPathConnection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostPathConnection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hostPathConnection")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 connection schema based on the context it is used
func (m *V1ConnectionSchema) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBucketConnection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClaimConnection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGitConnection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHostConnection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHostPathConnection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ConnectionSchema) contextValidateBucketConnection(ctx context.Context, formats strfmt.Registry) error {

	if m.BucketConnection != nil {
		if err := m.BucketConnection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bucketConnection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bucketConnection")
			}
			return err
		}
	}

	return nil
}

func (m *V1ConnectionSchema) contextValidateClaimConnection(ctx context.Context, formats strfmt.Registry) error {

	if m.ClaimConnection != nil {
		if err := m.ClaimConnection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("claimConnection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("claimConnection")
			}
			return err
		}
	}

	return nil
}

func (m *V1ConnectionSchema) contextValidateGitConnection(ctx context.Context, formats strfmt.Registry) error {

	if m.GitConnection != nil {
		if err := m.GitConnection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gitConnection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gitConnection")
			}
			return err
		}
	}

	return nil
}

func (m *V1ConnectionSchema) contextValidateHostConnection(ctx context.Context, formats strfmt.Registry) error {

	if m.HostConnection != nil {
		if err := m.HostConnection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostConnection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hostConnection")
			}
			return err
		}
	}

	return nil
}

func (m *V1ConnectionSchema) contextValidateHostPathConnection(ctx context.Context, formats strfmt.Registry) error {

	if m.HostPathConnection != nil {
		if err := m.HostPathConnection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostPathConnection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hostPathConnection")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ConnectionSchema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ConnectionSchema) UnmarshalBinary(b []byte) error {
	var res V1ConnectionSchema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
