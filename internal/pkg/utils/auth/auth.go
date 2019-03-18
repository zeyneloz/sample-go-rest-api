package auth

import (
	"github.com/jinzhu/gorm"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/errors"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/models"
	"github.com/zeyneloz/sample-go-rest-api/internal/svc/db"
	"github.com/zeyneloz/sample-go-rest-api/pkg/crypt"
)

// UserToJWT converts user model instance to a map to be used as JWT data.
func UserToJWT(user models.User) crypt.JWTData {
	return map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
	}
}

// JWTToUser takes JWT data and returns correspanding user.
func JWTToUser(data crypt.JWTData) (*models.User, error) {
	var user models.User
	err := db.Instance.Where("id = ?", int(data["id"].(float64))).First(&user).Error
	// Check if user really exsists with that email.
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.NewHTTPValidationError(map[string]string{
			"email": "User wtih this email is not found.",
		})
	}
	// If there is any other error, return internal error.
	if err != nil {
		return nil, errors.NewInternalError(err, "GORM error")
	}
	return &user, nil
}
