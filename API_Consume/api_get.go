package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/websocket"
)

//https://finnhub.io/docs/api/websocket-trades
func getTickerPrice(apikey string, apiaddress string, tickers []string) {
	formatted_websocket_address := apiaddress + "?token=" + apikey
	websock, _, err := websocket.DefaultDialer.Dial(formatted_websocket_address, nil)
	err_val(err)
	defer websock.Close()
	for _, s := range tickers {
		msg, _ := json.Marshal(map[string]interface{}{"type": "subscribe", "symbol": s})
		websock.WriteMessage(websocket.TextMessage, msg)
	}
	var msg interface{}
	for {
		err := websock.ReadJSON(&msg)
		err_val(err)
		encoded_json, err := json.Marshal(msg)
		msg_formatted := fmt.Sprintf("%s", encoded_json)
		convert_price_json(msg_formatted)
		break
	}
}

//https://finnhub.io/docs/api/symbol-search
func get_ticker_codes(apikey string, httpaddr string, company_names []string) {
	for _, company_name := range company_names {
		http_address := httpaddr + company_name + "&token=" + apikey
		http_ret, err := http.Get(http_address)
		err_val(err)
		body, err := ioutil.ReadAll(http_ret.Body)
		err_val(err)
		convert_ticker_json(string(body))
	}
}
