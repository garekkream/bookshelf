package Shelf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/garekkream/BookShelf/Settings"
)

type Shelf struct {
	ShelfName string `json:"name"`
	ShelfPath string `json:"path"`
	items     int
}

const (
	debugMarker = string("BookShelf::Shelf: ")
)

var defaultPath = Settings.GetConfigPath()[:strings.LastIndex(Settings.GetConfigPath(), "/")]

func debugPrintln(text string) {
	if Settings.GetDebugMode() {
		fmt.Println(debugMarker + text)
	}
}

func NewShelf(name string, path string) {
	s := new(Shelf)

	if len(name) != 0 {
		s.Name(name)
	} else {
		debugPrintln("Missing new Shelf name!")
		s.Name("Bookshelf")
	}

	if len(path) != 0 {
		if filepath.IsAbs(path) {
			s.ShelfPath = path
		} else {
			p, _ := filepath.Abs(path)
			s.ShelfPath = p
		}
	} else {
		t := time.Now()
		file := fmt.Sprintf("%d_%d_%d_bookshelf.shelf", t.Year(), t.Month(), t.Day())

		s.ShelfPath = defaultPath + "/" + file
	}

	os.Create(s.ShelfPath)

	s.saveShelf()
	s.addShelfToConfig()

	debugPrintln("New shelf: " + s.ShelfName + " in " + s.ShelfPath)
}

func (shelf *Shelf) saveShelf() {
	b, _ := json.Marshal(shelf)

	ioutil.WriteFile(shelf.GetPath(), b, os.ModeAppend)
}

func DelShelf(name string) {
	conf := Settings.GetConfig()

	for i, n := range conf.Shelfs {
		if n.Name == name {

			// If removing currently active shelf, activate first one
			if n.Active && len(conf.Shelfs) > 0 {
				conf.Shelfs[0].Active = true
			}

			_, err := os.Stat(n.Path)
			if err != nil {
				fmt.Printf("Failed to remove Shelf! File %s doesn't exists!\n", n.Path)
			} else {
				os.Remove(n.Path)
			}

			conf.Shelfs[i] = conf.Shelfs[len(conf.Shelfs)-1]
			conf.Shelfs = conf.Shelfs[:len(conf.Shelfs)-1]

			debugPrintln("Removed Shelf: " + name)
			Settings.WriteConfig()

			return
		} else {
			debugPrintln("Failed to find shelf: " + name)
		}
	}
}

func (shelf *Shelf) Name(name string) {
	debugPrintln("Set Shelf name to: " + name)
	shelf.ShelfName = name
}

func (shelf *Shelf) GetName() string {
	return shelf.ShelfName
}

func (shelf *Shelf) GetPath() string {
	return shelf.ShelfPath
}

func (shelf *Shelf) addShelfToConfig() {
	conf := Settings.GetConfig()

	//When new shelf is available make it active by default
	for i, _ := range conf.Shelfs {
		conf.Shelfs[i].Active = false
	}

	conf.Shelfs = append(conf.Shelfs, Settings.ShelfList{shelf.GetName(), shelf.GetPath(), true})

	Settings.WriteConfig()
}
