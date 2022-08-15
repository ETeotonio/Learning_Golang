package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type configapi struct {
	apiKey     string
	apiAddress string

	dbKey string
	dbUri string
}

var data = make(map[string]configapi)
var config_file = "config.yml"

func main() {
	_, err := os.Stat(config_file)
	if (os.Args[1] == "createconfig") || errors.Is(err, os.ErrNotExist) {
		fmt.Println("creating config file...")
		createConfigFile()
	} else {

	}
}

func createConfigFile() {
	config := "apiKey:\n"
	config += "apiAddress:\n"
	config += "dbKey:\n"
	config += "dbUri:"
	err := os.WriteFile(config_file, []byte(config), 0644)
	if err != nil {
		panic(err)
	}
}

func getConfig() {
	config_file, err := ioutil.ReadFile(config_file)
	if err != nil {
		panic(err)
	}
	err2 := yaml.Unmarshal(config_file, &data)
	if err2 != nil {
		panic(err2)
	}

}

// func getTickerPrice string(ticker string){
// 	return ""
// }
