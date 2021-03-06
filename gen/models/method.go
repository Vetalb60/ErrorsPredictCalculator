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

// Method method
//
// swagger:model Method
type Method struct {

	// method
	// Required: true
	Method *string `json:"method"`
}

// Validate validates this method
func (m *Method) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Method) validateMethod(formats strfmt.Registry) error {

	if err := validate.Required("method", "body", m.Method); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this method based on context it is used
func (m *Method) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Method) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Method) UnmarshalBinary(b []byte) error {
	var res Method
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
