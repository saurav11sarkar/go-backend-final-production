package routes

import (
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/auth"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/category"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/link"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/product"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/upload"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/user"
)

type Dependencies struct {
	Auth     *auth.Handler
	User     *user.Handler
	Product  *product.Handler
	Category *category.Handler
	Link     *link.Handler
	Upload   *upload.Handler
}
