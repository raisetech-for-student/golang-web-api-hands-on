package usecase

import (
	"context"
	"errors"

	"golang-web-api-hands-on/domain/model"
	"golang-web-api-hands-on/domain/repository"
)

type Book interface {
	Get(ctx context.Context,
		id string) (*model.Book, error)
}

type bookUseCase struct {
	bookRepo repository.Book
}

func NewBook(bookRepo repository.Book) Book {
	return &bookUseCase{
		bookRepo: bookRepo,
	}
}

func (b *bookUseCase) Get(ctx context.Context, id string) (*model.Book, error) {
	book := b.bookRepo.FindByID(id)
	if book == nil {
		return nil, errors.New("could not find the book")
	}

	return book, nil
}
