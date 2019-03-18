package app

import (
	"github.com/go-chi/chi"
	"github.com/zeyneloz/sample-go-rest-api/internal/app/books"
	"github.com/zeyneloz/sample-go-rest-api/internal/app/users"
)

// GetRouter returns main router for the application
func GetRouter() chi.Router {
	r := chi.NewRouter()

	r.Mount("/users", users.GetRouter())
	r.Mount("/books", books.GetRouter())

	return r
}
