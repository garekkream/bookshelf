package main

import (
	"github.com/garekkream/BookShelf/Settings"
	"github.com/garekkream/BookShelf/Shelf"
	parser "gopkg.in/alecthomas/kingpin.v2"
)

var (
	ver  = "none"
	date = "note"

	shelfNew  = parser.Flag("shelfNew", "Creates new shelf").Short('n').Bool()
	shelfName = parser.Flag("shelfName", "New shelf name").Default("").String()
	shelfPath = parser.Flag("shelfPath", "New shelf config path").Default("").String()
)

func main() {
	Settings.DebugMode(true)

	parser.Version(date + ver)
	parser.Parse()

	if *shelfNew {
		Shelf.NewShelf(*shelfName, *shelfPath)
	}
}
