// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// License License for digging.
//
// swagger:model license
type License struct {

	// dig allowed
	// Required: true
	DigAllowed *Amount `json:"digAllowed"`

	// dig used
	// Required: true
	DigUsed *Amount `json:"digUsed"`

	// id
	// Required: true
	ID *int64 `json:"id"`
}

// Validate validates this license
func (m *License) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDigAllowed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDigUsed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *License) validateDigAllowed(formats strfmt.Registry) error {

	if err := validate.Required("digAllowed", "body", m.DigAllowed); err != nil {
		return err
	}

	if err := validate.Required("digAllowed", "body", m.DigAllowed); err != nil {
		return err
	}

	if m.DigAllowed != nil {
		if err := m.DigAllowed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("digAllowed")
			}
			return err
		}
	}

	return nil
}

func (m *License) validateDigUsed(formats strfmt.Registry) error {

	if err := validate.Required("digUsed", "body", m.DigUsed); err != nil {
		return err
	}

	if err := validate.Required("digUsed", "body", m.DigUsed); err != nil {
		return err
	}

	if m.DigUsed != nil {
		if err := m.DigUsed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("digUsed")
			}
			return err
		}
	}

	return nil
}

func (m *License) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this license based on the context it is used
func (m *License) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDigAllowed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDigUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *License) contextValidateDigAllowed(ctx context.Context, formats strfmt.Registry) error {

	if m.DigAllowed != nil {
		if err := m.DigAllowed.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("digAllowed")
			}
			return err
		}
	}

	return nil
}

func (m *License) contextValidateDigUsed(ctx context.Context, formats strfmt.Registry) error {

	if m.DigUsed != nil {
		if err := m.DigUsed.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("digUsed")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *License) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *License) UnmarshalBinary(b []byte) error {
	var res License
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
