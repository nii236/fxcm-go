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

// SubscribeTradingTablesDefinition subscribe trading tables definition
// swagger:model SubscribeTradingTablesDefinition
type SubscribeTradingTablesDefinition struct {

	// Name of the table model to be subscribed to
	Models string `json:"models,omitempty"`
}

// Validate validates this subscribe trading tables definition
func (m *SubscribeTradingTablesDefinition) Validate(formats strfmt.Registry) error {
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

var subscribeTradingTablesDefinitionTypeModelsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Offer","OpenPosition","ClosedPosition","Order","Account","Summary","LeverageProfile","Properties"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscribeTradingTablesDefinitionTypeModelsPropEnum = append(subscribeTradingTablesDefinitionTypeModelsPropEnum, v)
	}
}

const (
	// SubscribeTradingTablesDefinitionModelsOffer captures enum value "Offer"
	SubscribeTradingTablesDefinitionModelsOffer string = "Offer"
	// SubscribeTradingTablesDefinitionModelsOpenPosition captures enum value "OpenPosition"
	SubscribeTradingTablesDefinitionModelsOpenPosition string = "OpenPosition"
	// SubscribeTradingTablesDefinitionModelsClosedPosition captures enum value "ClosedPosition"
	SubscribeTradingTablesDefinitionModelsClosedPosition string = "ClosedPosition"
	// SubscribeTradingTablesDefinitionModelsOrder captures enum value "Order"
	SubscribeTradingTablesDefinitionModelsOrder string = "Order"
	// SubscribeTradingTablesDefinitionModelsAccount captures enum value "Account"
	SubscribeTradingTablesDefinitionModelsAccount string = "Account"
	// SubscribeTradingTablesDefinitionModelsSummary captures enum value "Summary"
	SubscribeTradingTablesDefinitionModelsSummary string = "Summary"
	// SubscribeTradingTablesDefinitionModelsLeverageProfile captures enum value "LeverageProfile"
	SubscribeTradingTablesDefinitionModelsLeverageProfile string = "LeverageProfile"
	// SubscribeTradingTablesDefinitionModelsProperties captures enum value "Properties"
	SubscribeTradingTablesDefinitionModelsProperties string = "Properties"
)

// prop value enum
func (m *SubscribeTradingTablesDefinition) validateModelsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subscribeTradingTablesDefinitionTypeModelsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SubscribeTradingTablesDefinition) validateModels(formats strfmt.Registry) error {

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
func (m *SubscribeTradingTablesDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscribeTradingTablesDefinition) UnmarshalBinary(b []byte) error {
	var res SubscribeTradingTablesDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
