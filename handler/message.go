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
	message := []string{
		"Change before you have to.",
		"There is always light behind the clouds.",
		"If you can dream it, you can do it.",
		"Love the life you live. Live the life you love.",
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(messages))

	message := messages[randomIndex]
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{
		"message": message,
	})
}
