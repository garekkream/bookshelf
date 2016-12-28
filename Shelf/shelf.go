package Shelf

import (
	"fmt"
	"strings"

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
var shelfList = []Shelf{}

func debugPrintln(text string) {
	if Settings.GetDebugMode() == true {
		fmt.Println(debugMarker + text)
	}
}

func NewShelf(name string, path string) {
	s := new(Shelf)

	if len(name) < 1 {
		debugPrintln("Missing new Shelf name!")
	}

	if len(path) > 0 {
		s.ShelfPath = path
	} else {
		fileName := time.
	}
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
