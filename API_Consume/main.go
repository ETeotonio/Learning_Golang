package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

type configapi struct {
	apiaddress string
	apikey     string
	dbkey      string
	dburi      string
	httpaddr   string
	ticker     string
}

var config_file = "config.ini"
var confs configapi

func main() {
	_, err := os.Stat(config_file)
	if (os.Args[1] == "createconfig") || errors.Is(err, os.ErrNotExist) {
		fmt.Println("creating config file...")
		createConfigFile()
	} else {
		getConfig()
		// getTickerPrice(confs.apikey, confs.apiaddress, strings.Split(confs.ticker, ","))
		get_ticker_codes(confs.apikey, confs.httpaddr, strings.Split("Apple,Petrobras,Stellantis,Coffee", ","))
	}
}

func createConfigFile() {
	config := "[appconfig]"
	config += "apiKey:\n"
	config += "apiAddress:\n"
	config += "dbKey:\n"
	config += "dbUri:"
	err := os.WriteFile(config_file, []byte(config), 0644)
	err_val(err)
}

func getConfig() {
	cfg, err := ini.Load(config_file)
	err_val(err)
	confs.apiaddress = cfg.Section("appconfig").Key("apiaddress").String()
	confs.apikey = cfg.Section("appconfig").Key("apikey").String()
	confs.dbkey = cfg.Section("appconfig").Key("dbkey").String()
	confs.dburi = cfg.Section("appconfig").Key("dburi").String()
	confs.httpaddr = cfg.Section("appconfig").Key("http_address").String()
	confs.ticker = cfg.Section("tickers").Key("ticker_list").String()
}

func err_val(err error) {
	if err != nil {
		panic(err)
	}
}
