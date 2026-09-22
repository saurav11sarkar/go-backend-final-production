package product

import (
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/config"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/middleware"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, h *Handler, cfg config.Config) {
	mux.Handle("GET /api/v1/products", middleware.Chain(http.HandlerFunc(h.List), middleware.Auth(cfg)))
	mux.Handle("GET /api/v1/products/{id}", middleware.Chain(http.HandlerFunc(h.Get), middleware.Auth(cfg)))
	mux.Handle("POST /api/v1/products", middleware.Chain(http.HandlerFunc(h.Create), middleware.Auth(cfg), middleware.Role("admin")))
	mux.Handle("PATCH /api/v1/products/{id}", middleware.Chain(http.HandlerFunc(h.Update), middleware.Auth(cfg), middleware.Role("admin")))
	mux.Handle("DELETE /api/v1/products/{id}", middleware.Chain(http.HandlerFunc(h.Delete), middleware.Auth(cfg), middleware.Role("admin")))
}
