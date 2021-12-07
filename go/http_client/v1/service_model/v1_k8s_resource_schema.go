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

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1K8sResourceSchema K8S resource schema validation
//
// swagger:model v1K8sResourceSchema
type V1K8sResourceSchema struct {

	// default mode
	DefaultMode int32 `json:"defaultMode,omitempty"`

	// items
	Items []string `json:"items"`

	// mount path
	MountPath string `json:"mountPath,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this v1 k8s resource schema
func (m *V1K8sResourceSchema) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 k8s resource schema based on context it is used
func (m *V1K8sResourceSchema) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1K8sResourceSchema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1K8sResourceSchema) UnmarshalBinary(b []byte) error {
	var res V1K8sResourceSchema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}