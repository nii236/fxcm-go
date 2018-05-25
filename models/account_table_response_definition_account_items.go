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

// AccountTableResponseDefinitionAccountItems account table response definition account items
// swagger:model accountTableResponseDefinitionAccountItems
type AccountTableResponseDefinitionAccountItems struct {

	// The unique identification number of the account the position is opened on. The number is unique within the database where the account is stored.
	AccountID string `json:"accountId,omitempty"`

	// The unique name of the account the position is opened on. The name is unique within the database where the account is stored.
	AccountName string `json:"accountName,omitempty"`

	// The amount of funds on the account. This amount does not include floating profit and loss
	Balance float64 `json:"balance,omitempty"`

	// The amount of profits and losses (both floating and realized) of the current trading day.
	DayPL float64 `json:"dayPL,omitempty"`

	// The amount of funds on the account, including profits and losses of all open positions (the floating balance of the account).
	Equity float64 `json:"equity,omitempty"`

	// The amount of profits and losses of all open positions on the account.
	GrossPL float64 `json:"grossPL,omitempty"`

	// The type of the position maintenance. It defines how trade operations can be performed on the account. The possible values are
	// Y – Hedging is allowed. In other words, both buy and sell positions can be opened for the same instrument at the same time. To close each buy or sell position, an individual order is required.
	// N – Hedging is not allowed. In other words, either a buy or a sell position can be opened for the same instrument at a time. Opening a position for the instrument that already has open position(s) of the opposite trade operation always causes closing or partial closing of the open position(s).
	// 0 – Netting only. In other words, for each instrument there exists only one open position. The amount of the position is the total amount of the instrument, either bought or sold, that has not yet been offset by opposite trade operations.
	// D – Day netting. In other words, for each instrument there exists only one open position. Same as Netting only, but within a trading day. If the position is not offset during the same trading day it is opened, it is closed automatically on simulated delivery date.
	// F – FIFO. Positions open and close in accordance with the FIFO (First-in, First-out) rule. Hedging is not allowed.
	//
	Hedging string `json:"hedging,omitempty"`

	// Indicates the row is a summary of for whole table.
	IsTotal bool `json:"isTotal,omitempty"`

	// The limitation state of the account. Each state defines the operations that can be performed on the account. The possible values are
	// Y – Margin call (all positions are liquidated, new positions cannot be opened).
	// W – Warning of a possible margin call (positions may be closed, new positions cannot be opened).
	// Q – Equity stop (all positions are liquidated, new positions cannot be opened up to the end of the trading day).
	// A – Equity alert (positions may be closed, new positions cannot be opened up to the end of the trading day).
	// N – No limitations (no limitations are imposed on the account operations).
	//
	Mc string `json:"mc,omitempty"`

	// The price precision of the instrument. It defines number of digits after the decimal point in the instrument price quote.
	RatePrecision float64 `json:"ratePrecision,omitempty"`

	// ID number of the table
	T float64 `json:"t,omitempty"`

	// The amount of funds available to open new positions or to absorb losses of the existing positions.
	UsableMargin float64 `json:"usableMargin,omitempty"`

	// UNKNOWN
	UsableMargin3 float64 `json:"usableMargin3,omitempty"`

	// UNKNOWN
	UsableMargin3Perc float64 `json:"usableMargin3Perc,omitempty"`

	// UNKNOWN
	UsableMarginPerc float64 `json:"usableMarginPerc,omitempty"`

	// The amount of funds used to maintain all open positions on the account.
	UsdMr float64 `json:"usdMr,omitempty"`

	// The amount of funds used to maintain all open positions on the account with the three-level margin policy.
	UsdMr3 float64 `json:"usdMr3,omitempty"`
}

// Validate validates this account table response definition account items
func (m *AccountTableResponseDefinitionAccountItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHedging(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMc(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var accountTableResponseDefinitionAccountItemsTypeHedgingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Y","N","O","D","F"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountTableResponseDefinitionAccountItemsTypeHedgingPropEnum = append(accountTableResponseDefinitionAccountItemsTypeHedgingPropEnum, v)
	}
}

const (
	// AccountTableResponseDefinitionAccountItemsHedgingY captures enum value "Y"
	AccountTableResponseDefinitionAccountItemsHedgingY string = "Y"
	// AccountTableResponseDefinitionAccountItemsHedgingN captures enum value "N"
	AccountTableResponseDefinitionAccountItemsHedgingN string = "N"
	// AccountTableResponseDefinitionAccountItemsHedgingO captures enum value "O"
	AccountTableResponseDefinitionAccountItemsHedgingO string = "O"
	// AccountTableResponseDefinitionAccountItemsHedgingD captures enum value "D"
	AccountTableResponseDefinitionAccountItemsHedgingD string = "D"
	// AccountTableResponseDefinitionAccountItemsHedgingF captures enum value "F"
	AccountTableResponseDefinitionAccountItemsHedgingF string = "F"
)

// prop value enum
func (m *AccountTableResponseDefinitionAccountItems) validateHedgingEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, accountTableResponseDefinitionAccountItemsTypeHedgingPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccountTableResponseDefinitionAccountItems) validateHedging(formats strfmt.Registry) error {

	if swag.IsZero(m.Hedging) { // not required
		return nil
	}

	// value enum
	if err := m.validateHedgingEnum("hedging", "body", m.Hedging); err != nil {
		return err
	}

	return nil
}

var accountTableResponseDefinitionAccountItemsTypeMcPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Y","W","Q","A","N"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountTableResponseDefinitionAccountItemsTypeMcPropEnum = append(accountTableResponseDefinitionAccountItemsTypeMcPropEnum, v)
	}
}

const (
	// AccountTableResponseDefinitionAccountItemsMcY captures enum value "Y"
	AccountTableResponseDefinitionAccountItemsMcY string = "Y"
	// AccountTableResponseDefinitionAccountItemsMcW captures enum value "W"
	AccountTableResponseDefinitionAccountItemsMcW string = "W"
	// AccountTableResponseDefinitionAccountItemsMcQ captures enum value "Q"
	AccountTableResponseDefinitionAccountItemsMcQ string = "Q"
	// AccountTableResponseDefinitionAccountItemsMcA captures enum value "A"
	AccountTableResponseDefinitionAccountItemsMcA string = "A"
	// AccountTableResponseDefinitionAccountItemsMcN captures enum value "N"
	AccountTableResponseDefinitionAccountItemsMcN string = "N"
)

// prop value enum
func (m *AccountTableResponseDefinitionAccountItems) validateMcEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, accountTableResponseDefinitionAccountItemsTypeMcPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccountTableResponseDefinitionAccountItems) validateMc(formats strfmt.Registry) error {

	if swag.IsZero(m.Mc) { // not required
		return nil
	}

	// value enum
	if err := m.validateMcEnum("mc", "body", m.Mc); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountTableResponseDefinitionAccountItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountTableResponseDefinitionAccountItems) UnmarshalBinary(b []byte) error {
	var res AccountTableResponseDefinitionAccountItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
