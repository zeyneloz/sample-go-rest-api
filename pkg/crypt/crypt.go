package crypt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTData is used as data part od JWT token
type JWTData map[string]interface{}

// GenerateJWT returns new json web token, expires in given minutes.
func GenerateJWT(secret string, expires int, data JWTData) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	for k, v := range data {
		claims[k] = v
	}

	// Set expire time.
	exp := int(time.Minute) * expires
	claims["exp"] = time.Now().Add(time.Duration(exp)).Unix()

	// Sign and generate token.
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWT parses given jwt and returns the claims part.
func ParseJWT(secret string, tokenString string) (JWTData, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Invalid JWT")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Claim parse error")
	}

	data := JWTData{}
	for k, v := range claims {
		data[k] = v
	}

	return data, nil
}
