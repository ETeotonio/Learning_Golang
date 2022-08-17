package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigApi struct {
	apiaddress string
	apikey     string
	dbkey      string
	dburi      string
}

var config_file = "config.yml"

func main() {
	_, err := os.Stat(config_file)
	if (os.Args[1] == "createconfig") || errors.Is(err, os.ErrNotExist) {
		fmt.Println("creating config file...")
		createConfigFile()
	} else {
		getConfig()
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
	config_file1, err := ioutil.ReadFile(config_file)
	if err != nil {
		panic(err)
	}
	// data := make(map[interface{}]interface{})
	data := make(map[string]ConfigApi)
	err2 := yaml.Unmarshal(config_file1, &data)
	if err2 != nil {
		panic(err2)
	}
	for k, v := range data {
		fmt.Print(k, v)
	}
}
