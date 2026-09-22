package middleware

import (
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/config"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/utils"
	"net/http"
	"strings"
)

func Auth(cfg config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			parts := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
			if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
				utils.Error(w, 401, "authorization token is required")
				return
			}
			claims, err := utils.ParseToken(parts[1], cfg.AccessSecret, "access")
			if err != nil {
				utils.Error(w, 401, "invalid or expired access token")
				return
			}
			next.ServeHTTP(w, SetUser(r, claims.UserID, claims.Role))
		})
	}
}
func Role(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, role := CurrentUser(r)
			for _, allowed := range roles {
				if role == allowed {
					next.ServeHTTP(w, r)
					return
				}
			}
			utils.Error(w, 403, "forbidden")
		})
	}
}
