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

// GetTradingTablesSnapshotDefinition get trading tables snapshot definition
// swagger:model GetTradingTablesSnapshotDefinition
type GetTradingTablesSnapshotDefinition struct {

	// Gets current content snapshot of the specified data models.
	Models string `json:"models,omitempty"`
}

// Validate validates this get trading tables snapshot definition
func (m *GetTradingTablesSnapshotDefinition) Validate(formats strfmt.Registry) error {
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

var getTradingTablesSnapshotDefinitionTypeModelsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Offer","OpenPosition","ClosedPosition","Order","Account","Summary","LeverageProfile","Properties"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getTradingTablesSnapshotDefinitionTypeModelsPropEnum = append(getTradingTablesSnapshotDefinitionTypeModelsPropEnum, v)
	}
}

const (
	// GetTradingTablesSnapshotDefinitionModelsOffer captures enum value "Offer"
	GetTradingTablesSnapshotDefinitionModelsOffer string = "Offer"
	// GetTradingTablesSnapshotDefinitionModelsOpenPosition captures enum value "OpenPosition"
	GetTradingTablesSnapshotDefinitionModelsOpenPosition string = "OpenPosition"
	// GetTradingTablesSnapshotDefinitionModelsClosedPosition captures enum value "ClosedPosition"
	GetTradingTablesSnapshotDefinitionModelsClosedPosition string = "ClosedPosition"
	// GetTradingTablesSnapshotDefinitionModelsOrder captures enum value "Order"
	GetTradingTablesSnapshotDefinitionModelsOrder string = "Order"
	// GetTradingTablesSnapshotDefinitionModelsAccount captures enum value "Account"
	GetTradingTablesSnapshotDefinitionModelsAccount string = "Account"
	// GetTradingTablesSnapshotDefinitionModelsSummary captures enum value "Summary"
	GetTradingTablesSnapshotDefinitionModelsSummary string = "Summary"
	// GetTradingTablesSnapshotDefinitionModelsLeverageProfile captures enum value "LeverageProfile"
	GetTradingTablesSnapshotDefinitionModelsLeverageProfile string = "LeverageProfile"
	// GetTradingTablesSnapshotDefinitionModelsProperties captures enum value "Properties"
	GetTradingTablesSnapshotDefinitionModelsProperties string = "Properties"
)

// prop value enum
func (m *GetTradingTablesSnapshotDefinition) validateModelsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getTradingTablesSnapshotDefinitionTypeModelsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetTradingTablesSnapshotDefinition) validateModels(formats strfmt.Registry) error {

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
func (m *GetTradingTablesSnapshotDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTradingTablesSnapshotDefinition) UnmarshalBinary(b []byte) error {
	var res GetTradingTablesSnapshotDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
