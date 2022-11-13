package response

import (
	"net/http"
	"strconv"

	"dependency-injection-sample/domain/model"
)

type Book struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

func NewBook(book *model.Book) *Book {
	return &Book{
		ID:    strconv.Itoa(book.ID()),
		Name:  book.Name(),
		Price: strconv.Itoa(book.Price()),
	}
}

func (b *Book) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
