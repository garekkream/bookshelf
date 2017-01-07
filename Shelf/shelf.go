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
			_, err := os.Stat(n.Path)
			if err != nil {
				fmt.Printf("Failed to remove Shelf! File %s doesn't exists!\n", n.Path)
			} else {
				os.Remove(n.Path)

				conf.Shelfs[i] = conf.Shelfs[len(conf.Shelfs)-1]
				conf.Shelfs = conf.Shelfs[:len(conf.Shelfs)-1]

				debugPrintln("Removed Shelf: " + name)
			}
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
	conf.Shelfs = append(conf.Shelfs, Settings.ShelfList{shelf.GetName(), shelf.GetPath()})

	Settings.WriteConfig()
}
