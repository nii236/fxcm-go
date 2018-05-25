// Code generated by go-swagger; DO NOT EDIT.

package historical_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new historical data API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for historical data API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Candle candles offer id period id

This command will request historical price data.
*/
func (a *Client) Candle(params *CandleParams, authInfo runtime.ClientAuthInfoWriter) (*CandleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCandleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "candle",
		Method:             "GET",
		PathPattern:        "/candles/{offer_id}/{period_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CandleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CandleOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
