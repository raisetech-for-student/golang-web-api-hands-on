package main

import (
	"log"
	"net/http"

	"github.com/go-chi/render"

	"golang-web-api-hands-on/handler"
	"golang-web-api-hands-on/infra/dao"
	"golang-web-api-hands-on/usecase"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	messageHandler := handler.NewMessage()
	bookRepo := dao.NewBook()
	bookUseCase := usecase.NewBook(bookRepo)
	bookHandler := handler.NewBookHandler(bookUseCase)

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{
			"message": "hello world",
		})
	})
	r.Get("/message", messageHandler.Get)
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/books/{bookID}", func(r chi.Router) {
			r.Get("/", bookHandler.Get)
		})
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("failed to ListenAndServe: %v", err)
	}
}
