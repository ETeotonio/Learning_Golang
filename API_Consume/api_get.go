package main

import (
	"github.com/gorilla/websocket"
)

//https://finnhub.io/docs/api/websocket-trades
func getTickerPrice(apikey string, apiaddress string) {
	formatted_websocket_address := apiaddress + "?token=" + apikey
	websock, err := websocket.DefaultDialer.Dial(formatted_websocket_address, nil)
	if err != nil {
		panic(err)
	}
	defer websock.Close()
}
