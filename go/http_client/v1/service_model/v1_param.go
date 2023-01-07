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

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Param Param specification
//
// swagger:model v1Param
type V1Param struct {

	// A flag to signal to Polyaxon that this param is used with a connection
	Connection string `json:"connection,omitempty"`

	// A flag to signal to Polyaxon that this param should not be validated against io
	ContextOnly bool `json:"contextOnly,omitempty"`

	// Ref corresponds to a reference of an object
	Ref string `json:"ref,omitempty"`

	// A flag to signal to Polyaxon that this io must be tranformed to the environment variable passed
	ToEnv string `json:"toEnv,omitempty"`

	// A flag to signal to Polyaxon that this param must be tranformed to an init container
	ToInit bool `json:"toInit,omitempty"`

	// The value to pass, if no ref or search is passed then it corresponds to a literal value
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this v1 param
func (m *V1Param) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 param based on context it is used
func (m *V1Param) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1Param) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Param) UnmarshalBinary(b []byte) error {
	var res V1Param
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
