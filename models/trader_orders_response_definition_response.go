// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TraderOrdersResponseDefinitionResponse trader orders response definition response
// swagger:model traderOrdersResponseDefinitionResponse
type TraderOrdersResponseDefinitionResponse struct {

	// Execution successful
	Excuted bool `json:"excuted,omitempty"`
}

// Validate validates this trader orders response definition response
func (m *TraderOrdersResponseDefinitionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TraderOrdersResponseDefinitionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TraderOrdersResponseDefinitionResponse) UnmarshalBinary(b []byte) error {
	var res TraderOrdersResponseDefinitionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
