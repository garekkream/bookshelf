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
	if Settings.GetDebugMode() == true {
		fmt.Println(debugMarker + text)
	}
}

func NewShelf(name string, path string) {
	s := new(Shelf)

	if len(name) < 1 {
		debugPrintln("Missing new Shelf name!")
		s.Name("Bookshelf")
	} else {
		s.Name(name)
	}

	if len(path) > 0 {
		s.ShelfPath = path
	} else {
		t := time.Now()
		file := fmt.Sprintf("%d_%d_%d_bookshelf.shelf", t.Year(), t.Month(), t.Day())

		s.ShelfPath = defaultPath + "/" + file
	}

	debugPrintln("New shelf: " + s.ShelfName + " in " + s.ShelfPath)
}

func (shelf *Shelf) Name(name string) {
	debugPrintln("Set Shelf name to: " + name)
	shelf.ShelfName = name
}

func (shelf *Shelf) GetName() string {
	return shelf.ShelfName
}

func (Shelf *Shelf) GetPath() string {
	return Shelf.ShelfPath
}
