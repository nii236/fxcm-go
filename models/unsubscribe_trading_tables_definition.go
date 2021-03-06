// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UnsubscribeTradingTablesDefinition unsubscribe trading tables definition
// swagger:model UnsubscribeTradingTablesDefinition
type UnsubscribeTradingTablesDefinition struct {

	// Name of the table model to be unsubscribed to
	Models string `json:"models,omitempty"`
}

// Validate validates this unsubscribe trading tables definition
func (m *UnsubscribeTradingTablesDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModels(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var unsubscribeTradingTablesDefinitionTypeModelsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Offer","OpenPosition","ClosedPosition","Order","Account","Summary","LeverageProfile","Properties"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		unsubscribeTradingTablesDefinitionTypeModelsPropEnum = append(unsubscribeTradingTablesDefinitionTypeModelsPropEnum, v)
	}
}

const (
	// UnsubscribeTradingTablesDefinitionModelsOffer captures enum value "Offer"
	UnsubscribeTradingTablesDefinitionModelsOffer string = "Offer"
	// UnsubscribeTradingTablesDefinitionModelsOpenPosition captures enum value "OpenPosition"
	UnsubscribeTradingTablesDefinitionModelsOpenPosition string = "OpenPosition"
	// UnsubscribeTradingTablesDefinitionModelsClosedPosition captures enum value "ClosedPosition"
	UnsubscribeTradingTablesDefinitionModelsClosedPosition string = "ClosedPosition"
	// UnsubscribeTradingTablesDefinitionModelsOrder captures enum value "Order"
	UnsubscribeTradingTablesDefinitionModelsOrder string = "Order"
	// UnsubscribeTradingTablesDefinitionModelsAccount captures enum value "Account"
	UnsubscribeTradingTablesDefinitionModelsAccount string = "Account"
	// UnsubscribeTradingTablesDefinitionModelsSummary captures enum value "Summary"
	UnsubscribeTradingTablesDefinitionModelsSummary string = "Summary"
	// UnsubscribeTradingTablesDefinitionModelsLeverageProfile captures enum value "LeverageProfile"
	UnsubscribeTradingTablesDefinitionModelsLeverageProfile string = "LeverageProfile"
	// UnsubscribeTradingTablesDefinitionModelsProperties captures enum value "Properties"
	UnsubscribeTradingTablesDefinitionModelsProperties string = "Properties"
)

// prop value enum
func (m *UnsubscribeTradingTablesDefinition) validateModelsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, unsubscribeTradingTablesDefinitionTypeModelsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UnsubscribeTradingTablesDefinition) validateModels(formats strfmt.Registry) error {

	if swag.IsZero(m.Models) { // not required
		return nil
	}

	// value enum
	if err := m.validateModelsEnum("models", "body", m.Models); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UnsubscribeTradingTablesDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnsubscribeTradingTablesDefinition) UnmarshalBinary(b []byte) error {
	var res UnsubscribeTradingTablesDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
