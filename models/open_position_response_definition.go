// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenPositionResponseDefinition open position response definition
// swagger:model OpenPositionResponseDefinition
type OpenPositionResponseDefinition struct {

	// open position
	OpenPosition OpenPositionResponseDefinitionOpenPosition `json:"openPosition"`
}

// Validate validates this open position response definition
func (m *OpenPositionResponseDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenPositionResponseDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenPositionResponseDefinition) UnmarshalBinary(b []byte) error {
	var res OpenPositionResponseDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
