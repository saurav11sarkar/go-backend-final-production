package user

import (
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/config"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/middleware"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, h *Handler, cfg config.Config) {
	mux.Handle("GET /api/v1/users", middleware.Chain(http.HandlerFunc(h.List), middleware.Auth(cfg), middleware.Role("admin")))
	mux.Handle("GET /api/v1/users/me", middleware.Chain(http.HandlerFunc(h.Me), middleware.Auth(cfg)))
	mux.Handle("PATCH /api/v1/users/me", middleware.Chain(http.HandlerFunc(h.UpdateProfile), middleware.Auth(cfg)))
}
