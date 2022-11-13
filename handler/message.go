package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

type Message interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type messageHandler struct{}

func NewMessage() Message {
	return &messageHandler{}
}

func (m *messageHandler) Get(w http.ResponseWriter, r *http.Request) {
	message := "There is always light behind the clouds."
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{
		"message": message,
	})
}
