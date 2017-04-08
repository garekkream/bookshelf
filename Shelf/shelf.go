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
	ShelfID   string      `json:"id"`
	ShelfName string      `json:"name"`
	ShelfPath string      `json:"path"`
	Books     []Book.Book `json:"books"`
	items     int
}

var (
	defaultPath  = Settings.GetConfigPath()[:strings.LastIndex(Settings.GetConfigPath(), "/")]
	currentShelf = new(Shelf)
)

func generateShelfID(chain string) string {
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

	s.ShelfID = generateShelfID(name + path)

	os.Create(s.ShelfPath)

	s.saveShelf()
	s.addShelfToConfig()

	ReadShelf(s.GetPath())

	Settings.Log().Debugln("New shelf: " + s.ShelfName + " in " + s.ShelfPath)

	return s.ShelfID
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

func isShelfListEmpty() (bool, error) {
	conf := Settings.GetConfig()

	if len(conf.Shelfs) > 0 {
		return false, nil
	}

	return true, fmt.Errorf("Shelf list is empty")
}

func activateFirstShelf(shelf Settings.ShelfList) {
	conf := Settings.GetConfig()

	if shelf.Active && len(conf.Shelfs) > 0 {
		conf.Shelfs[0].Active = true
	}
}

//DelShelfByName removes shelf by name provided by user
// if success function returns nil, otherwise error messeage
func DelShelfByName(name string) error {
	conf := Settings.GetConfig()

	_, err := isShelfListEmpty()
	if err != nil {
		return err
	}

	for i, n := range conf.Shelfs {
		if n.Name == name {
			activateFirstShelf(n)
			return delShelf(n.Path, i)
		}
	}

	return fmt.Errorf("Shelf %s not found", name)
}

//DelShelfByID removes shelf by id provided by user
// if success function returns nil, otherwise error messeage
func DelShelfByID(id string) error {
	conf := Settings.GetConfig()

	_, err := isShelfListEmpty()
	if err != nil {
		return err
	}

	for i, n := range conf.Shelfs {
		if n.Id == id {
			activateFirstShelf(n)
			return delShelf(n.Path, i)
		}
	}

	return fmt.Errorf("Shelf %s not found", id)
}

func delShelf(path string, index int) error {
	conf := Settings.GetConfig()

	_, err := os.Stat(path)
	if err != nil {
		Settings.Log().Warningf("Failed to remove Shelf! File %s doesn't exists! (err: %v)\n", path, err)
		return err
	}

	err = os.Remove(path)
	if err != nil {
		return err
	}

	conf.Shelfs[index] = conf.Shelfs[len(conf.Shelfs)-1]
	conf.Shelfs = conf.Shelfs[:len(conf.Shelfs)-1]

	Settings.Log().Debugln("Removed Shelf: " + path)
	Settings.WriteConfig()

	return nil
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

func (shelf *Shelf) GetId() string {
	return shelf.ShelfID
}

func (shelf *Shelf) addShelfToConfig() {
	conf := Settings.GetConfig()

	//When new shelf is available make it active by default
	for i := range conf.Shelfs {
		conf.Shelfs[i].Active = false
	}

	conf.Shelfs = append(conf.Shelfs, Settings.ShelfList{shelf.GetId(), shelf.GetName(), shelf.GetPath(), true})

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

//ListShelfs will print out all shelfs into console
func ListShelfs() error {
	_, err := isShelfListEmpty()
	if err != nil {
		Settings.Log().Errorln(err)
		return err
	}
	for _, n := range Settings.GetConfig().Shelfs {
		fmt.Printf("%s \t%s \t%s ", n.Id, n.Name, n.Path)
		if n.Active {
			fmt.Print("(Active)")
		}
		fmt.Print("\n")
	}

	return nil
}

func ActivateShelf(id string) {
	shelfs := Settings.GetConfig().Shelfs

	for i, n := range shelfs {
		if n.Id == id {
			shelfs[i].Active = true
			ReadShelf(shelfs[i].Path)
		} else {
			shelfs[i].Active = false
		}
	}

	Settings.WriteConfig()
}
