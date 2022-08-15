package main

import (
	"os"
)

type config_api struct {
	apiKey     string
	apiAddress string

	dbKey string
	dbUri string
}

func main() {
	if os.Args[1] == "createconfig" {
		createConfigFile()
	}
}

func createConfigFile() {
	config := "API:\n"
	config += "\tapiKey:\n"
	config += "\tapiAddress:\n"
	config += "DB:\n"
	config += "\tdbKey:\n"
	config += "\tdbUri:"
	err := os.WriteFile("config.yml", []byte(config), 0644)
	if err != nil {
		panic(err)
	}
}
