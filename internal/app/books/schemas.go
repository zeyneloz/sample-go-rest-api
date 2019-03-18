package books

import (
	ozzo "github.com/go-ozzo/ozzo-validation"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/errors"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/models"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/utils/validation"
	"github.com/zeyneloz/sample-go-rest-api/internal/svc/db"
)

// Author Schemas

type authorCreateSchema struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}

func (s *authorCreateSchema) Validate() error {
	// Basic validation.
	if err := validation.Validate(s,
		ozzo.Field(&s.Name, ozzo.Required, ozzo.Length(1, 70)),
		ozzo.Field(&s.LastName, ozzo.Required, ozzo.Length(3, 70)),
	); err != nil {
		return err
	}
	return nil
}

type authorResponseSchema struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}

// Book Schemas

type bookCreateSchema struct {
	Name   string `json:"name"`
	Author int    `json:"author"`
}

func (s *bookCreateSchema) Validate() error {
	// Basic validation.
	if err := validation.Validate(s,
		ozzo.Field(&s.Name, ozzo.Required),
		ozzo.Field(&s.Author, ozzo.Required, ozzo.Min(0)),
	); err != nil {
		return err
	}

	// Check if author really exists.
	var count int
	err := db.Instance.Model(&models.Author{}).Where("id = ?", s.Author).Count(&count).Error
	if err != nil {
		return errors.NewInternalError(err, "Error while getting author count.")
	}
	if count == 0 {
		return errors.NewHTTPValidationError(map[string]string{
			"author": "Author does not exists.",
		})
	}

	return nil
}

type bookCreateResponseSchema struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Author    int    `json:"author"`
	CreatedBy int    `json:"createdBy"`
}

type userDetailSchema struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type bookDetailSchema struct {
	ID        int                  `json:"id"`
	Name      string               `json:"name"`
	Author    authorResponseSchema `json:"author"`
	CreatedBy userDetailSchema     `json:"createdBy"`
}
