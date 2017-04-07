package main

import (
	"os"

	"github.com/garekkream/bookshelf/Settings"
	"github.com/garekkream/bookshelf/Shelf"

	parser "gopkg.in/alecthomas/kingpin.v2"
)

var (
	ver  = "none"
	date = "none"
)

type settingsParams struct {
	cmd           *parser.CmdClause
	printSettings *bool
	debugMode     *string
}

type shelfsParams struct {
	cmd    *parser.CmdClause
	new    *parser.CmdClause
	list   *parser.CmdClause
	del    *parser.CmdClause
	active *parser.CmdClause
	name   *string
	id     *string
	path   *string
	index  *int
}

type booksParams struct {
	cmd    *parser.CmdClause
	new    *parser.CmdClause
	del    *parser.CmdClause
	list   *parser.CmdClause
	title  *string
	author *string
	ID     *int
}

func cliParse() {
	setting := new(settingsParams)
	shelf := new(shelfsParams)
	book := new(booksParams)

	// *shelf.index = -1
	// *book.ID = -1

	// debug := parser.Flag("debug", "Enable debug mode for this session").Bool()

	setting.cmd = parser.Command("settings", "Settings manipulation command")
	setting.printSettings = setting.cmd.Flag("print-config", "Print current settings").Short('p').Bool()
	setting.debugMode = setting.cmd.Flag("set-debug", "Enable/Disable debug mode [y/n]").Short('d').String()

	shelf.cmd = parser.Command("shelf", "Shelf manipulation command")
	shelf.new = shelf.cmd.Command("new", "Creates new shelf")
	shelf.list = shelf.cmd.Command("list", "List available shelfs")
	shelf.del = shelf.cmd.Command("del", "Delete existing shelf")
	shelf.active = shelf.cmd.Command("active", "Activate selected shelf")
	shelf.name = shelf.cmd.Flag("name", "Shelf name").String()
	shelf.id = shelf.cmd.Flag("name", "Shelf id").String()
	shelf.path = shelf.cmd.Flag("path", "Storage path for new shelf").String()
	shelf.index = shelf.cmd.Flag("index", "Access shelfs using index").Int()

	book.cmd = parser.Command("book", "Book manipulation command")
	book.new = book.cmd.Command("new", "Create new book in active shelf")
	book.del = book.cmd.Command("del", "Delete book from active shelf")
	book.list = book.cmd.Command("list", "List all books in active shelf")
	book.title = book.cmd.Flag("title", "Set book title").Short('t').String()
	book.author = book.cmd.Flag("author", "Set book author").Short('a').String()
	book.ID = book.cmd.Flag("id", "Access book using id").Short('i').Int()

	switch parser.Parse() {
	case "settings":
		cliParseSettingsHndl(setting)
		break
	case "shelf new":
		cliParseShelfNewHndl(shelf)
		break
	case "shelf del":
		cliParseShelfDelHndl(shelf)
		break
	}
}

func cliParseSettingsHndl(s *settingsParams) {
	var mode bool

	if *s.printSettings {
		Settings.PrintConfig()
	}

	if len(*s.debugMode) > 0 {
		switch *s.debugMode {
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
}

func cliParseShelfNewHndl(s *shelfsParams) {
	var n string
	var p string

	if len(*s.name) != 0 {
		n = *s.name
	}

	if len(*s.path) != 0 {
		p = *s.path
	}

	Shelf.NewShelf(n, p)
}

func cliParseShelfDelHndl(s *shelfsParams) {
	if len(*s.name) > 0 {
		Shelf.DelShelfByName(*s.name)
		return
	}

	if len(*s.id) > 0 {
		Shelf.DelShelfByName(*s.id)
		return
	}

	Settings.Log().Errorln("Failed to remove Shelf. Missing shelf name!")
}

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
		// case "shelf list":
		// 	shelfs := Settings.GetConfig().Shelfs
		//
		// 	if len(shelfs) > 0 {
		// 		for _, n := range Settings.GetConfig().Shelfs {
		// 			fmt.Printf("%s %s\n", n.Name, n.Path)
		// 		}
		// 	} else {
		// 		Settings.Log().Errorln("No Shelfs available!")
		// 	}
		// 	break
		//
		// case "shelf active":
		// 	if len(*shelfName) != 0 {
		// 		shelfs := Settings.GetConfig().Shelfs
		//
		// 		for i, n := range shelfs {
		// 			if n.Name == *shelfName {
		// 				Settings.ActivateShelf(i)
		// 				Settings.WriteConfig()
		// 				Shelf.ReadShelf(n.Path)
		// 				break
		// 			}
		// 		}
		// 		break
		// 	}
		//
		// 	if *shelfIndex != -1 {
		// 		shelfs := Settings.GetConfig().Shelfs
		//
		// 		if *shelfIndex < len(shelfs) &&
		// 			*shelfIndex >= 0 {
		//
		// 			Settings.ActivateShelf(*shelfIndex)
		// 		}
		//
		// 		Settings.WriteConfig()
		// 		Shelf.ReadShelf(shelfs[*shelfIndex].Path)
		// 		break
		// 	}
		//
		// case "book new":
		// 	b := Book.AddBook(*bookTitle, *bookAuthor)
		// 	Shelf.AddBookToShelf(b)
		//
		// case "book del":
		// 	if *bookId != -1 {
		// 		Shelf.RemoveBookFromShelf(*bookId)
		// 	}
		// 	Settings.Log().Error("Unable to remove book! Book id missing!")
		//
		// case "book list":
		// 	Shelf.ListBooks()
	}
}
