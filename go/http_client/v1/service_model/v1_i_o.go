// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1IO Inputs/Outputs specification
//
// swagger:model v1IO
type V1IO struct {

	// An optional argFromat of the input/output to be used instead of passing the value as is
	ArgFormat string `json:"argFormat,omitempty"`

	// A flag to signal to Polyaxon that this io is used with a connection
	Connection string `json:"connection,omitempty"`

	// A flag to tell if param validation for this input/output should be delayed
	DelayValidation bool `json:"delayValidation,omitempty"`

	// Description for the input/output
	Description string `json:"description,omitempty"`

	// A flag to tell if this input/output is flag, only valid for bool type
	IsFlag bool `json:"isFlag,omitempty"`

	// A flag to tell if this input/output is list
	IsList bool `json:"isList,omitempty"`

	// A flag to tell if this input/output is optional
	IsOptional bool `json:"isOptional,omitempty"`

	// Name for the input/output
	Name string `json:"name,omitempty"`

	// An optional field to provide possible values for validation
	Options []interface{} `json:"options"`

	// A flag to signal to Polyaxon that this io must be tranformed to the environment variable passed
	ToEnv string `json:"toEnv,omitempty"`

	// A flag to signal to Polyaxon that this io must be tranformed to an init container
	ToInit bool `json:"toInit,omitempty"`

	// The type of the input/output
	Type string `json:"type,omitempty"`

	// The value of the input/output should be compatible with the type
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this v1 i o
func (m *V1IO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 i o based on context it is used
func (m *V1IO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1IO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IO) UnmarshalBinary(b []byte) error {
	var res V1IO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
