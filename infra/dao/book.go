package dao

import (
	"dependency-injection-sample/domain/model"
	"dependency-injection-sample/domain/repository"
)

type bookRepository struct{}

func NewBook() repository.Book {
	return &bookRepository{}
}

func (b *bookRepository) FindByID(id string) *model.Book {
	if id == "1" {
		book := model.NewBook(1, "The Lord of the Rings", 1600)
		return &book
	} else if id == "2" {
		book := model.NewBook(2, "Harry Potter and the Philosopher's Stone, Book 1", 1200)
		return &book
	} else if id == "3" {
		book := model.NewBook(3, "The Chronicles of Narnia: 1", 4000)
		return &book
	} else if id == "4" {
		book := model.NewBook(4, "The Saga of Darren Shan", 1500)
		return &book
	} else {
		return nil
	}
}
