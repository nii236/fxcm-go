// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenPositionResponseDefinitionOpenPositionItems open position response definition open position items
// swagger:model openPositionResponseDefinitionOpenPositionItems
type OpenPositionResponseDefinitionOpenPositionItems struct {

	// The unique identification number of the account the position is opened on. The number is unique within the database where the account is stored.
	AccountID string `json:"accountId,omitempty"`

	// The unique name of the account the position is opened on. The name is unique within the database where the account is stored.
	AccountName string `json:"accountName,omitempty"`

	// The amount of the position in thousand units.
	AmountK float64 `json:"amountK,omitempty"`

	// The price at which the position can be closed at the moment.
	Close float64 `json:"close,omitempty"`

	// The amount of funds subtracted from the account balance to pay for the broker's service in accordance with the terms and conditions of the account trading agreement.
	Com float64 `json:"com,omitempty"`

	// The symbol of the instrument.
	Currency string `json:"currency,omitempty"`

	// unknown
	CurrencyPoint float64 `json:"currencyPoint,omitempty"`

	// The current profit/loss of the position. It is expressed in the account currency.
	GrossPL float64 `json:"grossPL,omitempty"`

	// The trade operation the position is opened by. The possible values are
	// True – Buy
	// False – Sell
	//
	IsBuy float64 `json:"isBuy,omitempty"`

	// unknown
	IsDisabled bool `json:"isDisabled,omitempty"`

	// Indicates the row is a summary of for whole table.
	IsTotal bool `json:"isTotal,omitempty"`

	// The price of the associated limit order (profit limit level).
	Limit float64 `json:"limit,omitempty"`

	// The price the position is opened at.
	Open float64 `json:"open,omitempty"`

	// The price precision of the instrument. It defines number of digits after the decimal point in the instrument price quote.
	RatePrecision float64 `json:"ratePrecision,omitempty"`

	// The cumulative amount of funds that is added the account balance for holding the position overnight.
	Roll float64 `json:"roll,omitempty"`

	// The price of the associated stop order (loss limit level).
	Stop float64 `json:"stop,omitempty"`

	// The number of pips the market should move before the stop order moves the same number of pips after it. If the trailing order is dynamic (automatically updates every 0.1 of a pip), then the value of this field is 1. If the order is not trailing, the value of this field is 0.
	StopMove float64 `json:"stopMove,omitempty"`

	// ID number of the table
	T float64 `json:"t,omitempty"`

	// The date and time when the position was opened.
	Time string `json:"time,omitempty"`

	// The unique identification number of the open position. The number is unique within the same database that stores the account the position is opened on.
	TradeID string `json:"tradeId,omitempty"`

	// The amount of funds currently committed to maintain the position.
	UsedMargin float64 `json:"usedMargin,omitempty"`

	// The simulated delivery date. The date when the position could be automatically closed. The date is provided in the yyyyMMdd format. It is applicable only for positions opened on accounts with the day netting trading mode. Otherwise, the value of this field is blank.
	ValueDate string `json:"valueDate,omitempty"`

	// The current profit/loss per one lot of the position. It is expressed in the account currency.
	VisiblePL float64 `json:"visiblePL,omitempty"`
}

// Validate validates this open position response definition open position items
func (m *OpenPositionResponseDefinitionOpenPositionItems) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenPositionResponseDefinitionOpenPositionItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenPositionResponseDefinitionOpenPositionItems) UnmarshalBinary(b []byte) error {
	var res OpenPositionResponseDefinitionOpenPositionItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
