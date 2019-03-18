package books

import (
	"github.com/go-chi/chi"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/middleware"
)

// GetRouter for user app routes.
func GetRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/authors", middleware.RootHandler(middleware.AuthMiddleware(authorListHandler)))
	r.Post("/authors", middleware.RootHandler(middleware.AuthMiddleware(authorCreateHandler)))
	r.Get("/{bookID}", middleware.RootHandler(middleware.AuthMiddleware(bookDetailHandler)))
	r.Post("/", middleware.RootHandler(middleware.AuthMiddleware(bookCreateHandler)))
	return r
}
