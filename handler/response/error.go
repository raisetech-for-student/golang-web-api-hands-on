package response

import (
	"net/http"

	"github.com/go-chi/render"
)

var (
	ErrNotFound    = &ErrResponse{HTTPStatusCode: 404, StatusText: "resource not found"}
	ErrSystemError = &ErrResponse{HTTPStatusCode: 500, StatusText: "internal server error"}
)

type ErrResponse struct {
	HTTPStatusCode int    `json:"-"`       // http response status code
	StatusText     string `json:"message"` // user-level status message
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}
