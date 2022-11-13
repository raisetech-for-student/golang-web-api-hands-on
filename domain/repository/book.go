package repository

import "dependency-injection-sample/domain/model"

type Book interface {
	FindByID(id string) *model.Book
}
