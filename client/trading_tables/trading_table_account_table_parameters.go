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

// NewTradingTableAccountTableParams creates a new TradingTableAccountTableParams object
// with the default values initialized.
func NewTradingTableAccountTableParams() *TradingTableAccountTableParams {
	var ()
	return &TradingTableAccountTableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTradingTableAccountTableParamsWithTimeout creates a new TradingTableAccountTableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTradingTableAccountTableParamsWithTimeout(timeout time.Duration) *TradingTableAccountTableParams {
	var ()
	return &TradingTableAccountTableParams{

		timeout: timeout,
	}
}

// NewTradingTableAccountTableParamsWithContext creates a new TradingTableAccountTableParams object
// with the default values initialized, and the ability to set a context for a request
func NewTradingTableAccountTableParamsWithContext(ctx context.Context) *TradingTableAccountTableParams {
	var ()
	return &TradingTableAccountTableParams{

		Context: ctx,
	}
}

// NewTradingTableAccountTableParamsWithHTTPClient creates a new TradingTableAccountTableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTradingTableAccountTableParamsWithHTTPClient(client *http.Client) *TradingTableAccountTableParams {
	var ()
	return &TradingTableAccountTableParams{
		HTTPClient: client,
	}
}

/*TradingTableAccountTableParams contains all the parameters to send to the API endpoint
for the trading table account table operation typically these are written to a http.Request
*/
type TradingTableAccountTableParams struct {

	/*Body*/
	Body *models.GetTradingTablesSnapshotDefinition

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the trading table account table params
func (o *TradingTableAccountTableParams) WithTimeout(timeout time.Duration) *TradingTableAccountTableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the trading table account table params
func (o *TradingTableAccountTableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the trading table account table params
func (o *TradingTableAccountTableParams) WithContext(ctx context.Context) *TradingTableAccountTableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the trading table account table params
func (o *TradingTableAccountTableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the trading table account table params
func (o *TradingTableAccountTableParams) WithHTTPClient(client *http.Client) *TradingTableAccountTableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the trading table account table params
func (o *TradingTableAccountTableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the trading table account table params
func (o *TradingTableAccountTableParams) WithBody(body *models.GetTradingTablesSnapshotDefinition) *TradingTableAccountTableParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the trading table account table params
func (o *TradingTableAccountTableParams) SetBody(body *models.GetTradingTablesSnapshotDefinition) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TradingTableAccountTableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
