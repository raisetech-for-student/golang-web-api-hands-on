package repository

import "golang-web-api-hands-on/domain/model"

type Book interface {
	FindByID(id string) *model.Book
}
