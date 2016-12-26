package Settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

const (
	clitext = iota
	litesql = iota
	mongodb = iota
)

type Config struct {
	ConfigPath string `json:"configPath"`
	Debug      bool   `json:"debug"`
	DBEngine   int    `json:"db_engine"`
}

var config *Config

const (
	debugMarker = string("BookShelf::Settings: ")
)

func init() {
	config = new(Config)

	user, err := user.Current()
	if err != nil {
		fmt.Println(debugMarker + "Failed to obrain user info!")
		return
	}

	config.ConfigPath = user.HomeDir + "/.config/BookShelf/config.json"

	_, err = os.Stat(config.ConfigPath)
	if err != nil {
		createConfig()
	} else {
		readConfig()
	}
}

func createConfig() {
	index := strings.LastIndex(config.ConfigPath, "/")

	os.Mkdir(config.ConfigPath[:index], os.ModePerm)
	os.Create(config.ConfigPath)

	config.DBEngine = litesql
	config.Debug = false

	b, _ := json.Marshal(*config)

	ioutil.WriteFile(config.ConfigPath, b, os.ModeAppend)
}

func readConfig() {
	file, err := ioutil.ReadFile(config.ConfigPath)
	if err != nil {
		fmt.Println(debugMarker + "Failed to open config file!")
	}

	json.Unmarshal(file, config)
}

func DebugMode(mode bool) {
	config.Debug = mode
}

func getDebugMode() bool {
	return config.Debug
}
