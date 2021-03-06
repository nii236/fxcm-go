// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Error401Unauthorized Unauthorized
// swagger:model Error401_Unauthorized
type Error401Unauthorized struct {

	// code
	Code int32 `json:"Code,omitempty"`

	// The access token is not valid for this.
	Message string `json:"Message,omitempty"`
}

// Validate validates this error401 unauthorized
func (m *Error401Unauthorized) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Error401Unauthorized) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Error401Unauthorized) UnmarshalBinary(b []byte) error {
	var res Error401Unauthorized
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
