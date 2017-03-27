package Shelf

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/garekkream/bookshelf/Book"
	"github.com/garekkream/bookshelf/Settings"
)

type Shelf struct {
	ShelfId   string      `json:"id"`
	ShelfName string      `json:"name"`
	ShelfPath string      `json:"path"`
	Books     []Book.Book `json:"books"`
	items     int
}

var (
	defaultPath  = Settings.GetConfigPath()[:strings.LastIndex(Settings.GetConfigPath(), "/")]
	currentShelf = new(Shelf)
)

func generateShelfId(chain string) string {
	str := []byte(chain + string(time.Now().Minute()) + string(time.Now().Second()))

	sum := md5.New()
	sum.Write(str)

	return hex.EncodeToString(sum.Sum(nil)[:5])
}

func NewShelf(name string, path string) string {
	s := new(Shelf)

	if len(name) != 0 {
		s.Name(name)
	} else {
		Settings.Log().Warningln("Missing new Shelf name!")
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

	s.ShelfId = generateShelfId(name + path)

	os.Create(s.ShelfPath)

	s.saveShelf()
	s.addShelfToConfig()

	ReadShelf(s.GetPath())

	Settings.Log().Debugln("New shelf: " + s.ShelfName + " in " + s.ShelfPath)

	return s.ShelfId
}

func (shelf *Shelf) saveShelf() {
	b, _ := json.Marshal(shelf)

	ioutil.WriteFile(shelf.GetPath(), b, os.ModeAppend)
	Settings.Log().Debugf("Shelf %s saved as %s.\n", shelf.GetName(), shelf.GetPath())
}

func ReadShelf(path string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		Settings.Log().Errorln(err)
		return
	}

	json.Unmarshal(file, currentShelf)
}

func isShelfListEmpty() bool {
	conf := Settings.GetConfig()

	if len(conf.Shelfs) > 0 {
		return false
	} else {
		return true
	}
}

func DelShelf(name string) error {
	conf := Settings.GetConfig()

	if isShelfListEmpty() {
		errorStr := "Shelf list is empty!"

		Settings.Log().Errorln(errorStr)
		return fmt.Errorf(errorStr)
	}

	for i, n := range conf.Shelfs {
		if n.Name == name {
			// If removing currently active shelf, activate first one
			if n.Active && len(conf.Shelfs) > 0 {
				conf.Shelfs[0].Active = true
			}

			_, err := os.Stat(n.Path)
			if err != nil {
				Settings.Log().Warningf("Failed to remove Shelf! File %s doesn't exists! (err: %v)\n", n.Path, err)
				return err
			}

			os.Remove(n.Path)

			conf.Shelfs[i] = conf.Shelfs[len(conf.Shelfs)-1]
			conf.Shelfs = conf.Shelfs[:len(conf.Shelfs)-1]

			Settings.Log().Debugln("Removed Shelf: " + name)
			Settings.WriteConfig()

			return nil
		}
	}

	errorStr := "Failed to find shelf: " + name
	Settings.Log().Errorln(errorStr)
	return fmt.Errorf(errorStr)
}

func (shelf *Shelf) Name(name string) {
	Settings.Log().Debugln("Set Shelf name to: " + name)
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

	//When new shelf is available make it active by default
	for i := range conf.Shelfs {
		conf.Shelfs[i].Active = false
	}

	conf.Shelfs = append(conf.Shelfs, Settings.ShelfList{shelf.GetName(), shelf.GetPath(), true})

	Settings.WriteConfig()
}

func getActiveShelfPath() string {
	shelfs := Settings.GetConfig().Shelfs

	for i := range shelfs {
		if shelfs[i].Active == true {
			return shelfs[i].Path
		}
	}

	Settings.Log().Debugln("Failed to find active shelf!")
	return ""
}

func findFreeBookId() int {
	idx := len(currentShelf.Books)

	if (idx > 0) && (idx < Book.GetMaxBooksCnt()) {
		flag := false
		i := 0

		for i < Book.GetMaxBooksCnt() {
			for _, b := range currentShelf.Books {
				if b.GetId() == i {
					flag = true
					break
				}
			}

			if flag {
				i++
				flag = false
			} else {
				return i
			}
		}
	}

	Settings.Log().Errorln("Failed to find free ID!")
	return -1
}

func AddBookToShelf(book *Book.Book) {
	book.Id(findFreeBookId())
	currentShelf.Books = append(currentShelf.Books, *book)
}

func RemoveBookFromShelf(id int) {
	for i, book := range currentShelf.Books {
		if book.GetId() == id {
			l := len(currentShelf.Books)

			currentShelf.Books[i] = currentShelf.Books[l-1]
			currentShelf.Books = currentShelf.Books[:l-1]
		} else {
			Settings.Log().Errorf("Failed to find book with id = %d!\n", id)
		}
	}
}

func ListBooks() {
	fmt.Println(currentShelf.GetName())
	for _, book := range currentShelf.Books {
		fmt.Printf("\t%d \t%s \t%s\n", book.GetId(), book.GetTitle(), book.GetAuthor())
	}
}
