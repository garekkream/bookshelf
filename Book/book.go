package Book

import "github.com/garekkream/BookShelf/Settings"

type Book struct {
	BookID     int    `json:"id"`
	BookTitle  string `json:"title"`
	BookAuthor string `json:"author"`
}

const maxBooksCnt = 1024

func NewBook() *Book {
	return new(Book)
}

func AddBook(title string, author string) *Book {
	book := NewBook()

	book.Title(title)
	book.Author(author)

	return book
}

func (b *Book) Id(id int) {
	b.BookID = id
	Settings.Log().Debugf("Set book id to %d.\n", id)
}

func (b *Book) GetId() int {
	return b.BookID
}

func (b *Book) Title(title string) {
	b.BookTitle = title
	Settings.Log().Debugf("Set book title to %s.\n", title)
}

func (b *Book) GetTitle() string {
	return b.BookTitle
}

func (b *Book) Author(author string) {
	b.BookAuthor = author
	Settings.Log().Debugf("Set book title to %s.\n", author)
}

func (b *Book) GetAuthor() string {
	return b.BookAuthor
}

func GetMaxBooksCnt() int {
	return maxBooksCnt
}
