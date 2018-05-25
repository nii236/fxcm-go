// Code generated by go-swagger; DO NOT EDIT.

package trading_orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new trading orders API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for trading orders API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ChangeOrder changes order

This command will change an existing order that has not been executed.
*/
func (a *Client) ChangeOrder(params *ChangeOrderParams, authInfo runtime.ClientAuthInfoWriter) (*ChangeOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeOrderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeOrder",
		Method:             "POST",
		PathPattern:        "/trading/change_order",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChangeOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeOrderOK), nil

}

/*
ChangeOrderStopLimit changes order stop limit

This command will request the change of a stop loss or limit profit order attached to an entry order.
*/
func (a *Client) ChangeOrderStopLimit(params *ChangeOrderStopLimitParams, authInfo runtime.ClientAuthInfoWriter) (*ChangeOrderStopLimitOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeOrderStopLimitParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeOrderStopLimit",
		Method:             "POST",
		PathPattern:        "/trading/change_order_stop_limit",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChangeOrderStopLimitReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeOrderStopLimitOK), nil

}

/*
ChangeTradeStopLimit changes trade stop limit

This command will request the change of a stop loss or limit profit order attached to a trade.
*/
func (a *Client) ChangeTradeStopLimit(params *ChangeTradeStopLimitParams, authInfo runtime.ClientAuthInfoWriter) (*ChangeTradeStopLimitOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeTradeStopLimitParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeTradeStopLimit",
		Method:             "POST",
		PathPattern:        "/trading/change_trade_stop_limit",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChangeTradeStopLimitReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeTradeStopLimitOK), nil

}

/*
CloseTrade closes trade

This command will request immediate closure of a trade at the best available price.
*/
func (a *Client) CloseTrade(params *CloseTradeParams, authInfo runtime.ClientAuthInfoWriter) (*CloseTradeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCloseTradeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "closeTrade",
		Method:             "POST",
		PathPattern:        "/trading/close_trade",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CloseTradeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CloseTradeOK), nil

}

/*
CloseAllForSymbol closes all for a symbol

This command will request the closure of all positions for a specified security.
*/
func (a *Client) CloseAllForSymbol(params *CloseAllForSymbolParams, authInfo runtime.ClientAuthInfoWriter) (*CloseAllForSymbolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCloseAllForSymbolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "close_all_for_symbol",
		Method:             "POST",
		PathPattern:        "/trading/close_all_for_symbol",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CloseAllForSymbolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CloseAllForSymbolOK), nil

}

/*
CreateEntryOrder creates entry order

This command will request the creation of a standing order to be filled when market reaches the requested price.
*/
func (a *Client) CreateEntryOrder(params *CreateEntryOrderParams, authInfo runtime.ClientAuthInfoWriter) (*CreateEntryOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateEntryOrderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createEntryOrder",
		Method:             "POST",
		PathPattern:        "/trading/create_entry_order",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateEntryOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateEntryOrderOK), nil

}

/*
DeleteOrder deletes order

This command will request the removal of an existing order that has not been executed.
*/
func (a *Client) DeleteOrder(params *DeleteOrderParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteOrder",
		Method:             "POST",
		PathPattern:        "/trading/delete_order",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteOrderOK), nil

}

/*
EditOCO edits o c o

This command will request the addition or removal of an existing orders to/from an OCO group.
*/
func (a *Client) EditOCO(params *EditOCOParams, authInfo runtime.ClientAuthInfoWriter) (*EditOCOOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditOCOParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "editOCO",
		Method:             "POST",
		PathPattern:        "/trading/edit_oco",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EditOCOReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EditOCOOK), nil

}

/*
OpenTrade opens trade

This command will request immediate opening of a trade at the best available price.
*/
func (a *Client) OpenTrade(params *OpenTradeParams, authInfo runtime.ClientAuthInfoWriter) (*OpenTradeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOpenTradeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "openTrade",
		Method:             "POST",
		PathPattern:        "/trading/open_trade",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OpenTradeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OpenTradeOK), nil

}

/*
SimpleOCO simples o c o order

This command will request the creation of a pair of orders. Execution of one of the orders will cancel the other one.
*/
func (a *Client) SimpleOCO(params *SimpleOCOParams, authInfo runtime.ClientAuthInfoWriter) (*SimpleOCOOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSimpleOCOParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "simpleOCO",
		Method:             "POST",
		PathPattern:        "/trading/simple_oco",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SimpleOCOReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SimpleOCOOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
