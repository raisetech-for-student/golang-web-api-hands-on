package handler

import (
	"net/http"

	"golang-web-api-hands-on/handler/response"
	"golang-web-api-hands-on/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Book interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type bookHandler struct {
	useCase usecase.Book
}

func NewBookHandler(useCase usecase.Book) Book {
	return &bookHandler{
		useCase: useCase,
	}
}

func (b *bookHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	bookID := chi.URLParam(r, "bookID")

	if bookID == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, response.ErrNotFound)
	}

	result, err := b.useCase.Get(ctx, bookID)
	if err != nil {
		if err.Error() == "could not find the book" {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, response.ErrNotFound)
			return
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, response.ErrSystemError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response.NewBook(result))
}
