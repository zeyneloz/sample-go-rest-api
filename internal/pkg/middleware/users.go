package middleware

import (
	"log"
	"net/http"

	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/config"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/core"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/errors"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/utils/auth"
	"github.com/zeyneloz/sample-go-rest-api/pkg/crypt"
)

// AuthMiddleware uses `Token` header if exists, to authenticate the user request using JWT.
func AuthMiddleware(next core.Handler) core.Handler {
	return core.Handler(func(w http.ResponseWriter, c *core.Context) error {
		tokenString := c.Header.Get("Token")
		secret := config.GetConfig().SecretKey
		data, err := crypt.ParseJWT(secret, tokenString)
		if err != nil {
			log.Printf("[ERROR] %v\n", err)
			return errors.NewHTTPError(http.StatusUnauthorized, "Unauthorized", map[string]string{})
		}
		user, err := auth.JWTToUser(data)
		if err != nil {
			return err
		}
		c.User = user
		return next.ServeHTTP(w, c)
	})
}
