package upload

import (
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/config"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/middleware"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, h *Handler, cfg config.Config) {
	mux.Handle("POST /api/v1/uploads/image", middleware.Chain(http.HandlerFunc(h.Image), middleware.Auth(cfg)))
}
