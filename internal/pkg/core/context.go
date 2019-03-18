package core

import (
	"net/http"

	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/models"
)

// Context adds extra fields to http.Request.
type Context struct {
	*http.Request
	User *models.User
}

// Handler describes custom middleware function.
type Handler func(http.ResponseWriter, *Context) error

// ServeHTTP calls f(w, r).
func (f Handler) ServeHTTP(w http.ResponseWriter, c *Context) error {
	return f(w, c)
}
