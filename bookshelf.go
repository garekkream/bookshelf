package main

import (
	"os"

	"github.com/garekkream/bookshelf/Settings"
	parser "gopkg.in/alecthomas/kingpin.v2"
)

var (
	ver  = "none"
	date = "none"
)

func main() {
	Settings.Log().Infoln("Application started!")
	defer Settings.CloseLogFile()

	parser.Version(date + ver)

	if len(os.Args) < 2 {
		restInit()
		webkitInit()

		Settings.Log().Debugln("Initialization completed!")
	} else {
		cliParse()
	}
}
