// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SubscribeMarketDataResponseDefinitionResponse subscribe market data response definition response
// swagger:model subscribeMarketDataResponseDefinitionResponse
type SubscribeMarketDataResponseDefinitionResponse struct {

	// Type of socket
	Error string `json:"error,omitempty"`

	// Execution successful
	Executed bool `json:"executed,omitempty"`
}

// Validate validates this subscribe market data response definition response
func (m *SubscribeMarketDataResponseDefinitionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SubscribeMarketDataResponseDefinitionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscribeMarketDataResponseDefinitionResponse) UnmarshalBinary(b []byte) error {
	var res SubscribeMarketDataResponseDefinitionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
