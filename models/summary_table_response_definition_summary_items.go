// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SummaryTableResponseDefinitionSummaryItems summary table response definition summary items
// swagger:model summaryTableResponseDefinitionSummaryItems
type SummaryTableResponseDefinitionSummaryItems struct {

	// The sum of amounts of all positions in thousand units.
	AmountK float64 `json:"amountK,omitempty"`

	// The sum of amounts of Buy positions in thousand units.
	AmountKBuy float64 `json:"amountKBuy,omitempty"`

	// The sum of amounts of Sell positions in thousand units.
	AmountKSell float64 `json:"amountKSell,omitempty"`

	// The average open price of Buy positions.
	AvgBuy float64 `json:"avgBuy,omitempty"`

	// The average open price of Sell positions.
	AvgSell float64 `json:"avgSell,omitempty"`

	// The current market price, at which Sell positions can be closed.
	CloseBuy float64 `json:"closeBuy,omitempty"`

	// The current market price, at which Buy positions can be closed.
	CloseSell float64 `json:"closeSell,omitempty"`

	// The symbol of the instrument.
	Currency string `json:"currency,omitempty"`

	// UNKNOWN
	CurrencyPoint float64 `json:"currencyPoint,omitempty"`

	// The current profit/loss of all positions. It does not include commissions and interests.
	GrossPL float64 `json:"grossPL,omitempty"`

	// UNKNOWN
	IsBuyDisabled bool `json:"isBuyDisabled,omitempty"`

	// UNKNOWN
	IsSellDisabled bool `json:"isSellDisabled,omitempty"`

	// Indicates the row is a summary of for whole table.
	IsTotal bool `json:"isTotal,omitempty"`

	// UNKNOWN
	NetLimit float64 `json:"netLimit,omitempty"`

	// The current profit/loss of all positions. It includes commissions and interests.
	NetPL float64 `json:"netPL,omitempty"`

	// UNKNOWN
	NetStop float64 `json:"netStop,omitempty"`

	// UNKNOWN
	NetStopMove float64 `json:"netStopMove,omitempty"`

	// The unique identification number of the instrument.
	Offerid float64 `json:"offerid,omitempty"`

	// The current profit/loss of all Buy positions. It does not include commissions and interests.
	PlBuy float64 `json:"plBuy,omitempty"`

	// The current profit/loss of all Sell positions. It does not include commissions and interests.
	PlSell float64 `json:"plSell,omitempty"`

	// The price precision of the instrument. It defines number of digits after the decimal point in the instrument price quote.
	RatePrecision float64 `json:"ratePrecision,omitempty"`

	// The cumulative amount of funds that is added the account balance for holding the positions overnight.
	RollSum float64 `json:"rollSum,omitempty"`

	// ID number of the table
	T float64 `json:"t,omitempty"`

	// The amount of funds currently committed to maintain Buy positions.
	UsedMarginBuy float64 `json:"usedMarginBuy,omitempty"`

	// The amount of funds currently committed to maintain Sell positions.
	UsedMarginSell float64 `json:"usedMarginSell,omitempty"`
}

// Validate validates this summary table response definition summary items
func (m *SummaryTableResponseDefinitionSummaryItems) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SummaryTableResponseDefinitionSummaryItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SummaryTableResponseDefinitionSummaryItems) UnmarshalBinary(b []byte) error {
	var res SummaryTableResponseDefinitionSummaryItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}