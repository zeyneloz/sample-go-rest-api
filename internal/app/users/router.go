package users

import (
	"github.com/go-chi/chi"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/middleware"
)

// GetRouter for user app routes.
func GetRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/{userID}", middleware.RootHandler(middleware.AuthMiddleware(userDetailHandler)))
	r.Post("/", middleware.RootHandler(registerHandler))
	r.Post("/login", middleware.RootHandler(loginHandler))
	return r
}
