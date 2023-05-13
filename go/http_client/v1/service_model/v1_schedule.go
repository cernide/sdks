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

// V1Schedule All Schemas message
//
// swagger:model v1Schedule
type V1Schedule struct {

	// cron
	Cron *V1CronSchedule `json:"cron,omitempty"`

	// datetime
	Datetime *V1DateTimeSchedule `json:"datetime,omitempty"`

	// interval
	Interval *V1IntervalSchedule `json:"interval,omitempty"`
}

// Validate validates this v1 schedule
func (m *V1Schedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCron(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Schedule) validateCron(formats strfmt.Registry) error {
	if swag.IsZero(m.Cron) { // not required
		return nil
	}

	if m.Cron != nil {
		if err := m.Cron.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cron")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cron")
			}
			return err
		}
	}

	return nil
}

func (m *V1Schedule) validateDatetime(formats strfmt.Registry) error {
	if swag.IsZero(m.Datetime) { // not required
		return nil
	}

	if m.Datetime != nil {
		if err := m.Datetime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datetime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datetime")
			}
			return err
		}
	}

	return nil
}

func (m *V1Schedule) validateInterval(formats strfmt.Registry) error {
	if swag.IsZero(m.Interval) { // not required
		return nil
	}

	if m.Interval != nil {
		if err := m.Interval.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interval")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interval")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 schedule based on the context it is used
func (m *V1Schedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCron(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatetime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterval(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Schedule) contextValidateCron(ctx context.Context, formats strfmt.Registry) error {

	if m.Cron != nil {
		if err := m.Cron.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cron")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cron")
			}
			return err
		}
	}

	return nil
}

func (m *V1Schedule) contextValidateDatetime(ctx context.Context, formats strfmt.Registry) error {

	if m.Datetime != nil {
		if err := m.Datetime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datetime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datetime")
			}
			return err
		}
	}

	return nil
}

func (m *V1Schedule) contextValidateInterval(ctx context.Context, formats strfmt.Registry) error {

	if m.Interval != nil {
		if err := m.Interval.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interval")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interval")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Schedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Schedule) UnmarshalBinary(b []byte) error {
	var res V1Schedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
