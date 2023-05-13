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

// V1HpParams Hp Matrix specification
//
// swagger:model v1HpParams
type V1HpParams struct {

	// choice
	Choice *V1HpChoice `json:"choice,omitempty"`

	// daterange
	Daterange *V1HpDateRange `json:"daterange,omitempty"`

	// datetimerange
	Datetimerange *V1HpDateTimeRange `json:"datetimerange,omitempty"`

	// geomspace
	Geomspace *V1HpGeomSpace `json:"geomspace,omitempty"`

	// linspace
	Linspace *V1HpLinSpace `json:"linspace,omitempty"`

	// lognormal
	Lognormal *V1HpLogNormal `json:"lognormal,omitempty"`

	// logspace
	Logspace *V1HpLogSpace `json:"logspace,omitempty"`

	// loguniform
	Loguniform *V1HpLogUniform `json:"loguniform,omitempty"`

	// normal
	Normal *V1HpNormal `json:"normal,omitempty"`

	// pchoice
	Pchoice *V1HpPChoice `json:"pchoice,omitempty"`

	// qlognormal
	Qlognormal *V1HpQLogNormal `json:"qlognormal,omitempty"`

	// qloguniform
	Qloguniform *V1HpQLogUniform `json:"qloguniform,omitempty"`

	// qnormal
	Qnormal *V1HpQNormal `json:"qnormal,omitempty"`

	// quniform
	Quniform *V1HpQUniform `json:"quniform,omitempty"`

	// range
	Range *V1HpRange `json:"range,omitempty"`

	// uniform
	Uniform *V1HpUniform `json:"uniform,omitempty"`
}

