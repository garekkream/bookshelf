package Settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/Sirupsen/logrus"
)

type ShelfList struct {
	Id     string `json:"shelfId"`
	Name   string `json:"shelfName"`
	Path   string `json:"shelfPath"`
	Active bool   `json:"shelfActive"`
}

type Config struct {
	ConfigPath string      `json:"configPath"`
	LogPath    string      `json:"logPath"`
	Debug      bool        `json:"debug"`
	Shelfs     []ShelfList `json:"shelfs"`
}

var config *Config
var log *logrus.Logger
var logFile *os.File

func init() {
	config = new(Config)
	log = logrus.New()
	formatter := new(logrus.TextFormatter)

	user, err := user.Current()
	if err != nil {
		fmt.Println("Failed to obrain user info!")
		return
	}

	config.ConfigPath = user.HomeDir + "/.config/BookShelf/config.json"
	config.LogPath = user.HomeDir + "/.config/BookShelf/bookshelf.log"

	if _, err = os.Stat(config.LogPath); err != nil {
		os.Create(config.LogPath)
	}

	if logFile, err = os.OpenFile(config.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
		fmt.Println(err)
		return
	}

	formatter.FullTimestamp = true

	log.Out = logFile
	log.Formatter = formatter
	log.Level = logrus.InfoLevel

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

	config.Debug = false

	WriteConfig()
}

func readConfig() {
	file, err := ioutil.ReadFile(config.ConfigPath)
	if err != nil {
		log.Debugln("Failed to open config file!")
		return
	}

	json.Unmarshal(file, config)
}

func DebugMode(mode bool) {
	switch mode {
	case true:
		log.Debugln("Debug mode enabled!")
		log.Level = logrus.DebugLevel
		break
	case false:
		log.Debugln("Debug mode disabled!")
		log.Level = logrus.InfoLevel
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

	log.Debug("ConfigPath set to: " + config.ConfigPath)
}

func GetConfigPath() string {
	return config.ConfigPath
}

func GetConfigDir() string {
	return path.Dir(config.ConfigPath)
}

func PrintConfig() {
	fmt.Println("Bookshelf Settings:")
	fmt.Printf("\tConfig path: \t%s\n", config.ConfigPath)
	fmt.Printf("\tLog path: \t%s\n", config.LogPath)
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
	for i := range config.Shelfs {
		config.Shelfs[i].Active = false
	}

	config.Shelfs[index].Active = true
}

func CloseLogFile() {
	logFile.Close()
}

func Log() *logrus.Logger {
	return log
}
