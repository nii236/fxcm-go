// Code generated by go-swagger; DO NOT EDIT.

package trading_orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/nii236/fxcm-go/models"
)

// NewSimpleOCOParams creates a new SimpleOCOParams object
// with the default values initialized.
func NewSimpleOCOParams() *SimpleOCOParams {
	var ()
	return &SimpleOCOParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSimpleOCOParamsWithTimeout creates a new SimpleOCOParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSimpleOCOParamsWithTimeout(timeout time.Duration) *SimpleOCOParams {
	var ()
	return &SimpleOCOParams{

		timeout: timeout,
	}
}

// NewSimpleOCOParamsWithContext creates a new SimpleOCOParams object
// with the default values initialized, and the ability to set a context for a request
func NewSimpleOCOParamsWithContext(ctx context.Context) *SimpleOCOParams {
	var ()
	return &SimpleOCOParams{

		Context: ctx,
	}
}

// NewSimpleOCOParamsWithHTTPClient creates a new SimpleOCOParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSimpleOCOParamsWithHTTPClient(client *http.Client) *SimpleOCOParams {
	var ()
	return &SimpleOCOParams{
		HTTPClient: client,
	}
}

/*SimpleOCOParams contains all the parameters to send to the API endpoint
for the simple o c o operation typically these are written to a http.Request
*/
type SimpleOCOParams struct {

	/*AccountID
	  The trade‘s account identifier

	*/
	AccountID string
	/*Amount
	  The trade‘s amount in lots.

	*/
	Amount string
	/*AtMarket
	  Defines the market range.

	*/
	AtMarket float64
	/*Body*/
	Body models.SimpleOCORequestDefinition
	/*Expiration
	  The order's expiration date.

	*/
	Expiration *string
	/*IsBuy
	  Defines the trade‘s market side (if true, then buy trade,otherwise sell trade). Temporarily not required by the server and defaults to true but this will change.

	*/
	IsBuy *bool
	/*IsBuy2
	  Defines the second trade‘s market side (if true, then buy trade, otherwise sell trade).

	*/
	IsBuy2 bool
	/*IsInPips
	  Defines if the trade‘s stop/limit rate is in pips.

	*/
	IsInPips *bool
	/*Limit
	  The trade‘s limit rate.

	*/
	Limit *float64
	/*Limit2
	  The second trade‘s limit rate.

	*/
	Limit2 float64
	/*OrderType
	  The type of the order execution. Market Order type choices “AtMarket”, “MarketRange”.

	*/
	OrderType string
	/*Rate
	  The trade‘s rate.

	*/
	Rate *float64
	/*Rate2
	  The second trade‘s rate.

	*/
	Rate2 float64
	/*Stop
	  The trade‘s stop rate.

	*/
	Stop *float64
	/*Stop2
	  The secondtrade‘s stop rate.

	*/
	Stop2 float64
	/*Symbol
	  The trade‘s account identifier

	*/
	Symbol string
	/*TimeInForce
	  Time in force choices “IOC”, “GTC”, “FOK”, “DAY”, “GTD”.

	*/
	TimeInForce string
	/*TrailingStep
	  The trailing step for the stop rate.

	*/
	TrailingStep *float64
	/*TrailingStep2
	  The second trailing step for the stop rate.

	*/
	TrailingStep2 float64
	/*TrailingStopStep
	  The trailing step for the stop rate.

	*/
	TrailingStopStep *float64
	/*TrailingStopStep2
	  The second trailing step for the stop rate.

	*/
	TrailingStopStep2 float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the simple o c o params
func (o *SimpleOCOParams) WithTimeout(timeout time.Duration) *SimpleOCOParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the simple o c o params
func (o *SimpleOCOParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the simple o c o params
func (o *SimpleOCOParams) WithContext(ctx context.Context) *SimpleOCOParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the simple o c o params
func (o *SimpleOCOParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the simple o c o params
func (o *SimpleOCOParams) WithHTTPClient(client *http.Client) *SimpleOCOParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the simple o c o params
func (o *SimpleOCOParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the simple o c o params
func (o *SimpleOCOParams) WithAccountID(accountID string) *SimpleOCOParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the simple o c o params
func (o *SimpleOCOParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithAmount adds the amount to the simple o c o params
func (o *SimpleOCOParams) WithAmount(amount string) *SimpleOCOParams {
	o.SetAmount(amount)
	return o
}

// SetAmount adds the amount to the simple o c o params
func (o *SimpleOCOParams) SetAmount(amount string) {
	o.Amount = amount
}

// WithAtMarket adds the atMarket to the simple o c o params
func (o *SimpleOCOParams) WithAtMarket(atMarket float64) *SimpleOCOParams {
	o.SetAtMarket(atMarket)
	return o
}

// SetAtMarket adds the atMarket to the simple o c o params
func (o *SimpleOCOParams) SetAtMarket(atMarket float64) {
	o.AtMarket = atMarket
}

// WithBody adds the body to the simple o c o params
func (o *SimpleOCOParams) WithBody(body models.SimpleOCORequestDefinition) *SimpleOCOParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the simple o c o params
func (o *SimpleOCOParams) SetBody(body models.SimpleOCORequestDefinition) {
	o.Body = body
}

// WithExpiration adds the expiration to the simple o c o params
func (o *SimpleOCOParams) WithExpiration(expiration *string) *SimpleOCOParams {
	o.SetExpiration(expiration)
	return o
}

// SetExpiration adds the expiration to the simple o c o params
func (o *SimpleOCOParams) SetExpiration(expiration *string) {
	o.Expiration = expiration
}

// WithIsBuy adds the isBuy to the simple o c o params
func (o *SimpleOCOParams) WithIsBuy(isBuy *bool) *SimpleOCOParams {
	o.SetIsBuy(isBuy)
	return o
}

// SetIsBuy adds the isBuy to the simple o c o params
func (o *SimpleOCOParams) SetIsBuy(isBuy *bool) {
	o.IsBuy = isBuy
}

// WithIsBuy2 adds the isBuy2 to the simple o c o params
func (o *SimpleOCOParams) WithIsBuy2(isBuy2 bool) *SimpleOCOParams {
	o.SetIsBuy2(isBuy2)
	return o
}

// SetIsBuy2 adds the isBuy2 to the simple o c o params
func (o *SimpleOCOParams) SetIsBuy2(isBuy2 bool) {
	o.IsBuy2 = isBuy2
}

// WithIsInPips adds the isInPips to the simple o c o params
func (o *SimpleOCOParams) WithIsInPips(isInPips *bool) *SimpleOCOParams {
	o.SetIsInPips(isInPips)
	return o
}

// SetIsInPips adds the isInPips to the simple o c o params
func (o *SimpleOCOParams) SetIsInPips(isInPips *bool) {
	o.IsInPips = isInPips
}

// WithLimit adds the limit to the simple o c o params
func (o *SimpleOCOParams) WithLimit(limit *float64) *SimpleOCOParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the simple o c o params
func (o *SimpleOCOParams) SetLimit(limit *float64) {
	o.Limit = limit
}

// WithLimit2 adds the limit2 to the simple o c o params
func (o *SimpleOCOParams) WithLimit2(limit2 float64) *SimpleOCOParams {
	o.SetLimit2(limit2)
	return o
}

// SetLimit2 adds the limit2 to the simple o c o params
func (o *SimpleOCOParams) SetLimit2(limit2 float64) {
	o.Limit2 = limit2
}

// WithOrderType adds the orderType to the simple o c o params
func (o *SimpleOCOParams) WithOrderType(orderType string) *SimpleOCOParams {
	o.SetOrderType(orderType)
	return o
}

// SetOrderType adds the orderType to the simple o c o params
func (o *SimpleOCOParams) SetOrderType(orderType string) {
	o.OrderType = orderType
}

// WithRate adds the rate to the simple o c o params
func (o *SimpleOCOParams) WithRate(rate *float64) *SimpleOCOParams {
	o.SetRate(rate)
	return o
}

// SetRate adds the rate to the simple o c o params
func (o *SimpleOCOParams) SetRate(rate *float64) {
	o.Rate = rate
}

// WithRate2 adds the rate2 to the simple o c o params
func (o *SimpleOCOParams) WithRate2(rate2 float64) *SimpleOCOParams {
	o.SetRate2(rate2)
	return o
}

// SetRate2 adds the rate2 to the simple o c o params
func (o *SimpleOCOParams) SetRate2(rate2 float64) {
	o.Rate2 = rate2
}

// WithStop adds the stop to the simple o c o params
func (o *SimpleOCOParams) WithStop(stop *float64) *SimpleOCOParams {
	o.SetStop(stop)
	return o
}

// SetStop adds the stop to the simple o c o params
func (o *SimpleOCOParams) SetStop(stop *float64) {
	o.Stop = stop
}

// WithStop2 adds the stop2 to the simple o c o params
func (o *SimpleOCOParams) WithStop2(stop2 float64) *SimpleOCOParams {
	o.SetStop2(stop2)
	return o
}

// SetStop2 adds the stop2 to the simple o c o params
func (o *SimpleOCOParams) SetStop2(stop2 float64) {
	o.Stop2 = stop2
}

// WithSymbol adds the symbol to the simple o c o params
func (o *SimpleOCOParams) WithSymbol(symbol string) *SimpleOCOParams {
	o.SetSymbol(symbol)
	return o
}

// SetSymbol adds the symbol to the simple o c o params
func (o *SimpleOCOParams) SetSymbol(symbol string) {
	o.Symbol = symbol
}

// WithTimeInForce adds the timeInForce to the simple o c o params
func (o *SimpleOCOParams) WithTimeInForce(timeInForce string) *SimpleOCOParams {
	o.SetTimeInForce(timeInForce)
	return o
}

// SetTimeInForce adds the timeInForce to the simple o c o params
func (o *SimpleOCOParams) SetTimeInForce(timeInForce string) {
	o.TimeInForce = timeInForce
}

// WithTrailingStep adds the trailingStep to the simple o c o params
func (o *SimpleOCOParams) WithTrailingStep(trailingStep *float64) *SimpleOCOParams {
	o.SetTrailingStep(trailingStep)
	return o
}

// SetTrailingStep adds the trailingStep to the simple o c o params
func (o *SimpleOCOParams) SetTrailingStep(trailingStep *float64) {
	o.TrailingStep = trailingStep
}

// WithTrailingStep2 adds the trailingStep2 to the simple o c o params
func (o *SimpleOCOParams) WithTrailingStep2(trailingStep2 float64) *SimpleOCOParams {
	o.SetTrailingStep2(trailingStep2)
	return o
}

// SetTrailingStep2 adds the trailingStep2 to the simple o c o params
func (o *SimpleOCOParams) SetTrailingStep2(trailingStep2 float64) {
	o.TrailingStep2 = trailingStep2
}

// WithTrailingStopStep adds the trailingStopStep to the simple o c o params
func (o *SimpleOCOParams) WithTrailingStopStep(trailingStopStep *float64) *SimpleOCOParams {
	o.SetTrailingStopStep(trailingStopStep)
	return o
}

// SetTrailingStopStep adds the trailingStopStep to the simple o c o params
func (o *SimpleOCOParams) SetTrailingStopStep(trailingStopStep *float64) {
	o.TrailingStopStep = trailingStopStep
}

// WithTrailingStopStep2 adds the trailingStopStep2 to the simple o c o params
func (o *SimpleOCOParams) WithTrailingStopStep2(trailingStopStep2 float64) *SimpleOCOParams {
	o.SetTrailingStopStep2(trailingStopStep2)
	return o
}

// SetTrailingStopStep2 adds the trailingStopStep2 to the simple o c o params
func (o *SimpleOCOParams) SetTrailingStopStep2(trailingStopStep2 float64) {
	o.TrailingStopStep2 = trailingStopStep2
}

// WriteToRequest writes these params to a swagger request
func (o *SimpleOCOParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param account_id
	qrAccountID := o.AccountID
	qAccountID := qrAccountID
	if qAccountID != "" {
		if err := r.SetQueryParam("account_id", qAccountID); err != nil {
			return err
		}
	}

	// query param amount
	qrAmount := o.Amount
	qAmount := qrAmount
	if qAmount != "" {
		if err := r.SetQueryParam("amount", qAmount); err != nil {
			return err
		}
	}

	// query param at_market
	qrAtMarket := o.AtMarket
	qAtMarket := swag.FormatFloat64(qrAtMarket)
	if qAtMarket != "" {
		if err := r.SetQueryParam("at_market", qAtMarket); err != nil {
			return err
		}
	}

	if o.Expiration != nil {

		// query param expiration
		var qrExpiration string
		if o.Expiration != nil {
			qrExpiration = *o.Expiration
		}
		qExpiration := qrExpiration
		if qExpiration != "" {
			if err := r.SetQueryParam("expiration", qExpiration); err != nil {
				return err
			}
		}

	}

	if o.IsBuy != nil {

		// query param is_buy
		var qrIsBuy bool
		if o.IsBuy != nil {
			qrIsBuy = *o.IsBuy
		}
		qIsBuy := swag.FormatBool(qrIsBuy)
		if qIsBuy != "" {
			if err := r.SetQueryParam("is_buy", qIsBuy); err != nil {
				return err
			}
		}

	}

	// query param is_buy2
	qrIsBuy2 := o.IsBuy2
	qIsBuy2 := swag.FormatBool(qrIsBuy2)
	if qIsBuy2 != "" {
		if err := r.SetQueryParam("is_buy2", qIsBuy2); err != nil {
			return err
		}
	}

	if o.IsInPips != nil {

		// query param is_in_pips
		var qrIsInPips bool
		if o.IsInPips != nil {
			qrIsInPips = *o.IsInPips
		}
		qIsInPips := swag.FormatBool(qrIsInPips)
		if qIsInPips != "" {
			if err := r.SetQueryParam("is_in_pips", qIsInPips); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit float64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatFloat64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	// query param limit2
	qrLimit2 := o.Limit2
	qLimit2 := swag.FormatFloat64(qrLimit2)
	if qLimit2 != "" {
		if err := r.SetQueryParam("limit2", qLimit2); err != nil {
			return err
		}
	}

	// query param order_type
	qrOrderType := o.OrderType
	qOrderType := qrOrderType
	if qOrderType != "" {
		if err := r.SetQueryParam("order_type", qOrderType); err != nil {
			return err
		}
	}

	if o.Rate != nil {

		// query param rate
		var qrRate float64
		if o.Rate != nil {
			qrRate = *o.Rate
		}
		qRate := swag.FormatFloat64(qrRate)
		if qRate != "" {
			if err := r.SetQueryParam("rate", qRate); err != nil {
				return err
			}
		}

	}

	// query param rate2
	qrRate2 := o.Rate2
	qRate2 := swag.FormatFloat64(qrRate2)
	if qRate2 != "" {
		if err := r.SetQueryParam("rate2", qRate2); err != nil {
			return err
		}
	}

	if o.Stop != nil {

		// query param stop
		var qrStop float64
		if o.Stop != nil {
			qrStop = *o.Stop
		}
		qStop := swag.FormatFloat64(qrStop)
		if qStop != "" {
			if err := r.SetQueryParam("stop", qStop); err != nil {
				return err
			}
		}

	}

	// query param stop2
	qrStop2 := o.Stop2
	qStop2 := swag.FormatFloat64(qrStop2)
	if qStop2 != "" {
		if err := r.SetQueryParam("stop2", qStop2); err != nil {
			return err
		}
	}

	// query param symbol
	qrSymbol := o.Symbol
	qSymbol := qrSymbol
	if qSymbol != "" {
		if err := r.SetQueryParam("symbol", qSymbol); err != nil {
			return err
		}
	}

	// query param time_in_force
	qrTimeInForce := o.TimeInForce
	qTimeInForce := qrTimeInForce
	if qTimeInForce != "" {
		if err := r.SetQueryParam("time_in_force", qTimeInForce); err != nil {
			return err
		}
	}

	if o.TrailingStep != nil {

		// query param trailing_step
		var qrTrailingStep float64
		if o.TrailingStep != nil {
			qrTrailingStep = *o.TrailingStep
		}
		qTrailingStep := swag.FormatFloat64(qrTrailingStep)
		if qTrailingStep != "" {
			if err := r.SetQueryParam("trailing_step", qTrailingStep); err != nil {
				return err
			}
		}

	}

	// query param trailing_step2
	qrTrailingStep2 := o.TrailingStep2
	qTrailingStep2 := swag.FormatFloat64(qrTrailingStep2)
	if qTrailingStep2 != "" {
		if err := r.SetQueryParam("trailing_step2", qTrailingStep2); err != nil {
			return err
		}
	}

	if o.TrailingStopStep != nil {

		// query param trailing_stop_step
		var qrTrailingStopStep float64
		if o.TrailingStopStep != nil {
			qrTrailingStopStep = *o.TrailingStopStep
		}
		qTrailingStopStep := swag.FormatFloat64(qrTrailingStopStep)
		if qTrailingStopStep != "" {
			if err := r.SetQueryParam("trailing_stop_step", qTrailingStopStep); err != nil {
				return err
			}
		}

	}

	// query param trailing_stop_step2
	qrTrailingStopStep2 := o.TrailingStopStep2
	qTrailingStopStep2 := swag.FormatFloat64(qrTrailingStopStep2)
	if qTrailingStopStep2 != "" {
		if err := r.SetQueryParam("trailing_stop_step2", qTrailingStopStep2); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
