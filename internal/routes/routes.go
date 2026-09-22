package routes

import (
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/auth"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/category"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/config"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/link"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/middleware"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/product"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/upload"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/user"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/utils"
	"net/http"
)

func New(d Dependencies, cfg config.Config) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		utils.JSON(w, 200, "ok", map[string]string{"status": "up"})
	})
	auth.RegisterRoutes(mux, d.Auth)
	user.RegisterRoutes(mux, d.User, cfg)
	product.RegisterRoutes(mux, d.Product, cfg)
	category.RegisterRoutes(mux, d.Category, cfg)
	link.RegisterRoutes(mux, d.Link, cfg)
	upload.RegisterRoutes(mux, d.Upload, cfg)
	return middleware.Chain(mux, middleware.CORS(cfg.CorsOrigin), middleware.RequestID, middleware.Logger, middleware.Recover)
}
