package app

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/auth"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/category"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/config"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/email"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/link"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/product"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/routes"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/upload"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/user"
	"net/http"
)

func NewHandler(db *pgxpool.Pool, cfg config.Config) (http.Handler, error) {
	mail := email.New(cfg)
	cloud, err := upload.New(cfg)
	if err != nil {
		return nil, err
	}
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo, cfg, mail)
	authHandler := auth.NewHandler(authService)

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	productRepo := product.NewRepository(db)
	productService := product.NewService(productRepo)
	productHandler := product.NewHandler(productService)

	categoryRepo := category.NewRepository(db)
	categoryService := category.NewService(categoryRepo)
	categoryHandler := category.NewHandler(categoryService)

	linkRepo := link.NewRepository(db)
	linkService := link.NewService(linkRepo)
	linkHandler := link.NewHandler(linkService)

	uploadHandler := upload.NewHandler(cloud, cfg.MaxUploadMB)

	deps := routes.Dependencies{
		Auth: authHandler, User: userHandler, Product: productHandler,
		Category: categoryHandler, Link: linkHandler, Upload: uploadHandler,
	}

	return routes.New(deps, cfg), nil
}
