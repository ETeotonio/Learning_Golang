package main

import (
	"encoding/json"
	"fmt"

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
	}
	var msg interface{}
	for {
		err := websock.ReadJSON(&msg)
		if err != nil {
			panic(err)
		}
		encoded_json, err := json.Marshal(msg)
		msg_formatted := fmt.Sprintf("%s", encoded_json)
		convert_price_json(msg_formatted)
		break
	}
}

//https://finnhub.io/docs/api/symbol-search
// func ticker_search(apikey string, apiaddress string, company_names []string) {
// 	formatted_websocket_address := apiaddress + "?token=" + apikey + "/search?q="
// 	for company_name := range company_names {

// 	}
// }
