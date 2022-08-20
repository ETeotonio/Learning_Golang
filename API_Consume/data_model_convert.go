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

func convert_price_json(json_text string) {
	price_json := Price_Ticker{}
	err := json.Unmarshal([]byte(json_text), &price_json)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(price_json.Data); i++ {
		fmt.Println("Symbol: ", price_json.Data[i].S)
		fmt.Printf("Last Price: %f\r\n", price_json.Data[i].P)
		fmt.Printf("Volume: %f\r\n", price_json.Data[i].V)
		fmt.Println("Time: ", time.Unix(price_json.Data[i].T, 0))
	}

}
