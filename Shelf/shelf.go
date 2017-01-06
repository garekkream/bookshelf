package Shelf

import (
	"fmt"
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
		s.ShelfPath = path
	} else {
		t := time.Now()
		file := fmt.Sprintf("%d_%d_%d_bookshelf.shelf", t.Year(), t.Month(), t.Day())

		s.ShelfPath = defaultPath + "/" + file
	}

	s.AddShelfToConfig()

	debugPrintln("New shelf: " + s.ShelfName + " in " + s.ShelfPath)
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

func (shelf *Shelf) AddShelfToConfig() {
	conf := Settings.GetConfig()
	conf.Shelfs = append(conf.Shelfs, Settings.ShelfList{shelf.GetName(), shelf.GetPath()})

	Settings.WriteConfig()
}
