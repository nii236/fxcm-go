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

// NewChangeTradeStopLimitParams creates a new ChangeTradeStopLimitParams object
// with the default values initialized.
func NewChangeTradeStopLimitParams() *ChangeTradeStopLimitParams {
	var ()
	return &ChangeTradeStopLimitParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeTradeStopLimitParamsWithTimeout creates a new ChangeTradeStopLimitParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeTradeStopLimitParamsWithTimeout(timeout time.Duration) *ChangeTradeStopLimitParams {
	var ()
	return &ChangeTradeStopLimitParams{

		timeout: timeout,
	}
}

// NewChangeTradeStopLimitParamsWithContext creates a new ChangeTradeStopLimitParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeTradeStopLimitParamsWithContext(ctx context.Context) *ChangeTradeStopLimitParams {
	var ()
	return &ChangeTradeStopLimitParams{

		Context: ctx,
	}
}

// NewChangeTradeStopLimitParamsWithHTTPClient creates a new ChangeTradeStopLimitParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeTradeStopLimitParamsWithHTTPClient(client *http.Client) *ChangeTradeStopLimitParams {
	var ()
	return &ChangeTradeStopLimitParams{
		HTTPClient: client,
	}
}

/*ChangeTradeStopLimitParams contains all the parameters to send to the API endpoint
for the change trade stop limit operation typically these are written to a http.Request
*/
type ChangeTradeStopLimitParams struct {

	/*Body*/
	Body models.ChangeTradeStopLimitRequestDefinition
	/*IsInPips
	  Defines if the trade‘s stop/limit rate is in pips.

	*/
	IsInPips *bool
	/*IsStop
	  Defines stop or limit should be changed (if true, then stop should be changed, otherwise limit).

	*/
	IsStop bool
	/*Rate
	  The trade‘s rate.

	*/
	Rate *float64
	/*TradeID
	  The trade identifier

	*/
	TradeID string
	/*TrailingStep
	  The trailing step for the stop rate.

	*/
	TrailingStep *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change trade stop limit params
func (o *ChangeTradeStopLimitParams) WithTimeout(timeout time.Duration) *ChangeTradeStopLimitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change trade stop limit params
func (o *ChangeTradeStopLimitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change trade stop limit params
func (o *ChangeTradeStopLimitParams) WithContext(ctx context.Context) *ChangeTradeStopLimitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change trade stop limit params
func (o *ChangeTradeStopLimitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change trade stop limit params
func (o *ChangeTradeStopLimitParams) WithHTTPClient(client *http.Client) *ChangeTradeStopLimitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change trade stop limit params
func (o *ChangeTradeStopLimitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change trade stop limit params
func (o *ChangeTradeStopLimitParams) WithBody(body models.ChangeTradeStopLimitRequestDefinition) *ChangeTradeStopLimitParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change trade stop limit params
func (o *ChangeTradeStopLimitParams) SetBody(body models.ChangeTradeStopLimitRequestDefinition) {
	o.Body = body
}

// WithIsInPips adds the isInPips to the change trade stop limit params
func (o *ChangeTradeStopLimitParams) WithIsInPips(isInPips *bool) *ChangeTradeStopLimitParams {
	o.SetIsInPips(isInPips)
	return o
}

// SetIsInPips adds the isInPips to the change trade stop limit params
func (o *ChangeTradeStopLimitParams) SetIsInPips(isInPips *bool) {
	o.IsInPips = isInPips
}

// WithIsStop adds the isStop to the change trade stop limit params
func (o *ChangeTradeStopLimitParams) WithIsStop(isStop bool) *ChangeTradeStopLimitParams {
	o.SetIsStop(isStop)
	return o
}

// SetIsStop adds the isStop to the change trade stop limit params
func (o *ChangeTradeStopLimitParams) SetIsStop(isStop bool) {
	o.IsStop = isStop
}

// WithRate adds the rate to the change trade stop limit params
func (o *ChangeTradeStopLimitParams) WithRate(rate *float64) *ChangeTradeStopLimitParams {
	o.SetRate(rate)
	return o
}

// SetRate adds the rate to the change trade stop limit params
func (o *ChangeTradeStopLimitParams) SetRate(rate *float64) {
	o.Rate = rate
}

// WithTradeID adds the tradeID to the change trade stop limit params
func (o *ChangeTradeStopLimitParams) WithTradeID(tradeID string) *ChangeTradeStopLimitParams {
	o.SetTradeID(tradeID)
	return o
}

// SetTradeID adds the tradeId to the change trade stop limit params
func (o *ChangeTradeStopLimitParams) SetTradeID(tradeID string) {
	o.TradeID = tradeID
}

// WithTrailingStep adds the trailingStep to the change trade stop limit params
func (o *ChangeTradeStopLimitParams) WithTrailingStep(trailingStep *float64) *ChangeTradeStopLimitParams {
	o.SetTrailingStep(trailingStep)
	return o
}

// SetTrailingStep adds the trailingStep to the change trade stop limit params
func (o *ChangeTradeStopLimitParams) SetTrailingStep(trailingStep *float64) {
	o.TrailingStep = trailingStep
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeTradeStopLimitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// query param is_stop
	qrIsStop := o.IsStop
	qIsStop := swag.FormatBool(qrIsStop)
	if qIsStop != "" {
		if err := r.SetQueryParam("is_stop", qIsStop); err != nil {
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

	// query param trade_id
	qrTradeID := o.TradeID
	qTradeID := qrTradeID
	if qTradeID != "" {
		if err := r.SetQueryParam("trade_id", qTradeID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
