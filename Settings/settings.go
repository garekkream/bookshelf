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
	litesql
	mongodb
)

type ShelfList struct {
	Name   string `json:"shelfName"`
	Path   string `json:"shelfPath"`
	Active bool   `json:"shelfActive"`
}

type Config struct {
	ConfigPath string      `json:"configPath"`
	Debug      bool        `json:"debug"`
	DBEngine   int         `json:"db_engine"`
	Shelfs     []ShelfList `json:"shelfs"`
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

func WriteConfig() {
	b, _ := json.Marshal(*config)

	ioutil.WriteFile(config.ConfigPath, b, os.ModeAppend)
}

func createConfig() {
	index := strings.LastIndex(config.ConfigPath, "/")

	os.Mkdir(config.ConfigPath[:index], os.ModePerm)
	os.Create(config.ConfigPath)

	config.DBEngine = litesql
	config.Debug = false

	WriteConfig()
}

func readConfig() {
	file, err := ioutil.ReadFile(config.ConfigPath)
	if err != nil {
		fmt.Println(debugMarker + "Failed to open config file!")
		return
	}

	json.Unmarshal(file, config)
}

func debugPrintln(text string) {
	if GetDebugMode() {
		fmt.Println(debugMarker + text)
	}
}

func DebugMode(mode bool) {
	switch mode {
	case true:
		debugPrintln("Debug mode enabled!")
		break
	case false:
		debugPrintln("Debug mode disabled!")
		break
	}

	config.Debug = mode
}

func DebugModeSave(mode bool) {
	DebugMode(mode)

	WriteConfig()
}

func GetDebugMode() bool {
	return config.Debug
}

func ConfigPath(path string) {
	config.ConfigPath = path

	debugPrintln("ConfigPath set to: " + config.ConfigPath)
}

func GetConfigPath() string {
	return config.ConfigPath
}

func PrintConfig() {
	fmt.Println("Bookshelf Settings:")
	fmt.Printf("\tConfig path: \t%s\n", config.ConfigPath)
	fmt.Printf("\tDebug mode: \t%t\n", config.Debug)
	if len(config.Shelfs) != 0 {
		fmt.Printf("\tShelfs:\n")
		for _, shelf := range config.Shelfs {
			fmt.Printf("\t\t%s \t (%s) \t [Active = %t]\n", shelf.Name, shelf.Path, shelf.Active)
		}
	} else {
		fmt.Printf("\t\tNo shelfs available!\n")
	}
}

func GetConfig() *Config {
	return config
}

func ActivateShelf(index int) {
	for i, _ := range config.Shelfs {
		config.Shelfs[i].Active = false
	}

	config.Shelfs[index].Active = true
}
