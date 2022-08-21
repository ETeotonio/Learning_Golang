package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Price_Ticker struct {
	Data []Price_Data `json: "data"`
}
type Price_Data struct {
	c string  `json: c`
	P float64 `json: p`
	S string  `json: s`
	T int64   `json: t`
	V float64 `json: v`
}

type Ticker struct {
	Count  int64             `json: "count"`
	Result []TickerStructure `json: "result"`
}

type TickerStructure struct {
	Description   string `json: "description"`
	DisplaySymbol string `json: "displaySymbol"`
	Symbol        string `json: "symbol"`
	type_ticker   string `json: "type"`
}

func convert_price_json(json_text string) {
	price_json := Price_Ticker{}
	err := json.Unmarshal([]byte(json_text), &price_json)
	err_val(err)
	for i := 0; i < len(price_json.Data); i++ {
		fmt.Println("Symbol: ", price_json.Data[i].S)
		fmt.Printf("Last Price: %f\r\n", price_json.Data[i].P)
		fmt.Printf("Volume: %f\r\n", price_json.Data[i].V)
		fmt.Println("Time: ", time.UnixMilli(price_json.Data[i].T))
	}

}

func convert_ticker_json(json_text string) {
	ticker_data := Ticker{}
	err := json.Unmarshal([]byte(json_text), &ticker_data)
	err_val(err)
	for i := 0; i < len(ticker_data.Result); i++ {
		fmt.Println("Description: ", ticker_data.Result[i].Description)
		fmt.Println("Display Symbol: ", ticker_data.Result[i].DisplaySymbol)
		fmt.Println("Ticker Code: ", ticker_data.Result[i].Symbol)
	}
}
