package middleware

import (
	"log"
	"net/http"

	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/core"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/errors"
)

// Return internal error.
func internalEror(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(""))
}

// RootHandler must be the first middleware called, since it is the root middleware.
func RootHandler(next core.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := &core.Context{r, nil}
		// Call next middleware.
		err := next.ServeHTTP(w, c)

		// If middleware returns no error, early exit.
		if err == nil {
			return
		}

		log.Printf("[ERROR] %v\n", err)
		rootError := errors.Cause(err)

		// Check if this is a visible error to user.
		visibleError, ok := rootError.(errors.UserVisibleError)

		// If it is not visible to user, just return an Internal Server Error.
		if !ok {
			internalEror(w)
			return
		}

		body, err := visibleError.ResponseBody()
		if err != nil {
			log.Printf("[ERROR] %v\n", err)
			internalEror(w)
			return
		}
		status, messages := visibleError.ResponseHeaders()

		for k, v := range messages {
			w.Header().Set(k, v)
		}
		w.WriteHeader(status)
		w.Write(body)
	})
}
