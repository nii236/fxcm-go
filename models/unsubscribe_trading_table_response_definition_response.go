// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UnsubscribeTradingTableResponseDefinitionResponse unsubscribe trading table response definition response
// swagger:model unsubscribeTradingTableResponseDefinitionResponse
type UnsubscribeTradingTableResponseDefinitionResponse struct {

	// Execution successful
	Excuted bool `json:"excuted,omitempty"`
}

// Validate validates this unsubscribe trading table response definition response
func (m *UnsubscribeTradingTableResponseDefinitionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UnsubscribeTradingTableResponseDefinitionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnsubscribeTradingTableResponseDefinitionResponse) UnmarshalBinary(b []byte) error {
	var res UnsubscribeTradingTableResponseDefinitionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
