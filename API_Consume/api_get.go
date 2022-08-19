package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

//https://finnhub.io/docs/api/websocket-trades
func getTickerPrice(apikey string, apiaddress string, tickers []string) {
	formatted_websocket_address := apiaddress + "?token=" + apikey
	websock, _, err := websocket.DefaultDialer.Dial(formatted_websocket_address, nil)
	if err != nil {
		panic(err)
	}
	defer websock.Close()
	for _, s := range tickers {
		msg, _ := json.Marshal(map[string]interface{}{"type": "subscribe", "symbol": s})
		websock.WriteMessage(websocket.TextMessage, msg)
		fmt.Printf("%s", msg)
	}
	var msg interface{}
	for {
		err := websock.ReadJSON(&msg)
		if err != nil {
			panic(err)
		}
		fmt.Println("Message from server ", msg)
		time.Sleep(120)
	}
}
