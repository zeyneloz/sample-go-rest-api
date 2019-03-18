package users

import (
	ozzo "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/errors"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/models"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/utils/validation"
	"github.com/zeyneloz/sample-go-rest-api/internal/svc/db"
)

// Validation Schemas are below

type registerSchema struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *registerSchema) Validate() error {
	// Basic validation.
	if err := validation.Validate(s,
		ozzo.Field(&s.Name, ozzo.Required, ozzo.Length(1, 70)),
		ozzo.Field(&s.Password, ozzo.Required, ozzo.Length(3, 100)),
		ozzo.Field(&s.Email, ozzo.Required, is.Email),
	); err != nil {
		return err
	}

	// Check if the email already exists in db.
	var count int
	err := db.Instance.Model(&models.User{}).Where("email = ?", s.Email).Count(&count).Error
	if err != nil {
		return errors.NewInternalError(err, "Error while getting user count.")
	}
	if count > 0 {
		return errors.NewHTTPValidationError(map[string]string{
			"email": "This email already exists",
		})
	}
	return nil
}

type loginSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *loginSchema) Validate() error {
	// Basic validation.
	if err := validation.Validate(s,
		ozzo.Field(&s.Password, ozzo.Required),
		ozzo.Field(&s.Email, ozzo.Required, is.Email),
	); err != nil {
		return err
	}
	return nil
}

type loginResponseSchema struct {
	Token string `json:"token"`
}

type userDetailResponseSchema struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
