package main

import (
	"fmt"

	"github.com/garekkream/BookShelf/Settings"
	"github.com/garekkream/BookShelf/Shelf"

	parser "gopkg.in/alecthomas/kingpin.v2"
)

var (
	ver  = "none"
	date = "note"

	debug = parser.Flag("debug", "Enable debug mode for this session").Bool()

	settings      = parser.Command("settings", "Settings manipulation command")
	settingsPrint = settings.Flag("print-config", "Print current settings").Short('p').Bool()
	settingsDebug = settings.Flag("set-debug", "Enable/Disable debug mode [y/n]").Short('d').String()

	shelf       = parser.Command("shelf", "Shelf manipulation command")
	shelfNew    = shelf.Command("new", "Creates new shelf")
	shelfList   = shelf.Command("list", "List available shelfs")
	shelfDel    = shelf.Command("del", "Delete existing shelf")
	shelfActive = shelf.Command("active", "Activate selected shelf")
	shelfName   = shelf.Flag("name", "New shelf name").String()
	shelfPath   = shelf.Flag("path", "Storage path for new shelf").String()
	shelfIndex  = shelf.Flag("index", "Access shelfs using index").Int()
)

func init() {
	*shelfIndex = -1
}

func main() {
	var mode bool

	Settings.Log().Infoln("Application started!")

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
				Settings.Log().Warningln("Unknown debug status! Setting debug mode to false!")
				mode = false
				break
			}
			Settings.DebugModeSave(mode)
		}

	case "shelf new":
		var n string
		var p string

		if len(*shelfName) != 0 {
			n = *shelfName
		}

		if len(*shelfPath) != 0 {
			p = *shelfPath
		}

		Shelf.NewShelf(n, p)
		break

	case "shelf del":
		if len(*shelfName) != 0 {
			Shelf.DelShelf(*shelfName)
		} else {
			Settings.Log().Errorln("Failed to remove Shelf. Missing Shelf name!")
		}
		break

	case "shelf list":
		shelfs := Settings.GetConfig().Shelfs

		if len(shelfs) > 0 {
			for _, n := range Settings.GetConfig().Shelfs {
				fmt.Printf("%s %s\n", n.Name, n.Path)
			}
		} else {
			Settings.Log().Errorln("No Shelfs available!")
		}
		break

	case "shelf active":
		if len(*shelfName) != 0 {
			shelfs := Settings.GetConfig().Shelfs

			for i, n := range shelfs {
				if n.Name == *shelfName {
					Settings.ActivateShelf(i)
					Settings.WriteConfig()
					Shelf.ReadShelf(n.Path)
					break
				}
			}
			break
		}

		if *shelfIndex != -1 {
			shelfs := Settings.GetConfig().Shelfs

			if *shelfIndex < len(shelfs) &&
				*shelfIndex >= 0 {

				Settings.ActivateShelf(*shelfIndex)
			}

			Settings.WriteConfig()
			Shelf.ReadShelf(shelfs[*shelfIndex].Path)
			break
		}
	}

	Settings.Log().Debugln("Initialization completed!")
	Settings.CloseLogFile()
}
