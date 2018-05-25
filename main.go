package main

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

func main() {
	clientws()
	// c := client.NewHTTPClient(strfmt.Default)
	// _, err := c.MarketData.MarketDataSubscribeMarketData(&market_data.MarketDataSubscribeMarketDataParams{}, runtime.ClientAuthInfoWriterFunc(auth))
	// if err != nil {
	// 	panic(err)
	// }
}

func auth(r runtime.ClientRequest, reg strfmt.Registry) error {
	return r.SetQueryParam("access_token", "HI")
}

func clientws() {
	fmt.Println("connect")
	//connect to server, you can use your own transport settings
	c, err := gosocketio.Dial(
		gosocketio.GetUrl("api-demo.fxcm.com/?access_token=hi", 443, true),
		transport.GetDefaultWebsocketTransport(),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("connected")
	//do something, handlers and functions are same as server ones

	//close connection
	c.Close()
	fmt.Println("closed")
}
