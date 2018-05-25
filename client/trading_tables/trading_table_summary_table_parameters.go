// Code generated by go-swagger; DO NOT EDIT.

package trading_tables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/nii236/fxcm-go/models"
)

// NewTradingTableSummaryTableParams creates a new TradingTableSummaryTableParams object
// with the default values initialized.
func NewTradingTableSummaryTableParams() *TradingTableSummaryTableParams {
	var ()
	return &TradingTableSummaryTableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTradingTableSummaryTableParamsWithTimeout creates a new TradingTableSummaryTableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTradingTableSummaryTableParamsWithTimeout(timeout time.Duration) *TradingTableSummaryTableParams {
	var ()
	return &TradingTableSummaryTableParams{

		timeout: timeout,
	}
}

// NewTradingTableSummaryTableParamsWithContext creates a new TradingTableSummaryTableParams object
// with the default values initialized, and the ability to set a context for a request
func NewTradingTableSummaryTableParamsWithContext(ctx context.Context) *TradingTableSummaryTableParams {
	var ()
	return &TradingTableSummaryTableParams{

		Context: ctx,
	}
}

// NewTradingTableSummaryTableParamsWithHTTPClient creates a new TradingTableSummaryTableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTradingTableSummaryTableParamsWithHTTPClient(client *http.Client) *TradingTableSummaryTableParams {
	var ()
	return &TradingTableSummaryTableParams{
		HTTPClient: client,
	}
}

/*TradingTableSummaryTableParams contains all the parameters to send to the API endpoint
for the trading table summary table operation typically these are written to a http.Request
*/
type TradingTableSummaryTableParams struct {

	/*Body*/
	Body *models.GetTradingTablesSnapshotDefinition

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the trading table summary table params
func (o *TradingTableSummaryTableParams) WithTimeout(timeout time.Duration) *TradingTableSummaryTableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the trading table summary table params
func (o *TradingTableSummaryTableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the trading table summary table params
func (o *TradingTableSummaryTableParams) WithContext(ctx context.Context) *TradingTableSummaryTableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the trading table summary table params
func (o *TradingTableSummaryTableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the trading table summary table params
func (o *TradingTableSummaryTableParams) WithHTTPClient(client *http.Client) *TradingTableSummaryTableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the trading table summary table params
func (o *TradingTableSummaryTableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the trading table summary table params
func (o *TradingTableSummaryTableParams) WithBody(body *models.GetTradingTablesSnapshotDefinition) *TradingTableSummaryTableParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the trading table summary table params
func (o *TradingTableSummaryTableParams) SetBody(body *models.GetTradingTablesSnapshotDefinition) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TradingTableSummaryTableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}