package main

import (
	"fmt"

	"github.com/garekkream/BookShelf/Settings"
	_ "github.com/garekkream/BookShelf/Shelf"
	parser "gopkg.in/alecthomas/kingpin.v2"
)

var (
	ver  = "none"
	date = "note"

	debug = parser.Flag("debug", "Enable debug mode for this session").Bool()

	settings      = parser.Command("settings", "Settings manipulation command")
	settingsPrint = settings.Flag("print-config", "Print current settings").Short('p').Bool()
	settingsDebug = settings.Flag("set-debug", "Enable/Disable debug mode [y/n]").Short('d').String()
)

const (
	debugMarker = string("BookShelf::Main")
)

func debugPrintln(text string) {
	if Settings.GetDebugMode() {
		fmt.Println(debugMarker + text)
	}
}

func main() {
	var mode bool

	debugPrintln("Application started!")

	parser.Version(date + ver)

	switch parser.Parse() {
	case "settings":
		if *settingsPrint {
			Settings.PrintConfig()
		}

		if len(*settingsDebug) > 0 {
			switch *settingsDebug {
			case "y":
				mode = true
				break
			case "n":
				mode = false
				break
			default:
				fmt.Println("Unknown debug status! Setting debug mode to false!")
				mode = false
				break
			}
			Settings.DebugModeSave(mode)
		}

	}

	debugPrintln("Initialization completed!")

}
