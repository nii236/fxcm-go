// Code generated by go-swagger; DO NOT EDIT.

package trading_tables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new trading tables API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for trading tables API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
TradingTableAccountTable snapshots of account table

the Accounts table that contains the trading account data such as funds used in trading, idle funds, profits/losses, certain account limitations, and so on.
*/
func (a *Client) TradingTableAccountTable(params *TradingTableAccountTableParams, authInfo runtime.ClientAuthInfoWriter) (*TradingTableAccountTableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTradingTableAccountTableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tradingTable/AccountTable",
		Method:             "GET",
		PathPattern:        "/trading/get_model&model=Account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TradingTableAccountTableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TradingTableAccountTableOK), nil

}

/*
TradingTableClosedPositions snapshots of closed position

the Closed Positions table that contains information about the positions closed during the current trading day such as realized profit/loss, charged commission, cumulative interest, and so on.
*/
func (a *Client) TradingTableClosedPositions(params *TradingTableClosedPositionsParams, authInfo runtime.ClientAuthInfoWriter) (*TradingTableClosedPositionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTradingTableClosedPositionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tradingTable/ClosedPositions",
		Method:             "GET",
		PathPattern:        "/trading/get_model&model=ClosedPosition",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TradingTableClosedPositionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TradingTableClosedPositionsOK), nil

}

/*
TradingTableOfferTable snapshots of offer

the Offers table that contains information about trading instruments, current prices, and high/low trading day prices.
*/
func (a *Client) TradingTableOfferTable(params *TradingTableOfferTableParams, authInfo runtime.ClientAuthInfoWriter) (*TradingTableOfferTableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTradingTableOfferTableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tradingTable/OfferTable",
		Method:             "GET",
		PathPattern:        "/trading/get_model&model=Offer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TradingTableOfferTableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TradingTableOfferTableOK), nil

}

/*
TradingTableOrdersTable snapshots of orders table

the Orders table that contains information about orders. The data is kept in this table until all the orders are executed
*/
func (a *Client) TradingTableOrdersTable(params *TradingTableOrdersTableParams, authInfo runtime.ClientAuthInfoWriter) (*TradingTableOrdersTableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTradingTableOrdersTableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tradingTable/OrdersTable",
		Method:             "GET",
		PathPattern:        "/trading/get_model&model=Order",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TradingTableOrdersTableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TradingTableOrdersTableOK), nil

}

/*
TradingTableSummaryTable snapshots of summary table

the Summary table that contains summarized information such as the average entry price, profit/loss, and so on for every instrument currently traded.
*/
func (a *Client) TradingTableSummaryTable(params *TradingTableSummaryTableParams, authInfo runtime.ClientAuthInfoWriter) (*TradingTableSummaryTableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTradingTableSummaryTableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tradingTable/SummaryTable",
		Method:             "GET",
		PathPattern:        "/trading/get_model&model=Summary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TradingTableSummaryTableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TradingTableSummaryTableOK), nil

}

/*
TradingTableSnapshot snapshots of open posisiont tables

In case continuous updates of the trading tables is not needed, it is possible to request a one-time snapshot. Gets current content snapshot of the specified data models.
*/
func (a *Client) TradingTableSnapshot(params *TradingTableSnapshotParams, authInfo runtime.ClientAuthInfoWriter) (*TradingTableSnapshotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTradingTableSnapshotParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tradingTable/snapshot",
		Method:             "GET",
		PathPattern:        "/trading/get_model&model=OpenPosition",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TradingTableSnapshotReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TradingTableSnapshotOK), nil

}

/*
TradingTableSubscribeTradingTable subscribes to trading tables

Subscribes to the updates of the data models. Update will be pushed to client via the socket.
*/
func (a *Client) TradingTableSubscribeTradingTable(params *TradingTableSubscribeTradingTableParams, authInfo runtime.ClientAuthInfoWriter) (*TradingTableSubscribeTradingTableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTradingTableSubscribeTradingTableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tradingTable/subscribeTradingTable",
		Method:             "POST",
		PathPattern:        "/trading/subscribe",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TradingTableSubscribeTradingTableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TradingTableSubscribeTradingTableOK), nil

}

/*
TradingTableUnsubscribeTradingTable unsubscribes from trading tables

Subscribes to the updates of the data models. Update will be pushed to client via the socket.
*/
func (a *Client) TradingTableUnsubscribeTradingTable(params *TradingTableUnsubscribeTradingTableParams, authInfo runtime.ClientAuthInfoWriter) (*TradingTableUnsubscribeTradingTableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTradingTableUnsubscribeTradingTableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tradingTable/unsubscribeTradingTable",
		Method:             "POST",
		PathPattern:        "/trading/unsubscribe",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TradingTableUnsubscribeTradingTableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TradingTableUnsubscribeTradingTableOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