// Validate validates this v1 hp params
func (m *V1HpParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChoice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDaterange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatetimerange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeomspace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinspace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLognormal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogspace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoguniform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNormal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePchoice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQlognormal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQloguniform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQnormal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuniform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUniform(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1HpParams) validateChoice(formats strfmt.Registry) error {
	if swag.IsZero(m.Choice) { // not required
		return nil
	}

	if m.Choice != nil {
		if err := m.Choice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("choice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("choice")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) validateDaterange(formats strfmt.Registry) error {
	if swag.IsZero(m.Daterange) { // not required
		return nil
	}

	if m.Daterange != nil {
		if err := m.Daterange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("daterange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("daterange")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) validateDatetimerange(formats strfmt.Registry) error {
	if swag.IsZero(m.Datetimerange) { // not required
		return nil
	}

	if m.Datetimerange != nil {
		if err := m.Datetimerange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datetimerange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datetimerange")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) validateGeomspace(formats strfmt.Registry) error {
	if swag.IsZero(m.Geomspace) { // not required
		return nil
	}

	if m.Geomspace != nil {
		if err := m.Geomspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("geomspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("geomspace")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) validateLinspace(formats strfmt.Registry) error {
	if swag.IsZero(m.Linspace) { // not required
		return nil
	}

	if m.Linspace != nil {
		if err := m.Linspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("linspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("linspace")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) validateLognormal(formats strfmt.Registry) error {
	if swag.IsZero(m.Lognormal) { // not required
		return nil
	}

	if m.Lognormal != nil {
		if err := m.Lognormal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lognormal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lognormal")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) validateLogspace(formats strfmt.Registry) error {
	if swag.IsZero(m.Logspace) { // not required
		return nil
	}

	if m.Logspace != nil {
		if err := m.Logspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logspace")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) validateLoguniform(formats strfmt.Registry) error {
	if swag.IsZero(m.Loguniform) { // not required
		return nil
	}

	if m.Loguniform != nil {
		if err := m.Loguniform.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loguniform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("loguniform")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) validateNormal(formats strfmt.Registry) error {
	if swag.IsZero(m.Normal) { // not required
		return nil
	}

	if m.Normal != nil {
		if err := m.Normal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("normal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("normal")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) validatePchoice(formats strfmt.Registry) error {
	if swag.IsZero(m.Pchoice) { // not required
		return nil
	}

	if m.Pchoice != nil {
		if err := m.Pchoice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pchoice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pchoice")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) validateQlognormal(formats strfmt.Registry) error {
	if swag.IsZero(m.Qlognormal) { // not required
		return nil
	}

	if m.Qlognormal != nil {
		if err := m.Qlognormal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qlognormal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("qlognormal")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) validateQloguniform(formats strfmt.Registry) error {
	if swag.IsZero(m.Qloguniform) { // not required
		return nil
	}

	if m.Qloguniform != nil {
		if err := m.Qloguniform.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qloguniform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("qloguniform")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) validateQnormal(formats strfmt.Registry) error {
	if swag.IsZero(m.Qnormal) { // not required
		return nil
	}

	if m.Qnormal != nil {
		if err := m.Qnormal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qnormal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("qnormal")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) validateQuniform(formats strfmt.Registry) error {
	if swag.IsZero(m.Quniform) { // not required
		return nil
	}

	if m.Quniform != nil {
		if err := m.Quniform.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quniform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quniform")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) validateRange(formats strfmt.Registry) error {
	if swag.IsZero(m.Range) { // not required
		return nil
	}

	if m.Range != nil {
		if err := m.Range.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("range")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) validateUniform(formats strfmt.Registry) error {
	if swag.IsZero(m.Uniform) { // not required
		return nil
	}

	if m.Uniform != nil {
		if err := m.Uniform.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uniform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uniform")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 hp params based on the context it is used
func (m *V1HpParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChoice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDaterange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatetimerange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGeomspace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinspace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLognormal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogspace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLoguniform(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNormal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePchoice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQlognormal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQloguniform(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQnormal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuniform(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUniform(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1HpParams) contextValidateChoice(ctx context.Context, formats strfmt.Registry) error {

	if m.Choice != nil {
		if err := m.Choice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("choice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("choice")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) contextValidateDaterange(ctx context.Context, formats strfmt.Registry) error {

	if m.Daterange != nil {
		if err := m.Daterange.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("daterange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("daterange")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) contextValidateDatetimerange(ctx context.Context, formats strfmt.Registry) error {

	if m.Datetimerange != nil {
		if err := m.Datetimerange.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datetimerange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datetimerange")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) contextValidateGeomspace(ctx context.Context, formats strfmt.Registry) error {

	if m.Geomspace != nil {
		if err := m.Geomspace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("geomspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("geomspace")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) contextValidateLinspace(ctx context.Context, formats strfmt.Registry) error {

	if m.Linspace != nil {
		if err := m.Linspace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("linspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("linspace")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) contextValidateLognormal(ctx context.Context, formats strfmt.Registry) error {

	if m.Lognormal != nil {
		if err := m.Lognormal.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lognormal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lognormal")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) contextValidateLogspace(ctx context.Context, formats strfmt.Registry) error {

	if m.Logspace != nil {
		if err := m.Logspace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logspace")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) contextValidateLoguniform(ctx context.Context, formats strfmt.Registry) error {

	if m.Loguniform != nil {
		if err := m.Loguniform.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loguniform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("loguniform")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) contextValidateNormal(ctx context.Context, formats strfmt.Registry) error {

	if m.Normal != nil {
		if err := m.Normal.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("normal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("normal")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) contextValidatePchoice(ctx context.Context, formats strfmt.Registry) error {

	if m.Pchoice != nil {
		if err := m.Pchoice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pchoice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pchoice")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) contextValidateQlognormal(ctx context.Context, formats strfmt.Registry) error {

	if m.Qlognormal != nil {
		if err := m.Qlognormal.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qlognormal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("qlognormal")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) contextValidateQloguniform(ctx context.Context, formats strfmt.Registry) error {

	if m.Qloguniform != nil {
		if err := m.Qloguniform.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qloguniform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("qloguniform")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) contextValidateQnormal(ctx context.Context, formats strfmt.Registry) error {

	if m.Qnormal != nil {
		if err := m.Qnormal.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qnormal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("qnormal")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) contextValidateQuniform(ctx context.Context, formats strfmt.Registry) error {

	if m.Quniform != nil {
		if err := m.Quniform.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quniform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quniform")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) contextValidateRange(ctx context.Context, formats strfmt.Registry) error {

	if m.Range != nil {
		if err := m.Range.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("range")
			}
			return err
		}
	}

	return nil
}

func (m *V1HpParams) contextValidateUniform(ctx context.Context, formats strfmt.Registry) error {

	if m.Uniform != nil {
		if err := m.Uniform.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uniform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uniform")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1HpParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1HpParams) UnmarshalBinary(b []byte) error {
	var res V1HpParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
