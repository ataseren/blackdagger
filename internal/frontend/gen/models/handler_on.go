// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HandlerOn handler on
//
// swagger:model handlerOn
type HandlerOn struct {

	// cancel
	Cancel *StepObject `json:"Cancel,omitempty"`

	// exit
	Exit *StepObject `json:"Exit,omitempty"`

	// failure
	Failure *StepObject `json:"Failure,omitempty"`

	// success
	Success *StepObject `json:"Success,omitempty"`
}

// Validate validates this handler on
func (m *HandlerOn) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCancel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HandlerOn) validateCancel(formats strfmt.Registry) error {
	if swag.IsZero(m.Cancel) { // not required
		return nil
	}

	if m.Cancel != nil {
		if err := m.Cancel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Cancel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Cancel")
			}
			return err
		}
	}

	return nil
}

func (m *HandlerOn) validateExit(formats strfmt.Registry) error {
	if swag.IsZero(m.Exit) { // not required
		return nil
	}

	if m.Exit != nil {
		if err := m.Exit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Exit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Exit")
			}
			return err
		}
	}

	return nil
}

func (m *HandlerOn) validateFailure(formats strfmt.Registry) error {
	if swag.IsZero(m.Failure) { // not required
		return nil
	}

	if m.Failure != nil {
		if err := m.Failure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Failure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Failure")
			}
			return err
		}
	}

	return nil
}

func (m *HandlerOn) validateSuccess(formats strfmt.Registry) error {
	if swag.IsZero(m.Success) { // not required
		return nil
	}

	if m.Success != nil {
		if err := m.Success.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Success")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Success")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this handler on based on the context it is used
func (m *HandlerOn) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCancel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFailure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSuccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HandlerOn) contextValidateCancel(ctx context.Context, formats strfmt.Registry) error {

	if m.Cancel != nil {

		if swag.IsZero(m.Cancel) { // not required
			return nil
		}

		if err := m.Cancel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Cancel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Cancel")
			}
			return err
		}
	}

	return nil
}

func (m *HandlerOn) contextValidateExit(ctx context.Context, formats strfmt.Registry) error {

	if m.Exit != nil {

		if swag.IsZero(m.Exit) { // not required
			return nil
		}

		if err := m.Exit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Exit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Exit")
			}
			return err
		}
	}

	return nil
}

func (m *HandlerOn) contextValidateFailure(ctx context.Context, formats strfmt.Registry) error {

	if m.Failure != nil {

		if swag.IsZero(m.Failure) { // not required
			return nil
		}

		if err := m.Failure.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Failure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Failure")
			}
			return err
		}
	}

	return nil
}

func (m *HandlerOn) contextValidateSuccess(ctx context.Context, formats strfmt.Registry) error {

	if m.Success != nil {

		if swag.IsZero(m.Success) { // not required
			return nil
		}

		if err := m.Success.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Success")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Success")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HandlerOn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HandlerOn) UnmarshalBinary(b []byte) error {
	var res HandlerOn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
