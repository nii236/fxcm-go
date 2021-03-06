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

// NewTradingTableSubscribeTradingTableParams creates a new TradingTableSubscribeTradingTableParams object
// with the default values initialized.
func NewTradingTableSubscribeTradingTableParams() *TradingTableSubscribeTradingTableParams {
	var ()
	return &TradingTableSubscribeTradingTableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTradingTableSubscribeTradingTableParamsWithTimeout creates a new TradingTableSubscribeTradingTableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTradingTableSubscribeTradingTableParamsWithTimeout(timeout time.Duration) *TradingTableSubscribeTradingTableParams {
	var ()
	return &TradingTableSubscribeTradingTableParams{

		timeout: timeout,
	}
}

// NewTradingTableSubscribeTradingTableParamsWithContext creates a new TradingTableSubscribeTradingTableParams object
// with the default values initialized, and the ability to set a context for a request
func NewTradingTableSubscribeTradingTableParamsWithContext(ctx context.Context) *TradingTableSubscribeTradingTableParams {
	var ()
	return &TradingTableSubscribeTradingTableParams{

		Context: ctx,
	}
}

// NewTradingTableSubscribeTradingTableParamsWithHTTPClient creates a new TradingTableSubscribeTradingTableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTradingTableSubscribeTradingTableParamsWithHTTPClient(client *http.Client) *TradingTableSubscribeTradingTableParams {
	var ()
	return &TradingTableSubscribeTradingTableParams{
		HTTPClient: client,
	}
}

/*TradingTableSubscribeTradingTableParams contains all the parameters to send to the API endpoint
for the trading table subscribe trading table operation typically these are written to a http.Request
*/
type TradingTableSubscribeTradingTableParams struct {

	/*Accept
	  Acceptable response MIME type

	*/
	Accept string
	/*Authorization
	  Authorization string containing “Bearer “, ID of socke.io connection and persistent token

	*/
	Authorization string
	/*ContentType
	  Media type of the request

	*/
	ContentType string
	/*UserAgent
	  Identification of the client software

	*/
	UserAgent string
	/*Body*/
	Body *models.SubscribeTradingTablesDefinition

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the trading table subscribe trading table params
func (o *TradingTableSubscribeTradingTableParams) WithTimeout(timeout time.Duration) *TradingTableSubscribeTradingTableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the trading table subscribe trading table params
func (o *TradingTableSubscribeTradingTableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the trading table subscribe trading table params
func (o *TradingTableSubscribeTradingTableParams) WithContext(ctx context.Context) *TradingTableSubscribeTradingTableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the trading table subscribe trading table params
func (o *TradingTableSubscribeTradingTableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the trading table subscribe trading table params
func (o *TradingTableSubscribeTradingTableParams) WithHTTPClient(client *http.Client) *TradingTableSubscribeTradingTableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the trading table subscribe trading table params
func (o *TradingTableSubscribeTradingTableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccept adds the accept to the trading table subscribe trading table params
func (o *TradingTableSubscribeTradingTableParams) WithAccept(accept string) *TradingTableSubscribeTradingTableParams {
	o.SetAccept(accept)
	return o
}

// SetAccept adds the accept to the trading table subscribe trading table params
func (o *TradingTableSubscribeTradingTableParams) SetAccept(accept string) {
	o.Accept = accept
}

// WithAuthorization adds the authorization to the trading table subscribe trading table params
func (o *TradingTableSubscribeTradingTableParams) WithAuthorization(authorization string) *TradingTableSubscribeTradingTableParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the trading table subscribe trading table params
func (o *TradingTableSubscribeTradingTableParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithContentType adds the contentType to the trading table subscribe trading table params
func (o *TradingTableSubscribeTradingTableParams) WithContentType(contentType string) *TradingTableSubscribeTradingTableParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the trading table subscribe trading table params
func (o *TradingTableSubscribeTradingTableParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithUserAgent adds the userAgent to the trading table subscribe trading table params
func (o *TradingTableSubscribeTradingTableParams) WithUserAgent(userAgent string) *TradingTableSubscribeTradingTableParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the trading table subscribe trading table params
func (o *TradingTableSubscribeTradingTableParams) SetUserAgent(userAgent string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the trading table subscribe trading table params
func (o *TradingTableSubscribeTradingTableParams) WithBody(body *models.SubscribeTradingTablesDefinition) *TradingTableSubscribeTradingTableParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the trading table subscribe trading table params
func (o *TradingTableSubscribeTradingTableParams) SetBody(body *models.SubscribeTradingTablesDefinition) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TradingTableSubscribeTradingTableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Accept
	if err := r.SetHeaderParam("Accept", o.Accept); err != nil {
		return err
	}

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// header param Content-Type
	if err := r.SetHeaderParam("Content-Type", o.ContentType); err != nil {
		return err
	}

	// header param User-Agent
	if err := r.SetHeaderParam("User-Agent", o.UserAgent); err != nil {
		return err
	}

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
