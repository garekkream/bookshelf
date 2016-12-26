package main

import (
	"github.com/garekkream/BookShelf/Settings"
	parser "gopkg.in/alecthomas/kingpin.v2"
)

var (
	ver  = "none"
	date = "note"
)

func main() {
	Settings.DebugMode(true)

	parser.Version(date + ver)
	parser.Parse()
}
